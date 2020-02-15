package models

import (
	"database/sql"
)

type GetUserScanner struct {
	ID       sql.NullInt64  `db:"id"`
	Username sql.NullString `db:"username"`
	Password sql.NullString `db:"password"`
	FullName sql.NullString `db:"full_name"`
	Picture  sql.NullString `db:"picture"`
}

type GetProfileResult struct {
	ID       int64  `json:"id"`
	UserName string `json:"user_name"`
	FullName string `json:"full_name"`
	Picture  string `json:"picture"`
}


type FormUpdateProfile struct {
	Picture string `json:"picture"`
}