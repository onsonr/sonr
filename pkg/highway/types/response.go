package types

// LoginStartResponse is the response to a login start request
type LoginStartResponse struct {
	Challenge      string                            `json:"challenge"`
	RequestOptions PublicKeyCredentialRequestOptions `json:"requestOptions"`
	Origin         string                            `json:"origin"`
	ExpiresAt      int64                             `json:"expiresAt"`
}

// LoginFinishResponse is the response to a login finish request
type LoginFinishResponse struct{}

// RegistrationStartResponse is the response to a registration start request
type RegistrationStartResponse struct {
	Challenge       string                             `json:"challenge"`
	CreationOptions PublicKeyCredentialCreationOptions `json:"creationOptions"`
	Origin          string                             `json:"origin"`
	ExpiresAt       int64                              `json:"expiresAt"`
}

// RegistrationFinishResponse is the response to a registration finish request
type RegistrationFinishResponse struct {
	Address   string `json:"address"`
	TxHash    string `json:"txHash"`
	Timestamp string `json:"timestamp"`
}
