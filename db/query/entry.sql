-- name: CreateEntry :one
INSERT INTO entries 
(entry_type_id, release_version, priority) 
values ($1, $2, $3)
RETURNING *;

-- name: GetEntry :one
SELECT * FROM entries
WHERE id= $1;

-- name: GetEntryById :one
SELECT * FROM entries
WHERE id= $1;

-- name: GetEntries :many
SELECT * FROM entries;

-- name: UpdateEntry :one
UPDATE entries
SET entry_type_id = $1, release_version = $2, priority = $3
WHERE id = $4
RETURNING *;

-- name: DeleteEntry :one
DELETE FROM entries
WHERE id = $1
RETURNING *;

-- name: DeleteAllEntries :many
DELETE FROM entries
RETURNING *;