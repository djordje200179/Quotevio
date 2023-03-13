package database

import (
	"backend/models"
	"backend/storage"
	"database/sql"
)

func readQuote(row *sql.Row) (models.Quote, error) {
	var quote models.Quote
	err := row.Scan(&quote.Id, &quote.Text, &quote.Author, &quote.Created, &quote.Likes, &quote.Dislikes)
	if err != nil {
		if err == sql.ErrNoRows {
			err = storage.QuoteNotFoundError
		}

		return models.Quote{}, err
	}

	return quote, nil
}

func (database Database) getNumberOfQuotes() (uint, error) {
	statement := `
		SELECT COUNT(*) 
		FROM Quotes
	`
	row := database.db.QueryRow(statement)

	var count uint
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
