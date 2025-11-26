package types

// Location represents a geographical coordinate.
type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// Step represents a single navigation instruction.
type Step struct {
	Instruction string `json:"instruction"`
	Distance    string `json:"distance"`
}

// Route contains aggregated navigation details between two locations.
type Route struct {
	Distance string `json:"distance"`
	Duration string `json:"duration"`
	Polyline string `json:"polyline"`
	Steps    []Step `json:"steps"`
}

// Place represents a point of interest that can be returned in search results.
type Place struct {
	Name     string   `json:"name"`
	Address  string   `json:"address"`
	Location Location `json:"location"`
	Rating   float64  `json:"rating"`
}

// GeocodeResult maps a formatted address to a specific location.
type GeocodeResult struct {
	FormattedAddress string   `json:"formatted_address"`
	Location         Location `json:"location"`
}

// DistanceElement captures the distance and duration between an origin/destination pair.
type DistanceElement struct {
	Distance string `json:"distance"`
	Duration string `json:"duration"`
}

// DistanceRow contains the computed elements for a single origin against multiple destinations.
type DistanceRow struct {
	Elements []DistanceElement `json:"elements"`
}

// OptimizedRoute describes the result of stop-order optimisation.
type OptimizedRoute struct {
	OptimizedOrder []int32 `json:"optimized_order"`
	TotalDistance  string  `json:"total_distance"`
	TotalDuration  string  `json:"total_duration"`
}

// StaticMap provides a URL for a static map preview.
type StaticMap struct {
	URL string `json:"url"`
}

// TrackLocation contains a location snapshot that may be streamed over time.
type TrackLocation struct {
	Location  Location `json:"location"`
	Timestamp string   `json:"timestamp"`
}

// GetDirectionsRequest is the payload for requesting a mock route.
type GetDirectionsRequest struct {
	Origin      *Location `json:"origin"`
	Destination *Location `json:"destination"`
	Mode        string    `json:"mode"`
}

// SearchPlacesRequest is the payload for retrieving nearby places.
type SearchPlacesRequest struct {
	Query    string    `json:"query"`
	Location *Location `json:"location"`
	Radius   int32     `json:"radius"`
	Type     string    `json:"type"`
}

// SearchPlacesResponse bundles search results.
type SearchPlacesResponse struct {
	Results []Place `json:"results"`
}

// GeocodeRequest represents an address lookup.
type GeocodeRequest struct {
	Address string `json:"address"`
}

// GeocodeResponse wraps geocode results.
type GeocodeResponse struct {
	Results []GeocodeResult `json:"results"`
}

// ReverseGeocodeRequest represents a reverse lookup payload.
type ReverseGeocodeRequest struct {
	Location Location `json:"location"`
}

// ReverseGeocodeResponse contains the resolved address string.
type ReverseGeocodeResponse struct {
	Address string `json:"address"`
}

// DistanceMatrixRequest contains collections of origins and destinations.
type DistanceMatrixRequest struct {
	Origins      []Location `json:"origins"`
	Destinations []Location `json:"destinations"`
	Mode         string     `json:"mode"`
}

// DistanceMatrixResponse wraps rows of computed distances and durations.
type DistanceMatrixResponse struct {
	Rows []DistanceRow `json:"rows"`
}

// OptimizeRouteRequest is the payload for ordering multiple stops.
type OptimizeRouteRequest struct {
	Stops []Location `json:"stops"`
}

// TrackResponse returns a list of tracked locations.
type TrackResponse struct {
	Locations []TrackLocation `json:"locations"`
}

// StaticMapRequest defines the parameters required to obtain a static map URL.
type StaticMapRequest struct {
	Center Location `json:"center"`
	Zoom   int32    `json:"zoom"`
	Size   string   `json:"size"`
}

// Marker represents a pin on the map
type Marker struct {
	Location Location `json:"location"`
	Label    string   `json:"label,omitempty"`
	Color    string   `json:"color,omitempty"`
}

// MapSVGRequest defines parameters for SVG map generation with markers
type MapSVGRequest struct {
	Center  Location `json:"center"`
	Zoom    int32    `json:"zoom"`
	Width   int32    `json:"width"`
	Height  int32    `json:"height"`
	Markers []Marker `json:"markers,omitempty"`
}

// MapSVGResponse contains the SVG content
type MapSVGResponse struct {
	SVG string `json:"svg"`
}

// ErrorResponse standardises error payloads returned by the service.
type ErrorResponse struct {
	Error string `json:"error"`
}
