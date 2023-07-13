package controllers

import (
	"github.com/RakibSiddiquee/gin-gorm-crud/initializers"
	"github.com/RakibSiddiquee/gin-gorm-crud/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatePost(c *gin.Context) {
	// Get data from req body
	var request struct {
		Title string
		Body  string
	}

	c.Bind(&request)

	// Create a post
	post := models.Post{Title: request.Title, Body: request.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return the post
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func GetPosts(c *gin.Context) {
	// Get the posts
	var posts []models.Post

	initializers.DB.Find(&posts)

	// Return the posts
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func FindPost(c *gin.Context) {
	// Get the id from url
	id := c.Param("id")

	// Find the post
	var post models.Post
	initializers.DB.First(&post, id)

	// Return the post
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {
	// Get the id from url
	id := c.Param("id")

	// Get the data from request body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// Find the post by id
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// Return the post
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context) {
	// Get the id from url
	id := c.Param("id")

	// Delete the post
	initializers.DB.Delete(&models.Post{}, id)

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "The post deleted successfully!",
	})
}
