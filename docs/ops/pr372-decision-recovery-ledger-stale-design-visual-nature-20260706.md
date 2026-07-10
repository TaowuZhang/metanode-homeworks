# PR #372 判断链追回：失效节点、设计目录与视觉自然边界（2026-07-06）

## 0. 边界

本文件补充追回 PR #372 中三个剩余专项的判断链：

1. 当前 TopHub 失效节点清理；
2. `TopHub > 设计`目录逐项复审；
3. `视觉与自然`分组实测与设计 / 摄影 / 自然边界。

纳入文件：

- `docs/ops/tophub-current-subscriptions-stale-node-review-20260705.md`
- `docs/ops/tophub-stale-node-removal-platform-verified-20260705.md`
- `docs/ops/tophub-design-reassessment-pages-01-03-20260705.md`
- `docs/ops/tophub-design-final-closure-20260705.md`
- `docs/ops/tophub-visual-nature-audit-20260704.md`
- `资源/雷达/TopHub 全部订阅顺序.md`
- `资源/雷达/TopHub 使用配置.md`

当前不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 失效节点总规则：旧题材不等于失效，失效也不自动补位

```yaml
- source_id: "TopHub 当前订阅失效节点复核整体"
  current_route: platform_verified_stale_node_removal_no_backfill
  decision_status: recovered_from_pr_and_platform_verified
  selected_reason: "对原有 84 个订阅进行失效节点初筛时，严格区分旧题材、节点是否持续抓取新内容、低频告警节点的低更新特性，以及标题长期停留旧事件并经用户确认不再更新的真实失效。"
  rejected_alternatives:
    - "看到旧年份标题就直接删除"
    - "低频节点没有新条目就判定失效"
    - "取消后为了维持整数或分组数量找替代"
    - "把停止更新节点继续留在公共雷达"
  route_reason: "只有标题长期停留旧事件，且用户确认不再更新，才进入取消清单。四个节点已在 TopHub 平台真实取消，新增 0，当前真实 TopHub 从 84 调整为 80。"
  reuse_rule: "判断失效必须同时看内容日期、主题性质、节点角色、是否低频、用户界面确认和平台核验。取消后是否补位由功能缺口决定，不由数量决定。"
  evidence_locator: "docs/ops/tophub-current-subscriptions-stale-node-review-20260705.md#触发方式; docs/ops/tophub-stale-node-removal-platform-verified-20260705.md#结论"
  next_action: "作为失效节点处理规则。"
```

---

## 2. 四个取消节点：为什么取消、为什么不补

```yaml
- source_id: "微信｜热词"
  current_route: retired_stale_public_temperature_no_backfill
  decision_status: recovered_from_pr_and_platform_verified
  selected_reason: "原角色是公共温度中的微信热词入口。"
  rejected_alternatives:
    - "继续保留"
    - "用另一个微信热词或平台词条补位"
  route_reason: "当前标题混有冷冬、秋收等明显不合 2026 年 7 月时令的旧内容，用户确认节点不再更新；同时它与 `微信｜24h热文榜` 和 `今日热榜｜实时榜中榜`重复。"
  reuse_rule: "平台热词类节点失效后，不必补另一个热词；若公共温度仍足够覆盖，就压缩。"
  evidence_locator: "docs/ops/tophub-current-subscriptions-stale-node-review-20260705.md#公共温度; docs/ops/tophub-stale-node-removal-platform-verified-20260705.md#平台证据"
  next_action: "已取消，不替换。"

- source_id: "后续｜跟踪 Live"
  current_route: retired_stale_public_temperature_no_backfill
  decision_status: recovered_from_pr_and_platform_verified
  selected_reason: "原角色是事件后续追踪入口。"
  rejected_alternatives:
    - "继续保留旧 Live 节点"
    - "寻找新的后续追踪节点补位"
  route_reason: "当前前列仍是养殖貂新冠、苏宁收购家乐福、黄冈卫健委、瑞幸会不会成功、波音 737 MAX 等 2019—2020 年旧事件，用户确认节点失效。"
  reuse_rule: "事件追踪源如果停止更新，就不能再承担公共温度。真正需要跟踪事件时，应建立具体事件、官方来源或精确追踪，而不是保留旧 Live。"
  evidence_locator: "docs/ops/tophub-current-subscriptions-stale-node-review-20260705.md#公共温度; docs/ops/tophub-stale-node-removal-platform-verified-20260705.md#平台证据"
  next_action: "已取消，不替换。"

- source_id: "国家留学网｜综合项目专栏"
  current_route: retired_stale_life_community_no_backfill_on_demand_official_tool
  decision_status: recovered_from_pr_and_platform_verified
  selected_reason: "原角色是生活与社区中的留学项目入口。"
  rejected_alternatives:
    - "继续保留停止更新的项目专栏"
    - "用其他留学项目榜补位"
  route_reason: "当前仍以 2025、2024、2023 年项目为主，未出现 2026 年项目，用户确认节点不再更新。国家留学网本身保留为按需官网工具，但不再作为常驻节点。"
  reuse_rule: "官方项目站点如果具体 TopHub 节点失效，站点仍可按任务使用；不要因官网重要就保留失效聚合节点。"
  evidence_locator: "docs/ops/tophub-current-subscriptions-stale-node-review-20260705.md#生活与社区; docs/ops/tophub-stale-node-removal-platform-verified-20260705.md#数量核验"
  next_action: "已取消，不替换。"

- source_id: "Yoho!潮流志｜每日潮闻"
  current_route: retired_stale_life_community_no_backfill
  decision_status: recovered_from_pr_and_platform_verified
  selected_reason: "原角色是生活与社区中的潮流入口。"
  rejected_alternatives:
    - "继续保留旧潮流节点"
    - "用另一个潮流 / 时尚热榜补位"
  route_reason: "当前前列仍出现 2024 年活动和旧品牌公关稿，用户确认节点不再更新；该角色也不构成当前结构中的必要缺口。娱乐与数据结构中已有华丽志承担消费产业，视觉 / 作品 / 时尚灵感则按需。"
  reuse_rule: "潮流来源若只是旧品牌公关稿或停止更新，不补位。潮流不是必须常驻的生活节点。"
  evidence_locator: "docs/ops/tophub-current-subscriptions-stale-node-review-20260705.md#生活与社区; docs/ops/tophub-stale-node-removal-platform-verified-20260705.md#平台证据"
  next_action: "已取消，不替换。"
```

