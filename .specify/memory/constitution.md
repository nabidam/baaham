<!--
Sync Impact Report
- Version change: unknown -> 1.0.0
- Modified principles:
  - PRINCIPLE_1_NAME -> Coding Standards
  - PRINCIPLE_2_NAME -> Real-Time Principles
  - PRINCIPLE_3_NAME -> Security and Auth
  - PRINCIPLE_4_NAME -> Interoperability & Protocols (new title)
  - PRINCIPLE_5_NAME -> Simplicity
- Added sections:
  - Tech Stack and Architecture (section)
- Removed sections: none
- Templates requiring updates:
  - .specify/templates/plan-template.md: ⚠ pending (file not present)
  - .specify/templates/spec-template.md: ⚠ pending (file not present)
  - .specify/templates/tasks-template.md: ⚠ pending (file not present)
  - .specify/templates/commands/*.md: ⚠ pending (no commands directory found)
- Runtime/docs checked: README.md, docs/quickstart.md: ⚠ pending (files not present)
- Follow-up TODOs:
  - RATIFICATION_DATE is unknown — TODO(RATIFICATION_DATE): provide original adoption date
-->

# Baaham Constitution

<!-- Example: Spec Constitution, TaskFlow Constitution, etc. -->

## Core Principles

### Coding Standards
Go: Follow idiomatic Go project layout. Ensure strong error handling.

React: Use Functional Components and Hooks. Use a distinct folder structure for components, hooks, and context.

Styling: Use Tailwind utility classes exclusively. Avoid external CSS files where possible.

### Real-Time Principles
Single Source of Truth: The server maintains the master state for video playback timestamp and pause/play status.

Optimistic UI: Clients update immediately on user action but revert if the server contradicts.

Simplicity Constraint: Optimize logic for rooms with exactly 2 users to simplify synchronization, but write code that doesn't hard-crash if a 3rd joins.

### Security and Auth
Auth: JWT (JSON Web Tokens) for stateless authentication.

Roles: Strict separation. Admins manage data; Users consume it.

Sockets: Authenticate WebSocket connections via JWT query parameter or initial handshake.

### Interoperability & Protocols
Protocol: Strict JSON format for all WebSocket messages. Follow consistent message envelopes and version the envelope schema when changes occur.

### Simplicity
Prioritize clear, minimal solutions that are easy to reason about and maintain. Prefer small, well-documented components and avoid premature optimization.

## Tech Stack and Architecture
Backend: Go (Golang). Use standard library or a lightweight router like Chi.
Frontend: Vite with React (TypeScript) and Tailwind CSS.
Communication: REST API for authentication and management. WebSockets for all real-time features (Chat, Sync, Whiteboard).
Protocol: Strict JSON format for all WebSocket messages.

## Development Workflow
Code reviews, CI checks, and automated tests MUST validate compliance with this constitution for relevant changes. Major or breaking changes to principles or protocols require a documented migration plan and explicit approval by project maintainers.

## Governance
All contributions MUST demonstrate how they adhere to the constitution. Amendments follow semantic versioning rules: MAJOR for incompatible governance changes, MINOR for added principles or material expansions, PATCH for wording/clarity fixes.

Amendment procedure:
- Propose change in a PR with rationale, tests/migrations, and impact analysis.
- Two maintainer approvals required for MINOR/PATCH; three required for MAJOR changes.
- Final ratification date recorded in the constitution.

Compliance review expectations:
- PRs touching protocol, auth, or synchronization code MUST include compatibility tests and a roll-forward/rollback plan.

**Version**: 1.0.0 | **Ratified**: TODO(RATIFICATION_DATE): original adoption date needed | **Last Amended**: 2025-12-15
