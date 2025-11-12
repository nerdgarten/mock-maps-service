package data

import (
	"fmt"
	"math"
	"strings"

	"github.com/nerdgarten/mock-maps/service/types"
)

// MockPlaces contains a list of mock places for search
var MockPlaces = []types.Place{
	{Name: "Pizza Planet", Address: "123 Main St", Location: types.Location{Lat: 13.7563, Lng: 100.5018}, Rating: 4.5},
	{Name: "Cheesy Bites", Address: "45 King Rd", Location: types.Location{Lat: 13.7570, Lng: 100.5030}, Rating: 4.2},
	{Name: "Burger Barn", Address: "78 Elm Ave", Location: types.Location{Lat: 13.7580, Lng: 100.5040}, Rating: 4.0},
	{Name: "Taco Town", Address: "90 Oak St", Location: types.Location{Lat: 13.7590, Lng: 100.5050}, Rating: 4.3},
	{Name: "Sushi Spot", Address: "12 Pine Rd", Location: types.Location{Lat: 13.7600, Lng: 100.5060}, Rating: 4.7},
	{Name: "Coffee Corner", Address: "34 Maple Ln", Location: types.Location{Lat: 13.7610, Lng: 100.5070}, Rating: 4.1},
	{Name: "Ice Cream Island", Address: "56 Birch Blvd", Location: types.Location{Lat: 13.7620, Lng: 100.5080}, Rating: 4.6},
	{Name: "Pasta Palace", Address: "78 Cedar Ct", Location: types.Location{Lat: 13.7630, Lng: 100.5090}, Rating: 4.4},
}

// MockGeocodes contains mock geocoding results
var MockGeocodes = map[string]*types.GeocodeResult{
	"Chulalongkorn University": {FormattedAddress: "254 Phayathai Rd, Pathum Wan, Bangkok", Location: types.Location{Lat: 13.7383, Lng: 100.5323}},
	"Central World":            {FormattedAddress: "999/9 Rama I Rd, Pathum Wan, Bangkok", Location: types.Location{Lat: 13.7465, Lng: 100.5393}},
	"Grand Palace":             {FormattedAddress: "Na Phra Lan Rd, Phra Borom Maha Ratchawang, Phra Nakhon, Bangkok", Location: types.Location{Lat: 13.7500, Lng: 100.4925}},
	"Siam Paragon":             {FormattedAddress: "991/1 Rama I Rd, Pathum Wan, Bangkok", Location: types.Location{Lat: 13.7462, Lng: 100.5347}},
}

// MockAddresses contains mock reverse geocoding results
var MockAddresses = map[string]string{
	"13.7563,100.5018": "123 Main St, Bangkok, Thailand",
	"13.7383,100.5323": "254 Phayathai Rd, Pathum Wan, Bangkok",
	"13.7465,100.5393": "999/9 Rama I Rd, Pathum Wan, Bangkok",
	"13.7500,100.4925": "Na Phra Lan Rd, Phra Borom Maha Ratchawang, Phra Nakhon, Bangkok",
}

// MockRoutes contains mock routes for different modes
var MockRoutes = map[string]*types.Route{
	"driving": {
		Distance: "5.3 km",
		Duration: "12 mins",
		Polyline: "mock_encoded_polyline_driving",
		Steps: []types.Step{
			{Instruction: "Head north on Main St", Distance: "200m"},
			{Instruction: "Turn right onto 1st Ave", Distance: "500m"},
			{Instruction: "Continue straight on 2nd St", Distance: "3.6 km"},
		},
	},
	"walking": {
		Distance: "5.3 km",
		Duration: "65 mins",
		Polyline: "mock_encoded_polyline_walking",
		Steps: []types.Step{
			{Instruction: "Walk north on Main St", Distance: "200m"},
			{Instruction: "Turn right onto 1st Ave", Distance: "500m"},
			{Instruction: "Continue on sidewalk along 2nd St", Distance: "3.6 km"},
		},
	},
	"bicycling": {
		Distance: "5.5 km",
		Duration: "18 mins",
		Polyline: "mock_encoded_polyline_bicycling",
		Steps: []types.Step{
			{Instruction: "Bike north on Main St", Distance: "200m"},
			{Instruction: "Turn right onto 1st Ave", Distance: "500m"},
			{Instruction: "Follow bike path on 2nd St", Distance: "3.8 km"},
		},
	},
}

// MockDistanceMatrix contains mock distance matrix data
var MockDistanceMatrix = []types.DistanceRow{
	{
		Elements: []types.DistanceElement{
			{Distance: "1.2 km", Duration: "4 mins"},
			{Distance: "3.8 km", Duration: "10 mins"},
			{Distance: "2.5 km", Duration: "7 mins"},
		},
	},
}

// MockOptimizedRoutes contains mock optimized routes
var MockOptimizedRoutes = []*types.OptimizedRoute{
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
var MockTrackLocations = []types.TrackLocation{
	{Location: types.Location{Lat: 13.7563, Lng: 100.5018}, Timestamp: "2025-10-18T14:30:00Z"},
	{Location: types.Location{Lat: 13.7570, Lng: 100.5025}, Timestamp: "2025-10-18T14:30:10Z"},
	{Location: types.Location{Lat: 13.7575, Lng: 100.5030}, Timestamp: "2025-10-18T14:30:20Z"},
	{Location: types.Location{Lat: 13.7580, Lng: 100.5035}, Timestamp: "2025-10-18T14:30:30Z"},
	{Location: types.Location{Lat: 13.7585, Lng: 100.5040}, Timestamp: "2025-10-18T14:30:40Z"},
}

// MockStaticMaps contains mock static map URLs
var MockStaticMaps = map[string]string{
	"13.7563,100.5018_14_600x400": "https://mockmaps.local/static/preview1.png",
	"13.7383,100.5323_15_800x600": "https://mockmaps.local/static/preview2.png",
}

// GetMockPlaces returns filtered places based on query
func GetMockPlaces(query string, location *types.Location, radius int32, placeType string) []types.Place {
	var results []types.Place
	for _, place := range MockPlaces {
		if query != "" && !containsIgnoreCase(place.Name, query) {
			continue
		}
		if placeType != "" && !containsIgnoreCase(place.Name, placeType) {
			continue
		}
		if location != nil && radius > 0 {
			if calculateDistance(*location, place.Location) > float64(radius) {
				continue
			}
		}
		results = append(results, place)
	}
	return results
}

// GetMockGeocode returns geocode result for address
func GetMockGeocode(address string) *types.GeocodeResult {
	return MockGeocodes[address]
}

// GetMockReverseGeocode returns address for location
func GetMockReverseGeocode(lat, lng float64) string {
	key := fmt.Sprintf("%.4f,%.4f", lat, lng)
	return MockAddresses[key]
}

// GetMockRoute returns route based on mode
func GetMockRoute(mode string) *types.Route {
	if route, ok := MockRoutes[strings.ToLower(mode)]; ok {
		return route
	}
	return MockRoutes["driving"]
}

// Helper functions
func containsIgnoreCase(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

func calculateDistance(loc1, loc2 types.Location) float64 {
	dx := loc1.Lng - loc2.Lng
	dy := loc1.Lat - loc2.Lat
	return math.Sqrt(dx*dx+dy*dy) * 111000
}
