package models

import (
	"github.com/raphael-foliveira/studentManagementSystem/db"
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

func (c *Class) All() []Class {
	var classes []Class
	db.Db.Model(c).Preload("Students").Preload("Teacher").Find(&classes)
	return classes
}

func (c *Class) Find(id string) {
	db.Db.Model(c).Preload("Students").Preload("Teacher").First(c, id)
}

func (c *Class) Create() {
	db.Db.Create(c)
}

func (c *Class) Delete(id string) {
	db.Db.Delete(c, id)
}
