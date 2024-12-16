package main

import (
	"fmt"
	"go_programming/config"
	"go_programming/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Dev code app running...")

	// Initialize database connection
	db, err := config.DatabaseConfig()
	if err != nil {
		panic(err)
	}

	// Initialize Fiber app
	app := fiber.New()

	// Setup CORS middleware
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Set("Access-Control-Allow-Credentials", "true")

		if c.Method() == "OPTIONS" {
			c.SendStatus(fiber.StatusNoContent)
			return nil
		}

		return c.Next()
	})

	// Setup routes
	router.SetupRoutes(app, db)

	// Start the server
	app.Listen(":8080")
}
