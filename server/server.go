package server

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/nerdgarten/mock-maps/service/data"
	pb "github.com/nerdgarten/mock-maps/service/proto"
)

// MapsServer implements the MapsServiceServer interface
type MapsServer struct {
	pb.UnimplementedMapsServiceServer
}

// NewMapsServer creates a new MapsServer instance
func NewMapsServer() *MapsServer {
	return &MapsServer{}
}

// GetDirections returns mock directions
func (s *MapsServer) GetDirections(ctx context.Context, req *pb.GetDirectionsRequest) (*pb.Route, error) {
	log.Printf("GetDirections called with origin: %v, destination: %v, mode: %s", req.Origin, req.Destination, req.Mode)
	route := data.GetMockRoute(req.Mode)
	if route == nil {
		return nil, fmt.Errorf("unsupported mode: %s", req.Mode)
	}
	return route, nil
}

// SearchPlaces returns mock places based on search criteria
func (s *MapsServer) SearchPlaces(ctx context.Context, req *pb.SearchPlacesRequest) (*pb.SearchPlacesResponse, error) {
	log.Printf("SearchPlaces called with query: %s, location: %v, radius: %d, type: %s", req.Query, req.Location, req.Radius, req.Type)
	places := data.GetMockPlaces(req.Query, req.Location, req.Radius, req.Type)
	return &pb.SearchPlacesResponse{Results: places}, nil
}

// Geocode converts address to coordinates
func (s *MapsServer) Geocode(ctx context.Context, req *pb.GeocodeRequest) (*pb.GeocodeResponse, error) {
	log.Printf("Geocode called with address: %s", req.Address)
	result := data.GetMockGeocode(req.Address)
	if result == nil {
		return &pb.GeocodeResponse{Results: []*pb.GeocodeResult{}}, nil // Return empty for unknown addresses
	}
	return &pb.GeocodeResponse{Results: []*pb.GeocodeResult{result}}, nil
}

// ReverseGeocode converts coordinates to address
func (s *MapsServer) ReverseGeocode(ctx context.Context, req *pb.ReverseGeocodeRequest) (*pb.ReverseGeocodeResponse, error) {
	log.Printf("ReverseGeocode called with location: %v", req.Location)
	address := data.GetMockReverseGeocode(req.Location.Lat, req.Location.Lng)
	if address == "" {
		address = fmt.Sprintf("%.4f, %.4f", req.Location.Lat, req.Location.Lng) // Fallback
	}
	return &pb.ReverseGeocodeResponse{Address: address}, nil
}

// DistanceMatrix calculates distances between origins and destinations
func (s *MapsServer) DistanceMatrix(ctx context.Context, req *pb.DistanceMatrixRequest) (*pb.DistanceMatrixResponse, error) {
	log.Printf("DistanceMatrix called with %d origins and %d destinations", len(req.Origins), len(req.Destinations))
	// For simplicity, return the mock data regardless of input
	return &pb.DistanceMatrixResponse{Rows: data.MockDistanceMatrix}, nil
}

// OptimizeRoute optimizes the order of stops
func (s *MapsServer) OptimizeRoute(ctx context.Context, req *pb.OptimizeRouteRequest) (*pb.OptimizedRoute, error) {
	log.Printf("OptimizeRoute called with %d stops", len(req.Stops))
	// Randomly select one of the mock optimized routes
	index := rand.Intn(len(data.MockOptimizedRoutes))
	return data.MockOptimizedRoutes[index], nil
}

// Track streams mock location updates
func (s *MapsServer) Track(req *pb.TrackRequest, stream pb.MapsService_TrackServer) error {
	log.Println("Track called, starting location stream")
	for _, loc := range data.MockTrackLocations {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		default:
			if err := stream.Send(&pb.TrackLocation{
				Location:  &pb.Location{Lat: loc.Lat, Lng: loc.Lng},
				Timestamp: loc.Timestamp,
			}); err != nil {
				return err
			}
			time.Sleep(1 * time.Second)
		}
	}
	return nil
}

// GetStaticMap returns a mock static map URL
func (s *MapsServer) GetStaticMap(ctx context.Context, req *pb.StaticMapRequest) (*pb.StaticMap, error) {
	log.Printf("GetStaticMap called with center: %v, zoom: %d, size: %s", req.Center, req.Zoom, req.Size)
	key := fmt.Sprintf("%.4f,%.4f_%d_%s", req.Center.Lat, req.Center.Lng, req.Zoom, req.Size)
	url, ok := data.MockStaticMaps[key]
	if !ok {
		url = "https://mockmaps.local/static/default.png" // Default fallback
	}
	return &pb.StaticMap{Url: url}, nil
}
