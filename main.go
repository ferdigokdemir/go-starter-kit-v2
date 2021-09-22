package main

import (
	"go_starter_kit/config"
	"go_starter_kit/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// database connect
	config.ConnectDB()
}

func main() {
	app := fiber.New()

	routes.SetupRoutes(app)

	err := app.Listen(":" + os.Getenv("PORT"))

	if err != nil {
		panic(err)
	}
}
