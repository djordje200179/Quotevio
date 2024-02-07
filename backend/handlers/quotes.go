package handlers

import (
	"backend/db"
	"errors"
	"net/http"
	"strconv"
)

type quotesController struct {
	db db.DB

	*http.ServeMux
}

func QuotesMux(db db.DB) *http.ServeMux {
	ctrl := &quotesController{
		db: db,

		ServeMux: http.NewServeMux(),
	}

	ctrl.HandleFunc("GET /", ctrl.getQuotes)
	ctrl.HandleFunc("GET /{id}", ctrl.getQuote)
	ctrl.HandleFunc("POST /", ctrl.createQuote)
	ctrl.HandleFunc("POST /{id}/like", ctrl.likeQuote)
	ctrl.HandleFunc("POST /{id}/dislike", ctrl.dislikeQuote)

	return ctrl.ServeMux
}

func (ctrl *quotesController) getQuotes(w http.ResponseWriter, _ *http.Request) {
	quotes, err := ctrl.db.GetQuotes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	returnJSON(w, http.StatusOK, quotes)
}

const invalidQuoteID = "id must be an integer"

func (ctrl *quotesController) getQuote(w http.ResponseWriter, r *http.Request) {
	idRaw := r.PathValue("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		http.Error(w, invalidQuoteID, http.StatusBadRequest)
		return
	}

	quote, err := ctrl.db.GetQuote(uint(id))
	if err != nil {
		if errors.Is(err, db.ErrQuoteNotFound) {
			http.NotFound(w, r)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	returnJSON(w, http.StatusOK, quote)
}

func (ctrl *quotesController) createQuote(w http.ResponseWriter, r *http.Request) {
	var quoteData struct {
		text   string
		author string
	}

	err := readBody(r, &quoteData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	quote, err := ctrl.db.CreateQuote(quoteData.text, quoteData.author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	returnJSON(w, http.StatusCreated, quote)
}

func (ctrl *quotesController) likeQuote(w http.ResponseWriter, r *http.Request) {
	idRaw := r.PathValue("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		http.Error(w, invalidQuoteID, http.StatusBadRequest)
		return
	}

	quote, err := ctrl.db.IncrementQuoteLikes(uint(id))
	if err != nil {
		if errors.Is(err, db.ErrQuoteNotFound) {
			http.NotFound(w, r)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	returnJSON(w, http.StatusOK, quote)
}

func (ctrl *quotesController) dislikeQuote(w http.ResponseWriter, r *http.Request) {
	idRaw := r.PathValue("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		http.Error(w, invalidQuoteID, http.StatusBadRequest)
		return
	}

	quote, err := ctrl.db.IncrementQuoteDislikes(uint(id))
	if err != nil {
		if errors.Is(err, db.ErrQuoteNotFound) {
			http.NotFound(w, r)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	returnJSON(w, http.StatusOK, quote)
}
