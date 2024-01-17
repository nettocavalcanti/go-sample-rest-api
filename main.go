package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	router := fiber.New()
	app := fiber.New()

	app.Mount("/api", router)

	router.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status": "up",
			"message": "welcome",
		})
	})

	log.Fatal(app.Listen(":8090"))
}
