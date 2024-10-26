package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jamadeu/accounts/services/account"
	"github.com/jamadeu/accounts/services/user"
	"gorm.io/gorm"
)

const (
	basePath = "api"
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

	userRepo := user.NewUserRepository(s.db)
	userHandler := user.NewUserHandler(userRepo)
	userHandler.RegisterRoutes(router, basePath)

	accountRepo := account.NewAccountRepository(s.db)
	accountHandler := account.NewAccountHandler(accountRepo, userRepo)
	accountHandler.RegisterRoutes(router, basePath)

	return router.Run(s.port)
}
