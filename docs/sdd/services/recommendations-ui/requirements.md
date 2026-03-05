# Recommendations UI Requirements

## Scope

Web interface for entity creation and recommendation consumption.

## Owned Features

- `recommendations` (home screen)
- `user-management` (user creation)
- `book-catalog` (book creation)
- `book-purchases` (purchase registration)
- `model-training` (manual trigger action)

## Functional Requirements

- `REQ-FE-001`: On app start, show user selection screen (no login).
- `REQ-FE-002`: Load user list automatically at startup.
- `REQ-FE-003`: Allow selecting a user to enter the recommendations home.
- `REQ-FE-004`: Home screen must focus on recommendation actions/results only.
- `REQ-FE-005`: Allow active user switch from recommendation screen without full app reload.
- `REQ-FE-006`: Navigation menu must provide:
  - Home (Recommendations)
  - Create User
  - Create Book
- `REQ-FE-007`: Allow user creation.
- `REQ-FE-008`: Allow book creation.
- `REQ-FE-009`: Allow manual training trigger.
- `REQ-FE-010`: Allow recommendation retrieval for selected user.
- `REQ-FE-011`: Display loading/success/error/empty states for main actions.
- `REQ-FE-012`: Provide a dedicated "Buy Book" flow that registers purchases.
- `REQ-FE-013`: Render recommendations as visual cards instead of raw JSON.
- `REQ-FE-014`: Show recommendation score as percentage in each card.

## UI/UX Requirements

- `UI-FE-001`: Use Bootstrap for core layout and styling.
- `UI-FE-002`: Responsive layout for desktop/tablet.
- `UI-FE-003`: Menu remains visible and consistent across screens.
- `UI-FE-004`: User selector is highlighted on entry and accessible from recommendations screen.

## Non-Functional Requirements

- `NFR-FE-001`: Run on `PORT=3000` using Node.js server.
- `NFR-FE-002`: Consume API from configurable `API_URL`.
- `NFR-FE-003`: Show integration errors without breaking the page.
- `NFR-FE-004`: Keep authentication out of scope for this MVP phase.
