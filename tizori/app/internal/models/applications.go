package models

import (
	tizoriCrypto "github.com/GDGVIT/Tizori-backend/internal/crypto"
	"github.com/GDGVIT/Tizori-backend/internal/database"
)

type Application struct {
	Id          string                 `gorm:"primaryKey"`
	Name        string                 `gorm:"not null"`
	Credentials ApplicationCredentials `gorm:"serializer:json"`
}

type ApplicationCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a *Application) Save() error {
	return database.DB.Save(&a).Error
}

func (a *Application) Delete() error {
	return database.DB.Delete(&a).Error
}

func GetApplicationById(id string) (*Application, error) {
	var application Application
	err := database.DB.Where("id = ?", id).First(&application).Error
	return &application, err
}

func GetApplications() ([]Application, error) {
	var applications []Application
	err := database.DB.Find(&applications).Error
	return applications, err
}

func StoreApplicationCredentials(id string, credentials ApplicationCredentials) error {
	application, err := GetApplicationById(id)
	if err != nil {
		return err
	}

	application.Credentials.Username = credentials.Username

	// Encrypt password
	application.Credentials.Password, err = tizoriCrypto.EncryptPassword(credentials.Password, tizoriCrypto.AESKey)
	if err != nil {
		return err
	}

	return application.Save()
}

func GetApplicationCredentials(id string) (ApplicationCredentials, error) {
	application, err := GetApplicationById(id)
	if err != nil {
		return ApplicationCredentials{}, err
	}

	// Decrypt password
	password, err := tizoriCrypto.DecryptPassword(application.Credentials.Password, tizoriCrypto.AESKey)
	if err != nil {
		return ApplicationCredentials{}, err
	}

	return ApplicationCredentials{
		Username: application.Credentials.Username,
		Password: password,
	}, nil
}
