# SDD Workspace

This directory defines the project baseline using **Spec-Driven Development (SDD)**.

## Goal

Every meaningful change must include:

1. business context;
2. testable requirements;
3. technical specification;
4. traceability between requirements, implementation, and validation.

## Structure

```text
docs/sdd/
  README.md
  service-feature-mapping.md
  services/
    recommendations-api/
      requirements.md
      spec.md
    ml-recommendations-api/
      requirements.md
      spec.md
    recommendations-ui/
      requirements.md
      spec.md
    recommendations-data/
      requirements.md
      spec.md
  features/
    book-catalog/
      requirements.md
      spec.md
    user-management/
      requirements.md
      spec.md
    model-training/
      requirements.md
      spec.md
    recommendations/
      requirements.md
      spec.md
```

## Service Ownership

Use `service-feature-mapping.md` to identify the owning service and integrated services for each feature.

## Recommended Workflow

1. Update `requirements.md` (service + feature scope).
2. Update `spec.md` with design, contracts, and rollout notes.
3. Implement code referencing requirement IDs.
4. Validate acceptance criteria with technical and functional tests.

## SDD Conventions

- Requirements use stable IDs (`REQ-*`, `ARC-*`, `NFR-*`, `UI-*`).
- Specs document explicit architecture and design decisions.
- Changes must preserve end-to-end traceability.
- New features must not be implemented without approved requirements/specs.
