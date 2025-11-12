package server

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"

	"github.com/nerdgarten/mock-maps/service/data"
	"github.com/nerdgarten/mock-maps/service/types"
)

// MapsServer exposes HTTP handlers that emulate a maps API.
type MapsServer struct{}

// NewMapsServer creates a new MapsServer instance.
func NewMapsServer() *MapsServer {
	return &MapsServer{}
}

// RegisterRoutes wires all HTTP endpoints into the provided mux.
func (s *MapsServer) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/directions", s.handleGetDirections)
	mux.HandleFunc("/places/search", s.handleSearchPlaces)
	mux.HandleFunc("/geocode", s.handleGeocode)
	mux.HandleFunc("/reverse-geocode", s.handleReverseGeocode)
	mux.HandleFunc("/distance-matrix", s.handleDistanceMatrix)
	mux.HandleFunc("/optimize-route", s.handleOptimizeRoute)
	mux.HandleFunc("/track", s.handleTrack)
	mux.HandleFunc("/static-map", s.handleGetStaticMap)
}

func (s *MapsServer) handleGetDirections(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeMethodNotAllowed(w)
		return
	}
	var req types.GetDirectionsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	mode := req.Mode
	if mode == "" {
		mode = "driving"
	}
	log.Printf("REST GetDirections called with mode=%s", mode)
	route := data.GetMockRoute(mode)
	if route == nil {
		writeError(w, http.StatusNotFound, "unsupported mode")
		return
	}
	writeJSON(w, http.StatusOK, route)
}

func (s *MapsServer) handleSearchPlaces(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeMethodNotAllowed(w)
		return
	}
	var req types.SearchPlacesRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	log.Printf("REST SearchPlaces called with query=%s", req.Query)
	places := data.GetMockPlaces(req.Query, req.Location, req.Radius, req.Type)
	writeJSON(w, http.StatusOK, types.SearchPlacesResponse{Results: places})
}

func (s *MapsServer) handleGeocode(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeMethodNotAllowed(w)
		return
	}
	var req types.GeocodeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	log.Printf("REST Geocode called for address=%s", req.Address)
	result := data.GetMockGeocode(req.Address)
	if result == nil {
		writeJSON(w, http.StatusOK, types.GeocodeResponse{Results: []types.GeocodeResult{}})
		return
	}
	writeJSON(w, http.StatusOK, types.GeocodeResponse{Results: []types.GeocodeResult{*result}})
}

func (s *MapsServer) handleReverseGeocode(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeMethodNotAllowed(w)
		return
	}
	var req types.ReverseGeocodeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	log.Printf("REST ReverseGeocode called for lat=%.4f lng=%.4f", req.Location.Lat, req.Location.Lng)
	address := data.GetMockReverseGeocode(req.Location.Lat, req.Location.Lng)
	if address == "" {
		address = formatFallbackAddress(req.Location)
	}
	writeJSON(w, http.StatusOK, types.ReverseGeocodeResponse{Address: address})
}

func (s *MapsServer) handleDistanceMatrix(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeMethodNotAllowed(w)
		return
	}
	var req types.DistanceMatrixRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	log.Printf("REST DistanceMatrix called with %d origins and %d destinations", len(req.Origins), len(req.Destinations))
	writeJSON(w, http.StatusOK, types.DistanceMatrixResponse{Rows: data.MockDistanceMatrix})
}

func (s *MapsServer) handleOptimizeRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeMethodNotAllowed(w)
		return
	}
	var req types.OptimizeRouteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	log.Printf("REST OptimizeRoute called with %d stops", len(req.Stops))
	if len(data.MockOptimizedRoutes) == 0 {
		writeError(w, http.StatusInternalServerError, "no mock optimized routes configured")
		return
	}
	index := rand.Intn(len(data.MockOptimizedRoutes))
	writeJSON(w, http.StatusOK, data.MockOptimizedRoutes[index])
}

func (s *MapsServer) handleTrack(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeMethodNotAllowed(w)
		return
	}
	log.Print("REST Track called")
	writeJSON(w, http.StatusOK, types.TrackResponse{Locations: data.MockTrackLocations})
}

func (s *MapsServer) handleGetStaticMap(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeMethodNotAllowed(w)
		return
	}
	var req types.StaticMapRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	log.Printf("REST GetStaticMap called for lat=%.4f lng=%.4f zoom=%d size=%s", req.Center.Lat, req.Center.Lng, req.Zoom, req.Size)
	key := formatStaticMapKey(req.Center.Lat, req.Center.Lng, req.Zoom, req.Size)
	url, ok := data.MockStaticMaps[key]
	if !ok {
		url = "https://mockmaps.local/static/default.png"
	}
	writeJSON(w, http.StatusOK, types.StaticMap{URL: url})
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("failed to encode response: %v", err)
	}
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, types.ErrorResponse{Error: message})
}

func writeMethodNotAllowed(w http.ResponseWriter) {
	writeError(w, http.StatusMethodNotAllowed, "method not allowed")
}

func formatFallbackAddress(loc types.Location) string {
	return fmt.Sprintf("%.4f, %.4f", loc.Lat, loc.Lng)
}

func formatStaticMapKey(lat, lng float64, zoom int32, size string) string {
	key := fmt.Sprintf("%.4f,%.4f_%d_%s", lat, lng, zoom, size)
	return strings.TrimRight(key, "_")
}
