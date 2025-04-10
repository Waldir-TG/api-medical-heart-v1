package models

// CREATE TABLE role_permissions (
//     role_id UUID REFERENCES roles(id) ON DELETE CASCADE,
//     permission_id UUID REFERENCES permissions(id) ON DELETE CASCADE,
//     PRIMARY KEY (role_id, permission_id)
// );

type RolePermission struct {
	RoleID       string `json:"role_id"`
	PermissionID string `json:"permission_id"`
}
