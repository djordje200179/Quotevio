package database

import (
	"backend/models"
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

func (db database) AddQuote(quote models.Quote) (models.QuoteId, error) {
	statement := `
		INSERT INTO quotes
	    (text, author) VALUES (?, ?)
		RETURNING Id
    `
	row := db.QueryRow(statement, quote.Text, quote.Author)

	var nextId models.QuoteId
	err := row.Scan(&nextId)
	if err != nil {
		return 0, err
	}

	return nextId, nil
}

func (db database) GetSingleQuote(id models.QuoteId) (models.Quote, error) {
	statement := `
		SELECT * 
		FROM Quotes 
		WHERE id = ?
	`
	row := db.QueryRow(statement, id)

	return readQuote(row)
}

func (db database) GetAllQuotes() ([]models.Quote, error) {
	count, err := db.getNumberOfQuotes()
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return []models.Quote{}, nil
	}

	statement := `
		SELECT * 
		FROM Quotes
	`
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	quotes := make([]models.Quote, count)
	i := 0
	for rows.Next() {
		err := rows.Scan(&quotes[i].Id, &quotes[i].Text, &quotes[i].Author, &quotes[i].Created, &quotes[i].Likes, &quotes[i].Dislikes)
		if err != nil {
			return nil, err
		}
		i++
	}

	return quotes, nil
}

func (db database) IncrementQuoteLikes(id models.QuoteId) (models.Quote, error) {
	statement := `
		UPDATE Quotes
		SET Likes = Likes + 1
		WHERE id = ?
		RETURNING *
	`
	row := db.QueryRow(statement, id)

	return readQuote(row)
}

func (db database) IncrementQuoteDislikes(id models.QuoteId) (models.Quote, error) {
	statement := `
		UPDATE Quotes
		SET Dislikes = Dislikes + 1
		WHERE id = ?
		RETURNING *
	`
	row := db.QueryRow(statement, id)

	return readQuote(row)
}

func (db database) Close() {
	err := db.DB.Close()
	if err != nil {
		log.Panicln(err)
	}
}
