-- name: GetUser :one
SELECT * FROM user WHERE id = ? LIMIT 1;

-- name: CreateUser :execresult
INSERT INTO user (username, email, password) VALUES (?, ?, ?);

-- name: GetUserByEmail :one
SELECT * FROM user WHERE email = ? LIMIT 1;
