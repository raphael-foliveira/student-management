package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/raphael-foliveira/studentManagementSystem/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDb() (*gorm.DB, error) {
	godotenv.Load()
	dsn := fmt.Sprint(
		"host="+os.Getenv("DB_HOST"),
		" user="+os.Getenv("DB_USER"),
		" password="+os.Getenv("DB_PASSWORD"),
		" dbname="+os.Getenv("DB_NAME"),
		" port="+os.Getenv("DB_PORT"),
		" sslmode="+os.Getenv("DB_SSL_MODE"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(
		&models.Student{},
		&models.Teacher{},
		&models.Class{},
		&models.StudentClasses{},
	)

	return db, nil
}
