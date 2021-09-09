package discovery

import "github.com/ethereum/go-ethereum/common"

//go:generate stringer -type=IQContractName

type IQContractName int

type IQContractAddresses = map[IQContractName]common.Address
type IQChains = map[ChainID]ChainInfo
type IQContractDeployments = map[ChainID]IQContractAddresses

const (
	BorrowToken IQContractName = iota
	DefaultConverter
	Enterprise
	EnterpriseFactory
	InterestToken
	PowerToken
)

type ChainExplorer struct {
	Name     string  `json:"name,omitempty"`
	Url      string  `json:"url,omitempty"`
	Standard string  `json:"standard,omitempty"`
	Icon     *string `json:"icon,omitempty"`
}

// CAIP-2 blockchain ID
// see: https://github.com/ChainAgnostic/CAIPs/blob/master/CAIPs/caip-2.md
type ChainID string

const (
	BinanceSmartChainMainnet ChainID = "eip155:56"
	BinanceSmartChainTestnet ChainID = "eip155:97"
)

type ChainInfo struct {
	Name           string
	ChainID        int64
	ShortName      string
	Chain          string
	Network        string
	NetworkID      int64
	NativeCurrency Currency
	Rpc            []string
	Explorers      []ChainExplorer
	Faucets        []string
	InfoURL        string
}

type Currency struct {
	Name     string
	Symbol   string
	Decimals int64
}
