package repository

import (
	"github.com/raphael-foliveira/studentManagementSystem/pkg/models"
	"gorm.io/gorm"
)

type TeacherRepository struct {
	Db *gorm.DB
}

func NewTeacherRepository(db *gorm.DB) *TeacherRepository {
	return &TeacherRepository{Db: db}
}

func (t *TeacherRepository) All() ([]models.Teacher, error) {
	var teachers []models.Teacher
	tx := t.Db.Model(&models.Teacher{}).Preload("Classes").Find(&teachers)
	return teachers, tx.Error
}

func (t *TeacherRepository) Create(teacher *models.Teacher) error {
	return t.Db.Model(&models.Teacher{}).Preload("Classes").Create(teacher).Error
}

func (t *TeacherRepository) Find(id uint) (*models.Teacher, error) {
	var teacher models.Teacher
	tx := t.Db.First(&teacher, id)
	return &teacher, tx.Error
}
func (t *TeacherRepository) Delete(id uint) error {
	teacher := &models.Teacher{ID: id}
	return t.Db.Delete(teacher).Error
}
