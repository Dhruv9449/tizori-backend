package api

import (
	v1 "github.com/GDGVIT/Tizori-backend/api/v1"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewWebApi() *fiber.App {
	// New fiber app
	fiberApp := fiber.New()

	fiberApp = fiber.New()
	fiberApp.Use(logger.New())
	fiberApp.Use(cors.New(
		cors.Config{
			AllowOrigins: "*",
			AllowHeaders: "*",
			AllowMethods: "GET,POST,DELETE,PATCH,PUT,OPTIONS",
		},
	))

	// Root endpoint
	fiberApp.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("Welcome to Tizori API!🎉")
	})

	// Ping endpoint
	fiberApp.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(
			fiber.Map{
				"detail": "pong",
			})
	})

	api := fiberApp.Group("/api")
	v1.V1Handler(api)

	return fiberApp
}
