package helpers

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func ResponseSuccess(status int, message string, ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(status).JSON(fiber.Map{
		"status":  status,
		"message": message,
		"data":    data,
	})
}

func ResponseError(status int, message string, ctx *fiber.Ctx) error {
	return ctx.Status(status).JSON(fiber.Map{
		"status":  status,
		"message": message,
	})
}

/**
* Function for hashing password
 */
func HashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(bytes), err
}

/**
* Function for checking password
 */
func CheckPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if err != nil {
		return false, nil
	}

	return true, nil
}
