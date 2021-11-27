package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/golang-jwt/jwt/v4"

	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "github.com/rainforest-tools/personal_backend/docs"
)

// @title Rainforest Personal Backend
// @version 0.1

// @contact.name Rainforest
// @contact.url http://rainforest.tools
// @contact.email rainforestnick@gmail.com

// @host localhost:3000
// @BasePath /
func main() {
	app := fiber.New(fiber.Config{Prefork: true})
	// middleware
	app.Use(recover.New())
	app.Use(cors.New())

	app.Get("/about", About)

	// docs
	app.Get("/swagger/*", swagger.Handler)
	app.Get("/health", HealthCheck)

	// auth
	app.Post("/login", login)

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

// Auth
func login(c *fiber.Ctx) error {
	user := c.FormValue("user")
	pass := c.FormValue("pass")

	if user != "rainforest" || pass != "rainforest" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	claims := jwt.MapClaims{
		"name":  "Rainforest",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

// About
// @Summary About
// @Router /about [get]
// @Produce text/plain
// @Success      200  {object}  string
func About(c *fiber.Ctx) error {
	return c.SendString("Hello, this is Rainforest 🌧🌲🌲")
}

// HealthCheck
// @Summary Show the status of server
// @Router /health [get]
func HealthCheck(c *fiber.Ctx) error {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}
	if err := c.JSON(res); err != nil {
		return err
	}
	return nil
}
