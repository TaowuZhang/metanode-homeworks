# PR #372 判断链追回 Source Index（2026-07-06）

## 0. 用途

本文件是判断链追回的母表，不是判断账本。

它只回答：

- 当前有哪些 TopHub 节点；
- 当前有哪些版本化可见的 Folo 来源；
- 哪些来源已经有 ledger；
- 哪些来源还没有判断链；
- 哪些来源因为账本故意省略名称，不能由当前模型补编。

本文件不得被用来替代逐源判断链。

---

## 1. 来源文件

- `资源/雷达/TopHub 全部订阅顺序.md`
- `资源/雷达/订阅源账本.yml`
- `docs/ops/pr372-decision-recovery-ledger-draft-20260706.md`
- `docs/ops/pr372-decision-recovery-ledger-comprehensive-20260706.md`
- `docs/ops/pr372-decision-recovery-ledger-technology-20260706.md`
- `docs/ops/pr372-decision-recovery-ledger-community-20260706.md`

---

## 2. TopHub 当前真实节点索引

当前平台真实订阅为 80 项：79 个内容节点 + 1 个系统级节点 `专栏｜订阅聚合`。

### 2.1 公共温度（01—16）

| order | source_id | current_route | ledger_status | next_action |
|---:|---|---|---|---|
| 1 | 微博｜热搜榜 | tophub_public_temperature | needs_recovery | 回综合 / 微博 closure |
| 2 | 小红书｜热榜 | tophub_public_temperature | needs_recovery | 回娱乐 / 社区或综合相关 closure |
| 3 | 知乎｜热榜 | tophub_public_temperature | recovered_from_pr | 已在 comprehensive ledger |
| 4 | 抖音｜总榜 | tophub_public_temperature | needs_recovery | 回综合 / 娱乐相关 closure |
| 5 | 微博｜话题榜 | tophub_public_temperature | needs_recovery | 回综合 / 微博 closure |
| 6 | 今日热榜｜实时榜中榜 | tophub_public_temperature | needs_recovery | 回综合目录判断链 |
| 7 | 微信｜24h热文榜 | tophub_public_temperature | needs_recovery | 回综合 / 微信 closure |
| 8 | 澎湃｜热榜 | tophub_public_temperature | needs_recovery | 回综合新闻源 closure |
| 9 | 中国新闻周刊｜十大热文 | tophub_public_temperature | needs_recovery | 回综合新闻源 closure |
| 10 | 新京报｜排行 | tophub_public_temperature | needs_recovery | 回综合新闻源 closure |
| 11 | 南方周末｜热门文章 | tophub_public_temperature | needs_recovery | 回综合 / 南方周末 closure |
| 12 | 观察者网｜要闻 | tophub_public_temperature | needs_recovery | 回综合新闻源 closure |
| 13 | 参考消息｜滚动新闻 | tophub_public_temperature | needs_recovery | 回综合滚动新闻 closure |
| 14 | 中央电视台｜新闻联播 | tophub_public_temperature | needs_recovery | 回综合 / 央视 closure |
| 15 | 半月谈｜今日谈 | tophub_public_temperature | needs_recovery | 回综合 / 半月谈 closure |
| 16 | 新华每日电讯｜电子报 | tophub_public_temperature | recovered_from_pr | 已在 draft ledger / coverage audit；可补更细行号 |

### 2.2 科技雷达（17—28）

| order | source_id | current_route | ledger_status | next_action |
|---:|---|---|---|---|
| 17 | 极客公园｜每日最新 | tophub_technology_radar | needs_recovery | 回科技 general tech closure |
| 18 | 少数派｜热门文章 | tophub_technology_radar | needs_recovery | 回科技少数派 closure；与 Folo 移出关系需补 |
| 19 | TechCrunch｜Today | tophub_technology_radar | needs_recovery | 回科技 global / general tech closure |
| 20 | The Verge｜Today | tophub_technology_radar | needs_recovery | 回科技 global / general tech closure |
| 21 | IT之家｜日榜 | tophub_technology_radar | needs_recovery | 回 IT 之家 closure |
| 22 | The Register｜Latest | tophub_technology_radar | recovered_from_pr | 已在 technology ledger；仍可补同族 closure |
| 23 | 先知社区｜精华推荐 | tophub_technology_radar | recovered_from_pr | 已在 community ledger；需补 security closure |
| 24 | 科普中国网｜头条 | tophub_technology_radar | partial | 已在 technology ledger；需补科普 closure |
| 25 | 果壳｜科学人 | tophub_technology_radar | needs_recovery | 回科技科学 / 科普 closure |
| 26 | Science Magazine｜Latest News | tophub_technology_radar | partial | 已在 technology ledger；需补 science closure |
| 27 | HelloGitHub｜月刊 | tophub_technology_radar | needs_recovery | 回开发 / 科技 closure |
| 28 | Apple 支持｜更换和维修扩展计划 | tophub_technology_radar | recovered_from_pr | 已在 technology ledger；可补 Apple closure |

