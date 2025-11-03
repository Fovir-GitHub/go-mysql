package main

import (
	"fmt"

	"github.com/Fovir-GitHub/go-mysql/handlers"
	"github.com/Fovir-GitHub/go-mysql/router"
	"github.com/Fovir-GitHub/go-mysql/utils"
)

func main() {
	const PORT = 8080

	db := utils.ConnectToDatabase("root", "tcp", "127.0.0.1:3306", "recordings")
	fmt.Println("Connected!")
	defer db.Close()

	albumHandler := &handlers.AlbumHandler{DB: db}
	r := router.SetupRouter(albumHandler)

	fmt.Printf("Server running at http://localhost:%v", PORT)
	r.Run(fmt.Sprintf(":%v", PORT))
}
