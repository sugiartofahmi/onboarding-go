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
	"event-backend/infrastructure/utils"
	eventDtos "event-backend/presentation/http/event/dtos"
)

type EventServiceMock struct{}

func (m *EventServiceMock) Pagination(ctx context.Context, dto *eventDtos.EventQueryRequestDto) *infradtos.PaginationResultDto[entities.EventEntity] {
	return &infradtos.PaginationResultDto[entities.EventEntity]{
		Data:  []*entities.EventEntity{},
		Count: 0,
	}
}

func (m *EventServiceMock) Detail(ctx context.Context, id uuid.UUID) *entities.EventEntity {
	return &entities.EventEntity{
		Id:              id,
		OrganizerUserId: uuid.New(),
		CategoryId:      uuid.New(),
		Title:           "Test Event",
		Slug:            "test-event",
		StartDate:       time.Now(),
		EndDate:         time.Now().Add(24 * time.Hour),
		Status:          1,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}

func (m *EventServiceMock) Create(ctx context.Context, dto *eventDtos.EventCreateRequestDto) *entities.EventEntity {
	return &entities.EventEntity{
		Id:              uuid.New(),
		OrganizerUserId: uuid.New(),
		CategoryId:      dto.CategoryId,
		Title:           dto.Title,
		Slug:            "test-event",
		StartDate:       dto.StartDate,
		EndDate:         dto.EndDate,
		Status:          1,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}

func (m *EventServiceMock) Update(ctx context.Context, dto *eventDtos.EventUpdateRequestDto) *entities.EventEntity {
	return &entities.EventEntity{
		Id:              dto.Id,
		OrganizerUserId: uuid.New(),
		CategoryId:      dto.CategoryId,
		Title:           dto.Title,
		Slug:            "updated-event",
		StartDate:       dto.StartDate,
		EndDate:         dto.EndDate,
		Status:          1,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}

func (m *EventServiceMock) SoftDelete(ctx context.Context, id uuid.UUID, currentUserId uuid.UUID) {}

func TestEventController_Pagination_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &EventServiceMock{}
	controller := &EventController{
		eventService: mockService,
	}

	router := gin.New()
	router.GET("/api/v1/events", controller.Pagination())

	req, _ := http.NewRequest("GET", "/api/v1/events?page=1&per_page=10", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, float64(http.StatusOK), response["status_code"])
}

func TestEventController_Detail_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &EventServiceMock{}
	controller := &EventController{
		eventService: mockService,
	}

	router := gin.New()
	router.GET("/api/v1/events/:id", controller.Detail())

	eventID := uuid.New()
	req, _ := http.NewRequest("GET", "/api/v1/events/"+eventID.String(), nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, float64(http.StatusOK), response["status_code"])
}

func TestEventController_Create_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &EventServiceMock{}
	controller := &EventController{
		eventService: mockService,
	}

	router := gin.New()

	router.Use(func(c *gin.Context) {
		c.Set("user", utils.JWTUser{
			Id:       uuid.New(),
			Name:     "Test User",
			Email:    "test@example.com",
			RoleId:   uuid.New(),
			RoleName: "organizer",
		})
		c.Next()
	})
	router.POST("/api/v1/events", controller.Create())

	startDate := time.Now().Add(24 * time.Hour)
	endDate := time.Now().Add(48 * time.Hour)
	body := `{"category_id":"` + uuid.New().String() + `","title":"Test Event","start_date":"` + startDate.Format(time.RFC3339) + `","end_date":"` + endDate.Format(time.RFC3339) + `"}`
	req, _ := http.NewRequest("POST", "/api/v1/events", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, float64(http.StatusCreated), response["status_code"])
}

func TestEventController_Update_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &EventServiceMock{}
	controller := &EventController{
		eventService: mockService,
	}

	router := gin.New()
	router.Use(func(c *gin.Context) {
		c.Set("user", utils.JWTUser{
			Id:       uuid.New(),
			Name:     "Test User",
			Email:    "test@example.com",
			RoleId:   uuid.New(),
			RoleName: "organizer",
		})
		c.Next()
	})
	router.PUT("/api/v1/events/:id", controller.Update())

	eventID := uuid.New()
	startDate := time.Now().Add(24 * time.Hour)
	endDate := time.Now().Add(48 * time.Hour)
	body := `{"category_id":"` + uuid.New().String() + `","title":"Updated Event","start_date":"` + startDate.Format(time.RFC3339) + `","end_date":"` + endDate.Format(time.RFC3339) + `"}`
	req, _ := http.NewRequest("PUT", "/api/v1/events/"+eventID.String(), bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, float64(http.StatusOK), response["status_code"])
}

func TestEventController_Delete_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &EventServiceMock{}
	controller := &EventController{
		eventService: mockService,
	}

	router := gin.New()
	router.Use(func(c *gin.Context) {
		c.Set("user", utils.JWTUser{
			Id:       uuid.New(),
			Name:     "Test User",
			Email:    "test@example.com",
			RoleId:   uuid.New(),
			RoleName: "organizer",
		})
		c.Next()
	})
	router.DELETE("/api/v1/events/:id", controller.Delete())

	eventID := uuid.New()
	req, _ := http.NewRequest("DELETE", "/api/v1/events/"+eventID.String(), nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, float64(http.StatusOK), response["status_code"])
}
