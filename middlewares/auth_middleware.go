// middlewares/auth_middleware.go
package middlewares

import (
	"net/http"
	"your_project_name/helpers"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implementasi otentikasi JWT
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			helpers.RespondJSON(c, http.StatusUnauthorized, gin.H{"error": "Token is missing"})
			c.Abort()
			return
		}
		// Validasi token
		token, err := helpers.ValidateToken(tokenString)
		if err != nil {
			helpers.RespondJSON(c, http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		if token.Valid {
			c.Next()
		} else {
			helpers.RespondJSON(c, http.StatusUnauthorized, gin.H{"error": "Token is not valid"})
			c.Abort()
			return
		}
	}
}
