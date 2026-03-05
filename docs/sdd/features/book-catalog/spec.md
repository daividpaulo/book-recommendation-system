# Feature Spec: Book Catalog

## Solution

- Create endpoint writes book to PostgreSQL.
- List endpoint queries the `books` table.
- Async integration with ML API through `POST /embed/books`.

## Data

- `Book` entity: `id`, `title`, `author`, `category`, `subject`, `area`, `description`.

## Risks and Mitigation

- ML dependency failures must not compromise transactional book writes.
- Keep minimum payload validation to prevent invalid records.

## Traceability

- Covers: `REQ-FBK-001` to `REQ-FBK-004`.
