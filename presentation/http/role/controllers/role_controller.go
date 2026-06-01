package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	roleConstants "event-backend/app/role/constants"
	roleConstants2 "event-backend/app/role/constants"
	roleInterfaces "event-backend/app/role/interfaces"
	"event-backend/infrastructure/guards"
	"event-backend/infrastructure/middlewares"
	"event-backend/infrastructure/utils"
	uuidValidator "event-backend/infrastructure/validators"
	roleDtos "event-backend/app/role/dtos"
)

type RoleController struct {
	roleService roleInterfaces.RoleServiceInterface
}

func NewRoleController(router *gin.Engine, roleService roleInterfaces.RoleServiceInterface) {
	controller := &RoleController{
		roleService: roleService,
	}

	roleRoute := router.Group("/api/v1/roles", middlewares.AuthorizationMiddleware(), guards.RoleGuard([]string{roleConstants.ADMIN}))
	roleRoute.GET("", controller.Pagination())
	roleRoute.GET("/:id", controller.Detail())
	roleRoute.POST("", middlewares.ValidateRequestJson[roleDtos.RoleCreateRequestDto](), controller.Create())
	roleRoute.PUT("/:id", middlewares.ValidateRequestJson[roleDtos.RoleUpdateRequestDto](), controller.Update())
	roleRoute.DELETE("/:id", controller.Delete())
}

func (controller *RoleController) Pagination() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		dto := roleDtos.AssignRoleQueryRequestDto(httpContext)
		result := controller.roleService.Pagination(ctx, dto)
		meta := utils.PaginationMetaBuilder(dto.Page, dto.PerPage, int(result.Count))
		items := roleDtos.RoleResponseDtoFromEntities(result.Data)
		response := utils.SuccessResponsePagination(http.StatusOK, roleConstants.ROLE_PAGINATION_SUCCESS, items, *meta)

		httpContext.JSON(http.StatusOK, response)
	}
}

func (controller *RoleController) Detail() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		id := uuidValidator.ValidateUUID(httpContext.Param("id"))
		result := controller.roleService.Detail(ctx, id)
		response := utils.SuccessResponse(http.StatusOK, roleConstants.ROLE_DETAIL_SUCCESS, roleDtos.RoleDetailResponseDtoFromEntity(*result))

		httpContext.JSON(http.StatusOK, response)
	}
}

func (controller *RoleController) Create() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		dto := httpContext.MustGet(middlewares.RequestBodyJsonKey).(*roleDtos.RoleCreateRequestDto)
		result := controller.roleService.Create(ctx, dto)
		response := utils.SuccessResponse(http.StatusCreated, roleConstants2.ROLE_CREATE_SUCCESS, roleDtos.RoleDetailResponseDtoFromEntity(*result))

		httpContext.JSON(http.StatusCreated, response)
	}
}

func (controller *RoleController) Update() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		id := uuidValidator.ValidateUUID(httpContext.Param("id"))
		dto := httpContext.MustGet(middlewares.RequestBodyJsonKey).(*roleDtos.RoleUpdateRequestDto)
		dto.Id = id
		result := controller.roleService.Update(ctx, dto)
		response := utils.SuccessResponse(http.StatusOK, roleConstants2.ROLE_UPDATE_SUCCESS, roleDtos.RoleDetailResponseDtoFromEntity(*result))

		httpContext.JSON(http.StatusOK, response)
	}
}

func (controller *RoleController) Delete() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		id := uuidValidator.ValidateUUID(httpContext.Param("id"))
		controller.roleService.Delete(ctx, id)
		response := utils.SuccessResponse(http.StatusOK, roleConstants2.ROLE_DELETE_SUCCESS, nil)

		httpContext.JSON(http.StatusOK, response)
	}
}
