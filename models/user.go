package models

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Username     string `gorm:"unique;not null"`
    PasswordHash string `gorm:"not null"`
    Email        string `gorm:"unique;not null"`
    RoleID       uint
    Role         Role `gorm:"foreignKey:RoleID;constraint:onUpdate:CASCADE,onDelete:SET NULL; default:2"`
    UserDetails  *UserDetails `gorm:"foreignKey:UserID"`
    Projects     []Project   `gorm:"foreignKey:UserID"`
}
