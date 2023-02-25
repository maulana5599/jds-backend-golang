package main

import (
	"jdsapp/config"
	"jdsapp/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	config.DatabaseConnection()
	godotenv.Load(".env")

	app.Use(cors.New())

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	routes.Routes(app)

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Selamat datang!")
	})

	routes.RestrictedRoutes(app)

	app.Listen(":3001")
}
