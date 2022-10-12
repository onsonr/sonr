package schemas

import (
	st "github.com/sonr-io/sonr/x/schema/types"
)

func (s *SchemaImpl) GetDID() string {
	return s.whatIs.Did
}

func (s *SchemaImpl) GetField(name string) (*st.SchemaField, bool) {
	for _, f := range s.fields {
		if f.Name == name {
			return f, true
		}
	}
	return nil, false
}

func (s *SchemaImpl) GetFields() []*st.SchemaField {
	return s.fields
}

func (s *SchemaImpl) GetLabel() string {
	return s.whatIs.Schema.Label
}
