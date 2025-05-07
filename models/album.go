package models

// Album represents album data stored in the database
type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

// AlbumDTO is the data transfer object for Album
type AlbumDTO struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
