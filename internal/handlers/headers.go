package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func HeaderHandler(w http.ResponseWriter, r *http.Request) {
	// Nastavení typu obsahu odpovědi na JSON
	w.Header().Set("Content-Type", "application/json")

	// Zpracování hlaviček pomocí package
	jsonResponse, err := processHeaders(r)
	if err != nil {
		http.Error(w, "Error converting headers to JSON", http.StatusInternalServerError)
		return
	}

	// Odeslání odpovědi s hlavičkami klientovi
	w.Write(jsonResponse)

	// Zalogovani na STDOUT
	log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
}

func processHeaders(r *http.Request) ([]byte, error) {
	headers := r.Header

	headerMap := make(map[string][]string)
	for key, values := range headers {
		headerMap[key] = values
	}

	return json.Marshal(headerMap)
}