package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type APIServer struct {
	port string
	db   *gorm.DB
}

func NewApiServer(addr string, db *gorm.DB) *APIServer {
	return &APIServer{
		port: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := gin.Default()

	return router.Run(s.port)
}
