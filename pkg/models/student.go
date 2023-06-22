package models

import (
	"gorm.io/gorm"
)

type Student struct {
	ID        uint           `gorm:"primaryKey"`
	FirstName string         `json:"firstName"`
	LastName  string         `json:"lastName"`
	Email     string         `json:"email"`
	Course    string         `json:"course"`
	Classes   []Class        `json:"classes,omitempty" gorm:"many2many:StudentClasses"`
	Semester  int            `json:"semester"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
