package models

import (
	"errors"

	"github.com/GDGVIT/Tizori-backend/internal/database"
)

type GlobalPermission string

// GlobalScopes is a type that represents global scopes
const (
	ReadRoles         GlobalPermission = "read_roles"
	ReadUsers         GlobalPermission = "read_users"
	ReadApplications  GlobalPermission = "read_applications"
	WriteRoles        GlobalPermission = "write_roles"
	WriteUsers        GlobalPermission = "write_users"
	WriteApplications GlobalPermission = "write_applications"
)

type ApplicationPermission string

// ApplicationScopes is a type that represents application scopes
const (
	ReadApplicationCredentials  ApplicationPermission = "read_application_credentials"
	WriteApplicationCredentials ApplicationPermission = "write_application_credentials"
)

type PermissionType interface {
	isPermissionType()
}

func (GlobalPermission) isPermissionType() {}

func (ApplicationPermission) isPermissionType() {}

type Permission struct {
	Scope      string `json:"scope"`
	Permission string `json:"permission"`
}

type Role struct {
	Id          string       `gorm:"primaryKey"`
	Name        string       `gorm:"not null,unique"`
	Permissions []Permission `gorm:"serializer:json"`
}

func (r *Role) Save() error {
	// Validate permissions
	for _, p := range r.Permissions {
		err := validatePermission(p)
		if err != nil {
			return err
		}
	}
	return database.DB.Save(&r).Error
}

func (r *Role) Delete() error {
	return database.DB.Delete(&r).Error
}

func GetRoles() ([]Role, error) {
	var roles []Role
	err := database.DB.Find(&roles).Error
	return roles, err
}

func GetRoleById(id string) (*Role, error) {
	var role Role
	err := database.DB.Where("id = ?", id).First(&role).Error
	return &role, err
}

func GetRoleByName(name string) (*Role, error) {
	var role Role
	err := database.DB.Where("name = ?", name).First(&role).Error
	return &role, err
}

func validatePermission(p Permission) error {
	// Scope should be either "global" or name of the application
	if p.Scope != "global" {
		_, err := GetApplicationById(p.Scope)
		if err != nil {
			return err
		}
	}
	// Permission should be one of the predefined permissions
	switch p.Scope {
	case "global":
		switch GlobalPermission(p.Permission) {
		case ReadRoles, ReadUsers, ReadApplications, WriteRoles, WriteUsers, WriteApplications:
			return nil
		default:
			return errors.New("invalid permission")
		}
	default:
		switch ApplicationPermission(p.Permission) {
		case ReadApplicationCredentials, WriteApplicationCredentials:
			return nil
		default:
			return errors.New("invalid permission")
		}
	}
}
