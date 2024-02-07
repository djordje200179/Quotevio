package handlers

import (
	"net/http"
	"strconv"
	"time"
)

func ping(w http.ResponseWriter, r *http.Request) {
	returnString(w, http.StatusOK, "pong")
}

func timestamp(w http.ResponseWriter, r *http.Request) {
	unixTime := time.Now().UnixMilli()

	returnString(w, http.StatusOK, strconv.Itoa(int(unixTime)))
}

func HealthMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /ping", ping)
	mux.HandleFunc("GET /timestamp", timestamp)

	return mux
}
