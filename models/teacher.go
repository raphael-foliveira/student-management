package models

import (
	"strconv"

	"github.com/raphael-foliveira/studentManagementSystem/db"
	"gorm.io/gorm"
)

type Teacher struct {
	ID        uint           `gorm:"primaryKey"`
	FirstName string         `json:"firstName"`
	LastName  string         `json:"lastName"`
	Classes   []Class        `json:"classes,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty"`
}

func (t *Teacher) All() []Teacher {
	var teachers []Teacher
	db.Db.Model(t).Preload("Classes").Find(&teachers)
	return teachers
}

func (t *Teacher) Create() {
	db.Db.Model(t).Preload("Classes").Create(t)
}

func (t *Teacher) Find(id string) error {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	db.Db.First(t, intId)
	return nil
}
func (t *Teacher) Delete(id string) error {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	db.Db.Delete(t, intId)
	return nil
}
