package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"go-25-27/dto"
	"go-25-27/model"
	"go-25-27/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthHandler_Login(t *testing.T) {
	mockAuth := new(service.MockServiceAuth)

	svc := service.Service{
		AuthService: mockAuth,
	}

	h := NewAuthHandler(svc)

	t.Run("invalid request body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer([]byte("{invalid-json")))
		w := httptest.NewRecorder()

		h.Login(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("validation failed", func(t *testing.T) {
		loginReq := dto.LoginRequest{
			Email:    "",
			Password: "password",
		}
		bodyBytes, _ := json.Marshal(loginReq)

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(bodyBytes))
		w := httptest.NewRecorder()

		h.Login(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("auth service returns error", func(t *testing.T) {
		// Reset mock calls before subtest
		mockAuth.ExpectedCalls = nil
		mockAuth.Calls = nil

		loginReq := dto.LoginRequest{
			Email:    "user@example.com",
			Password: "wrongpass",
		}

		mockAuth.
			On("Login", loginReq.Email, loginReq.Password).
			Return(nil, errors.New("invalid credentials")).Once()

		bodyBytes, _ := json.Marshal(loginReq)

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(bodyBytes))
		w := httptest.NewRecorder()

		h.Login(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		mockAuth.AssertExpectations(t)
	})

	t.Run("login success", func(t *testing.T) {
		// Reset mock calls before subtest
		mockAuth.ExpectedCalls = nil
		mockAuth.Calls = nil

		loginReq := dto.LoginRequest{
			Email:    "user@example.com",
			Password: "password123",
		}

		mockUser := &model.User{
			Email: "user@example.com",
			Name:  "Test User",
		}

		mockAuth.
			On("Login", loginReq.Email, loginReq.Password).
			Return(mockUser, nil).Once()

		bodyBytes, _ := json.Marshal(loginReq)

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(bodyBytes))
		w := httptest.NewRecorder()

		h.Login(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		// decode response
		var responseBody map[string]interface{}
		err := json.NewDecoder(resp.Body).Decode(&responseBody)
		assert.NoError(t, err)

		// t.Logf("Response body: %#v", responseBody)

		assert.Equal(t, "Login success", responseBody["message"])

		data, ok := responseBody["data"].(map[string]interface{})
		assert.True(t, ok, "data should be a map")

		// Check correct field names (lowercase)
		assert.Equal(t, "user@example.com", data["email"])
		assert.Equal(t, "Test User", data["name"])

		mockAuth.AssertExpectations(t)
	})
}
