package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	DisplayName string
	Name        string
	Origin      string
	Credentials []Credential
}
