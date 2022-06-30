package schemas

import (
	"errors"

	"github.com/ipld/go-ipld-prime/datamodel"
	rt "github.com/sonr-io/sonr/x/registry/types"
	st "github.com/sonr-io/sonr/x/schema/types"
)

var (
	errAccountNotProvided        = errors.New("No Acct active")
	errSchemaFieldsInvalid       = errors.New("Supplied Schema is invalid")
	errVerficationMethodNotFound = errors.New("Supplied Schema is invalid")
	endpointVaultLive            = "https://vault.sonr.ws"
)

type AppSchemaInternal interface {
	GetAllWhatIs()
	GetAllSchemaDefinitions()
	BuildSchemaFromDefinition() (datamodel.NodeBuilder, error)
	VerifyObject(doc map[string]interface{}, def st.SchemaDefinition) error
	WithAcct(whoIs rt.WhoIs)
	GetDocFromAcct() (*rt.DIDDocument, error)
	GetVerificationFromAccount(id string) (*rt.VerificationMethod, error)
}

type appSchemaInternalImpl struct {
	vaultEndpoint     string
	schemaDefinitions map[string]*st.SchemaDefinition
	WhatIs            []*st.WhatIs
	Acct              *rt.WhoIs
	nodes             map[string]*datamodel.NodeBuilder
}

func New() AppSchemaInternal {
	return &appSchemaInternalImpl{
		vaultEndpoint:     endpointVaultLive,
		schemaDefinitions: make(map[string]*st.SchemaDefinition),
		WhatIs:            make([]*st.WhatIs, 0),
		Acct:              nil,
	}
}

func NewWithAcct(whoIs *rt.WhoIs) AppSchemaInternal {
	return &appSchemaInternalImpl{
		vaultEndpoint:     endpointVaultLive,
		schemaDefinitions: make(map[string]*st.SchemaDefinition),
		WhatIs:            make([]*st.WhatIs, 0),
		Acct:              whoIs,
	}
}
