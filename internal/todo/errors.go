package todo

import (
	"project/pkg/errorhandler"

	"github.com/gofiber/fiber/v2"
)

var (
	ErrTodoNotFound = &errorhandler.BusinessError{Code: 2004, Message: "Todo not found", HTTPStatus: fiber.StatusNotFound}
	ErrInvalidID    = &errorhandler.BusinessError{Code: 2002, Message: "Invalid ID", HTTPStatus: fiber.StatusBadRequest}
	ErrInvalidBody  = &errorhandler.BusinessError{Code: 2003, Message: "Invalid Request Body", HTTPStatus: fiber.StatusBadRequest}
)
