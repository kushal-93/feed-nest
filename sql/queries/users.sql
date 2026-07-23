-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, username, api_key, name)
VALUES ($1, $2, $3, $4, encode(sha256(random()::text::bytea),'hex'), $5)
RETURNING *;

-- name: GetUserByAPIKey :one
SELECT * FROM users where api_key = $1;