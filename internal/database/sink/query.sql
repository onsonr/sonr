-- name: InsertCredential :one
INSERT INTO credentials (
    handle,
    credential_id,
    origin,
    type,
    transports
) VALUES (?, ?, ?, ?, ?)
RETURNING *;

-- name: InsertProfile :one
INSERT INTO profiles (
    address,
    handle,
    origin,
    name
) VALUES (?, ?, ?, ?)
RETURNING *;

-- name: GetProfileByID :one
SELECT * FROM profiles
WHERE id = ? AND deleted_at IS NULL
LIMIT 1;

-- name: GetProfileByAddress :one
SELECT * FROM profiles
WHERE address = ? AND deleted_at IS NULL
LIMIT 1;

-- name: GetChallengeBySessionID :one
SELECT challenge FROM sessions
WHERE id = ? AND deleted_at IS NULL
LIMIT 1;

-- name: GetHumanVerificationNumbers :one
SELECT is_human_first, is_human_last FROM sessions
WHERE id = ? AND deleted_at IS NULL
LIMIT 1;

-- name: GetSessionByID :one
SELECT * FROM sessions
WHERE id = ? AND deleted_at IS NULL
LIMIT 1;

-- name: GetSessionByClientIP :one
SELECT * FROM sessions
WHERE client_ipaddr = ? AND deleted_at IS NULL
LIMIT 1;

-- name: UpdateSessionHumanVerification :one
UPDATE sessions
SET 
    is_human_first = ?,
    is_human_last = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?
RETURNING *;

-- name: UpdateSessionWithProfileID :one
UPDATE sessions
SET 
    profile_id = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?
RETURNING *;

-- name: CheckHandleExists :one
SELECT COUNT(*) > 0 as handle_exists FROM profiles 
WHERE handle = ? 
AND deleted_at IS NULL;

-- name: GetCredentialsByHandle :many
SELECT * FROM credentials
WHERE handle = ?
AND deleted_at IS NULL;

-- name: GetCredentialByID :one
SELECT * FROM credentials
WHERE credential_id = ?
AND deleted_at IS NULL
LIMIT 1;

-- name: SoftDeleteCredential :exec
UPDATE credentials
SET deleted_at = CURRENT_TIMESTAMP
WHERE credential_id = ?;

-- name: SoftDeleteProfile :exec
UPDATE profiles
SET deleted_at = CURRENT_TIMESTAMP
WHERE address = ?;

-- name: UpdateProfile :one
UPDATE profiles
SET 
    name = ?,
    handle = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE address = ? 
AND deleted_at IS NULL
RETURNING *;

-- name: GetProfileByHandle :one
SELECT * FROM profiles
WHERE handle = ? 
AND deleted_at IS NULL
LIMIT 1;

-- name: GetVaultConfigByCID :one
SELECT * FROM vaults
WHERE cid = ? 
AND deleted_at IS NULL
LIMIT 1;

-- name: GetVaultRedirectURIBySessionID :one
SELECT redirect_uri FROM vaults
WHERE session_id = ? 
AND deleted_at IS NULL
LIMIT 1;

-- name: CreateSession :one
INSERT INTO sessions (
    id,
    browser_name,
    browser_version,
    client_ipaddr,
    platform,
    is_desktop,
    is_mobile,
    is_tablet,
    is_tv,
    is_bot,
    challenge,
    is_human_first,
    is_human_last,
    profile_id
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )
RETURNING *;
