package models

import (
    "gorm.io/gorm"
)

type Technology struct {
    gorm.Model
    Name string `gorm:"unique;size:50"`
    ProjectTechnologies []ProjectTechnologies `gorm:"foreignKey:TechnologyID"`
}
