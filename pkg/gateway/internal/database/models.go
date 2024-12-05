package database

type Session struct {
	ID               string `json:"id"`
	BrowserName      string `json:"browserName"`
	BrowserVersion   string `json:"browserVersion"`
	UserArchitecture string `json:"userArchitecture"`
	Platform         string `json:"platform"`
	PlatformVersion  string `json:"platformVersion"`
	DeviceModel      string `json:"deviceModel"`
}

type User struct {
	Address   string `json:"address"`
	Handle    string `json:"handle"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	VaultCID  string `json:"vaultCID"`
}
