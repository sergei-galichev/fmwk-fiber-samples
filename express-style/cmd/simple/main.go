package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {

	app := fiber.New()

	app.Static("/", "./web/static")
	app.Use("/hello", handler()).Get("/hello", hello)

	app.Get("/middleware", handler())

	// Middleware: not matching other routes
	app.Use(
		func(c *fiber.Ctx) error {
			err := c.SendStatus(fiber.StatusNotFound)
			if err != nil {
				log.Fatal(err)
				return err
			}

			return nil
		},
	)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func handler() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.SendString("Hello, from middleware!")
	}
}
