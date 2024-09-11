package db

import (
	"crypto/rand"

	"github.com/ncruces/go-sqlite3/gormlite"
	"golang.org/x/crypto/argon2"
	"gorm.io/gorm"
	"lukechampine.com/adiantum/hbsh"
	"lukechampine.com/adiantum/hpolyc"

	_ "github.com/ncruces/go-sqlite3/embed"
)

type DBOption func(config *DBConfig)

func WithDir(dir string) DBOption {
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

	fileName string
}

func (config *DBConfig) ConnectionString() string {
	connStr := "file:"
	connStr += config.Dir + "/" + config.fileName
	return connStr
}

// GormDialector creates a gorm dialector for the database.
func (config *DBConfig) Open() (*DB, error) {
	db, err := gorm.Open(gormlite.Open(config.ConnectionString()))
	if err != nil {
		return nil, err
	}
	return createInitialTables(db)
}

// HBSH creates an HBSH cipher given a key.
func (c *DBConfig) HBSH(key []byte) *hbsh.HBSH {
	if len(key) != 32 {
		// Key is not appropriate, return nil.
		return nil
	}
	return hpolyc.New(key)
}

// KDF gets a key from a secret.
func (c *DBConfig) KDF(secret string) []byte {
	if secret == "" {
		// No secret is given, generate a random key.
		key := make([]byte, 32)
		n, _ := rand.Read(key)
		return key[:n]
	}
	// Hash the secret with a KDF.
	return argon2.IDKey([]byte(secret), []byte("hpolyc"), 3, 64*1024, 4, 32)
}
