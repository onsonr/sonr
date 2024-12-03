package login

type LoginRequest struct {
	Subject     string
	Action      string
	Origin      string
	Status      string
	Ping        string
	BlockSpeed  string
	BlockHeight string
}
