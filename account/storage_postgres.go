package account

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/lib/pq"
)

type PostgresStorage struct {
	accountTableName, stateTableName string
	db                               *sql.DB
	timeout                          time.Duration
}

func NewPostgresStorage(db *sql.DB, accountTableName, stateTableName string, timeout time.Duration) *PostgresStorage {
	return &PostgresStorage{
		accountTableName: pq.QuoteIdentifier(accountTableName),
		stateTableName:   pq.QuoteIdentifier(stateTableName),
		db:               db,
		timeout:          timeout,
	}
}

func (ps *PostgresStorage) Init() error {
	ctx, cancel := context.WithTimeout(context.Background(), ps.timeout)
	defer cancel()
	_, err := ps.db.ExecContext(ctx, fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		 			id varchar NOT NULL PRIMARY KEY,
		 			data jsonb)`, ps.accountTableName))
	if err != nil {
		return fmt.Errorf("could not create account table %w", err)
	}

	_, err = ps.db.ExecContext(ctx, fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		service_id varchar NOT NULL,
		account_id varchar NOT NULL REFERENCES %s ON UPDATE RESTRICT ON DELETE RESTRICT,
		gap_halving_period varchar NOT NULL,
		power varchar NOT NULL,
		locked_power varchar NOT NULL,
		energy_cap varchar NOT NULL,
		energy varchar NOT NULL,
		energy_calculated_at timestamp NOT NULL,
		PRIMARY KEY (service_id, account_id))`, ps.stateTableName, ps.accountTableName))
	if err != nil {
		return fmt.Errorf("could not create state table %w", err)
	}

	return nil
}

func (ps *PostgresStorage) GetAccount(id string) (Account, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ps.timeout)
	defer cancel()

	row := ps.db.QueryRowContext(ctx, fmt.Sprintf(`SELECT id, data
	FROM %s
	WHERE id = $1`, ps.accountTableName), id)
	return ps.rowToAccount(row)
}

func (ps *PostgresStorage) rowToAccount(row *sql.Row) (Account, error) {
	var id string
	data := []byte{}
	err := row.Scan(&id, &data)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Account{}, ErrNotFound
		}
		return Account{}, err
	}
	claim := OwnershipClaim{}
	err = json.Unmarshal(data, &claim)
	if err != nil {
		return Account{}, err
	}
	return Account{
		ID:             id,
		OwnershipClaim: claim,
	}, nil
}

func (ps *PostgresStorage) DeleteAccount(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), ps.timeout)
	defer cancel()
	tx, err := ps.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.Exec(fmt.Sprintf(`DELETE FROM %s WHERE account_id = $1`, ps.stateTableName), id)
	if err != nil {
		tx.Rollback()
		return err
	}

	res, err := tx.Exec(fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, ps.accountTableName), id)
	if err != nil {
		tx.Rollback()
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	if affected == 0 {
		tx.Rollback()
		return ErrNotFound
	}
	return tx.Commit()
}

func (ps *PostgresStorage) GetAccountState(serviceID string, accountID string) (AccountState, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ps.timeout)
	defer cancel()

	query := fmt.Sprintf(`SELECT
			service_id,
            account_id,
            gap_halving_period,
            power,
            locked_power,
            energy_cap,
            energy,
            energy_calculated_at
          FROM %s
          WHERE service_id = $1
            AND account_id = $2
        `, ps.stateTableName)

	row := ps.db.QueryRowContext(ctx, query, serviceID, accountID)
	return ps.rowToAccountState(row)
}

func (ps *PostgresStorage) DeleteAccountState(serviceID string, accountID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), ps.timeout)
	defer cancel()

	res, err := ps.db.ExecContext(ctx, fmt.Sprintf(`DELETE FROM %s WHERE account_id = $1 and service_id =$2`, ps.stateTableName), accountID, serviceID)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return ErrNotFound
	}
	return nil
}

func (ps *PostgresStorage) SaveAccount(account Account) (Account, error) {
	marshaled, err := json.Marshal(account.OwnershipClaim)
	if err != nil {
		return Account{}, err
	}

	query := fmt.Sprintf(`INSERT INTO %s (
		id,
		data
	  ) VALUES (
		$1,$2
	  )
	  ON CONFLICT (id)
	  DO UPDATE SET data = EXCLUDED.data
	  RETURNING *`, ps.accountTableName)

	ctx, cancel := context.WithTimeout(context.Background(), ps.timeout)
	defer cancel()

	row := ps.db.QueryRowContext(ctx, query, account.ID, marshaled)
	if err != nil {
		return Account{}, err
	}

	return ps.rowToAccount(row)
}

