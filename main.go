package main

import (
	"github.com/RakibSiddiquee/gin-gorm-crud/initializers"
	"github.com/RakibSiddiquee/gin-gorm-crud/router"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()

	router.GetRoutes(r)

	r.Run()

}
