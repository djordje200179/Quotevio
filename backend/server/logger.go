package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

func InitLogger(filePath string) error {
	gin.DisableConsoleColor()

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	gin.DefaultWriter = io.MultiWriter(file)

	return nil
}

func logger(params gin.LogFormatterParams) string {
	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		params.ClientIP,
		params.TimeStamp.Format(time.RFC1123),
		params.Method,
		params.Path,
		params.Request.Proto,
		params.StatusCode,
		params.Latency,
		params.Request.UserAgent(),
		params.ErrorMessage,
	)
}
