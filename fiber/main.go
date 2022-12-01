package main

import (
	"github.com/gofiber/fiber/v2"
)

var HelloWorld = func(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func main() {
	app := fiber.New()

	app.Get("/", HelloWorld)

	app.Listen(":3000")
}
