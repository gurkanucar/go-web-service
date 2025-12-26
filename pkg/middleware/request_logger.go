package middleware

import (
	"errors"
	"log/slog"
	"time"

	"project/pkg/errorhandler"

	"github.com/gofiber/fiber/v2"
)

func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		latency := time.Since(start)

		status := c.Response().StatusCode()
		if err != nil {
			var busErr *errorhandler.BusinessError
			if errors.As(err, &busErr) {
				status = busErr.HTTPStatus
			} else if fiberErr, ok := err.(*fiber.Error); ok {
				status = fiberErr.Code
			} else {
				status = fiber.StatusInternalServerError
			}
		}

		slog.InfoContext(c.UserContext(), "incoming request",
			"method", c.Method(),
			"path", c.Path(),
			"status", status,
			"latency_ms", latency.Milliseconds(),
			"ip", c.IP(),
		)

		return err
	}
}
