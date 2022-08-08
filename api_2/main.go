package main

import "github.com/piovani/go_api/cmd"

func main() {
	api := cmd.NewApi()
	api.Start()
}
