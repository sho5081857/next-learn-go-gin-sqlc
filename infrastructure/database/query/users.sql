-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;
-- name: CreateUser :one
INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING *;
-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;
