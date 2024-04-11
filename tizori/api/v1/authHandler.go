package v1

import (
	"strings"

	"github.com/GDGVIT/Tizori-backend/api/serializers"
	"github.com/GDGVIT/Tizori-backend/internal/auth"
	"github.com/GDGVIT/Tizori-backend/internal/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func authHandler(api fiber.Router) {
	group := api.Group("/auth")
	group.Post("/check-username", checkUsernameValidity)
	group.Post("/check-user-exists", checkUserExists)
	group.Post("/login", login)
}

func checkUsernameValidity(c *fiber.Ctx) error {
	type RequestBody struct {
		Username string `json:"username"`
	}
	var body RequestBody
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}
	username := strings.ToLower(body.Username)
	if val, msg := models.ValidateUsername(username); !val {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": msg,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"detail": "Username is valid",
	})
}

func checkUserExists(c *fiber.Ctx) error {
	type RequestBody struct {
		UUID string `json:"uuid"`
	}

	var body RequestBody
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	if !models.CheckUserUUIDExists(body.UUID) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "User does not exist",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"detail": "User exists",
	})
}

func login(c *fiber.Ctx) error {
	type RequestBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var body RequestBody
	err := c.BodyParser(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	// Get user from database
	user, err := models.GetUserByUsername(body.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	// Check if the password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"detail": "Invalid credentials",
		})
	}

	// Create JWT token
	token, err := auth.CreateJWTToken(user.Username, user.Email, auth.JWTSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(serializers.UserLoginSerializer(*user, token))
}
