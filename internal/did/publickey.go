package did

import (
	"github.com/okx/go-wallet-sdk/coins/cosmos"
)

func GetAddressByPublicKey(pubKeyHex string, HRP string) (string, error) {
	return cosmos.GetAddressByPublicKey(pubKeyHex, HRP)
}
