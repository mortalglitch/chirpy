-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email, hashed_password, is_chirpy_red)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6
)
RETURNING *;

-- name: DeleteUsers :exec
DELETE FROM users;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: UpdateUserInfo :exec
UPDATE users
SET email = $1, hashed_password = $2, updated_at = $3
WHERE id = $4;

-- name: UpgradeUserToRed :exec
UPDATE users
SET is_chirpy_red = TRUE, updated_at = $1
WHERE id = $2;
