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
	router.PUT("/albums/:id", controllers.UpdateAlbum)
	router.DELETE("/albums/:id", controllers.DeleteAlbum)

	return router
}
