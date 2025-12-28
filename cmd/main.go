package main

import (
	"log"
	"net/http"

	"github.com/Clkoud-for-You/debugger-pod/internal/db"
	"github.com/Clkoud-for-You/debugger-pod/internal/handlers"
)

func main() {
	err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.CloseDB()

	http.HandleFunc("/", handlers.HeaderHandler)
	http.HandleFunc("/db_connect", handlers.DBConnectHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}