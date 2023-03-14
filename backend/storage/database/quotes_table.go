package database

import (
	"backend/storage"
	"backend/storage/entities"
	"database/sql"
)

type quoteScanner interface {
	Scan(destinations ...interface{}) error
}

func readQuote(scanner quoteScanner, quote *entities.Quote) error {
	err := scanner.Scan(&quote.Id, &quote.Text, &quote.Author, &quote.Created, &quote.Likes, &quote.Dislikes)
	if err != nil {
		if err == sql.ErrNoRows {
			err = storage.QuoteNotFoundError
		}

		return err
	}

	return nil
}

func (db database) getNumberOfQuotes() (uint, error) {
	statement := `
		SELECT COUNT(*) 
		FROM quotes
	`
	row := db.QueryRow(statement)

	var count uint
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (db database) CreateQuote(quote entities.Quote) (entities.Quote, error) {
	statement := `
		INSERT INTO quotes
	    (text, author) VALUES (?, ?)
		RETURNING *
    `
	row := db.QueryRow(statement, quote.Text, quote.Author)

	err := readQuote(row, &quote)
	if err != nil {
		return entities.Quote{}, err
	}

	return quote, nil
}

func (db database) GetSingleQuote(id uint) (entities.Quote, error) {
	statement := `
		SELECT * 
		FROM quotes 
		WHERE id = ?
	`
	row := db.QueryRow(statement, id)

	var quote entities.Quote
	err := readQuote(row, &quote)
	if err != nil {
		return entities.Quote{}, err
	}

	return quote, nil
}

func (db database) GetAllQuotes() ([]entities.Quote, error) {
	count, err := db.getNumberOfQuotes()
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return []entities.Quote{}, nil
	}

	statement := `
		SELECT * 
		FROM quotes
	`
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	quotes := make([]entities.Quote, count)
	i := 0
	for rows.Next() {
		err := readQuote(rows, &quotes[i])
		if err != nil {
			return nil, err
		}

		i++
	}

	return quotes, nil
}

func (db database) IncrementQuoteLikes(id uint) (entities.Quote, error) {
	statement := `
		UPDATE quotes
		SET likes = likes + 1
		WHERE id = ?
		RETURNING *
	`

	row := db.QueryRow(statement, id)

	var quote entities.Quote
	err := readQuote(row, &quote)
	if err != nil {
		if err == sql.ErrNoRows {
			err = storage.QuoteNotFoundError
		}

		return entities.Quote{}, err
	}

	return quote, nil
}

func (db database) IncrementQuoteDislikes(id uint) (entities.Quote, error) {
	statement := `
		UPDATE quotes
		SET dislikes = dislikes + 1
		WHERE id = ?
		RETURNING *
	`

	row := db.QueryRow(statement, id)

	var quote entities.Quote
	err := readQuote(row, &quote)
	if err != nil {
		if err == sql.ErrNoRows {
			err = storage.QuoteNotFoundError
		}

		return entities.Quote{}, err
	}

	return quote, nil
}
