# 状态锁相流：从 issue labels 到 主星盘 fields

> 状态：v0，本地手动脚本。  
> 目标：把已经整理好的 GitHub 中枢户口 labels，同步成 GitHub 主星盘字段值。

## 为什么需要它

户口锁相流 v0 已经做到：

```text
Notion 工作台 / 今日开工后台
↔ GitHub 中枢户口
↔ account-phase-lock-map.json
→ dry-run 对账报告
```

但还有一层低风险字段需要补齐：

```text
GitHub issue labels → GitHub 主星盘 fields
```

## 同步规则

### waterline → 水位

| Label | 主星盘字段 |
|---|---|
| `waterline:今日` | `水位 = 今日` |
| `waterline:候补` | `水位 = 候补` |
| `waterline:水下` | `水位 = 水下` |

### lane → 晚上列

| Label | 主星盘字段 |
|---|---|
| `lane:待裁决` | `晚上列 = 待裁决` |
| `lane:谷WIP` | `晚上列 = 谷 WIP` |
| `lane:AI可推进` | `晚上列 = AI 可推进` |
| `lane:暂留` | `晚上列 = 暂留` |
| `lane:代谢台` | `晚上列 = 代谢台` |
| `lane:完成退相干` | `晚上列 = 完成·退相干` |

### waterline + lane → 早晨列

| 条件 | 早晨列 |
|---|---|
| `waterline:水下` | `不进早上` |
| `waterline:今日` + `lane:谷WIP` | `直接做` |
| `waterline:今日` + `lane:AI可推进` | `交给克` |
| `waterline:今日` + 其他 lane | `今天占位` |
| `waterline:候补` | `今天占位` |

## 使用方法

```bash
gh auth refresh -s project
bash scripts/status-phase-lock-from-labels.sh
RUN_WRITE=1 bash scripts/status-phase-lock-from-labels.sh
```

## 边界

- 只同步 `水位`、`晚上列`、`早晨列`、`Last touched`。
- 不关闭 issue。
- 不创建 / 删除 issue。
- 不写 Notion。
- 不判断候选是否独立开户。
- 不改 `接续之问`、`给克一句话`、`晚上判断` 这类需要语义判断的字段。

## 和户口锁相流的关系

```text
account-phase-lock-dry-run = 对账报告
status-phase-lock-from-labels = 低风险字段补齐
```
