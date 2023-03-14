package handlers

import (
	"backend/storage"
	"backend/storage/entities"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type QuotesHandler struct {
	storage storage.Storage
}

func NewQuotesHandler(storage storage.Storage, engine *gin.Engine) *QuotesHandler {
	handler := &QuotesHandler{
		storage: storage,
	}

	group := engine.Group("/quotes")

	group.POST("/", handler.AddQuote)

	group.GET("/", handler.GetAllQuotes)
	group.GET("/:id", handler.GetSingleQuote)

	group.PUT("/:id", handler.ApplyActionOnQuote)

	return handler
}

var quoteAddingParamsMissingError = errors.New("text and author are required")
var quoteIdParamError = errors.New("id must be an integer")
var quoteActionParamError = errors.New("valid action is required")

func (handler *QuotesHandler) AddQuote(context *gin.Context) {
	text := context.PostForm("text")
	author := context.PostForm("author")

	if text == "" || author == "" {
		returnError(context, quoteAddingParamsMissingError, http.StatusBadRequest)
		return
	}

	quote := entities.Quote{
		Text:   text,
		Author: author,
	}

	_, err := handler.storage.CreateQuote(quote)
	if err != nil {
		returnServerError(context, err)
		return
	}

	context.Status(http.StatusCreated)
}

func (handler *QuotesHandler) GetAllQuotes(context *gin.Context) {
	quotes, err := handler.storage.GetAllQuotes()
	if err != nil {
		returnServerError(context, err)
		return
	}

	context.JSON(http.StatusOK, quotes)
}

func (handler *QuotesHandler) GetSingleQuote(context *gin.Context) {
	rawId := context.Param("id")

	id, err := strconv.ParseUint(rawId, 10, 64)
	if err != nil {
		returnError(context, quoteIdParamError, http.StatusBadRequest)
		return
	}

	quote, err := handler.storage.GetSingleQuote(uint(id))
	if err != nil {
		if err == storage.QuoteNotFoundError {
			context.Status(http.StatusNotFound)
		} else {
			returnError(context, err, http.StatusInternalServerError)
		}

		return
	}

	context.JSON(http.StatusOK, quote)
}

func (handler *QuotesHandler) ApplyActionOnQuote(context *gin.Context) {
	action := context.Query("action")
	rawId := context.Param("id")

	id, err := strconv.ParseUint(rawId, 10, 64)
	if err != nil {
		returnError(context, quoteIdParamError, http.StatusBadRequest)
	}

	var quote entities.Quote
	switch action {
	case "like":
		err = handler.storage.IncrementQuoteLikes(uint(id))
	case "dislike":
		err = handler.storage.IncrementQuoteDislikes(uint(id))
	default:
		returnError(context, quoteActionParamError, http.StatusBadRequest)
		return
	}

	if err != nil {
		if err == storage.QuoteNotFoundError {
			context.Status(http.StatusNotFound)
		} else {
			returnError(context, err, http.StatusInternalServerError)
		}

		return
	}

	context.JSON(http.StatusOK, quote)
}
