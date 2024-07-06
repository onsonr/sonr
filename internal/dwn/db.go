//go:build (linux || darwin || windows || freebsd || illumos) && !sqlite3_nosys

package main

import (
	"database/sql"
	"fmt"

	"github.com/onsonr/hway/crypto"
	"github.com/onsonr/hway/crypto/secret"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

const kVaultDBFileName = "file:demo.db?_pragma=busy_timeout(10000)"

type Database interface {
	ExistsCredential(did string) bool
	ExistsProfile(did string) bool
	ExistsWallet(did string) bool

	GetCredential(did string) (*Credential, error)
	GetProfile(did string) (*Profile, error)
	GetWallet(did string) (*Wallet, error)

	InsertCredentials(credentials ...*Credential) error
	InsertProfiles(profiles ...*Profile) error
	InsertWallets(wallets ...*Wallet) error

	ListCredentials() ([]*Credential, error)
	ListProfiles() ([]*Profile, error)
	ListWallets() ([]*Wallet, error)
}

type Credential struct {
	Transport       string
	Origin          string
	Controller      string
	DID             string
	DisplayName     string
	AttestationType string
	Attachment      string
	AAGUID          []byte
	PublicKey       []byte
	CredentialID    []byte
	ID              int64
	SignCount       uint32
	BackupEligible  bool
	BackupState     bool
	UserVerified    bool
	UserPresent     bool
}

type Profile struct {
	DID         string
	DisplayName string
	Name        string
	Origin      string
	Controller  string
	ID          int64
}

type Wallet struct {
	Address    string
	Controller string
	Name       string
	ChainID    string
	Network    string
	Label      string
	DID        string
	PublicKey  []byte
	ID         int64
	Index      int
	CoinType   int64
}

func seedDB() (Database, error) {
	db, err := sql.Open("sqlite", kVaultDBFileName)
	if err != nil {
		return nil, err
	}

	// Create tables
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS credentials (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			display_name TEXT,
			origin TEXT,
			controller TEXT,
			attestation_type TEXT,
			did TEXT UNIQUE,
			credential_id BLOB,
			public_key BLOB,
			transport TEXT,
			user_present BOOLEAN,
			user_verified BOOLEAN,
			backup_eligible BOOLEAN,
			backup_state BOOLEAN,
			aaguid BLOB,
			sign_count INTEGER,
			attachment TEXT
		);

		CREATE TABLE IF NOT EXISTS profiles (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			did TEXT UNIQUE,
			display_name TEXT,
			name TEXT,
			origin TEXT,
			controller TEXT
		);

		CREATE TABLE IF NOT EXISTS wallets (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			address TEXT,
			controller TEXT,
			name TEXT,
			chain_id TEXT,
			network TEXT,
			label TEXT,
			did TEXT UNIQUE,
			public_key BLOB,
			index_num INTEGER,
			coin_type INTEGER
		);
	`)
	if err != nil {
		return nil, err
	}

	return &embedDB{DB: db}, nil
}

type embedDB struct {
	DB *sql.DB
}

func (db *embedDB) GetCredential(did string) (*Credential, error) {
	credential := new(Credential)
	err := db.DB.QueryRow("SELECT * FROM credentials WHERE did = ?", did).Scan(
		&credential.ID, &credential.DisplayName, &credential.Origin, &credential.Controller,
		&credential.AttestationType, &credential.DID, &credential.CredentialID, &credential.PublicKey,
		&credential.Transport, &credential.UserPresent, &credential.UserVerified,
		&credential.BackupEligible, &credential.BackupState, &credential.AAGUID,
		&credential.SignCount, &credential.Attachment,
	)
	if err != nil {
		return nil, err
	}
	return credential, nil
}

func (db *embedDB) GetProfile(did string) (*Profile, error) {
	profile := new(Profile)
	err := db.DB.QueryRow("SELECT * FROM profiles WHERE did = ?", did).Scan(
		&profile.ID, &profile.DID, &profile.DisplayName, &profile.Name,
		&profile.Origin, &profile.Controller,
	)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (db *embedDB) GetWallet(did string) (*Wallet, error) {
	wallet := new(Wallet)
	err := db.DB.QueryRow("SELECT * FROM wallets WHERE did = ?", did).Scan(
		&wallet.ID, &wallet.Address, &wallet.Controller, &wallet.Name,
		&wallet.ChainID, &wallet.Network, &wallet.Label, &wallet.DID,
		&wallet.PublicKey, &wallet.Index, &wallet.CoinType,
	)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

func (db *embedDB) ExistsCredential(did string) bool {
	var count int
	db.DB.QueryRow("SELECT COUNT(*) FROM credentials WHERE did = ?", did).Scan(&count)
	return count > 0
}

func (db *embedDB) ExistsProfile(did string) bool {
	var count int
	db.DB.QueryRow("SELECT COUNT(*) FROM profiles WHERE did = ?", did).Scan(&count)
	return count > 0
}

func (db *embedDB) ExistsWallet(did string) bool {
	var count int
	db.DB.QueryRow("SELECT COUNT(*) FROM wallets WHERE did = ?", did).Scan(&count)
	return count > 0
}

func (db *embedDB) InsertCredentials(credentials ...*Credential) error {
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`
		INSERT INTO credentials (
			display_name, origin, controller, attestation_type, did, credential_id,
			public_key, transport, user_present, user_verified, backup_eligible,
			backup_state, aaguid, sign_count, attachment
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, c := range credentials {
		_, err = stmt.Exec(
			c.DisplayName, c.Origin, c.Controller, c.AttestationType, c.DID, c.CredentialID,
			c.PublicKey, c.Transport, c.UserPresent, c.UserVerified, c.BackupEligible,
			c.BackupState, c.AAGUID, c.SignCount, c.Attachment,
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (db *embedDB) InsertProfiles(profiles ...*Profile) error {
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`
		INSERT INTO profiles (did, display_name, name, origin, controller)
		VALUES (?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, p := range profiles {
		_, err = stmt.Exec(p.DID, p.DisplayName, p.Name, p.Origin, p.Controller)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (db *embedDB) InsertWallets(wallets ...*Wallet) error {
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`
		INSERT INTO wallets (
			address, controller, name, chain_id, network, label, did,
			public_key, index_num, coin_type
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, w := range wallets {
		_, err = stmt.Exec(
			w.Address, w.Controller, w.Name, w.ChainID, w.Network, w.Label, w.DID,
			w.PublicKey, w.Index, w.CoinType,
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (db *embedDB) ListCredentials() ([]*Credential, error) {
	rows, err := db.DB.Query("SELECT * FROM credentials")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var credentials []*Credential
	for rows.Next() {
		c := new(Credential)
		err := rows.Scan(
			&c.ID, &c.DisplayName, &c.Origin, &c.Controller, &c.AttestationType,
			&c.DID, &c.CredentialID, &c.PublicKey, &c.Transport, &c.UserPresent,
			&c.UserVerified, &c.BackupEligible, &c.BackupState, &c.AAGUID,
			&c.SignCount, &c.Attachment,
		)
		if err != nil {
			return nil, err
		}
		credentials = append(credentials, c)
	}
	return credentials, nil
}

func (db *embedDB) ListProfiles() ([]*Profile, error) {
	rows, err := db.DB.Query("SELECT * FROM profiles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var profiles []*Profile
	for rows.Next() {
		p := new(Profile)
		err := rows.Scan(&p.ID, &p.DID, &p.DisplayName, &p.Name, &p.Origin, &p.Controller)
		if err != nil {
			return nil, err
		}
		profiles = append(profiles, p)
	}
	return profiles, nil
}

func (db *embedDB) ListWallets() ([]*Wallet, error) {
	rows, err := db.DB.Query("SELECT * FROM wallets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wallets []*Wallet
	for rows.Next() {
		w := new(Wallet)
		err := rows.Scan(
			&w.ID, &w.Address, &w.Controller, &w.Name, &w.ChainID, &w.Network,
			&w.Label, &w.DID, &w.PublicKey, &w.Index, &w.CoinType,
		)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, w)
	}
	return wallets, nil
}

func (db *embedDB) WitnessCredential(publicKey crypto.PublicKey, did string) ([]byte, error) {
	if !db.ExistsCredential(did) {
		return nil, fmt.Errorf("credential with DID %s does not exist", did)
	}

	pk, err := secret.NewKey("credentials", publicKey)
	if err != nil {
		return nil, err
	}

	creds, err := db.ListCredentials()
	if err != nil {
		return nil, err
	}

	credIDStrs := make([]string, len(creds))
	for i, c := range creds {
		credIDStrs[i] = c.DID
	}

	acc, err := pk.CreateAccumulator(credIDStrs...)
	if err != nil {
		return nil, err
	}

	witness, err := pk.CreateWitness(acc, did)
	if err != nil {
		return nil, err
	}
	return witness.MarshalBinary()
}

func (db *embedDB) WitnessProfile(publicKey crypto.PublicKey, did string) ([]byte, error) {
	if !db.ExistsProfile(did) {
		return nil, fmt.Errorf("profile with DID %s does not exist", did)
	}

	pk, err := secret.NewKey("profiles", publicKey)
	if err != nil {
		return nil, err
	}

	profiles, err := db.ListProfiles()
	if err != nil {
		return nil, err
	}

	profileIDs := make([]string, len(profiles))
	for i, p := range profiles {
		profileIDs[i] = p.DID
	}

	acc, err := pk.CreateAccumulator(profileIDs...)
	if err != nil {
		return nil, err
	}

	witness, err := pk.CreateWitness(acc, did)
	if err != nil {
		return nil, err
	}
	return witness.MarshalBinary()
}

func (db *embedDB) WitnessWallet(publicKey crypto.PublicKey, did string) ([]byte, error) {
	if !db.ExistsWallet(did) {
		return nil, fmt.Errorf("wallet with DID %s does not exist", did)
	}

	pk, err := secret.NewKey("wallets", publicKey)
	if err != nil {
		return nil, err
	}

	wallets, err := db.ListWallets()
	if err != nil {
		return nil, err
	}

	walletIDs := make([]string, len(wallets))
	for i, w := range wallets {
		walletIDs[i] = w.DID
	}

	acc, err := pk.CreateAccumulator(walletIDs...)
	if err != nil {
		return nil, err
	}

	witness, err := pk.CreateWitness(acc, did)
	if err != nil {
		return nil, err
	}
	return witness.MarshalBinary()
}

func main() {}
