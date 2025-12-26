package response

import (
	"math"

	"github.com/gofiber/fiber/v2"
)

type BaseResponse struct {
	Error   bool   `json:"error"`
	TraceID string `json:"traceId"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type Pageable struct {
	TotalElements    int64 `json:"totalElements"`
	NumberOfElements int   `json:"numberOfElements"`
	TotalPages       int   `json:"totalPages"`
	Size             int   `json:"size"`
	Last             bool  `json:"last"`
	First            bool  `json:"first"`
	Empty            bool  `json:"empty"`
}

type PageData struct {
	Content  any      `json:"content"`
	Pageable Pageable `json:"pageable"`
}

func Success(c *fiber.Ctx, message string, data any) error {
	traceID, _ := c.Locals("traceId").(string)
	return c.JSON(BaseResponse{
		Error:   false,
		TraceID: traceID,
		Message: message,
		Data:    data,
	})
}

func Page(c *fiber.Ctx, message string, content any, total int64, page int, size int, itemCount int) error {
	traceID, _ := c.Locals("traceId").(string)

	if size <= 0 {
		size = 10
	}
	totalPages := int(math.Ceil(float64(total) / float64(size)))
	if totalPages < 0 {
		totalPages = 0
	}

	first := page == 1
	last := page >= totalPages

	return c.JSON(BaseResponse{
		Error:   false,
		TraceID: traceID,
		Message: message,
		Data: PageData{
			Content: content,
			Pageable: Pageable{
				TotalElements:    total,
				NumberOfElements: itemCount,
				TotalPages:       totalPages,
				Size:             size,
				Last:             last,
				First:            first,
				Empty:            itemCount == 0,
			},
		},
	})
}
