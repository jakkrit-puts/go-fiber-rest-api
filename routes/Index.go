package routes

import "github.com/gofiber/fiber/v2"

func IndexRoutes(rg fiber.Router) {

	// Prefix
	routerGroup := rg.Group("")

	routerGroup.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"API_INFO": "Fiber Basic API V.1.0.0",
		})
	})
}
