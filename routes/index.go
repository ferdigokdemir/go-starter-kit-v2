package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(r *fiber.App) {
	r.Static("/", "./public")
	api := r.Group("/api")
	TodoRoute(api.Group("/todos"))
}
