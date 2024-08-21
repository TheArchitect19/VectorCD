package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"vectordb/connect"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Default port if not specified
	}
	client := connect.ConnectDB()
	
	app := fiber.New()

	// Define a simple GET route using the environment variable
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!")
	})
	log.Fatal(app.Listen(":" + port))
}
