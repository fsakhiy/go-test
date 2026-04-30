-- name: CreateUser :one
INSERT INTO users (username, password) VALUES(?, ?) RETURNING id, username, password, created_at, updated_at;

-- name: FindUserByUsername :one
SELECT * FROM users WHERE username = ? LIMIT 1;