### 2.3 数据与结构（29—40）

| order | source_id | current_route | ledger_status | next_action |
|---:|---|---|---|---|
| 29 | 国家统计局｜最新数据发布 | tophub_data_and_structure | needs_recovery | 回财经 / 综合数据源判断 |
| 30 | 国家统计局｜数据解读 | tophub_data_and_structure | needs_recovery | 回财经 / 综合数据源判断 |
| 31 | 财新网｜首页推荐 | tophub_data_and_structure | recovered_from_pr | 已在 draft ledger / coverage audit |
| 32 | 日经中文网｜每日最新 | tophub_data_and_structure | recovered_from_pr | 已在 draft ledger / coverage audit |
| 33 | FT中文网｜十大热门文章 | tophub_data_and_structure | partial | draft ledger 有保留关系；需补财经页账本 |
| 34 | 中央电视台｜经济半小时 | tophub_data_and_structure | needs_recovery | 回综合 / 央视 / 财经 closure |
| 35 | 中央电视台｜每周质量报告 | tophub_data_and_structure | needs_recovery | 回综合 / 央视 closure |
| 36 | 中央电视台｜新闻调查 | tophub_data_and_structure | recovered_from_pr | 已在 comprehensive ledger |
| 37 | 36氪出海｜热门推荐 | tophub_data_and_structure | partial | 已在 technology ledger；需补 36kr / globalization closure |
| 38 | 第一财经｜汽车新闻 | tophub_data_and_structure | recovered_from_pr | 已在 technology / draft ledger；可补 automotive closure |
| 39 | Counterpoint Research｜最新见解 | tophub_data_and_structure | recovered_from_pr | 已在 technology ledger；可补 reports / digital closure |
| 40 | 华丽志 | tophub_data_and_structure | recovered_from_pr | 已在 draft ledger；可补娱乐 / 消费产业来源理由 |

### 2.4 慢读与思想（41—53）

| order | source_id | current_route | ledger_status | next_action |
|---:|---|---|---|---|
| 41 | 三联生活周刊｜封面故事 | tophub_slow_read_and_thought | needs_recovery | 回综合 / 报刊 / 慢读来源判断 |
| 42 | 爱思想｜每周文章排行 | tophub_slow_read_and_thought | needs_recovery | 回专栏 / 思想来源判断 |
| 43 | 知乎日报｜Today | tophub_slow_read_and_thought | recovered_from_pr | 已在 comprehensive ledger |
| 44 | 豆瓣｜热门图书-虚构类 | tophub_slow_read_and_thought | needs_recovery | 回娱乐 / 书籍 closure |
| 45 | 微信读书｜新书榜 | tophub_slow_read_and_thought | needs_recovery | 回娱乐 / 书籍 closure |
| 46 | 豆瓣｜新书速递 | tophub_slow_read_and_thought | needs_recovery | 回娱乐 / 书籍 closure |
| 47 | 豆瓣｜最受欢迎的书评 | tophub_slow_read_and_thought | needs_recovery | 回娱乐 / 书评 closure |
| 48 | Cinephilia 迷影 | tophub_slow_read_and_thought | needs_recovery | 回娱乐 / 电影评论 closure |
| 49 | 豆瓣电影｜最受欢迎的影评 | tophub_slow_read_and_thought | recovered_from_pr | 已在 draft ledger；需补娱乐单项时暂缓 |
| 50 | 书格｜每日好书 | tophub_slow_read_and_thought | recovered_from_pr | 已在 draft ledger |
| 51 | 器物于我 | tophub_slow_read_and_thought | needs_recovery | 回专栏 / 慢读 closure |
| 52 | 爱思想｜每日文章排行 | tophub_slow_read_and_thought | needs_recovery | 回专栏 / 思想来源判断 |
| 53 | 历史上的今天｜Today | tophub_slow_read_and_thought | needs_recovery | 回综合 / 工具型历史入口判断 |

### 2.5 生活与社区（54—60）

| order | source_id | current_route | ledger_status | next_action |
|---:|---|---|---|---|
| 54 | 下厨房｜24小时最佳 | tophub_life_and_community | needs_recovery | 回综合 / 生活来源判断 |
| 55 | 科普中国｜今日辟谣文章 | tophub_life_and_community | recovered_from_pr | 已在 comprehensive ledger |
| 56 | 半月谈｜健康 | tophub_life_and_community | recovered_from_pr | 已在 comprehensive ledger |
| 57 | 豆瓣话题广场｜热门话题 | tophub_life_and_community | needs_recovery | 回社区 / 豆瓣 closure |
| 58 | 豆瓣小组｜讨论精选 | tophub_life_and_community | needs_recovery | 回社区 / 豆瓣 closure |
| 59 | V2EX 周报 | tophub_life_and_community | recovered_from_pr | 已在 community ledger；需补 V2EX closure |
| 60 | 马蜂窝｜热门游记 | tophub_life_and_community | recovered_from_pr | 已在 community ledger；需补 travel closure |

