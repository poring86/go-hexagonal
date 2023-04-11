package db

import (
	"database/sql"
	"log"
)

var Db *sql.DB

func setUp() {
	// Db :=
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products ("id" string, "name" string, "price" float, "status" string);`
	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}
