package main

import (
	"encoding/json"
	"log"
	"net/http"
)

/*
JSON response wrapper for success response
*/
func respondWithJSON(w http.ResponseWriter, httpStatusCode int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	w.Write(data)
}

/*
JSON response wrapper for error response
*/
func respondWithError(w http.ResponseWriter, httpStatusCode int, message string) {
	if httpStatusCode >= 500 {
		log.Fatalf("Error from server: %s", message)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, httpStatusCode, errResponse{Error: message})
}
