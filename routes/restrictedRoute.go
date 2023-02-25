package routes

import (
	"jdsapp/handler"
	"jdsapp/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RestrictedRoutes(app *fiber.App) {
	Product(app)
}

func Product(app *fiber.App) {
	api := app.Group("api")
	api.Get("/product", middlewares.Protected(), handler.GetProduct)
}
