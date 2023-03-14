package storage

import (
	"backend/storage/entities"
	"errors"
)

var QuoteNotFoundError = errors.New("quote not found")

type Storage interface {
	CreateQuote(quote entities.Quote) (uint, error)
	GetSingleQuote(id uint) (entities.Quote, error)
	GetAllQuotes() ([]entities.Quote, error)
	IncrementQuoteLikes(id uint) error
	IncrementQuoteDislikes(id uint) error

	Close()
}
