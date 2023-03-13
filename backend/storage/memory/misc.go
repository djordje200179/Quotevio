package memory

import "backend/models"

func (memory *InMemory) findQuote(id models.QuoteId) int {
	for index, quote := range memory.quotes {
		if quote.Id == id {
			return index
		}
	}

	return -1
}
