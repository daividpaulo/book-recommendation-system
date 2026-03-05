# Feature Requirements: Model Training

## Objective

Train recommendation model using registered users and books.

## Involved Services

- Technical owner: `ml-recommendations-api`
- Orchestrator: `recommendations-api`
- Integrated: `recommendations-ui` (manual trigger), `recommendations-data`

## Requirements

- `REQ-FTRN-001`: Recommendations API aggregates local data and calls `POST /train`.
- `REQ-FTRN-002`: ML API creates training data from age, interests, and purchase-behavior overlap logic.
- `REQ-FTRN-003`: Trained model is persisted for future inference.
- `REQ-FTRN-004`: Training response includes execution metadata.
- `REQ-FTRN-005`: Training remains manual in this phase (operator-triggered).
- `REQ-FTRN-006`: Artifact persistence strategy allows migration from local volume to object storage without contract changes.
- `REQ-FTRN-007`: Training payload must include purchase-derived profile fields (`purchase_count`, `purchased_book_ids`).

## Acceptance

- `POST /api/v1/recommendations/train` returns success with `trained=true` when there is valid training data.
- Model persists across `ml-recommendations-api` container restarts (volume).
