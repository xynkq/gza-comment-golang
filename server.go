package main

import (
	"gza/comment/controllers"
	"gza/comment/models"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	commentAPI := r.Group("/comment")
	{
		// GET: http://localhost:8080/comment
		commentAPI.GET("/", controllers.GetComments)
		// GET: http://localhost:8080/comment/:id
		commentAPI.GET("/:id", controllers.GetComment)
	}
	r.Use()
	r.NoRoute(func(c *gin.Context) {
		// In gin this is how you return a JSON response
		c.JSON(404, gin.H{"message": "Not found"})
	})

	DB = models.ConnectDatabase()
	r.Run(":8080")
}
