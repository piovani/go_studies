package controllers

import (
	"net/http"
	"strconv"

	"go_gin/database"
	"go_gin/models"

	"github.com/gin-gonic/gin"
)

func ShowBook(c *gin.Context) {
	id := c.Param("id")
	newId, err := strconv.ParseInt(id, 10, 64)
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

func CreateBook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Has a error internal: " + err.Error(),
		})

		return
	}

	err = db.Create(&book).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Has a error internal: " + err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, book)
}

func ShowBooks(c *gin.Context) {
	db := database.GetDatabase()

	var books []models.Book
	err := db.Find(&books).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Has a error internal: " + err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, books)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	newId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID has to be integer",
		})
	}

	var book models.Book
	err = c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Has a error internal: " + err.Error(),
		})

		return
	}

	book.ID = newId

	db := database.GetDatabase()
	err = db.Save(&book).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Has a error internal: " + err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	newId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID has to be integer",
		})
	}

	db := database.GetDatabase()
	err = db.Delete(models.Book{}, newId).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Has a error internal: " + err.Error(),
		})

		return
	}

	c.Status(http.StatusNoContent)
}
