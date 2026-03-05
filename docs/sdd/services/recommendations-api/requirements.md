# Recommendations API Requirements

## Scope

Orchestration and persistence service for books, users, purchases, training trigger, and recommendation gateway endpoints.

## Owned Features

- `book-catalog`
- `user-management`
- `model-training` (orchestration)
- `recommendations` (ML HTTP gateway)
- `book-purchases`

## Functional Requirements

- `REQ-API-001`: Expose `GET /health` with service status.
- `REQ-API-002`: Create books via `POST /api/v1/books`.
- `REQ-API-003`: List books via `GET /api/v1/books`.
- `REQ-API-004`: Create users via `POST /api/v1/users`.
- `REQ-API-005`: List users via `GET /api/v1/users`.
- `REQ-API-006`: Trigger training via `POST /api/v1/recommendations/train`.
- `REQ-API-007`: Fetch recommendations via `GET /api/v1/recommendations/{userId}`.
- `REQ-API-008`: Validate user existence before requesting recommendations.
- `REQ-API-009`: Call `POST /embed/books` after book creation without blocking create response.
- `REQ-API-010`: Call `POST /embed/users` after user creation without blocking create response.
- `REQ-API-011`: Register purchases via `POST /api/v1/purchases`.
- `REQ-API-012`: Expose user purchase history via `GET /api/v1/users/{userId}/purchases`.
- `REQ-API-013`: Enrich training payload with purchase profile fields (`purchase_count`, `purchased_book_ids`).

## Architecture Requirements (Clean Architecture)

- `ARC-API-001`: `cmd/server/main.go` must contain bootstrap/wiring only.
- `ARC-API-002`: Business rules must live in `internal/usecase`.
- `ARC-API-003`: Entities and contracts must live in `internal/domain`.
- `ARC-API-004`: HTTP inbound adapters must live in `internal/delivery/http`.
- `ARC-API-005`: Outbound adapters (Postgres/ML client) must live in `internal/repository`.
- `ARC-API-006`: Dependencies must point inward (`delivery -> usecase -> domain`).

## Non-Functional Requirements

- `NFR-API-001`: Start only after database and ML API are healthy.
- `NFR-API-002`: Return JSON on all public routes.
- `NFR-API-003`: Handle ML API failures without crashing the API process.
- `NFR-API-004`: Ensure per-layer testability (unit tests for use cases, integration tests for adapters).
