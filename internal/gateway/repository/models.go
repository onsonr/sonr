// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package repository

import (
	"database/sql"
	"time"
)

type Credential struct {
	ID           int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    sql.NullTime
	Handle       string
	CredentialID string
	Origin       string
	Type         string
	Transports   string
}

type Session struct {
	ID             string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      sql.NullTime
	BrowserName    string
	BrowserVersion string
	Platform       string
	IsDesktop      int64
	IsMobile       int64
	IsTablet       int64
	IsTv           int64
	IsBot          int64
	Challenge      string
	IsHumanFirst   int64
	IsHumanLast    int64
}

type User struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
	Address   string
	Handle    string
	Origin    string
	Name      string
	Cid       string
}