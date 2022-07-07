package schemas

import (
	"errors"

	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/sonr-io/sonr/pkg/crypto"
	"github.com/sonr-io/sonr/pkg/did"
	st "github.com/sonr-io/sonr/x/schema/types"
)

var (
	errAccountNotProvided        = errors.New("no Acct active")
	errAccountAlreadyDefined     = errors.New("cannot reassign account once assigned")
	errSchemaFieldsInvalid       = errors.New("supplied Schema is invalid")
	errVerficationMethodNotFound = errors.New("supplied Schema is invalid")
	errIdNotFound                = errors.New("Id not found")
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
		Acessor method for node map related by definition did
	*/
	GetNodeMap() map[string]datamodel.Node

	/*
		Acessor method for node map related by did
	*/
	GetWhatIsMap() map[string]*st.WhatIs

	/*
		Builds a linkage of IPLD nodes from the provided schema definition
		returns the `Node` and assigns it to the given id internally.
	*/
	BuildNodesFromDefinition(id string, def *st.SchemaDefinition, object map[string]interface{}) (datamodel.Node, error)

	/*
		Returns an error if any of the keys within provided data dont match the given schema definition
		useful for verifying
	*/
	VerifyObject(doc map[string]interface{}, def *st.SchemaDefinition) error

	/*
		Adds an account to be used for operations, once defined will raise `errAccountAlreadyDefined`
	*/
	WithAcct(wallet *crypto.MPCWallet) error

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
	GetTopLevelNodeById(id string) (datamodel.Node, error)

	/*
		Get an IPLD object as a flat path of nodes.
	*/
	GetPath(id string) (datamodel.ListIterator, error)
}

/*
	Creates a relationship between a schema definition and the 'whatIs' definition
*/
type SchemaRelationShip struct {
	definition *st.SchemaDefinition
	did        did.DID
}

type appSchemaInternalImpl struct {
	schemaDefinitions map[string]*st.SchemaDefinition
	WhatIs            map[string]*st.WhatIs
	Acct              *crypto.MPCWallet
	nodes             map[string]datamodel.Node
}

func New() AppSchemaInternal {
	asi := &appSchemaInternalImpl{
		schemaDefinitions: make(map[string]*st.SchemaDefinition),
		WhatIs:            make(map[string]*st.WhatIs),
		Acct:              nil,
		nodes:             make(map[string]datamodel.Node),
	}

	return asi
}

func NewWithAcct(wallet *crypto.MPCWallet) AppSchemaInternal {
	asi := &appSchemaInternalImpl{
		schemaDefinitions: make(map[string]*st.SchemaDefinition),
		WhatIs:            make(map[string]*st.WhatIs, 0),
		Acct:              wallet,
		nodes:             make(map[string]datamodel.Node),
	}

	return asi
}
