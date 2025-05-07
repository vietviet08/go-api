package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"vietquoc/connect-db/models"
	"vietquoc/connect-db/services"
)

func GetAlbums(c *gin.Context) {
	artistName := c.Query("name")

	if artistName != "" {
		albums, err := services.GetAlbumsByArtist(artistName)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to get albums by artist"})
			return
		}

		albumDTOs := services.ConvertToDTOs(albums)
		c.IndentedJSON(http.StatusOK, albumDTOs)
		return
	}

	albums, err := services.GetAllAlbums()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to get albums"})
		return
	}

	albumDTOs := services.ConvertToDTOs(albums)
	c.IndentedJSON(http.StatusOK, albumDTOs)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid album ID"})
		return
	}

	album, err := services.GetAlbumByID(idInt)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to get album"})
		return
	}

	albumDTO := services.ConvertToDTO(album)
	c.IndentedJSON(http.StatusOK, albumDTO)
}

func CreateAlbum(c *gin.Context) {
	var newAlbumDTO models.AlbumDTO

	if err := c.BindJSON(&newAlbumDTO); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid album data"})
		return
	}

	album := models.Album{
		Title:  newAlbumDTO.Title,
		Artist: newAlbumDTO.Artist,
		Price:  float32(newAlbumDTO.Price),
	}

	albumID, err := services.AddAlbum(album)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to add album"})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{
		"id":     albumID,
		"title":  newAlbumDTO.Title,
		"artist": newAlbumDTO.Artist,
		"price":  newAlbumDTO.Price,
	})
}

func UpdateAlbum(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid album ID"})
		return
	}

	album, err := services.GetAlbumByID(idInt)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to get album"})
		return
	}

	var newAlbumDTO models.AlbumDTO

	if err := c.BindJSON(&newAlbumDTO); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid album data"})
		return
	}

	album.Title = newAlbumDTO.Title
	album.Artist = newAlbumDTO.Artist
	album.Price = float32(newAlbumDTO.Price)

	albumID, err := services.UpdateAlbum(idInt, album)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to update album"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"id":     albumID,
		"title":  album.Title,
		"artist": album.Artist,
		"price":  album.Price,
	})

}

func DeleteAlbum(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid album ID"})
		return
	}

	albumID, err := services.DeleteAlbum(idInt)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete album"})
		return
	}

	c.IndentedJSON(http.StatusNoContent, gin.H{"message": "Album deleted successfully ", "id": albumID})

}