---

## 3. 两个被怀疑但保留的节点：不能只看标题年份

```yaml
- source_id: "新华每日电讯｜电子报"
  current_route: keep_public_temperature_updated_official_broad_newspaper
  decision_status: recovered_from_pr_and_platform_verified
  selected_reason: "用户确认仍在更新，当前同时出现电子商务法修正草案、快递数据和当期世界杯等新内容。"
  rejected_alternatives:
    - "因报刊类曾有失效问题而一并取消"
  route_reason: "它是报刊等量替换后的官方广谱日报入口，且本次失效复核确认仍在更新。"
  reuse_rule: "同类曾有失效不代表新节点失效；要逐节点确认。"
  evidence_locator: "docs/ops/tophub-current-subscriptions-stale-node-review-20260705.md#明确保留; docs/ops/tophub-stale-node-removal-platform-verified-20260705.md#平台证据"
  next_action: "保留。"

- source_id: "喷嚏网｜乐影"
  current_route: keep_movie_game_music_documentary_discovery
  decision_status: recovered_from_pr_and_platform_verified
  selected_reason: "用户确认仍在更新，标题中的 2020、2022、2024 多为影视作品年份，不代表节点冻结。"
  rejected_alternatives:
    - "看到标题中旧年份就取消"
  route_reason: "纪录片、电影、历史影像和作品推荐标题中天然会出现作品年份。这里旧年份是作品对象，不是抓取停滞证据。"
  reuse_rule: "作品发现源的标题年份要区分作品年份与发布时间。不能把作品年份当节点失效证据。"
  evidence_locator: "docs/ops/tophub-current-subscriptions-stale-node-review-20260705.md#明确保留; docs/ops/tophub-stale-node-removal-platform-verified-20260705.md#平台证据"
  next_action: "保留。"
```

---

## 4. 为什么取消后不补位

```yaml
- source_id: "失效节点取消后不补位规则"
  current_route: no_backfill_after_stale_removal
  decision_status: recovered_from_pr_and_platform_verified
  selected_reason: "公共温度剩余 16 个节点仍足够覆盖社交热度、新闻媒体、官方叙事和全国日报；生活与社区剩余 7 个节点仍覆盖烹饪、辟谣、健康、豆瓣社区、V2EX 和旅行。"
  rejected_alternatives:
    - "为了总量整齐补回 84"
    - "为了公共温度保持 18 而找替代"
    - "为了生活与社区保持 9 而找替代"
  route_reason: "为维持整数或分组数量而寻找替代，会重新制造无明确角色的订阅。有效节点优先于总量。"
  reuse_rule: "删除失效节点后先问现有分组是否仍能完成角色；若能，不补。"
  evidence_locator: "docs/ops/tophub-current-subscriptions-stale-node-review-20260705.md#为什么不补位; docs/ops/tophub-stale-node-removal-platform-verified-20260705.md#数量核验"
  next_action: "不补位。"
```

