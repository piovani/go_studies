package main

import (
	"go_gin/database"
	"go_gin/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}
