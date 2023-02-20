package routes

import (
	"jdsapp/handler"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	AuthHandler(app)
}

func AuthHandler(app *fiber.App) {
	app.Post("/auth/login", handler.Auth)
}
