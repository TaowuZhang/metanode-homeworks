# Steam → 资源/游戏 只读探针 v0

## 负责范围

这个探针只验证 Steam 数据能否作为 `资源/游戏` 的外部事实源，并把可读到的事实转成候选事件。

它负责：

- 读取 Steam owned games。
- 读取 recently played games。
- 可选读取已游玩游戏的 achievement 摘要。
- 与 `资源/游戏.csv` 做轻量标题匹配。
- 输出 `资源/_steam/steam_snapshot.json` 与 `资源/_steam/steam_candidate_events.csv`。
- 可选把候选事件整理为人工复核队列 `资源/_steam/steam_review_queue.csv`。

## 不负责范围

它不负责：

- 不写 Steam。
- 不发布或修改 Steam 评价。
- 不读取私密数据；只按 Steam API 与账号公开 / 授权可见范围读取。
- 不更新 `资源/游戏.csv` 主表。
- 不自动判断游戏价值。
- 不自动判断通关。
- 不自动把 Steam 库全量搬入沃壤。
- 不把游玩时长等同于重要性。

## 游戏领域判断口径

`有限 / 无限的游戏` 在这里只作为游戏领域内部的形态判断，不作为沃壤或腐海的总纲。

先判游戏形态，再判状态：

| 形态 | 可写完成态吗 | 主要状态 |
| --- | --- | --- |
| 有限型 | 可以 | 未开始 / 进行中 / 通关候选 / 主线通关 / 多结局完成 / 全成就 |
| 无限型 | 通常不写通关 | 未接触 / 想玩 / 已入库 / 试过 / 在玩 / 周期回访 / 暂停 / 退坑 |
| 混合型 | 拆双字段 | `finite_completion_status` + `ongoing_play_status` |

Steam API 只提供外部事实，不替谷判断这些字段。

## 候选事件

v0 只生成这些事件：

| event_type | 触发 | 候选含义 |
| --- | --- | --- |
| `steam_owned_detected` | Steam 库出现该游戏 | 已入库 / 补全 steam_appid 候选 |
| `steam_played_detected` | 总时长 > 0 | 试过候选；不等于重要，不等于通关 |
| `steam_recently_played` | 最近玩过 / 近两周有时长 | 在玩 / 回访候选 |

后续再考虑：

- `steam_achievement_seen`
- `steam_achievement_unlocked`
- `steam_review_seen`
- `steam_wishlist_seen`

## 人工复核队列

`steam_candidate_events.csv` 是事件表，一款游戏可能有多条事件。复核时可以先生成按游戏聚合的一行一游戏队列：

```bash
node scripts/steam-review-queue.mjs --dry-run
node scripts/steam-review-queue.mjs
```

默认输出：

```plain text
资源/_steam/steam_review_queue.csv
```

复核桶：

| review_bucket | 含义 |
| --- | --- |
| `existing_match_candidate` | 已匹配到 `资源/游戏.csv` 条目，可复核后补 Steam 字段 |
| `new_entry_candidate` | Steam 有，但 `资源/游戏.csv` 暂无匹配；不自动入主表 |
| `needs_match_review` | 模糊匹配，必须人工核对，避免错配污染主表 |

## 本地使用

不要把 Steam API key 粘到聊天里，也不要提交到 Git。

```bash
export STEAM_API_KEY="..."
export STEAM_ID="..."
node scripts/steam-probe.mjs --dry-run
node scripts/steam-probe.mjs
```

可选成就探针：

```bash
STEAM_PROBE_ACHIEVEMENTS=1 node scripts/steam-probe.mjs --achievement-limit 25
```

默认输出：

```plain text
资源/_steam/steam_snapshot.json
资源/_steam/steam_candidate_events.csv
```

## 需要谷确认的场景

以下场景必须保持候选，不得自动写主表：

- Steam 游戏与 `资源/游戏.csv` 模糊匹配。
- Steam 显示有时长，但谷不认为它“玩过”。
- Steam 显示最近玩过，但谷不认为它“在玩”。
- 成就像通关，但该游戏形态或完成条件不明确。
- 游戏可能属于有限型 / 无限型 / 混合型，需要具体判断。
- 游戏可能沉入 `领域/游戏设计`、`领域/作品`、`成为/项目` 或 `语境`。

## 下一步

1. 在本地填入 `STEAM_API_KEY` 与 `STEAM_ID`。
2. 先跑 `--dry-run` 看 Steam API 是否可读。
3. 若可读，生成 `资源/_steam/` 两个候选文件。
4. 生成 `steam_review_queue.csv`，把事件表压成一行一游戏的人工复核队列。
5. 挑少量具体游戏复核形态：有限型 / 无限型 / 混合型。
6. 再决定是否给 `资源/游戏.csv` 增加稳定字段。
