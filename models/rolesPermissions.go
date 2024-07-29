package models

type RolesPermissions struct {
	RoleID       uint       `gorm:"primaryKey" json:"role_id"`
	PermissionID uint       `gorm:"primaryKey" json:"permission_id"`
	Role         Role       `gorm:"foreignKey:RoleID"`
	Permission   Permission `gorm:"foreignKey:PermissionID"`
}
