# WQB Authorized Batch Dispatch Runbook v1

状态：逐批授权后的执行手册  
生成时间：2026-06-22T10:20:00+08:00

## 0. Hard gates

任何 simulation 派发前必须同时满足：

- 用户本轮明确授权某一个 batch，例如：`授权 Batch B，confirmSimulation=true`。
- batch 存在于 `runs/wqb-current/next-batch-proposal.json`。
- batch candidate 数量不超过 `safe_batch_size`，默认 6。
- 当前 `cooldown-state.json` 没有对应 platform blocker。
- `scripts/wqb-validate-proposal.mjs` 对 proposal 返回 `ok=true`。
- 不包含 submit 动作；`submitEnabled=false`。
- 不沉 credentials / cookies / private alpha full text。

如果授权话术只说“继续”，不能视为 simulation 授权。

## 1. Dispatch sequence

1. Read `runs/wqb-current/run-state.json` and `runs/wqb-current/cooldown-state.json`.
2. Run `node scripts/wqb-validate-proposal.mjs runs/wqb-current/next-batch-proposal.json`.
3. Confirm selected batch id and `confirmSimulation=true` are explicit in the current user message.
4. Materialize FASTEXPR only in transient execution context from:
   - `field_ref`
   - `operator_family`
   - `horizon`
   - `settings`
5. Call WQB `create_multi_simulation` once, with at most 6 candidates.
6. Poll conservatively; respect `Retry-After`.
7. Write returned simulation ids/urls into a new `runs/wqb-batches/BATCH-*.json` result file.
8. Do not start another batch without another explicit authorization.

## 2. Materialization rules

| operator_family | expression template |
|---|---|
| `rank_ts_zscore` | `rank(ts_zscore(field, zscore))` |
| `rank_ts_delta_ts_zscore` | `rank(ts_delta(ts_zscore(field, zscore), delta))` |
| `rank_ts_rank` | `rank(ts_rank(field, rank))` |
| `inverse_rank_ts_zscore` | `-rank(ts_zscore(field, zscore))` |
| `vector_avg_rank_ts_zscore` | `rank(ts_zscore(vec_avg(field), zscore))` |

Templates are implementation mechanics, not private Alpha ledger storage. Long-term Git records should keep `expression_ref`, not full private expression text.

## 3. Proposal validation rules

The validator enforces:

- proposal is still `proposal-only` and has not dispatched simulation;
- `submitEnabled=false` and `autonomousBulkMining=false`;
- each batch requires explicit confirmation;
- each batch stays within `safe_batch_size`;
- no `expression` or `regular` is stored in Git proposal candidates;
- each candidate has `expression_ref`, `field_ref`, `operator_family`, `horizon`, and `settings`;
- behavioral VECTOR candidates use a `vector_*` adapter;
- helper/control fields do not enter the first behavioral exploration batch.

## 4. Cooldown rules

- `429`: stop platform calls, write cooldown event, switch to ledger/proposal/docs.
- `CAPTCHA`: stop platform calls; no bypass.
- `AUTH`: stop platform calls; user restores auth.
- `CONCURRENT_SIMULATION_LIMIT_EXCEEDED`: stop or reduce next authorized batch to 4–6; do not immediately retry.
- correlation reads returning `data=null`: do not upgrade state; keep as unrefreshed.

## 5. Post-run ledger minimum

Each dispatched candidate must record:

- `batch_id`
- `candidate_id`
- `simulation_id` or `simulation_url`
- `field_ref`
- `operator_family`
- `horizon`
- `settings`
- `expression_ref`
- `status`
- `cooldown_events[]` if any

Only after completed results are read should gate status update `terrain-atlas.jsonl` and `family-summary.json`.

## 6. Final-click rule

A result can only become `final_click_candidate` if official submission check and correlation/platform gates all pass or are non-blocking. Otherwise it remains `repair`, `breed`, `gate_check`, `cooldown`, or `discard`.
