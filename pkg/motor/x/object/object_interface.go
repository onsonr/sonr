package object

type ObjectReference struct {
	Did   string
	Label string
	Cid   string
}

/*
	Object definition to be returned after object creation
*/
type ObjectUploadResult struct {
	Code      int32
	Reference *ObjectReference
	Message   string
}

/*
	Underlying api definition of internal implementation of Objects
	Higher level APIs implementing ObjectClient features

	Version 0.1.0
*/
type ObjectClient interface {
	/*
		Persists an object definition to the storage configured within the module.
	*/
	CreateObject(
		label string,
		object map[string]interface{}) (*ObjectUploadResult, error)

	/*
		Retrieves an object for the data store
	*/
	GetObject(cid string) (map[string]interface{}, error)
}
