package errorhandler

import (
	"github.com/gofiber/fiber/v2"
)

type BusinessError struct {
	Code       int    `json:"businessErrorCode"`
	Message    string `json:"message"`
	HTTPStatus int    `json:"-"`
}

func (e *BusinessError) Error() string {
	return e.Message
}

func NewBusinessError(code int, message string, httpStatus int) *BusinessError {
	return &BusinessError{
		Code:       code,
		Message:    message,
		HTTPStatus: httpStatus,
	}
}

var (
	ErrRequestValidation = &BusinessError{Code: 400, Message: "Validation failed", HTTPStatus: fiber.StatusBadRequest}

	ErrNotFound = &BusinessError{Code: 404, Message: "Resource not found", HTTPStatus: fiber.StatusNotFound}
)
