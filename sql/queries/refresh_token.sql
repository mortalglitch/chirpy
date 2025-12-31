-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens (token, created_at, updated_at, user_id, expires_at, revoked_at)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6
)
RETURNING *;

-- name: GetUserFromRefreshToken :one
SELECT * FROM refresh_tokens
WHERE token = $1 LIMIT 1;

-- name: MarkRefreshRevoked :exec
UPDATE refresh_tokens
SET updated_at = $1, revoked_at = $1
WHERE token = $2;
