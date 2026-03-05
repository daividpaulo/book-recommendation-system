# Feature Requirements: Recommendations

## Objective

Return ranked book recommendations by user affinity.

## Involved Services

- Technical ranking owner: `ml-recommendations-api`
- Inbound gateway and input validation: `recommendations-api`
- User experience: `recommendations-ui`
- Data support: `recommendations-data`

## Requirements

- `REQ-FREC-001`: Validate user before requesting recommendations.
- `REQ-FREC-002`: Retrieve candidates from vector index (Qdrant).
- `REQ-FREC-003`: Apply model-based re-ranking.
- `REQ-FREC-004`: Return list sorted by score.
- `REQ-FREC-005`: Return explicit error when model/context is not ready.
- `REQ-FREC-006`: Re-ranking must consider age and purchase-derived user profile signals.
- `REQ-FREC-007`: UI must present recommendation score as percentage.

## Acceptance

- Recommendation flow returns ranked books after successful training.
- Invalid flows return explicit errors without crashing services.
