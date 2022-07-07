package schemas

import "github.com/sonr-io/sonr/pkg/crypto"

func (as *appSchemaInternalImpl) WithAcct(wallet *crypto.MPCWallet) error {
	if as.Acct != nil {
		return errAccountAlreadyDefined
	}

	as.Acct = wallet

	return nil
}
