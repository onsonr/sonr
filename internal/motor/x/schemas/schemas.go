package schemas

import (
	"errors"

	"github.com/ipld/go-ipld-prime/datamodel"
	rt "github.com/sonr-io/sonr/x/registry/types"
	st "github.com/sonr-io/sonr/x/schema/types"
)

var (
	errAccountNotProvided        = errors.New("no Acct active")
	errAccountAlreadyDefined     = errors.New("cannot reassign account once assigned")
	errSchemaFieldsInvalid       = errors.New("supplied Schema is invalid")
	errVerficationMethodNotFound = errors.New("supplied Schema is invalid")
	errIdNotFound                = errors.New("Id not found")
	endpointVaultLive            = "https://vault.sonr.ws"
)

/*
	Underyling api definition of internal implementation of Schemas.
	Higher level APIs implementing schema features

	Version: 0.1.0
*/
type AppSchemaInternal interface {

	/*
		Gets all `whatIs` objects for the account `whoIs` an error if the query fails
	*/
	GetAllWhatIs() error

	/*
		Gets all `whatIs` objects for the account `whoIs` an error if the query fails
	*/
	GetAllSchemaDefinitions() error

	/*
		Builds a linkage of IPLD nodes from the provided schema definition
		returns the `Node` and assigns it to the given id internally.
	*/
	BuildNodesFromDefinition(id string, def *st.SchemaDefinition) (*datamodel.Node, error)

	/*
		Returns an error if any of the keys within provided data dont match the given schema definition
		useful for verifying
	*/
	VerifyObject(doc map[string]interface{}, def st.SchemaDefinition) error

	/*
		Adds an account to be used for operations, once defined will raise `errAccountAlreadyDefined`
	*/
	WithAcct(whoIs rt.WhoIs) error

	/*
		Returns the did Document relating to the `whoIs` instance
	*/
	GetDocFromAcct() (*rt.DIDDocument, error)

	/*
		Returns a verification method defined within the `whoIs` did document
	*/
	GetVerificationFromAccount(id string) (*rt.VerificationMethod, error)

	/*
		Encodes a given IPLD Node as JSON
	*/
	EncodeDagJson(node datamodel.Node) ([]byte, error)

	/*
		Encodes a given IPLD Node as CBOR
	*/
	EncodeDagCbor(node datamodel.Node) ([]byte, error)

	/*
		Get a top level IPLD node by the associated WhatIs did
	*/
	GetNodeById(id string) (datamodel.Node, error)
}

type appSchemaInternalImpl struct {
	vaultEndpoint     string
	schemaDefinitions map[string]*st.SchemaDefinition
	WhatIs            map[string]*st.WhatIs
	Acct              *rt.WhoIs
	nodes             map[string]*datamodel.Node
}

func New() AppSchemaInternal {
	return &appSchemaInternalImpl{
		vaultEndpoint:     endpointVaultLive,
		schemaDefinitions: make(map[string]*st.SchemaDefinition),
		WhatIs:            make(map[string]*st.WhatIs),
		Acct:              nil,
	}
}

func NewWithAcct(whoIs *rt.WhoIs) AppSchemaInternal {
	return &appSchemaInternalImpl{
		vaultEndpoint:     endpointVaultLive,
		schemaDefinitions: make(map[string]*st.SchemaDefinition),
		WhatIs:            make(map[string]*st.WhatIs, 0),
		Acct:              whoIs,
	}
}
