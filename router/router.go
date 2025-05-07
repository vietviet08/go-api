package router

import (
	"github.com/gin-gonic/gin"

	"vietquoc/connect-db/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/albums", controllers.GetAlbums)
	router.GET("/albums/:id", controllers.GetAlbumByID)
	router.POST("/albums", controllers.CreateAlbum)

	return router
}
