# ML Recommendations API Spec

## Responsibilities

- Build user and book vectors.
- Generate synthetic supervised training data.
- Train and persist neural ranking model.
- Re-rank vector candidates returned from Qdrant.

## Technical Design

- Encoder converts entities into numeric vectors using a shared context, including age and purchase intensity.
- Training uses binary classification with scores in `[0, 1]`.
- Recommendation pipeline combines Qdrant candidate retrieval with model re-ranking.

## Target Architecture (Clean Architecture)

```text
ml-recommendations-api/
  src/
    domain/
      entities/
    application/
      train_model.py
      recommend.py
      embed_user.py
      embed_book.py
    interfaces/
      encoders/base.py
      model_repository/base.py
      vector_store/base.py
    infrastructure/
      encoders/tensorflow_encoder.py
      model_repository/tensorflow_model.py
      vector_store/qdrant_store.py
  api/
    main.py
```

## Dependency Rules

- `domain` <- `application` <- `api` and `infrastructure`.
- `application` imports only contracts from `interfaces`.
- `api/main.py` performs dependency composition.

## Core Contracts

- `POST /train`: returns training status and metadata (`samples`, `input_dim`).
- `POST /recommend`: accepts `user_id` and returns ordered ranking.
- `POST /embed/*`: updates vector index for users/books.

## Training Strategy (PoC)

- Manual trigger via `POST /train` from API/UI.
- No automatic retraining in this phase.
- Future evolution can add scheduler/event-driven triggers without contract changes.

## Persistence

- Model artifact: `.keras` object stored in MinIO bucket.
- Encoding context: `encoding-context.json` object stored in MinIO bucket.
- User/book embeddings: Qdrant collections (`users`, `books`).
- Model artifact is not stored in Qdrant.
- MinIO is mandatory for model/context persistence in runtime flow.

## Acceptance Criteria

- `REQ-ML-*`, `ARC-ML-*`, and `NFR-ML-*` are covered.
- Layer dependency rules are respected.
- Persistence responsibilities are clearly separated by data type.
- Minimum tests:
  - unit tests in `application`;
  - integration tests in `api` and `infrastructure`.
- End-to-end flow works: create entities/purchases -> manual train -> recommend.
