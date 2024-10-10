package models

// ╭───────────────────────────────────────────────────────────╮
// │                  Credentials Management                   │
// ╰───────────────────────────────────────────────────────────╯

type PublicKeyParams struct {
	Algorithm string `json:"algorithm"`
	Encoding  string `json:"encoding"`
	Curve     string `json:"curve"`
}

type RelayingPartyEntity struct {
	ID          []byte `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Icon        string `json:"icon"`
}

type UserEntity struct {
	ID          []byte `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

type WebauthnExtensions struct{}
