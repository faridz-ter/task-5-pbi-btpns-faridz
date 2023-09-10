// main.go
package main

import (
	"your_project_name/database"
	"your_project_name/router"
)

func main() {
	database.Init()
	r := router.SetupRouter()
	r.Run(":8080")
}
