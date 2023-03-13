package main

import (
	"backend/server"
	"backend/storage/database"
	"log"
)

func main() {
	var err error

	err = server.InitLogger("./server.log")
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.New("./database.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	srv := server.New(db, 80)

	err = srv.Start()
	if err != nil {
		return
	}
}
