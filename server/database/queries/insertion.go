package queries

import (
	"database/sql"
)

func RegisterAccount(db *sql.DB) *sql.Stmt {
	stmt, err := db.Prepare("INSERT INTO accounts (name, email, password) VALUES ($1, $2, $3) RETURNING id")

	checkError(err)

	return stmt
}
