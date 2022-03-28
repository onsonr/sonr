package crypto

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/sonr-io/core/config"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func GenerateKeyring(cnfg *config.SonrConfig, kr keyring.Keyring, options ...GenerateOption) (KeySet, string, error) {
	g := defaultOptions(cnfg)
	for _, option := range options {
		if err := option(g); err != nil {
			return nil, "", err
		}
	}

	// Add keys and see they return in alphabetical order
	_, mnemonic, err := kr.NewMnemonic(ProvisionUid(g.sname), keyring.English, sdk.FullFundraiserPath, g.passphrase, hd.Secp256k1)
	if err != nil {
		return nil, "", err
	}

	// Create default sonr key
	ks, err := CreateKeySet(mnemonic)
	if err != nil {
		return nil, "", err
	}

	// Copy keys to keyring if not already there
	_, err = ks.CopyToKeyring(kr, g.sname)
	if err == nil {
		err = ks.Export(g.config.WalletFolder())
		if err != nil {
			return nil, "", err
		}
	}

	return ks, mnemonic, nil
}

func ProvisionUid(s string) string {
	return fmt.Sprintf("%s.provision", s)
}
