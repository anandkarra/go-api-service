## Requirements:
### Primary:
- (As given in problem statement)

### Other:
- Containerization for simplified deployment

## API Design:
- Create endpoint
  - Method: `POST`
  - Endpoint: `/v1/risks`
  - Query parameters: None
  - Request body: Title (string), description (string), state (string; one of: open, closed, accepted, investigating)
  - Response codes:
    - 200: Successfully created Risk entry
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
  - Response: JSON (containing Risk UUID, title, description and state)
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
