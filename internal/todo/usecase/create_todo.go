package usecase

import (
	"context"
	"log/slog"
	"project/internal/todo"
)

func (s *useCase) CreateTodo(ctx context.Context, title string) (todo.Todo, error) {
	slog.InfoContext(ctx, "UseCase: CreateTodo", "title", title)
	return s.repo.Create(ctx, title)
}
