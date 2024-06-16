package main

import (
	"os"

	"flag"
	"log"

	"boilerplate/database"
	_ "boilerplate/docs"
	"boilerplate/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	_ "github.com/joho/godotenv/autoload"
)

var (
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

// @title						Fiber Example API
// @version					    1.0
// @description				    This is a sample swagger for Fiber
// @host						localhost:8080
// @BasePath					/api/v1
// @securityDefinitions.apikey	JWT
// @in							header
// @name						Authorization
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

	// Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

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
	log.Fatal(app.Listen(":" + port)) // go run app.go -port=:8080
}
