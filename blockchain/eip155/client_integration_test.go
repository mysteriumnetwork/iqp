package eip155

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mysteriumnetwork/iqp/bindings"
	"github.com/stretchr/testify/assert"
)

func TestEip155ClientIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping eip 155 client testing in short mode")
	}

	var ethClient *ethclient.Client
	var eip155Signer func(address common.Address, tx *types.Transaction) (*types.Transaction, error)
	var deployerAddress = common.HexToAddress("0x354Bd098B4eF8c9E70B7F21BE2d455DF559705d7")
	var erc20Address common.Address
	var powertTokenAddress common.Address
	var interestTokenAddress common.Address
	var borrowTokenAddress common.Address
	var enterpriseAddress common.Address
	var converterAddress common.Address
	var enterpriseFactoryAddress common.Address
	var ownerPK *ecdsa.PrivateKey
	var chainID *big.Int
	oneToken, _ := big.NewInt(0).SetString("1000000000000000000", 10)

	var getTransactOpts = func() (*bind.TransactOpts, func()) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		return &bind.TransactOpts{
			Context: ctx,
			Signer:  eip155Signer,
			From:    deployerAddress,
		}, cancel
	}

	t.Run("setup", func(t *testing.T) {
		t.Run("connect to local ganache", func(t *testing.T) {
			assert.Eventually(t, func() bool {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()
				ethc, err := ethclient.DialContext(ctx, "ws://localhost:8548")
				if err != nil {
					return false
				}
				ethClient = ethc
				return true
			}, time.Second*10, time.Millisecond*300)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			chid, err := ethClient.NetworkID(ctx)
			assert.NoError(t, err)
			chainID = chid
		})
		t.Run("setup private key", func(t *testing.T) {
			ownerPrivateKey := "45bb96530f3d1972fdcd2005c1987a371d0b6d378b77561c6beeaca27498f46b"
			privateKey, err := crypto.HexToECDSA(ownerPrivateKey)
			assert.NoError(t, err)
			ownerPK = privateKey

			chainID, err := ethClient.NetworkID(context.Background())
			assert.NoError(t, err)

			eip155Signer = func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
				return types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
			}
		})
		t.Run("deploy smart contracts", func(t *testing.T) {
			t.Run("deploy erc20", func(t *testing.T) {
				opts, cancel := getTransactOpts()
				defer cancel()
				addr, tx, _, err := bindings.DeployERC20Mock(opts, ethClient, "Testing", "TST", 18, big.NewInt(0).Mul(big.NewInt(1000000000), oneToken))
				assert.NoError(t, err)
				assertTxSuccess(t, tx, ethClient)

				erc20Address = addr
			})
			t.Run("deploy ExpMathMock", func(t *testing.T) {
				opts, cancel := getTransactOpts()
				defer cancel()
				_, tx, _, err := bindings.DeployExpMathMock(opts, ethClient)
				assert.NoError(t, err)
				assertTxSuccess(t, tx, ethClient)
			})
			t.Run("deploy power token", func(t *testing.T) {
				opts, cancel := getTransactOpts()
				defer cancel()
				addr, tx, _, err := bindings.DeployPowerToken(opts, ethClient)
				assert.NoError(t, err)
				assertTxSuccess(t, tx, ethClient)

				powertTokenAddress = addr
			})
			t.Run("deploy interest token", func(t *testing.T) {
				opts, cancel := getTransactOpts()
				defer cancel()
				addr, tx, _, err := bindings.DeployStakeToken(opts, ethClient)
				assert.NoError(t, err)
				assertTxSuccess(t, tx, ethClient)

				interestTokenAddress = addr
			})
			t.Run("deploy interest token", func(t *testing.T) {
				opts, cancel := getTransactOpts()
				defer cancel()
				addr, tx, _, err := bindings.DeployRentalToken(opts, ethClient)
				assert.NoError(t, err)
				assertTxSuccess(t, tx, ethClient)

				borrowTokenAddress = addr
			})
			t.Run("deploy enterprise", func(t *testing.T) {
				opts, cancel := getTransactOpts()
				defer cancel()
				addr, tx, _, err := bindings.DeployEnterprise(opts, ethClient)
				assert.NoError(t, err)
				assertTxSuccess(t, tx, ethClient)

				enterpriseAddress = addr
			})
			t.Run("deploy converter", func(t *testing.T) {
				opts, cancel := getTransactOpts()
				defer cancel()
				addr, tx, _, err := bindings.DeployDefaultConverter(opts, ethClient)
				assert.NoError(t, err)
				assertTxSuccess(t, tx, ethClient)

				converterAddress = addr
			})
			t.Run("deploy enterprise factory", func(t *testing.T) {
				opts, cancel := getTransactOpts()
				defer cancel()
				addr, tx, _, err := bindings.DeployEnterpriseFactory(opts, ethClient, enterpriseAddress, powertTokenAddress, interestTokenAddress, borrowTokenAddress)
				assert.NoError(t, err)
				assertTxSuccess(t, tx, ethClient)

				enterpriseFactoryAddress = addr
			})
		})
	})

	t.Run("test client", func(t *testing.T) {
		baseEnterpriseParams := EnterpriseParams{
			Name:                  "Test Enterprise",
			BaseUri:               "https://iq.space",
			GCFeePercent:          2 * 100,
			LiquidityTokenAddress: erc20Address,
			Converter:             converterAddress,
		}
		client := NewClient(ethClient, time.Second, deployerAddress, eip155Signer)

		var deployedEnterprise common.Address
		t.Run("deploys enterprise", func(t *testing.T) {
			tx, err := client.DeployEnterprise(enterpriseFactoryAddress, baseEnterpriseParams)
			assert.NoError(t, err)
			assertTxSuccess(t, tx, ethClient)

			filterer, err := bindings.NewEnterpriseFactoryFilterer(enterpriseFactoryAddress, ethClient)
			assert.NoError(t, err)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			iter, err := filterer.FilterEnterpriseDeployed(&bind.FilterOpts{
				Context: ctx,
				Start:   0,
			}, []common.Address{deployerAddress}, []common.Address{baseEnterpriseParams.LiquidityTokenAddress})
			assert.NoError(t, err)

			for iter.Next() {
				deployedEnterprise = iter.Event.Deployed
			}
		})
		t.Run("returns chain ID", func(t *testing.T) {
			chid, err := client.GetChainID()
			assert.NoError(t, err)
			assert.Equal(t, int64(31337), chid)
		})

		t.Run("retrieves an arbitrary token balance", func(t *testing.T) {
			b, err := client.GetTokenBalance(baseEnterpriseParams.LiquidityTokenAddress, deployerAddress)
			assert.NoError(t, err)

			assert.Equal(t, big.NewInt(0).Mul(big.NewInt(1000000000), oneToken), b)

			t.Run("retrieves 0 if no balance exists", func(t *testing.T) {
				b, err := client.GetTokenBalance(baseEnterpriseParams.LiquidityTokenAddress, common.HexToAddress("0x17CdaD42F88fcecF77F53467B3c15613e8063adD"))
				assert.NoError(t, err)
				assert.Equal(t, int64(0), b.Int64())
			})
		})

		t.Run("retrieves enterprise data", func(t *testing.T) {
			info, err := client.GetEnterpriseInfo(deployedEnterprise)
			assert.NoError(t, err)
			assert.Equal(t, baseEnterpriseParams.Name, info.Name)
			assert.Equal(t, baseEnterpriseParams.BaseUri, info.BaseUri)
			assert.Equal(t, uint32(24*3600*7), info.StreamingReserveHalvingPeriod)
			assert.Equal(t, uint32(12*3600), info.RenterOnlyReturnPeriod)
			assert.Equal(t, uint32(24*3600), info.EnterpriseOnlyCollectionPeriod)
			assert.Equal(t, baseEnterpriseParams.GCFeePercent, info.GCFeePercent)
			assert.Equal(t, uint32(0), info.StreamingReserveUpdated)
			assert.Equal(t, big.NewInt(0).Int64(), info.TotalShares.Int64())
			assert.Equal(t, big.NewInt(0).Int64(), info.FixedReserve.Int64())
			assert.Equal(t, big.NewInt(0).Int64(), info.UsedReserve.Int64())
			assert.Equal(t, big.NewInt(0).Int64(), info.StreamingReserve.Int64())
			assert.Equal(t, big.NewInt(0).Int64(), info.StreamingReserveTarget.Int64())
		})

		t.Run("retrieves enterprise liquidity token metadata", func(t *testing.T) {
			metadata, err := client.GetLiquidityTokenMetadata(deployedEnterprise)
			assert.NoError(t, err)

			assert.EqualValues(t, ERC20Metadata{
				Address:  baseEnterpriseParams.LiquidityTokenAddress,
				Name:     "Testing",
				Symbol:   "TST",
				Decimals: 18,
			}, metadata)
		})

		baseServiceParams := ServiceParams{
			Name:              "Test Service",
			Symbol:            "IQPT",
			GapHalvingPeriod:  86400,
			BaseRate:          big.NewInt(6405119470038),
			BaseToken:         baseEnterpriseParams.LiquidityTokenAddress,
			ServiceFeePercent: 3 * 100,
			MinLoanDuration:   12 * 3600,
			MaxLoanDuration:   60 * 3600 * 24,
			MinGCFee:          oneToken,
			SwappingEnabled:   true,
		}

		t.Run("allows to set collector address", func(t *testing.T) {
			collector := common.HexToAddress("0x17CdaD42F88fcecF77F53467B3c15613e8063adD")
			tx, err := client.SetEnterpriseCollectorAddress(deployedEnterprise, collector)
			assert.NoError(t, err)

			assertTxSuccess(t, tx, ethClient)

			fetchedCollector, err := client.GetEnterpriseCollectorAddress(deployedEnterprise)
			assert.NoError(t, err)

			assert.Equal(t, collector, fetchedCollector)
		})

		t.Run("allows to set an enterprise vault address", func(t *testing.T) {
			vault := common.HexToAddress("0x1E60bE47921E15FBc9e9246daC71F771d4b78a6c")
			tx, err := client.SetEnterpriseWalletAddress(deployedEnterprise, vault)

			assert.NoError(t, err)

			assertTxSuccess(t, tx, ethClient)

			fetchedVault, err := client.GetEnterpriseWalletAddress(deployedEnterprise)
			assert.NoError(t, err)

			assert.Equal(t, vault, fetchedVault)
		})

		t.Run("allows to set a converter address", func(t *testing.T) {
			newAddress := common.HexToAddress("0x312D566FABE323BA574fFf724d29c08F46b746b0")
			tx, err := client.SetConverterAddress(deployedEnterprise, newAddress)

			assert.NoError(t, err)

			assertTxSuccess(t, tx, ethClient)

			fetchedAddress, err := client.GetConverterAddress(deployedEnterprise)
			assert.NoError(t, err)

			assert.Equal(t, newAddress, fetchedAddress)
		})

		t.Run("allows to set a bonding curve", func(t *testing.T) {
			tx, err := client.SetBondingCurve(deployedEnterprise, big.NewInt(10), big.NewInt(5))
			assert.NoError(t, err)

			assertTxSuccess(t, tx, ethClient)

			curve, err := client.GetBondingCurve(deployedEnterprise)
			assert.NoError(t, err)

			assert.Equal(t, curve.Pole, big.NewInt(10))
			assert.Equal(t, curve.Slope, big.NewInt(5))
		})

		t.Run("allows to set a borrower return grace period", func(t *testing.T) {
			var period uint32 = 3 * 3600
			tx, err := client.SetRenterOnlyReturnPeriod(deployedEnterprise, period)
			assert.NoError(t, err)

			assertTxSuccess(t, tx, ethClient)

			p, err := client.GetRenterOnlyReturnPeriod(deployedEnterprise)
			assert.NoError(t, err)
			assert.Equal(t, period, p)
		})

		t.Run("allows to set a collector grace period", func(t *testing.T) {
			var period uint32 = 14 * 3600
			tx, err := client.SetEnterpriseLoanCollectionPeriod(deployedEnterprise, period)
			assert.NoError(t, err)

			assertTxSuccess(t, tx, ethClient)

			p, err := client.GetEnterpriseLoanCollectionPeriod(deployedEnterprise)
			assert.NoError(t, err)
			assert.Equal(t, period, p)
		})

		t.Run("allows to set enterprise base URI", func(t *testing.T) {
			uri := "https://iq.space"
			tx, err := client.SetBaseUri(deployedEnterprise, uri)
			assert.NoError(t, err)

			assertTxSuccess(t, tx, ethClient)

			p, err := client.GetBaseUri(deployedEnterprise)
			assert.NoError(t, err)
			assert.Equal(t, uri, p)
		})

		t.Run("allows to set interest gap halving period", func(t *testing.T) {
			var period uint32 = 6 * 3600
			tx, err := client.SetInterestReserveHalvingPeriod(deployedEnterprise, period)
			assert.NoError(t, err)

			assertTxSuccess(t, tx, ethClient)

			p, err := client.GetInterestReserveHalvingPeriod(deployedEnterprise)
			assert.NoError(t, err)
			assert.Equal(t, period, p)
		})

		t.Run("allows to set GC fee percent", func(t *testing.T) {
			var percent uint16 = 200
			tx, err := client.SetGCFeePercent(deployedEnterprise, percent)
			assert.NoError(t, err)

			assertTxSuccess(t, tx, ethClient)

			p, err := client.GetGCFeePercent(deployedEnterprise)
			assert.NoError(t, err)
			assert.Equal(t, percent, p)
		})

		t.Run("when enterprise has registered services", func(t *testing.T) {
			var service1Params ServiceParams
			var service1Address common.Address
			var service2Address common.Address
			t.Run("registers two services", func(t *testing.T) {
				t.Run("registers service one", func(t *testing.T) {
					params := baseServiceParams
					params.Name = "Test Service 1"
					params.Symbol = "IQPT1"
					tx, err := client.RegisterService(deployedEnterprise, params)
					assert.NoError(t, err)
					assertTxSuccess(t, tx, ethClient)

					service1Params = params

					addr, err := getPowerToken(t, tx, ethClient, deployedEnterprise)
					assert.NoError(t, err)

					service1Address = addr
				})
				t.Run("registers service two", func(t *testing.T) {
					params := baseServiceParams
					params.Name = "Test Service 2"
					params.Symbol = "IQPT2"
					tx, err := client.RegisterService(deployedEnterprise, params)
					assert.NoError(t, err)
					assertTxSuccess(t, tx, ethClient)

					addr, err := getPowerToken(t, tx, ethClient, deployedEnterprise)
					assert.NoError(t, err)

					service2Address = addr
				})

				t.Run("lists enterprise services", func(t *testing.T) {
					svcs, err := client.GetServices(deployedEnterprise)
					assert.NoError(t, err)

					assert.Len(t, svcs, 2)
					assert.Equal(t, service1Address, svcs[0])
					assert.Equal(t, service2Address, svcs[1])
				})

				t.Run("retrieves service data", func(t *testing.T) {
					info, err := client.GetServiceInfo(service1Address)
					assert.NoError(t, err)
					assert.Equal(t, ServiceInfo{
						Name:                   service1Params.Name,
						Symbol:                 fmt.Sprintf("%v %v", "TST", service1Params.Symbol),
						BaseRate:               service1Params.BaseRate,
						MinGCFee:               service1Params.MinGCFee,
						EnergyGapHalvingPeriod: service1Params.GapHalvingPeriod,
						Index:                  0,
						BaseToken:              service1Params.BaseToken,
						MinRentalPeriod:        service1Params.MinLoanDuration,
						MaxRentalPeriod:        service1Params.MaxLoanDuration,
						ServiceFeePercent:      service1Params.ServiceFeePercent,
						SwappingEnabled:        service1Params.SwappingEnabled,
					}, info)
				})

				t.Run("retrieves account state for specific service", func(t *testing.T) {
					accountAddress := common.HexToAddress("0x6F5E1207671091337abA2F2cda040EB16E757092")
					state, err := client.GetAccountState(service1Address, accountAddress)
					assert.NoError(t, err)

					assert.Equal(t, accountAddress, state.AccountAddress)
					assert.Equal(t, service1Address, state.ServiceAddress)
					assert.Equal(t, big.NewInt(0).Int64(), state.Balance.Int64())
					assert.Equal(t, big.NewInt(0).Int64(), state.Energy.Int64())
					assert.Equal(t, big.NewInt(0).Int64(), state.LockedBalance.Int64())
					assert.Equal(t, int64(0), state.Timestamp)
				})

				t.Run("retrieves the service gap halving period", func(t *testing.T) {
					state, err := client.GetEnergyGapHalvingPeriod(service1Address)
					assert.NoError(t, err)

					assert.Equal(t, service1Params.GapHalvingPeriod, uint32(state))
				})

				t.Run("retrieves the service index", func(t *testing.T) {
					idx, err := client.GetServiceIndex(service1Address)
					assert.NoError(t, err)

					assert.Equal(t, uint16(0), idx)
				})

				t.Run("retrieves the service fee percent", func(t *testing.T) {
					prc, err := client.GetServiceFeePercent(service1Address)
					assert.NoError(t, err)

					assert.Equal(t, baseServiceParams.ServiceFeePercent, uint16(prc))
				})

				t.Run("retrieves the service perpetual token flag", func(t *testing.T) {
					allows, err := client.GetAllowsSwapping(service1Address)
					assert.NoError(t, err)

					assert.Equal(t, baseServiceParams.SwappingEnabled, allows)
				})

				t.Run("allows to set service base rate", func(t *testing.T) {
					baseRate := big.NewInt(10)
					baseToken := common.HexToAddress("0xBC7024d93ae7db4a60E0720c09127A49477aea80")
					minGCFee := big.NewInt(2)
					tx, err := client.SetBaseRate(service1Address, baseToken, baseRate, minGCFee)
					assert.NoError(t, err)

					assertTxSuccess(t, tx, ethClient)

					rate, err := client.GetBaseRate(service1Address)
					assert.NoError(t, err)
					assert.Equal(t, baseRate, rate)

					addr, err := client.GetBaseTokenAddress(service1Address)
					assert.NoError(t, err)
					assert.Equal(t, baseToken, addr)

					minFee, err := client.GetMinGCFee(service1Address)
					assert.NoError(t, err)
					assert.Equal(t, minGCFee, minFee)
				})

				t.Run("allows to set service fee percent", func(t *testing.T) {
					var feePercent uint16 = 12 * 100
					tx, err := client.SetServiceFeePercent(service1Address, feePercent)
					assert.NoError(t, err)

					assertTxSuccess(t, tx, ethClient)

					newPercent, err := client.GetServiceFeePercent(service1Address)
					assert.NoError(t, err)

					assert.Equal(t, feePercent, uint16(newPercent))
				})

				t.Run("allows to set service loan duration limits", func(t *testing.T) {
					var min uint32 = 8 * 3600
					var max uint32 = 10 * 3600 * 24
					tx, err := client.SetRentalPeriodLimits(service1Address, min, max)
					assert.NoError(t, err)
					assertTxSuccess(t, tx, ethClient)

					newMin, err := client.GetMinRentalDuration(service1Address)
					assert.NoError(t, err)

					newMax, err := client.GetMaxRentalDuration(service1Address)
					assert.NoError(t, err)

					assert.Equal(t, min, uint32(newMin))
					assert.Equal(t, max, uint32(newMax))
				})

				t.Run("allows to approve liquidity tokens to a service", func(t *testing.T) {
					tx, err := client.ApproveLiquidityTokensToService(service1Address, oneToken)
					assert.NoError(t, err)
					assertTxSuccess(t, tx, ethClient)

					caller, err := bindings.NewERC20Caller(baseEnterpriseParams.LiquidityTokenAddress, ethClient)
					assert.NoError(t, err)

					ctx, cancel := context.WithTimeout(context.Background(), time.Second)
					defer cancel()
					allowance, err := caller.Allowance(&bind.CallOpts{Context: ctx}, deployerAddress, service1Address)
					assert.NoError(t, err)

					assert.Equal(t, oneToken, allowance)
				})

				t.Run("retrieves liquidity token allowance to service", func(t *testing.T) {
					allowance, err := client.GetLiquidityTokenServiceAllowance(service1Address, deployerAddress)
					assert.NoError(t, err)
					assert.Equal(t, oneToken, allowance)
				})

				t.Run("allows to wrap liquidity tokens to get power tokens", func(t *testing.T) {
					amount := big.NewInt(0).Mul(oneToken, big.NewInt(15))
					caller, err := bindings.NewERC20Caller(baseEnterpriseParams.LiquidityTokenAddress, ethClient)
					assert.NoError(t, err)

					ctx, cancel := context.WithTimeout(context.Background(), time.Second)
					defer cancel()
					liquidityTokenBalanceBefore, err := caller.BalanceOf(&bind.CallOpts{Context: ctx}, deployerAddress)
					assert.NoError(t, err)

					powerTokenBalanceBefore, err := client.GetTokenBalance(service1Address, deployerAddress)
					assert.NoError(t, err)

					tx, err := client.ApproveLiquidityTokensToService(service1Address, amount)
					assert.NoError(t, err)
					assertTxSuccess(t, tx, ethClient)

					tx, err = client.SwapIn(service1Address, amount)
					assert.NoError(t, err)
					assertTxSuccess(t, tx, ethClient)

					liquidityTokenBalanceAfter, err := caller.BalanceOf(&bind.CallOpts{Context: ctx}, deployerAddress)
					assert.NoError(t, err)
					powerTokenBalanceAfter, err := client.GetTokenBalance(service1Address, deployerAddress)
					assert.NoError(t, err)

					assert.Equal(t, big.NewInt(0).Int64(), powerTokenBalanceBefore.Int64())
					assert.Equal(t, amount, powerTokenBalanceAfter)
					assert.Equal(t, liquidityTokenBalanceAfter, new(big.Int).Sub(liquidityTokenBalanceBefore, amount))
				})

				destAccountAddress := client.signerAddress

				t.Run("when there are wrapped tokens", func(t *testing.T) {
					t.Run("retrieves power token balance for account", func(t *testing.T) {
						balance, err := client.GetPowerTokenBalance(service1Address, destAccountAddress)
						assert.NoError(t, err)
						assert.Equal(t, big.NewInt(0).Mul(oneToken, big.NewInt(15)), balance)

						balance, err = client.GetPowerTokenAvailableBalance(service1Address, destAccountAddress)
						assert.NoError(t, err)
						assert.Equal(t, big.NewInt(0).Mul(oneToken, big.NewInt(15)), balance)
					})
					t.Run("allows to unwrap power tokens to get liquidity tokens", func(t *testing.T) {
						caller, err := bindings.NewERC20Caller(baseEnterpriseParams.LiquidityTokenAddress, ethClient)
						assert.NoError(t, err)

						ctx, cancel := context.WithTimeout(context.Background(), time.Second)
						defer cancel()
						liquidityTokenBalanceBefore, err := caller.BalanceOf(&bind.CallOpts{Context: ctx}, deployerAddress)
						assert.NoError(t, err)

						powerTokenBalanceBefore, err := client.GetTokenBalance(service1Address, deployerAddress)
						assert.NoError(t, err)

						tx, err := client.SwapOut(service1Address, powerTokenBalanceBefore)
						assert.NoError(t, err)

						assertTxSuccess(t, tx, ethClient)
						liquidityTokenBalanceAfter, err := caller.BalanceOf(&bind.CallOpts{Context: ctx}, deployerAddress)
						assert.NoError(t, err)

						powerTokenBalanceAfter, err := client.GetTokenBalance(service1Address, deployerAddress)
						assert.NoError(t, err)

						assert.Equal(t, int64(0), powerTokenBalanceAfter.Int64())
						assert.Equal(t, new(big.Int).Add(liquidityTokenBalanceBefore, powerTokenBalanceBefore), liquidityTokenBalanceAfter)
					})
				})
				t.Run("When there is a liquidity provider", func(t *testing.T) {
					var lpClient *Client
					var liquidityProviderAddress common.Address
					t.Run("setup", func(t *testing.T) {
						t.Run("setup client", func(t *testing.T) {
							ownerPrivateKey := "45bb96530f3d1972fdcd2005c1987a371d0b6d378b77561c6beeaca27498f46c"
							lpClient = privkeyToClient(t, ownerPrivateKey, ethClient, chainID)
							liquidityProviderAddress = lpClient.signerAddress
						})
						t.Run("transfer eth to liquidity provider", func(t *testing.T) {
							ctx, cancel := context.WithTimeout(context.Background(), time.Second)
							defer cancel()

							nonce, err := ethClient.PendingNonceAt(ctx, deployerAddress)
							assert.NoError(t, err)

							tx := types.NewTransaction(nonce, liquidityProviderAddress, oneToken, 100000, nil, nil)

							signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), ownerPK)
							assert.NoError(t, err)

							err = ethClient.SendTransaction(ctx, signedTx)
							assert.NoError(t, err)
						})
						t.Run("transfer erc20 to liquidity provider", func(t *testing.T) {
							opts, cancel := getTransactOpts()
							defer cancel()
							transactor, err := bindings.NewERC20Transactor(baseEnterpriseParams.LiquidityTokenAddress, ethClient)
							assert.NoError(t, err)

							tx, err := transactor.Transfer(opts, liquidityProviderAddress, new(big.Int).Mul(oneToken, big.NewInt(10000)))
							assert.NoError(t, err)
							assertTxSuccess(t, tx, ethClient)
						})
					})

					t.Run("allows to approve liquidity tokens to enterprise", func(t *testing.T) {
						tx, err := lpClient.ApproveLiquidityTokensToEnterprise(deployedEnterprise, oneToken)
						assert.NoError(t, err)
						assertTxSuccess(t, tx, ethClient)

						caller, err := bindings.NewERC20Caller(baseEnterpriseParams.LiquidityTokenAddress, ethClient)
						assert.NoError(t, err)

						ctx, cancel := context.WithTimeout(context.Background(), time.Second)
						defer cancel()

						res, err := caller.Allowance(&bind.CallOpts{Context: ctx}, liquidityProviderAddress, deployedEnterprise)
						assert.NoError(t, err)

						assert.Equal(t, oneToken, res)
					})
					t.Run("retrieves liquidity token allowance to enterprise", func(t *testing.T) {
						res, err := lpClient.GetLiquidityTokenEnterpriseAllowance(deployedEnterprise, liquidityProviderAddress)
						assert.NoError(t, err)

						assert.Equal(t, oneToken, res)
					})
					t.Run("allows to add liquidity", func(t *testing.T) {
						tx, err := lpClient.Stake(deployedEnterprise, oneToken)
						assert.NoError(t, err)
						assertTxSuccess(t, tx, ethClient)
					})

					t.Run("when liquidity is added", func(t *testing.T) {
						t.Run("setup", func(t *testing.T) {
							tx, err := lpClient.ApproveLiquidityTokensToEnterprise(deployedEnterprise, new(big.Int).Mul(oneToken, big.NewInt(2000)))
							assert.NoError(t, err)
							assertTxSuccess(t, tx, ethClient)

							tx, err = lpClient.Stake(deployedEnterprise, new(big.Int).Mul(oneToken, big.NewInt(1000)))
							assert.NoError(t, err)
							assertTxSuccess(t, tx, ethClient)

							tx, err = lpClient.Stake(deployedEnterprise, new(big.Int).Mul(oneToken, big.NewInt(799)))
							assert.NoError(t, err)
							assertTxSuccess(t, tx, ethClient)
						})
						t.Run("retrieves a list of interest tokens", func(t *testing.T) {
							ids, err := lpClient.GetPaymentTokenIDs(deployedEnterprise, liquidityProviderAddress, baseEnterpriseParams.LiquidityTokenAddress)
							assert.NoError(t, err)
							assert.Len(t, ids, 3)
						})
						t.Run("retrieves liquidity info by interest token ID", func(t *testing.T) {
							ids, err := lpClient.GetPaymentTokenIDs(deployedEnterprise, liquidityProviderAddress, baseEnterpriseParams.LiquidityTokenAddress)
							assert.NoError(t, err)
							info, err := lpClient.GetStake(deployedEnterprise, ids[1])
							assert.NoError(t, err)
							assert.Equal(t, ids[1], info.TokenID)
							assert.Equal(t, big.NewInt(0).Mul(oneToken, big.NewInt(1000)), info.Shares)
							assert.Equal(t, big.NewInt(0).Mul(oneToken, big.NewInt(1000)), info.Amount)
						})
						t.Run("allows to increase liquidity", func(t *testing.T) {
							ids, err := lpClient.GetPaymentTokenIDs(deployedEnterprise, liquidityProviderAddress, baseEnterpriseParams.LiquidityTokenAddress)
							assert.NoError(t, err)
							tx, err := lpClient.IncreaseStake(deployedEnterprise, ids[1], new(big.Int).Mul(oneToken, big.NewInt(200)))
							assert.NoError(t, err)
							assertTxSuccess(t, tx, ethClient)

							info, err := lpClient.GetStake(deployedEnterprise, ids[1])
							assert.NoError(t, err)
							assert.Equal(t, ids[1], info.TokenID)
							assert.Equal(t, big.NewInt(0).Mul(oneToken, big.NewInt(1200)), info.Shares)
							assert.Equal(t, big.NewInt(0).Mul(oneToken, big.NewInt(1200)), info.Amount)
						})
						t.Run("allows to decrease liquidity", func(t *testing.T) {
							ids, err := lpClient.GetPaymentTokenIDs(deployedEnterprise, liquidityProviderAddress, baseEnterpriseParams.LiquidityTokenAddress)
							assert.NoError(t, err)
							tx, err := lpClient.DecreaseStake(deployedEnterprise, ids[1], new(big.Int).Mul(oneToken, big.NewInt(200)))
							assert.NoError(t, err)
							assertTxSuccess(t, tx, ethClient)

							info, err := lpClient.GetStake(deployedEnterprise, ids[1])
							assert.NoError(t, err)
							assert.Equal(t, ids[1], info.TokenID)
							assert.Equal(t, big.NewInt(0).Mul(oneToken, big.NewInt(1000)), info.Shares)
							assert.Equal(t, big.NewInt(0).Mul(oneToken, big.NewInt(1000)), info.Amount)
						})
						t.Run("allows to remove liquidity", func(t *testing.T) {
							ids, err := lpClient.GetPaymentTokenIDs(deployedEnterprise, liquidityProviderAddress, baseEnterpriseParams.LiquidityTokenAddress)
							assert.NoError(t, err)
							tx, err := lpClient.Unstake(deployedEnterprise, ids[2])
							assert.NoError(t, err)
							assertTxSuccess(t, tx, ethClient)
							removedID := ids[2]

							ids, err = lpClient.GetPaymentTokenIDs(deployedEnterprise, liquidityProviderAddress, baseEnterpriseParams.LiquidityTokenAddress)
							assert.NoError(t, err)
							assert.Len(t, ids, 2)

							_, err = lpClient.GetStake(deployedEnterprise, removedID)
							assert.Error(t, err)
						})
						t.Run("When there is a borrower", func(t *testing.T) {
							var borrowerClient *Client
							var borrowerAddress, service3Address common.Address
							t.Run("setup", func(t *testing.T) {
								t.Run("initiate private key", func(t *testing.T) {
									borrowerClient = privkeyToClient(t, "337d1d2404bfe027ec0464f61e57f9398c56ec511645da01fe13bf937eaa71c0", ethClient, chainID)
									borrowerAddress = borrowerClient.signerAddress
								})
								t.Run("transfer eth to borrower", func(t *testing.T) {
									ctx, cancel := context.WithTimeout(context.Background(), time.Second)
									defer cancel()
									nonce, err := ethClient.PendingNonceAt(ctx, deployerAddress)
									assert.NoError(t, err)

									tx := types.NewTransaction(nonce, borrowerAddress, oneToken, 100000, nil, nil)
									signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), ownerPK)
									assert.NoError(t, err)
									err = ethClient.SendTransaction(ctx, signedTx)
									assert.NoError(t, err)
								})
								t.Run("transfer erc20 to borrower", func(t *testing.T) {
									opts, cancel := getTransactOpts()
									defer cancel()
									transactor, err := bindings.NewERC20Transactor(baseEnterpriseParams.LiquidityTokenAddress, ethClient)
									assert.NoError(t, err)

									tx, err := transactor.Transfer(opts, borrowerAddress, new(big.Int).Mul(oneToken, big.NewInt(500)))
									assert.NoError(t, err)
									assertTxSuccess(t, tx, ethClient)
								})
								t.Run("setup enterprise for testing", func(t *testing.T) {
									params := baseServiceParams
									params.Name = "Test Service 3"
									params.Symbol = "IQPT3"
									tx, err := client.RegisterService(deployedEnterprise, params)
									assert.NoError(t, err)
									assertTxSuccess(t, tx, ethClient)

									addr, err := getPowerToken(t, tx, ethClient, deployedEnterprise)
									assert.NoError(t, err)

									service3Address = addr
									wrappedAmount := big.NewInt(0).Mul(oneToken, big.NewInt(100))
									tx, err = client.ApproveLiquidityTokensToService(service3Address, wrappedAmount)
									assert.NoError(t, err)
									assertTxSuccess(t, tx, ethClient)

									tx, err = client.SwapIn(service3Address, wrappedAmount)
									assert.NoError(t, err)
									assertTxSuccess(t, tx, ethClient)

									tx, err = client.SetConverterAddress(deployedEnterprise, converterAddress)
									assert.NoError(t, err)
									assertTxSuccess(t, tx, ethClient)

									tx, err = client.SetEnterpriseLoanCollectionPeriod(deployedEnterprise, 24*3600)
									assert.NoError(t, err)
									assertTxSuccess(t, tx, ethClient)

									tx, err = client.SetRenterOnlyReturnPeriod(deployedEnterprise, 12*3600)
									assert.NoError(t, err)
									assertTxSuccess(t, tx, ethClient)

									tx, err = client.SetInterestReserveHalvingPeriod(deployedEnterprise, 24*3600*7)
									assert.NoError(t, err)
									assertTxSuccess(t, tx, ethClient)
								})
							})
							t.Run("allows to estimate a loan via enterprise", func(t *testing.T) {
								estimate, err := borrowerClient.EstimateRentalFee(service3Address, baseEnterpriseParams.LiquidityTokenAddress, new(big.Int).Mul(oneToken, big.NewInt(1000)), 10*24*3600)
								assert.NoError(t, err)

								est := estimate.GcFee.Add(estimate.GcFee, estimate.PoolFee)
								est = est.Add(est, estimate.ServiceFee)
								shouldBe, _ := new(big.Int).SetString("305999999999998231422", 10)
								assert.Equal(t, shouldBe, est)
							})
							t.Run("allows to get loan estimation breakdown via service", func(t *testing.T) {
								details, err := borrowerClient.EstimateRentalFee(service3Address, baseEnterpriseParams.LiquidityTokenAddress, new(big.Int).Mul(oneToken, big.NewInt(1000)), 10*24*3600)
								assert.NoError(t, err)

								interest, _ := new(big.Int).SetString("290999999999998318117", 10)
								gcFee, _ := new(big.Int).SetString("5999999999999965322", 10)
								serviceFee, _ := new(big.Int).SetString("8999999999999947983", 10)
								assert.Equal(t, interest, details.PoolFee)
								assert.Equal(t, gcFee, details.GcFee)
								assert.Equal(t, serviceFee, details.ServiceFee)
							})
							t.Run("allows to borrow power tokens", func(t *testing.T) {
								powerTokenBalanceBefore, err := borrowerClient.GetTokenBalance(service3Address, borrowerAddress)
								assert.NoError(t, err)

								amount := new(big.Int).Mul(oneToken, big.NewInt(100))
								var duration uint32 = 10 * 24 * 3600
								estimate, err := borrowerClient.EstimateRentalFee(service3Address, baseEnterpriseParams.LiquidityTokenAddress, amount, duration)
								assert.NoError(t, err)

								est := big.NewInt(0).Add(estimate.GcFee, estimate.PoolFee)
								est = big.NewInt(0).Add(est, estimate.ServiceFee)

								tx, err := borrowerClient.ApproveLiquidityTokensToEnterprise(deployedEnterprise, est)
								assert.NoError(t, err)
								assertTxSuccess(t, tx, ethClient)

								tx, err = borrowerClient.Rent(deployedEnterprise, service3Address, baseEnterpriseParams.LiquidityTokenAddress, amount, est, duration)
								assert.NoError(t, err)
								assertTxSuccess(t, tx, ethClient)

								powerTokenBalanceAfter, err := borrowerClient.GetTokenBalance(service3Address, borrowerAddress)
								assert.NoError(t, err)

								assert.Equal(t, new(big.Int).Add(powerTokenBalanceBefore, amount), powerTokenBalanceAfter)
							})
							t.Run("when loans are taken", func(t *testing.T) {
								t.Run("retrieves pwoer token balance", func(t *testing.T) {
									balance, err := borrowerClient.GetPowerTokenBalance(service3Address, borrowerAddress)
									assert.NoError(t, err)

									assert.Equal(t, new(big.Int).Mul(oneToken, big.NewInt(100)), balance)

									availableBalance, err := borrowerClient.GetPowerTokenAvailableBalance(service3Address, borrowerAddress)
									assert.NoError(t, err)
									assert.Equal(t, big.NewInt(0).Int64(), availableBalance.Int64())
								})
								t.Run("retrieves energy cap value at the specific time", func(t *testing.T) {
									energy, err := borrowerClient.GetEnergyAt(service3Address, borrowerAddress, uint32(time.Now().Unix()+3600))
									assert.NoError(t, err)

									assert.True(t, energy.Cmp(big.NewInt(0)) > 0)
								})
								t.Run("retrieves a list of borrow tokens", func(t *testing.T) {
									res, err := borrowerClient.GetRentalTokenIDs(deployedEnterprise, borrowerAddress)
									assert.NoError(t, err)
									assert.Len(t, res, 1)
								})
								t.Run("retrieves loan info by borrow token ID", func(t *testing.T) {
									res, err := borrowerClient.GetRentalTokenIDs(deployedEnterprise, borrowerAddress)
									assert.NoError(t, err)

									info, err := borrowerClient.GetRentalAgreement(deployedEnterprise, res[0])
									assert.NoError(t, err)
									assert.Equal(t, new(big.Int).Mul(oneToken, big.NewInt(100)), info.RentalAmount)
								})
								t.Run("allows to return a loan", func(t *testing.T) {
									res, err := borrowerClient.GetRentalTokenIDs(deployedEnterprise, borrowerAddress)
									assert.NoError(t, err)

									tx, err := borrowerClient.ReturnRental(deployedEnterprise, res[0])
									assert.NoError(t, err)
									assertTxSuccess(t, tx, ethClient)
								})
							})
						})
					})
				})

			})
		})
	})
}

