package object

import st "github.com/sonr-io/sonr/x/schema/types"

/*
	Underlying api definition of internal implementation of Objects
	Higher level APIs implementing Object features

	Version 0.1.0
*/
type AppObjectInternal interface {
	/*
		Persists an object definition to the storage configured within the module.
	*/
	CreateObject(label string, fields []*st.SchemaKindDefinition, object map[string]interface{}) (*ObjectUploadResult, error)

	/*
		Retrieves an
	*/
	GetObject(cid string) (map[string]interface{}, error)
}
