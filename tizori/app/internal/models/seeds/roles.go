package seeds

import (
	"github.com/GDGVIT/Tizori-backend/internal/database"
	"github.com/GDGVIT/Tizori-backend/internal/models"
	"github.com/google/uuid"
)

var DEFAULT_ROLES = []models.Role{
	{
		Name: "super-admin",
		Permissions: []models.Permission{
			{
				Scope:      "global",
				Permission: string(models.ReadUsers),
			},
			{
				Scope:      "global",
				Permission: string(models.WriteUsers),
			},
			{
				Scope:      "global",
				Permission: string(models.ReadRoles),
			},
			{
				Scope:      "global",
				Permission: string(models.WriteRoles),
			},
			{
				Scope:      "global",
				Permission: string(models.ReadApplications),
			},
			{
				Scope:      "global",
				Permission: string(models.WriteApplications),
			},
		},
	},
	{
		Name: "admin",
		Permissions: []models.Permission{
			{
				Scope:      "global",
				Permission: string(models.ReadUsers),
			},
			{
				Scope:      "global",
				Permission: string(models.ReadRoles),
			},
			{
				Scope:      "global",
				Permission: string(models.WriteRoles),
			},
			{
				Scope:      "global",
				Permission: string(models.ReadApplications),
			},
			{
				Scope:      "global",
				Permission: string(models.WriteApplications),
			},
		},
	},
	{
		Name: "user",
		Permissions: []models.Permission{
			{
				Scope:      "global",
				Permission: string(models.ReadUsers),
			},
			{
				Scope:      "global",
				Permission: string(models.ReadRoles),
			},
			{
				Scope:      "global",
				Permission: string(models.ReadApplications),
			},
		},
	},
}

// seedRoles seeds the default roles in the database
func seedRoles() error {
	tx := database.DB.Begin() // Start a new database transaction
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback() // Rollback the transaction if a panic occurs
		}
	}()

	for _, role := range DEFAULT_ROLES {
		uuid, err := uuid.NewUUID()
		if err != nil {
			return err
		}
		role.Id = uuid.String()
		if err = tx.Create(&role).Error; err != nil {
			tx.Rollback() // Rollback the transaction on error
			return err
		}
	}

	return tx.Commit().Error // Commit the transaction
}
