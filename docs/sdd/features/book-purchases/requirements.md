# Feature Requirements: Book Purchases

## Objective

Register book purchases and use this signal to enrich user profile quality for recommendations.

## Involved Services

- Functional owner: `recommendations-api`
- Technical model owner: `ml-recommendations-api`
- User journey owner: `recommendations-ui`
- Data support: `recommendations-data`

## Requirements

- `REQ-FPUR-001`: Create purchase via `POST /api/v1/purchases` with `user_id`, `book_id`, and `quantity`.
- `REQ-FPUR-002`: Validate user and book existence before persisting purchase.
- `REQ-FPUR-003`: List user purchases via `GET /api/v1/users/{userId}/purchases`.
- `REQ-FPUR-004`: Include purchase history in training payload sent to ML API.
- `REQ-FPUR-005`: Consider purchase history as recommendation signal after manual retraining.
- `REQ-FPUR-006`: UI must provide a dedicated flow to register purchases.

## Acceptance

- Purchase creation succeeds for valid references and fails with explicit errors for missing user/book.
- Manual training consumes purchase history and recommendations change according to profile and purchase behavior.
