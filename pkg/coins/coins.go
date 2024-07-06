package coins

// Coin represents a cryptocurrency
type Coin interface {
	// FormatAddress formats a public key into an address
	FormatAddress(pubKey []byte) (string, error)

	// GetIndex returns the coin type index
	GetIndex() int64

	// GetPath returns the coin component path
	GetPath() uint32

	// GetSymbol returns the coin symbol
	GetSymbol() string

	// GetMethod returns the coin DID method
	GetMethod() string

	// GetName returns the coin name
	GetName() string
}

// DefaultCoins is a list of default coins used in the vault
var DefaultCoins = []Coin{
	CoinBTC,
	CoinETH,
	CoinSNR,
}

var (
	// Bitcoin mainnet
	CoinBTC = &coin{
		Name:   "Bitcoin",
		Index:  0,
		Path:   0x80000000,
		Symbol: "BTC",
		Hrp:    "bc",
		Method: "btcr",
	}

	// Ethereum
	CoinETH = &coin{
		Name:   "Ethereum",
		Index:  60,
		Path:   0x8000003c,
		Symbol: "ETH",
		Method: "ethr",
	}

	// Sonr
	CoinSNR = &coin{
		Name:   "Sonr",
		Index:  703,
		Path:   0x800002bf,
		Symbol: "SNR",
		Hrp:    "idx",
		Method: "sonr",
	}
)

// CoinBTCType is the coin type for BTC
const CoinBTCType = int64(0)

// CoinETHType is the coin type for ETH
const CoinETHType = int64(60)

// CoinSNRType is the coin type for SNR
const CoinSNRType = int64(703)
