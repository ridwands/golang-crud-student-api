package infrastructure

import (
	"database/sql"
	"log"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/golang_api")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
