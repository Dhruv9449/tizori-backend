package v1

import (
	"github.com/GDGVIT/Tizori-backend/api/middlewares"
	"github.com/GDGVIT/Tizori-backend/api/serializers"
	"github.com/GDGVIT/Tizori-backend/internal/auth"
	"github.com/GDGVIT/Tizori-backend/internal/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// userHandler is a function to handle all the user related routes
func usersHandler(api fiber.Router) {
	group := api.Group("/users")

	group.Use(middlewares.JWTAuthMiddleware)
	group.Use(middlewares.GlobalPermissionsMiddleware(models.ReadUsers))
	group.Get("/", getUsers)
	group.Get("/:username", getUser)

	group.Use(middlewares.GlobalPermissionsMiddleware(models.WriteUsers))
	group.Post("/", createUser)
	group.Patch("/:username", updateUser)
	group.Delete("/:username", deleteUser)
}

// getUsers is a function to get all the users
func getUsers(c *fiber.Ctx) error {
	users, err := models.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(serializers.UsersListSerializer(users))
}

// getUser is a function to get a user by username
func getUser(c *fiber.Ctx) error {
	username := c.Params("username")
	user, err := models.GetUserByUsername(username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(serializers.UserSerializer(*user))
}

// createUser is a function to create a new user
func createUser(c *fiber.Ctx) error {
	type requestBody struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Name     string `json:"name"`
	}

	var body requestBody
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	// Check if the username is valid
	if val, msg := models.ValidateUsername(body.Username); !val {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": msg,
		})
	}

	// Check if the email exists
	if models.CheckEmailExists(body.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "Email already exists",
		})
	}

	// Create a new user
	user := models.User{
		Username: body.Username,
		Email:    body.Email,
		Name:     body.Name,
	}

	// Generate a random password
	password := auth.GeneratePassword(12, true, true, true)

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	user.Password = string(hashedPassword)

	// Save the user to the database
	err = user.Save()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"detail":   "User created successfully",
		"username": user.Username,
		"password": password,
	})
}

// updateUser is a function to update a user
func updateUser(c *fiber.Ctx) error {
	type requestBody struct {
		Email string   `json:"email"`
		Name  string   `json:"name"`
		Roles []string `json:"roles"`
	}

	var body requestBody
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	username := c.Params("username")
	user, err := models.GetUserByUsername(username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	if body.Email != "" {
		// Check if the email exists
		user.Email = body.Email
	}
	if body.Name != "" {
		user.Name = body.Name
	}

	// Update the roles
	var roles []models.Role
	for _, roleId := range body.Roles {
		role, err := models.GetRoleById(roleId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"detail": err.Error(),
			})
		}
		roles = append(roles, *role)
	}
	user.Roles = roles

	err = user.Save()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"detail": "User updated successfully",
	})
}

// deleteUser is a function to delete a user
func deleteUser(c *fiber.Ctx) error {
	username := c.Params("username")
	user, err := models.GetUserByUsername(username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	err = user.Delete()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"detail": "User deleted successfully",
	})
}
