package controllers

import (
	"net/http"
	"project-trihech-backend/config"
	"project-trihech-backend/models"

	"github.com/gin-gonic/gin"
)

func GetTechnologies(c *gin.Context) {
    var technologies []models.Technology
    if err := config.DB.Find(&technologies).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, technologies)
}

func GetTechnology(c *gin.Context) {
    var technology models.Technology
    if err := config.DB.First(&technology, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Technology not found"})
        return
    }
    c.JSON(http.StatusOK, technology)
}

func CreateTechnology(c *gin.Context) {
    var input models.Technology
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

func UpdateTechnology(c *gin.Context) {
    var technology models.Technology
    if err := config.DB.First(&technology, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Technology not found"})
        return
    }

    var input models.Technology
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB.Model(&technology).Updates(input).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, technology)
}

func DeleteTechnology(c *gin.Context) {
    var technology models.Technology
    if err := config.DB.First(&technology, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Technology not found"})
        return
    }

    if err := config.DB.Delete(&technology).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Technology deleted"})
}
