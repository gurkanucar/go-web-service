package middleware

import (
	"context"
	"project/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Trace() fiber.Handler {
	return func(c *fiber.Ctx) error {
		traceID := c.Get("X-Trace-ID")
		if traceID == "" {
			traceID = uuid.New().String()
		}

		c.Locals("traceId", traceID)
		c.Set("X-Trace-ID", traceID)

		ctx := context.WithValue(c.UserContext(), logger.TraceIDKey, traceID)
		c.SetUserContext(ctx)

		return c.Next()
	}
}
