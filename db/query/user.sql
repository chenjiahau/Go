-- name: CreateUser :one
INSERT INTO
users (username, password, full_name) 
values ($1, $2, $3)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username= $1 and  password= $2;

-- name: GetUserById :one
SELECT * FROM users
WHERE id= $1;

-- name: GetUsers :many
SELECT * FROM users;

-- name: Update :one
UPDATE users
SET username = $1, password = $2, full_name = $3
WHERE id = $4
RETURNING *;

-- name: Delete :one
DELETE FROM users
WHERE id = $1
RETURNING *;