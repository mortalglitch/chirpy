
-- name: CreateChirp :one
INSERT INTO chirps (id, created_at, updated_at, body, user_id) 
VALUES ( 
  $1,
  $2,
  $3,
  $4,
  $5
)
RETURNING *;

-- name: DeleteChirps :exec
DELETE FROM chirps;

-- name: GetChirpsByCreated :many
SELECT * FROM chirps
ORDER BY created_at;

-- name: GetChirpByID :one
SELECT * FROM chirps
WHERE ID = $1 LIMIT 1;
