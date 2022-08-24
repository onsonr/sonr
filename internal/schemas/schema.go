package schemas

import (
	"github.com/sonr-io/sonr/pkg/motor/x/object"
)

func (s *schemaImpl) GetField(name string) (object.SchemaField, bool) {
	for _, f := range s.fields {
		if f.Name == name {
			return f, true
		}
	}
	return nil, false
}

func (s *schemaImpl) GetFields() []object.SchemaField {
	result := make([]object.SchemaField, len(s.fields))
	for i, f := range s.fields {
		result[i] = object.SchemaField(f)
	}
	return result
}

func (s *schemaImpl) GetLabel() string {
	return s.whatIs.Schema.Label
}
