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

func AssignPaginationQueryRequestDto(c *gin.Context) *PaginationQueryRequestDto {
	q := &PaginationQueryRequestDto{}

	queryParamsPage := c.Query("page")
	if queryParamsPage == "" {
		queryParamsPage = "1"
	}
	q.Page, _ = strconv.Atoi(queryParamsPage)

	queryParamsPerPage := c.Query("per_page")
	if queryParamsPerPage == "" {
		queryParamsPerPage = "10"
	}
	q.PerPage, _ = strconv.Atoi(queryParamsPerPage)

	q.Search = c.Query("search")

	q.SortBy = c.Query("sort_by")
	if q.SortBy == "" {
		q.SortBy = "created_at"
	}

	q.Order = enums.SortOrderEnum(c.Query("order"))
	if q.Order == "" {
		q.Order = enums.SortOrderDesc
	}

	return q
}
