package resourcetype

// Resource is a unique identifier for a thing, usually stored state. Resources
// are organized by string types
type Resource interface {
	Type() ResourceType
	Value() string
	Contains(b Resource) bool
}

type stringResource struct {
	t ResourceType
	v string
}

func (r stringResource) Type() ResourceType {
	return r.t
}

func (r stringResource) Value() string {
	return r.v
}

func (r stringResource) Contains(b Resource) bool {
	return r.Type() == b.Type() && len(r.Value()) <= len(b.Value())
}

func NewResource(typ ResourceType, val string) Resource {
	return stringResource{
		t: typ,
		v: val,
	}
}
