# Build stage
FROM golang:1.24.4-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main .

# Runtime stage
FROM alpine:latest

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Expose the port
EXPOSE 50051

# Set environment variable for port
ENV PORT=50051

# Run the application
CMD ["./main"]