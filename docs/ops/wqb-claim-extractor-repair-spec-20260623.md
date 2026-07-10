# WQB Claim Extractor Repair Spec — 2026-06-23

Status: repair specification only. No forum writes. No simulation.

## Defect

The previous deterministic extractor overproduced generic claims and did not ask topic-specific questions. It also made `read_status=read` too easy to confuse with understanding.

## Required output fields

`source_id`, `review_record_ref`, `extractor_type`, `claim_quality`, `body_chars_seen_in_memory`, `method_summary`, `mechanism_hypothesis`, `allowed_use`, `forbidden_use`, `gate_mapping`, `copy_risk`, `policy_risk`, `confidence_reason`, `expression_detail_removed=true`.

## Banned output fields

The extractor must not output fields named `expression`, `full_expression`, `fast_expression`, `regular`, `formula`, or `raw_template`.

If expression-like material appears, discard it, set `expression_detail_removed=true`, and retain only abstract method structure.

## Extractor types

- `turnover`: entry / hold / exit / information-arrival / decay boundary.
- `prod_corr`: originality, self/prod correlation, mechanism/data-family pivot, anti-bulk-mining.
- `operator_grammar`: operator roles, input types, complexity penalty, economic hypothesis link.
- `portfolio_sa`: pool quality, low-correlation method, combine metric decomposition, method-only.
- `research_paper`: research prior, field family, hypothesis, validation requirements.
- `workflow_automation`: anti-waste, recovery, logging, no autonomous bulk mining.
- `policy_boundary`: program boundary, allowed/disallowed behavior, no scheduler content.

## Repair priority queue

1. SCR-003 / forum:32226888249239 — gate standard extractor.
2. SCR-008 / forum:30927669645207 — turnover extractor; body evidence currently zero.
3. SCR-004 / forum:26054361848343 — template sanitizer; high copy risk.
4. SCR-005 / forum:37266277327767 — operator grammar extractor.
5. SCR-006 / forum:34696235484567 — economic meaning extractor.
6. SCR-016 / forum:35928620962839 — portfolio/combined score extractor.
7. SCR-020 / forum:37483774843671 — single-data extractor.
8. SCR-021 / forum:35377811169175 — SA method-only extractor.
9. SCR-013 / forum:32034293019671 — SA pool quality extractor.
10. SCR-034 / forum:15152019662487 — evergreen rulebook extractor.
11. SCR-002 / forum:19273239621399 — research-prior extractor.
12. SCR-044 / forum:29085671898775 — operator extractor.