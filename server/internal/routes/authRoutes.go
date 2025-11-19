package routes

import "github.com/gofiber/fiber/v3"

func SetupAuthRoutes(baseRoute fiber.Router) {
	authGroup := baseRoute.Group("/auth")

	authGroup.Post("/login", func(c fiber.Ctx) error {
		// Handle login
		return c.SendString("Login endpoint")
	})

	authGroup.Post("/register", func(c fiber.Ctx) error {
		// Handle registration
		return c.SendString("Register endpoint")
	})

	authGroup.Post("/logout", func(c fiber.Ctx) error {
		// Handle logout
		return c.SendString("Logout endpoint")
	})

}
