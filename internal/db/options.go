package db

import (
	_ "github.com/ncruces/go-sqlite3/embed"
	"github.com/onsonr/sonr/internal/db/orm"
)

type DBOption func(config *DBConfig)

func WitDir(dir string) DBOption {
	return func(config *DBConfig) {
		config.Dir = dir
	}
}

func WithSecretKey(secretKey string) DBOption {
	return func(config *DBConfig) {
		config.SecretKey = secretKey
	}
}

type DBConfig struct {
	Dir       string
	SecretKey string

	fileName           string
	initialAccounts    []*orm.Account
	initialAssets      []*orm.Asset
	initialCredentials []*orm.Credential
	initialKeyshares   []*orm.Keyshare
	initialPermissions []*orm.Permission
	initialProfiles    []*orm.Profile
	initialProperties  []*orm.Property
}

func (config *DBConfig) ConnectionString() string {
	connStr := "file:"
	connStr += config.fileName
	return connStr
}
