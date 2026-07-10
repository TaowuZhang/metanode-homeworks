# Runs｜原始输出归档

每个模型一次运行应保留原始输出，不要只保存评分。

命名校正：**克** 是谷给 Notion 中所有 AI 的在场名，不是模型名，也不专属于 GPT；运行目录与 metadata 必须登记具体模型 / 版本 / 入口。

建议目录：

```text
runs/
  2026-06-04_m0_v0.1/
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

metadata.yaml 最小字段：

```yaml
model:
model_version_or_entry:
date:
prompt_version: v0.1
mother_topic: M0｜第 0 期外显入口
one_shot: true
context_condition: 新会话 / 有上下文 / 不确定
notion_run_page:
notes:
```
