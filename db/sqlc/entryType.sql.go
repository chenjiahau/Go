// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: entryType.sql

package db

import (
	"context"
)

const createEntryType = `-- name: CreateEntryType :one
INSERT INTO 
entry_types (name) 
values ($1)
RETURNING id, name, created_at
`

func (q *Queries) CreateEntryType(ctx context.Context, name string) (EntryTypes, error) {
	row := q.db.QueryRowContext(ctx, createEntryType, name)
	var i EntryTypes
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}

const deleteAllEntryTypes = `-- name: DeleteAllEntryTypes :many
DELETE FROM entry_types
RETURNING id, name, created_at
`

func (q *Queries) DeleteAllEntryTypes(ctx context.Context) ([]EntryTypes, error) {
	rows, err := q.db.QueryContext(ctx, deleteAllEntryTypes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []EntryTypes
	for rows.Next() {
		var i EntryTypes
		if err := rows.Scan(&i.ID, &i.Name, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const deleteEntryType = `-- name: DeleteEntryType :one
DELETE FROM entry_types
WHERE id = $1
RETURNING id, name, created_at
`

func (q *Queries) DeleteEntryType(ctx context.Context, id int32) (EntryTypes, error) {
	row := q.db.QueryRowContext(ctx, deleteEntryType, id)
	var i EntryTypes
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}

const getEntryType = `-- name: GetEntryType :one
SELECT id, name, created_at FROM entry_types
WHERE id= $1
`

func (q *Queries) GetEntryType(ctx context.Context, id int32) (EntryTypes, error) {
	row := q.db.QueryRowContext(ctx, getEntryType, id)
	var i EntryTypes
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}

const getEntryTypeById = `-- name: GetEntryTypeById :one
SELECT id, name, created_at FROM entry_types
WHERE id= $1
`

func (q *Queries) GetEntryTypeById(ctx context.Context, id int32) (EntryTypes, error) {
	row := q.db.QueryRowContext(ctx, getEntryTypeById, id)
	var i EntryTypes
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}

const getEntryTypes = `-- name: GetEntryTypes :many
SELECT id, name, created_at FROM entry_types
`

func (q *Queries) GetEntryTypes(ctx context.Context) ([]EntryTypes, error) {
	rows, err := q.db.QueryContext(ctx, getEntryTypes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []EntryTypes
	for rows.Next() {
		var i EntryTypes
		if err := rows.Scan(&i.ID, &i.Name, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateEntryType = `-- name: UpdateEntryType :one
UPDATE entry_types
SET name = $1
WHERE id = $2
RETURNING id, name, created_at
`

type UpdateEntryTypeParams struct {
	Name string `db:"name"`
	ID   int32  `db:"id"`
}

func (q *Queries) UpdateEntryType(ctx context.Context, arg UpdateEntryTypeParams) (EntryTypes, error) {
	row := q.db.QueryRowContext(ctx, updateEntryType, arg.Name, arg.ID)
	var i EntryTypes
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}
