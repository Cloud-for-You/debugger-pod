package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Získání všech HTTP hlaviček z požadavku
	headers := r.Header

  // Převod hlaviček do formátu JSON
	headerMap := make(map[string][]string)
	for key, values := range headers {
		headerMap[key] = values
	}

  // Nastavení typu obsahu odpovědi na JSON
	w.Header().Set("Content-Type", "application/json")

	// Konverze mapy do JSON
	jsonResponse, err := json.Marshal(headerMap)
	if err != nil {
		http.Error(w, "Error converting headers to JSON", http.StatusInternalServerError)
		return
	}

	// Odeslání odpovědi s hlavičkami klientovi
	w.Write(jsonResponse)
	
	// Zalogovani na STDOUT
	log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
}
