package main

import (
	"fmt"
	"log"
)

const (
	apiKey = ""
	city   = "Maringá, Brazil"
)

func main() {
	openweatherProvider := NewProvider(apiKey)
	weatherService := NewService(openweatherProvider)

	outfit, err := weatherService.WhatToWear(city)
	if err != nil {
		log.Fatalf("couldn't get what to wear in %s: %v", city, err)
	}

	fmt.Printf("you should wear in %s: %s\n", city, outfit)
}
