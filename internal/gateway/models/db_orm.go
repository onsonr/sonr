package models

import (
	"gorm.io/gorm"
)

type Credential struct {
	gorm.Model
	Handle     string `json:"handle"`
	ID         string `json:"id"`
	Origin     string `json:"origin"`
	Type       string `json:"type"`
	Transports string `json:"transports"`
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
	Challenge      string `json:"challenge"`
}

type User struct {
	gorm.Model
	Address string `json:"address"`
	Handle  string `json:"handle"`
	Origin  string `json:"origin"`
	Name    string `json:"name"`
	CID     string `json:"cid"`
}
