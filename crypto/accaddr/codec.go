package bip32

type CoinType uint32

const (
	// Hardened offset for BIP-44 derivation
	HardenedOffset uint32 = 0x80000000

	// Registered coin types for BIP-44
	CoinTypeBitcoin  CoinType = CoinType(0 + HardenedOffset)
	CoinTypeEthereum CoinType = CoinType(60 + HardenedOffset)
	CoinTypeSonr     CoinType = CoinType(703 + HardenedOffset)
)

// Uint32 returns the coin type as a uint32.
func (c CoinType) Uint32() uint32 {
	return uint32(c)
}
