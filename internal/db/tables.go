package db

const (
	createAccountsTable = `
		CREATE TABLE IF NOT EXISTS accounts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			address TEXT NOT NULL UNIQUE,
			public_key BLOB NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`

	createAssetsTable = `
		CREATE TABLE IF NOT EXISTS assets (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			symbol TEXT NOT NULL,
			decimals INTEGER NOT NULL,
			chain_id INTEGER,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (chain_id) REFERENCES chains(id)
		)
	`

	createChainsTable = `
		CREATE TABLE IF NOT EXISTS chains (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			network_id TEXT NOT NULL UNIQUE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`

	createCredentialsTable = `
		CREATE TABLE IF NOT EXISTS credentials (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			handle TEXT NOT NULL,
			controller TEXT NOT NULL,
			attestation_type TEXT NOT NULL,
			origin TEXT NOT NULL,
			credential_id BLOB NOT NULL,
			public_key BLOB NOT NULL,
			transport TEXT NOT NULL,
			sign_count INTEGER NOT NULL,
			user_present BOOLEAN NOT NULL,
			user_verified BOOLEAN NOT NULL,
			backup_eligible BOOLEAN NOT NULL,
			backup_state BOOLEAN NOT NULL,
			clone_warning BOOLEAN NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`

	createProfilesTable = `
		CREATE TABLE IF NOT EXISTS profiles (
			id TEXT PRIMARY KEY,
			subject TEXT NOT NULL,
			controller TEXT NOT NULL,
			origin_uri TEXT,
			public_metadata TEXT,
			private_metadata TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`

	createPropertiesTable = `
		CREATE TABLE IF NOT EXISTS properties (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			profile_id TEXT NOT NULL,
			key TEXT NOT NULL,
			accumulator BLOB NOT NULL,
			property_key BLOB NOT NULL,
			FOREIGN KEY (profile_id) REFERENCES profiles(id)
		)
	`

	createKeysharesTable = `
		CREATE TABLE IF NOT EXISTS keyshares (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			metadata TEXT NOT NULL,
			payloads TEXT NOT NULL,
			protocol TEXT NOT NULL,
			public_key BLOB NOT NULL,
			role INTEGER NOT NULL,
			version INTEGER NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`

	createPermissionsTable = `
		CREATE TABLE IF NOT EXISTS permissions (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			service_id TEXT NOT NULL,
			grants TEXT NOT NULL,
			scopes TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (service_id) REFERENCES services(id)
		)
	`
)
