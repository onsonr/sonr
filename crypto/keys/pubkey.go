package keys

type PubKey interface {
	Type() string
	Value() string
}
