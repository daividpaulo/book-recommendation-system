# Service x Feature Mapping

## Objective

Define explicit ownership for each feature to support SDD planning, implementation, and testing.

## Mapping

| Service | Owned Features | Integrated Features |
|---|---|---|
| `recommendations-api` | `book-catalog`, `user-management`, `book-purchases`, `model-training` (orchestration), `recommendations` (HTTP gateway) | `recommendations` (ML ranking dependency) |
| `ml-recommendations-api` | `model-training` (training), `recommendations` (inference/re-ranking) | `book-catalog`, `user-management`, `book-purchases` (profile enrichment) |
| `recommendations-ui` | `recommendations` (home), `user-management` (create), `book-catalog` (create), `book-purchases` (register purchase) | `model-training` (manual trigger action) |
| `recommendations-data` | relational/vector/object data platform support | all features |

## Working Rule

For any change:

1. update feature requirements/spec;
2. update owning service requirements/spec;
3. update cross-service validation plan when integrations are affected.
