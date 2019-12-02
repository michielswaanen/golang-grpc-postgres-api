package queries

import (
	"database/sql"
	"log"
)

func AuthenticateAccount(db *sql.DB) *sql.Stmt {
	stmt, err := db.Prepare("SELECT * FROM account WHERE email = $1 AND password = $2")

	if err != nil {
		log.Fatal(err)
	}
	
	return stmt
}
