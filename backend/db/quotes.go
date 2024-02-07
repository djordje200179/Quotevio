package db

import (
	"backend/entities"
	"database/sql"
	"errors"
)

var ErrQuoteNotFound = errors.New("quote not found")

func (db DB) readQuote(sc interface{ Scan(...interface{}) error }, quote *entities.Quote) error {
	err := sc.Scan(&quote.Id, &quote.Text, &quote.Author, &quote.CreatedAt, &quote.Likes, &quote.Dislikes)
	if err != nil {
		return err
	}

	return nil
}

func (db DB) getQuotesCount() (uint, error) {
	stmt := `
		SELECT COUNT(*) 
		FROM quotes
	`
	row := db.conn.QueryRow(stmt)

	var cnt uint
	err := row.Scan(&cnt)
	if err != nil {
		return 0, err
	}

	return cnt, nil
}

func (db DB) CreateQuote(text, author string) (entities.Quote, error) {
	stmt := `
		INSERT INTO quotes
	    (text, author) VALUES (?, ?)
    `
	res, err := db.conn.Exec(stmt, text, author)
	if err != nil {
		return entities.Quote{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return entities.Quote{}, err
	}

	return db.GetQuote(uint(id))
}

func (db DB) GetQuote(id uint) (entities.Quote, error) {
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
			err = ErrQuoteNotFound
		}

		return entities.Quote{}, err
	}

	return quote, nil
}

func (db DB) GetQuotes() ([]entities.Quote, error) {
	cnt, err := db.getQuotesCount()
	if err != nil {
		return nil, err
	}

	if cnt == 0 {
		return []entities.Quote{}, nil
	}

	stmt := `
		SELECT * 
		FROM quotes
	`
	rows, err := db.conn.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	quotes := make([]entities.Quote, cnt)
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
	stmt := `
		UPDATE quotes
		SET likes = likes + 1
		WHERE id = ?
	`

	_, err := db.conn.Exec(stmt, id)
	if err != nil {
		return entities.Quote{}, err
	}

	return db.GetQuote(id)
}

func (db DB) IncrementQuoteDislikes(id uint) (entities.Quote, error) {
	stmt := `
		UPDATE quotes
		SET dislikes = dislikes + 1
		WHERE id = ?
	`

	_, err := db.conn.Exec(stmt, id)
	if err != nil {
		return entities.Quote{}, err
	}

	return db.GetQuote(id)
}
