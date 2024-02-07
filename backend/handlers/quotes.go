package handlers

import (
	"backend/db"
	"encoding/json"
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

	ctrl.HandleFunc("GET /", ctrl.getAllQuotes)
	ctrl.HandleFunc("GET /{id}", ctrl.getQuote)
	ctrl.HandleFunc("POST /", ctrl.addQuote)
	ctrl.HandleFunc("POST /{id}/like", ctrl.likeQuote)
	ctrl.HandleFunc("POST /{id}/dislike", ctrl.dislikeQuote)

	return ctrl.ServeMux
}

func (ctrl *quotesController) getAllQuotes(w http.ResponseWriter, r *http.Request) {
	quotes, err := ctrl.db.GetQuotes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	returnJSON(w, http.StatusOK, quotes)
}

const invalidQuoteID = "id must be an integer"
const invalidBody = "invalid request body"

func (ctrl *quotesController) getQuote(w http.ResponseWriter, r *http.Request) {
	idRaw := r.PathValue("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		http.Error(w, invalidQuoteID, http.StatusBadRequest)
		return
	}

	quote, err := ctrl.db.GetQuote(uint(id))
	if err != nil {
		var statusCode int
		if errors.Is(err, db.ErrQuoteNotFound) {
			statusCode = http.StatusNotFound
		} else {
			statusCode = http.StatusInternalServerError
		}

		http.Error(w, err.Error(), statusCode)
		return
	}

	returnJSON(w, http.StatusOK, quote)
}

func (ctrl *quotesController) addQuote(w http.ResponseWriter, r *http.Request) {
	var quoteData struct {
		Text   string `json:"text"`
		Author string `json:"author"`
	}

	err := json.NewDecoder(r.Body).Decode(&quoteData)
	if err != nil || quoteData.Text == "" || quoteData.Author == "" {
		http.Error(w, invalidBody, http.StatusBadRequest)
		return
	}

	quote, err := ctrl.db.CreateQuote(quoteData.Text, quoteData.Author)
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
		var statusCode int
		if errors.Is(err, db.ErrQuoteNotFound) {
			statusCode = http.StatusNotFound
		} else {
			statusCode = http.StatusInternalServerError
		}

		http.Error(w, err.Error(), statusCode)
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
		var statusCode int
		if errors.Is(err, db.ErrQuoteNotFound) {
			statusCode = http.StatusNotFound
		} else {
			statusCode = http.StatusInternalServerError
		}

		http.Error(w, err.Error(), statusCode)
		return
	}

	returnJSON(w, http.StatusOK, quote)
}
