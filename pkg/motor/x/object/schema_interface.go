package object

import "github.com/sonr-io/sonr/x/schema/types"

type SchemaField interface {
	GetName() string
	GetKind() types.SchemaKind
	TryValue(interface{}) bool
}

type Schema interface {
	GetLabel() string
	GetFields() []SchemaField
	GetField(string) (SchemaField, bool)
}
