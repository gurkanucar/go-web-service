package todo

import (
	"context"

	"gorm.io/gorm"
)

type gormRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &gormRepository{db: db}
}

func (r *gormRepository) GetAll(ctx context.Context, filter TodoFilter) ([]Todo, int64, error) {
	var todos []Todo
	var total int64

	query := r.db.WithContext(ctx).Model(&Todo{})

	query = query.Scopes(
		ScopeSearch(filter.Search),
		ScopeCompleted(filter.Completed),
	)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Scopes(filter.Pagination.Paginate("id", "desc")).Find(&todos).Error

	if err != nil {
		return nil, 0, err
	}

	return todos, total, nil
}

func (r *gormRepository) GetByID(ctx context.Context, id int) (Todo, error) {
	var todo Todo
	if err := r.db.WithContext(ctx).First(&todo, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return Todo{}, ErrTodoNotFound
		}
		return Todo{}, err
	}
	return todo, nil
}

func (r *gormRepository) Create(ctx context.Context, title string) (Todo, error) {
	todo := Todo{
		Title:     title,
		Completed: false,
	}
	if err := r.db.WithContext(ctx).Create(&todo).Error; err != nil {
		return Todo{}, err
	}
	return todo, nil
}

func (r *gormRepository) Update(ctx context.Context, t Todo) (Todo, error) {
	// GORM Save updates all fields. Ensure ID is set.
	if err := r.db.WithContext(ctx).Save(&t).Error; err != nil {
		return Todo{}, err
	}
	return t, nil
}
