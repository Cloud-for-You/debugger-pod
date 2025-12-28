package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(connStr string) error {
	var err error
	DB, err = sql.Open("postgres", connStr)
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