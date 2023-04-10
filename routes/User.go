package routes

import (
	"github.com/gofiber/fiber/v2"
	userctrl "github.com/jakkrit-puts/go-fiber-rest-api/controllers/user"
)

func UserRoutes(rg fiber.Router) {

	// Prefix
	routerGroup := rg.Group("/users")

	routerGroup.Get("/", userctrl.GetAll)

	routerGroup.Post("/login", userctrl.Login)

	routerGroup.Post("/register", userctrl.Register)
}
