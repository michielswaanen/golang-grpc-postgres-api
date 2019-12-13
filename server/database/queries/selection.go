package queries

import (
	"database/sql"
	"log"
)

func AuthenticateAccount(db *sql.DB) *sql.Stmt {
	stmt, err := db.Prepare("SELECT name, email FROM accounts WHERE email = $1 AND password = $2")

	checkError(err)
	
	return stmt
}

func IsAccountAvailable(db *sql.DB) *sql.Stmt {
	stmt, err := db.Prepare("SELECT * FROM accounts WHERE email = $1")

	checkError(err)

	return stmt
}

func FetchAccount(db *sql.DB) *sql.Stmt {
	stmt, err := db.Prepare("SELECT name, email, createdAt FROM accounts WHERE id = $1")

	checkError(err)

	return stmt;
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
