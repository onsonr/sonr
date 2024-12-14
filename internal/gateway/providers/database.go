package providers

import (
	"gorm.io/gorm"
)

type Database interface{}

type DatabaseService struct {
	db *gorm.DB
}

func NewDatabaseService(db *gorm.DB) Database {
	return &DatabaseService{
		db: db,
	}
}
