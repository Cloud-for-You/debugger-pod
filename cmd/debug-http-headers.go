package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Získání všech HTTP hlaviček z požadavku
	headers := r.Header

	// Vytvoření řetězce obsahujícího všechny hlavičky
	var headerStr string
	for key, values := range headers {
		for _, value := range values {
			headerStr += fmt.Sprintf("%s: %s\n", key, value)
		}
	}

	// Odeslání odpovědi s hlavičkami klientovi
	fmt.Fprintf(w, "%s", headerStr)
	
	// Zalogovani na STDOUT
	log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
}
