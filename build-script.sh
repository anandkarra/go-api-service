#!/bin/bash

# Function to print usage
usage() {
    echo "Usage: $0 {build|run|stop|clean}"
    exit 1
}

# Build Docker image
build() {
    echo "Building Docker image for Risk API Service..."
    docker build -t risk-api-service .
}

# Run Docker container
run() {
    echo "Starting Risk API Service container..."
    docker-compose up -d
}

# Stop and remove containers
stop() {
    echo "Stopping Risk API Service container..."
    docker-compose down
}

# Clean up Docker resources
clean() {
    echo "Cleaning up Docker resources..."
    docker-compose down --rmi all --volumes
}

# Parse command
case "$1" in 
    build)
        build
        ;;
    run)
        run
        ;;
    stop)
        stop
        ;;
    clean)
        clean
        ;;
    *)
        usage
        ;;
esac

exit 0