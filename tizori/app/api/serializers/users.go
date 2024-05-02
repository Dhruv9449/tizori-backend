package serializers

import "github.com/GDGVIT/Tizori-backend/internal/models"

func UserLoginSerializer(user models.User, token string) map[string]interface{} {
	return map[string]interface{}{
		"username": user.Username,
		"name":     user.Name,
		"token":    token,
	}
}

func UserSerializer(user models.User) map[string]interface{} {
	return map[string]interface{}{
		"username": user.Username,
		"name":     user.Name,
		"email":    user.Email,
		"roles":    RolesListSerializer(user.Roles),
	}
}

func UserBlockSerializer(user models.User) map[string]interface{} {
	return map[string]interface{}{
		"username": user.Username,
		"name":     user.Name,
		"email":    user.Email,
	}
}

func UsersListSerializer(users []models.User) []map[string]interface{} {
	var serializedUsers []map[string]interface{}
	for _, user := range users {
		serializedUsers = append(serializedUsers, UserBlockSerializer(user))
	}
	return serializedUsers
}
