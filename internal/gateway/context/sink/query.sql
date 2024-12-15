-- name: InsertCredential :one
INSERT INTO credentials (
    handle,
    credential_id,
    origin,
    type,
    transports
) VALUES (?, ?, ?, ?, ?)
RETURNING *;

-- name: InsertUser :one
INSERT INTO users (
    address,
    handle,
    origin,
    name
) VALUES (?, ?, ?, ?)
RETURNING *;

-- name: GetUserByAddress :one
SELECT * FROM users
WHERE address = ? AND deleted_at IS NULL
LIMIT 1;

-- name: GetSessionByID :one
SELECT * FROM sessions
WHERE id = ? AND deleted_at IS NULL
LIMIT 1;

-- name: UpdateSessionHumanVerification :one
UPDATE sessions
SET 
    is_human_first = ?,
    is_human_last = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?
RETURNING *;

-- name: CheckHandleExists :one
SELECT COUNT(*) > 0 as handle_exists FROM users 
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

-- name: SoftDeleteUser :exec
UPDATE users
SET deleted_at = CURRENT_TIMESTAMP
WHERE address = ?;

-- name: UpdateUser :one
UPDATE users
SET 
    name = ?,
    handle = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE address = ? 
AND deleted_at IS NULL
RETURNING *;

-- name: GetUserByHandle :one
SELECT * FROM users
WHERE handle = ? 
AND deleted_at IS NULL
LIMIT 1;

-- name: CreateSession :one
INSERT INTO sessions (
    id,
    browser_name,
    browser_version,
    platform,
    is_desktop,
    is_mobile,
    is_tablet,
    is_tv,
    is_bot,
    challenge,
    is_human_first,
    is_human_last
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
    ABS(RANDOM() % 5) + 1,  -- Random number between 1-5
    ABS(RANDOM() % 4) + 1   -- Random number between 1-4
)
RETURNING *;

-- name: InitializeProfile :one
SELECT 
    ABS(RANDOM() % 5) + 1 as first_number,
    ABS(RANDOM() % 4) + 1 as last_number;
