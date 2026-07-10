# Steam → 资源/游戏 样本事实表策略 v0

## 为什么先不用正式旁表

直接给 `资源/游戏.csv` 追加十多个字段，会让 CSV 中每一行都被重写，`git diff` 变成全表噪声。把事实表直接放进 `资源/游戏/` 也会让 v0 探针提前变成正式资源结构调整。

因此 v0 收口为样本事实表：

```plain text
docs/ops/steam-game-resource-facts-sample-v0.csv
```

主表 `资源/游戏.csv` 和资源目录 `资源/游戏/` 暂时保持干净。样本事实表只保存 Steam 与人工确认事实的 v0 口径，后续稳定后再决定是否建立正式旁表，或是否迁回资源目录。

## 执行方式

```bash
node scripts/write-steam-game-facts-v0.mjs
node scripts/write-steam-game-facts-v0.mjs --write
```

默认 dry-run；只有 `--write` 才会创建或更新：

```plain text
docs/ops/steam-game-resource-facts-sample-v0.csv
```

脚本只写 docs/ops 下的 v0 样本 CSV，不修改 `资源/游戏.csv`，也不修改 `资源/游戏/`。

## 当前样本事实 / 候选

| title | 来源 | 写入内容 |
| --- | --- | --- |
| Cities: Skylines | Steam 复核队列候选 | `steam_appid=255710`、owned、时长、最近游玩；城市模拟 / 建造经营类不按通关处理；机制研究价值待谷确认 |
| Sid Meier's Civilization VI | Steam 复核队列候选 | `steam_appid=289070`、owned、时长、最近游玩；单局有限 + 长期策略循环；不由高时长推出通关、重要性或系统沉积 |
| 极乐迪斯科 Disco Elysium | Steam 探针 | `steam_appid=632470`、owned、时长、最近游玩；不判断通关或价值 |
| 史丹利的寓言：终级豪华版 The Stanley Parable: Ultra Deluxe | Steam 探针 | `steam_appid=1703340`、owned、时长、最近游玩；混合型口径待确认 |
| 粘粘世界 World of Goo | 谷人工确认 | PC + iPad 玩过，并已通关；非 Steam 来源，不得映射 `World of Warships` |

## 本地修正旧旁表位置

如果已经运行过旧脚本并在 `资源/游戏/` 下生成了旁表，先恢复资源目录里的样本文件：

```bash
git restore -- 资源/游戏/steam_facts.csv
```

然后重新走 v0 样本事实表：

```bash
git pull
node scripts/write-steam-game-facts-v0.mjs
node scripts/write-steam-game-facts-v0.mjs --write
git diff -- docs/ops/steam-game-resource-facts-sample-v0.csv
```

## 后续原则

- `资源/游戏.csv` 保持原有索引职责。
- `资源/游戏/` 暂不新增正式 Steam 旁表。
- `docs/ops/steam-game-resource-facts-sample-v0.csv` 只承接 v0 Steam / manual 样本事实。
- Steam 新条目候选继续留在 `资源/_steam/steam_review_queue.csv`，不自动入主表。
- 样本表里的 `steam_candidate` 不代表谷确认玩过、喜欢、重要、通关或沉积。
- 等字段稳定后，再考虑主表结构调整、正式旁表或迁移。
