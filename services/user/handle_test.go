package user

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jamadeu/accounts/schemas"
	"github.com/stretchr/testify/assert"
)

func TestUserHandlers(t *testing.T) {
	userRepo := &mockUserRepository{}
	handler := NewUserHandler(userRepo)

	router := gin.Default()
	handler.RegisterRoutes(router, "/api")
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)

}

type mockUserRepository struct{}

func (m *mockUserRepository) FindById(id string) (*schemas.User, error) {
	return &schemas.User{}, nil
}

func (m *mockUserRepository) ListUsers() (*[]schemas.User, error) {
	return &[]schemas.User{}, nil
}
func (m *mockUserRepository) Create(user *schemas.User) (schemas.User, error) {
	return schemas.User{}, nil
}
func (m *mockUserRepository) Update(user *schemas.User) error {
	return nil
}
func (m *mockUserRepository) Delete(user *schemas.User) error {
	return nil
}
