package usecase

import "project/internal/todo"

type useCase struct {
	repo todo.Repository
}

func NewUseCase(repo todo.Repository) todo.UseCase {
	return &useCase{repo: repo}
}
