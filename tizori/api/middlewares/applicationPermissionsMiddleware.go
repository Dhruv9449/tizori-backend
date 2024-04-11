package middlewares

import (
	"fmt"

	"github.com/GDGVIT/Tizori-backend/internal/models"
	"github.com/gofiber/fiber/v2"
)

// ApplicationPermissionsMiddleware is a go-fiber middleware to check if the user has a specific permission (global or application-specific)
func ApplicationPermissionsMiddleware(permission models.ApplicationPermission) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extract roles from fiber.Ctx locals
		roles, ok := c.Locals("roles").([]models.Role)
		if !ok {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"detail": "Roles not found in context",
			})
		}

		// Get application ID from URL params
		applicationID := c.Params("id")

		fmt.Println("Application ID: ", applicationID)
		fmt.Println("Permission: ", permission)

		// Check if the user has the required permission
		for _, role := range roles {
			for _, rolePermission := range role.Permissions {
				if rolePermission.Scope == "global" &&
					models.GlobalPermission(rolePermission.Permission) == models.WriteApplications {
					// User has global read applications permission, proceed to next middleware/handler
					return c.Next()
				}
				if rolePermission.Scope == applicationID &&
					models.ApplicationPermission(rolePermission.Permission) == permission {
					// User has the required permission, proceed to next middleware/handler
					fmt.Println("User has permission")
					return c.Next()
				}
			}
		}

		// If permission not found in user's roles, deny access
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"detail": "You do not have permission to perform this action",
		})
	}
}
