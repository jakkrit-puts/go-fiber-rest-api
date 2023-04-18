package routes

import (
	"github.com/gofiber/fiber/v2"
	userctrl "github.com/jakkrit-puts/go-fiber-rest-api/controllers/user"
	"github.com/jakkrit-puts/go-fiber-rest-api/middlewares"
)

func UserRoutes(rg fiber.Router) {

	// Prefix
	routerGroup := rg.Group("/users")

	routerGroup.Get("/", userctrl.GetAll)

	routerGroup.Get("/me", middlewares.AuthJWT, userctrl.GetProfile)

	routerGroup.Get("/search", userctrl.GetBySearch)

	routerGroup.Get("/:id", userctrl.GetById)

	routerGroup.Post("/login", userctrl.Login)

	routerGroup.Post("/register", userctrl.Register)

}
