# Feature Spec: Recommendations

## Solution

- Recommendations API receives `userId` and validates user existence.
- Recommendations API calls ML API with `POST /recommend`.
- ML Recommendations API:
  - loads user from runtime context;
  - retrieves candidates from Qdrant;
  - computes scores via `model.predict` using age and purchase-enriched profile vectors;
  - sorts and returns ranking.

## Dependencies

- Trained model and loaded context.
- Indexed user/book vectors.

## Expected Errors

- Unknown user.
- Model not trained.
- Vector dependency failure.

## Traceability

- Covers: `REQ-FREC-001` to `REQ-FREC-007`.
