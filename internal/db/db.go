package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		dbURL = "postgres://user:password@localhost/dbname?sslmode=disable" // default for testing
	}
	var err error
	DB, err = sql.Open("postgres", dbURL)
	if err != nil {
		return err
	}

	// Test connection
	if err = DB.Ping(); err != nil {
		log.Printf("DB connection failed: %v", err)
		return err
	} else {
		log.Println("DB connection established")
	}

	return nil
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}

func PingDB() error {
	if DB == nil {
		return sql.ErrConnDone
	}
	return DB.Ping()
}