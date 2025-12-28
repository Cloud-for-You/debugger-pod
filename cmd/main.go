package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Clkoud-for-You/debugger-pod/internal/db"
	"github.com/Clkoud-for-You/debugger-pod/internal/handlers"
)

func main() {
	connStr := os.Getenv("DB_CONNECTION_STRING")

	http.HandleFunc("/headers", handlers.HeaderHandler)

	if connStr != "" {
		err := db.InitDB(connStr)
		if err != nil {
			log.Fatal(err)
		}
		defer db.CloseDB()
		http.HandleFunc("/db_connect", handlers.DBConnectHandler)
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}