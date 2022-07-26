package schemas

import (
	"errors"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/sonr-io/sonr/pkg/client"
	st "github.com/sonr-io/sonr/x/schema/types"
)

var (
	errCidInvalid          = errors.New("attempted to load a non CID link")
	errSchemaFieldsInvalid = errors.New("supplied Schema is invalid")
	errIdNotFound          = errors.New("Id not found")
)

/*
	Underyling api definition of internal implementation of Schemas.
	Higher level APIs implementing schema features

	Version: 0.1.0
*/
type AppSchemaInternal interface {

	/*
		Builds a linkage of IPLD nodes from the provided schema definition
		returns the `Node` and assigns it to the given id internally.
	*/
	BuildNodesFromDefinition(fields []*st.SchemaKindDefinition, object map[string]interface{}) (datamodel.Node, error)

	/*
		Returns an error if any of the keys within provided data dont match the given schema definition
		useful for verifying
	*/
	VerifyObject(doc map[string]interface{}, fields []*st.SchemaKindDefinition) error

	/*
		Encodes a given IPLD Node as JSON
	*/
	EncodeDagJson(node datamodel.Node) ([]byte, error)

	/*
		Encodes a given IPLD Node as CBOR
	*/
	EncodeDagCbor(node datamodel.Node) ([]byte, error)

	/*
		Encodes a given IPLD Node as CBOR
	*/
	DecodeDagJson(buffer []byte) (datamodel.Node, error)

	/*
		Decodes a given IPLD Node as CBOR
	*/
	DecodeDagCbor(buffer []byte) (datamodel.Node, error)

	/*
		Get an IPLD object as a flat path of nodes.
	*/
	GetPath(node datamodel.Node) (datamodel.ListIterator, error)
}

/*
	Interface for implementing querying logic for chain interactions
*/
type SchemaDataResolver interface {
	/*
		Gets all `whatIs` objects for the account `whoIs` an error if the query fails
	*/
	GetWhatIs(creator string, did string) (*st.WhatIs, error)

	/*
		Gets all `whatIs` objects for the account `whoIs` an error if the query fails
	*/
	GetSchemaByCid(cid string) ([]*st.SchemaKindDefinition, error)
}

/*
	Creates a relationship between a schema definition and the 'whatIs' definition
*/
type SchemaRelationShip struct {
	definition *st.SchemaDefinition
	cid        string
}

type appSchemaInternalImpl struct {
	schemas map[string][]*st.SchemaKindDefinition
	whatIs  map[string]*st.WhatIs
	client  *client.Client
	shell   shell.Shell
}

func New(persistenceUri string, endpointType client.ConnEndpointType) AppSchemaInternal {
	asi := &appSchemaInternalImpl{
		// Holds relation of schema's to CID
		schemas: make(map[string][]*st.SchemaKindDefinition),
		// Holds relation of WhatIs's to DID
		whatIs: make(map[string]*st.WhatIs),
		client: client.NewClient(endpointType),
		shell:  *shell.NewShell(persistenceUri),
	}

	return asi
}
