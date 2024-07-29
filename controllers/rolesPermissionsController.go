package controllers

import (
	"fmt"
	"net/http"
	"project-trihech-backend/config"
	"project-trihech-backend/models"

	"github.com/gin-gonic/gin"
)

func GetRolesPermissions(c *gin.Context) {
    var rolesPermissions []models.RolesPermissions
    if err := config.DB.Find(&rolesPermissions).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, rolesPermissions)
}

func GetRolePermission(c *gin.Context) {
    var rolesPermissions models.RolesPermissions
    if err := config.DB.First(&rolesPermissions, "role_id = ? AND permission_id = ?", c.Param("role_id"), c.Param("permission_id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "RolePermission not found"})
        return
    }
    c.JSON(http.StatusOK, rolesPermissions)
}

func CreateRolePermission(c *gin.Context) {
    var input models.RolesPermissions
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    
    fmt.Println(input, "test")

    if err := config.DB.Create(&input).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, input)
}

func DeleteRolePermission(c *gin.Context) {
    var rolesPermissions models.RolesPermissions
    if err := config.DB.First(&rolesPermissions, "role_id = ? AND permission_id = ?", c.Param("role_id"), c.Param("permission_id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "RolePermission not found"})
        return
    }

    if err := config.DB.Delete(&rolesPermissions).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "RolePermission deleted"})
}
