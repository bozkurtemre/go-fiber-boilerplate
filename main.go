package main

import (
	"os"

	"flag"
	"log"

	"boilerplate/database"
	"boilerplate/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/joho/godotenv/autoload"
)

var (
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

func main() {
	// Get port from .env file
	port := os.Getenv("APP_PORT")

	// Parse command-line flags
	flag.Parse()

	// Connect with database
	database.Connect()

	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork: *prod, // go run app.go -prod
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// Health Check
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "OK",
		})
	})

	// Group
	v1 := app.Group("/api/v1")

	// Routes
	users := v1.Group("/users")
	users.Get("/", handlers.UserList)
	users.Post("/", handlers.UserCreate)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"message": "Not Found",
		})
	})

	// Listen on port 8080
	log.Fatal(app.Listen(":" + port))
}
