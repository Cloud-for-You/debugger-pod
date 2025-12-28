package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(connStr string) error {
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	// Configure connection pool
	DB.SetMaxIdleConns(1)           // Minimal idle connections
	DB.SetMaxOpenConns(1)           // Only one open connection
	DB.SetConnMaxLifetime(time.Minute) // Connections last max 1 minute

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