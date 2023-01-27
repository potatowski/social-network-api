package database

import (
	"database/sql"
	"log"
	"social-api/src/config"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

// Connect open conection with the database and return
func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
