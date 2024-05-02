package serializers

import "github.com/GDGVIT/Tizori-backend/internal/models"

func ApplicationSerializer(application models.Application) map[string]interface{} {
	return map[string]interface{}{
		"id":   application.Id,
		"name": application.Name,
	}
}

func ApplicationsListSerializer(applications []models.Application) []map[string]interface{} {
	var serializedApplications []map[string]interface{}
	for _, application := range applications {
		serializedApplications = append(serializedApplications, ApplicationSerializer(application))
	}
	return serializedApplications
}

func ApplicationCredentialsSerializer(applications models.Application, credentials models.ApplicationCredentials) map[string]interface{} {
	return map[string]interface{}{
		"id":   applications.Id,
		"name": applications.Name,
		"credentials": map[string]string{
			"username": credentials.Username,
			"password": credentials.Password,
		},
	}
}
