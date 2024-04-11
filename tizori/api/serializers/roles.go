package serializers

import "github.com/GDGVIT/Tizori-backend/internal/models"

func PermissionsSerlializer(permission []models.Permission) []map[string]interface{} {
	var serializedPermissions []map[string]interface{}
	for _, p := range permission {
		serializedPermissions = append(serializedPermissions, map[string]interface{}{
			"scope":      p.Scope,
			"permission": p.Permission,
		})
	}
	return serializedPermissions
}

func RoleSerializer(role models.Role) map[string]interface{} {
	return map[string]interface{}{
		"id":          role.Id,
		"name":        role.Name,
		"permissions": PermissionsSerlializer(role.Permissions),
	}
}

func RolesListSerializer(roles []models.Role) []map[string]interface{} {
	var serializedRoles []map[string]interface{}
	for _, role := range roles {
		serializedRoles = append(serializedRoles, RoleSerializer(role))
	}
	return serializedRoles
}
