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
	categoryDtos "event-backend/presentation/http/category/dtos"
)

type CategoryServiceMock struct{}

func (m *CategoryServiceMock) Pagination(ctx context.Context, dto *categoryDtos.CategoryQueryRequestDto) *infradtos.PaginationResultDto[entities.CategoryEntity] {
	return &infradtos.PaginationResultDto[entities.CategoryEntity]{
		Data:  []*entities.CategoryEntity{},
		Count: 0,
	}
}

func (m *CategoryServiceMock) Detail(ctx context.Context, id uuid.UUID) *entities.CategoryEntity {
	return &entities.CategoryEntity{
		Id:        id,
		Name:      "Test Category",
		Slug:      "test-category",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (m *CategoryServiceMock) Create(ctx context.Context, dto *categoryDtos.CategoryCreateRequestDto) *entities.CategoryEntity {
	return &entities.CategoryEntity{
		Id:        uuid.New(),
		Name:      dto.Name,
		Slug:      "test-category",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (m *CategoryServiceMock) Update(ctx context.Context, dto *categoryDtos.CategoryUpdateRequestDto) *entities.CategoryEntity {
	return &entities.CategoryEntity{
		Id:        dto.Id,
		Name:      dto.Name,
		Slug:      "updated-category",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (m *CategoryServiceMock) SoftDelete(ctx context.Context, id uuid.UUID) {}

func TestCategoryController_Pagination_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &CategoryServiceMock{}
	controller := &CategoryController{
		categoryService: mockService,
	}

	router := gin.New()
	router.GET("/api/v1/categories", controller.Pagination())

	req, _ := http.NewRequest("GET", "/api/v1/categories?page=1&per_page=10", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, float64(http.StatusOK), response["status_code"])
}

func TestCategoryController_Detail_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &CategoryServiceMock{}
	controller := &CategoryController{
		categoryService: mockService,
	}

	router := gin.New()
	router.GET("/api/v1/categories/:id", controller.Detail())

	categoryID := uuid.New()
	req, _ := http.NewRequest("GET", "/api/v1/categories/"+categoryID.String(), nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, float64(http.StatusOK), response["status_code"])
}

func TestCategoryController_Create_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &CategoryServiceMock{}
	controller := &CategoryController{
		categoryService: mockService,
	}

	router := gin.New()
	router.POST("/api/v1/categories", controller.Create())

	body := `{"name":"Test Category","slug":"test-category"}`
	req, _ := http.NewRequest("POST", "/api/v1/categories", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, float64(http.StatusCreated), response["status_code"])
}

func TestCategoryController_Update_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &CategoryServiceMock{}
	controller := &CategoryController{
		categoryService: mockService,
	}

	router := gin.New()
	router.PUT("/api/v1/categories/:id", controller.Update())

	categoryID := uuid.New()
	body := `{"name":"Updated Category","slug":"updated-category"}`
	req, _ := http.NewRequest("PUT", "/api/v1/categories/"+categoryID.String(), bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, float64(http.StatusOK), response["status_code"])
}

func TestCategoryController_Delete_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &CategoryServiceMock{}
	controller := &CategoryController{
		categoryService: mockService,
	}

	router := gin.New()
	router.DELETE("/api/v1/categories/:id", controller.Delete())

	categoryID := uuid.New()
	req, _ := http.NewRequest("DELETE", "/api/v1/categories/"+categoryID.String(), nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, float64(http.StatusOK), response["status_code"])
}
