-- Profiles represent user identities
CREATE TABLE profiles (
    id TEXT PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,
    address TEXT NOT NULL,
    handle TEXT NOT NULL UNIQUE,
    origin TEXT NOT NULL,
    name TEXT NOT NULL,
    status TEXT NOT NULL,
    UNIQUE(address, origin)
);

-- Accounts represent blockchain accounts
CREATE TABLE accounts (
    id TEXT PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,
    number BIGINT NOT NULL,
    sequence INT NOT NULL DEFAULT 0,
    address TEXT NOT NULL UNIQUE,
    public_key TEXT NOT NULL,
    chain_id TEXT NOT NULL,
    controller TEXT NOT NULL,
    is_subsidiary BOOLEAN NOT NULL DEFAULT FALSE,
    is_validator BOOLEAN NOT NULL DEFAULT FALSE,
    is_delegator BOOLEAN NOT NULL DEFAULT FALSE,
    is_accountable BOOLEAN NOT NULL DEFAULT TRUE
);

-- Assets represent tokens and coins
CREATE TABLE assets (
    id TEXT PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,
    name TEXT NOT NULL,
    symbol TEXT NOT NULL,
    decimals INT NOT NULL CHECK(decimals >= 0),
    chain_id TEXT NOT NULL,
    channel TEXT NOT NULL,
    asset_type TEXT NOT NULL,
    coingecko_id TEXT,
    UNIQUE(chain_id, symbol)
);

-- Credentials store WebAuthn credentials
CREATE TABLE credentials (
    id TEXT PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,
    handle TEXT NOT NULL,
    credential_id TEXT NOT NULL UNIQUE,
    authenticator_attachment TEXT NOT NULL,
    origin TEXT NOT NULL,
    type TEXT NOT NULL,
    transports TEXT NOT NULL
);

-- Sessions track user authentication state
CREATE TABLE sessions (
    id TEXT PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,
    browser_name TEXT NOT NULL,
    browser_version TEXT NOT NULL,
    client_ipaddr TEXT NOT NULL,
    platform TEXT NOT NULL,
    is_desktop BOOLEAN NOT NULL DEFAULT FALSE,
    is_mobile BOOLEAN NOT NULL DEFAULT FALSE,
    is_tablet BOOLEAN NOT NULL DEFAULT FALSE,
    is_tv BOOLEAN NOT NULL DEFAULT FALSE,
    is_bot BOOLEAN NOT NULL DEFAULT FALSE,
    challenge TEXT NOT NULL,
    is_human_first BOOLEAN NOT NULL DEFAULT FALSE,
    is_human_last BOOLEAN NOT NULL DEFAULT FALSE,
    profile_id TEXT NOT NULL REFERENCES profiles(id),
    FOREIGN KEY (profile_id) REFERENCES profiles(id)
);

-- Vaults store encrypted data
CREATE TABLE vaults (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,
    handle TEXT NOT NULL,
    origin TEXT NOT NULL,
    address TEXT NOT NULL,
    cid TEXT NOT NULL UNIQUE,
    config JSONB NOT NULL,
    session_id BIGINT NOT NULL,
    redirect_uri TEXT NOT NULL,
    FOREIGN KEY (session_id) REFERENCES sessions(id)
);

-- Indexes for common queries
CREATE INDEX idx_profiles_handle ON profiles(handle);
CREATE INDEX idx_profiles_address ON profiles(address);
CREATE INDEX idx_profiles_origin ON profiles(origin);
CREATE INDEX idx_profiles_status ON profiles(status);
CREATE INDEX idx_profiles_deleted_at ON profiles(deleted_at);

CREATE INDEX idx_accounts_address ON accounts(address);
CREATE INDEX idx_accounts_chain_id ON accounts(chain_id);
CREATE INDEX idx_accounts_deleted_at ON accounts(deleted_at);

CREATE INDEX idx_assets_symbol ON assets(symbol);
CREATE INDEX idx_assets_chain_id ON assets(chain_id);
CREATE INDEX idx_assets_deleted_at ON assets(deleted_at);

CREATE INDEX idx_credentials_handle ON credentials(handle);
CREATE INDEX idx_credentials_origin ON credentials(origin);
CREATE INDEX idx_credentials_deleted_at ON credentials(deleted_at);

CREATE INDEX idx_sessions_profile_id ON sessions(profile_id);
CREATE INDEX idx_sessions_client_ipaddr ON sessions(client_ipaddr);
CREATE INDEX idx_sessions_deleted_at ON sessions(deleted_at);
