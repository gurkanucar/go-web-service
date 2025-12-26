package todo

import (
	"context"
	"project/pkg/database/scopes"
	"project/pkg/database/shared"
)

type Todo struct {
	shared.BaseEntity
	Title     string `gorm:"size:255;not null"`
	Completed bool   `gorm:"default:false;not null"`
}

func (Todo) TableName() string {
	return "todos"
}

type TodoFilter struct {
	scopes.Pagination

	Search    string `query:"q"`
	Completed *bool  `query:"completed"`
}

type Repository interface {
	GetAll(ctx context.Context, filter TodoFilter) ([]Todo, int64, error)
	GetByID(ctx context.Context, id int) (Todo, error)
	Create(ctx context.Context, title string) (Todo, error)
	Update(ctx context.Context, todo Todo) (Todo, error)
}

type UseCase interface {
	GetTodos(ctx context.Context, filter TodoFilter) ([]Todo, int64, error)
	GetTodo(ctx context.Context, id int) (Todo, error)
	CreateTodo(ctx context.Context, title string) (Todo, error)
	UpdateTodo(ctx context.Context, id int, title string, completed bool) (Todo, error)
}
