package coins

import (
	"github.com/okx/go-wallet-sdk/coins/cosmos"
)

// NewSonrAddress returns a new sonr address from a public key
func NewSonrAddress(pubKeyHex string) (string, error) {
    addr, err := cosmos.GetAddressByPublicKey(pubKeyHex, "idx")
    if err!= nil {
        return "", err
    }
    return addr, nil
}
