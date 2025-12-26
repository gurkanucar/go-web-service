package usecase

import (
	"context"
	"log/slog"
	"project/internal/todo"
)

func (u *useCase) GetTodos(ctx context.Context, filter todo.TodoFilter) ([]todo.Todo, int64, error) {
	slog.InfoContext(ctx, "UseCase: GetTodos")
	return u.repo.GetAll(ctx, filter)
}

// GetTodo retrieves a single todo by ID
// Added here or in a separate file if split, but keeping context from previous edits
func (u *useCase) GetTodo(ctx context.Context, id int) (todo.Todo, error) {
	slog.InfoContext(ctx, "UseCase: GetTodo", "id", id)
	return u.repo.GetByID(ctx, id)
}
