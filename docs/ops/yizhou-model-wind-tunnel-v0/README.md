# 易舟 · 模型风洞 v0

本目录是 Notion「易舟 · 模型风洞 v0」的 GitHub 厚源仓入口。

它不是通用模型排行榜，也不是一次性测评报告；它负责保存可复现材料：冻结 prompt、测试集快照、不变量、评分规则、原始输出归档位置与首轮读数报告。

## v0 范围

- 母题：M0｜第 0 期外显入口
- 探针：A / B / C / D / E1 / E2 / F
- 评分：0 / 1 / 2 粗评
- 读数：V_intent / V_constrained / V_articulation / V_noise → 刚度 / 强度 / 韧度暂读
- 状态：测试集冻结；等待 3–5 个模型原始输出

## 目录

```text
docs/ops/yizhou-model-wind-tunnel-v0/
  README.md
  datasets/
    m0_external_entry/
      v0.1/
        invariant.md
        source_snapshot.md
        test_cases.yaml
        prompt_full.md
  scoring/
    rubric_v0.1.md
    noise_tags_v0.1.yaml
  runs/
    README.md
  reports/
    m0_v0.1_initial_readout.md
```

## Notion / GitHub 分居

- Notion：现场控制台、测试集可读表、模型输出摘要、评分读数、派船判断。
- GitHub：冻结 prompt、测试集 YAML、母题快照、原始输出、评分规则、报告与未来脚本。

## 硬边界

- 不做模型排行榜。
- 不把一次输出写成稳定规则。
- 不用高承重、身体、关系材料做 v0 红队。
- 不临场给不同模型补不同解释。
- 不在没有原始输出的情况下读坐标。

## 当前下一步

1. 用 `datasets/m0_external_entry/v0.1/prompt_full.md` 跑 3–5 个模型。
2. 将原始输出按模型放入 `runs/` 子目录。
3. 按 `scoring/rubric_v0.1.md` 粗评。
4. 生成第一版三轴暂读报告。
