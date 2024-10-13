package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-fiber-jwt-task/model"
)

// instance database postgres
var DB *gorm.DB

// connect to postgres
func Connect() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
	)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect DB")
	}

	fmt.Println("Success connect to DB")

	//Run migration DB
	err = DB.AutoMigrate(&model.Task{}, &model.User{})
	if err != nil {
		panic("Failed to run migration DB")
	}

	fmt.Println("Migration DB successfully")

}
