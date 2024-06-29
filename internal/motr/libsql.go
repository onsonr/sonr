package motr

import (
	"github.com/di-dao/sonr/internal/motr/models"
	"github.com/di-dao/sonr/pkg/fs"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Vault struct {
	DB *gorm.DB
}

func SeedTables(file fs.File) (*Vault, error) {
	db, err := gorm.Open(sqlite.Open(file.Path()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.Wallet{}, &models.Credential{}, &models.Profile{})
	return &Vault{DB: db}, nil
}
