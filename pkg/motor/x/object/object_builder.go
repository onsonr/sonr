package object

import (
	"errors"
	"fmt"

	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
)

type ObjectBuilder struct {
	schema       Schema
	objectClient ObjectClient

	label  string
	values Object
}

// NewBuilder creates an ObjectBuilder to build uploadable objects
func NewBuilder(schema Schema, objectClient ObjectClient) *ObjectBuilder {
	return &ObjectBuilder{
		schema:       schema,
		objectClient: objectClient,

		values: make(map[string]interface{}),
	}
}

// SetLabel sets the label for the object
func (ob *ObjectBuilder) SetLabel(label string) {
	ob.label = label
}

// Set sets the value of the given field, ensuring it matches in type
func (ob *ObjectBuilder) Set(fieldName string, value interface{}) error {
	var field SchemaField
	if f, found := ob.schema.GetField(fieldName); !found {
		return fmt.Errorf("no field '%s' in schema '%s'", field, ob.schema.GetLabel())
	} else {
		field = f
	}

	if !field.TryValue(value) {
		return fmt.Errorf("value '%s' not of kind '%s'", value, field.GetKind())
	}

	ob.values[fieldName] = value
	return nil
}

// Remove removes a field value, returning false if the field did not exist and true otherwise
func (ob *ObjectBuilder) Remove(field string) bool {
	if _, ok := ob.values[field]; ok {
		delete(ob.values, field)
		return true
	}
	return false
}

func (ob *ObjectBuilder) Get(field string) interface{} {
	if v, ok := ob.values[field]; ok {
		return v
	}
	return nil
}

func (ob *ObjectBuilder) Has(field string) bool {
	_, ok := ob.values[field]
	return ok
}

// Build checks that the object is properly built and returns the map
func (ob *ObjectBuilder) Build() (Object, error) {
	if ob.label == "" {
		return ob.values, errors.New("object is missing a label")
	}

	missingFields := make([]string, 0)
	for _, field := range ob.schema.GetFields() {
		if _, ok := ob.values[field.GetName()]; !ok {
			missingFields = append(missingFields, field.GetName())
		}
	}

	if len(missingFields) != 0 {
		return ob.values, fmt.Errorf("missing fields in object: %s", missingFields)
	}

	return ob.values, nil
}

func (ob *ObjectBuilder) Upload() (*mt.UploadObjectResponse, error) {
	toUpload, err := ob.Build()
	if err != nil {
		return nil, err
	}

	return ob.objectClient.CreateObject(ob.label, toUpload)
}

func (ob *ObjectBuilder) GetByCID(cid string) (Object, error) {
	return ob.objectClient.GetObject(cid)
}
