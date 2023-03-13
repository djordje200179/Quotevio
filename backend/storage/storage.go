package storage

import (
	"backend/models"
	"errors"
)

var QuoteNotFoundError = errors.New("quote not found")

type Storage interface {
	AddQuote(quote models.Quote) (models.QuoteId, error)

	GetSingleQuote(id models.QuoteId) (models.Quote, error)
	GetAllQuotes() ([]models.Quote, error)

	IncrementQuoteLikes(id models.QuoteId) (models.Quote, error)
	IncrementQuoteDislikes(id models.QuoteId) (models.Quote, error)

	Close()
}
