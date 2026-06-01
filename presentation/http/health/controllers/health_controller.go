package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	healthInterfaces "event-backend/app/health/interfaces"
	"event-backend/infrastructure/utils"
)

type HealthController struct {
	healthService healthInterfaces.HealthServiceInterface
}

func NewHealthController(router *gin.Engine, healthService healthInterfaces.HealthServiceInterface) {
	controller := &HealthController{
		healthService: healthService,
	}

	router.GET("/health", controller.Check())
}

func (ctrl *HealthController) Check() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		result := ctrl.healthService.Check()
		response := utils.SuccessResponse(http.StatusOK, "Service is running", result)

		httpContext.JSON(http.StatusOK, response)
	}
}
