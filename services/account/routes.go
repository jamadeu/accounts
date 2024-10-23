package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jamadeu/accounts/types"
)

type AccountHandler struct {
	repo types.AccountRepository
}

func NewAccountHandler(repo types.AccountRepository) *AccountHandler {
	return &AccountHandler{repo: repo}
}

func (ah *AccountHandler) RegisterRoutes(router *gin.Engine, basePath string) {
	v1 := router.Group(basePath)
	{
		v1.POST("/account", ah.handleCreateAccount)
	}
}

func (ah *AccountHandler) handleCreateAccount(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Teste handle",
	})
}
