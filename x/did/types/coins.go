package types

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

// CoinBTCType is the coin type for BTC
const CoinBTCType = int64(0)

// CoinETHType is the coin type for ETH
const CoinETHType = int64(60)

// CoinSNRType is the coin type for SNR
const CoinSNRType = int64(703)
