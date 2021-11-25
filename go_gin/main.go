package main

import (
	"github.com/gin-gonic/gin"
)

// func main() {
// 	server := server.NewServer()g

// 	server.Run()
// }

func main() {
	route := gin.Default()
	route.Any("/testing", startPage)
	route.Run(":8085")
}

func startPage(c *gin.Context) {
	c.String(200, "Success")
}
