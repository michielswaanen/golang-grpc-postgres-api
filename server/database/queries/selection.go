package queries

import (
	"database/sql"
	"log"
)

func AuthenticateAccount(db *sql.DB) *sql.Stmt {
	stmt, err := db.Prepare("SELECT * FROM account WHERE email = $1 AND password = $2")

	checkError(err)
	
	return stmt
}

func IsAccountAvailable(db *sql.DB) *sql.Stmt {
	stmt, err := db.Prepare("SELECT * FROM account WHERE email = $1")

	checkError(err)

	return stmt
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
