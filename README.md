# Book Recommendation System

Distributed book recommendation system with:

- Go API (`recommendations-api`)
- Python ML API + TensorFlow (`ml-recommendations-api`)
- HTML + Node.js UI (`recommendations-ui`)
- PostgreSQL + Qdrant
- MinIO (object storage simulation)
- Docker Compose for local execution

## What is implemented

- User, book, and purchase management.
- Manual model training from current relational data.
- Recommendation ranking based on age, interests, and purchase behavior.
- Recommendation UI with visual cards, score percentage, and explanation signals.
- Qdrant vector indexing with automatic reindex when vector dimensions evolve.

## Run the full stack

```bash
docker compose up --build
```

To recreate data volumes and bootstrap from scratch:

```bash
docker compose down -v
docker compose up --build
```

## Local URLs

- UI: http://localhost:3000
- Recommendations API: http://localhost:8080
- ML Recommendations API: http://localhost:5000
- Qdrant: http://localhost:6333/dashboard
- MinIO API: http://localhost:9000
- MinIO Console: http://localhost:9001

## Initial endpoints

- API health: `GET /health`
- Books: `POST /api/v1/books`, `GET /api/v1/books`
- Users: `POST /api/v1/users`, `GET /api/v1/users`
- Purchases: `POST /api/v1/purchases`, `GET /api/v1/users/{userId}/purchases`
- Recommendations: `GET /api/v1/recommendations/{userId}`
- Manual training trigger: `POST /api/v1/recommendations/train`

## MinIO artifact layout

ML artifacts are stored in MinIO bucket `ml-models`:

- `book-recommendation.keras`: trained TensorFlow model.
- `encoding-context.json`: encoding context used by training and inference.

Runtime behavior:

- `POST /api/v1/recommendations/train` writes/overwrites both objects.
- `POST /recommend` in `ml-recommendations-api` loads model/context from MinIO.
- If MinIO is unavailable, ML service startup/training fails (mandatory dependency).

## C4 Model (GitHub-renderable)

### Container Diagram

```mermaid
flowchart LR
    U["👤 Person<br/>User"]

    subgraph UI_SYS
      direction TB
      UI["recommendations-ui<br/>Node.js + HTML + Bootstrap"]
    end

    subgraph API_SYS
      direction TB
      API["recommendations-api<br/>Go HTTP API"]
      PG[("PostgreSQL<br/>Users, Books, Purchases")]
      API -->|CRUD + orchestration| PG
    end

    subgraph ML_SYS
      direction TB
      ML["ml-recommendations-api<br/>FastAPI + TensorFlow"]
      QD[("Qdrant<br/>Embeddings Index")]
      MI[("MinIO<br/>Model & Context Artifacts")]
      ML -->|read/write vectors| QD
      ML -->|save/load model + context| MI
    end

    U -->|uses| UI
    UI -->|HTTP/JSON| API
    API -->|train/recommend orchestration| ML

    classDef person fill:#08427B,color:#FFFFFF,stroke:#052E56,stroke-width:1.5px;
    classDef container fill:#438DD5,color:#FFFFFF,stroke:#2E6295,stroke-width:1.5px;
    classDef database fill:#2E6295,color:#FFFFFF,stroke:#1B4F72,stroke-width:1.5px;

    class U person;
    class UI,API,ML container;
    class PG,QD,MI database;

    style UI_SYS fill:transparent,stroke:#9CA3AF,stroke-width:1.5px
    style API_SYS fill:transparent,stroke:#9CA3AF,stroke-width:1.5px
    style ML_SYS fill:transparent,stroke:#9CA3AF,stroke-width:1.5px
```

### Current bootstrap status

- Recommendations API persists books, users, and purchases in PostgreSQL.
- ML Recommendations API follows a clean-architecture-inspired structure.
- Initial user/book/purchase seed is loaded from `db/init/002_seed.sql`.
- Encoding context and model are persisted in MinIO bucket (`ml-models`).
- Training is manual through `POST /api/v1/recommendations/train`.
- Recommendation flow has two stages:
  1. Qdrant retrieves ANN candidates by vector similarity.
  2. TensorFlow re-ranks candidates with `model.predict` using age and purchase-enriched profile signals.
- On vector dimension mismatch, embeddings are reindexed (resaved) before retrying recommendation.
- UI renders recommendations as cards with score percentage.
- MinIO is required and used as the active model artifact backend.
- Healthchecks are chained across `postgres -> ml-recommendations-api -> recommendations-api -> recommendations-ui`.

> This bootstrap is a starting point for iterative evolution under SDD.

## Spec-Driven Development (SDD)

Specifications are organized by service and by feature in:

- `docs/sdd/README.md`
- `docs/sdd/services/`
- `docs/sdd/features/`

Use these documents as the official source for new requirements and changes.
`SPEC-BOOK-RECOMMENDATIONS.md` is now an executive index pointing to SDD.
