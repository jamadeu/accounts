package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
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

var updatedUserTest = userTest

func jsonToString(s interface{}) string {
	b, err := json.Marshal(s)
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

	t.Run("handle find should get user by ID", func(t *testing.T) {
		w := httptest.NewRecorder()
		userId := strconv.Itoa(int(userTest.ID))
		req, err := http.NewRequest("GET", "/api/v1/user?id="+userId, nil)
		if err != nil {
			t.Fatal(err)
		}
		router.ServeHTTP(w, req)

		expectedBody := "{" +
			"\"data\":" + jsonToString(userTest) + "," +
			"\"message\":\"operation from handler: find-user-by-id successfull\"" +
			"}"

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, expectedBody, w.Body.String())
	})

	t.Run("handle find should return 404 when user is not found", func(t *testing.T) {
		w := httptest.NewRecorder()
		userId := "2"
		req, err := http.NewRequest("GET", "/api/v1/user?id="+userId, nil)
		if err != nil {
			t.Fatal(err)
		}
		router.ServeHTTP(w, req)

		expectedResponseBody := "{\"errorCode\":404,\"message\":\"user with id: " + userId + " not found\"}"
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Equal(t, expectedResponseBody, w.Body.String())
	})

	t.Run("handle list should return a list of users", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/api/v1/users", nil)
		if err != nil {
			t.Fatal(err)
		}
		router.ServeHTTP(w, req)

		expectedResponseBody := "{\"data\":[" + jsonToString(userTest) + "]," +
			"\"message\":\"operation from handler: list-users successfull\"}"
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, expectedResponseBody, w.Body.String())
	})

	t.Run("handle create should return created user", func(t *testing.T) {
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

		expectedResponseBody := "{\"data\":" + jsonToString(userTest) + "," +
			"\"message\":\"operation from handler: create-user successfull\"}"
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, expectedResponseBody, w.Body.String())
	})

	t.Run("handle create should return 400 when request boddy is empty", func(t *testing.T) {
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

	t.Run("handle create should return 400 when name is empty", func(t *testing.T) {
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

	t.Run("handle create should return 400 when document is empty", func(t *testing.T) {
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

	t.Run("handle create should return 400 when email is empty", func(t *testing.T) {
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

	t.Run("handle create should return 400 when email is invalid", func(t *testing.T) {
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

	t.Run("handle update should return updated user", func(t *testing.T) {
		w := httptest.NewRecorder()
		payload := UpdateUserRequest{
			Name:     "Updated Name",
			Document: "10987654321",
			Email:    "updated_email@test.com",
		}
		updatedUserTest.Name = payload.Name
		updatedUserTest.Document = payload.Document
		updatedUserTest.Email = payload.Email

		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		userId := strconv.Itoa(int(userTest.ID))
		req, err := http.NewRequest("PUT", "/api/v1/user?id="+userId, bytes.NewBuffer(b))
		if err != nil {
			t.Fatal(err)
		}
		router.ServeHTTP(w, req)

		expectedResponseBody := "{\"data\":" + jsonToString(updatedUserTest) + "," +
			"\"message\":\"operation from handler: update-user successfull\"}"
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, expectedResponseBody, w.Body.String())
	})

	t.Run("should handle update return 400 when user id is empty", func(t *testing.T) {
		w := httptest.NewRecorder()
		payload := UpdateUserRequest{
			Name:     "Updated Name",
			Document: "10987654321",
			Email:    "updated_email@test.com",
		}
		updatedUserTest.Name = payload.Name
		updatedUserTest.Document = payload.Document
		updatedUserTest.Email = payload.Email

		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("PUT", "/api/v1/user", bytes.NewBuffer(b))
		if err != nil {
			t.Fatal(err)
		}
		router.ServeHTTP(w, req)

		expectedResponseBody := "{\"errorCode\":400,\"message\":\"param: id (type: queryParameter) is required\"}"
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, expectedResponseBody, w.Body.String())
	})

	t.Run("handle update should return 404 when user is not found", func(t *testing.T) {
		w := httptest.NewRecorder()
		payload := UpdateUserRequest{
			Name:     "Updated Name",
			Document: "10987654321",
			Email:    "updated_email@test.com",
		}
		updatedUserTest.Name = payload.Name
		updatedUserTest.Document = payload.Document
		updatedUserTest.Email = payload.Email

		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		userId := "2"
		req, err := http.NewRequest("PUT", "/api/v1/user?id="+userId, bytes.NewBuffer(b))
		if err != nil {
			t.Fatal(err)
		}
		router.ServeHTTP(w, req)

		expectedResponseBody := "{\"errorCode\":404,\"message\":\"user with id: " + userId + " not found\"}"
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Equal(t, expectedResponseBody, w.Body.String())
	})

	t.Run("handle update should return 400 when payload is empty", func(t *testing.T) {
		w := httptest.NewRecorder()
		payload := UpdateUserRequest{}

		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		userId := strconv.Itoa(int(userTest.ID))
		req, err := http.NewRequest("PUT", "/api/v1/user?id="+userId, bytes.NewBuffer(b))
		if err != nil {
			t.Fatal(err)
		}
		router.ServeHTTP(w, req)

		expectedResponseBody := "{\"errorCode\":400,\"message\":\"at least one valid field must be provided\"}"
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, expectedResponseBody, w.Body.String())
	})

	t.Run("handle update should return 400 when email is invalid", func(t *testing.T) {
		w := httptest.NewRecorder()
		payload := UpdateUserRequest{
			Email: "invalid email",
		}

		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		userId := strconv.Itoa(int(userTest.ID))
		req, err := http.NewRequest("PUT", "/api/v1/user?id="+userId, bytes.NewBuffer(b))
		if err != nil {
			t.Fatal(err)
		}
		router.ServeHTTP(w, req)

		expectedResponseBody := "{\"errorCode\":400,\"message\":\"param: email (type: string) is required\"}"
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, expectedResponseBody, w.Body.String())
	})

	t.Run("handle delete should user by ID", func(t *testing.T) {
		w := httptest.NewRecorder()
		userId := strconv.Itoa(int(userTest.ID))
		req, err := http.NewRequest("DELETE", "/api/v1/user?id="+userId, nil)
		if err != nil {
			t.Fatal(err)
		}
		router.ServeHTTP(w, req)

		expectedBody := "{" +
			"\"data\":\"id: " + userId + "\"," +
			"\"message\":\"operation from handler: delete-user successfull\"" +
			"}"

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, expectedBody, w.Body.String())
	})

	t.Run("handle delete should return 400 when user id is empty", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, err := http.NewRequest("DELETE", "/api/v1/user", nil)
		if err != nil {
			t.Fatal(err)
		}
		router.ServeHTTP(w, req)

		expectedResponseBody := "{\"errorCode\":400,\"message\":\"param: id (type: queryParameter) is required\"}"
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, expectedResponseBody, w.Body.String())
	})

	t.Run("handle delete should return 404 when user is not found", func(t *testing.T) {
		w := httptest.NewRecorder()
		userId := "2"
		req, err := http.NewRequest("DELETE", "/api/v1/user?id="+userId, nil)
		if err != nil {
			t.Fatal(err)
		}
		router.ServeHTTP(w, req)

		expectedResponseBody := "{\"errorCode\":404,\"message\":\"user with id: " + userId + " not found\"}"
		assert.Equal(t, http.StatusNotFound, w.Code)
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
