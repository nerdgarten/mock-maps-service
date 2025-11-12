package main

import (
	"log"
	"net/http"
	"os"

	"github.com/nerdgarten/mock-maps/service/server"
)

func main() {
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
