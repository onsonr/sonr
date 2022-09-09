package object

import motor "github.com/sonr-io/sonr/third_party/types/motor/api/v1"

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
		object map[string]interface{}) (*motor.UploadObjectResponse, error)

	/*
		Retrieves an object for the data store
	*/
	GetObject(cid string) (map[string]interface{}, error)
}
