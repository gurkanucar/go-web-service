package todo

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(router fiber.Router, handler *Handler) {
	router.Get("/todos", handler.GetTodos)
	router.Get("/todos/:id", handler.GetTodo)
	router.Post("/todos", handler.CreateTodo)
	router.Put("/todos/:id", handler.UpdateTodo)
}
