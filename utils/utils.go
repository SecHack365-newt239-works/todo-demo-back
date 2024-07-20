package utils

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDBConnection() (*gorm.DB, error) {
	url := os.Getenv("POSTGRES_URL")
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	return db, err
}
