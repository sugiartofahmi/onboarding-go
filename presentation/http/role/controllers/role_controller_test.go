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
	roleDtos "event-backend/presentation/http/role/dtos"
)

type RoleServiceMock struct{}

func (m *RoleServiceMock) Pagination(ctx context.Context, dto *roleDtos.RoleQueryRequestDto) *infradtos.PaginationResultDto[entities.RoleEntity] {
	return &infradtos.PaginationResultDto[entities.RoleEntity]{
		Data:  []*entities.RoleEntity{},
		Count: 0,
	}
}

func (m *RoleServiceMock) Detail(ctx context.Context, id uuid.UUID) *entities.RoleEntity {
	return &entities.RoleEntity{
		Id:        id,
		Name:      "admin",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (m *RoleServiceMock) Create(ctx context.Context, dto *roleDtos.RoleCreateRequestDto) *entities.RoleEntity {
	return &entities.RoleEntity{
		Id:        uuid.New(),
		Name:      dto.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (m *RoleServiceMock) Update(ctx context.Context, dto *roleDtos.RoleUpdateRequestDto) *entities.RoleEntity {
	return &entities.RoleEntity{
		Id:        dto.Id,
		Name:      dto.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (m *RoleServiceMock) SoftDelete(ctx context.Context, id uuid.UUID) {}

func TestRoleController_Pagination_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &RoleServiceMock{}
	controller := &RoleController{
		roleService: mockService,
	}

	router := gin.New()
	router.GET("/api/v1/roles", controller.Pagination())

	req, _ := http.NewRequest("GET", "/api/v1/roles?page=1&per_page=10", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, float64(http.StatusOK), response["status_code"])
}

func TestRoleController_Detail_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &RoleServiceMock{}
	controller := &RoleController{
		roleService: mockService,
	}

	router := gin.New()
	router.GET("/api/v1/roles/:id", controller.Detail())

	roleID := uuid.New()
	req, _ := http.NewRequest("GET", "/api/v1/roles/"+roleID.String(), nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, float64(http.StatusOK), response["status_code"])
}

func TestRoleController_Create_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &RoleServiceMock{}
	controller := &RoleController{
		roleService: mockService,
	}

	router := gin.New()
	router.POST("/api/v1/roles", controller.Create())

	body := `{"name":"admin"}`
	req, _ := http.NewRequest("POST", "/api/v1/roles", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, float64(http.StatusCreated), response["status_code"])
}

func TestRoleController_Update_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &RoleServiceMock{}
	controller := &RoleController{
		roleService: mockService,
	}

	router := gin.New()
	router.PUT("/api/v1/roles/:id", controller.Update())

	roleID := uuid.New()
	body := `{"name":"organizer"}`
	req, _ := http.NewRequest("PUT", "/api/v1/roles/"+roleID.String(), bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, float64(http.StatusOK), response["status_code"])
}

func TestRoleController_Delete_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &RoleServiceMock{}
	controller := &RoleController{
		roleService: mockService,
	}

	router := gin.New()
	router.DELETE("/api/v1/roles/:id", controller.Delete())

	roleID := uuid.New()
	req, _ := http.NewRequest("DELETE", "/api/v1/roles/"+roleID.String(), nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, float64(http.StatusOK), response["status_code"])
}
