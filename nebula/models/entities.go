package models

// ╭───────────────────────────────────────────────────────────╮
// │                  Credentials Management                   │
// ╰───────────────────────────────────────────────────────────╯

type PublicKeyParams struct {
	Type      string `json:"type"`
	Algorithm string `json:"alg"`
}

type RelayingPartyEntity struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type UserEntity struct {
	ID          []byte `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

type WebauthnExtensions struct{}

// ╭───────────────────────────────────────────────────────────╮
// │                  Generic Methods                          │
// ╰───────────────────────────────────────────────────────────╯

func DefaultPublicKeyParams() *PublicKeyParams {
	return &PublicKeyParams{
		Type:      "public-key",
		Algorithm: "-7",
	}
}
