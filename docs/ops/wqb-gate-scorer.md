# WQB Gate Scorer v1

状态：排序器，不是提交许可  
生成时间：2026-06-21T23:36:28+08:00

## Formula

score = 0.25 signal_strength + 0.20 stability + 0.15 tradability + 0.15 originality + 0.15 platform_fit + 0.10 debuggability - blocker_penalty

## Components

| component | reads |
|---|---|
| signal_strength | Sharpe, Fitness, return, drawdown, lift over family baseline |
| stability | LOW_2Y_SHARPE, LOW_SUB_UNIVERSE_SHARPE, horizon consistency |
| tradability | turnover range, concentration, truncation sanity |
| originality | self/prod/power-pool correlation, data diversity |
| platform_fit | submission check, regular submission, matches pyramid |
| debuggability | simple operator_count, clear mutation path, no type fragility |

## Hard rule

if any blocking gate fails: state != final_click_candidate

Score only ranks next actions. It never overrides official WQB gates.
