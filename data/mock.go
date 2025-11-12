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

// GenerateMapSVG generates an SVG representation of a map with markers
func GenerateMapSVG(center types.Location, zoom int32, width, height int32, markers []types.Marker) string {
	// Calculate scale based on zoom (higher zoom = more detail)
	scale := math.Pow(2, float64(zoom-10)) * 100000

	// Build SVG header
	svg := fmt.Sprintf(`<svg width="%d" height="%d" xmlns="http://www.w3.org/2000/svg">`, width, height)

	// Add background (map area)
	svg += fmt.Sprintf(`<rect width="%d" height="%d" fill="#e8f4f8"/>`, width, height)

	// Add grid pattern for map feel
	svg += `<defs><pattern id="grid" width="40" height="40" patternUnits="userSpaceOnUse">
		<path d="M 40 0 L 0 0 0 40" fill="none" stroke="#d0e8f0" stroke-width="1"/>
	</pattern></defs>`
	svg += fmt.Sprintf(`<rect width="%d" height="%d" fill="url(#grid)"/>`, width, height)

	// Add some streets/paths
	centerX := float64(width) / 2
	centerY := float64(height) / 2

	// Horizontal streets
	for i := -2; i <= 2; i++ {
		y := centerY + float64(i*80)
		if y >= 0 && y <= float64(height) {
			svg += fmt.Sprintf(`<line x1="0" y1="%.1f" x2="%d" y2="%.1f" stroke="#c4d8e0" stroke-width="3"/>`,
				y, width, y)
		}
	}

	// Vertical streets
	for i := -3; i <= 3; i++ {
		x := centerX + float64(i*80)
		if x >= 0 && x <= float64(width) {
			svg += fmt.Sprintf(`<line x1="%.1f" y1="0" x2="%.1f" y2="%d" stroke="#c4d8e0" stroke-width="3"/>`,
				x, x, height)
		}
	}

	// Add some park areas (green spaces)
	svg += fmt.Sprintf(`<circle cx="%.1f" cy="%.1f" r="50" fill="#b8e6b8" opacity="0.6"/>`,
		centerX-120, centerY-100)
	svg += fmt.Sprintf(`<circle cx="%.1f" cy="%.1f" r="35" fill="#b8e6b8" opacity="0.6"/>`,
		centerX+150, centerY+80)

	// Add buildings
	buildings := []struct{ x, y, w, h float64 }{
		{centerX - 60, centerY - 60, 40, 40},
		{centerX + 20, centerY - 80, 50, 50},
		{centerX - 100, centerY + 20, 35, 35},
		{centerX + 80, centerY - 20, 45, 45},
	}

	for _, b := range buildings {
		if b.x >= 0 && b.x+b.w <= float64(width) && b.y >= 0 && b.y+b.h <= float64(height) {
			svg += fmt.Sprintf(`<rect x="%.1f" y="%.1f" width="%.1f" height="%.1f" fill="#d4d4d4" stroke="#999" stroke-width="1"/>`,
				b.x, b.y, b.w, b.h)
		}
	}

	// Add markers
	for _, marker := range markers {
		// Convert lat/lng to SVG coordinates relative to center
		dx := (marker.Location.Lng - center.Lng) * scale
		dy := (center.Lat - marker.Location.Lat) * scale

		x := centerX + dx
		y := centerY + dy

		// Only draw marker if it's within bounds
		if x >= 0 && x <= float64(width) && y >= 0 && y <= float64(height) {
			color := marker.Color
			if color == "" {
				color = "#ff4444"
			}

			// Draw pin (teardrop shape)
			svg += fmt.Sprintf(`<g transform="translate(%.1f,%.1f)">`, x, y)
			svg += fmt.Sprintf(`<path d="M 0,-20 C -8,-20 -15,-13 -15,-5 C -15,0 0,20 0,20 C 0,20 15,0 15,-5 C 15,-13 8,-20 0,-20 Z" fill="%s" stroke="#fff" stroke-width="2"/>`, color)
			svg += `<circle cx="0" cy="-8" r="5" fill="#fff"/>`
			svg += `</g>`

			// Add label if provided
			if marker.Label != "" {
				svg += fmt.Sprintf(`<text x="%.1f" y="%.1f" text-anchor="middle" font-size="12" fill="#333" font-weight="bold">%s</text>`,
					x, y+25, marker.Label)
			}
		}
	}

	// Add center indicator (small circle)
	svg += fmt.Sprintf(`<circle cx="%.1f" cy="%.1f" r="3" fill="#4285f4" stroke="#fff" stroke-width="1"/>`,
		centerX, centerY)

	svg += `</svg>`
	return svg
}
