package album

import (
	"go-app/pkg/logging"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)
}

func getAlbums(c *gin.Context) {
	logging.Log.Infof("Fetching albums handler")
	albums, err := GetAlbums()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to get albums"})
		return
	}
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	if err := AddAlbum(newAlbum); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to add album"})
		return
	}
	logging.Log.Infof("Added new album: %+v", newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	logging.Log.Infof("Fetching album with ID: %s", id)

	album, err := GetAlbumByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to get album"})
		return
	}
	if album.ID == "" {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}