---

## 5. 设计目录整体：不缺作品瀑布流，缺设计文化和材料空间视角

```yaml
- source_id: "TopHub > 设计 30 节点整体"
  current_route: closed_no_tophub_change_one_folo_review_candidate
  decision_status: recovered_from_pr
  selected_reason: "设计目录 30/30 已逐项判断完成。它并不缺作品和灵感瀑布流，缺的是有编辑判断、能解释材料、工艺、空间与设计文化的来源；不被 AI 教程、作品集营销和商业软文占满的方法来源；以及将视觉漫游和工作任务分开的规则。"
  rejected_alternatives:
    - "加入站酷、Dribbble、Behance、500px 等作品流"
    - "把优设、站酷、UI 中国等 AI 教程和工具清单设为持续订阅"
    - "为了补设计来源扩张 TopHub 到 85"
    - "在 2026-07-17 前增加 Folo"
  route_reason: "最终 TopHub 新增 0、取消 0、替换 0、调序 0；Folo 当前新增 0；2026-07-17 新增复查候选 `designboom`。"
  reuse_rule: "设计来源先分作品瀑布流、教程工具、品牌识别、UX 方法档案、设计系统、广告案例、设计文化 / 材料 / 空间。只有最后一类可能补系统缺口；其他按任务。"
  evidence_locator: "docs/ops/tophub-design-reassessment-pages-01-03-20260705.md#总体判断; docs/ops/tophub-design-final-closure-20260705.md#最终判断"
  next_action: "2026-07-17 复查 designboom。"
```

---

## 6. designboom：唯一 A1，但不立即执行

```yaml
- source_id: "designboom"
  current_route: folo_review_candidate_design_material_space_culture
  decision_status: recovered_from_pr
  selected_reason: "它持续覆盖建筑、产品、家具、材料、工艺、公共空间、展览和设计文化，当前文章具有编辑判断，而不是单纯作品瀑布流，能补当前视觉体系中缺少的“设计如何进入物、空间与生活”的视角。"
  rejected_alternatives:
    - "立即加入 TopHub"
    - "立即加入 Folo"
    - "用站酷、Dribbble、Behance 或 500px 替代"
  route_reason: "列为 A1，进入 2026-07-17 Folo 统一复查。当前不修改 TopHub，也不立即增加 Folo。若通过复查，它承担开放天线，不进入 TopHub 热榜排序。"
  reuse_rule: "设计文化源即使是唯一强候选，也要经过更新频率、未读负担、内容比例和与现有来源互补性复查。"
  evidence_locator: "docs/ops/tophub-design-reassessment-pages-01-03-20260705.md#第2页; docs/ops/tophub-design-final-closure-20260705.md#Folo"
  next_action: "2026-07-17 Folo 复查。"
```

---

## 7. 设计旁边来源为什么按需或排除

```yaml
- source_id: "设计 A2 低频与任务来源"
  current_route: on_demand_design_task_sources
  decision_status: recovered_from_pr
  selected_reason: "标志情报局、月球背面设计素材周刊、腾讯CDC、腾讯社交用户体验设计、百度用户体验中心、WebdesignerNews、广告门、Behance Featured Projects 各有低频或任务价值。"
  rejected_alternatives:
    - "全部加入 TopHub 或 Folo"
    - "因为它们有用就占持续未读位"
  route_reason: "这些分别对应品牌识别、工具素材、产品设计与无障碍档案、社交产品体验、UX 评估方法、网页设计外链雷达、品牌传播和视觉灵感。价值明确但任务性强或更新不连续，不适合持续订阅。"
  reuse_rule: "任务源不等于失败。只有真实设计项目、品牌识别、UX 研究、网页设计或视觉灵感任务出现时调用。"
  evidence_locator: "docs/ops/tophub-design-reassessment-pages-01-03-20260705.md#A2低频或任务来源; docs/ops/tophub-design-final-closure-20260705.md#任务与低频来源"
  next_action: "按需。"

- source_id: "设计明确排除重复流"
  current_route: rejected_repetitive_or_polluted_design_streams
  decision_status: recovered_from_pr
  selected_reason: "无持续订阅价值。"
  rejected_alternatives:
    - "Dribbble 实时、本周、本月、年度和历史总榜"
    - "站酷作品总榜、全部推荐、文章总榜及重复推荐流"
    - "500px 热门作品"
    - "Designspiration Popular"
    - "Designmodo 当前邮件营销内容"
    - "UI 中国推荐文章"
  route_reason: "Dribbble 与站酷多个榜单高度重复或历史流量驱动；500px 与 CNU、中国国家地理、胶片的味道重叠；Designspiration 偏历史收藏；Designmodo 已偏邮件营销；UI 中国后部被 SEO 软文和服务稿污染。"
  reuse_rule: "作品平台和教程平台只要核心是展示稿、模板、历史流量、AI 教程、SEO 软文或营销工具，就不进入持续订阅。"
  evidence_locator: "docs/ops/tophub-design-reassessment-pages-01-03-20260705.md#明确排除的重复流; docs/ops/tophub-design-final-closure-20260705.md#明确不订阅"
  next_action: "排除。"
```

