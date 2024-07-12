package album

import (
	"database/sql"
	"go-app/internal/database"
	"go-app/pkg/logging"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func GetAlbums() ([]Album, error) {
	logging.Log.Infof("album.go GetAlbums")
	rows, err := database.DB.Query("SELECT id, title, artist, price FROM albums")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var albums []Album
	for rows.Next() {
		var album Album
		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}
	return albums, nil
}

func GetAlbumByID(id string) (Album, error) {
	logging.Log.Infof("album.go GetAlbumByID")
	var album Album
	row := database.DB.QueryRow("SELECT id, title, artist, price FROM albums WHERE id = ?", id)
	err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
	if err == sql.ErrNoRows {
		return album, nil
	} else if err != nil {
		return album, err
	}
	return album, nil
}

func AddAlbum(album Album) error {
	_, err := database.DB.Exec("INSERT INTO albums (id, title, artist, price) VALUES (?, ?, ?, ?)", album.ID, album.Title, album.Artist, album.Price)
	if err != nil {
		return err
	}
	return nil
}
