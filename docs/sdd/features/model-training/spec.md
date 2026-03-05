# Feature Spec: Model Training

## Solution

- Recommendations API loads `users` and `books` from PostgreSQL.
- Recommendations API also loads `purchases` and enriches users with purchase profile fields.
- Recommendations API calls `ml-recommendations-api:/train` with consolidated payload.
- ML Recommendations API:
  - builds encoding context;
  - builds synthetic training dataset `(user, book, label)` including age and purchase signals;
  - trains binary neural ranking model;
  - saves model and context.
- Training trigger is manual (UI/API), with no scheduled retraining in PoC.

## Return Contract

- `trained` (bool), `samples` (int), `input_dim` (int), `context` (object).

## Business Rules

- If no valid dataset exists, return `trained=false` with reason.
- Active model is updated only after successful manual training.

## Traceability

- Covers: `REQ-FTRN-001` to `REQ-FTRN-007`.
