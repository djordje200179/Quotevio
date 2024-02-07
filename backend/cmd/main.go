package main

import (
	"backend/database"
	"backend/handlers"
	"log"
	"net/http"
)

func main() {
	var err error

	db, err := database.New("db", "root", "root", "quotes", log.Default())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.Handle("/health/", http.StripPrefix("/health", handlers.HealthMux()))
	http.Handle("/quotes/", http.StripPrefix("/quotes", handlers.QuotesMux(db)))

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
