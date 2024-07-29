package models

import (
    "gorm.io/gorm"
)

type Role struct {
    gorm.Model
    Nama             string `gorm:"unique;size:50"`
    Users            []User `gorm:"foreignKey:RoleID"`
    RolesPermissions []RolesPermissions `gorm:"foreignKey:RoleID"`
}
