# Feature Requirements: Book Catalog

## Objective

Enable book creation and listing to feed training and recommendations.

## Involved Services

- Owner: `recommendations-api`
- Integrated: `ml-recommendations-api`, `recommendations-ui`, `recommendations-data`

## Requirements

- `REQ-FBK-001`: Create book with minimum fields (`title`, `id` or auto-generated id).
- `REQ-FBK-002`: Persist ML-relevant attributes (`area`, `category`, `subject`).
- `REQ-FBK-003`: List books ordered by creation time.
- `REQ-FBK-004`: Publish book embedding trigger through current ML integration flow.

## Acceptance

- Created book appears in `GET /api/v1/books`.
- Book creation flow does not fail if async embedding call fails downstream.
