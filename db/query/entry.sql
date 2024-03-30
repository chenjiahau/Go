-- name: CreateEntry :one
INSERT INTO entries 
(entry_id, type_id, release_version, priority) 
values ($1, $2, $3, $4)
RETURNING *;

-- name: GetEntry :one
SELECT * FROM entries
WHERE id= $1;

-- name: GetEntries :many
SELECT * FROM entries;

-- name: UpdateEntry :one
UPDATE entries
SET release_version = $1, priority = $2
WHERE id = $3
RETURNING *;

-- name: DeleteEntry :one
DELETE FROM entries
WHERE id = $1
RETURNING *;