package schemas

import (
	rt "github.com/sonr-io/sonr/x/registry/types"
)

func (as *appSchemaInternalImpl) WithAcct(whoIs rt.WhoIs) error {
	if as.Acct != nil {
		return errAccountAlreadyDefined
	}

	as.Acct = &whoIs

	return nil
}
