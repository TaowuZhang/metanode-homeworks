# 绝对基岩｜2026-06-13

> 这份记录收束 2026-06-13 的 Notion × GitHub 中枢户口锁相施工。它不是新设计，而是绝对基岩确认。

## 1. 已部署的前台 / 承重面

### Notion

- `今日开工`：早晨取景器。
- `今日开工后台`：早晨卡与 Notion 活现场字段后台。
- 泊坞工作台：资料源、服务台、外显母表、视觉索引等活现场。

### GitHub

- GitHub 主星盘：全天行动承重总账 / 中枢户口集合。
- GitHub 中枢户口：每条行动户口的 hard anchor（技术上仍是 label: cockpit 的 issues）。
- GitHub 解理坞：晚上裁决样机，不是最终总账。
- `cockpit/ignition-wire` 分支：自动烘焙发车引线。

定位：

```text
GitHub 主星盘 = 行动总账
GitHub issues = 中枢户口
GitHub docs = 可版本化事实地层
```

## 2. 已部署的自动化 / 脚本

### `cockpit-ignition-wire`

作用：

```text
open 中枢户口
→ 按 waterline 过滤
→ docs/cockpit/ignition-wire.md
→ cockpit/ignition-wire 分支
```

### `account-phase-lock-dry-run`

作用：

```text
GitHub 中枢户口
+ docs/ops/account-phase-lock-map.json
→ 户口锁相流 dry-run 对账报告
```

边界：不写 Notion，不改主星盘字段，不裁决开户 / 拆户口。

### `status-phase-lock-from-labels.sh`

作用：

```text
issue labels
→ 主星盘 fields
```

同步字段：`水位`、`晚上列`、`早晨列`、`Last touched`。

### `status-phase-lock-inspector-from-labels.sh`

作用：

```text
主星盘 fields
↔ issue labels
→ 漂移报告
```

## 3. 当前边界

已经自动 / 半自动完成：

```text
GitHub 中枢户口 → 发车引线
GitHub 中枢户口 + mapping → 户口锁相流 dry-run 报告
GitHub issue labels → 主星盘 fields
主星盘 fields ↔ GitHub issue labels 巡检
```

尚未打开：

```text
Notion API 自动回写
主星盘字段自动定时写入
候选自动拆户口
issue 自动关闭 / 开户
PR 自动合并
```

这些不应在 v0 贸然打开。

## 4. 最终结论

```text
Notion 早晨面板已部署。
GitHub 主星盘已部署。
GitHub 中枢户口已对齐。
发车引线自动化已跑通。
户口锁相流 dry-run 对账已跑通。
状态锁相流已写入成功。
巡检仪已确认 18/18 对齐。
```
