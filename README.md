# Mock Maps Service

A lightweight REST API that emulates core map-provider features for local development and automated testing. The service returns deterministic mock data and avoids any external API calls.

## Features

- Pure HTTP+JSON interface – no gRPC tooling required
- Mocked responses for directions, geocoding, places search, distance matrix and more
- Deterministic datasets suitable for contract/integration tests
- Example REST client demonstrating common calls

## REST Endpoints

| Method | Path               | Description                                                        |
| ------ | ------------------ | ------------------------------------------------------------------ |
| `POST` | `/directions`      | Returns a mock route between an origin and destination.            |
| `POST` | `/places/search`   | Filters the mock place catalogue by query, type and radius.        |
| `POST` | `/geocode`         | Resolves an address string to mock coordinates.                    |
| `POST` | `/reverse-geocode` | Converts coordinates into a mock postal address.                   |
| `POST` | `/distance-matrix` | Provides mock distance/duration rows for origin/destination pairs. |
| `POST` | `/optimize-route`  | Returns one of the predefined optimised stop orders.               |
| `GET`  | `/track`           | Returns a list of mock driver location updates.                    |
| `POST` | `/static-map`      | Supplies a static map preview URL for the requested location.      |

All POST endpoints expect a JSON body that mirrors the structures in `types/types.go`. Errors are returned in the form:

```json
{
  "error": "description"
}
```

## Running the Server

```bash
go run main.go
```

The server listens on `:50051` by default. Override with `PORT`:

```bash
PORT=8080 go run main.go
```

## Example Requests

### Directions

```bash
curl -X POST http://localhost:50051/directions \
  -H "Content-Type: application/json" \
  -d '{
    "origin": {"lat": 13.7563, "lng": 100.5018},
    "destination": {"lat": 13.7383, "lng": 100.5323},
    "mode": "driving"
  }'
```

### Place Search

```bash
curl -X POST http://localhost:50051/places/search \
  -H "Content-Type: application/json" \
  -d '{
    "query": "pizza",
    "location": {"lat": 13.7563, "lng": 100.5018},
    "radius": 3000,
    "type": "restaurant"
  }'
```

### Track Locations

```bash
curl http://localhost:50051/track
```

## Project Structure

- `types/` – shared request/response models used by the server, client and data layers
- `data/` – curated mock data and helper functions
- `server/` – HTTP handlers and routing
- `client/` – sample REST client that exercises the API
- `main.go` – entrypoint wiring routes and launching the HTTP server

## Running the Sample Client

Ensure the server is running, then execute:

```bash
go run client/client.go
```

The client connects to `http://localhost:50051` by default (override via `PORT`) and logs the responses from each endpoint.
