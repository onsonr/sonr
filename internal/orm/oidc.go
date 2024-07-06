package orm

type UserInfo struct {
	DID   string
	Sub   string `json:"sub"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	// Add other claims as needed
}
