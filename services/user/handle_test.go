package user

import (
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jamadeu/accounts/schemas"
	"github.com/stretchr/testify/assert"
)

var today = time.Now()
var user = schemas.User{
	Model: gorm.Model{
		ID:        1,
		CreatedAt: today,
		UpdatedAt: today,
	},
	Name:     "Test",
	Document: "12345678901",
	Email:    "test@test.com",
}

func getUserString() string {
	b, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func TestUserHandlers(t *testing.T) {
	userRepo := &mockUserRepository{}
	handler := NewUserHandler(userRepo)
	router := gin.Default()
	handler.RegisterRoutes(router, "/api")

	t.Run("should handle get user by ID", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/api/v1/user?id=1", nil)
		if err != nil {
			t.Fatal(err)
		}
		router.ServeHTTP(w, req)

		expectedBody := "{" +
			"\"data\":" + getUserString() + "," +
			"\"message\":\"operation from handler: find-user-by-id successfull\"" +
			"}"

		assert.Equal(t, 200, w.Code)
		assert.NotEmpty(t, w.Body.String())
		assert.Equal(t, expectedBody, w.Body.String())
	})

	t.Run("should handle return 404 if user not found", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/api/v1/user?id=2", nil)
		if err != nil {
			t.Fatal(err)
		}
		router.ServeHTTP(w, req)

		expectedResponseBody := "{\"errorCode\":404,\"message\":\"user with id: 2 not found\"}"
		assert.Equal(t, 404, w.Code)
		assert.Equal(t, expectedResponseBody, w.Body.String())
	})

	t.Run("should handle return a list of users", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/api/v1/users", nil)
		if err != nil {
			t.Fatal(err)
		}
		router.ServeHTTP(w, req)
		expectedResponseBody := "{\"data\":[" + getUserString() + "]," +
			"\"message\":\"operation from handler: list-users successfull\"}"
		assert.Equal(t, 200, w.Code)
		assert.Equal(t, expectedResponseBody, w.Body.String())
	})

}

type mockUserRepository struct{}

func (m *mockUserRepository) FindById(id string) (*schemas.User, error) {
	if id == "1" {
		return &user, nil
	} else {
		return nil, errors.New("user not found")
	}
}

func (m *mockUserRepository) ListUsers() (*[]schemas.User, error) {
	listUser := []schemas.User{user}
	return &listUser, nil
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
