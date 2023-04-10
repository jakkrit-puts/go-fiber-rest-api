package userctrl

import "github.com/gofiber/fiber/v2"

func GetAll(c *fiber.Ctx) error {
	err := c.SendString("User....")
	return err
}

func Login(c *fiber.Ctx) error {
	err := c.SendString("login.")
	return err
}

func Register(c *fiber.Ctx) error {
	err := c.SendString("register.")
	return err
}
