package db

import (
	"fmt"

	"github.com/ncruces/go-sqlite3"
	_ "github.com/ncruces/go-sqlite3/embed"
)

type DB struct {
	*sqlite3.Conn
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

func Open(config *DBConfig) (*DB, error) {
	conn, err := sqlite3.Open(config.ConnectionString())
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	db := &DB{
		Conn: conn,
	}

	if err := createTables(db); err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to create tables: %w", err)
	}
	return db, nil
}

func createTables(db *DB) error {
	tables := []string{
		createAccountsTable,
		createAssetsTable,
		createChainsTable,
		createCredentialsTable,
		createKeysharesTable,
		createProfilesTable,
		createPropertiesTable,
		createPermissionsTable,
	}

	for _, table := range tables {
		err := db.Exec(table)
		if err != nil {
			return fmt.Errorf("failed to create table: %w", err)
		}
	}

	return nil
}

// AddAccount adds a new account to the database
func (db *DB) AddAccount(name, address string) error {
	return db.Exec(insertAccountQuery(name, address))
}

// AddAsset adds a new asset to the database
func (db *DB) AddAsset(name, symbol string, decimals int, chainID int64) error {
	return db.Exec(insertAssetQuery(name, symbol, decimals, chainID))
}

// AddChain adds a new chain to the database
func (db *DB) AddChain(name, networkID string) error {
	return db.Exec(insertChainQuery(name, networkID))
}

// AddCredential adds a new credential to the database
func (db *DB) AddCredential(
	handle, controller, attestationType, origin string,
	credentialID, publicKey []byte,
	transport string,
	signCount uint32,
	userPresent, userVerified, backupEligible, backupState, cloneWarning bool,
) error {
	return db.Exec(insertCredentialQuery(
		handle,
		controller,
		attestationType,
		origin,
		credentialID,
		publicKey,
		transport,
		signCount,
		userPresent,
		userVerified,
		backupEligible,
		backupState,
		cloneWarning,
	))
}

//
// // AddProfile adds a new profile to the database
// func (db *DB) AddProfile(
// 	id, subject, controller, originURI string,
// 	publicMetadata, privateMetadata string,
// ) error {
// 	return db.statements["insertProfile"].Exec(
// 		id,
// 		subject,
// 		controller,
// 		originURI,
// 		publicMetadata,
// 		privateMetadata,
// 	)
// }
//
// // AddProperty adds a new property to the database
// func (db *DB) AddProperty(profileID, key string, accumulator, propertyKey []byte) error {
// 	return db.statements["insertProperty"].Exec(profileID, key, accumulator, propertyKey)
// }
//
// // AddPermission adds a new permission to the database
// func (db *DB) AddPermission(
// 	serviceID string,
// 	grants []DIDNamespace,
// 	scopes []PermissionScope,
// ) error {
// 	grantsJSON, err := json.Marshal(grants)
// 	if err != nil {
// 		return fmt.Errorf("failed to marshal grants: %w", err)
// 	}
//
// 	scopesJSON, err := json.Marshal(scopes)
// 	if err != nil {
// 		return fmt.Errorf("failed to marshal scopes: %w", err)
// 	}
//
// 	return db.statements["insertPermission"].Exec(
// 		serviceID,
// 		string(grantsJSON),
// 		string(scopesJSON),
// 	)
// }
//
// // GetPermission retrieves a permission from the database
// func (db *DB) GetPermission(serviceID string) ([]DIDNamespace, []PermissionScope, error) {
// 	stmt := db.statements["getPermission"]
// 	if err := stmt.Exec(serviceID); err != nil {
// 		return nil, nil, fmt.Errorf("failed to execute statement: %w", err)
// 	}
//
// 	if !stmt.Step() {
// 		return nil, nil, fmt.Errorf("permission not found")
// 	}
//
// 	grantsJSON := stmt.ColumnText(0)
// 	scopesJSON := stmt.ColumnText(1)
//
// 	var grants []DIDNamespace
// 	err := json.Unmarshal([]byte(grantsJSON), &grants)
// 	if err != nil {
// 		return nil, nil, fmt.Errorf("failed to unmarshal grants: %w", err)
// 	}
//
// 	var scopes []PermissionScope
// 	err = json.Unmarshal([]byte(scopesJSON), &scopes)
// 	if err != nil {
// 		return nil, nil, fmt.Errorf("failed to unmarshal scopes: %w", err)
// 	}
//
// 	return grants, scopes, nil
// }
