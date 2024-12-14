# Use the official Golang image as the base image
FROM golang:1.22-alpine AS builder

# Setup the working directory in the container
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o go-api-service cmd/main.go

# Use a slim alpine image for the final container
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the pre-built binary
COPY --from=builder /app/go-api-service .

# Expose the port 8080
EXPOSE 8080

# Run the executable
ENTRYPOINT ["./go-api-service"]