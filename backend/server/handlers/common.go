package handlers

import (
	"backend/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getStorage(context *gin.Context) storage.Storage {
	return context.MustGet("storage").(storage.Storage)
}

func returnError(context *gin.Context, err error, status int) {
	dict := gin.H{
		"error": err.Error(),
	}

	context.JSON(status, dict)
}

func returnServerError(context *gin.Context, err error) {
	returnError(context, err, http.StatusInternalServerError)
}
