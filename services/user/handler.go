package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jamadeu/accounts/schemas"
	s "github.com/jamadeu/accounts/services"
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
		v1.GET("/users", h.handleListUsers)
		v1.PUT("/user", h.handleUpdateUser)
		v1.DELETE("/user", h.handleDeleteUser)
	}
}

func (h *UserHandler) handleCreateUser(ctx *gin.Context) {
	var err error
	request := CreateUserRequest{}
	ctx.BindJSON(&request)
	if err = request.Validate(); err != nil {
		fmt.Printf("validation error: %v", err.Error())
		s.SendError(ctx, http.StatusBadRequest, err.Error())
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
		s.SendError(ctx, http.StatusInternalServerError, "creating account on database")
		return
	}
	s.SendSuccess(ctx, "create-user", user)
}

func (h *UserHandler) handleFindUserById(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		s.SendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	user, err := h.userRepo.FindById(id)
	if err != nil {
		s.SendError(ctx, http.StatusNotFound, fmt.Sprintf("user with id: %s not found", id))
		return
	}
	s.SendSuccess(ctx, "find-user-by-id", user)
}

func (h *UserHandler) handleListUsers(ctx *gin.Context) {
	users, err := h.userRepo.ListUsers()
	if err != nil {
		s.SendError(ctx, http.StatusInternalServerError, "error to list users")
	}
	s.SendSuccess(ctx, "list-users", users)
}

func (h *UserHandler) handleUpdateUser(ctx *gin.Context) {
	request := UpdateUserRequest{}
	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		s.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	id := ctx.Query("id")
	if id == "" {
		s.SendError(ctx, http.StatusBadRequest, errParamIsRequired("id",
			"queryParameter").Error())
		return
	}
	user, err := h.userRepo.FindById(id)
	if err != nil {
		s.SendError(ctx, http.StatusNotFound, fmt.Sprintf("user with id: %s not found", id))
		return
	}
	if request.Name != "" {
		user.Name = request.Name
	}
	if request.Document != "" {
		user.Document = request.Document
	}
	if request.Email != "" {
		user.Email = request.Email
	}

	if err = h.userRepo.Update(user); err != nil {
		s.SendError(ctx, http.StatusInternalServerError, "error updating user")
		return
	}
	s.SendSuccess(ctx, "update-user", user)
}

func (h *UserHandler) handleDeleteUser(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		s.SendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	user, err := h.userRepo.FindById(id)
	if err != nil {
		s.SendError(ctx, http.StatusNotFound, fmt.Sprintf("user with id: %s not found", id))
		return
	}
	err = h.userRepo.Delete(user)
	if err != nil {
		s.SendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleteing car with id: %s", id))
		return
	}
	s.SendSuccess(ctx, "delete-user", fmt.Sprintf("id: %s", id))
}
