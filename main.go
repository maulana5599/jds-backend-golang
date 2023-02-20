package main

import (
	"jdsapp/config"
	"jdsapp/handler"
	"jdsapp/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v3"
)

func main() {
	app := fiber.New()
	config.DatabaseConnection()

	app.Use(cors.New())

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	routes.Routes(app)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("MooAZePGCWmF8M4SIXeqDkhd1OKNN9DMd1N2AVuQDm3xqdCXJwqcIdYH61UibOxC"),
	}))

	routes.RestrictedRoutes(app)

	app.Get("/test", handler.TestJwt)
	app.Listen(":3001")
}
