package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
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

	r := gin.Default()
	server.NewMapsServer().RegisterRoutes(r)

	addr := ":" + port
	log.Printf("Mock Maps REST server listening on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
