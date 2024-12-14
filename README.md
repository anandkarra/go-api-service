# Risks API Service

## Overview
This Go-based Risk management application uses the Gin web framework to provide REST endpoints for creating and retrieving risks.

## Prerequisites
- Go 1.23 or higher

## Installation and Startup
- Clone the repository
- Download dependencies by running: `go mod tidy`
- Start the server on `http://localhost:8080` by running: `go run main.go` 

## Endpoints
### Create a Risk

```sh
$ curl -X POST http://localhost:8080/v1/risks \
  -H "Content-Type: application/json" \
  -d '{"state":"<open|closed|accepted|investigating>","title":"<title>","description":"<description>"}'
```

Example:
```sh
$ curl -X POST http://localhost:8080/v1/risks \
  -H "Content-Type: application/json" \
  -d '{"state":"open","title":"Privilage escalation","description":"Potential privilage escalation vulnerability"}'
{"id":"c61ff24a-e33a-485e-8e83-60a1d0a42906","state":"open","title":"Privilage escalation","description":"Potential privilage escalation vulnerability"}
```

### Get a specific risk: 

`$ curl http://localhost:8080/v1/risks/<uuid>`

Example:
```sh
$ curl http://localhost:8080/v1/risks/c61ff24a-e33a-485e-8e83-60a1d0a42906
{"id":"c61ff24a-e33a-485e-8e83-60a1d0a42906","state":"open","title":"Privilage escalation","description":"Potential privilage escalation vulnerability"}
```

### List risks: 

`$ curl http://localhost:8080/v1/risks`

Example:
```sh
$ curl http://localhost:8080/v1/risks
[{"id":"c61ff24a-e33a-485e-8e83-60a1d0a42906","state":"open","title":"Privilage escalation","description":"Potential privilage escalation vulnerability"}]
```

## Docker Deployment

## Prerequisites
- Docker
- Docker Compose

## Build and Run through build script

```sh
# Make the build script executable
chmod +x build-script.sh

# Build the Docker image
./build-script.sh build

# Run the application
./build-script.sh run

# Stop the application
./build-script.sh stop

# Clean up all resources
./build-script.sh clean
```

## Direct Docker build and run

```sh
# Build the image
docker build -t go-api-service .

# Run the container
docker run -p 8080:8080 go-api-service
```

## Unit Testing

- Run tests: `go test ./...`
- Get test coverage: `go test ./... -cover`
- Generate coverage profile and view it in a web browser: `go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out -o coverage.html`

Current test coverage:
```sh
$ go test ./... -cover
        go-api-service          coverage: 0.0% of statements
ok      go-api-service/handlers 0.007s  coverage: 100.0% of statements
ok      go-api-service/models   (cached)        coverage: 96.6% of statements
```
