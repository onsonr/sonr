package crypto

import (
	"log"

	"github.com/cosmos/cosmos-sdk/crypto"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/sonr-io/core/config"
)

type GenerateOption func(*options) error

func WithPassphrase(s string) GenerateOption {
	return func(o *options) error {
		o.passphrase = s
		return nil
	}
}

type options struct {
	sname      string
	config     *config.SonrConfig
	passphrase string
}

func defaultOptions(cnfg *config.SonrConfig) *options {
	return &options{
		sname:      cnfg.AccountName,
		passphrase: "bad-passphrase",
		config:     cnfg,
	}
}

// ExportWallet returns armored private key and public key
func ExportWallet(kr keyring.Keyring, sname string, passphrase string) (string, error) {
	armor, err := kr.ExportPrivKeyArmor(sname, passphrase)
	if err != nil {
		return "", err
	}
	return armor, nil
}

// RestoreWallet restores a private key from ASCII armored format.
func RestoreWallet(sname string, armor string, passphrase string) (keyring.Keyring, error) {
	privKey, algo, err := crypto.UnarmorDecryptPrivKey(armor, passphrase)
	if err != nil {
		return nil, err
	}
	kr := keyring.NewInMemory()
	if err := kr.ImportPrivKey(sname, algo, passphrase); err != nil {
		return nil, err
	}
	log.Println(privKey.PubKey())
	return kr, nil
}
