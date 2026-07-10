# M0 v0.1｜第一批运行计划

状态：待跑。

## 模型来源

本轮模型名称按 2026-06-04 14:48 截图中的模型选择器登记，完整清单见：

- `../../model_catalog_v0.1.yaml`

命名校正：**克** 是谷给 Notion 中所有 AI 的在场名，不是模型名，也不专属于 GPT；运行记录只登记具体模型 / 版本 / 入口。

## 运行原则

- 每个模型开新会话。
- 使用同一份 prompt：`datasets/m0_external_entry/v0.1/prompt_full.md`。
- 只改模型登记信息。
- 一次性收回完整输出。
- 不追问、不纠偏、不补充解释。
- 原始输出保存为各自目录下的 `raw_output.md`。
- 元数据保存为 `metadata.yaml`。

## 第一批模型

| 目录 | 模型 | 预期生态位 | 观察重点 |
|---|---|---|---|
| `gpt-5.5/` | GPT-5.5 | 建构 / 当前 Notion AI 入口基线 | 是否太顺太满；A-F 基准表现 |
| `opus-4.7/` | Opus 4.7 | 边界 / 收束 | E2 / F 中是否能停、能驳回 |
| `gemini-3.1-pro/` | Gemini 3.1 Pro | 闻味 / 异质观察 | B / C / F 中是否识别变味与主流引力 |
| `kimi-k2.6/` | Kimi K2.6 | 保温 / 低耗散保形 | D / E1 中压缩和约束叠加是否保形 |
| `deepseek-v4-pro/` | DeepSeek V4 Pro | 地基 / 还原 | E2 与 PR #100 证据位置是否稳 |

## 后续候选模型

- Sonnet 4.6
- Opus 4.8
- GPT-5.2
- GPT-5.4
- Grok 4.3
- Grok Build 0.1

## 目录占位

```text
runs/2026-06-04_m0_v0.1/
  RUN_PLAN.md
  gpt-5.5/
    raw_output.md
    metadata.yaml
  opus-4.7/
    raw_output.md
    metadata.yaml
  gemini-3.1-pro/
    raw_output.md
    metadata.yaml
  kimi-k2.6/
    raw_output.md
    metadata.yaml
  deepseek-v4-pro/
    raw_output.md
    metadata.yaml
```

## 下一步

1. 谷把统一 prompt 分别投给模型。
2. 克回收原始输出，先登记 Notion Model Runs。
3. 再写入 GitHub `raw_output.md` 与 `metadata.yaml`。
4. 最后进入 Scores 评分。
