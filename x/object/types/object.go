package types

import (
	"strings"

)

func (o *ObjectDoc) Validate(b *ObjectDoc) bool {
	if o.GetLabel() != b.GetLabel() {
		return false
	}

	for k, v := range o.GetFields() {
		if b.GetFields()[k] != v {
			return false
		}
	}
	return true
}

// AddFields takes a list of fields and adds it to ObjectDoc
func (o *ObjectDoc) AddFields(l ...*TypeField) {
	for i, v := range o.GetFields() {
		if strings.EqualFold(v.GetName(), l[i].GetName()) {
			o.Fields[i] = l[i]
		}
	}
}

// RemoveFields takes a list of ObjectFields
// and removes the matching label from the ObjectDoc
func (o *ObjectDoc) RemoveFields(l ...*TypeField) {
	for i, v := range o.GetFields() {
		if strings.EqualFold(v.GetName(), l[i].GetName()) {
			o.Fields = append(o.Fields[:i], o.Fields[i+1:]...)
		}
	}
}
