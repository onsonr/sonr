CREATE TABLE accounts (
    id TEXT PRIMARY KEY,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME,
    number INTEGER NOT NULL,
    sequence INTEGER NOT NULL,
    address TEXT NOT NULL,
    public_key JSON NOT NULL,
    chain_id TEXT NOT NULL,
    controller TEXT NOT NULL,
    is_subsidiary INTEGER NOT NULL,
    is_validator INTEGER NOT NULL,
    is_delegator INTEGER NOT NULL,
    is_accountable INTEGER NOT NULL
);

CREATE TABLE assets (
    id TEXT PRIMARY KEY,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME,
    name TEXT NOT NULL,
    symbol TEXT NOT NULL,
    decimals INTEGER NOT NULL,
    chain_id TEXT NOT NULL,
    channel TEXT NOT NULL,
    asset_type TEXT NOT NULL,
    coingecko_id TEXT NOT NULL
);

CREATE TABLE credentials (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME,
    handle TEXT NOT NULL,
    credential_id TEXT NOT NULL,
    authenticator_attachment TEXT NOT NULL,
    origin TEXT NOT NULL,
    type TEXT NOT NULL,
    transports TEXT NOT NULL
);

CREATE TABLE profiles (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME,
    address TEXT NOT NULL,
    handle TEXT NOT NULL,
    origin TEXT NOT NULL,
    name TEXT NOT NULL
);

CREATE TABLE sessions (
    id TEXT PRIMARY KEY,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME,
    browser_name TEXT NOT NULL,
    browser_version TEXT NOT NULL,
    client_ipaddr TEXT NOT NULL,
    platform TEXT NOT NULL,
    is_desktop INTEGER NOT NULL, -- SQLite doesn't have boolean, using INTEGER (0/1)
    is_mobile INTEGER NOT NULL,
    is_tablet INTEGER NOT NULL,
    is_tv INTEGER NOT NULL,
    is_bot INTEGER NOT NULL,
    challenge TEXT NOT NULL,
    is_human_first INTEGER NOT NULL,
    is_human_last INTEGER NOT NULL,
    profile_id INTEGER NOT NULL
);

CREATE TABLE vaults (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME,
    handle TEXT NOT NULL,
    origin TEXT NOT NULL,
    address TEXT NOT NULL,
    cid TEXT NOT NULL,
    config JSON NOT NULL,
    session_id TEXT NOT NULL,
    redirect_uri TEXT NOT NULL
);

-- Add indexes for common query patterns
CREATE INDEX idx_credentials_handle ON credentials(handle);
CREATE INDEX idx_credentials_origin ON credentials(origin);
CREATE INDEX idx_profiles_handle ON profiles(handle);
CREATE INDEX idx_profiles_address ON profiles(address);
CREATE INDEX idx_sessions_deleted_at ON sessions(deleted_at);
CREATE INDEX idx_credentials_deleted_at ON credentials(deleted_at);
CREATE INDEX idx_profiles_deleted_at ON profiles(deleted_at);
CREATE INDEX idx_vaults_deleted_at ON vaults(deleted_at);
CREATE INDEX idx_sessions_profile_id ON sessions(profile_id);
CREATE INDEX idx_sessions_client_ipaddr ON sessions(client_ipaddr);
