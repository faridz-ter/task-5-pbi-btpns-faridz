// controllers/photo_controller.go
package controllers

import (
	"net/http"
	"your_project_name/database"
	"your_project_name/helpers"
	"your_project_name/models"

	"github.com/gin-gonic/gin"
)

func CreatePhoto(c *gin.Context) {
	// Implementasi pembuatan foto
	var photo models.Photo
	if err := c.ShouldBindJSON(&photo); err != nil {
		helpers.RespondJSON(c, http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Pastikan Anda menetapkan UserID dengan ID user yang saat ini masuk
	userID := c.GetInt("user_id")
	photo.UserID = uint(userID)
	database.DB.Create(&photo)
	helpers.RespondJSON(c, http.StatusCreated, photo)
}
