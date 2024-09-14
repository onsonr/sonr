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
		&orm.Credential{},
		&orm.Keyshare{},
		&orm.Permission{},
		&orm.Profile{},
		&orm.Property{},
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

// GetAccount gets an account from the database
func (db *DB) GetAccount(account *orm.Account) error {
	tx := db.First(account)
	if tx.Error != nil {
		return fmt.Errorf("failed to get account: %w", tx.Error)
	}

	return nil
}

// UpdateAccount updates an existing account in the database
func (db *DB) UpdateAccount(account *orm.Account) error {
	tx := db.Save(account)
	if tx.Error != nil {
		return fmt.Errorf("failed to update account: %w", tx.Error)
	}

	return nil
}

// DeleteAccount deletes an existing account from the database
func (db *DB) DeleteAccount(account *orm.Account) error {
	tx := db.Delete(account)
	if tx.Error != nil {
		return fmt.Errorf("failed to delete account: %w", tx.Error)
	}

	return nil
}

// AddAsset adds a new asset to the database
func (db *DB) AddAsset(asset *orm.Asset) error {
	tx := db.Create(asset)

	if tx.Error != nil {
		return fmt.Errorf("failed to add asset: %w", tx.Error)
	}

	return nil
}

// GetAsset gets an asset from the database
func (db *DB) GetAsset(asset *orm.Asset) error {
	tx := db.First(asset)

	if tx.Error != nil {
		return fmt.Errorf("failed to get asset: %w", tx.Error)
	}

	return nil
}

// UpdateAsset updates an existing asset in the database
func (db *DB) UpdateAsset(asset *orm.Asset) error {
	tx := db.Save(asset)

	if tx.Error != nil {
		return fmt.Errorf("failed to update asset: %w", tx.Error)
	}

	return nil
}

// DeleteAsset deletes an existing asset from the database
func (db *DB) DeleteAsset(asset *orm.Asset) error {
	tx := db.Delete(asset)

	if tx.Error != nil {
		return fmt.Errorf("failed to delete asset: %w", tx.Error)
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

// GetCredential gets an credential from the database
func (db *DB) GetCredential(credential *orm.Credential) error {
	tx := db.First(credential)

	if tx.Error != nil {
		return fmt.Errorf("failed to get credential: %w", tx.Error)
	}

	return nil
}

// UpdateCredential updates an existing credential in the database
func (db *DB) UpdateCredential(credential *orm.Credential) error {
	tx := db.Save(credential)

	if tx.Error != nil {
		return fmt.Errorf("failed to update credential: %w", tx.Error)
	}

	return nil
}

// DeleteCredential deletes an existing credential from the database
func (db *DB) DeleteCredential(credential *orm.Credential) error {
	tx := db.Delete(credential)

	if tx.Error != nil {
		return fmt.Errorf("failed to delete credential: %w", tx.Error)
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

// GetKeyshare gets an keyshare from the database
func (db *DB) GetKeyshare(keyshare *orm.Keyshare) error {
	tx := db.First(keyshare)

	if tx.Error != nil {
		return fmt.Errorf("failed to get keyshare: %w", tx.Error)
	}

	return nil
}

// UpdateKeyshare updates an existing keyshare in the database
func (db *DB) UpdateKeyshare(keyshare *orm.Keyshare) error {
	tx := db.Save(keyshare)

	if tx.Error != nil {
		return fmt.Errorf("failed to update keyshare: %w", tx.Error)
	}

	return nil
}

// DeleteKeyshare deletes an existing keyshare from the database
func (db *DB) DeleteKeyshare(keyshare *orm.Keyshare) error {
	tx := db.Delete(keyshare)

	if tx.Error != nil {
		return fmt.Errorf("failed to delete keyshare: %w", tx.Error)
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

// GetPermission gets an permission from the database
func (db *DB) GetPermission(permission *orm.Permission) error {
	tx := db.First(permission)

	if tx.Error != nil {
		return fmt.Errorf("failed to get permission: %w", tx.Error)
	}

	return nil
}

// UpdatePermission updates an existing permission in the database
func (db *DB) UpdatePermission(permission *orm.Permission) error {
	tx := db.Save(permission)

	if tx.Error != nil {
		return fmt.Errorf("failed to update permission: %w", tx.Error)
	}

	return nil
}

// DeletePermission deletes an existing permission from the database
func (db *DB) DeletePermission(permission *orm.Permission) error {
	tx := db.Delete(permission)

	if tx.Error != nil {
		return fmt.Errorf("failed to delete permission: %w", tx.Error)
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

// GetProfile gets an profile from the database
func (db *DB) GetProfile(profile *orm.Profile) error {
	tx := db.First(profile)

	if tx.Error != nil {
		return fmt.Errorf("failed to get profile: %w", tx.Error)
	}

	return nil
}

// UpdateProfile updates an existing profile in the database
func (db *DB) UpdateProfile(profile *orm.Profile) error {
	tx := db.Save(profile)

	if tx.Error != nil {
		return fmt.Errorf("failed to update profile: %w", tx.Error)
	}

	return nil
}

// DeleteProfile deletes an existing profile from the database
func (db *DB) DeleteProfile(profile *orm.Profile) error {
	tx := db.Delete(profile)

	if tx.Error != nil {
		return fmt.Errorf("failed to delete profile: %w", tx.Error)
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

// GetProperty gets an property from the database
func (db *DB) GetProperty(property *orm.Property) error {
	tx := db.First(property)

	if tx.Error != nil {
		return fmt.Errorf("failed to get property: %w", tx.Error)
	}

	return nil
}

// UpdateProperty updates an existing property in the database
func (db *DB) UpdateProperty(property *orm.Property) error {
	tx := db.Save(property)

	if tx.Error != nil {
		return fmt.Errorf("failed to update property: %w", tx.Error)
	}

	return nil
}

// DeleteProperty deletes an existing property from the database
func (db *DB) DeleteProperty(property *orm.Property) error {
	tx := db.Delete(property)

	if tx.Error != nil {
		return fmt.Errorf("failed to delete property: %w", tx.Error)
	}

	return nil
}
