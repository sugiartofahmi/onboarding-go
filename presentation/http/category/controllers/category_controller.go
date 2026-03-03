package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	categoryConstants "event-backend/app/category/constants"
	categoryInterfaces "event-backend/app/category/interfaces"
	"event-backend/infrastructure/utils"
	categoryDtos "event-backend/presentation/http/category/dtos"

	uuidValidator "event-backend/infrastructure/validators"
)

type CategoryController struct {
	categoryService categoryInterfaces.CategoryServiceInterface
}

func NewCategoryController(router *gin.Engine, categoryService categoryInterfaces.CategoryServiceInterface) {
	categoryRoute := router.Group("/api/v1/categories")

	controller := &CategoryController{
		categoryService: categoryService,
	}

	categoryRoute.GET("", controller.Pagination())
	categoryRoute.POST("", controller.Create())
	categoryRoute.GET("/:id", controller.Detail())
	categoryRoute.PUT("/:id", controller.Update())
	categoryRoute.DELETE("/:id", controller.Delete())
}

func (controller *CategoryController) Pagination() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		dto := categoryDtos.AssignCategoryQueryRequestDTO(httpContext)
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
		dto := &categoryDtos.CategoryCreateRequestDTO{}
		httpContext.ShouldBindJSON(dto)
		result := controller.categoryService.Create(ctx, dto)
		response := utils.SuccessResponse(http.StatusCreated, categoryConstants.CATEGORY_CREATE_SUCCESS, categoryDtos.CategoryDetailResponseDtoFromEntity(*result))

		httpContext.JSON(http.StatusCreated, response)
	}
}

func (controller *CategoryController) Update() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		id := uuidValidator.ValidateUUID(httpContext.Param("id"))
		dto := &categoryDtos.CategoryUpdateRequestDTO{}
		httpContext.ShouldBindJSON(dto)
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
