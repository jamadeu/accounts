package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jamadeu/accounts/schemas"
	"github.com/jamadeu/accounts/services"
)

type UserHandler struct {
	userRepo schemas.UserRepository
}

func NewUserHandler(ur schemas.UserRepository) *UserHandler {
	return &UserHandler{ur}
}

func (h *UserHandler) RegisterRoutes(router *gin.Engine, basePath string) {
	v1 := router.Group(basePath + "/v1")
	{
		v1.POST("/user", h.handleCreateUser)
		v1.GET("/user", h.handleFindUserById)
	}
}

func (h *UserHandler) handleCreateUser(ctx *gin.Context) {
	var err error
	request := CreateUserRequest{}
	ctx.BindJSON(&request)
	if err = request.Validate(); err != nil {
		fmt.Printf("validation error: %v", err.Error())
		services.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	user := schemas.User{
		Name:      request.Name,
		Document:  request.Document,
		Email:     request.Email,
		AccountID: 0,
	}

	user, err = h.userRepo.Create(&user)
	if err != nil {
		services.SendError(ctx, http.StatusInternalServerError, "creating account on database")
		return
	}
	services.SendSuccess(ctx, "create-user", user)
}

func (h *UserHandler) handleFindUserById(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		services.SendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	user, err := h.userRepo.FindById(id)
	if err != nil {
		services.SendError(ctx, http.StatusNotFound, fmt.Sprintf("user with id: %s not found", id))
		return
	}
	services.SendSuccess(ctx, "find-user-by-id", user)
}
