package database

import "database/sql"

const (
	DB_USER = "postgres"
	DB_PASSWORD = "1234"
	DB_NAME = "postgres"
	DB_SSLMODE = "disable"
	DB_HOST = "localhost"
)

func getConnectionCredentials() string {
	return "host=" + DB_HOST +  "dbname=" + DB_NAME + " user=" + DB_USER + " password=" + DB_PASSWORD + " sslmode=" + DB_SSLMODE
}

func GetConnection() *sql.DB {
	credentials := getConnectionCredentials()

	db, err := sql.Open("postgres", credentials)

	if err != nil {
		panic(err.Error())
	}

	return db
}
