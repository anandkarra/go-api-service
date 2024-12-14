# Risks API Service

## Overview
This Go-based Risk management application uses the Gin web framework to provide REST endpoints for creating and retrieving risks.

## Prerequisites
- Go 1.21 or higher

## Installation and Startup
- Clone the repository
- Download dependencies by running: `go mod tidy`
- Start the server on `http://localhost:8080` by running: `go run main.go` 

## Endpoints
### Create a Risk

```json
$ curl -X POST http://localhost:8080/v1/risks \
  -H "Content-Type: application/json" \
  -d '{"state":"<open|closed|accepted|investigating>","title":"<title>","description":"<description>"}'
```

Example:
```json
$ curl -X POST http://localhost:8080/v1/risks \
  -H "Content-Type: application/json" \
  -d '{"state":"open","title":"Privilage escalation","description":"Potential privilage escalation vulnerability"}'
{"id":"c61ff24a-e33a-485e-8e83-60a1d0a42906","state":"open","title":"Privilage escalation","description":"Potential privilage escalation vulnerability"}
```

### Get a specific risk: 

`$ curl http://localhost:8080/v1/risks/<uuid>`

Example:
```json
$ curl http://localhost:8080/v1/risks/c61ff24a-e33a-485e-8e83-60a1d0a42906
{"id":"c61ff24a-e33a-485e-8e83-60a1d0a42906","state":"open","title":"Privilage escalation","description":"Potential privilage escalation vulnerability"}
```

### List risks: 

`$ curl http://localhost:8080/v1/risks`

Example:
```json
$ curl http://localhost:8080/v1/risks
[{"id":"c61ff24a-e33a-485e-8e83-60a1d0a42906","state":"open","title":"Privilage escalation","description":"Potential privilage escalation vulnerability"}]
```