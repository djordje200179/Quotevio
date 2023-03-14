package handlers

import (
	st "backend/storage"
	models2 "backend/storage/entities"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var quoteAddingParamsMissingError = errors.New("text and author are required")
var quoteIdParamError = errors.New("id must be an integer")
var quoteActionParamError = errors.New("valid action is required")

func RegisterQuotes(engine *gin.Engine) {
	group := engine.Group("/quotes")

	group.POST("/", addQuote)

	group.GET("/", getAllQuotes)
	group.GET("/:id", getSingleQuote)

	group.PUT("/:id", applyActionOnQuote)
}

func addQuote(context *gin.Context) {
	text := context.PostForm("text")
	author := context.PostForm("author")

	if text == "" || author == "" {
		returnError(context, quoteAddingParamsMissingError, http.StatusBadRequest)
		return
	}

	quote := models2.Quote{
		Text:    text,
		Author:  author,
		Created: models2.Today(),
	}

	storage := getStorage(context)

	id, err := storage.CreateQuote(quote)
	if err != nil {
		returnServerError(context, err)
		return
	}

	context.Header("Location", "/quotes/"+strconv.FormatUint(uint64(id), 10))
	context.Status(http.StatusCreated)
}

func getAllQuotes(context *gin.Context) {
	storage := getStorage(context)

	quotes, err := storage.GetAllQuotes()
	if err != nil {
		returnServerError(context, err)
		return
	}

	context.JSON(http.StatusOK, quotes)
}

func getSingleQuote(context *gin.Context) {
	rawId := context.Param("id")

	id, err := strconv.ParseUint(rawId, 10, 64)
	if err != nil {
		returnError(context, quoteIdParamError, http.StatusBadRequest)
		return
	}

	storage := getStorage(context)

	quote, err := storage.GetSingleQuote(models2.QuoteId(id))
	if err != nil {
		if err == st.QuoteNotFoundError {
			context.Status(http.StatusNotFound)
		} else {
			returnError(context, err, http.StatusInternalServerError)
		}

		return
	}

	context.JSON(http.StatusOK, quote)
}

func applyActionOnQuote(context *gin.Context) {
	action := context.Query("action")
	rawId := context.Param("id")

	id, err := strconv.ParseUint(rawId, 10, 64)
	if err != nil {
		returnError(context, quoteIdParamError, http.StatusBadRequest)
	}

	storage := getStorage(context)

	var quote models2.Quote
	switch action {
	case "like":
		quote, err = storage.IncrementQuoteLikes(models2.QuoteId(id))
	case "dislike":
		quote, err = storage.IncrementQuoteDislikes(models2.QuoteId(id))
	default:
		returnError(context, quoteActionParamError, http.StatusBadRequest)
		return
	}

	if err != nil {
		if err == st.QuoteNotFoundError {
			context.Status(http.StatusNotFound)
		} else {
			returnError(context, err, http.StatusInternalServerError)
		}

		return
	}

	context.JSON(http.StatusOK, quote)
}
