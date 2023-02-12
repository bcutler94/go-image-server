package routes

import (
	"example/web-service-gin/controllers"

	"github.com/gin-gonic/gin"
)

func AlbumRoute(router *gin.Engine) {
	// TODO
	group := router.Group("/v1")
	group.POST("/album", controllers.CreateAlbum())
	group.GET("/album/:id", controllers.GetAlbum())
}
