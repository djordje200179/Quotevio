package memory

import (
	"backend/models"
	"backend/storage"
)

type InMemory struct {
	quotes []models.Quote
}

func New() storage.Storage {
	return &InMemory{make([]models.Quote, 0)}
}

func (memory *InMemory) AddQuote(quote models.Quote) (models.QuoteId, error) {
	nextId := memory.quotes[len(memory.quotes)-1].Id + 1
	quote.Id = nextId

	memory.quotes = append(memory.quotes, quote)

	return nextId, nil
}

func (memory *InMemory) GetSingleQuote(id models.QuoteId) (models.Quote, error) {
	index := memory.findQuote(id)
	if index == -1 {
		return models.Quote{}, storage.QuoteNotFoundError
	}

	return memory.quotes[index], nil
}

func (memory *InMemory) GetAllQuotes() ([]models.Quote, error) {
	return memory.quotes, nil
}

func (memory *InMemory) IncrementQuoteLikes(id models.QuoteId) (models.Quote, error) {
	index := memory.findQuote(id)
	if index == -1 {
		return models.Quote{}, storage.QuoteNotFoundError
	}

	memory.quotes[index].Likes++

	return memory.quotes[index], nil
}

func (memory *InMemory) IncrementQuoteDislikes(id models.QuoteId) (models.Quote, error) {
	index := memory.findQuote(id)
	if index == -1 {
		return models.Quote{}, storage.QuoteNotFoundError
	}

	memory.quotes[index].Dislikes++

	return memory.quotes[index], nil
}

func (memory *InMemory) Close() {
	memory.quotes = nil
}
