package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

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
	categoryRoute.GET("/:id", controller.Detail())
}

func (controller *CategoryController) Pagination() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		dto := categoryDtos.AssignCategoryQueryRequestDTO(httpContext)
		result := controller.categoryService.Pagination(ctx, dto)
		meta := utils.PaginationMetaBuilder(dto.Page, dto.PerPage, int(result.Count))
		items := categoryDtos.CategoryResponseDtoFromEntities(result.Data)
		response := utils.SuccessResponsePagination(http.StatusOK, items, *meta)

		httpContext.JSON(http.StatusOK, response)
	}
}

func (controller *CategoryController) Detail() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		id := uuidValidator.ValidateUUID(httpContext.Param("id"))
		result := controller.categoryService.Detail(ctx, id)
		response := utils.SuccessResponse(http.StatusOK, categoryDtos.CategoryDetailResponseDtoFromEntity(*result))

		httpContext.JSON(http.StatusOK, response)
	}
}