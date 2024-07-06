package orm

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	DID         string
	DisplayName string
	Name        string
	Origin      string
	Controller  string
	Credentials []Credential
}
