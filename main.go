package main

import (
	"log"

	"errors"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	router := fiber.New()
	app := fiber.New()

	app.Mount("/api", router)

	router.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "up",
			"message": "welcome",
			"detail":  "Detailed message to display",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal(app.Listen(":8090"))
	} else {
		log.Fatal(app.Listen(":" + port))
	}
}

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return name, errors.New("empty name")
	}
	// Create a message using a random format.
	// message := fmt.Sprintf(randomFormat(), name)
	message := "Hi, " + name
	return message, nil
}
