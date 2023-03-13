package database

import (
	"backend/models"
	"backend/storage"
	"database/sql"
	"log"
	_ "modernc.org/sqlite"
)

type Database struct {
	db *sql.DB
}

func New(path string) (storage.Storage, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return Database{}, err
	}

	return Database{db}, nil
}

func (database Database) AddQuote(quote models.Quote) (models.QuoteId, error) {
	statement := `
		INSERT INTO Quotes 
	    (Text, Author, Created) VALUES (?, ?, ?)
		RETURNING Id
    `
	row := database.db.QueryRow(statement, quote.Text, quote.Author, quote.Created)

	var nextId models.QuoteId
	err := row.Scan(&nextId)
	if err != nil {
		return 0, err
	}

	return nextId, nil
}

func (database Database) GetSingleQuote(id models.QuoteId) (models.Quote, error) {
	statement := `
		SELECT * 
		FROM Quotes 
		WHERE id = ?
	`
	row := database.db.QueryRow(statement, id)

	return readQuote(row)
}

func (database Database) GetAllQuotes() ([]models.Quote, error) {
	count, err := database.getNumberOfQuotes()
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
	rows, err := database.db.Query(statement)
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

func (database Database) IncrementQuoteLikes(id models.QuoteId) (models.Quote, error) {
	statement := `
		UPDATE Quotes
		SET Likes = Likes + 1
		WHERE id = ?
		RETURNING *
	`
	row := database.db.QueryRow(statement, id)

	return readQuote(row)
}

func (database Database) IncrementQuoteDislikes(id models.QuoteId) (models.Quote, error) {
	statement := `
		UPDATE Quotes
		SET Dislikes = Dislikes + 1
		WHERE id = ?
		RETURNING *
	`
	row := database.db.QueryRow(statement, id)

	return readQuote(row)
}

func (database Database) Close() {
	err := database.db.Close()
	if err != nil {
		log.Panicln(err)
	}
}
