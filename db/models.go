// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
)

type Todo struct {
	ID      int64          `db:"id"`
	Title   sql.NullString `db:"title"`
	Tags    []string       `db:"tags"`
	Content sql.NullString `db:"content"`
}
