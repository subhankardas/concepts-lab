# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy and download dependency using go mod
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/main /app/main

# Expose port 8080 
EXPOSE 8080

# Command to run the executable
CMD ["./main"]