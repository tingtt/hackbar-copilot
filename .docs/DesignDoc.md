# Design Doc: hackbar-copilot

## Objective

Provides support for bartender operations copilot.

## Goal, Non goal

### Goal

- **Mobile order system**
- **Recipe book**
- **Manuals**

### Non goal

- AI Chatting
- Disrespect for the efforts of bartenders
- To be the noise of the space provided by the bar.

## High Level Structure

```sh
.
├── .docs/      # Documents
│   └── DesignDoc.md
├── client/     # Client code (TypeScript)
├── cmd/        # Entrypoints (Go)
│   └── copilot/
├── internal/   # Internal packages (organizing with "Clean Architecture") (Go)
│   ├── infrastructure      # Scopes: infrastructure, security and persistence data
│   │   ├── api/
│   │   │   └── http/
│   │   └── datasource/
│   │       └── filesystem/
│   ├── interface-adapter/  # Scopes: adaption between infrastructure and usecase
│   │   └── handler/
│   │       ├── middleware/
│   │       └── graphql/
│   ├── usecase/            # Scopes: application bussiness rules
│   │   ├── copilot/
│   │   ├── order/
│   │   └── cashout/
│   └── domain/             # Scopes: enterprise bussiness rules
│       ├── menu/
│       ├── order/
│       ├── ordersummary/
│       ├── recipe/
│       └── stock/
└── test/
    ├── e2e/    # E2E test environments
    └── ci/     # CI test environments
```

## Open Issues

## References
