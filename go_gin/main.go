package main

import (
	"go_gin/server"
)

func main() {
	server := server.NewServer()

	server.Run()
}
