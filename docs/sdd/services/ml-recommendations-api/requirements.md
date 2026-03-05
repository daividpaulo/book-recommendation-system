# ML Recommendations API Requirements

## Scope

Embedding, training, and inference service for book recommendations.

## Owned Features

- `model-training` (technical owner)
- `recommendations` (technical owner)
- embedding support for `book-catalog` and `user-management`
- profile enrichment support for `book-purchases`

## Functional Requirements

- `REQ-ML-001`: Expose `GET /health` with model/context readiness.
- `REQ-ML-002`: Accept users and generate embeddings via `POST /embed/users`.
- `REQ-ML-003`: Accept books and generate embeddings via `POST /embed/books`.
- `REQ-ML-004`: Train model via `POST /train` with `{ users, books }`, where users include purchase profile fields.
- `REQ-ML-005`: Persist trained model and encoding context.
- `REQ-ML-006`: Generate recommendations via `POST /recommend`.
- `REQ-ML-007`: Support manual training as the default PoC flow.
- `REQ-ML-008`: Persist model artifact and encoding context in MinIO as mandatory storage backend.
- `REQ-ML-009`: Recommendation scoring must consider age and purchase-derived profile signals after training.

## Architecture Requirements (Clean Architecture)

- `ARC-ML-001`: `domain` cannot depend on `tensorflow`, `fastapi`, or `qdrant-client`.
- `ARC-ML-002`: `application` depends only on `domain` and `interfaces` (ports).
- `ARC-ML-003`: `interfaces` defines contracts for encoder, model repository, and vector store.
- `ARC-ML-004`: `infrastructure` only implements `interfaces` contracts.
- `ARC-ML-005`: `api` acts as inbound adapter and must not contain business rules.
- `ARC-ML-006`: Dependency composition occurs at startup boundary (`api/main.py`).
- `ARC-ML-007`: Model persistence is handled by `model_repository`, not `vector_store`.

## Non-Functional Requirements

- `NFR-ML-001`: Persist model/context in MinIO and fail fast when MinIO is unavailable.
- `NFR-ML-002`: Inference failures return HTTP errors without process termination.
- `NFR-ML-003`: CPU-only operation is supported.
- `NFR-ML-004`: Ensure per-layer testability (unit tests in `application`, integration tests in `api`/`infrastructure`).
- `NFR-ML-005`: MinIO is a required dependency in Docker Compose for ML service startup.
