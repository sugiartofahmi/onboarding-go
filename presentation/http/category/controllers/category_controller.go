package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	categoryConstants "event-backend/app/category/constants"
	categoryInterfaces "event-backend/app/category/interfaces"
	roleConstants "event-backend/app/role/constants"
	"event-backend/infrastructure/guards"
	"event-backend/infrastructure/middlewares"
	"event-backend/infrastructure/utils"
	uuidValidator "event-backend/infrastructure/validators"
	categoryDtos "event-backend/presentation/http/category/dtos"
)

type CategoryController struct {
	categoryService categoryInterfaces.CategoryServiceInterface
}

func NewCategoryController(router *gin.Engine, categoryService categoryInterfaces.CategoryServiceInterface) {
	controller := &CategoryController{
		categoryService: categoryService,
	}

	categoryRoute := router.Group("/api/v1/categories", middlewares.AuthorizationMiddleware())
	categoryRoute.GET("", controller.Pagination())
	categoryRoute.GET("/:id", controller.Detail())

	protected := categoryRoute.Group("", guards.RoleGuard([]string{roleConstants.ADMIN}))
	protected.POST("", middlewares.ValidateRequestJson[categoryDtos.CategoryCreateRequestDto](), controller.Create())
	protected.PUT("/:id", middlewares.ValidateRequestJson[categoryDtos.CategoryUpdateRequestDto](), controller.Update())
	protected.DELETE("/:id", controller.Delete())
}

func (controller *CategoryController) Pagination() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		dto := categoryDtos.AssignCategoryQueryRequestDto(httpContext)
		result := controller.categoryService.Pagination(ctx, dto)
		meta := utils.PaginationMetaBuilder(dto.Page, dto.PerPage, int(result.Count))
		items := categoryDtos.CategoryResponseDtoFromEntities(result.Data)
		response := utils.SuccessResponsePagination(http.StatusOK, categoryConstants.CATEGORY_PAGINATION_SUCCESS, items, *meta)

		httpContext.JSON(http.StatusOK, response)
	}
}

func (controller *CategoryController) Detail() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		id := uuidValidator.ValidateUUID(httpContext.Param("id"))
		result := controller.categoryService.Detail(ctx, id)
		response := utils.SuccessResponse(http.StatusOK, categoryConstants.CATEGORY_DETAIL_SUCCESS, categoryDtos.CategoryDetailResponseDtoFromEntity(*result))

		httpContext.JSON(http.StatusOK, response)
	}
}

func (controller *CategoryController) Create() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		dto := httpContext.MustGet(middlewares.RequestBodyJsonKey).(*categoryDtos.CategoryCreateRequestDto)
		result := controller.categoryService.Create(ctx, dto)
		response := utils.SuccessResponse(http.StatusCreated, categoryConstants.CATEGORY_CREATE_SUCCESS, categoryDtos.CategoryDetailResponseDtoFromEntity(*result))

		httpContext.JSON(http.StatusCreated, response)
	}
}

func (controller *CategoryController) Update() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		id := uuidValidator.ValidateUUID(httpContext.Param("id"))
		dto := httpContext.MustGet(middlewares.RequestBodyJsonKey).(*categoryDtos.CategoryUpdateRequestDto)
		dto.Id = id
		result := controller.categoryService.Update(ctx, dto)
		response := utils.SuccessResponse(http.StatusOK, categoryConstants.CATEGORY_UPDATE_SUCCESS, categoryDtos.CategoryDetailResponseDtoFromEntity(*result))

		httpContext.JSON(http.StatusOK, response)
	}
}

func (controller *CategoryController) Delete() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		id := uuidValidator.ValidateUUID(httpContext.Param("id"))
		controller.categoryService.SoftDelete(ctx, id)
		response := utils.SuccessResponse(http.StatusOK, categoryConstants.CATEGORY_DELETE_SUCCESS, nil)

		httpContext.JSON(http.StatusOK, response)
	}
}
