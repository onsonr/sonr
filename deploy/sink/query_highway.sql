-- name: InsertCredential :one
INSERT INTO credentials (
    handle,
    credential_id,
    origin,
    type,
    transports
) VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: InsertProfile :one
INSERT INTO profiles (
    address,
    handle,
    origin,
    name
) VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetProfileByID :one
SELECT * FROM profiles
WHERE id = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: GetProfileByAddress :one
SELECT * FROM profiles
WHERE address = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: GetChallengeBySessionID :one
SELECT challenge FROM sessions
WHERE id = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: GetHumanVerificationNumbers :one
SELECT is_human_first, is_human_last FROM sessions
WHERE id = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: GetSessionByID :one
SELECT * FROM sessions
WHERE id = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: GetSessionByClientIP :one
SELECT * FROM sessions
WHERE client_ipaddr = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: UpdateSessionHumanVerification :one
UPDATE sessions
SET 
    is_human_first = $1,
    is_human_last = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $3
RETURNING *;

-- name: UpdateSessionWithProfileID :one
UPDATE sessions
SET 
    profile_id = $1,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $2
RETURNING *;

-- name: CheckHandleExists :one
SELECT COUNT(*) > 0 as handle_exists FROM profiles 
WHERE handle = $1 
AND deleted_at IS NULL;

-- name: GetCredentialsByHandle :many
SELECT * FROM credentials
WHERE handle = $1
AND deleted_at IS NULL;

-- name: GetCredentialByID :one
SELECT * FROM credentials
WHERE credential_id = $1
AND deleted_at IS NULL
LIMIT 1;

-- name: SoftDeleteCredential :exec
UPDATE credentials
SET deleted_at = CURRENT_TIMESTAMP
WHERE credential_id = $1;

-- name: SoftDeleteProfile :exec
UPDATE profiles
SET deleted_at = CURRENT_TIMESTAMP
WHERE address = $1;

-- name: UpdateProfile :one
UPDATE profiles
SET 
    name = $1,
    handle = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE address = $3 
AND deleted_at IS NULL
RETURNING *;

-- name: GetProfileByHandle :one
SELECT * FROM profiles
WHERE handle = $1 
AND deleted_at IS NULL
LIMIT 1;

-- name: GetVaultConfigByCID :one
SELECT * FROM vaults
WHERE cid = $1 
AND deleted_at IS NULL
LIMIT 1;

-- name: GetVaultRedirectURIBySessionID :one
SELECT redirect_uri FROM vaults
WHERE session_id = $1 
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
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
RETURNING *;
