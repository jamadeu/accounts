package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jamadeu/accounts/schemas"
)

type AccountHandler struct {
	accountRepo    schemas.AccountRepository
	userRepository schemas.UserRepository
}

func NewAccountHandler(ar schemas.AccountRepository, ur schemas.UserRepository) *AccountHandler {
	return &AccountHandler{accountRepo: ar, userRepository: ur}
}

func (ah *AccountHandler) RegisterRoutes(router *gin.Engine, basePath string) {
	v1 := router.Group(basePath)
	{
		v1.POST("/v1/account", ah.handleCreateAccount)
	}
}

func (ah *AccountHandler) handleCreateAccount(ctx *gin.Context) {
	request := CreateAccountRequest{}
	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		// fmt.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	user, err := ah.userRepository.FindById(request.UserId)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, "user not found")
		return
	}
	account := schemas.Account{
		Balance:      request.Balance,
		User:         *user,
		Transactions: []schemas.Transaction{},
	}
	if err := ah.accountRepo.CreateAccount(account); err != nil {
		sendError(ctx, http.StatusInternalServerError, "creating account on database")
		return
	}
	sendSuccess(ctx, "create-account", account)
}
