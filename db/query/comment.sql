-- name: CreateComment :one
INSERT INTO comments
(entry_id, user_id, content)
values ($1, $2, $3)
RETURNING *;

-- name: GetComment :one
SELECT * FROM comments
WHERE id= $1;

-- name: GetCommentById :one
SELECT * FROM comments
WHERE id= $1;

-- name: GetComments :many
SELECT * FROM comments;

-- name: UpdateComment :one
UPDATE comments
SET content = $1
WHERE id = $2
RETURNING *;

-- name: DeleteComment :one
DELETE FROM comments
WHERE id = $1
RETURNING *;

-- name: DeleteComments :many
DELETE FROM comments
RETURNING *;