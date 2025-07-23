package db

import (
	"fmt"
	"log"
	"os"
	"todolist/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() {
	// connect to database
	var err error
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// migragte the schema
	err = DB.AutoMigrate(&models.User{}, &models.Todolist{})
	if err != nil {
		log.Fatal("failed to migrate the schema", err)
	}

	fmt.Println("connected to db successfully")
}
