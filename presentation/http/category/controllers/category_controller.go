package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	categoryDtos "event-backend/app/category/dtos"
	categoryInterfaces "event-backend/app/category/interfaces"
	"event-backend/infrastructure/utils"
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
