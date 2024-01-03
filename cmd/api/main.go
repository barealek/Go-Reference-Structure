package main

import (
	"log"
	"os"

	"go-reference-structure/internal/database"
	"go-reference-structure/internal/server"
	"go-reference-structure/internal/server/routes"
)

// The apps entrypoint

func main() {
	db, err := database.NewMongo()

	if err != nil {
		log.Fatal(err)
	}

	server := server.New(db)

	routes.Register(server)

	log.Fatal(server.Start(os.Getenv("HOST")))
}
