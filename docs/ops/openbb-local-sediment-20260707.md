# OpenBB local repository sediment · 2026-07-07

## Source and status

- local path: `/Users/jiajia/Documents/GitHub/OpenBB`
- recorded remote: `TaowuZhang/OpenBB`; authenticated GitHub API returns HTTP 404
- local state: clean `main` at `c53cf94`, with no configured upstream

## Durable shape

The unique part is a small FastAPI adapter, not the vendored Python environment:

- `/api/equity/*` routes wrap OpenBB for US equity historical data and quotes;
- `/api/a-share/*` routes wrap AkShare for A-share historical data and quotes;
- `/api/health` and the root endpoint expose basic service health and discovery;
- service modules separate US-date validation/OpenBB calls from China-date validation/AkShare calls.

This is a useful boundary pattern: one HTTP surface, provider-specific services, explicit market-specific validation, and normalized responses.

## Excluded material

- 15,455 tracked `.venv` files;
- Python bytecode caches;
- the tracked Python installer package;
- the full pinned environment;
- any runtime data or credentials.

These are rebuildable or unsafe migration material and are not canonical knowledge.

## Disposition

The architecture and asset boundary are preserved here. The eight source Python files remain unique local implementation and were not copied. Before local checkout removal, either migrate those source files into an explicitly chosen maintained project or mark the implementation discardable. The missing/renamed remote must also be resolved.

## Eight-file source manifest

Both repository commits were authored by `TaowuZhang`; these files are a personal prototype rather than copied OpenBB upstream source.

| source file | durable behavior |
|---|---|
| `app/__init__.py` | package identity for the REST adapter |
| `app/main.py` | FastAPI assembly, development CORS, root discovery, health/version endpoint |
| `app/routers/__init__.py` | router package boundary |
| `app/routers/equity.py` | US equity historical/quote HTTP routes and 400/503 error boundary |
| `app/routers/a_share.py` | A-share historical/quote HTTP routes and 400/503 error boundary |
| `app/services/__init__.py` | provider-service package boundary |
| `app/services/openbb_service.py` | US symbol/date validation, OpenBB+yfinance calls, normalized historical/quote results |
| `app/services/akshare_service.py` | six-digit A-share/date validation, AkShare calls, normalized Chinese-market fields |

The reusable asset is the provider-neutral boundary, validation split, and normalized result contract—not the exact dependency snapshot. This manifest and the durable-shape section preserve that design. Exact source remains in the clean local checkout until the user chooses either promotion into maintained code or explicit discard; no source deletion occurred.

## Round 5 retention decision

| source file | role | migrate? | reason | target |
|---|---|---:|---|---|
| `app/__init__.py` | package marker | summary only | no reusable behavior beyond package identity | this manifest |
| `app/main.py` | FastAPI/CORS/health assembly | summary only | useful boundary, but prototype-specific implementation | durable-shape section |
| `app/routers/__init__.py` | router package marker | summary only | no standalone reusable behavior | this manifest |
| `app/routers/equity.py` | US equity HTTP adapter | summary only | thin provider route; exact code is not a maintained contract | eight-file manifest |
| `app/routers/a_share.py` | A-share HTTP adapter | summary only | thin provider route; exact code is not a maintained contract | eight-file manifest |
| `app/services/__init__.py` | service package marker | summary only | no standalone reusable behavior | this manifest |
| `app/services/openbb_service.py` | validation and OpenBB/yfinance normalization | summary only | provider versions and API surface can drift | durable-shape + manifest |
| `app/services/akshare_service.py` | A-share validation and AkShare normalization | summary only | hard-coded dates/provider fields need redesign before reuse | durable-shape + manifest |

Codex-actionable retention is complete: architecture, file identity, and reusable decisions are preserved without copying unstable code. Decision `D-openbb-source-retention` remains for the user: keep exact source as historical local evidence, promote it into a maintained project, or explicitly discard it before removing the checkout.
