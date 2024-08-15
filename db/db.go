package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDatabase() *sql.DB {

	db_url := os.Getenv("DATABASE_URL")
	// connect to database
	db, err := sql.Open("postgres", db_url)
	if err != nil {
		log.Fatal(err)
	}
	// check connection
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Printf("Connected to database")

	return db
}
