package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Fovir-GitHub/go-mysql/models"
	"github.com/Fovir-GitHub/go-mysql/utils"
)

func main() {
	var db *sql.DB
	var err error

	db = utils.ConnectToDatabase("root", "tcp", "127.0.0.1:3306", "recordings")
	fmt.Println("Connected!")

	albums, err := utils.QueryAlbumByArtist(db, "John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	alb, err := utils.QueryAlbumByID(db, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", alb)

	albID, err := utils.AddAlbum(db, models.Album{Title: "The Modern Sound of Betty Carter", Artist: "Betty Carter", Price: 49.99})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albID)

	deletedAlbID, err := utils.DeleteAlbumByID(db, albID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID removed: %v\n", deletedAlbID)
}
