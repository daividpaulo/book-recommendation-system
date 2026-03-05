# Feature Requirements: User Management

## Objective

Enable user creation and listing with personalization attributes.

## Involved Services

- Owner: `recommendations-api`
- Integrated: `ml-recommendations-api`, `recommendations-ui`, `recommendations-data`

## Requirements

- `REQ-FUSR-001`: Create user with `name`, `age`, `profession`, `interest_areas`.
- `REQ-FUSR-002`: Persist `interest_areas` as a string array.
- `REQ-FUSR-003`: List users for recommendation selection flow.
- `REQ-FUSR-004`: Publish user embedding trigger through current ML integration flow.

## Acceptance

- Created user appears in `GET /api/v1/users`.
- Interest data is returned in the same structure as submitted.
