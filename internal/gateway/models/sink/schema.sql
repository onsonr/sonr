
CREATE TABLE credentials (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME,
    handle TEXT NOT NULL,
    credential_id TEXT NOT NULL,
    origin TEXT NOT NULL,
    type TEXT NOT NULL,
    transports TEXT NOT NULL
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
    is_human_last INTEGER NOT NULL
);

CREATE TABLE profiles (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME,
    address TEXT NOT NULL,
    handle TEXT NOT NULL,
    origin TEXT NOT NULL,
    name TEXT NOT NULL,
    cid TEXT NOT NULL
);

-- Add indexes for common query patterns
CREATE INDEX idx_credentials_handle ON credentials(handle);
CREATE INDEX idx_credentials_origin ON credentials(origin);
CREATE INDEX idx_profiles_handle ON profiles(handle);
CREATE INDEX idx_profiles_address ON profiles(address);
CREATE INDEX idx_sessions_deleted_at ON sessions(deleted_at);
CREATE INDEX idx_credentials_deleted_at ON credentials(deleted_at);
CREATE INDEX idx_profiles_deleted_at ON profiles(deleted_at);
