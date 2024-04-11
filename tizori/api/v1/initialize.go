package v1

import (
	"github.com/gofiber/fiber/v2"
)

func V1Handler(api fiber.Router) {
	group := api.Group("/v1")
	authHandler(group)
	usersHandler(group)
	rolesHandler(group)
	applicationsHandler(group)
	credentialsHandler(group)
}
