package routes

import "github.com/gofiber/fiber/v2"

func IndexRoutes(rg fiber.Router) {

	// Prefix
	routerGroup := rg.Group("")

	routerGroup.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("Fiber Stock Mini API V.1.0.0")
		return err
	})
}
