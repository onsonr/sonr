package common

func CoinTypeFromIndex(i int32) CoinType {
	switch i {
	case 0:
		return CoinType_CoinType_BITCOIN
	case 60:
		return CoinType_CoinType_ETHEREUM
	case 2:
		return CoinType_CoinType_LITECOIN
	case 3:
		return CoinType_CoinType_DOGE
	case 703:
		return CoinType_CoinType_SONR
	case 1:
		return CoinType_CoinType_TESTNET
	default:
		return CoinType_CoinType_TESTNET
	}
}

func CoinTypeFromString(str string) CoinType {
	switch str {
	case "btc":
		return CoinType_CoinType_BITCOIN
	case "0x":
		return CoinType_CoinType_ETHEREUM
	case "ltc":
		return CoinType_CoinType_LITECOIN
	case "doge":
		return CoinType_CoinType_DOGE
	case "snr":
		return CoinType_CoinType_SONR
	case "test":
		return CoinType_CoinType_TESTNET
	default:
		return CoinType_CoinType_TESTNET
	}
}

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
	case CoinType_CoinType_TESTNET:
		return "test"
	default:
		return "test"
	}
}

// Index returns the index for the given coin type.
func (ct CoinType) Index() int {
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
	case CoinType_CoinType_TESTNET:
		return "Testnet"
	default:
		return "Testnet"
	}
}

// PathComponent returns the path component for the given coin type.
func (ct CoinType) PathComponent() uint32 {
	switch ct {
	case CoinType_CoinType_BITCOIN:
		return 0x80000000
	case CoinType_CoinType_ETHEREUM:
		return 0x8000003c
	case CoinType_CoinType_LITECOIN:
		return 0x80000002
	case CoinType_CoinType_DOGE:
		return 0x80000003
	case CoinType_CoinType_SONR:
		return 0x800002bf
	case CoinType_CoinType_TESTNET:
		return 0x80000001
	default:
		return 0x80000001
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
	case CoinType_CoinType_TESTNET:
		return "TEST"
	default:
		return "TESTNET"
	}
}
