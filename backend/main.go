package main

import (
	"backend/server"
	"backend/storage/database"
	"log"
	"net"
)

func main() {
	var err error

	err = server.InitLogger("./server.log")
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.New(
		net.IPv4(192, 168, 1, 26),
		"djordje200179", "Djole2001",
		"quotevio",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	srv := server.New(db, 8080)

	err = srv.Start()
	if err != nil {
		return
	}
}
