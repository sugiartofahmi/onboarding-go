package utils

import (
	"gorm.io/gorm"

	infradtos "event-backend/infrastructure/dtos"
)

func Paginate(q *infradtos.PaginationQueryRequestDto) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := q.Page
		if page <= 0 {
			page = 1
		}

		perPage := q.PerPage
		if perPage <= 0 || perPage > 100 {
			perPage = 10
		}

		offset := (page - 1) * perPage
		return db.Offset(offset).Limit(perPage)
	}
}
