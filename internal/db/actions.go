package db

import (
	"fmt"

	"github.com/onsonr/sonr/internal/db/orm"
	"gorm.io/gorm"
)

// createInitialTables creates the initial tables in the database.
func createInitialTables(db *gorm.DB) (*DB, error) {
	err := db.AutoMigrate(
		&orm.Account{},
		&orm.Asset{},
		&orm.Keyshare{},
		&orm.Credential{},
		&orm.Profile{},
		&orm.Property{},
		&orm.Permission{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create table: %w", err)
	}

	return &DB{db}, nil
}

// AddAccount adds a new account to the database
func (db *DB) AddAccount(account *orm.Account) error {
	tx := db.Create(account)
	if tx.Error != nil {
		return fmt.Errorf("failed to add account: %w", tx.Error)
	}

	return nil
}

// AddKeyshare adds a new keyshare to the database
func (db *DB) AddKeyshare(keyshare *orm.Keyshare) error {
	tx := db.Create(keyshare)

	if tx.Error != nil {
		return fmt.Errorf("failed to add keyshare: %w", tx.Error)
	}

	return nil
}

// AddCredential adds a new credential to the database
func (db *DB) AddCredential(credential *orm.Credential) error {
	tx := db.Create(credential)

	if tx.Error != nil {
		return fmt.Errorf("failed to add credential: %w", tx.Error)
	}

	return nil
}

// AddProfile adds a new profile to the database
func (db *DB) AddProfile(profile *orm.Profile) error {
	tx := db.Create(profile)

	if tx.Error != nil {
		return fmt.Errorf("failed to add profile: %w", tx.Error)
	}

	return nil
}

// AddProperty adds a new property to the database
func (db *DB) AddProperty(property *orm.Property) error {
	tx := db.Create(property)

	if tx.Error != nil {
		return fmt.Errorf("failed to add property: %w", tx.Error)
	}

	return nil
}

// AddPermission adds a new permission to the database
func (db *DB) AddPermission(permission *orm.Permission) error {
	tx := db.Create(permission)

	if tx.Error != nil {
		return fmt.Errorf("failed to add permission: %w", tx.Error)
	}

	return nil
}
