# Recommendations Data Spec

## Components

- PostgreSQL: transactional source of truth for domain entities.
- Qdrant: vector index for recommendation candidates.
- MinIO: S3-compatible object storage for artifact simulation.

## Bootstrap Design

- SQL scripts in `db/init/` create schema and load seed data.
- Named Docker volumes preserve state across restarts.
- Startup order is managed in `docker-compose.yml`.

## Operational Contracts

- PostgreSQL: `5432`
- Qdrant API/dashboard: `6333`
- MinIO S3 API: `9000`
- MinIO console: `9001`

## ML Persistence Strategy

- Current PoC: model/context persisted in local `ml-recommendations-api` volume.
- Evolution-ready: MinIO available for model artifact version migration.
- Qdrant remains exclusive to user/book embeddings and vector search.

## Acceptance Criteria

- Containers start in valid dependency order.
- Seed data is available after volume reset.
- API/ML integration works with relational and vector components.
