package models

import (
	"github.com/GDGVIT/Tizori-backend/internal/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	Username string `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique,not null"`
	Password string `gorm:"not null"`
	Roles    []Role `gorm:"many2many:user_roles;"`
}

func (u *User) Save() error {
	return database.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&u).Error
}

func (u *User) Delete() error {
	return database.DB.Delete(&u).Error
}

func GetUsers() ([]User, error) {
	var users []User
	err := database.DB.Find(&users).Error
	return users, err
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	err := database.DB.Where("username = ?", username).Preload(clause.Associations).First(&user).Error
	return &user, err
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	err := database.DB.Where("email = ?", email).Preload(clause.Associations).First(&user).Error
	return &user, err
}

func GetUserByUUID(uuid string) (*User, error) {
	var user User
	err := database.DB.Where("firebase_uuid = ?", uuid).Preload(clause.Associations).First(&user).Error
	return &user, err
}

func CheckUsernameExists(username string) bool {
	var count int64
	database.DB.Model(&User{}).Where("username = ?", username).Count(&count)
	return count != 0
}

func CheckEmailExists(email string) bool {
	var count int64
	database.DB.Model(&User{}).Where("email = ?", email).Count(&count)
	return count != 0
}

func CheckUserUUIDExists(uuid string) bool {
	var count int64
	database.DB.Model(&User{}).Where("firebase_uuid = ?", uuid).Count(&count)
	return count != 0
}

func ValidateUsername(username string) (bool, string) {
	// Username should be between 3 and 20 characters
	if len(username) < 3 || len(username) > 20 {
		return false, "Username should be between 3 and 20 characters"
	}
	// Username should be alphanumeric with no spaces, only . and _
	for _, char := range username {
		if !(char >= 'a' && char <= 'z') &&
			!(char >= 'A' && char <= 'Z') &&
			!(char >= '0' && char <= '9') &&
			char != '.' &&
			char != '_' {
			return false, "Username should be alphanumeric with no spaces, only . and _"
		}
	}

	// Username should have at max 1 . or _ in a row
	for i := 0; i < len(username)-1; i++ {
		if (username[i] == '.' || username[i] == '_') && (username[i+1] == '.' || username[i+1] == '_') {
			return false, "Username should have at max 1 . or _ in a row"
		}
	}

	if CheckUsernameExists(username) {
		return false, "Username already exists"
	}

	return true, ""
}
