package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type DB struct {
	conn   *sql.DB
	logger *log.Logger
}

func New(host, username, password, databaseName string, logger *log.Logger) (DB, error) {
	var err error

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		username, password, host, databaseName,
	)
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return DB{}, err
	}

	return DB{conn, logger}, nil
}

func (db DB) Close() {
	err := db.conn.Close()
	if err != nil {
		db.logger.Println(err)
	}
}
