package models

import (
    "time"
    "gorm.io/gorm"
)

type Project struct {
    gorm.Model
    NamaProject string
    Deskripsi   string
    CreatedAt   time.Time
    UpdatedAt   time.Time
    UserID      uint
    User        User `gorm:"foreignKey:UserID"`
    ProjectTechnologies []ProjectTechnologies `gorm:"foreignKey:ProjectID"`
}
