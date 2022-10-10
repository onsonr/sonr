package document

import (
	"errors"
	"fmt"

	id "github.com/sonr-io/sonr/internal/document"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	st "github.com/sonr-io/sonr/x/schema/types"
)

type DocumentBuilder struct {
	schema    id.Schema
	docClient DocumentClient

	label  string
	values Document
}

// NewBuilder creates an ObjectBuilder to build uploadable objects
func NewBuilder(schema id.Schema, objectClient DocumentClient) *DocumentBuilder {
	return &DocumentBuilder{
		schema:    schema,
		docClient: objectClient,

		values: make(map[string]interface{}),
	}
}

// SetLabel sets the label for the object
func (db *DocumentBuilder) SetLabel(label string) {
	db.label = label
}

// Set sets the value of the given field, ensuring it matches in type
func (db *DocumentBuilder) Set(fieldName string, value interface{}) error {
	var field *st.SchemaField
	if f, found := db.schema.GetField(fieldName); !found {
		return fmt.Errorf("no field '%s' in schema '%s'", field, db.schema.GetLabel())
	} else {
		field = f
	}

	if !field.TryValue(value) {
		return fmt.Errorf("value '%s' not of kind '%s'", value, field.GetKind())
	}

	db.values[fieldName] = value
	return nil
}

// Remove removes a field value, returning false if the field did not exist and true otherwise
func (db *DocumentBuilder) Remove(field string) bool {
	if _, ok := db.values[field]; ok {
		delete(db.values, field)
		return true
	}
	return false
}

func (db *DocumentBuilder) Get(field string) interface{} {
	if v, ok := db.values[field]; ok {
		return v
	}
	return nil
}

func (db *DocumentBuilder) Has(field string) bool {
	_, ok := db.values[field]
	return ok
}

// Build checks that the object is properly built and returns the map
func (db *DocumentBuilder) Build() (Document, error) {
	if db.label == "" {
		return db.values, errors.New("object is missing a label")
	}

	missingFields := make([]string, 0)
	for _, field := range db.schema.GetFields() {
		if _, ok := db.values[field.GetName()]; !ok {
			missingFields = append(missingFields, field.GetName())
		}
	}

	if len(missingFields) != 0 {
		return db.values, fmt.Errorf("missing fields in object: %s", missingFields)
	}

	return db.values, nil
}

func (db *DocumentBuilder) Upload() (*mt.UploadDocumentResponse, error) {
	toUpload, err := db.Build()
	if err != nil {
		return nil, err
	}

	return db.docClient.CreateDocument(db.label, db.schema.GetDID(), toUpload)
}

func (db *DocumentBuilder) GetByCID(cid string) (*st.Document, error) {
	return db.docClient.GetDocument(cid)
}
