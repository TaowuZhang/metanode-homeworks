# WQB premium forum claims review — 2026-06-23

## What happened

The local premium-source extraction completed successfully and was pushed as commit `1b208ed`.

Local run result:

```text
sources_attempted = 12
read_status_counts.read = 12
output = runs/wqb-community/premium-source-claims.json
safety scan = no matches
integration = ok
new_claims_added = 11
usable_claim_count = 11
blocked_or_deferred_claim_count = 1
```

This confirms the source-layer workflow now works end-to-end:

```text
premium queue
→ local authenticated CNHKMCP reads
→ claim-only artifact
→ safety scan
→ integration into claim registry / starmap / Batch F packet
→ GitHub push
```

No simulation was dispatched.

## Sources read

### P0 high-vote sources

1. `forum:15238206176279` — **Term #14 : Price Momentum**
   - vote_count: 91
   - status: full_read_claim_draft_needs_review
   - effect: price momentum must become source-backed scout priors and operator grammar, not blind price-volume reuse.

2. `forum:41254129745815` — **Antigravity-cli migration tutorial**
   - vote_count: 55
   - status: full_read_claim_draft_needs_review
   - effect: workflow/tooling source only; it may improve local runner/process design but cannot authorize autonomous bulk mining.

### P0 official evergreen community sources

3. `official-community:wqb-community-reduce-correlation`
   - status: full_read_claim_draft_needs_review
   - effect: correlation/prod_corr is reinforced as mechanism/data-family/robustness/originality problem, not just parameter repair.

4. `official-community:wqb-community-avoid-overfitting`
   - status: full_read_claim_draft_needs_review
   - effect: overfitting and stability require robustness/regime diagnostics; do not blindly expand parameters after 2Y/sub-universe failure.

5. `official-community:wqb-community-evaluate-new-dataset`
   - status: full_read_claim_draft_needs_review
   - effect: scout slots require dataset/field evaluation checklist before simulation.

6. `official-community:wqb-community-weight-coverage`
   - status: metadata_only_or_empty_body
   - effect: still not strong enough. Keep P0 queued for retry or replace with official alpha-submission/data claims until full body is available.

7. `official-community:wqb-community-reduce-turnover-trade-when`
   - status: full_read_claim_draft_needs_review
   - effect: turnover repair should be structural entry/exit design via trade_when-style grammar, not just larger decay.

### P1 high-value sources

8. `forum:41284360565399` — automated alpha mining workflow
   - status: full_read_claim_draft_needs_review
   - effect: process/workflow insights only; anti-bulk-mining and per-batch authorization stay mandatory.

9. `forum:41225167988503` — forum posting/comment management notice
   - status: low-confidence generic draft
   - effect: needs better manual/stronger summary if it is to affect workflow norms.

10. `forum:15770966939671` — Research Paper 12: Stock Recommendations from Stochastic Discounted Cash Flows
   - status: low-confidence generic draft
   - effect: not yet useful as a source-backed alpha design prior; needs targeted paper claim extraction.

11. `forum:41285657542295` — Super Alpha Combo expressions
   - status: full_read_claim_draft_needs_review
   - effect: method-only correlation/originality insight; expression copying remains forbidden.

12. `forum:41284510093591` — Fourier/Python Alpha signal-processing post
   - status: low-confidence generic draft
   - effect: needs targeted Fourier/frequency-domain method extraction before affecting candidate grammar.

## Quality assessment

The extraction was safe and useful, but still deterministic and conservative.

Strong enough to immediately affect scheduler:

- reduce correlation / prod_corr evergreen
- avoid overfitting evergreen
- evaluate new dataset evergreen
- trade_when / turnover evergreen
- price momentum high-vote term as scout prior
- automation/workflow high-vote post as process-only claim
- Super Alpha combo as method-only correlation/originality claim

Not strong enough yet:

- weight coverage evergreen returned empty body and must be retried or supported by official docs.
- forum governance post produced generic claim only.
- DCF research paper produced generic claim only.
- Fourier post produced generic claim only.

## Design consequences

### Batch F remains blocked, but for a better reason

The old Batch F packet was already marked stale. After premium forum intake, it should not simply be unblocked. It must be redesigned around the following rules:

1. **Correlation first**
   - Reduce local micro-repair budget.
   - Prefer mechanism/data-family/robustness/originality pivots.

2. **Dataset scout discipline**
   - Add field profiling / coverage / frequency / bounds / distribution checks before expensive scout simulations.

3. **Overfitting/stability discipline**
   - Add robustness/regime diagnostics before exploiting any single lucky parent.

4. **Turnover grammar**
   - Turnover repair requires entry/hold/exit structure, not decay-only mutation.

5. **Expression-copy firewall**
   - Super Alpha combo source is method-only. No expression or formula copying.

6. **Workflow-source separation**
   - Tooling/automation posts may improve process but never override WQB policy: no autonomous bulk mining, no submit, per-batch authorization required.

## Immediate next actions

1. Fix the extractor/integrator weakness where failure_modes stayed empty for most forum claims.
2. Retry or replace the weak `weight coverage` P0 source.
3. Add a targeted extraction pass for the three generic P1 sources if they remain relevant:
   - forum governance
   - DCF research paper
   - Fourier/Python Alpha
4. Redesign Batch F as a source-backed proposal, not as the previous stale packet.
5. Run duplicate/pre-dispatch validators before any future simulation.

## Current state

```text
premium_forum_claims_integrated_reviewed_no_simulation
final_click_candidate_found=false
Batch F = stale / blocked / redesign required
```
