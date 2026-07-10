# WQB Platform Rulebook v1

状态：执行脚手架 / 平台边界  
生成时间：2026-06-21T23:36:28+08:00  
来源：WQB Alpha 生产脚手架批准执行

## 0. 硬边界

- 不自动 submit；当前工具不暴露 submit_alpha，submitEnabled=false。
- 不绕过 captcha、auth、rate-limit、Retry-After 或并发限制。
- 不沉 credentials、cookies、private alpha full text。
- 不做 autonomous bulk mining。
- simulation 必须逐批显式授权；每批记录 confirmSimulation=true。
- create_multi_simulation 安全 batch size 默认 6，平台稳定才可升到 8，不触碰约 10 的硬上限。

## 1. 平台事件处理

| 事件 | 处理 |
|---|---|
| 429 / Retry-After | 立即停止平台撞击，记录 cooldown-state，转 ledger / proposal / 文档工作 |
| CAPTCHA | 停止平台动作，不绕过，等待用户人工处理 |
| auth false | 停止平台动作，不反复重试，标记 AUTH cooldown |
| CONCURRENT_SIMULATION_LIMIT_EXCEEDED | 降 batch size 到 4–6，延长 polling 间隔 |
| simulation pending 过久 | 不开新批次，先回收可读结果并补 ledger |
| submit 需求 | 标 blocked_by_policy，只能停在 UI Submit 前 |

## 2. 批次纪律

每个 batch 必须有：

- batch_id
- purpose：explore / exploit / repair / diagnostic / legacy_import
- family_id
- authorized_by_user
- confirmSimulation
- max_batch_size
- candidates[]
- results[]
- cooldown_events[]

不得出现“单票游离 Alpha”：每个 candidate 必须归属 batch、family、cell。

## 3. final-click 命名禁区

只有 official submission check、self/prod/power-pool correlation、data diversity、regular submission、matches pyramid 等 blocking gates 全部 pass 或 non-blocking，才可进入 final_click_candidate。任何 blocking gate fail 时，状态只能是 weak_signal / promising / breed / repair / gate_check / discard / cooldown / blocked_by_policy。
