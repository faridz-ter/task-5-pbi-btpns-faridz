// controllers/user_controller.go
package controllers

import (
	"net/http"
	"your_project_name/database"
	"your_project_name/helpers"
	"your_project_name/models"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	// Implementasi registrasi user
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		helpers.RespondJSON(c, http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		helpers.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = hashedPassword
	database.DB.Create(&user)
	token, err := helpers.CreateToken(user.ID)
	if err != nil {
		helpers.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
		return
	}
	helpers.RespondJSON(c, http.StatusOK, gin.H{"token": token})
}
