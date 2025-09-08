package utils

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func Measure(w http.ResponseWriter, r *http.Request) {
	// Serialize the request to calculate its length
	var buf bytes.Buffer
	if err := r.Write(&buf); err != nil {
		log.Printf("Failed to serialize request: %v", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	// Get the serialized request and calculate its length
	requestBytes := buf.Bytes()
	requestLength := len(requestBytes)
	log.Printf("Received request, protocol: %s", r.Proto)
	log.Printf("Full Request:\n%s", string(requestBytes))
	log.Printf("Total Request Length: %d bytes", requestLength)

	// Respond to the client
	fmt.Fprintf(w, "Request received, length: %d bytes", requestLength)
}