### 2.6 影游音乐（61—72）

| order | source_id | current_route | ledger_status | next_action |
|---:|---|---|---|---|
| 61 | Apple Music｜专辑排行 | tophub_film_game_music | needs_recovery | 回娱乐音乐 closure |
| 62 | 豆瓣｜华语新碟榜 | tophub_film_game_music | recovered_from_pr | 已在 draft ledger |
| 63 | 网易云音乐｜原创榜 | tophub_film_game_music | recovered_from_pr | 已在 draft ledger |
| 64 | 豆瓣电影｜一周口碑榜 | tophub_film_game_music | needs_recovery | 回娱乐电影 closure |
| 65 | IMDb｜热门电影榜 | tophub_film_game_music | needs_recovery | 回娱乐电影 closure |
| 66 | 豆瓣｜全球口碑剧集榜 | tophub_film_game_music | recovered_from_pr | 已在 draft ledger |
| 67 | 新京报｜娱乐 | tophub_film_game_music | recovered_from_pr | 已在 draft ledger |
| 68 | 哔哩哔哩｜每周必看 | tophub_film_game_music | needs_recovery | 回娱乐 / 视频 closure |
| 69 | 游研社｜首页推荐 | tophub_film_game_music | needs_recovery | 回娱乐 / 游戏 closure |
| 70 | INDIENOVA｜文章 | tophub_film_game_music | recovered_from_pr | 已在 draft ledger |
| 71 | 机核网｜每日最新 | tophub_film_game_music | needs_recovery | 回娱乐 / 游戏 closure |
| 72 | 喷嚏网｜乐影 | tophub_film_game_music | recovered_from_pr | 已在 draft ledger；平台更新状态已确认 |

### 2.7 视觉与自然（73—79）

| order | source_id | current_route | ledger_status | next_action |
|---:|---|---|---|---|
| 73 | iDaily · 每日环球视野｜Today | tophub_visual_and_nature | needs_recovery | 回娱乐 / 视觉 closure |
| 74 | 中国国家地理网｜热度榜 | tophub_visual_and_nature | needs_recovery | 回视觉 / 自然 closure |
| 75 | CNU视觉联盟｜每日精选 | tophub_visual_and_nature | recovered_from_pr | 已在 draft ledger |
| 76 | 胶片的味道 | tophub_visual_and_nature | recovered_from_pr | 已在 draft ledger |
| 77 | NASA 🌍｜每日星球 | tophub_visual_and_nature | needs_recovery | 回视觉 / 自然 closure；与 Folo NASA APOD 区分 |
| 78 | 北京天文馆｜每日一图 | tophub_visual_and_nature | needs_recovery | 回视觉 / 自然 closure |
| 79 | 果壳｜物种日历 | tophub_visual_and_nature | needs_recovery | 回视觉 / 自然 closure |

### 2.8 系统聚合入口（80）

| order | source_id | current_route | ledger_status | next_action |
|---:|---|---|---|---|
| 80 | 专栏｜订阅聚合 | tophub_system_aggregate | recovered_from_pr | 已在 community ledger 标为系统级总览入口；专栏目录仍需单独追回 |

---

## 3. Folo 当前可见来源索引

Folo 当前总订阅 27 项。版本化账本中可见 20 项，另有 7 项名称故意省略，不得由当前模型补编。

### 3.1 trigger（3/3 可见）

| source_id | folo_category | ledger_status | next_action |
|---|---|---|---|
| Google Developers Blog | trigger | partial | 2026-07-17 复查；需与 Google AI Developers 比较 |
| Google DeepMind News | trigger | needs_recovery | 回 Folo / AI 来源判断 |
| Obsidian Plugins | trigger | needs_recovery | 回工具 / Obsidian 来源判断 |

### 3.2 deep_read（7/7 可见）

| source_id | folo_category | ledger_status | next_action |
|---|---|---|---|
| 小众软件 | deep_read | partial | 已在科技 / 反斗软件拒绝关系中被引用；需补 Folo 原始理由 |
| 晚点 - 长报道 | deep_read | recovered_from_pr_partial | 综合 / 知乎想法热榜已说明不重复添加晚点；需补 Folo 原始理由 |
| Last Week in AI | deep_read | partial | AI ledger 中与 Interconnects 覆盖关系出现；需补 Folo 原始理由 |
| darkreading | deep_read | partial | 科技 / iThome 拒绝关系中被引用；需补安全来源理由 |
| Krebs on Security | deep_read | partial | 科技 / iThome 拒绝关系中被引用；需补安全来源理由 |
| Steph Ango | deep_read | needs_recovery | 回旧对话或 Folo 分类重组 |
| 有知有行 | deep_read | recovered_from_pr | 已在财经 ledger / draft ledger |

