package did

import (
	"encoding/hex"

	"github.com/okx/go-wallet-sdk/coins/cosmos"

	modulev1 "github.com/sonrhq/sonr/api/sonr/identity/module/v1"
)

func GetAddressByPublicKey(pubKey []byte, coinType modulev1.CoinType) (string, error) {
	hrp := GetCoinTypeHRP(coinType)
	pubKeyHex := hex.EncodeToString(pubKey)
	return cosmos.GetAddressByPublicKey(pubKeyHex, hrp)
}
