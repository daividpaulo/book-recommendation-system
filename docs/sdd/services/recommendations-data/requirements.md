# Recommendations Data Requirements

## Scope

Data platform infrastructure with PostgreSQL, Qdrant, and MinIO.

## Functional Requirements

- `REQ-DATA-001`: Provision PostgreSQL with initial `users` and `books` schema.
- `REQ-DATA-002`: Load local seed data for bootstrap.
- `REQ-DATA-003`: Provision Qdrant for vector persistence/search.
- `REQ-DATA-004`: Expose Qdrant dashboard locally.
- `REQ-DATA-005`: Provision MinIO for mandatory ML model/context artifact persistence.
- `REQ-DATA-006`: Expose MinIO console locally for bucket/object inspection.

## Non-Functional Requirements

- `NFR-DATA-001`: Use Docker volumes for persistence.
- `NFR-DATA-002`: Keep relational DB healthcheck enabled.
- `NFR-DATA-003`: Application services must depend on data platform readiness.
- `NFR-DATA-004`: ML service must depend on MinIO readiness before startup.
