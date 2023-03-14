package database

import (
	"backend/storage"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net"
)

type database struct {
	*sql.DB
}

func New(address net.IP, username, password, databaseName string) (storage.Storage, error) {
	var err error

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		username, password, address, databaseName,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return database{db}, nil
}

func (db database) Close() {
	err := db.DB.Close()
	if err != nil {
		log.Panicln(err)
	}
}
