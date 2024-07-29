package controllers

import (
    "net/http"
    "project-trihech-backend/config"
    "project-trihech-backend/models"

    "github.com/gin-gonic/gin"
)

func GetRoles(c *gin.Context) {
    var roles []models.Role
    if err := config.DB.Find(&roles).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, roles)
}

func GetRole(c *gin.Context) {
    var role models.Role
    if err := config.DB.First(&role, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
        return
    }
    c.JSON(http.StatusOK, role)
}

func CreateRole(c *gin.Context) {
    var input models.Role
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB.Create(&input).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, input)
}

func UpdateRole(c *gin.Context) {
    var role models.Role
    if err := config.DB.First(&role, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
        return
    }

    var input models.Role
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB.Model(&role).Updates(input).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, role)
}

func DeleteRole(c *gin.Context) {
    var role models.Role
    if err := config.DB.First(&role, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
        return
    }

    if err := config.DB.Delete(&role).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Role deleted"})
}
