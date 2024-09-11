package db

import "gorm.io/gorm"

type DB struct {
	*gorm.DB
}

func New(opts ...DBOption) *DBConfig {
	config := &DBConfig{
		fileName: "vault.db",
	}
	for _, opt := range opts {
		opt(config)
	}
	return config
}
