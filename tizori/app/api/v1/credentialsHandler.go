package v1

import (
	"github.com/GDGVIT/Tizori-backend/api/middlewares"
	"github.com/GDGVIT/Tizori-backend/api/serializers"
	"github.com/GDGVIT/Tizori-backend/internal/models"
	"github.com/gofiber/fiber/v2"
)

func credentialsHandler(api fiber.Router) {
	group := api.Group("/credentials")

	group.Use(middlewares.JWTAuthMiddleware)
	group.Get("/:id",
		middlewares.ApplicationPermissionsMiddleware(models.ReadApplicationCredentials),
		getCredentials)

	group.Patch("/:id",
		middlewares.ApplicationPermissionsMiddleware(models.WriteApplicationCredentials),
		updateCredentials)
}

func getCredentials(c *fiber.Ctx) error {
	id := c.Params("id")
	application, err := models.GetApplicationById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}
	credentials, err := models.GetApplicationCredentials(application.Id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(serializers.ApplicationCredentialsSerializer(*application, credentials))
}

func updateCredentials(c *fiber.Ctx) error {
	id := c.Params("id")
	application, err := models.GetApplicationById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	var body models.ApplicationCredentials
	err = c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	err = models.StoreApplicationCredentials(application.Id, body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"detail": "Credentials updated successfully",
	})
}
