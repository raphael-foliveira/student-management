package repository

import (
	"github.com/raphael-foliveira/studentManagementSystem/pkg/models"
	"gorm.io/gorm"
)

type ClassRepository struct {
	Db *gorm.DB
}

func NewClassRepository(db *gorm.DB) *ClassRepository {
	return &ClassRepository{Db: db}
}

func (c *ClassRepository) All() ([]models.Class, error) {
	var classes []models.Class
	tx := c.Db.Model(&models.Class{}).Preload("Students").Preload("Teacher").Find(&classes)
	return classes, tx.Error
}

func (c *ClassRepository) Find(id uint) (*models.Class, error) {
	var class models.Class
	tx := c.Db.Model(&models.Class{}).Preload("Students").Preload("Teacher").First(&class, id).Error
	return &class, tx
}

func (c *ClassRepository) Create(class *models.Class) error {
	return c.Db.Create(class).Error
}

func (c *ClassRepository) Delete(id uint) error {
	return c.Db.Delete(&models.Class{}, id).Error
}
