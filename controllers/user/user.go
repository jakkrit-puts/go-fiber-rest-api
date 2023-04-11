package userctrl

import "github.com/gofiber/fiber/v2"

func GetAll(c *fiber.Ctx) error {
	err := c.SendString("User....")
	return err
}

func GetById(c *fiber.Ctx) error {
	id := c.Params("id")

	err := c.SendString("Use ID:" + id)
	return err
}

func Search(c *fiber.Ctx) error {
	q := c.Query("q")
	err := c.SendString("result :" + q)
	return err
}

func Login(c *fiber.Ctx) error {
	err := c.SendString("login...")
	return err
}

func Register(c *fiber.Ctx) error {
	err := c.SendString("register.")
	return err
}
