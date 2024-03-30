-- name: CreateEntryType :one
INSERT INTO 
entry_types (name) 
values ($1)
RETURNING *;

-- name: GetEntryType :one
SELECT * FROM entry_types
WHERE id= $1;

-- name: GetEntryTypeById :one
SELECT * FROM entry_types
WHERE id= $1;

-- name: GetEntryTypes :many
SELECT * FROM entry_types;

-- name: UpdateEntryType :one
UPDATE entry_types
SET name = $1
WHERE id = $2
RETURNING *;

-- name: DeleteEntryType :one
DELETE FROM entry_types
WHERE id = $1
RETURNING *;

-- name: DeleteAllEntryTypes :many
DELETE FROM entry_types
RETURNING *;