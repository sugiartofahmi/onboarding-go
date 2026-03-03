package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	authConstants "event-backend/app/auth/constants"
	authInterfaces "event-backend/app/auth/interfaces"
	"event-backend/infrastructure/utils"
	authDtos "event-backend/presentation/http/auth/dtos"
)

type AuthController struct {
	authService authInterfaces.AuthServiceInterface
}

func NewAuthController(router *gin.Engine, authService authInterfaces.AuthServiceInterface) {
	authRoute := router.Group("/api/v1/auth")

	controller := &AuthController{
		authService: authService,
	}

	authRoute.POST("/login", controller.Login())
	authRoute.POST("/register", controller.Register())
}

func (controller *AuthController) Login() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		dto := &authDtos.AuthLoginRequestDto{}
		httpContext.ShouldBindJSON(dto)
		result := controller.authService.Login(ctx, dto)
		response := utils.SuccessResponse(http.StatusOK, authConstants.AUTH_LOGIN_SUCCESS, result)

		httpContext.JSON(http.StatusOK, response)
	}
}

func (controller *AuthController) Register() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		dto := &authDtos.AuthRegisterRequestDto{}
		httpContext.ShouldBindJSON(dto)
		result := controller.authService.Register(ctx, dto)
		response := utils.SuccessResponse(http.StatusCreated, authConstants.AUTH_REGISTER_SUCCESS, result)

		httpContext.JSON(http.StatusCreated, response)
	}
}
