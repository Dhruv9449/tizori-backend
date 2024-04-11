package middlewares

import (
	"github.com/GDGVIT/Tizori-backend/internal/auth"
	"github.com/gofiber/fiber/v2"
)

// JWTAuthMiddleware is a go-fiber middleware to authenticate the user and add the user to the context if authenticated
func JWTAuthMiddleware(c *fiber.Ctx) error {
	// Get the JWT token from the Authorization header
	authorizationString := c.Get("Authorization")

	// If the Authorization header is not present, return an error
	if authorizationString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"detail": "Authorization header not present",
		})
	}

	// Get user from JWT token
	user, parseErr := auth.GetUserFromJWTToken(authorizationString, auth.JWTSecret)
	if parseErr != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"detail": parseErr.Error(),
		})
	}

	// Add the user to the context
	c.Locals("user", user)
	// Add the user's roles to the context
	c.Locals("roles", user.Roles)
	return c.Next()
}
