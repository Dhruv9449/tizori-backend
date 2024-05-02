package commands

import (
	"fmt"

	"github.com/GDGVIT/Tizori-backend/internal/models"
	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/bcrypt"
)

var userCommands = []*cli.Command{
	{
		Name:    "createsuperuser",
		Aliases: []string{"csu"},
		Usage:   "Create a superuser",
		Action:  createSuperuser,
	},
}

func createSuperuser(c *cli.Context) error {
	fmt.Println("Enter username: ")
	var username string
	fmt.Scanln(&username)
	var user *models.User
	var err error

	if !models.CheckUsernameExists(username) {
		fmt.Println("Enter email: ")
		var email string
		fmt.Scanln(&email)
		fmt.Println("Enter password: ")
		var password string
		fmt.Scanln(&password)
		// Hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user = &models.User{
			Username: username,
			Email:    email,
			Password: string(hashedPassword),
		}
		err = user.Save()
		if err != nil {
			return err
		}
	} else {
		user, err = models.GetUserByUsername(username)
		if err != nil {
			return err
		}
	}

	suAdminRole, err := models.GetRoleByName("super-admin")
	if err != nil {
		return err
	}

	user.Roles = append(user.Roles, *suAdminRole)
	err = user.Save()
	if err != nil {
		return err
	}
	fmt.Println("Superuser created successfully")
	return nil
}
