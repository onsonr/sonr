package schemas

import (
	"errors"

	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/sonr-io/sonr/pkg/did"
	st "github.com/sonr-io/sonr/x/schema/types"
)

var (
	errAccountNotProvided        = errors.New("no Acct active")
	errAccountAlreadyDefined     = errors.New("cannot reassign account once assigned")
	errCidInvalid                = errors.New("attempted to load a non CID link")
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
		Builds a linkage of IPLD nodes from the provided schema definition
		returns the `Node` and assigns it to the given id internally.
	*/
	BuildNodesFromDefinition(def *st.SchemaDefinition, object map[string]interface{}) (datamodel.Node, error)

	/*
		Returns an error if any of the keys within provided data dont match the given schema definition
		useful for verifying
	*/
	VerifyObject(doc map[string]interface{}, def *st.SchemaDefinition) error

	/*
		Encodes a given IPLD Node as JSON
	*/
	EncodeDagJson(node datamodel.Node) ([]byte, error)

	/*
		Encodes a given IPLD Node as CBOR
	*/
	EncodeDagCbor(node datamodel.Node) ([]byte, error)

	DecodeDagJson(buffer []byte) (datamodel.Node, error)

	DecodeDagCbor(buffer []byte) (datamodel.Node, error)

	/*
		Get an IPLD object as a flat path of nodes.
	*/
	GetPath(node datamodel.Node) (datamodel.ListIterator, error)
}

type SchemaDataResolver interface {
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
}

/*
	Creates a relationship between a schema definition and the 'whatIs' definition
*/
type SchemaRelationShip struct {
	definition *st.SchemaDefinition
	did        did.DID
}

type appSchemaInternalImpl struct{}

func New() AppSchemaInternal {
	asi := &appSchemaInternalImpl{}

	return asi
}
