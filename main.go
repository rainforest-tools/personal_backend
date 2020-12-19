package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendString("Hello, this is Rainforest 🌧🌲🌲")
	})

	port := getPort()
	app.Listen(fmt.Sprintf(":%s", port))
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
		log.Printf("Defaulting to port %s", port)
	}

	return port
}
