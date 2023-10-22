package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

func JsonErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(map[string]string{"message": message}); err != nil {
		log.Fatal(err)
	}
}
