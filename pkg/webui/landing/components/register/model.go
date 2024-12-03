package register

// Types for credential data
type PublicKeyCredentialCreationOptions struct {
	Challenge       string `json:"challenge"`
	RpName          string `json:"rpName"`
	RpID            string `json:"rpId"`
	UserID          string `json:"userId"`
	UserName        string `json:"userName"`
	UserDisplayName string `json:"userDisplayName"`
	Timeout         int    `json:"timeout,omitempty"`
	AttestationType string `json:"attestationType,omitempty"`
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
