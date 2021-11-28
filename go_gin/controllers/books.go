package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowBooks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"value": "ok",
	})
}
