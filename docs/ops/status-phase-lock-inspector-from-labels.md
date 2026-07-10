# 巡检仪：labels ↔ 主星盘 fields

> 状态：v0，本地审计脚本。  
> 目标：检查 GitHub 中枢户口 labels 与 GitHub 主星盘字段是否漂移。

## 为什么需要它

现在已经有两层：

```text
status-phase-lock-from-labels.sh
GitHub issue labels → 主星盘 fields
```

但写完以后还需要一层只读巡检：

```text
主星盘 fields 是否仍然等于 issue labels 推导结果
```

否则后续手动改主星盘或 issue labels 时，两边可能再次漂移。

## 巡检规则

- `waterline:*` → `水位`
- `lane:*` → `晚上列`
- `waterline + lane` → `早晨列`

## 使用方法

```bash
gh auth refresh -s project
bash scripts/status-phase-lock-inspector-from-labels.sh
```

输出：

```text
docs/cockpit/status-phase-lock-inspector/YYYY-MM-DD/report.md
```

## 边界

这个脚本只读：

- 不写主星盘；
- 不写 issue；
- 不写 Notion；
- 不裁决候选是否拆户口；
- 只报告缺失 / 冲突 / 已对齐。

## 和状态锁相流的关系

```text
status-phase-lock-from-labels.sh = 写入 / 补齐
status-phase-lock-inspector-from-labels.sh = 只读 / 巡检
```
