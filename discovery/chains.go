package discovery

var Chains IQChains = IQChains{
	BinanceSmartChainMainnet: {
		Name:    "Binance Smart Chain Mainnet",
		Chain:   "BSC",
		Network: "mainnet",
		Rpc: []string{
			"https://bsc-dataseed1.binance.org",
			"https://bsc-dataseed2.binance.org",
			"https://bsc-dataseed3.binance.org",
			"https://bsc-dataseed4.binance.org",
			"https://bsc-dataseed1.defibit.io",
			"https://bsc-dataseed2.defibit.io",
			"https://bsc-dataseed3.defibit.io",
			"https://bsc-dataseed4.defibit.io",
			"https://bsc-dataseed1.ninicoin.io",
			"https://bsc-dataseed2.ninicoin.io",
			"https://bsc-dataseed3.ninicoin.io",
			"https://bsc-dataseed4.ninicoin.io",
			"wss://bsc-ws-node.nariox.org",
		},
		Faucets: []string{},
		NativeCurrency: Currency{
			Name:     "Binance Chain Native Token",
			Symbol:   "BNB",
			Decimals: 18,
		},
		InfoURL:   "https://www.binance.org",
		ShortName: "bnb",
		ChainID:   56,
		NetworkID: 56,
		Explorers: []ChainExplorer{
			{
				Name:     "bscscan",
				Url:      "https://bscscan.com",
				Standard: "EIP3091",
			},
		},
	},
	BinanceSmartChainTestnet: {
		Name:    "Binance Smart Chain Testnet",
		Chain:   "BSC",
		Network: "Chapel",
		Rpc: []string{
			"https://data-seed-prebsc-1-s1.binance.org:8545",
			"https://data-seed-prebsc-2-s1.binance.org:8545",
			"https://data-seed-prebsc-1-s2.binance.org:8545",
			"https://data-seed-prebsc-2-s2.binance.org:8545",
			"https://data-seed-prebsc-1-s3.binance.org:8545",
			"https://data-seed-prebsc-2-s3.binance.org:8545",
		},
		Faucets: []string{"https://testnet.binance.org/faucet-smart"},
		NativeCurrency: Currency{
			Name:     "Binance Chain Native Token",
			Symbol:   "tBNB",
			Decimals: 18,
		},
		InfoURL:   "https://testnet.binance.org",
		ShortName: "bnb",
		ChainID:   97,
		NetworkID: 97,
		Explorers: []ChainExplorer{
			{
				Name:     "bscscan-testnet",
				Url:      "https://testnet.bscscan.com",
				Standard: "EIP3091",
			},
		},
	},
}
