package main

import (
	"log"
	"os"

	"messenger-module/confs"
	"messenger-module/db"
	"messenger-module/server"
)

func main() {
	if err := confs.LoadConfig(); err != nil {
		log.Fatalf("failed to load .env: %v", err)
	}
	database, err := db.Connect()
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}
	port := os.Getenv("PORT")
	if port == "" { port = "8080" }
	srv := server.NewServer()
	if err := srv.Start(database, port); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
