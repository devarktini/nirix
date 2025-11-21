package routes

import "github.com/gofiber/fiber/v3"

func SetupUserRoutes(baseRoute fiber.Router) {
	userGroup := baseRoute.Group("/user")
	userGroup.Add([]string{"GET"}, "", func(c fiber.Ctx) error {
		// Handle login
		return c.SendString("Login endpoint")
	})
}
