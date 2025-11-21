package routes

import "github.com/gofiber/fiber/v3"

func SetupOpsRoutes(rg fiber.Router) {
	opsGroup := rg.Group("/ops")
	opsGroup.Get("/health", healthCheckHandler)
}

func healthCheckHandler(c fiber.Ctx) error {
	return c.JSON(map[string]string{
		"status": "ok",
	})
}
