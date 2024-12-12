package sessions

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Address string `json:"address"`
	Handle  string `json:"handle"`
	Name    string `json:"name"`
	CID     string `json:"cid"`
}

type Session struct {
	gorm.Model
	ID               string `json:"id" gorm:"primaryKey"`
	BrowserName      string `json:"browserName"`
	BrowserVersion   string `json:"browserVersion"`
	UserArchitecture string `json:"userArchitecture"`
	Platform         string `json:"platform"`
	PlatformVersion  string `json:"platformVersion"`
	DeviceModel      string `json:"deviceModel"`
	UserHandle       string `json:"userHandle"`
	FirstName        string `json:"firstName"`
	LastInitial      string `json:"lastInitial"`
	VaultAddress     string `json:"vaultAddress"`
	HumanSum         int    `json:"humanSum"`
	Challenge        string `json:"challenge"`
}
