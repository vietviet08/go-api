package repositories

import (
	"database/sql"
	"fmt"

	"vietquoc/connect-db/database"
	"vietquoc/connect-db/models"
)

func GetAllAlbums() ([]models.Album, error) {
	var albums []models.Album

	rows, err := database.DB.Query("SELECT * FROM album")
	if err != nil {
		return nil, fmt.Errorf("GetAllAlbums: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var alb models.Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("GetAllAlbums: %v", err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetAllAlbums: %v", err)
	}
	return albums, nil
}

func GetAlbumsByArtist(name string) ([]models.Album, error) {
	var albums []models.Album

	rows, err := database.DB.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("GetAlbumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	for rows.Next() {
		var alb models.Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("GetAlbumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetAlbumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

func GetAlbumByID(id int64) (models.Album, error) {
	var alb models.Album

	row := database.DB.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("GetAlbumByID %d: no such album", id)
		}
		return alb, fmt.Errorf("GetAlbumByID %d: %v", id, err)
	}
	return alb, nil
}

func AddAlbum(alb models.Album) (int64, error) {
	result, err := database.DB.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("AddAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("AddAlbum: %v", err)
	}
	return id, nil
}

func UpdateAlbum(id int64, alb models.Album) (int64, error) {
	_, err := database.DB.Exec("UPDATE album SET title = ?, artist = ?, price = ? WHERE id = ?", alb.Title, alb.Artist, alb.Price, id)
	if err != nil {
		return 0, fmt.Errorf("UpdateAlbum %d: %v", id, err)
	}
	return id, nil
}
