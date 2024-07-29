package config

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "project-trihech-backend/models"
)

var DB *gorm.DB

func ConnectDatabase() {
    dsn := "root:@tcp(127.0.0.1:3306)/trihech-backend?charset=utf8mb4&parseTime=True&loc=Local"
    database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database!")
    }

    database.AutoMigrate(&models.User{}, &models.UserDetails{}, &models.Role{}, &models.Permission{}, &models.RolesPermissions{}, &models.Project{}, &models.Technology{}, &models.ProjectTechnologies{})
    DB = database
}
