package schemas

import (
	"github.com/ipld/go-ipld-prime/datamodel"
	st "github.com/sonr-io/sonr/x/schema/types"
)

/*
	Underyling api definition of internal implementation of Schemas.
	Higher level APIs implementing schema features

	Version: 0.1.0
*/
type Schema interface {

	/*
		Builds a linkage of IPLD nodes from the provided schema definition
		returns the `Node` and assigns it to the given id internally.
	*/
	BuildNodesFromDefinition(object map[string]interface{}) error

	/*
		Returns an error if any of the keys within provided data dont match the given schema definition
		useful for verifying
	*/
	VerifyObject(doc map[string]interface{}) error

	/*
		Encodes a given IPLD Node as JSON
	*/
	EncodeDagJson() ([]byte, error)

	/*
		Encodes a given IPLD Node as CBOR
	*/
	EncodeDagCbor() ([]byte, error)

	/*
		Encodes a given IPLD Node as CBOR
	*/
	DecodeDagJson(buffer []byte) error

	/*
		Decodes a given IPLD Node as CBOR
	*/
	DecodeDagCbor(buffer []byte) error

	/*
		Returns a list of SchemaKindDefinitions, composing the schema
	*/
	GetSchema() ([]*st.SchemaKindDefinition, error)

	/*
		Returns top level node of a hydrated schema
	*/
	GetNode() (datamodel.Node, error)

	/*
		Get an IPLD object as an iterator
	*/
	GetPath() (datamodel.ListIterator, error)
}
