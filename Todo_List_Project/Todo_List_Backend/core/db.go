package core

import (
	"fmt"
	"todo-list/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDatabase() {
	fmt.Println("Connecting to database...")
	config := LoadConfig()

	// Create the DSN (Data Source Name) string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", config.DBHost, config.DBUsername, config.DBPassword, config.DBName, config.DBPort)

	// Use postgres.Open instead of passing a string
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("There is an error while connecting to the database:", err)
		panic(err)
	}

	db.AutoMigrate(&model.TodoList{})

	Db = db
	fmt.Println("Successfully connected to database!")
}
