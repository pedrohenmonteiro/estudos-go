// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"database/sql"
)

type Category struct {
	ID          string
	Name        sql.NullString
	Description sql.NullString
}

type Course struct {
	ID          string
	Name        sql.NullString
	CategoryID  sql.NullString
	Description sql.NullString
	Price       sql.NullString
}
