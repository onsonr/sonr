package services

import "gorm.io/gorm"

type UserService struct {
	DB *gorm.DB
}
