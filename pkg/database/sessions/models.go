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
	ID             string `json:"id" gorm:"primaryKey"`
	BrowserName    string `json:"browserName"`
	BrowserVersion string `json:"browserVersion"`
	Platform       string `json:"platform"`
	IsDesktop      bool   `json:"isDesktop"`
	IsMobile       bool   `json:"isMobile"`
	IsTablet       bool   `json:"isTablet"`
	IsTV           bool   `json:"isTV"`
	IsBot          bool   `json:"isBot"`
	UserHandle     string `json:"userHandle"`
	VaultAddress   string `json:"vaultAddress"`
	Challenge      string `json:"challenge"`
}
