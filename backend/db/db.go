package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

type DB struct {
	conn *sql.DB
}

func Init() DB {
	var err error

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"), os.Getenv("DB_NAME"),
	)
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	return DB{conn: conn}
}

func (db DB) Close() {
	err := db.conn.Close()
	if err != nil {
		log.Fatal(err)
	}
}
