package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func GetDb() (err error) {
	godotenv.Load()
	dsn := fmt.Sprint(
		"host="+os.Getenv("DB_HOST"),
		" user="+os.Getenv("DB_USER"),
		" password="+os.Getenv("DB_PASSWORD"),
		" dbname="+os.Getenv("DB_NAME"),
		" port="+os.Getenv("DB_PORT"),
		" sslmode="+os.Getenv("DB_SSL_MODE"),
	)
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}
