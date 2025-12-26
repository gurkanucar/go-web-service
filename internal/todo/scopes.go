package todo

import (
	"strings"

	"gorm.io/gorm"
)

func ScopeSearch(search string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if search == "" {
			return db
		}
		term := "%" + strings.ToLower(search) + "%"
		return db.Where("LOWER(title) LIKE ?", term)
	}
}

func ScopeCompleted(completed *bool) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if completed == nil {
			return db
		}
		return db.Where("completed = ?", *completed)
	}
}
