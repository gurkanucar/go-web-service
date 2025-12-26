package shared

import (
	"time"
)

type BaseEntity struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	// CreatedBy string         `gorm:"size:100" json:"created_by,omitempty"`
	// UpdatedBy string `gorm:"size:100" json:"updated_by,omitempty"`
}
