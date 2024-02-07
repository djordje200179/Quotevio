package database

import (
	"backend/entities"
	"database/sql"
	"errors"
)

type quoteScanner interface {
	Scan(destinations ...interface{}) error
}

var ErrQuoteNotFound = errors.New("quote not found")
var ErrSql = errors.New("sql error")

func (db DB) readQuote(scanner quoteScanner, quote *entities.Quote) error {
	err := scanner.Scan(&quote.Id, &quote.Text, &quote.Author, &quote.Likes, &quote.Dislikes, &quote.CreatedAt)
	if err != nil {
		db.logger.Println(err)
		return errors.Join(err, ErrSql)
	}

	return nil
}

func (db DB) getNumberOfQuotes() (uint, error) {
	statement := `
		SELECT COUNT(*) 
		FROM quotes
	`
	row := db.conn.QueryRow(statement)

	var count uint
	err := row.Scan(&count)
	if err != nil {
		db.logger.Println(err)
		return 0, errors.Join(err, ErrSql)
	}

	return count, nil
}

func (db DB) CreateQuote(text, author string) (entities.Quote, error) {
	statement := `
		INSERT INTO quotes
	    (text, author) VALUES (?, ?)
    `
	result, err := db.conn.Exec(statement, text, author)
	if err != nil {
		db.logger.Println(err)
		return entities.Quote{}, errors.Join(err, ErrSql)
	}

	id, err := result.LastInsertId()
	if err != nil {
		db.logger.Println(err)
		return entities.Quote{}, errors.Join(err, ErrSql)
	}

	return db.GetSingleQuote(uint(id))
}

func (db DB) GetSingleQuote(id uint) (entities.Quote, error) {
	statement := `
		SELECT * 
		FROM quotes 
		WHERE id = ?
	`
	row := db.conn.QueryRow(statement, id)

	var quote entities.Quote
	err := db.readQuote(row, &quote)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.Quote{}, ErrQuoteNotFound
		} else {
			return entities.Quote{}, err
		}
	}

	return quote, nil
}

func (db DB) GetAllQuotes() ([]entities.Quote, error) {
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
	rows, err := db.conn.Query(statement)
	if err != nil {
		db.logger.Println(err)
		return nil, errors.Join(err, ErrSql)
	}
	defer rows.Close()

	quotes := make([]entities.Quote, count)
	i := 0
	for rows.Next() {
		err = db.readQuote(rows, &quotes[i])
		if err != nil {
			return nil, err
		}

		i++
	}

	return quotes, nil
}

func (db DB) IncrementQuoteLikes(id uint) (entities.Quote, error) {
	statement := `
		UPDATE quotes
		SET likes = likes + 1
		WHERE id = ?
	`

	_, err := db.conn.Exec(statement, id)
	if err != nil {
		db.logger.Println(err)
		return entities.Quote{}, errors.Join(err, ErrSql)
	}

	return db.GetSingleQuote(id)
}

func (db DB) IncrementQuoteDislikes(id uint) (entities.Quote, error) {
	statement := `
		UPDATE quotes
		SET dislikes = dislikes + 1
		WHERE id = ?
	`

	_, err := db.conn.Exec(statement, id)
	if err != nil {
		db.logger.Println(err)
		return entities.Quote{}, errors.Join(err, ErrSql)
	}

	return db.GetSingleQuote(id)
}
