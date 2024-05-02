package models

import (
	"log"

	"github.com/GDGVIT/Tizori-backend/internal/database"
)

func InitializeModels() {
	MODELS := map[string]interface{}{
		"User":         &User{},
		"Role":         &Role{},
		"Applications": &Application{},
	}

	for name, model := range MODELS {
		err := database.DB.AutoMigrate(model)
		if err != nil {
			log.Fatal("Failed to initialize model: ", name)
		} else {
			log.Println("Successfully initialized model: ", name)
		}
	}
}
