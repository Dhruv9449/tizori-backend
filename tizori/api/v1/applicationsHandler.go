package v1

import (
	"github.com/GDGVIT/Tizori-backend/api/middlewares"
	"github.com/GDGVIT/Tizori-backend/api/serializers"
	"github.com/GDGVIT/Tizori-backend/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func applicationsHandler(api fiber.Router) {
	group := api.Group("/applications")

	group.Use(middlewares.JWTAuthMiddleware)
	group.Use(middlewares.GlobalPermissionsMiddleware(models.ReadApplications))
	group.Get("/", getApplications)
	group.Get("/:id", getApplication)

	group.Use(middlewares.GlobalPermissionsMiddleware(models.WriteApplications))
	group.Post("/", createApplication)
	group.Patch("/:id", updateApplication)
	group.Delete("/:id", deleteApplication)
}

func getApplications(c *fiber.Ctx) error {
	applications, err := models.GetApplications()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(serializers.ApplicationsListSerializer(applications))
}

func getApplication(c *fiber.Ctx) error {
	id := c.Params("id")
	application, err := models.GetApplicationById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(serializers.ApplicationSerializer(*application))
}

func createApplication(c *fiber.Ctx) error {
	type requestBody struct {
		Name  string `json:"name"`
		Owner string `json:"owner"`
	}

	var body requestBody
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	uuid, err := uuid.NewUUID()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}
	application := models.Application{
		Id:   uuid.String(),
		Name: body.Name,
	}

	err = application.Save()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(serializers.ApplicationSerializer(application))
}

func updateApplication(c *fiber.Ctx) error {
	type requestBody struct {
		Name string `json:"name"`
	}

	var body requestBody
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	id := c.Params("id")
	application, err := models.GetApplicationById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	application.Name = body.Name

	err = application.Save()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(serializers.ApplicationSerializer(*application))
}

func deleteApplication(c *fiber.Ctx) error {
	id := c.Params("id")
	application, err := models.GetApplicationById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	err = application.Delete()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"detail": "Application deleted successfully",
	})
}
