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

type PublicKeyCredentialRequestOptions struct {
	Challenge        string                 `json:"challenge"`
	RpID             string                 `json:"rpId"`
	Timeout          int                    `json:"timeout,omitempty"`
	UserVerification string                 `json:"userVerification,omitempty"`
	AllowCredentials []CredentialDescriptor `json:"allowCredentials,omitempty"`
}

type CredentialDescriptor struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}
