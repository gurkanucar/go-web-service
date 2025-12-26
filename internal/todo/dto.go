package todo

import (
	"time"
)

type CreateTodoRequest struct {
	Title string `json:"title" validate:"required,min=3,max=100"`
}

func (r *CreateTodoRequest) ToEntity() Todo {
	return Todo{
		Title: r.Title,
	}
}

type UpdateTodoRequest struct {
	Title     string `json:"title" validate:"required,min=3,max=100"`
	Completed bool   `json:"completed"`
}

func (r *UpdateTodoRequest) ToEntity() Todo {
	return Todo{
		Title:     r.Title,
		Completed: r.Completed,
	}
}

type TodoResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewTodoResponse(t Todo) TodoResponse {
	return TodoResponse{
		ID:        t.ID,
		Title:     t.Title,
		Completed: t.Completed,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

func NewTodoListResponse(todos []Todo) []TodoResponse {
	responses := make([]TodoResponse, len(todos))
	for i, todo := range todos {
		responses[i] = NewTodoResponse(todo)
	}
	return responses
}
