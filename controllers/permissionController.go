package controllers

import (
    "net/http"
    "project-trihech-backend/config"
    "project-trihech-backend/models"

    "github.com/gin-gonic/gin"
)

func GetPermissions(c *gin.Context) {
    var permissions []models.Permission
    if err := config.DB.Find(&permissions).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, permissions)
}

func GetPermission(c *gin.Context) {
    var permission models.Permission
    if err := config.DB.First(&permission, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Permission not found"})
        return
    }
    c.JSON(http.StatusOK, permission)
}

func CreatePermission(c *gin.Context) {
    var input models.Permission
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

func UpdatePermission(c *gin.Context) {
    var permission models.Permission
    if err := config.DB.First(&permission, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Permission not found"})
        return
    }

    var input models.Permission
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB.Model(&permission).Updates(input).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, permission)
}

func DeletePermission(c *gin.Context) {
    var permission models.Permission
    if err := config.DB.First(&permission, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Permission not found"})
        return
    }

    if err := config.DB.Delete(&permission).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Permission deleted"})
}
