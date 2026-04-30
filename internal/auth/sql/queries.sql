-- name: CreateUser :execresult
INSERT INTO users (username, password) VALUES(?, ?);

-- name: FindUserByUsername :one
SELECT * FROM users WHERE username = ? LIMIT 1;