func (ps *PostgresStorage) InitAccountState(accountState AccountState) (AccountState, error) {
	accountState.FillDefaults()

	ctx, cancel := context.WithTimeout(context.Background(), ps.timeout)
	defer cancel()
	query := fmt.Sprintf(`INSERT INTO %s (
			  account_id,
			  service_id,
			  gap_halving_period,
			  power,
			  locked_power,
			  energy_cap,
			  energy,
			  energy_calculated_at
			) VALUES (
			  $1,
			  $2,
			  $3,
			  $4,
			  $5,
			  $6,
			  $7,
			  to_timestamp($8)
			)
			RETURNING *
		  `, ps.stateTableName)

	row := ps.db.QueryRowContext(ctx,
		query,
		accountState.AccountID,
		accountState.ServiceID,
		fmt.Sprint(accountState.GapHalvingPeriod),
		accountState.Power.String(),
		accountState.LockedPower.String(),
		accountState.EnergyCap.String(),
		accountState.Energy.String(),
		accountState.EnergyCalculatedAt,
	)

	return ps.rowToAccountState(row)
}

func (ps *PostgresStorage) rowToAccountState(row *sql.Row) (AccountState, error) {
	as := AccountState{}
	var gapHalvingPeriod, power, lockedPower, energyCap, energy string
	calculatedAt := time.Time{}
	err := row.Scan(
		&as.ServiceID,
		&as.AccountID,
		&gapHalvingPeriod,
		&power,
		&lockedPower,
		&energyCap,
		&energy,
		&calculatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) || strings.Contains(err.Error(), "violates foreign key constraint") {
			return AccountState{}, ErrNotFound
		} else if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return AccountState{}, ErrAlreadyInitialized
		}
		return AccountState{}, err
	}

	gapHalvingP, err := strconv.Atoi(gapHalvingPeriod)
	if err != nil {
		return AccountState{}, err
	}
	as.GapHalvingPeriod = int64(gapHalvingP)

	powerBig, ok := big.NewInt(0).SetString(power, 10)
	if !ok {
		return AccountState{}, fmt.Errorf("could not parse power from %v", power)
	}
	as.Power = powerBig

	lockedPowerBig, ok := big.NewInt(0).SetString(lockedPower, 10)
	if !ok {
		return AccountState{}, fmt.Errorf("could not parse locked power from %v", power)
	}
	as.LockedPower = lockedPowerBig

	energyCapBig, ok := big.NewInt(0).SetString(energyCap, 10)
	if !ok {
		return AccountState{}, fmt.Errorf("could not parse energy cap from %v", power)
	}
	as.EnergyCap = energyCapBig

	energyBig, ok := big.NewInt(0).SetString(energy, 10)
	if !ok {
		return AccountState{}, fmt.Errorf("could not parse energy from %v", power)
	}
	as.Energy = energyBig

	as.EnergyCalculatedAt = calculatedAt.Unix()
	return as, nil
}

func (ps *PostgresStorage) ChangeAccountState(previous, new AccountState) (AccountState, error) {
	new.FillDefaults()

	ctx, cancel := context.WithTimeout(context.Background(), ps.timeout)
	defer cancel()
	query := fmt.Sprintf(`UPDATE %s
            SET
              power = $1,
              locked_power = $2,
              energy = $3,
              energy_calculated_at = to_timestamp($4)
          WHERE
            account_id = $5
            AND service_id = $6
            AND gap_halving_period = $7
            AND power = $8
            AND locked_power = $9
            AND energy_cap = $10
            AND energy = $11
            AND energy_calculated_at = to_timestamp($12)
          RETURNING *`,
		ps.stateTableName)

	row := ps.db.QueryRowContext(ctx,
		query,
		new.Power.String(),
		new.LockedPower.String(),
		new.Energy.String(),
		new.EnergyCalculatedAt,
		previous.AccountID,
		previous.ServiceID,
		fmt.Sprint(previous.GapHalvingPeriod),
		previous.Power.String(),
		previous.LockedPower.String(),
		previous.EnergyCap.String(),
		previous.Energy.String(),
		previous.EnergyCalculatedAt,
	)
	return ps.rowToAccountState(row)
}
