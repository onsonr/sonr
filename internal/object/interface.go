package object

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

	/*
		Builds Object with schema definition,
		Returns an error if the object is invalid with the given Schema
		Calling assert will not perform persistence operations
	*/
	assert(object map[string]interface{}) error
}
