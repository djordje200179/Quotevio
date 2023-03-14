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

	group.GET("/", handler.GetAllQuotes)
	group.POST("/", handler.AddQuote)

	group.GET("/:id", handler.GetSingleQuote)
	group.POST("/:id/like", handler.LikeQuote)
	group.POST("/:id/dislike", handler.DislikeQuote)

	return handler
}

var quoteAddingParamsMissingError = errors.New("text and author are required")
var quoteIdParamError = errors.New("id must be an integer")

func (handler *QuotesHandler) AddQuote(context *gin.Context) {
	var quote entities.Quote

	err := context.ShouldBindJSON(&quote)
	if err != nil {
		returnError(context, quoteAddingParamsMissingError, http.StatusBadRequest)
		return
	}

	quote, err = handler.storage.CreateQuote(quote)
	if err != nil {
		returnServerError(context, err)
		return
	}

	context.JSON(http.StatusCreated, quote)
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

func (handler *QuotesHandler) LikeQuote(context *gin.Context) {
	rawId := context.Param("id")

	id, err := strconv.ParseUint(rawId, 10, 64)
	if err != nil {
		returnError(context, quoteIdParamError, http.StatusBadRequest)
		return
	}

	quote, err := handler.storage.IncrementQuoteLikes(uint(id))
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

func (handler *QuotesHandler) DislikeQuote(context *gin.Context) {
	rawId := context.Param("id")

	id, err := strconv.ParseUint(rawId, 10, 64)
	if err != nil {
		returnError(context, quoteIdParamError, http.StatusBadRequest)
		return
	}

	quote, err := handler.storage.IncrementQuoteDislikes(uint(id))
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
