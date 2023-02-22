package handler

import (
	"jdsapp/entity"
	"jdsapp/helpers"
	"jdsapp/repository"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Auth(ctx *fiber.Ctx) error {

	requestAuth := new(entity.RequestAuth)

	if err := ctx.BodyParser(requestAuth); err != nil {
		return helpers.ResponseError(http.StatusBadRequest, err.Error(), ctx)
	}
	// Cek user
	check, passwordUser := repository.Authentication(requestAuth.Nik, requestAuth.Password)

	if !check {
		return fiber.ErrUnauthorized
	}

	if passwordUser == "" {
		return fiber.ErrUnauthorized
	}

	match, _ := helpers.CheckPasswordHash(requestAuth.Password, passwordUser)

	if !match {
		return fiber.ErrUnauthorized
	}

	// Create claims
	claims := jwt.MapClaims{
		"admin": true,
		"role":  []string{"admin", "user"},
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	privateKey := helpers.GoDotEnvVariable("JWT_SECRET")

	t, err := token.SignedString([]byte(privateKey))
	if err != nil {
		log.Printf("token.SignedString: %v", err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	result := fiber.Map{
		"token": t,
	}
	return helpers.ResponseSuccess(http.StatusOK, "Authentication Successfully", ctx, result)
}

func TestJwt(c *fiber.Ctx) error {
	return c.JSON(fiber.StatusOK)
}
