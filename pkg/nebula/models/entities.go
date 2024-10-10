package models

// ╭───────────────────────────────────────────────────────────╮
// │                  Credentials Management                   │
// ╰───────────────────────────────────────────────────────────╯

type UserEntity struct {
	ID          []byte `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

type RelayingPartyEntity struct {
	ID          []byte `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Icon        string `json:"icon"`
}
