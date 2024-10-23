package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jamadeu/accounts/services/account"
	"gorm.io/gorm"
)

const (
	basePath = "api/v1"
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

	accountRepo := account.NewAccountRepository(s.db)
	accountHandler := account.NewAccountHandler(accountRepo)
	accountHandler.RegisterRoutes(router, basePath)

	return router.Run(s.port)
}
