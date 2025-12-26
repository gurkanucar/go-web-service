package errorhandler

import (
	"encoding/json"
	"errors"
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Error             bool              `json:"error"`
	TraceID           string            `json:"traceId"`
	BusinessErrorCode int               `json:"businessErrorCode"`
	Message           string            `json:"message"`
	ValidationErrors  map[string]string `json:"validationErrors,omitempty"`
}

func FiberErrorHandler(c *fiber.Ctx, err error) error {
	resp := ErrorResponse{
		Error:             true,
		TraceID:           getTraceID(c),
		BusinessErrorCode: 0,
		Message:           "Internal Server Error",
	}
	code := fiber.StatusInternalServerError

	var busErr *BusinessError
	if errors.As(err, &busErr) {
		code = busErr.HTTPStatus
		resp.BusinessErrorCode = busErr.Code
		resp.Message = busErr.Message
	} else if fiberErr, ok := err.(*fiber.Error); ok {
		code = fiberErr.Code
		resp.Message = fiberErr.Message

		if code == fiber.StatusBadRequest {
			resp.BusinessErrorCode = fiber.StatusBadRequest
			if validErrs := extractValidationErrors(fiberErr.Message); validErrs != nil {
				resp.ValidationErrors = validErrs
				resp.Message = "Validation failed"
				resp.BusinessErrorCode = fiber.ErrBadRequest.Code
			}
		}
	} else {
		resp.Message = err.Error()
	}

	logError(c, code, resp.BusinessErrorCode, err, resp.TraceID)

	return c.Status(code).JSON(resp)
}

func getTraceID(c *fiber.Ctx) string {
	if traceID, ok := c.Locals("traceId").(string); ok {
		return traceID
	}
	return "unknown"
}

func extractValidationErrors(message string) map[string]string {
	var validErrs map[string]string
	if err := json.Unmarshal([]byte(message), &validErrs); err == nil {
		return validErrs
	}
	return nil
}

func logError(c *fiber.Ctx, code, businessCode int, err error, traceID string) {
	slog.ErrorContext(c.UserContext(), err.Error(),
		"path", c.Path(),
		"status", code,
		"business_code", businessCode,
	)
}
