package exceptions

import "net/http"

type Exception struct {
	StatusCode       int
	ErrorMessage     string
	ValidationErrors map[string]string
}

// 400 Bad Request
func BadRequestException(message string) *Exception {
	return &Exception{
		ErrorMessage: message,
		StatusCode:   http.StatusBadRequest,
	}
}

// 400 Bad Request - Validation Error
func ValidationException(errors map[string]string) *Exception {
	return &Exception{
		ErrorMessage:     "Validation failed",
		StatusCode:       http.StatusBadRequest,
		ValidationErrors: errors,
	}
}

// 401 Unauthorized
func UnauthenticatedException(message string) *Exception {
	return &Exception{
		ErrorMessage: message,
		StatusCode:   http.StatusUnauthorized,
	}
}

// 403 Forbidden
func ForbiddenException(message string) *Exception {
	return &Exception{
		ErrorMessage: message,
		StatusCode:   http.StatusForbidden,
	}
}

// 404 Not Found
func NotFoundException(message string) *Exception {
	return &Exception{
		ErrorMessage: message,
		StatusCode:   http.StatusNotFound,
	}
}

// 409 Conflict
func ConflictException(message string) *Exception {
	return &Exception{
		ErrorMessage: message,
		StatusCode:   http.StatusConflict,
	}
}

// 422 Unprocessable Entity
func UnprocessableEntityException(message string) *Exception {
	return &Exception{
		ErrorMessage: message,
		StatusCode:   http.StatusUnprocessableEntity,
	}
}

// 500 Internal Server Error
func ServerErrorException(err error) *Exception {
	return &Exception{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}
}

// 503 Service Unavailable
func ServiceUnavailableException(message string) *Exception {
	return &Exception{
		ErrorMessage: message,
		StatusCode:   http.StatusServiceUnavailable,
	}
}
