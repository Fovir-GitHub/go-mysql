package router

import (
	"github.com/Fovir-GitHub/go-mysql/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(albumHandler *handlers.AlbumHandler) *gin.Engine {
	r := gin.Default()
	r.GET("/albums/:id", albumHandler.GetByID)

	return r
}
