package main

import (
	"log"
	"net/http"

	"market-exchange.com/utils"
)

func main() {
	// Define a handler to process and measure the request
	http.HandleFunc("/", utils.Measure)

	// Start the HTTP server on port 8080
	log.Println("Starting server on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
