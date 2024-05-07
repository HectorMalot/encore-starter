-- name: GetToken :one
SELECT * FROM tokens WHERE token_hash = $1;

-- name: StoreToken :one
INSERT INTO tokens (user_id, description, token_hash, valid_until) VALUES ($1, $2, $3, $4) RETURNING id;

-- name: ListTokens :many
SELECT * FROM tokens WHERE user_id = $1;

-- name: DeleteToken :exec
DELETE FROM tokens WHERE id = $1;
