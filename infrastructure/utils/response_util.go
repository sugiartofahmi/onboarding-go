package utils

import "math"

const ApiVersion = "1.0.0"

type BaseResponse struct {
	Version    string             `json:"version"`
	StatusCode int                `json:"statusCode"`
	Data       *interface{}       `json:"data,omitempty"`
	Items      *interface{}       `json:"items,omitempty"`
	Count      *int               `json:"count,omitempty"`
	Meta       *MetaResponse      `json:"meta,omitempty"`
	Error      string             `json:"errorMessage,omitempty"`
	Errors     *map[string]string `json:"errors,omitempty"`
}

type MetaResponse struct {
	Page      int `json:"page"`
	PerPage   int `json:"perPage"`
	Total     int `json:"total"`
	TotalPage int `json:"totalPage"`
}

func SuccessResponse(statusCode int, data interface{}) *BaseResponse {
	return &BaseResponse{
		Version:    ApiVersion,
		StatusCode: statusCode,
		Data:       &data,
	}
}

func SuccessResponseList(statusCode int, items interface{}, count int) *BaseResponse {
	return &BaseResponse{
		Version:    ApiVersion,
		StatusCode: statusCode,
		Items:      &items,
		Count:      &count,
	}
}

func SuccessResponsePagination(statusCode int, items interface{}, meta MetaResponse) *BaseResponse {
	return &BaseResponse{
		Version:    ApiVersion,
		StatusCode: statusCode,
		Items:      &items,
		Meta:       &meta,
	}
}

func ErrorResponse(statusCode int, errorMessage string, errors *map[string]string) *BaseResponse {
	return &BaseResponse{
		Version:    ApiVersion,
		StatusCode: statusCode,
		Error:      errorMessage,
		Errors:     errors,
	}
}

func PaginationMetaBuilder(page int, perPage int, total int) *MetaResponse {
	totalPage := int(math.Ceil(float64(total) / float64(perPage)))

	return &MetaResponse{
		Page:      page,
		PerPage:   perPage,
		Total:     total,
		TotalPage: totalPage,
	}
}
