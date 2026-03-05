# Recommendations API Spec

## Responsibilities

- Expose the public HTTP API.
- Persist books and users in PostgreSQL.
- Persist purchases in PostgreSQL.
- Orchestrate training and recommendation calls to the ML API.

## Covered Features

- `book-catalog`
- `user-management`
- `model-training` (data aggregation + remote trigger)
- `recommendations` (validation + semantic proxy)
- `book-purchases`

## Target Architecture (Clean Architecture)

```text
recommendations-api/
  cmd/server/main.go                  # bootstrap only
  internal/
    domain/
      entities/
      ports/
    usecase/
      book/
      user/
      recommendation/
      training/
    delivery/http/
      handlers/
      router/
      dto/
    repository/
      postgres/
      mlservice/
    config/
```

## Layer Rules

- `main.go` only wires dependencies and starts the server.
- Handlers map HTTP requests/responses to use-case inputs/outputs.
- Use cases contain business logic and depend on ports/interfaces.
- Repositories and gateways implement Postgres/ML integrations.

## Contract Rules

- Input/output must be JSON.
- Error semantics:
  - `400` invalid request payload
  - `404` missing resource
  - `5xx` internal processing failure
  - `502` downstream dependency failure

## Main Flows

1. Book/user creation -> persist in Postgres -> fire async embedding call.
2. Purchase creation -> persist in Postgres.
3. Training trigger -> load local users/books/purchases -> call ML `/train`.
4. Recommendation -> validate user existence -> call ML `/recommend`.

## Dependencies

- PostgreSQL for relational persistence.
- `ml-recommendations-api` for training/inference.

## Acceptance Criteria

- `REQ-API-*`, `ARC-API-*`, and `NFR-API-*` are covered.
- `cmd/server/main.go` has no SQL/domain business logic.
- Minimum tests:
  - unit tests for use cases;
  - integration tests for HTTP handlers/repositories.