### 3.3 wander（5/6 可见，1 项省略）

| source_id | folo_category | ledger_status | next_action |
|---|---|---|---|
| Nat Geo Photo of the Day | wander | partial | 与 TopHub 视觉 / 自然区分待补 |
| NASA APOD | wander | partial | 账本记录已替换为官方 RSS；需补与 TopHub NASA 每日星球区分 |
| Magnum Photos | wander | needs_recovery | 回视觉 / 摄影来源判断 |
| 中国爬楼联盟 | wander | partial | 科技 final closeout 提到与街拍中国最多保留一个；需补理由 |
| 街拍中国 | wander | partial | 同上 |
| 省略项 1 | wander | omitted_by_choice | 不得补编名称；只能保持省略 |

### 3.4 reminder（0/4 可见，4 项省略）

| source_id | folo_category | ledger_status | next_action |
|---|---|---|---|
| 省略项 1—4 | reminder | omitted_by_choice | 不得补编名称；观察期后按真实账本复查 |

### 3.5 trial（5/7 可见，2 项省略）

| source_id | folo_category | ledger_status | next_action |
|---|---|---|---|
| Stack Overflow Blog | trial | partial | 科技 final closeout 列为 2026-07-17 复查对象；需补判断链 |
| Google AI Developers | trial | partial | 与 Google Developers Blog 是否重复待复查 |
| Logan Kilpatrick | trial | partial | 2026-07-17 复查；需补理由 |
| News Minimalist | trial | partial | 2026-07-17 复查；需补理由 |
| Hacker News | trial | partial | 已替换为官方 RSS；需补保留 / 退出条件 |
| 省略项 1—2 | trial | omitted_by_choice | 不得补编名称 |

---

## 4. Folo 已验证替换与移出来源

### 4.1 已验证替换

| source_id | old_feed_removed | new_route | ledger_status | next_action |
|---|---:|---|---|---|
| Hacker News | true | official_rss | partial | 补为什么保留 HN、何时退出 |
| NASA APOD | true | official_rss | partial | 补为什么留 Folo、与 TopHub NASA 区分 |

### 4.2 第二波移出

| source_id | next_state | ledger_status | next_action |
|---|---|---|---|
| Lifehacker | tophub_or_on_demand | needs_recovery | 补为什么退出 Folo |
| Cyber Security News | tophub_or_on_demand | needs_recovery | 补与 Dark Reading / Krebs / 先知关系 |
| 领研 论文 计算机 | on_demand | needs_recovery | 补为什么按需 |
| 中国气象局 每日天气提示 | on_demand | needs_recovery | 补为什么不做常驻提醒 |
| 东方财富网 策略报告 | on_demand | needs_recovery | 补为什么按需 |
| 财经日历 华尔街见闻 | on_demand | recovered_from_pr | 已在财经 ledger 中作为工具链接 |

### 4.3 已移出命名来源

| source_id | next_state | ledger_status | next_action |
|---|---|---|---|
| 少数派 | tophub | partial | TopHub 中保留，需补 Folo 退出理由 |
| 36氪 - 24小时热榜 | tophub | partial | TopHub 中被 36氪出海替换；需补完整链 |
| Bing 每日壁纸 | tophub | needs_recovery | 回视觉 / 自然判断 |
| InfoQ 推荐 | on_demand | needs_recovery | 回开发 / 技术来源判断 |
| 金十数据 | on_demand | needs_recovery | 回财经 / 高频噪声判断 |
| New Routes | on_demand | needs_recovery | 回旅行 / 航线判断 |
| Anthropic News | on_demand | needs_recovery | 回 AI / 官方来源判断 |
| 创意文章榜 | on_demand | needs_recovery | 回设计 / 创意来源判断 |
| 创意作品榜 | retired_invalid | needs_recovery | 回设计 / 创意来源判断 |

---

## 5. Source Index 的当前结论

1. TopHub 80 项已经全部登记。
2. Folo 可见 20 项已经登记，省略 7 项明确标 `omitted_by_choice`。
3. 当前 recovered / partial 的来源只是一部分，不得宣称逐源判断已完成。
4. 下一步不是继续汇报数量，而是按 `needs_recovery` 和 `partial` 回对应 ledger / closure 逐个补判断链。
5. 本 index 是路线图，不是完成态账本。
