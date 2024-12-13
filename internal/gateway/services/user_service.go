package services

import "gorm.io/gorm"

type UserService struct {
	db *gorm.DB
}
