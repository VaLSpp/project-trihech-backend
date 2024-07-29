package controllers

import (
	"net/http"
	"project-trihech-backend/config"
	"project-trihech-backend/models"

	"github.com/gin-gonic/gin"
)

func GetProjects(c *gin.Context) {
    var projects []models.Project
    if err := config.DB.Preload("User").Find(&projects).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, projects)
}

func GetProject(c *gin.Context) {
    var project models.Project
    if err := config.DB.Preload("User").Preload("ProjectTechnologies").First(&project, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
        return
    }
    c.JSON(http.StatusOK, project)
}

func CreateProject(c *gin.Context) {
    var input models.Project
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

func UpdateProject(c *gin.Context) {
    var project models.Project
    if err := config.DB.First(&project, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
        return
    }

    var input models.Project
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB.Model(&project).Updates(input).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, project)
}

func DeleteProject(c *gin.Context) {
    var project models.Project
    if err := config.DB.First(&project, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
        return
    }

    if err := config.DB.Delete(&project).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Project deleted"})
}
