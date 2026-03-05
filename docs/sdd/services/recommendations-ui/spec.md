# Recommendations UI Spec

## Responsibilities

- Serve the web user interface.
- Call `recommendations-api` endpoints.
- Show results and recoverable error messages to users.

## Covered Features

- `recommendations`
- `user-management`
- `book-catalog`
- `book-purchases`
- `model-training` (manual action)

## Interface Design (Bootstrap)

- Bootstrap-based layout (navbar, grid, forms, alerts, buttons, cards).
- Fixed navigation:
  - Home (Recommendations)
  - Create User
  - Create Book
  - Buy Book
- Reusable components for loading/error/empty states.

## UI Flows

1. **Entry (User Selection)**
   - App loads users from `GET /api/v1/users`.
   - User selects active context (no login credentials).

2. **Home (Recommendations)**
   - Home shows only recommendation controls/results:
     - active user;
     - change user action;
     - train model action;
    - recommendation cards with score percentage.
   - Calls:
     - `POST /api/v1/recommendations/train`
     - `GET /api/v1/recommendations/{userId}`

3. **Create User**
   - Dedicated form -> `POST /api/v1/users`
   - Refresh local users list after success.

4. **Create Book**
   - Dedicated form -> `POST /api/v1/books`

5. **Buy Book**
   - Dedicated form -> `POST /api/v1/purchases`
   - Purchases become profile signal after manual retraining.

## UI State Model

- `activeUserId`
- `users`
- `recommendations`
- `status` (`idle | loading | success | error`)

## Acceptance Criteria

- `REQ-FE-*`, `UI-FE-*`, and `NFR-FE-*` are covered.
- First screen is user selection (no authentication).
- Home is recommendation-focused.
- Navigation clearly separates create-user and create-book flows.
