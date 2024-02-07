package main

import (
	"backend/db"
	"backend/handlers"
	"backend/middlewares"
	"log"
	"net/http"
)

func main() {
	var err error

	db, err := db.New("db", "root", "root", "quotes", log.Default())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.Handle("/health/", http.StripPrefix("/health", handlers.HealthMux()))
	http.Handle("/quotes/", http.StripPrefix("/quotes", handlers.QuotesMux(db)))

	var handler http.Handler
	handler = http.DefaultServeMux
	handler = middlewares.Log(handler, log.Default())

	log.Println("Server is running on port 8080")
	err = http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatal(err)
	}
}
