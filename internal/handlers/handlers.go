package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Clkoud-for-You/debugger-pod/internal/db"
	"github.com/Clkoud-for-You/debugger-pod/internal/headers"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Nastavení typu obsahu odpovědi na JSON
	w.Header().Set("Content-Type", "application/json")

	// Zpracování hlaviček pomocí package
	jsonResponse, err := headers.ProcessHeaders(r)
	if err != nil {
		http.Error(w, "Error converting headers to JSON", http.StatusInternalServerError)
		return
	}

	// Odeslání odpovědi s hlavičkami klientovi
	w.Write(jsonResponse)

	// Zalogovani na STDOUT
	log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
}

func DBConnectHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{}

	err := db.PingDB()
	if err != nil {
		response["status"] = "failed"
		response["message"] = err.Error()
		log.Printf("DB ping failed: %v", err)
	} else {
		response["status"] = "success"
		response["message"] = "Spojeni k DB bylo uskutecneni"
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