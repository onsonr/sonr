package types

type LoginStartRequest struct {
	Username string `json:"username"`
	Origin   string `json:"origin"`
}

type LoginFinishRequest struct {
	Username string              `json:"username"`
	Response PublicKeyCredential `json:"response"`
	Origin   string              `json:"origin"`
}

type RegistrationStartRequest struct {
	Username string `json:"username"`
	Origin   string `json:"origin"`
}

type RegistrationFinishRequest struct {
	Username string              `json:"username"`
	Response PublicKeyCredential `json:"response"`
	Origin   string              `json:"origin"`
}
