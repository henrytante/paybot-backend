package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)



func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "H1000:Tadoido987@tcp(localhost:3306)/paybot")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	
	return db, err
}
