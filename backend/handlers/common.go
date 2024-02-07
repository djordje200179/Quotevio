package handlers

import (
	"encoding/json"
	"net/http"
)

func returnString(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "text/plain")
	_, err := w.Write([]byte(message))
	if err != nil {
		panic(err)
	}
}

func returnJSON(w http.ResponseWriter, statusCode int, data any) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err)
	}
}
