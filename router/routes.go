package router

import (
	"github.com/RakibSiddiquee/gin-gorm-crud/controllers"
	"github.com/gin-gonic/gin"
)

func GetRoutes(r *gin.Engine) {
	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.FindPost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)
}
