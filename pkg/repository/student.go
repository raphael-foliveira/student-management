package repository

import (
	"github.com/raphael-foliveira/studentManagementSystem/pkg/models"
	"gorm.io/gorm"
)

type StudentRepository struct {
	Db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *StudentRepository {
	return &StudentRepository{Db: db}
}

func (sr *StudentRepository) All() ([]models.Student, error) {
	var students []models.Student
	tx := sr.Db.Model(&models.Student{}).Preload("Classes").Find(&students)
	return students, tx.Error
}

func (s *StudentRepository) Find(id uint) (*models.Student, error) {
	var student models.Student
	tx := s.Db.Model(&models.Student{ID: id}).Preload("Classes.Teacher").First(&student)
	return &student, tx.Error
}

func (s *StudentRepository) Create(student *models.Student) error {
	return s.Db.Create(student).Error
}

func (s *StudentRepository) Update(student *models.Student) error {
	return s.Db.Save(&student).Error

}

func (s *StudentRepository) Delete(id uint) error {
	return s.Db.Delete(&models.Student{ID: id}).Error

}
