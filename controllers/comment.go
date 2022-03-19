package controllers

import (
	"fmt"
	"net/http"

	"gza/comment/models"

	"github.com/gin-gonic/gin"
)

func GetComments(c *gin.Context) {
	var comments []models.Comment
	db := models.GetDatabase()
	fmt.Println(db.Debug())
	r := db.Model(comments).Find(&comments)
	if r.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "데이터가 존재하지 않습니다",
		})
		return
	}
	c.JSON(http.StatusOK, &comments)
}

func GetComment(c *gin.Context) {
	var comment models.Comment
	db := models.GetDatabase()
	r := db.First(&comment, "id = ?", c.Param("id"))
	if r.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "데이터가 존재하지 않습니다",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   comment,
	})
}
