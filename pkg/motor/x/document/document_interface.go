package document

import (
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	st "github.com/sonr-io/sonr/x/schema/types"
)

/*
Underlying api definition of internal implementation of Objects
Higher level APIs implementing DocumentClient features

Version 0.1.0
*/
type DocumentClient interface {
	/*
		Persists an object definition to the storage configured within the module.
	*/
	CreateDocument(
		label string,
		schemaDid string,
		document map[string]interface{}) (*mt.UploadDocumentResponse, error)

	/*
		Retrieves an object for the data store
	*/
	GetDocument(cid string) (*st.Document, error)
}
