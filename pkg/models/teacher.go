package models

import (
	"gorm.io/gorm"
)

type Teacher struct {
	ID        uint           `gorm:"primaryKey"`
	FirstName string         `json:"firstName"`
	LastName  string         `json:"lastName"`
	Classes   []Class        `json:"classes,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty"`
}
