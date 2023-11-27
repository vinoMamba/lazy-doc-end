-- name: CreateProject :execresult
INSERT INTO project (project_name, project_description, is_public, is_deleted, created_by) VALUES (?, ?, ?, ?, ?);

-- name: GetProjectList :many
SELECT * FROM project WHERE is_deleted = false ORDER BY created_at DESC LIMIT ? OFFSET ?;
