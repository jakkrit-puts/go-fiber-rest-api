package userctrl

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jakkrit-puts/go-fiber-rest-api/configs"
	"github.com/jakkrit-puts/go-fiber-rest-api/models"
)

func GetAll(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": "Users",
	})
}

func GetById(c *fiber.Ctx) error {
	id := c.Params("id")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": "User:" + id,
	})
}

func Search(c *fiber.Ctx) error {
	q := c.Query("q")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": "Search:" + q,
	})
}

func Login(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": "Login",
	})
}

func Register(c *fiber.Ctx) error {

	var payload *models.CreateUserPayload

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	now := time.Now()
	user := models.User{
		Username:  payload.Username,
		Password:  payload.Password,
		Fullname:  payload.Fullname,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := configs.DB.Create(&user)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "Username already exist, please use another username"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": user})

}