---

## 8. 视觉与自然：视觉来源启动器，不是视觉恢复界面本身

```yaml
- source_id: "TopHub｜视觉与自然分组"
  current_route: visual_source_launcher_not_visual_recovery_interface
  decision_status: recovered_from_pr
  selected_reason: "视觉与自然分组可以承担发现值得进入的视觉与自然来源，但 TopHub 今日热榜首页把所有来源渲染成文字卡片，因此它本身不是视觉恢复界面，只是视觉入口。"
  rejected_alternatives:
    - "因为来源是视觉型，就认定聚合页本身能视觉恢复"
    - "在首页读完标题列表"
    - "把视觉恢复变成新的信息流任务"
  route_reason: "来源是视觉型，不等于聚合页本身具有视觉性。正确用法是一次只从中选择一个来源，进入原页看图。"
  reuse_rule: "视觉源的判断必须进入原页，看图像、版式、编辑与上下文；不能只读 TopHub 文字卡片标题。"
  evidence_locator: "docs/ops/tophub-visual-nature-audit-20260704.md#结论; #使用规则"
  next_action: "保持为视觉来源启动器。"
```

---

## 9. 视觉与自然各节点怎样分工

```yaml
- source_id: "北京天文馆每日一图"
  current_route: keep_visual_nature_astronomy_object_entry
  decision_status: recovered_from_pr
  selected_reason: "标题集中在星系、星云、行星、彗星与天文现象，主题边界稳定，是当前最接近每天进入一个具体视觉对象的节点。"
  rejected_alternatives:
    - "因 TopHub 首页文字化而取消"
  route_reason: "其价值要在原页具体图像中实现，不在 TopHub 标题卡片中实现。"
  reuse_rule: "想看宇宙时优先进入北京天文馆或 NASA 原页。"
  evidence_locator: "docs/ops/tophub-visual-nature-audit-20260704.md#当前成立的部分"
  next_action: "保留。"

- source_id: "NASA 每日星球"
  current_route: keep_visual_nature_space_and_earth_image_candidate
  decision_status: recovered_from_pr
  selected_reason: "能提供地球、航天任务、空间影像和任务人物等候选。"
  rejected_alternatives:
    - "只因首页英文标题而否定"
    - "自动与 Folo 的 NASA APOD 视为完全相同"
  route_reason: "是否形成视觉恢复取决于进入原页后的图像，而不是首页英文标题列表；TopHub NASA 与 Folo APOD 的媒介呈现和使用场景不能自动等同。"
  reuse_rule: "NASA 相关源要区分标题扫描、图像原页、RSS 体验和 Folo 媒介呈现。"
  evidence_locator: "docs/ops/tophub-visual-nature-audit-20260704.md#当前成立的部分; 资源/雷达/TopHub 能力地图.md#视觉摄影与每日图片"
  next_action: "保留。"

- source_id: "中国国家地理网｜热度榜"
  current_route: keep_visual_nature_geography_wander
  decision_status: recovered_from_pr
  selected_reason: "羌塘、喀喇昆仑、西藏、日出和海等主题适合自然与地理漫游。"
  rejected_alternatives:
    - "把热门等同于摄影质量或内容新鲜度"
  route_reason: "它是站内热度榜，热门不等于摄影质量或内容新鲜度，但仍能作为自然地理漫游入口。"
  reuse_rule: "自然地理源可以作漫游，但评价内容质量时要进入原文和图片。"
  evidence_locator: "docs/ops/tophub-visual-nature-audit-20260704.md#当前成立的部分"
  next_action: "保留。"

- source_id: "iDaily 每日环球视野"
  current_route: keep_visual_news_and_picture_narrative
  decision_status: recovered_from_pr
  selected_reason: "能从全球事件、自然、航天、文化和城市中提供图像报道入口。"
  rejected_alternatives:
    - "当作纯自然恢复源"
  route_reason: "它更接近视觉新闻与图片叙事，而不是纯自然恢复，因此同时带有公共事件和时事负荷。"
  reuse_rule: "iDaily 适合看全球图片报道，不适合状态很低时当作纯恢复。"
  evidence_locator: "docs/ops/tophub-visual-nature-audit-20260704.md#当前成立的部分"
  next_action: "保留。"

- source_id: "胶片的味道"
  current_route: keep_with_observation_photography_narrative_not_only_learning_source
  decision_status: recovered_from_pr
  selected_reason: "仍可能提供摄影叙事、城市生活、胶片器材和画幅方法。"
  rejected_alternatives:
    - "仅凭情绪化标题立即删除"
    - "把它当稳定摄影学习来源"
  route_reason: "视觉实测中发现它标题大量采用爱情、遗憾、孤独和城市怀旧式文案，真正摄影器材、胶片知识和拍摄方法只占部分；后续摄影 closure 又修正为不能仅凭标题判断漂移。最终保留，但先看具体文章类型。"
  reuse_rule: "摄影叙事源不能只凭标题风格定性；也不能因名称自动当摄影学习源。进入原文看作品结构和器材/方法比例。"
  evidence_locator: "docs/ops/tophub-visual-nature-audit-20260704.md#当前降级或偏离的部分; docs/ops/tophub-entertainment-photography-closure-20260705.md"
  next_action: "保留观察。"

- source_id: "果壳物种日历"
  current_route: keep_with_observation_nature_topic_candidate
  decision_status: recovered_from_pr
  selected_reason: "仍可用于自然话题发现。"
  rejected_alternatives:
    - "继续按旧名称理解为每日一种物种"
    - "因偏离原名立即删除"
  route_reason: "当前内容主要是环境新闻、动物趣闻、气候事件和读者来信，不再像严格意义上的每日一种物种或自然观察日历。保留，但降级为自然议题候选。"
  reuse_rule: "节点名称不能覆盖实际内容漂移。自然源要按当前内容角色判断。"
  evidence_locator: "docs/ops/tophub-visual-nature-audit-20260704.md#当前降级或偏离的部分"
  next_action: "保留观察。"
```

