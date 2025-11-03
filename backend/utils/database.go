package utils

import (
	"database/sql"
	"log"

	"fmt"

	"github.com/Fovir-GitHub/go-mysql/models"
	"github.com/go-sql-driver/mysql"
)

const DATABASE_NAME = "album"

func ConnectToDatabase(user string, net string, addr string, dbName string) *sql.DB {
	cfg := mysql.NewConfig()
	cfg.User = user
	cfg.Net = net
	cfg.Addr = addr
	cfg.DBName = dbName
	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	return db
}

func AddAlbum(db *sql.DB, alb models.Album) (int64, error) {
	result, err := db.Exec(fmt.Sprintf("INSERT INTO %v (title, artist, price) VALUES (?,?,?)", DATABASE_NAME), alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("AddAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("AddAlbum: %v", err)
	}
	return id, err
}

func DeleteAlbumByID(db *sql.DB, id int64) (int64, error) {
	result, err := db.Exec(fmt.Sprintf("DELETE FROM %v WHERE id = ?", DATABASE_NAME), id)
	if err != nil {
		return 0, fmt.Errorf("DeleteAlbumByID %v: %v", id, err)
	}
	rowsAffected, err := result.RowsAffected()
	if rowsAffected < 1 || err != nil {
		return 0, fmt.Errorf("DeleteAlbumByID %v: ID does not exist", id)
	}
	return id, nil
}

func QueryAllAlbums(db *sql.DB) ([]models.Album, error) {
	var albums []models.Album
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %v", DATABASE_NAME))
	if err != nil {
		return nil, fmt.Errorf("QueryAllAlbums: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var alb models.Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("QueryAllAlbums: %v", err)
		}
		albums = append(albums, alb)
	}

	return albums, nil
}

func QueryAlbumByArtist(db *sql.DB, name string) ([]models.Album, error) {
	var albums []models.Album
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %v WHERE artist = ?", DATABASE_NAME), name)
	if err != nil {
		return nil, fmt.Errorf("QueryAlbumByArtist: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var alb models.Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("QueryAlbumByArtist %p: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("QueryAlbumByArtist %p: %v", name, err)
	}

	return albums, nil
}

func QueryAlbumByID(db *sql.DB, id int64) (models.Album, error) {
	var alb models.Album
	row := db.QueryRow(fmt.Sprintf("SELECT * FROM %v WHERE id = ?", DATABASE_NAME), id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("QueryAlbumByID %d: no such album", id)
		}
		return alb, fmt.Errorf("albumByID %d: %v", id, err)
	}

	return alb, nil
}
