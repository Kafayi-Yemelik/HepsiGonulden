package main

import (
	"HepsiGonulden/Customer"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	Customer.NewHandler(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
