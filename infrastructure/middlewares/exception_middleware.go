package middlewares

import (
	"net/http"
	"strings"

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

	if bindingErr, ok := err.(gin.Error); ok {
		if bindingErr.Type == gin.ErrorTypeBind {
			validationErrors := parseBindingErrors(bindingErr.Err)
			return exceptions.Exception{
				ErrorMessage:     "Validation failed",
				StatusCode:       http.StatusBadRequest,
				ValidationErrors: validationErrors,
			}
		}
	}

	errMsg := "Internal Server Error"
	if e, ok := err.(error); ok {
		errMsg = e.Error()
	} else if s, ok := err.(string); ok {
		errMsg = s
	}

	return exceptions.Exception{
		ErrorMessage: errMsg,
		StatusCode:   http.StatusInternalServerError,
	}
}

func parseBindingErrors(err error) map[string]string {
	errors := make(map[string]string)

	if validationErrors, ok := err.(interface{ GetErrors() []FieldError }); ok {
		for _, fieldErr := range validationErrors.GetErrors() {
			errors[fieldErr.Field] = fieldErr.Message
		}
	}

	if strings.Contains(err.Error(), "Key: '") {
		lines := strings.Split(err.Error(), "\n")
		for _, line := range lines {
			if strings.Contains(line, "Key: '") {
				start := strings.Index(line, "'") + 1
				end := strings.LastIndex(line, "'")
				if start > 0 && end > start {
					field := line[start:end]
					msg := "Field " + field + " is required"
					if strings.Contains(line, "required") {
						msg = "Field " + field + " is required"
					} else if strings.Contains(line, "email") {
						msg = "Field " + field + " must be a valid email format"
					} else if strings.Contains(line, "min") {
						msg = "Field " + field + " must be at least 6 characters"
					} else {
						msg = "Field " + field + " is invalid"
					}
					errors[field] = msg
				}
			}
		}
	}

	return errors
}

type FieldError struct {
	Field   string
	Message string
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
