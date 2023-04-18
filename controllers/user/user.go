package userctrl

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/jakkrit-puts/go-fiber-rest-api/configs"
	"github.com/jakkrit-puts/go-fiber-rest-api/models"
	"github.com/jakkrit-puts/go-fiber-rest-api/utils"
	"github.com/matthewhartstonge/argon2"
)

func Login(c *fiber.Ctx) error {

	var payload LoginSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	user := models.User{
		Username: payload.Username,
		Password: payload.Password,
	}

	hasUser := configs.DB.Where("username = ?", payload.Username).First(&user)

	if hasUser.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "Username Not Found"})
	}

	// compare password from model
	passwordMatch, _ := argon2.VerifyEncoded([]byte(payload.Password), []byte(user.Password))
	if !passwordMatch {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Password invalid"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(), // set expire token (time.Hour * 24 * 7) = 7days
	})

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	tokenString, _ := token.SignedString([]byte(jwtSecretKey))

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Login Success", "access_token": tokenString})
}

func Register(c *fiber.Ctx) error {

	var payload CreateUserSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	now := time.Now()
	user := models.User{
		Username:  payload.Username,
		Password:  payload.Password,
		Fullname:  payload.Fullname,
		CreatedAt: now,
		UpdatedAt: now,
	}

	checkUserExist := configs.DB.Where("username = ?", payload.Username).First(&user)

	result := configs.DB.Create(&user)

	if checkUserExist.RowsAffected == 1 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "error", "message": "Username already exist."})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "result": user})
}

func GetAll(c *fiber.Ctx) error {

	var users []models.User
	results := configs.DB.Scopes(utils.Paginate(c)).Find(&users)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "total": len(users), "result": users})
}

func GetById(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.User

	result := configs.DB.First(&user, id)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "record not found."})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": user,
	})
}

func GetBySearch(c *fiber.Ctx) error {
	fullname := c.Query("fullname")

	var users []models.User
	var count int64

	result := configs.DB.Where("fullname LIKE ?", "%"+fullname+"%").Scopes(utils.Paginate(c)).Find(&users).Count(&count)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "record not found."})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": users,
		"total":  count,
	})
}

func GetProfile(c *fiber.Ctx) error {

	user := c.Locals("user").(models.User)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": user,
	})
}
