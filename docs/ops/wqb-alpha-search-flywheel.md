# WQB Alpha Search Flywheel v1

经济假设 → 字段地形 → 候选苗圃 → 预算闸门 → 异步派发 → 结果回收 → 判决沉积 → 下一轮搜索。

## Loop contract

- 每轮至少更新 batch ledger 和 terrain atlas。
- 每个 candidate 属于一个 family。
- 每个 family 有预算、状态、dominant_failure、next_action。
- cooldown 不等于停工；cooldown 时做非平台撞击工作。

## Budget defaults

| state | budget |
|---|---|
| explore | 6–12 candidates |
| promising | +12–24 candidates |
| repair | 2–3 repair batches |
| stop | no more simulation until new evidence |
