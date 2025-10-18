package main

import (
	"log"
	"net"
	"os"

	pb "github.com/nerdgarten/mock-maps/service/proto"
	"github.com/nerdgarten/mock-maps/service/server"
	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "50051"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	mapsServer := server.NewMapsServer()
	pb.RegisterMapsServiceServer(s, mapsServer)

	log.Printf("Mock Maps gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
