package main

import (
	"context"
	"log"
	"os"

	pb "github.com/nerdgarten/mock-maps/service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Connect to the gRPC server
	port := os.Getenv("PORT")
	if port == "" {
		port = "50051"
	}
	conn, err := grpc.Dial("localhost:"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewMapsServiceClient(conn)

	// Example 1: Get Directions
	directionsReq := &pb.GetDirectionsRequest{
		Origin:      &pb.Location{Lat: 13.7563, Lng: 100.5018},
		Destination: &pb.Location{Lat: 13.7383, Lng: 100.5323},
		Mode:        "driving",
	}
	directionsResp, err := client.GetDirections(context.Background(), directionsReq)
	if err != nil {
		log.Printf("GetDirections failed: %v", err)
	} else {
		log.Printf("Directions: Distance: %s, Duration: %s", directionsResp.Distance, directionsResp.Duration)
	}

	// Example 2: Search Places
	placesReq := &pb.SearchPlacesRequest{
		Query:    "pizza",
		Location: &pb.Location{Lat: 13.7563, Lng: 100.5018},
		Radius:   3000,
		Type:     "restaurant",
	}
	placesResp, err := client.SearchPlaces(context.Background(), placesReq)
	if err != nil {
		log.Printf("SearchPlaces failed: %v", err)
	} else {
		log.Printf("Found %d places", len(placesResp.Results))
		for _, place := range placesResp.Results {
			log.Printf("Place: %s at %s", place.Name, place.Address)
		}
	}

	// Example 3: Track (streaming)
	trackReq := &pb.TrackRequest{}
	stream, err := client.Track(context.Background(), trackReq)
	if err != nil {
		log.Printf("Track failed: %v", err)
	} else {
		log.Println("Starting to receive location updates...")
		for {
			location, err := stream.Recv()
			if err != nil {
				log.Printf("Stream ended: %v", err)
				break
			}
			log.Printf("Location: Lat: %.4f, Lng: %.4f at %s", location.Location.Lat, location.Location.Lng, location.Timestamp)
		}
	}

	// Example 4: Get Static Map
	staticMapReq := &pb.StaticMapRequest{
		Center: &pb.Location{Lat: 13.7563, Lng: 100.5018},
		Zoom:   14,
		Size:   "600x400",
	}
	staticMapResp, err := client.GetStaticMap(context.Background(), staticMapReq)
	if err != nil {
		log.Printf("GetStaticMap failed: %v", err)
	} else {
		log.Printf("Static Map URL: %s", staticMapResp.Url)
	}
}
