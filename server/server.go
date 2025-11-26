package server

import (
	"fmt"
	"log"
	"math/rand"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nerdgarten/mock-maps/service/data"
	"github.com/nerdgarten/mock-maps/service/types"
)

// MapsServer exposes HTTP handlers that emulate a maps API.
type MapsServer struct{}

// NewMapsServer creates a new MapsServer instance.
func NewMapsServer() *MapsServer {
	return &MapsServer{}
}

// RegisterRoutes wires all HTTP endpoints into the provided Gin router.
func (s *MapsServer) RegisterRoutes(r *gin.Engine) {
	r.POST("/directions", s.handleGetDirections)
	r.POST("/places/search", s.handleSearchPlaces)
	r.POST("/geocode", s.handleGeocode)
	r.POST("/reverse-geocode", s.handleReverseGeocode)
	r.POST("/distance-matrix", s.handleDistanceMatrix)
	r.POST("/optimize-route", s.handleOptimizeRoute)
	r.GET("/track", s.handleTrack)
	r.POST("/static-map", s.handleGetStaticMap)
	r.POST("/map-svg", s.handleGetMapSVG)
}

func (s *MapsServer) handleGetDirections(c *gin.Context) {
	var req types.GetDirectionsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request payload"})
		return
	}
	mode := req.Mode
	if mode == "" {
		mode = "driving"
	}
	log.Printf("REST GetDirections called with mode=%s", mode)
	route := data.GetMockRoute(mode)
	if route == nil {
		c.JSON(404, gin.H{"error": "unsupported mode"})
		return
	}
	c.JSON(200, route)
}

func (s *MapsServer) handleSearchPlaces(c *gin.Context) {
	var req types.SearchPlacesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request payload"})
		return
	}
	log.Printf("REST SearchPlaces called with query=%s", req.Query)
	places := data.GetMockPlaces(req.Query, req.Location, req.Radius, req.Type)
	c.JSON(200, types.SearchPlacesResponse{Results: places})
}

func (s *MapsServer) handleGeocode(c *gin.Context) {
	var req types.GeocodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request payload"})
		return
	}
	log.Printf("REST Geocode called for address=%s", req.Address)
	result := data.GetMockGeocode(req.Address)
	if result == nil {
		c.JSON(200, types.GeocodeResponse{Results: []types.GeocodeResult{}})
		return
	}
	c.JSON(200, types.GeocodeResponse{Results: []types.GeocodeResult{*result}})
}

func (s *MapsServer) handleReverseGeocode(c *gin.Context) {
	var req types.ReverseGeocodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request payload"})
		return
	}
	log.Printf("REST ReverseGeocode called for lat=%.4f lng=%.4f", req.Location.Lat, req.Location.Lng)
	address := data.GetMockReverseGeocode(req.Location.Lat, req.Location.Lng)
	if address == "" {
		address = formatFallbackAddress(req.Location)
	}
	c.JSON(200, types.ReverseGeocodeResponse{Address: address})
}

func (s *MapsServer) handleDistanceMatrix(c *gin.Context) {
	var req types.DistanceMatrixRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request payload"})
		return
	}
	log.Printf("REST DistanceMatrix called with %d origins and %d destinations", len(req.Origins), len(req.Destinations))
	c.JSON(200, types.DistanceMatrixResponse{Rows: data.MockDistanceMatrix})
}

func (s *MapsServer) handleOptimizeRoute(c *gin.Context) {
	var req types.OptimizeRouteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request payload"})
		return
	}
	log.Printf("REST OptimizeRoute called with %d stops", len(req.Stops))
	if len(data.MockOptimizedRoutes) == 0 {
		c.JSON(500, gin.H{"error": "no mock optimized routes configured"})
		return
	}
	index := rand.Intn(len(data.MockOptimizedRoutes))
	c.JSON(200, data.MockOptimizedRoutes[index])
}

func (s *MapsServer) handleTrack(c *gin.Context) {
	log.Print("REST Track called")
	c.JSON(200, types.TrackResponse{Locations: data.MockTrackLocations})
}

func (s *MapsServer) handleGetStaticMap(c *gin.Context) {
	var req types.StaticMapRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request payload"})
		return
	}
	log.Printf("REST GetStaticMap called for lat=%.4f lng=%.4f zoom=%d size=%s", req.Center.Lat, req.Center.Lng, req.Zoom, req.Size)
	key := formatStaticMapKey(req.Center.Lat, req.Center.Lng, req.Zoom, req.Size)
	url, ok := data.MockStaticMaps[key]
	if !ok {
		url = "https://mockmaps.local/static/default.png"
	}
	c.JSON(200, types.StaticMap{URL: url})
}

func formatFallbackAddress(loc types.Location) string {
	return fmt.Sprintf("%.4f, %.4f", loc.Lat, loc.Lng)
}

func formatStaticMapKey(lat, lng float64, zoom int32, size string) string {
	key := fmt.Sprintf("%.4f,%.4f_%d_%s", lat, lng, zoom, size)
	return strings.TrimRight(key, "_")
}

func (s *MapsServer) handleGetMapSVG(c *gin.Context) {
	var req types.MapSVGRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request payload"})
		return
	}

	// Default dimensions if not provided
	if req.Width == 0 {
		req.Width = 800
	}
	if req.Height == 0 {
		req.Height = 600
	}
	if req.Zoom == 0 {
		req.Zoom = 14
	}

	log.Printf("REST GetMapSVG called for lat=%.4f lng=%.4f zoom=%d size=%dx%d with %d markers",
		req.Center.Lat, req.Center.Lng, req.Zoom, req.Width, req.Height, len(req.Markers))

	svg := data.GenerateMapSVG(req.Center, req.Zoom, req.Width, req.Height, req.Markers)
	c.JSON(200, types.MapSVGResponse{SVG: svg})
}
