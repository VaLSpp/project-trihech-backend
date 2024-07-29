package controllers

import (
	"net/http"
	"project-trihech-backend/config"
	"project-trihech-backend/models"

	"github.com/gin-gonic/gin"
)

func GetProjectTechnologies(c *gin.Context) {
    var projectTechnologies []models.ProjectTechnologies
    if err := config.DB.Find(&projectTechnologies).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, projectTechnologies)
}

func GetProjectTechnology(c *gin.Context) {
    var projectTechnology models.ProjectTechnologies
    if err := config.DB.First(&projectTechnology, "project_id = ? AND technology_id = ?", c.Param("project_id"), c.Param("technology_id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "ProjectTechnology not found"})
        return
    }
    c.JSON(http.StatusOK, projectTechnology)
}

func CreateProjectTechnology(c *gin.Context) {
    var input models.ProjectTechnologies
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

func DeleteProjectTechnology(c *gin.Context) {
    var projectTechnology models.ProjectTechnologies
    if err := config.DB.First(&projectTechnology, "project_id = ? AND technology_id = ?", c.Param("project_id"), c.Param("technology_id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "ProjectTechnology not found"})
        return
    }

    if err := config.DB.Delete(&projectTechnology).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "ProjectTechnology deleted"})
}
