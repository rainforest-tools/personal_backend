package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New(fiber.Config{Prefork: true})
	// middleware
	app.Use(recover.New())
	app.Use(cors.New())
	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendString("Hello, this is Rainforest 🌧🌲🌲")
	})

	app.Static("/static", "./static")

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
