package main

import (
	"github.com/RakibSiddiquee/gin-gorm-crud/initializers"
	"github.com/RakibSiddiquee/gin-gorm-crud/models"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	err := initializers.DB.AutoMigrate(&models.Post{})

	if err != nil {
		log.Fatal("Migration failed")
	}
}