func getPowerToken(t *testing.T, tx *types.Transaction, ethClient *ethclient.Client, enterpriseAddress common.Address) (common.Address, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	rcp, err := ethClient.TransactionReceipt(ctx, tx.Hash())
	assert.NoError(t, err)

	block := rcp.BlockNumber

	filterer, err := bindings.NewEnterpriseFilterer(enterpriseAddress, ethClient)
	assert.NoError(t, err)

	to := block.Uint64()
	iter, err := filterer.FilterServiceRegistered(&bind.FilterOpts{
		Context: ctx,
		Start:   block.Uint64(),
		End:     &to,
	}, []common.Address{})
	assert.NoError(t, err)

	for iter.Next() {
		return iter.Event.PowerToken, nil
	}
	return common.Address{}, errors.New("no power token found")
}

func assertTxSuccess(t *testing.T, tx *types.Transaction, ethClient *ethclient.Client) {
	if tx == nil {
		assert.Fail(t, "tx is nil")
		return
	}

	assert.Eventually(t, func() bool {
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
		defer cancel()
		receipt, err := ethClient.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			return false
		}
		return receipt.Status == types.ReceiptStatusSuccessful
	}, time.Second*1, time.Millisecond*100)
}

func privkeyToClient(t *testing.T, pk string, ethClient *ethclient.Client, chainID *big.Int) *Client {
	privateKey, err := crypto.HexToECDSA(pk)
	assert.NoError(t, err)

	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	addr := crypto.PubkeyToAddress(*publicKey)

	liquidityProviderSigner := func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		return types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	}

	return NewClient(ethClient, time.Second, addr, liquidityProviderSigner)
}
