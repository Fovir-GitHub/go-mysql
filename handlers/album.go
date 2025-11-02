package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/Fovir-GitHub/go-mysql/utils"
	"github.com/gin-gonic/gin"
)

type AlbumHandler struct {
	DB *sql.DB
}

func (h *AlbumHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	album, err := utils.QueryAlbumByID(h.DB, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, album)
}

func (h *AlbumHandler) GetByArtist(c *gin.Context) {
	name := c.Param("artist")
	albums, err := utils.QueryAlbumByArtist(h.DB, name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, albums)
}
