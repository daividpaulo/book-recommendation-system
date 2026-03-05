# Feature Spec: User Management

## Solution

- Create endpoint writes user in PostgreSQL `users`.
- List endpoint returns user data for UI selection flow.
- Async integration with ML API through `POST /embed/users`.

## Data

- `User` entity: `id`, `name`, `age`, `profession`, `interest_areas`.

## Rules

- `interest_areas` must remain an array to support encoder and synthetic labels.

## Traceability

- Covers: `REQ-FUSR-001` to `REQ-FUSR-004`.
