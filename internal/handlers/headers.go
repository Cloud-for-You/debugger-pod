package handlers

import (
	"log"
	"net/http"

	"github.com/Clkoud-for-You/debugger-pod/internal/headers"
)

func HeaderHandler(w http.ResponseWriter, r *http.Request) {
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