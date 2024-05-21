package proxy

type Challenge struct {
	// ID is the ID of the challenge.
	ID string `json:"id"`

	// Value is the value of the challenge.
	Value []byte `json:"value"`
}
