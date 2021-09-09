package account

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type ChainIdentifier string

const (
	ChainMaticMumbai  = "eip155:80001"
	ChainMaticMainnet = "eip155:137"
	ChainEthGoerli    = "eip155:5"
	ChainEthMainnet   = "eip155:1"
)

type AccountID struct {
	chainIdentifier ChainIdentifier
	chainID         int64
	address         common.Address
}

func (a AccountID) String() string {
	return fmt.Sprintf("%v:%v", a.chainIdentifier, strings.ToLower(a.address.Hex()))
}

func (a AccountID) Address() common.Address {
	return a.address
}

func (a AccountID) ChainID() int64 {
	return a.chainID
}

func NewAccountID(chainID ChainIdentifier, address common.Address) (AccountID, error) {
	aid := AccountID{
		chainIdentifier: chainID,
		address:         address,
	}

	chid, err := extractChainID(chainID)
	if err != nil {
		return AccountID{}, err
	}
	aid.chainID = chid
	return aid, nil
}

func extractChainID(chid ChainIdentifier) (int64, error) {
	splits := strings.Split(string(chid), ":")
	return strconv.ParseInt(splits[len(splits)-1], 10, 64)
}
