// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql"
)

type Comments struct {
	ID        int32          `db:"id"`
	EntryID   int32          `db:"entry_id"`
	UserID    int32          `db:"user_id"`
	Content   sql.NullString `db:"content"`
	CreatedAt sql.NullTime   `db:"created_at"`
}

type Entries struct {
	ID             int32          `db:"id"`
	EntryTypeID    int32          `db:"entry_type_id"`
	ReleaseVersion sql.NullString `db:"release_version"`
	Priority       sql.NullInt32  `db:"priority"`
	CreatedAt      sql.NullTime   `db:"created_at"`
}

type EntryTypes struct {
	ID        int32        `db:"id"`
	Name      string       `db:"name"`
	CreatedAt sql.NullTime `db:"created_at"`
}

type Users struct {
	ID        int32          `db:"id"`
	Username  string         `db:"username"`
	Password  string         `db:"password"`
	FullName  sql.NullString `db:"full_name"`
	CreatedAt sql.NullTime   `db:"created_at"`
}
