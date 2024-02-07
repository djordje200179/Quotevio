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

	db, err := db.New("db", "root", "root", "quotes")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.Handle("/health/", http.StripPrefix("/health", handlers.HealthMux()))
	http.Handle("/quotes/", http.StripPrefix("/quotes", handlers.QuotesMux(db)))
	http.HandleFunc("/", handlers.NotFound)

	log.Println("Server is running on port 8080")
	err = http.ListenAndServe(
		":8080",
		middlewares.Log(http.DefaultServeMux),
	)
	if err != nil {
		log.Fatal(err)
	}
}
