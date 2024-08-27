package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDBConnection() (*gorm.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file")
	}

	fmt.Println("ENV: ", os.Getenv("ENV"))
	fmt.Println("ENV: ", os.Getenv("POSTGRES_USER"))
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=require TimeZone=Asia/Shanghai",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
