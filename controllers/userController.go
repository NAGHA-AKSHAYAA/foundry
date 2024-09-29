package controllers

import (
	"net/http"

	"foundry/dto"
	"foundry/initialisers"
	"foundry/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user dto.UserDetails
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userModel := models.User{
		Role:     user.Role,
		Password: user.Password,
	}

	if err := userModel.HashPassword(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	userModel.Create()

	c.JSON(http.StatusOK, userModel)
}

func Login(c *gin.Context) {
	var user dto.UserDetails
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userModel models.User
	if err := initialisers.DB.Where("role = ?", user.Role).First(&userModel).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	if !userModel.Authenticate(user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": userModel})

}

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "Server is running"})
}
