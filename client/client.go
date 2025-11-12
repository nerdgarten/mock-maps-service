package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/nerdgarten/mock-maps/service/types"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "50051"
	}
	baseURL := fmt.Sprintf("http://localhost:%s", port)
	client := &http.Client{Timeout: 5 * time.Second}

	// Example 1: Get Directions
	directionsReq := types.GetDirectionsRequest{
		Origin:      &types.Location{Lat: 13.7563, Lng: 100.5018},
		Destination: &types.Location{Lat: 13.7383, Lng: 100.5323},
		Mode:        "driving",
	}
	var route types.Route
	if err := postJSON(client, baseURL+"/directions", directionsReq, &route); err != nil {
		log.Printf("GetDirections failed: %v", err)
	} else {
		log.Printf("Directions: Distance: %s, Duration: %s", route.Distance, route.Duration)
	}

	// Example 2: Search Places
	placesReq := types.SearchPlacesRequest{
		Query:    "pizza",
		Location: &types.Location{Lat: 13.7563, Lng: 100.5018},
		Radius:   3000,
		Type:     "restaurant",
	}
	var placesResp types.SearchPlacesResponse
	if err := postJSON(client, baseURL+"/places/search", placesReq, &placesResp); err != nil {
		log.Printf("SearchPlaces failed: %v", err)
	} else {
		log.Printf("Found %d places", len(placesResp.Results))
		for _, place := range placesResp.Results {
			log.Printf("Place: %s at %s", place.Name, place.Address)
		}
	}

	// Example 3: Track (rest response)
	var trackResp types.TrackResponse
	if err := getJSON(client, baseURL+"/track", &trackResp); err != nil {
		log.Printf("Track failed: %v", err)
	} else {
		log.Printf("Received %d location updates", len(trackResp.Locations))
		for _, loc := range trackResp.Locations {
			log.Printf("Location: Lat: %.4f, Lng: %.4f at %s", loc.Location.Lat, loc.Location.Lng, loc.Timestamp)
		}
	}

	// Example 4: Get Static Map
	staticMapReq := types.StaticMapRequest{
		Center: types.Location{Lat: 13.7563, Lng: 100.5018},
		Zoom:   14,
		Size:   "600x400",
	}
	var staticMap types.StaticMap
	if err := postJSON(client, baseURL+"/static-map", staticMapReq, &staticMap); err != nil {
		log.Printf("GetStaticMap failed: %v", err)
	} else {
		log.Printf("Static Map URL: %s", staticMap.URL)
	}
}

func postJSON(client *http.Client, url string, payload any, out any) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal request: %w", err)
	}
	resp, err := client.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("perform request: %w", err)
	}
	defer resp.Body.Close()
	return decodeResponse(resp, out)
}

func getJSON(client *http.Client, url string, out any) error {
	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("perform request: %w", err)
	}
	defer resp.Body.Close()
	return decodeResponse(resp, out)
}

func decodeResponse(resp *http.Response, out any) error {
	if resp.StatusCode >= http.StatusBadRequest {
		data, _ := io.ReadAll(resp.Body)
		var apiErr types.ErrorResponse
		if err := json.Unmarshal(data, &apiErr); err == nil && apiErr.Error != "" {
			return fmt.Errorf("request failed (%d): %s", resp.StatusCode, apiErr.Error)
		}
		return fmt.Errorf("request failed (%d): %s", resp.StatusCode, strings.TrimSpace(string(data)))
	}
	if out == nil {
		return nil
	}
	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return fmt.Errorf("decode response: %w", err)
	}
	return nil
}
