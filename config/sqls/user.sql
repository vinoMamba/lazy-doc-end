-- name: GetUser :one
SELECT * FROM user WHERE id = ? LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM user WHERE username = ? LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM user WHERE email = ? LIMIT 1;

-- name: CreateUser :execresult
INSERT INTO user (username, email, password) VALUES (?, ?, ?);


-- name: UpdateUsernameById :execresult
UPDATE user SET username = ? WHERE id = ?;

-- name: UpdateEmailById :execresult
UPDATE user SET email = ? WHERE id = ?;
