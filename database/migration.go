package database

import (
    "project-trihech-backend/config"
    "project-trihech-backend/models"
)

func Migrate() {
    config.DB.AutoMigrate(
        &models.User{},
        &models. UserDetails{},
        &models.Role{},
        &models.Permission{},
        &models.RolesPermissions{},
        &models.Project{},
        &models.Technology{},
        &models.ProjectTechnologies{},
    )
}
