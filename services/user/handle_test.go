package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/jamadeu/accounts/schemas"
	"github.com/stretchr/testify/assert"
)

var today = time.Now()
var userTest = schemas.User{
	Model: gorm.Model{
		ID:        1,
		CreatedAt: today,
		UpdatedAt: today,
	},
	Name:     "Test",
	Document: "12345678901",
	Email:    "test@test.com",
}

func getUserTestString() string {
	b, err := json.Marshal(userTest)
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
			"\"data\":" + getUserTestString() + "," +
			"\"message\":\"operation from handler: find-user-by-id successfull\"" +
			"}"

		assert.Equal(t, http.StatusOK, w.Code)
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
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Equal(t, expectedResponseBody, w.Body.String())
	})

	t.Run("should handle return a list of users", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/api/v1/users", nil)
		if err != nil {
			t.Fatal(err)
		}
		router.ServeHTTP(w, req)

		expectedResponseBody := "{\"data\":[" + getUserTestString() + "]," +
			"\"message\":\"operation from handler: list-users successfull\"}"
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, expectedResponseBody, w.Body.String())
	})

	t.Run("should return created user", func(t *testing.T) {
		w := httptest.NewRecorder()
		payload := CreateUserRequest{
			Name:     userTest.Name,
			Document: userTest.Document,
			Email:    userTest.Email,
		}
		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/api/v1/user", bytes.NewBuffer(b))
		if err != nil {
			t.Fatal(err)
		}
		router.ServeHTTP(w, req)

		expectedResponseBody := "{\"data\":" + getUserTestString() + "," +
			"\"message\":\"operation from handler: create-user successfull\"}"
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, expectedResponseBody, w.Body.String())
	})

	t.Run("should return 400 when request boddy is empty", func(t *testing.T) {
		w := httptest.NewRecorder()
		b, err := json.Marshal("{}")
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest("POST", "/api/v1/user", bytes.NewBuffer(b))
		if err != nil {
			t.Fatal(err)
		}
		router.ServeHTTP(w, req)

		expectedResponseBody := "{\"errorCode\":400,\"message\":\"reqest body is empty or malformed\"}"
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, expectedResponseBody, w.Body.String())
	})

	t.Run("should return 400 when name is empty", func(t *testing.T) {
		w := httptest.NewRecorder()
		payload := CreateUserRequest{
			Document: userTest.Document,
			Email:    userTest.Email,
		}
		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/api/v1/user", bytes.NewBuffer(b))
		if err != nil {
			t.Fatal(err)
		}
		router.ServeHTTP(w, req)

		expectedResponseBody := "{\"errorCode\":400,\"message\":\"param: name (type: string) is required\"}"
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, expectedResponseBody, w.Body.String())
	})

	t.Run("should return 400 when document is empty", func(t *testing.T) {
		w := httptest.NewRecorder()
		payload := CreateUserRequest{
			Name:  userTest.Name,
			Email: userTest.Email,
		}
		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/api/v1/user", bytes.NewBuffer(b))
		if err != nil {
			t.Fatal(err)
		}
		router.ServeHTTP(w, req)

		expectedResponseBody := "{\"errorCode\":400,\"message\":\"param: document (type: string) is required\"}"
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, expectedResponseBody, w.Body.String())
	})

	t.Run("should return 400 when email is empty", func(t *testing.T) {
		w := httptest.NewRecorder()
		payload := CreateUserRequest{
			Name:     userTest.Name,
			Document: userTest.Document,
		}
		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/api/v1/user", bytes.NewBuffer(b))
		if err != nil {
			t.Fatal(err)
		}
		router.ServeHTTP(w, req)

		expectedResponseBody := "{\"errorCode\":400,\"message\":\"param: email (type: string) is required\"}"
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, expectedResponseBody, w.Body.String())
	})

	t.Run("should return 400 when email is invalid", func(t *testing.T) {
		w := httptest.NewRecorder()
		payload := CreateUserRequest{
			Name:     userTest.Name,
			Document: userTest.Document,
			Email:    "invalid email",
		}
		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/api/v1/user", bytes.NewBuffer(b))
		if err != nil {
			t.Fatal(err)
		}
		router.ServeHTTP(w, req)

		expectedResponseBody := "{\"errorCode\":400,\"message\":\"param: email (type: string) is required\"}"
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, expectedResponseBody, w.Body.String())
	})

}

type mockUserRepository struct{}

func (m *mockUserRepository) FindById(id string) (*schemas.User, error) {
	if id == "1" {
		return &userTest, nil
	} else {
		return nil, errors.New("user not found")
	}
}

func (m *mockUserRepository) ListUsers() (*[]schemas.User, error) {
	listUser := []schemas.User{userTest}
	return &listUser, nil
}
func (m *mockUserRepository) Create(user *schemas.User) (schemas.User, error) {
	return userTest, nil
}
func (m *mockUserRepository) Update(user *schemas.User) error {
	return nil
}
func (m *mockUserRepository) Delete(user *schemas.User) error {
	return nil
}
