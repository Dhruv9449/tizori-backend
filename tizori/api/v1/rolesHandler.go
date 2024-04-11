package v1

import (
	"github.com/GDGVIT/Tizori-backend/api/middlewares"
	"github.com/GDGVIT/Tizori-backend/api/serializers"
	"github.com/GDGVIT/Tizori-backend/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func rolesHandler(api fiber.Router) {
	group := api.Group("/roles")

	group.Use(middlewares.JWTAuthMiddleware)
	group.Use(middlewares.GlobalPermissionsMiddleware(models.ReadRoles))
	group.Get("/", getRoles)
	group.Get("/:id", getRole)

	group.Use(middlewares.GlobalPermissionsMiddleware(models.WriteRoles))
	group.Post("/", createRole)
	group.Patch("/:id", updateRole)
	group.Delete("/:id", deleteRole)
}

func getRoles(c *fiber.Ctx) error {
	roles, err := models.GetRoles()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(serializers.RolesListSerializer(roles))
}

func getRole(c *fiber.Ctx) error {
	id := c.Params("id")
	role, err := models.GetRoleById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(serializers.RoleSerializer(*role))
}

func createRole(c *fiber.Ctx) error {
	type requestBody struct {
		Name       string              `json:"name"`
		Permission []models.Permission `json:"permissions"`
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

	role := models.Role{
		Id:          uuid.String(),
		Name:        body.Name,
		Permissions: body.Permission,
	}

	err = role.Save()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(serializers.RoleSerializer(role))
}

func updateRole(c *fiber.Ctx) error {
	type requestBody struct {
		Name       string              `json:"name"`
		Permission []models.Permission `json:"permissions"`
	}

	var body requestBody
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	id := c.Params("id")
	role, err := models.GetRoleById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	role.Name = body.Name
	role.Permissions = body.Permission

	err = role.Save()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(serializers.RoleSerializer(*role))
}

func deleteRole(c *fiber.Ctx) error {
	id := c.Params("id")
	role, err := models.GetRoleById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	err = role.Delete()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"detail": "Role deleted successfully",
	})
}
