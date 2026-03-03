package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	roleConstants "event-backend/app/role/constants"
	userConstants "event-backend/app/user/constants"
	userInterfaces "event-backend/app/user/interfaces"
	"event-backend/infrastructure/guards"
	"event-backend/infrastructure/middlewares"
	"event-backend/infrastructure/utils"
	uuidValidator "event-backend/infrastructure/validators"
	userDtos "event-backend/presentation/http/user/dtos"
)

type UserController struct {
	userService userInterfaces.UserServiceInterface
}

func NewUserController(router *gin.Engine, userService userInterfaces.UserServiceInterface) {
	controller := &UserController{
		userService: userService,
	}

	userRoute := router.Group("/api/v1/users", middlewares.AuthorizationMiddleware(), guards.RoleGuard([]string{roleConstants.ADMIN}))
	userRoute.GET("", controller.Pagination())
	userRoute.GET("/:id", controller.Detail())
	userRoute.POST("", controller.Create())
	userRoute.PUT("/:id", controller.Update())
	userRoute.DELETE("/:id", controller.Delete())
}

func (controller *UserController) Pagination() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		dto := userDtos.AssignUserQueryRequestDto(httpContext)
		result := controller.userService.Pagination(ctx, dto)
		meta := utils.PaginationMetaBuilder(dto.Page, dto.PerPage, int(result.Count))
		items := userDtos.UserResponseDtoFromEntities(result.Data)
		response := utils.SuccessResponsePagination(http.StatusOK, userConstants.USER_PAGINATION_SUCCESS, items, *meta)

		httpContext.JSON(http.StatusOK, response)
	}
}

func (controller *UserController) Detail() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		id := uuidValidator.ValidateUUID(httpContext.Param("id"))
		result := controller.userService.Detail(ctx, id)
		response := utils.SuccessResponse(http.StatusOK, userConstants.USER_DETAIL_SUCCESS, userDtos.UserDetailResponseDtoFromEntity(result))

		httpContext.JSON(http.StatusOK, response)
	}
}

func (controller *UserController) Update() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		id := uuidValidator.ValidateUUID(httpContext.Param("id"))
		dto := &userDtos.UserUpdateRequestDto{}
		httpContext.ShouldBindJSON(dto)
		dto.Id = id
		result := controller.userService.Update(ctx, dto)
		response := utils.SuccessResponse(http.StatusOK, userConstants.USER_UPDATE_SUCCESS, userDtos.UserDetailResponseDtoFromEntity(result))

		httpContext.JSON(http.StatusOK, response)
	}
}

func (controller *UserController) Delete() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		id := uuidValidator.ValidateUUID(httpContext.Param("id"))
		controller.userService.Delete(ctx, id)
		response := utils.SuccessResponse(http.StatusOK, userConstants.USER_DELETE_SUCCESS, nil)

		httpContext.JSON(http.StatusOK, response)
	}
}

func (controller *UserController) Create() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		dto := &userDtos.UserCreateRequestDto{}
		httpContext.ShouldBindJSON(dto)
		result := controller.userService.Create(ctx, dto)
		response := utils.SuccessResponse(http.StatusCreated, userConstants.USER_CREATE_SUCCESS, userDtos.UserDetailResponseDtoFromEntity(result))

		httpContext.JSON(http.StatusCreated, response)
	}
}
