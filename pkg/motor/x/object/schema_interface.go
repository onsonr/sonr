package object

import "fmt"

type SchemaKind interface {
	fmt.Stringer
	GetValue() int32
}

type SchemaField interface {
	GetName() string
	GetKind() SchemaKind
	TryValue(interface{}) bool
}

type Schema interface {
	GetLabel() string
	GetFields() []SchemaField
	GetField(string) (SchemaField, bool)
}
