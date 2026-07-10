# Steam → 资源/游戏 季度候选同步

## 负责范围

本页由季度 GitHub Actions 同步生成，只承接 Steam 外部事实候选，不直接修改 `资源/游戏.csv`，也不在 `资源/游戏/` 下新增正式 Steam 旁表。

## 本次观察

- observed_at: `2026-07-03T05:27:07.693Z`
- observed_quarter: `2026-Q3`
- candidate_rows: 23

## 复核桶统计

```json
{
  "byBucket": {
    "existing_match_candidate": 2,
    "new_entry_candidate": 21
  },
  "byPlayed": {
    "yes": 16,
    "(empty)": 7
  },
  "byRecent": {
    "(empty)": 23
  }
}
```

## 高时长 / 优先复核候选

| steam_name | matched_worang_title | review_bucket | playtime_hours | last_played | suggested_action |
| --- | --- | --- | ---: | --- | --- |
| Sid Meier's Civilization VI |  | new_entry_candidate | 832.5 | 2025-11-03T10:13:44.000Z | 新条目候选；仅标试过候选，不判断价值 |
| Cities: Skylines |  | new_entry_candidate | 440.7 | 2024-12-29T09:19:57.000Z | 新条目候选；仅标试过候选，不判断价值 |
| Oxygen Not Included |  | new_entry_candidate | 53.2 | 2025-01-27T16:03:50.000Z | 新条目候选；仅标试过候选，不判断价值 |
| Disco Elysium | 极乐迪斯科 Disco Elysium | existing_match_candidate | 35.8 | 2023-09-30T05:43:05.000Z | 可复核后补 steam_appid / steam_owned；不自动改状态 |
| Cult of the Lamb |  | new_entry_candidate | 33.2 | 2025-01-23T14:34:29.000Z | 新条目候选；仅标试过候选，不判断价值 |
| Broforce |  | new_entry_candidate | 31.9 | 2025-05-01T17:22:52.000Z | 新条目候选；仅标试过候选，不判断价值 |
| Wallpaper Engine |  | new_entry_candidate | 30.6 | 2025-06-18T05:26:07.000Z | 新条目候选；仅标试过候选，不判断价值 |
| The Stanley Parable: Ultra Deluxe | 史丹利的寓言：终级豪华版 The Stanley Parable: Ultra Deluxe | existing_match_candidate | 13.9 | 2022-08-03T15:11:59.000Z | 可复核后补 steam_appid / steam_owned；不自动改状态 |
| It Takes Two Friend's Pass |  | new_entry_candidate | 9.1 | 2023-10-24T17:00:21.000Z | 新条目候选；仅标试过候选，不判断价值 |
| Muse Dash |  | new_entry_candidate | 3.4 | 2024-10-31T02:42:18.000Z | 新条目候选；仅标试过候选，不判断价值 |
| Townscaper |  | new_entry_candidate | 2.9 | 2024-08-26T11:29:29.000Z | 新条目候选；仅标试过候选，不判断价值 |
| Cultist Simulator |  | new_entry_candidate | 2.8 | 2025-05-01T05:24:48.000Z | 新条目候选；仅标试过候选，不判断价值 |

## 边界

- 不提交 `资源/_steam/` 本地探针输出。
- 不修改 `资源/游戏.csv`。
- 不修改 `资源/游戏/`。
- 不把 Steam owned / playtime / achievement 自动等同于喜欢、重要、通关或沉积。
- 所有主表、正式旁表、价值与完成状态判断仍需谷确认。
