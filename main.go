package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/nerdgarten/mock-maps/service/server"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file: %s", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "50051"
	}

	mux := http.NewServeMux()
	server.NewMapsServer().RegisterRoutes(mux)

	addr := ":" + port
	log.Printf("Mock Maps REST server listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
