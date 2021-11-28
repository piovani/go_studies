package controllers

import (
	"net/http"
	"strconv"

	"go_gin/database"
	"go_gin/models"

	"github.com/gin-gonic/gin"
)

func ShowBooks(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID has to be integer",
		})
	}

	db := database.GetDatabase()

	var book models.Book
	err = db.First(&book, newId).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Has a error internal: " + err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, book)
}
