// helpers/response.go
package helpers

import (
	"github.com/gin-gonic/gin"
)

func RespondJSON(c *gin.Context, status int, payload interface{}) {
	c.Header("Content-Type", "application/json")
	c.JSON(status, payload)
}
