package dtos

import (
	"strconv"

	"event-backend/infrastructure/enums"

	"github.com/gin-gonic/gin"
)

type PaginationQueryRequestDto struct {
	Search  string              `form:"search"`
	PerPage int                 `form:"per_page"`
	Page    int                 `form:"page"`
	SortBy  string              `form:"sort_by"`
	Order   enums.SortOrderEnum `form:"order"`
}

func NewPaginationQueryRequestDto(c *gin.Context) *PaginationQueryRequestDto {
	q := &PaginationQueryRequestDto{}

	if page, err := strconv.Atoi(c.DefaultQuery("page", "1")); err == nil && page >= 1 {
		q.Page = page
	} else {
		q.Page = 1
	}

	if perPage, err := strconv.Atoi(c.DefaultQuery("per_page", "10")); err == nil && perPage >= 1 && perPage <= 100 {
		q.PerPage = perPage
	} else {
		q.PerPage = 10
	}

	q.Search = c.Query("search")
	q.SortBy = c.DefaultQuery("sort_by", "created_at")

	order := enums.SortOrderEnum(c.DefaultQuery("order", string(enums.SortOrderDesc)))
	if order != enums.SortOrderAsc && order != enums.SortOrderDesc {
		order = enums.SortOrderDesc
	}
	q.Order = order

	return q
}
