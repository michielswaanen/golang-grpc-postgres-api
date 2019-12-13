package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

const (
	DB_USER = "postgres"
	DB_PASSWORD = "1234"
	DB_NAME = "postgres"
	DB_SSLMODE = "disable"
	DB_HOST = "localhost"
)

var Connection *sql.DB

func NewConnection() {
	credentials := getConnectionCredentials()

	db, err := sql.Open("postgres", credentials)

	if err != nil {
		panic(err.Error())
	}

	db.SetMaxOpenConns(2)
	Connection = db
}

func getConnectionCredentials() string {
	return "host=" + DB_HOST +  " dbname=" + DB_NAME + " user=" + DB_USER + " password=" + DB_PASSWORD + " sslmode=" + DB_SSLMODE
}

