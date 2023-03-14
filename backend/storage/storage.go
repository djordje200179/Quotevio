package storage

import (
	"backend/storage/entities"
	"errors"
)

var QuoteNotFoundError = errors.New("quote not found")

type Storage interface {
	CreateQuote(quote entities.Quote) (entities.Quote, error)
	GetSingleQuote(id uint) (entities.Quote, error)
	GetAllQuotes() ([]entities.Quote, error)
	IncrementQuoteLikes(id uint) (entities.Quote, error)
	IncrementQuoteDislikes(id uint) (entities.Quote, error)

	Close()
}
