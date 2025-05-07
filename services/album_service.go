package services

import (
	"strconv"

	"vietquoc/connect-db/models"
	"vietquoc/connect-db/repositories"
)

func GetAllAlbums() ([]models.Album, error) {
	return repositories.GetAllAlbums()
}

func GetAlbumsByArtist(artistName string) ([]models.Album, error) {
	return repositories.GetAlbumsByArtist(artistName)
}

func GetAlbumByID(id int64) (models.Album, error) {
	return repositories.GetAlbumByID(id)
}

func AddAlbum(album models.Album) (int64, error) {
	return repositories.AddAlbum(album)
}

func UpdateAlbum(id int64, album models.Album) (int64, error) {
	return repositories.UpdateAlbum(id, album)
}

func DeleteAlbum(id int64) (int64, error) {
	return repositories.DeleteAlbum(id)
}

func ConvertToDTO(album models.Album) models.AlbumDTO {
	return models.AlbumDTO{
		ID:     strconv.FormatInt(album.ID, 10),
		Title:  album.Title,
		Artist: album.Artist,
		Price:  float64(album.Price),
	}
}

func ConvertToDTOs(albums []models.Album) []models.AlbumDTO {
	dtos := make([]models.AlbumDTO, len(albums))
	for i, album := range albums {
		dtos[i] = ConvertToDTO(album)
	}
	return dtos
}
