package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"event-backend/infrastructure/exceptions"
	"event-backend/infrastructure/utils"
)

func ExceptionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer handlePanic(c)
		c.Next()
	}
}

func handlePanic(c *gin.Context) {
	if err := recover(); err != nil {
		panicException := createPanicException(err)
		errorMessage := panicException.ErrorMessage
		if errorMessage == "" {
			errorMessage = getErrorMessageByStatusCode(panicException.StatusCode)
		}

		var errors *map[string]string
		if panicException.ValidationErrors != nil {
			errors = &panicException.ValidationErrors
		}

		response := utils.ErrorResponse(panicException.StatusCode, errorMessage, errors)
		c.JSON(panicException.StatusCode, response)
		c.Abort()
	}
}

func createPanicException(err interface{}) exceptions.Exception {
	if ex, ok := err.(*exceptions.Exception); ok {
		return *ex
	}

	if ex, ok := err.(exceptions.Exception); ok {
		return ex
	}

	return exceptions.Exception{
		ErrorMessage: err.(error).Error(),
		StatusCode:   http.StatusInternalServerError,
	}
}

func getErrorMessageByStatusCode(statusCode int) string {
	switch statusCode {
	case http.StatusBadRequest:
		return "Bad Request"
	case http.StatusUnauthorized:
		return "Unauthorized"
	case http.StatusForbidden:
		return "Forbidden"
	case http.StatusNotFound:
		return "Not Found"
	case http.StatusUnprocessableEntity:
		return "Unprocessable Entity"
	case http.StatusServiceUnavailable:
		return "Service Unavailable"
	default:
		return "Internal Server Error"
	}
}
