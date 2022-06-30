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

func (as *appSchemaInternalImpl) GetDocFromAcct() (*rt.DIDDocument, error) {
	if as.Acct != nil {
		return nil, errAccountNotProvided
	}

	doc := as.Acct.GetDidDocument()

	return doc, nil
}

func (as *appSchemaInternalImpl) GetVerificationFromAccount(id string) (*rt.VerificationMethod, error) {
	if as.Acct != nil {
		return nil, errAccountNotProvided
	}

	doc := as.Acct.GetDidDocument()

	vms := doc.GetVerificationMethod()

	for i := 1; i < len(vms); i++ {
		if vms[i].Id == id {
			return vms[i], nil
		}
	}

	return nil, errVerficationMethodNotFound
}
