package types

// AddrPrefix returns the address prefix for the given coin type.
func (ct CoinType) AddrPrefix() string {
	switch ct {
	case CoinType_CoinType_BITCOIN:
		return "btc"
	case CoinType_CoinType_ETHEREUM:
		return "0x"
	case CoinType_CoinType_LITECOIN:
		return "ltc"
	case CoinType_CoinType_DOGE:
		return "doge"
	case CoinType_CoinType_SONR:
		return "snr"
	case CoinType_CoinType_COSMOS:
		return "cosmos"
	case CoinType_CoinType_FILECOIN:
		return "f"
	case CoinType_CoinType_HNS:
		return "hs"
	case CoinType_CoinType_TESTNET:
		return "test"
	default:
		return "test"
	}
}

// Index returns the index for the given coin type.
func (ct CoinType) Index() int32 {
	switch ct {
	case CoinType_CoinType_BITCOIN:
		return 0
	case CoinType_CoinType_ETHEREUM:
		return 60
	case CoinType_CoinType_LITECOIN:
		return 2
	case CoinType_CoinType_DOGE:
		return 3
	case CoinType_CoinType_SONR:
		return 703
	case CoinType_CoinType_COSMOS:
		return 118
	case CoinType_CoinType_FILECOIN:
		return 461
	case CoinType_CoinType_HNS:
		return 5353
	case CoinType_CoinType_TESTNET:
		return 1
	default:
		return 1
	}
}

// Name returns the name for the given coin type.
func (ct CoinType) Name() string {
	switch ct {
	case CoinType_CoinType_BITCOIN:
		return "Bitcoin"
	case CoinType_CoinType_ETHEREUM:
		return "Ethereum"
	case CoinType_CoinType_LITECOIN:
		return "Litecoin"
	case CoinType_CoinType_DOGE:
		return "Dogecoin"
	case CoinType_CoinType_SONR:
		return "Sonr"
	case CoinType_CoinType_COSMOS:
		return "Cosmos"
	case CoinType_CoinType_FILECOIN:
		return "Filecoin"
	case CoinType_CoinType_HNS:
		return "Handshake"
	default:
		return "Testnet"
	}
}

// Symbol returns the symbol for the given coin type.
func (ct CoinType) Symbol() string {
	switch ct {
	case CoinType_CoinType_BITCOIN:
		return "BTC"
	case CoinType_CoinType_ETHEREUM:
		return "ETH"
	case CoinType_CoinType_LITECOIN:
		return "LTC"
	case CoinType_CoinType_DOGE:
		return "DOGE"
	case CoinType_CoinType_SONR:
		return "SNR"
	case CoinType_CoinType_COSMOS:
		return "ATOM"
	case CoinType_CoinType_FILECOIN:
		return "FIL"
	case CoinType_CoinType_HNS:
		return "HNS"
	default:
		return "TESTNET"
	}
}
