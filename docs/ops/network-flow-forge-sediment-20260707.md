# Network-Flow-Forge sediment · 2026-07-07

## Source and status

- source: `TaowuZhang/Network-Flow-Forge`
- GitHub state: private, active, `archived=false`, last pushed 2025-12-01
- local checkout: none found under `/Users/jiajia` to depth 6

## Durable shape

This repository is an integration blueprint across three upstream projects:

- `me-release`: Vue-based topology modeling and visualization;
- `flowmanager-master`: OpenFlow control and execution;
- `nautobot-develop`: network source of truth.

Its original value is the proposed BFF/contract layer: normalized `/api/*` models, decimal DPID rules, text-response normalization, read-after-write confirmation, WebSocket event translation/throttling, contract tests, and requirements traceability. The BFF itself is explicitly not implemented.

The remote tree contains 85 files: 60 Markdown documents and 8 JavaScript overlay files. The three large upstream codebases are submodules and remain vendor/upstream material, not migration candidates.

## Disposition

This page preserves the project identity, integration boundary, and key contract decisions. It does not replace the full specification corpus, diagrams, requirement matrix, or overlay source. Archive remains blocked until those original assets are either promoted into a maintained project, explicitly marked historical-only, or accepted as preserved solely inside an archived repository. Submodule ownership must remain attributed to the upstream repositories.

## Specification and overlay index

Core original specifications:

- `API-and-Data-Spec.md`: unified public contract and FlowManager mapping;
- `BUSINESS-OVERVIEW.md` / `BUSINESS-MODEL-CANVAS.md`: system and product rationale;
- `PRD-remaster-me.md`: front-end intent, state, and API requirements;
- `FUTURE-ROADMAP.md`: staged delivery and release gates;
- `TECH-STACK-DECISIONS.md`: runtime, BFF, security, and observability boundaries;
- `requirements-traceability.csv`: requirement-to-API-to-acceptance mapping;
- `docs/mermaid/`: C4, sequence, data-flow, overlay, and state diagrams.

Overlay assets are local integration work, not upstream submodule code:

- `overlays/bridge/`: topology bridge and IndexedDB boundary;
- `overlays/me-release/adapters/`: FlowManager API/adapter layer;
- `overlays/me-release/components/` and `overlays/ui/`: connector UI prototypes;
- `overlays/me-release/vite-plugin-flowmanager.js`, Vite overlay config, patch, and apply script: build-time integration;
- `overlays/me-release/test-data/`: topology fixtures and local test server;
- `overlays/docs/` plus validation guides: implementation/test plan and operator instructions.

This index is the minimum canonical doorplate. The full 60-document corpus and overlay code remain historical implementation evidence in the repository; they are not silently promoted into maintained `worang` code.

## Round 5 asset preservation policy

| asset group | preserve as | target | blocker |
|---|---|---|---|
| API/PRD/business/roadmap/stack specs | durable summary and exact path index | this sediment | none for summary; full text stays remote |
| Mermaid architecture and sequences | diagram-family index | `docs/mermaid/` path entry above | full diagrams remain archive evidence |
| requirements traceability | named historical artifact | `requirements-traceability.csv` path entry above | promote only if project resumes |
| bridge/adapters/components | overlay purpose index | overlay index above | code depends on upstream projects |
| build patches/plugin/config | historical integration evidence | overlay index above | version-coupled; do not copy blindly |
| fixtures/test server | historical validation evidence | overlay index above | useful only with the original integration |
| three upstream submodules | vendor/reference links | original `.gitmodules` attribution | never copy as personal source |

Codex-actionable preservation is complete at summary/index level. Decision `D-network-preservation` remains: accept the full repository as archive-only historical evidence, or promote selected complete specs/overlays into a maintained project before archive.
