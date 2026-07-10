# Steam → 资源/游戏 样本判定 v0

## 负责范围

本页只承接 Steam 探针跑通后的第一批样本判定，用来校准 `资源/游戏` 后续字段与人工复核口径。

它负责：

- 把 Steam 候选队列里的少量代表游戏压成可讨论的字段候选。
- 区分 Steam 外部事实、谷人工事实、游戏形态判断、后续沉积判断。
- 明确哪些内容可以后续写入 `资源/游戏.csv`，哪些仍必须等待谷确认。

它不负责：

- 不直接更新 `资源/游戏.csv`。
- 不提交本地 `资源/_steam/` 输出。
- 不把游玩时长等同于重要性。
- 不把通关等同于沉积。
- 不把 Steam 事实覆盖谷的非 Steam 游戏经历。

## 字段口径候选

后续若扩展 `资源/游戏.csv`，优先考虑以下字段：

| 字段 | 含义 |
| --- | --- |
| `steam_appid` | Steam AppID；只在匹配可信时填写 |
| `steam_owned` | 当前 Steam 账号是否拥有 / 可见 |
| `steam_playtime_total_min` | Steam 记录的总游玩分钟 |
| `steam_last_played_at` | Steam 最近游玩时间 |
| `external_sources` | 来源，如 `douban;steam;manual` |
| `experience_source` | 当前状态判断来源，如 `steam_candidate` / `manual_confirmed` |
| `play_sources` | 非 Steam 平台来源，如 `PC;iPad` |
| `game_form` | 有限型 / 无限型 / 混合型 / 工具型 / 未判 |
| `finite_completion_status` | 未开始 / 进行中 / 通关候选 / 主线通关 / 多结局完成 / 全成就 / 不适用 / 待确认 |
| `ongoing_play_status` | 想玩 / 已入库 / 试过 / 在玩 / 周期回访 / 暂停 / 退坑 / 待确认 |
| `value_usage` | 无 / 机制研究候选 / 机制研究 / 作品参考候选 / 作品参考 / 项目相关 / 语境相关 |
| `status_confirmed` | 是否经谷确认 |
| `decision_note` | 人工判断说明 |

## 五个样本判定

| 游戏 | 当前来源 | Steam AppID | Steam 事实 | 游戏形态候选 | 完成 / 状态候选 | 价值用途 | 建议动作 |
| --- | --- | --- | --- | --- | --- | --- | --- |
| 极乐迪斯科 Disco Elysium | `资源/游戏.csv` 已有 + Steam 匹配 | `632470` | owned=yes；playtime=2145min；last_played=2023-09-30 | 有限型为主 | `finite_completion_status=待确认`；不由 Steam 时长推出通关 | 作品参考候选 / 叙事样本候选，需谷确认 | 可后续补 Steam 字段；不自动改评分、通关、价值 |
| 史丹利的寓言：终级豪华版 The Stanley Parable: Ultra Deluxe | `资源/游戏.csv` 已有 + Steam 匹配 | `1703340` | owned=yes；playtime=836min；last_played=2022-08-03 | 混合型：多结局有限层 + 元游戏回访层 | `finite_completion_status=待确认`；`ongoing_play_status=暂停候选` | 作品参考候选 / 交互叙事候选，需谷确认 | 可后续补 Steam 字段；通关口径需单独聊 |
| Cities: Skylines | Steam 新条目候选 | `255710` | owned=yes；playtime=26440min；last_played=2024-12-29 | 无限型 / 混合型：城市模拟、建造、系统经营 | 不按通关处理；`ongoing_play_status=长期玩过后暂停候选` | 机制研究候选，需谷确认 | 可列入新条目候选；高时长不等于重要性 |
| Sid Meier's Civilization VI | Steam 新条目候选 | `289070` | owned=yes；playtime=49950min；last_played=2025-11-03 | 混合型：单局胜利有限层 + 长期策略循环 | 不写“通关”；可写“单局胜利 / 多局经验待确认” | 机制研究候选 / 系统设计候选，需谷确认 | 可列入新条目候选；优先人工复核 |
| 粘粘世界 World of Goo | `资源/游戏.csv` 已有 + 谷人工补充 | 待查；不得由 `World of Warships` 错配 | 非 Steam 事实：小时候在电脑和 iPad 玩过，并已通关 | 有限型 | `finite_completion_status=主线通关`；`experience_source=manual_confirmed` | 价值用途待判，不由通关自动沉积 | 后续可补人工事实；不得写入 `steam_appid=552990` |

## 明确禁止的映射

| Steam 游戏 | 禁止映射到 | 原因 |
| --- | --- | --- |
| World of Warships (`552990`) | 粘粘世界 World of Goo | 仅英文泛词重叠导致的错配；已由收紧匹配规则消除 |
| Cult of the Lamb (`1313140`) | 丁丁历险记 | 仅 token 误伤，不可写主表 |
| The Awesome Adventures of Captain Spirit (`845070`) | 丁丁历险记 | 冒险类泛词误伤，不可写主表 |

## 复核顺序建议

第一轮只复核五个样本，不全量处理 23 个 Steam 条目：

1. 极乐迪斯科：验证已有条目补 Steam 字段的方式。
2. 史丹利的寓言：验证混合型 / 多结局口径。
3. Cities: Skylines：验证无限型 / 高时长不等于重要性。
4. Civilization VI：验证单局有限 + 长期无限循环。
5. 粘粘世界：验证非 Steam 人工通关事实。

## 后续写入原则

可自动或半自动写入的，只限外部事实字段：

- `steam_appid`
- `steam_owned`
- `steam_playtime_total_min`
- `steam_last_played_at`
- `external_sources` 中追加 `steam`

必须谷确认后才能写入的字段：

- `game_form`
- `finite_completion_status`
- `ongoing_play_status`
- `value_usage`
- `decision_note` 中的主观判断

例外：谷已经明确补充的人工事实可以作为 `manual_confirmed`，例如：

```plain text
粘粘世界 World of Goo：小时候在电脑与 iPad 玩过，并已通关。
```
