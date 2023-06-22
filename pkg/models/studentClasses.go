package models

import "gorm.io/gorm"

type StudentClasses struct {
	ID        uint `gorm:"primaryKey"`
	StudentID uint `json:"-"`
	Student   Student
	ClassID   uint `json:"-"`
	Class     Class
	DeletedAt gorm.DeletedAt `json:"-"`
}
