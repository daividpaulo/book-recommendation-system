# Feature Spec: Book Purchases

## Solution

- Recommendations API persists purchases in PostgreSQL (`purchases` table).
- Purchase contract:
  - `POST /api/v1/purchases`
  - `GET /api/v1/users/{userId}/purchases`
- During training orchestration, Recommendations API enriches each user with:
  - `purchase_count`
  - `purchased_book_ids`
- ML Recommendations API uses these fields to build training features and influence ranking.

## Business Rules

- Purchase references must point to existing `users` and `books`.
- `quantity` defaults to `1` when omitted or invalid.
- Purchases affect recommendations after training execution.

## Traceability

- Covers: `REQ-FPUR-001` to `REQ-FPUR-006`.
