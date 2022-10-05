package document

import (
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/sonr-io/sonr/internal/schemas"
	st "github.com/sonr-io/sonr/x/schema/types"
)

/*
	Underyling api definition of internal implementation of Schemas.
	Higher level APIs implementing schema features

	Version: 0.1.0
*/
type Schema interface {

	/*
		Returns the DID of the schema
	*/
	GetDID() string

	/*
		Returns the label of the schema
	*/
	GetLabel() string

	/*
		Builds a linkage of IPLD nodes from the provided schema definition
		returns the `Node` and assigns it to the given id internally.
	*/
	BuildNodesFromDefinition(label, schemaDid string, object map[string]interface{}) error

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
		Returns a specific SchemaFields for the schema
	*/
	GetField(name string) (*st.SchemaField, bool)

	/*
		Returns a list of SchemaFields for the schema
	*/
	GetFields() []*st.SchemaField

	/*
		Returns the underlying schema type Schema
	*/
	GetSchema() (*st.Schema, error)

	/*
		Returns a Schema interface contained by the Schema
	*/
	GetSubSchema(did string) (*schemas.SchemaImpl, error)

	/*
		Returns top level node of a hydrated schema
	*/
	GetNode() (datamodel.Node, error)

	/*
		Get an IPLD object as an iterator
	*/
	GetPath() (datamodel.ListIterator, error)
}
