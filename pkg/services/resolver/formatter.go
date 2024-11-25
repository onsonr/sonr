package resolver

type DID string

func (d DID) String() string {
	return string(d)
}
