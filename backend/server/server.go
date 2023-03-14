package server

import (
	"backend/server/handlers"
	"backend/storage"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port uint

	engine *gin.Engine
}

func newEngine(storage storage.Storage) *gin.Engine {
	engine := gin.New()

	engine.Use(gin.LoggerWithFormatter(logger))
	engine.Use(gin.Recovery())

	handlers.NewQuotesHandler(storage, engine)

	return engine
}

func New(storage storage.Storage, port uint) *Server {
	server := new(Server)

	server.port = port
	server.engine = newEngine(storage)

	return server
}

func (server *Server) Start() error {
	addr := fmt.Sprintf(":%v", server.port)

	err := server.engine.Run(addr)
	return err
}
