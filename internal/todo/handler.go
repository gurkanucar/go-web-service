package todo

import (
	"log/slog"
	"project/pkg/database/scopes"
	"project/pkg/errorhandler"
	"project/pkg/response"
	"project/pkg/validator"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service   UseCase
	validator *validator.Validator
}

func NewHandler(service UseCase, v *validator.Validator) *Handler {
	return &Handler{service: service, validator: v}
}

// GetTodos godoc
// @Summary Get all todos
// @Description Get a list of all todos with pagination, filtering, and sorting
// @Tags todos
// @Accept json
// @Produce json
// @Param q query string false "Search query (title)"
// @Param completed query bool false "Filter by completed status"
// @Param page query int false "Page number" default(1)
// @Param size query int false "Page size" default(10)
// @Param sort_by query string false "Sort by field (allowed: id, title, created_at)"
// @Param order query string false "Sort direction (asc, desc)" default(desc)
// @Success 200 {object} response.BaseResponse{data=response.PageData{content=[]TodoResponse}}
// @Failure 400 {object} response.BaseResponse "Validation error or invalid query params"
// @Router /api/todos [get]
func (h *Handler) GetTodos(c *fiber.Ctx) error {
	// Set defaults before parsing
	filter := TodoFilter{
		Pagination: scopes.Pagination{Page: 1, Size: 10},
	}

	if err := c.QueryParser(&filter); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid query parameters")
	}

	if errs := h.validator.Validate(filter); errs != nil {
		return errorhandler.ErrRequestValidation
	}

	slog.InfoContext(c.UserContext(), "Fetching todos", "filter", filter)

	todos, total, err := h.service.GetTodos(c.UserContext(), filter)
	if err != nil {
		return err
	}

	return response.Page(c, "Data retrieved successfully",
		NewTodoListResponse(todos), total, filter.Page, filter.Size, len(todos))
}

// GetTodo godoc
// @Summary Get a todo
// @Description Get a todo by ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} response.BaseResponse{data=TodoResponse}
// @Failure 404 {object} errorhandler.ErrorResponse
// @Router /api/todos/{id} [get]
func (h *Handler) GetTodo(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return ErrInvalidID
	}

	todo, err := h.service.GetTodo(c.UserContext(), id)
	if err != nil {
		return err
	}

	return response.Success(c, "Operation successfully completed", NewTodoResponse(todo))
}

// CreateTodo godoc
// @Summary Create a todo
// @Description Create a new todo
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body CreateTodoRequest true "Todo to create"
// @Success 201 {object} response.BaseResponse{data=TodoResponse}
// @Failure 400 {object} errorhandler.ErrorResponse
// @Router /api/todos [post]
func (h *Handler) CreateTodo(c *fiber.Ctx) error {
	req, err := validator.ValidatedBody[CreateTodoRequest](c, h.validator)
	if err != nil {
		return err
	}

	newTodo, err := h.service.CreateTodo(c.UserContext(), req.Title)
	if err != nil {
		return err
	}

	return response.Success(c, "Todo created successfully", NewTodoResponse(newTodo))
}

// UpdateTodo godoc
// @Summary Update a todo
// @Description Update an existing todo
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Param todo body UpdateTodoRequest true "Todo update data"
// @Success 200 {object} response.BaseResponse{data=TodoResponse}
// @Failure 400 {object} errorhandler.ErrorResponse
// @Router /api/todos/{id} [put]
func (h *Handler) UpdateTodo(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return errorhandler.ErrRequestValidation
	}

	req, err := validator.ValidatedBody[UpdateTodoRequest](c, h.validator)
	if err != nil {
		return err
	}

	todo, err := h.service.UpdateTodo(c.UserContext(), id, req.Title, req.Completed)
	if err != nil {
		return err
	}

	return response.Success(c, "Todo updated successfully", NewTodoResponse(todo))
}