---

## 10. 视觉与自然未来槽位：缺口不是立即新增清单

```yaml
- source_id: "视觉与自然未来槽位"
  current_route: gap_notes_not_immediate_subscription
  decision_status: recovered_from_pr
  selected_reason: "实测暴露出缺口：首页或原页能直接呈现高质量图片的来源，有稳定编辑判断的摄影作品与摄影史来源，植物、动物、地貌和季节观察长期来源，城市空间、建筑与公共设计视觉来源，低文字负担、适合短暂恢复注意力的图像源。"
  rejected_alternatives:
    - "立即新增一批图片站、壁纸站或设计站"
    - "把缺口清单当订阅清单"
  route_reason: "这些只是容器暴露出的缺口，不是立即新增动作。后来娱乐和设计复核分别加入 CNU、保留胶片、提出 designboom 复查，都是沿着这些缺口逐步补，而不是批量扩容。"
  reuse_rule: "缺口记录要等具体来源证据。缺口不是订阅命令。"
  evidence_locator: "docs/ops/tophub-visual-nature-audit-20260704.md#未来可补的来源槽位; docs/ops/tophub-design-final-closure-20260705.md"
  next_action: "按证据复查。"
```

---

## 11. 本补充 ledger 的复用规则

1. 旧题材不等于节点失效；低频告警源不能因低更新直接删除。
2. 只有标题长期停留旧事件且用户确认不再更新，才取消。
3. 取消失效节点不自动补位；有效节点优先于总量。
4. 作品年份、历史题材、纪录片年份不能当作抓取失效证据。
5. 设计目录不缺瀑布流，缺编辑判断、材料、工艺、空间与设计文化。
6. `designboom` 是 Folo 复查候选，不是已订阅事实，也不进 TopHub 排序。
7. 站酷、Dribbble、Behance、500px、优设、UI 中国等按任务，不进入持续流。
8. 视觉与自然是视觉来源启动器，不是视觉恢复界面本身。
9. 视觉源要进入原页看图像与上下文，不能只读 TopHub 标题卡。
10. 未来槽位是缺口记录，不是订阅命令。

---

## 12. 仍需继续追回

失效节点、设计目录与视觉自然边界已补成 ledger。

继续待办：

- 等用户本地提交 master checklist 与 coverage audit 后，再统一登记本文件和此前新增 ledger；
- 旧 ChatGPT 对话追索所有 partial 条目；
- 如需继续，可补 `TopHub 账号与通知审计 / App 与帮助 / Widgets / 首页与功能审计` 等剩余操作性文件。 
