package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	authDtos "event-backend/presentation/http/auth/dtos"
)

type AuthServiceMock struct{}

func (m *AuthServiceMock) Login(ctx context.Context, dto *authDtos.AuthLoginRequestDto) *authDtos.AuthLoginResponseDto {
	return &authDtos.AuthLoginResponseDto{
		Token:     "mock-token-123",
		TokenType: "Bearer",
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}
}

func (m *AuthServiceMock) Register(ctx context.Context, dto *authDtos.AuthRegisterRequestDto) *authDtos.AuthRegisterResponseDto {
	return &authDtos.AuthRegisterResponseDto{
		Id:        uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"),
		Name:      dto.Name,
		Email:     dto.Email,
		CreatedAt: time.Now(),
	}
}

func TestAuthController_Login_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &AuthServiceMock{}
	controller := &AuthController{
		authService: mockService,
	}

	router := gin.New()
	router.POST("/api/v1/auth/login", controller.Login())

	body := `{"email":"test@example.com","password":"password123"}`
	req, _ := http.NewRequest("POST", "/api/v1/auth/login", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, float64(http.StatusOK), response["status_code"])
	assert.NotNil(t, response["data"])
}

func TestAuthController_Register_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &AuthServiceMock{}
	controller := &AuthController{
		authService: mockService,
	}

	router := gin.New()
	router.POST("/api/v1/auth/register", controller.Register())

	body := `{"name":"John Doe","email":"john@example.com","password":"password123"}`
	req, _ := http.NewRequest("POST", "/api/v1/auth/register", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, float64(http.StatusCreated), response["status_code"])
	assert.NotNil(t, response["data"])
}
