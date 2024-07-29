package controllers

import (
    "net/http"
    "project-trihech-backend/config"
    "project-trihech-backend/models"

    "github.com/gin-gonic/gin"
)

func GetUserDetails(c *gin.Context) {
    var userDetails models.UserDetails
    if err := config.DB.First(&userDetails, "user_id = ?", c.Param("user_id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "UserDetails not found"})
        return
    }
    c.JSON(http.StatusOK, userDetails)
}

func CreateUserDetails(c *gin.Context) {
    var input models.UserDetails
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Get userID from content
    userID, exists := c.Get("UserID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "UserID Not Found"})
        return
    }
    input.UserID = userID.(uint)

    if err := config.DB.Create(&input).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, input)
}

func UpdateUserDetails(c *gin.Context) {
    var userDetails models.UserDetails
    if err := config.DB.First(&userDetails, "user_id = ?", c.Param("user_id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "UserDetails not found"})
        return
    }

    var input models.UserDetails
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB.Model(&userDetails).Updates(input).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, userDetails)
}

func DeleteUserDetails(c *gin.Context) {
    var userDetails models.UserDetails
    if err := config.DB.First(&userDetails, "user_id = ?", c.Param("user_id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "UserDetails not found"})
        return
    }

    if err := config.DB.Delete(&userDetails).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "UserDetails deleted"})
}
