package usecase

import (
	"context"
	"log/slog"
	"project/internal/todo"
)

func (u *useCase) UpdateTodo(ctx context.Context, id int, title string, completed bool) (todo.Todo, error) {
	slog.InfoContext(ctx, "Updating todo", "id", id)

	existing, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return todo.Todo{}, err
	}

	existing.Title = title
	existing.Completed = completed

	return u.repo.Update(ctx, existing)
}
