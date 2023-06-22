package models

import (
	"gorm.io/gorm"
)

type Class struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `json:"name"`
	Students  []Student      `json:"students" gorm:"many2many:StudentClasses"`
	TeacherID uint           `json:"teacherId"`
	Teacher   Teacher        `json:"teacher,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty"`
}
