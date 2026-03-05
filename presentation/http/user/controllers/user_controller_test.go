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

	"event-backend/entities"
	infradtos "event-backend/infrastructure/dtos"
	userDtos "event-backend/presentation/http/user/dtos"
)

type UserServiceMock struct{}

func (m *UserServiceMock) Create(ctx context.Context, dto *userDtos.UserCreateRequestDto) *entities.UserEntity {
	return &entities.UserEntity{
		Id:        uuid.New(),
		Name:      dto.Name,
		Email:     dto.Email,
		RoleId:    dto.RoleId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (m *UserServiceMock) Pagination(ctx context.Context, dto *userDtos.UserQueryRequestDto) *infradtos.PaginationResultDto[entities.UserEntity] {
	return &infradtos.PaginationResultDto[entities.UserEntity]{
		Data:  []*entities.UserEntity{},
		Count: 0,
	}
}

func (m *UserServiceMock) Detail(ctx context.Context, id uuid.UUID) *entities.UserEntity {
	return &entities.UserEntity{
		Id:        id,
		Name:      "Test User",
		Email:     "test@example.com",
		RoleId:    uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (m *UserServiceMock) Update(ctx context.Context, dto *userDtos.UserUpdateRequestDto) *entities.UserEntity {
	roleId := uuid.New()
	if dto.RoleId != nil {
		roleId = *dto.RoleId
	}
	return &entities.UserEntity{
		Id:        dto.Id,
		Name:      dto.Name,
		Email:     dto.Email,
		RoleId:    roleId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (m *UserServiceMock) Delete(ctx context.Context, id uuid.UUID) {}

func TestUserController_Pagination_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &UserServiceMock{}
	controller := &UserController{
		userService: mockService,
	}

	router := gin.New()
	router.GET("/api/v1/users", controller.Pagination())

	req, _ := http.NewRequest("GET", "/api/v1/users?page=1&per_page=10", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, float64(http.StatusOK), response["status_code"])
}

func TestUserController_Detail_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &UserServiceMock{}
	controller := &UserController{
		userService: mockService,
	}

	router := gin.New()
	router.GET("/api/v1/users/:id", controller.Detail())

	userID := uuid.New()
	req, _ := http.NewRequest("GET", "/api/v1/users/"+userID.String(), nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, float64(http.StatusOK), response["status_code"])
}

func TestUserController_Create_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &UserServiceMock{}
	controller := &UserController{
		userService: mockService,
	}

	router := gin.New()
	router.POST("/api/v1/users", controller.Create())

	body := `{"name":"Test User","email":"test@example.com","role_id":"` + uuid.New().String() + `"}`
	req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, float64(http.StatusCreated), response["status_code"])
}

func TestUserController_Update_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &UserServiceMock{}
	controller := &UserController{
		userService: mockService,
	}

	router := gin.New()
	router.PUT("/api/v1/users/:id", controller.Update())

	userID := uuid.New()
	body := `{"name":"Updated User","email":"updated@example.com","role_id":"` + uuid.New().String() + `"}`
	req, _ := http.NewRequest("PUT", "/api/v1/users/"+userID.String(), bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, float64(http.StatusOK), response["status_code"])
}

func TestUserController_Delete_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &UserServiceMock{}
	controller := &UserController{
		userService: mockService,
	}

	router := gin.New()
	router.DELETE("/api/v1/users/:id", controller.Delete())

	userID := uuid.New()
	req, _ := http.NewRequest("DELETE", "/api/v1/users/"+userID.String(), nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, float64(http.StatusOK), response["status_code"])
}
