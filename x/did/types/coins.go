package types

import (
	fmt "fmt"

	"github.com/cosmos/cosmos-sdk/types/bech32"

	ethcommon "github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

var (
	CoinBTC = &Coin{
		Name:   "Bitcoin",
		Index:  0,
		Path:   0x80000000,
		Symbol: "BTC",
		Hrp:    "bc",
	}

	CoinETH = &Coin{
		Name:   "Ethereum",
		Index:  60,
		Path:   0x8000003c,
		Symbol: "ETH",
	}

	CoinSNR = &Coin{
		Name:   "Sonr",
		Index:  703,
		Path:   0x800002bf,
		Symbol: "SNR",
		Hrp:    "idx",
	}
)

func DefaultCoins() []Coin {
	return []Coin{
		*CoinBTC,
		*CoinETH,
		*CoinSNR,
	}
}

func (c *Coin) FormatAddress(pubKey []byte) (string, error) {
	if c.Hrp != "" {
		return bech32.ConvertAndEncode(c.Hrp, pubKey)
	}
	if c.Index == 60 {
		return ethcommon.BytesToAddress(ethcrypto.Keccak256(pubKey[1:])[12:]).Hex(), nil
	}
	return "", fmt.Errorf("unsupported coin")
}

func (c *Coin) Equal(c2 *Coin) bool {
	return c.Index == c2.Index && c.Path == c2.Path
}
