package dtos

import (
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

func AssignPaginationQueryRequestDto(c *gin.Context) *PaginationQueryRequestDto {
	q := &PaginationQueryRequestDto{}
	c.ShouldBindQuery(q)
	return q
}
