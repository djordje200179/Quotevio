package main

import (
	"backend/db"
	"backend/handlers"
	"backend/middlewares"
	"fmt"
	"net/http"
)

func main() {
	db := db.Init()
	defer db.Close()

	http.Handle("/health/", http.StripPrefix("/health", handlers.HealthMux()))
	http.Handle("/quotes/", http.StripPrefix("/quotes", handlers.QuotesMux(db)))

	handler := http.Handler(http.DefaultServeMux)
	handler = middlewares.Log(handler)
	handler = middlewares.PanicRecover(handler)
	handler = middlewares.Limit(handler)

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(
		":8080",
		handler,
	)
	if err != nil {
		panic(err)
	}
}
