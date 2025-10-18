package data

import (
	"fmt"
	"math"
	"strings"

	pb "github.com/nerdgarten/mock-maps/service/proto"
)

// MockPlaces contains a list of mock places for search
var MockPlaces = []*pb.Place{
	{Name: "Pizza Planet", Address: "123 Main St", Location: &pb.Location{Lat: 13.7563, Lng: 100.5018}, Rating: 4.5},
	{Name: "Cheesy Bites", Address: "45 King Rd", Location: &pb.Location{Lat: 13.7570, Lng: 100.5030}, Rating: 4.2},
	{Name: "Burger Barn", Address: "78 Elm Ave", Location: &pb.Location{Lat: 13.7580, Lng: 100.5040}, Rating: 4.0},
	{Name: "Taco Town", Address: "90 Oak St", Location: &pb.Location{Lat: 13.7590, Lng: 100.5050}, Rating: 4.3},
	{Name: "Sushi Spot", Address: "12 Pine Rd", Location: &pb.Location{Lat: 13.7600, Lng: 100.5060}, Rating: 4.7},
	{Name: "Coffee Corner", Address: "34 Maple Ln", Location: &pb.Location{Lat: 13.7610, Lng: 100.5070}, Rating: 4.1},
	{Name: "Ice Cream Island", Address: "56 Birch Blvd", Location: &pb.Location{Lat: 13.7620, Lng: 100.5080}, Rating: 4.6},
	{Name: "Pasta Palace", Address: "78 Cedar Ct", Location: &pb.Location{Lat: 13.7630, Lng: 100.5090}, Rating: 4.4},
}

// MockGeocodes contains mock geocoding results
var MockGeocodes = map[string]*pb.GeocodeResult{
	"Chulalongkorn University": {FormattedAddress: "254 Phayathai Rd, Pathum Wan, Bangkok", Location: &pb.Location{Lat: 13.7383, Lng: 100.5323}},
	"Central World":            {FormattedAddress: "999/9 Rama I Rd, Pathum Wan, Bangkok", Location: &pb.Location{Lat: 13.7465, Lng: 100.5393}},
	"Grand Palace":             {FormattedAddress: "Na Phra Lan Rd, Phra Borom Maha Ratchawang, Phra Nakhon, Bangkok", Location: &pb.Location{Lat: 13.7500, Lng: 100.4925}},
	"Siam Paragon":             {FormattedAddress: "991/1 Rama I Rd, Pathum Wan, Bangkok", Location: &pb.Location{Lat: 13.7462, Lng: 100.5347}},
}

// MockAddresses contains mock reverse geocoding results
var MockAddresses = map[string]string{
	"13.7563,100.5018": "123 Main St, Bangkok, Thailand",
	"13.7383,100.5323": "254 Phayathai Rd, Pathum Wan, Bangkok",
	"13.7465,100.5393": "999/9 Rama I Rd, Pathum Wan, Bangkok",
	"13.7500,100.4925": "Na Phra Lan Rd, Phra Borom Maha Ratchawang, Phra Nakhon, Bangkok",
}

// MockRoutes contains mock routes for different modes
var MockRoutes = map[string]*pb.Route{
	"driving": {
		Distance: "5.3 km",
		Duration: "12 mins",
		Polyline: "mock_encoded_polyline_driving",
		Steps: []*pb.Step{
			{Instruction: "Head north on Main St", Distance: "200m"},
			{Instruction: "Turn right onto 1st Ave", Distance: "500m"},
			{Instruction: "Continue straight on 2nd St", Distance: "3.6 km"},
		},
	},
	"walking": {
		Distance: "5.3 km",
		Duration: "65 mins",
		Polyline: "mock_encoded_polyline_walking",
		Steps: []*pb.Step{
			{Instruction: "Walk north on Main St", Distance: "200m"},
			{Instruction: "Turn right onto 1st Ave", Distance: "500m"},
			{Instruction: "Continue on sidewalk along 2nd St", Distance: "3.6 km"},
		},
	},
	"bicycling": {
		Distance: "5.5 km",
		Duration: "18 mins",
		Polyline: "mock_encoded_polyline_bicycling",
		Steps: []*pb.Step{
			{Instruction: "Bike north on Main St", Distance: "200m"},
			{Instruction: "Turn right onto 1st Ave", Distance: "500m"},
			{Instruction: "Follow bike path on 2nd St", Distance: "3.8 km"},
		},
	},
}

// MockDistanceMatrix contains mock distance matrix data
var MockDistanceMatrix = []*pb.DistanceRow{
	{
		Elements: []*pb.DistanceElement{
			{Distance: "1.2 km", Duration: "4 mins"},
			{Distance: "3.8 km", Duration: "10 mins"},
			{Distance: "2.5 km", Duration: "7 mins"},
		},
	},
}

// MockOptimizedRoutes contains mock optimized routes
var MockOptimizedRoutes = []*pb.OptimizedRoute{
	{
		OptimizedOrder: []int32{0, 2, 1},
		TotalDistance:  "6.4 km",
		TotalDuration:  "15 mins",
	},
	{
		OptimizedOrder: []int32{1, 0, 2},
		TotalDistance:  "6.8 km",
		TotalDuration:  "16 mins",
	},
}

// MockTrackLocations contains mock tracking locations
var MockTrackLocations = []struct {
	Lat, Lng  float64
	Timestamp string
}{
	{13.7563, 100.5018, "2025-10-18T14:30:00Z"},
	{13.7570, 100.5025, "2025-10-18T14:30:10Z"},
	{13.7575, 100.5030, "2025-10-18T14:30:20Z"},
	{13.7580, 100.5035, "2025-10-18T14:30:30Z"},
	{13.7585, 100.5040, "2025-10-18T14:30:40Z"},
}

// MockStaticMaps contains mock static map URLs
var MockStaticMaps = map[string]string{
	"13.7563,100.5018_14_600x400": "https://mockmaps.local/static/preview1.png",
	"13.7383,100.5323_15_800x600": "https://mockmaps.local/static/preview2.png",
}

// GetMockPlaces returns filtered places based on query
func GetMockPlaces(query string, location *pb.Location, radius int32, placeType string) []*pb.Place {
	// Simple filtering logic
	var results []*pb.Place
	for _, place := range MockPlaces {
		if query != "" && !containsIgnoreCase(place.Name, query) {
			continue
		}
		if placeType != "" && !containsIgnoreCase(place.Name, placeType) {
			continue
		}
		// Simple distance check (mock)
		if location != nil && calculateDistance(location, place.Location) > float64(radius) {
			continue
		}
		results = append(results, place)
	}
	return results
}

// GetMockGeocode returns geocode result for address
func GetMockGeocode(address string) *pb.GeocodeResult {
	return MockGeocodes[address]
}

// GetMockReverseGeocode returns address for location
func GetMockReverseGeocode(lat, lng float64) string {
	key := fmt.Sprintf("%.4f,%.4f", lat, lng)
	return MockAddresses[key]
}

// GetMockRoute returns route based on mode
func GetMockRoute(mode string) *pb.Route {
	if route, ok := MockRoutes[mode]; ok {
		return route
	}
	return MockRoutes["driving"] // default
}

// Helper functions
func containsIgnoreCase(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

func calculateDistance(loc1, loc2 *pb.Location) float64 {
	// Simple Euclidean distance (not accurate for lat/lng, but for mock)
	dx := loc1.Lng - loc2.Lng
	dy := loc1.Lat - loc2.Lat
	return math.Sqrt(dx*dx+dy*dy) * 111000 // rough conversion to meters
}
