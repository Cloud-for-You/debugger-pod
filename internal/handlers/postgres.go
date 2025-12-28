package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Clkoud-for-You/debugger-pod/internal/db"
)

func DBConnectHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{}

	err := db.PingDB()
	if err != nil {
		response["status"] = "failed"
		response["message"] = err.Error()
		log.Printf("DB ping failed: %v", err)
	} else {
		response["status"] = "success"
		response["message"] = "Connection to DB was established"
		log.Printf("DB ping success")
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error converting response to JSON", http.StatusInternalServerError)
		return
	}
	w.Write(jsonResponse)
}