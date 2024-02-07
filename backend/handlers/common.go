package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
)

var errInvalidBody = errors.New("invalid request body")

func readBody[T any](r *http.Request, data *T) error {
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		return errors.Join(errInvalidBody, err)
	}

	return nil
}

func returnString(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(statusCode)
	_, err := w.Write([]byte(message))
	if err != nil {
		panic(err)
	}
}

func returnJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err)
	}
}
