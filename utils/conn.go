package utils

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func connCheck(db *sql.DB) (*sql.DB, error) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}

func ConnDB(DbEngine, DbSource string) (db *sql.DB, err error) {
	db, _ = sql.Open(DbEngine, DbSource)
	db, err = connCheck(db)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
