package main

import (
	"api/config"
	"api/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := routes.SetupRouter()
	fmt.Println("🚀 Server running on port 8080")
	r.Run(":8080")
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
}
