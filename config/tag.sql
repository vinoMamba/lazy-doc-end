-- name: GetTag :one
SELECT * FROM tag WHERE id = ? LIMIT 1;
