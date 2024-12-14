## Requirements
### Primary
- (As given in problem statement)

### Other
- Containerization for simplified deployment and scalability as a micro-service

## API Design

### Endpoint specifications
- Create endpoint
  - Method: `POST`
  - Endpoint: `/v1/risks`
  - Query parameters: None
  - Request body: Title (string), description (string), state (string; one of: open, closed, accepted or investigating)
  - Response codes:
    - 201: Successfully created Risk entry
    - 400: Invalid input
    - 500: Internal server error
- Get endpoint
  - Method: `GET`
  - Endpoint: `/v1/risks`
  - Query parameters: ID (UUID string)
  - Request body: None
  - Response codes:
    - 200: Successfully fetched Risk entry
    - 400: Invalid input
    - 404: Risk entry not found
    - 500: Internal server error
  - Response: JSON object (containing Risk UUID, title, description and state)
- List endpoint
  - Method: `GET`
  - Endpoint: `/v1/risks`
  - Query parameters: None
  - Request body: None
  - Response codes:
    - 200: Successfully listed Risk entries
    - 404: Risk entries not found
    - 500: Internal server error
  - Response: JSON list (with each entry containing Risk UUID, title, description and state)

### Input Validation
- All the three parameters of a Risk entry (title, description and state) are assumed to be mandatory.
- The title and description values are assumed to only contain alphanumeric, period, hyphen and underscore characters. The same is validated with the regex: `^[a-zA-Z0-9\.\-_ ]+$`
- The title and description values are assumed to be no longer than 100 characters.
- The input ID for the Get endpoint is validated against the standard UUID regex: `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`

## Versioning

Although only three endpoints are defined in the requirements, the codebase is structured such that more endpoints or newer versions can be added without any major changes in the execution flow.

## Data Model
As per the given requirements, the Risk items are stored in-memory in a map data structure with the ID as key.

```go
map[uuid.UUID]*models.Risk
```

With the ID being of type `uuid.UUID` (from the package `"github.com/google/uuid"`) and the rest of the parameters being strings.

As data transfer is expected to be done in JSON, each of these parameters have their JSON keys defined in their model. Further, the `JSON` and `ShouldBindJSON` methods from `gin.Context` is used for data serialization and deserialization respectively.

```go
type Risk struct {
  ID          uuid.UUID `json:"id"`
  State       string    `json:"state"`
  Title       string    `json:"title"`
  Description string    `json:"description"`
}
```

### Sample representations:

- Input payload for create endpoint:

```json
{
  "description": "Potential privilage escalation vulnerability",
  "state": "open",
  "title": "Privilage escalation"
}
```

- Output of get endpoint:

```json
{
  "description": "Potential privilage escalation vulnerability",
  "id": "c61ff24a-e33a-485e-8e83-60a1d0a42906",
  "state": "open",
  "title": "Privilage escalation"
}
```

- Output of list endpoint:

```json
[
  {
    "description": "Potential privilage escalation vulnerability",
    "id": "c61ff24a-e33a-485e-8e83-60a1d0a42906",
    "state": "open",
    "title": "Privilage escalation"
  }
]
```

### Memory Utilization Estimation

The memory consumed is approximately proportional: No. of entries * (Key + Value sizes) + Overhead

- Key (UUID) = 16 bytes (The github.com/google/uuid package in Go represents UUIDs as a 16-byte array ([16]byte) internally, not as strings)
- Values:
  - ID (UUID) = 16 bytes
  - State <= 13 bytes (Largest state value being: investigating)
  - Title <=100 bytes (The length of title strings are limited to 100 characters)
  - Description <=100 bytes (The length of description strings are limited to 100 characters)
  - Total struct size <= 229 bytes
- Total key + value size <= 16 + (16 + 13 + 100 + 100) <= 245 bytes

Approximate memory utilization estimations:
- For 1 million enties, memory usage = 1,000,000 * 245 bytes = 2.45 GB
- For 100 million enties, memory usage = 100,000,000 * 245 bytes = 245 GB
- For 1 billion enties, memory usage = 1,000,000,000 * 245 bytes = 2.45 TB


## Performance Considerations

Data is stored in-memory in a map data structure. The map data structure in Go is implemented with hash tables which provide constant-time read and write operations.

## Testing

Unit tests are included for all business logic methods in accompanying test files.

Over 95% code coverage has been achieved through these tests.

## Deployment

This service can be built or run directly as a Docker container with the provided build script (`build-script.sh`), Dockerfile and `docker-compose.yml`. The service is deployed using a slim Alpine image to reduce resource usage and attack surface.

## Future Extensibility
- **Pagination**: As the number of entries grows, pagination must be added to the list endpoint to manage system and network usage. Since Go maps use hash tables that are not thread-safe, a mutex is required for concurrent operations.
- **Intergration testing**: Despite high unit test coverage, integration tests are needed to verify the API service's real-world functionality. These can be implemented as a standalone utility in Go or with a framework like Pytest.
- **Mock service design**: Using interfaces for a mock service design improves separation of data and business logic, enables easy method implementation changes (e.g., switching from an in-memory store to a database), and simplifies dependency injection and unit testing.
- **Generalized error handling**: As the number endpoints increase, separately handling errors for each code path becomes tedious. Maintaining a common list of errors with their response codes and messages ensures a more scalable service.
- **Authtication and authorization mechanisms**: Authentication and authorization checks can be added to the handler methods or through a middleware to improve security and manage potential abuse.
- **CI/CD**: The testing, building and deployment of this service can be automated using CI/CD pipelines like GitHub Actions, Travis etc. to accelerate delivery and reduce developer workload.
- **Logging and monitoring**: A high-performance logger like `zap` (`github.com/uber-go/zap`) can be used to provide structured logging with reduced overhead across the service. Further, these logs along with other metrics can be pushed to logDNA and monitoring services like Prometheus respectively.