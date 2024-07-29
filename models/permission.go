package models

import (
    "gorm.io/gorm"
)

type Permission struct {
    gorm.Model
    Nama             string `gorm:"unique;size:50"`
    RolesPermissions []RolesPermissions `gorm:"foreignKey:PermissionID"`
}
