package main

import (
	"example/web-service-gin/configs"
	"example/web-service-gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Connect to database
	configs.ConnectDB()

	// routes
	routes.AlbumRoute(router)

	router.Run("localhost:8080")
}
