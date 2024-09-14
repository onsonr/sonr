package db

import (
	"crypto/rand"

	"github.com/ncruces/go-sqlite3/gormlite"
	"golang.org/x/crypto/argon2"
	"gorm.io/gorm"
	"lukechampine.com/adiantum/hbsh"
	"lukechampine.com/adiantum/hpolyc"
)

type DB struct {
	*gorm.DB
}

func New(opts ...DBOption) (*DB, error) {
	config := &DBConfig{
		fileName: "vault.db",
	}
	for _, opt := range opts {
		opt(config)
	}
	gormdb, err := gorm.Open(gormlite.Open(config.ConnectionString()))
	if err != nil {
		return nil, err
	}
	db, err := createInitialTables(gormdb)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// HBSH creates an HBSH cipher given a key.
func (c *DB) HBSH(key []byte) *hbsh.HBSH {
	if len(key) != 32 {
		// Key is not appropriate, return nil.
		return nil
	}
	return hpolyc.New(key)
}

// KDF gets a key from a secret.
func (c *DB) KDF(secret string) []byte {
	if secret == "" {
		// No secret is given, generate a random key.
		key := make([]byte, 32)
		n, _ := rand.Read(key)
		return key[:n]
	}
	// Hash the secret with a KDF.
	return argon2.IDKey([]byte(secret), []byte("hpolyc"), 3, 64*1024, 4, 32)
}
