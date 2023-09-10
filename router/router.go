// router/router.go
package router

import (
    "github.com/gin-gonic/gin"
    "your_project_name/controllers"
    "your_project_name/middlewares"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    // Endpoint untuk user
    user := r.Group("/users
