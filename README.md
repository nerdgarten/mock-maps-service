# Mock Maps Service

This is a mock implementation of a Maps Service using gRPC for testing and development purposes.
It simulates the behavior of a real maps service (like Google Maps) without requiring external API access.

## Features

- Mock gRPC methods for common maps functionalities
- Expanded mock data with multiple places, addresses, routes, and locations
- Structured codebase with separate data and server packages
- Predefined, static, or dynamic mock responses
- Easy to set up locally for testing backend or frontend integrations
- Example client for testing the gRPC service

---

## gRPC Methods

### **1. GetDirections**
`rpc GetDirections(GetDirectionsRequest) returns (Route)`

Returns mock directions, polyline, and estimated travel time between two coordinates.

**Request Fields**
- origin: Location (lat, lng)
- destination: Location (lat, lng)
- mode: string (e.g., "driving", "walking", "bicycling")

**Example Response**
```json
{
  "distance": "5.3 km",
  "duration": "12 mins",
  "polyline": "mock_encoded_polyline",
  "steps": [
    {"instruction": "Head north on Main St", "distance": "200m"},
    {"instruction": "Turn right onto 1st Ave", "distance": "500m"}
  ]
}
```

---

### **2. SearchPlaces**
`rpc SearchPlaces(SearchPlacesRequest) returns (SearchPlacesResponse)`

Returns a list of nearby mock places such as restaurants or cafes.

**Request Fields**
- query: string
- location: Location (lat, lng)
- radius: int32
- type: string

**Example Response**

```json
{
  "results": [
    {
      "name": "Pizza Planet",
      "address": "123 Main St",
      "location": {"lat": 13.7563, "lng": 100.5018},
      "rating": 4.5
    },
    {
      "name": "Cheesy Bites",
      "address": "45 King Rd",
      "location": {"lat": 13.7570, "lng": 100.5030},
      "rating": 4.2
    }
  ]
}
```

---

### **3. Geocode**
`rpc Geocode(GeocodeRequest) returns (GeocodeResponse)`

Converts an address into mock coordinates.

**Request Fields**
- address: string

**Example Response**

```json
{
  "results": [
    {
      "formatted_address": "254 Phayathai Rd, Pathum Wan, Bangkok",
      "location": {"lat": 13.7383, "lng": 100.5323}
    }
  ]
}
```

---

### **4. ReverseGeocode**
`rpc ReverseGeocode(ReverseGeocodeRequest) returns (ReverseGeocodeResponse)`

Converts coordinates into a mock address.

**Request Fields**
- location: Location (lat, lng)

**Example Response**

```json
{
  "address": "123 Main St, Bangkok, Thailand"
}
```

---

### **5. DistanceMatrix**
`rpc DistanceMatrix(DistanceMatrixRequest) returns (DistanceMatrixResponse)`

Calculates mock distances and travel times between multiple origins and destinations.

**Request Fields**
- origins: repeated Location
- destinations: repeated Location
- mode: string

**Example Response**

```json
{
  "rows": [
    {
      "elements": [
        {"distance": "1.2 km", "duration": "4 mins"},
        {"distance": "3.8 km", "duration": "10 mins"}
      ]
    }
  ]
}
```

---

### **6. OptimizeRoute**
`rpc OptimizeRoute(OptimizeRouteRequest) returns (OptimizedRoute)`

Simulates route optimization for multiple delivery stops.

**Request Fields**
- stops: repeated Location

**Example Response**

```json
{
  "optimized_order": [0, 2, 1],
  "total_distance": "6.4 km",
  "total_duration": "15 mins"
}
```

---

### **7. Track**
`rpc Track(TrackRequest) returns (stream TrackLocation)`

Streams mock location updates for a driver.

**Request Fields**
- (empty for now)

**Example Stream**

```json
{"location": {"lat":13.7563,"lng":100.5018}, "timestamp":"2025-10-18T14:30:00Z"}
{"location": {"lat":13.7570,"lng":100.5025}, "timestamp":"2025-10-18T14:30:10Z"}
```

---

### **8. GetStaticMap**
`rpc GetStaticMap(StaticMapRequest) returns (StaticMap)`

Returns a mock static map image URL for UI display.

**Request Fields**
- center: Location (lat, lng)
- zoom: int32
- size: string

**Example Response**

```json
{
  "url": "https://mockmaps.local/static/preview.png"
}
```

---

## 🚀 Running the Server

To start the gRPC server:

```bash
go run main.go
```

The server will listen on port 50051 by default. You can set the `PORT` environment variable to use a different port:

```bash
PORT=8080 go run main.go
```

## Project Structure

- `proto/`: Contains the protobuf definitions and generated Go code
- `data/`: Mock data definitions and helper functions
- `server/`: gRPC server implementation
- `client/`: Example client for testing the service
- `main.go`: Entry point to start the server

## 🧪 Running the Client Example

To run the example client (ensure the server is running first):

```bash
go run client/client.go
```

The client connects to `localhost:50051` by default. If the server is running on a different port, set the `PORT` environment variable:

```bash
PORT=8080 go run client/client.go
```

This will demonstrate various gRPC calls to the mock service.
