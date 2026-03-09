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

	userRoute := router.Group("/api/v1/users", middlewares.AuthorizationMiddleware())

	userAdminRoute := userRoute.Group("", guards.RoleGuard([]string{roleConstants.ADMIN}))
	userAdminRoute.GET("", controller.Pagination())
	userAdminRoute.GET("/:id", controller.Detail())
	userAdminRoute.POST("", middlewares.ValidateRequestJson[userDtos.UserCreateRequestDto](), controller.Create())
	userAdminRoute.PUT("/:id", middlewares.ValidateRequestJson[userDtos.UserUpdateRequestDto](), controller.Update())
	userAdminRoute.DELETE("/:id", controller.Delete())

	userAttendeeRoute := userRoute.Group("", guards.RoleGuard([]string{roleConstants.ATTENDEE}))
	userAttendeeRoute.POST("/upgrade-to-organizer", controller.UpgradeToOrganizer())
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
		dto := httpContext.MustGet(middlewares.RequestBodyJsonKey).(*userDtos.UserUpdateRequestDto)
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
		dto := httpContext.MustGet(middlewares.RequestBodyJsonKey).(*userDtos.UserCreateRequestDto)
		result := controller.userService.Create(ctx, dto)
		response := utils.SuccessResponse(http.StatusCreated, userConstants.USER_CREATE_SUCCESS, userDtos.UserDetailResponseDtoFromEntity(result))

		httpContext.JSON(http.StatusCreated, response)
	}
}

func (controller *UserController) UpgradeToOrganizer() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		userId := utils.GetAuthUserId(httpContext)
		controller.userService.UpgradeToOrganizer(ctx, userId)
		response := utils.SuccessResponse(http.StatusOK, userConstants.USER_UPGRADE_TO_ORGANIZER_SUCCESS, nil)

		httpContext.JSON(http.StatusOK, response)
	}
}
