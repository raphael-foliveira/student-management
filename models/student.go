package models

import (
	"strconv"

	"github.com/raphael-foliveira/studentManagementSystem/db"
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

func (s *Student) All() []Student {
	var students []Student
	db.Db.Model(s).Preload("Classes").Find(&students)
	return students
}

func (s *Student) Find(id string) error {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	db.Db.Model(s).Preload("Classes.Teacher").First(s, intId)
	return nil
}

func (s *Student) Create() {
	db.Db.Create(s)
}

func (s *Student) Update(id string) error {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	db.Db.Find(s, intId)
	db.Db.Save(s)
	return nil
}

func (s *Student) Delete(id string) {
	db.Db.Delete(s, id)
}
