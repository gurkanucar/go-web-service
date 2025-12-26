package scopes

import (
	"fmt"

	"gorm.io/gorm"
)

type Pagination struct {
	Page int `query:"page" validate:"min=1"`
	Size int `query:"size" validate:"min=1,max=100"`

	SortBy    string `query:"sort_by" validate:"omitempty,oneof=id title created_at"`
	SortOrder string `query:"order" validate:"omitempty,oneof=asc desc"`
}

func (p *Pagination) Paginate(defaultField, defaultDirection string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := p.Page
		if page < 1 {
			page = 1
		}

		size := p.Size
		if size < 1 {
			size = 10
		}
		offset := (page - 1) * size

		sortBy := p.SortBy
		if sortBy == "" {
			sortBy = defaultField
		}

		sortOrder := p.SortOrder
		if sortOrder == "" {
			sortOrder = defaultDirection
		}

		orderClause := fmt.Sprintf("%s %s", sortBy, sortOrder)

		return db.Order(orderClause).Offset(offset).Limit(size)
	}
}
