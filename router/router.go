package router

import (
	"github.com/Fovir-GitHub/go-mysql/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(albumHandler *handlers.AlbumHandler) *gin.Engine {
	r := gin.Default()
	r.GET("/albums/id/:id", albumHandler.GetByID)
	r.GET("/albums/artist/:artist", albumHandler.GetByArtist)
	r.POST("/albums/add", albumHandler.PostAddAlbum)
	r.POST("/albums/delete/:id", albumHandler.PostDeleteAlbumByID)

	return r
}
