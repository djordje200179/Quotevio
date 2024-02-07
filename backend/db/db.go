package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type DB struct {
	conn *sql.DB
}

func New(host, username, password, dbName string) (DB, error) {
	var err error

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		username, password, host, dbName,
	)
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return DB{}, err
	}

	return DB{conn}, nil
}

func (db DB) Close() {
	err := db.conn.Close()
	if err != nil {
		log.Fatal(err)
	}
}
