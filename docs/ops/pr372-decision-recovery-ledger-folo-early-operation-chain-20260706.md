# PR #372 判断链追回：Folo 早期操作链与复查边界（2026-07-06）

## 0. 边界

本文件追回 PR #372 中 Folo 早期清理、直接 RSS 替换、分类重组、引导式复查协议、AI 目录交接与订阅系统收口形成的判断链。

纳入文件：

- `docs/ops/folo-subscription-diff-20260703.md`
- `docs/ops/folo-subscription-route-review-v2.md`
- `docs/ops/tophub-folo-guided-review-20260704.md`
- `docs/ops/tophub-folo-handoff-after-ai-20260704.md`
- `docs/ops/subscription-system-closeout-20260704.md`
- `资源/雷达/Folo 分类重组.md`
- `资源/雷达/订阅源账本.yml`
- `资源/雷达/订阅系统.md`
- `资源/雷达/信息流最终分工.md`

边界说明：

- 本文件不是重新审查当前 27 个 Folo 来源。
- 本文件不是 Folo 复查日的执行结果；下一次系统复查仍是 2026-07-17。
- 本文件不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。
- 本文件追回的是：Folo 为什么从 45 收束到 27、为什么替换直接 RSS、为什么分成五类注意力用途、为什么观察期候选不等于新增、为什么不把 TopHub 候选或 AI 候选立即写入真实 Folo 账本。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 总规则：Folo 收长期关系，不收公共雷达

```yaml
- source_id: "Folo 早期清理整体"
  current_route: reduced_to_27_long_term_relationships_and_attention_modes
  decision_status: recovered_from_pr
  selected_reason: "Folo 原始订阅 45 项，经过两轮清理和直接 RSS 替换后变为 27 项。它保留少量、持续、独特的一手关系、深读与多媒体原貌。"
  rejected_alternatives:
    - "把 Folo 当 TopHub 的另一份公共雷达"
    - "用 Folo 免费额度填满来源"
    - "让 RSSHub 承担所有来源"
    - "把未读数当待办债务"
    - "把候选来源立即写成真实订阅"
  route_reason: "Folo 的职责不是热榜、快讯、报告库、天气日提示、财经日历或公共话题扫描。TopHub 承担公共注意力、榜单、跨领域发现、结构观察、官方行动入口和少量精确追踪；Folo 只保留少量长期关系。"
  reuse_rule: "任何新来源进入 Folo 前，必须证明它是少量、持续、独特、值得保留原媒介形态的关系；否则回 TopHub、原平台、按需或沃壤待机。"
  evidence_locator: "docs/ops/folo-subscription-diff-20260703.md#总体结果; docs/ops/subscription-system-closeout-20260704.md#系统分工"
  next_action: "2026-07-17 复查观察期。"
```

---

## 2. 从 45 到 27：为什么清理，不继续删除

```yaml
- source_id: "Folo 45 → 27"
  current_route: cleaned_then_observation_period
  decision_status: recovered_from_pr
  selected_reason: "清理后总订阅从 45 降到 27，RSSHub 从 33 降到 15，直接 RSS 保持 10，第三方桥接保持 2。"
  rejected_alternatives:
    - "继续无差别删除"
    - "清理阶段结束后继续追求更低数量"
    - "因未读多继续删"
  route_reason: "差异记录明确：清理阶段已经完成，下一阶段转为分类重组和观察，不再继续无差别删除。订阅系统收口也明确复查前不继续删除来源、不导入新的 OPML、不把未读数当成待办债务。"
  reuse_rule: "系统清理要有停止条件。达到可观察状态后进入复查期，而不是用继续删除来替代判断。"
  evidence_locator: "docs/ops/folo-subscription-diff-20260703.md#当前实际结果; #总体结果; docs/ops/subscription-system-closeout-20260704.md#Folo最终状态"
  next_action: "维持 27，直到复查日。"
```

---

## 3. 直接 RSS 替换：减少 RSSHub 依赖，不改变来源角色

```yaml
- source_id: "Hacker News"
  current_route: direct_rss_replacement_observation_source
  decision_status: recovered_from_pr
  selected_reason: "原订阅使用 RSSHub 路由，存在不必要的桥接依赖。"
  rejected_alternatives:
    - "继续使用 rsshub.app/hackernews"
    - "删除 Hacker News"
    - "把替换误写成新增来源"
  route_reason: "删除 RSSHub 地址，新增并验证官方/直接 RSS `https://news.ycombinator.com/rss`。这是同一来源的订阅通道优化，不是新增来源，也不改变其观察期身份。"
  reuse_rule: "有稳定直接 RSS 时优先直接 RSS；通道替换不等于来源新增。"
  evidence_locator: "docs/ops/folo-subscription-diff-20260703.md#Hacker News; docs/ops/folo-subscription-route-review-v2.md#05观察期"
  next_action: "观察期复查，没有真实打开或独特价值则退出常驻。"

- source_id: "NASA APOD"
  current_route: direct_rss_replacement_visual_wander_source
  decision_status: recovered_from_pr
  selected_reason: "原订阅使用 RSSHub 路由，NASA APOD 有直接 RSS。"
  rejected_alternatives:
    - "继续使用 rsshub.app/nasa/apod"
    - "因 TopHub 中也有 NASA 每日星球就直接删除 Folo APOD"
    - "把替换误写成新增来源"
  route_reason: "删除 RSSHub 地址，新增并验证 `https://apod.nasa.gov/apod.rss`。这是通道优化。TopHub NASA 与 Folo APOD 可能相近，但媒介呈现和使用场景不同，是否重复要到视觉/漫游复查时判断。"
  reuse_rule: "视觉来源是否重复，要比较呈现媒介和实际打开场景，而不只比较名称。"
  evidence_locator: "docs/ops/folo-subscription-diff-20260703.md#NASA APOD; docs/ops/folo-subscription-route-review-v2.md#03漫游"
  next_action: "漫游组复查。"
```

---

## 4. 第二波删除 6 项：为什么移出 Folo，但不等于无价值

```yaml
- source_id: "Lifehacker"
  current_route: removed_from_folo_to_tophub_or_on_demand
  decision_status: recovered_from_pr
  selected_reason: "原来可能承担泛生活技巧和工具提示。"
  rejected_alternatives:
    - "继续占 Folo 持续未读位"
    - "把生活技巧流当深读来源"
  route_reason: "第二波删除中移出。其角色更像泛技巧、消费和工具提示，适合按需或由 TopHub / 浏览器搜索承担，不适合 Folo 免费计划里的少量长期关系。"
  reuse_rule: "泛生活技巧源若没有长期作者关系和稳定高价值，不占 Folo。"
  evidence_locator: "docs/ops/folo-subscription-diff-20260703.md#第二波精确差异; docs/ops/subscription-system-closeout-20260704.md#系统分工"
  next_action: "按需。"

- source_id: "Cyber Security News"
  current_route: removed_from_folo_security_news_on_demand_or_tophub
  decision_status: recovered_from_pr
  selected_reason: "原来可能承担安全新闻快讯。"
  rejected_alternatives:
    - "继续用安全新闻流占 Folo"
    - "把泛安全快讯等同于安全研究"
  route_reason: "第二波删除中移出。安全深读已由 darkreading、Krebs on Security 等承担，TopHub 科技雷达另有先知社区等安全入口；泛安全快讯更适合按具体 CVE、产品、供应链事件或任务查询。"
  reuse_rule: "安全源分研究、威胁情报、漏洞数据库、快讯、论坛和任务工具。泛快讯不默认进 Folo。"
  evidence_locator: "docs/ops/folo-subscription-diff-20260703.md#第二波精确差异; docs/ops/subscription-system-closeout-20260704.md#Folo最终状态"
  next_action: "按需。"

- source_id: "领研 论文 计算机"
  current_route: removed_from_folo_paper_index_on_demand
  decision_status: recovered_from_pr
  selected_reason: "原来可能承担计算机论文索引。"
  rejected_alternatives:
    - "把论文索引长期放 Folo"
    - "把论文流当每日未读"
  route_reason: "第二波删除中移出。论文、模型、数据集和研究目录应按具体研究问题、关键词、作者、引用链或项目任务调用，不适合制造 Folo 未读流。"
  reuse_rule: "论文索引按研究任务调用；需要长期跟随时优先具体作者、实验室、会议、领域或项目。"
  evidence_locator: "docs/ops/folo-subscription-diff-20260703.md#第二波精确差异; docs/ops/tophub-folo-handoff-after-ai-20260704.md#最近完成人工智能"
  next_action: "按需。"

- source_id: "中国气象局 每日天气提示"
  current_route: removed_from_folo_weather_alert_on_demand_or_task
  decision_status: recovered_from_pr
  selected_reason: "原来可能承担每日天气提示。"
  rejected_alternatives:
    - "把每日天气提示当长期 Folo 阅读源"
    - "把天气信息变成未读债务"
  route_reason: "第二波删除中移出。天气是地点、日期和出行任务触发的信息，不适合在 Folo 中形成泛阅读流。需要天气时用天气工具、官方预警或具体城市查询。"
  reuse_rule: "天气、预警和出行信息按地点和任务触发，不做常规未读源。"
  evidence_locator: "docs/ops/folo-subscription-diff-20260703.md#第二波精确差异"
  next_action: "按需。"

- source_id: "东方财富网 策略报告"
  current_route: removed_from_folo_financial_report_index_on_demand
  decision_status: recovered_from_pr
  selected_reason: "原来可能承担投资策略报告索引。"
  rejected_alternatives:
    - "把券商/策略报告库放 Folo"
    - "把策略报告当行动建议"
  route_reason: "第二波删除中移出。策略报告和研报索引需要回原券商、机构或报告发布方核验，不应成为 Folo 持续流，也不应直接触发投资行为。"
  reuse_rule: "财经报告库按具体研究任务调用；不直接转为投资动作。"
  evidence_locator: "docs/ops/folo-subscription-diff-20260703.md#第二波精确差异; docs/ops/subscription-system-closeout-20260704.md#系统分工"
  next_action: "按需。"

- source_id: "财经日历 华尔街见闻"
  current_route: removed_from_folo_calendar_tool_on_demand
  decision_status: recovered_from_pr
  selected_reason: "原来可能承担宏观数据、央行、财报和市场事件日历。"
  rejected_alternatives:
    - "把财经日历当 Folo 内容流"
    - "用日历更新制造未读"
  route_reason: "第二波删除中移出。财经日历是工具，不是阅读源；用于安排一周、核对央行、数据、财报、政策事件时按需打开。"
  reuse_rule: "日历型来源是工具，除非明确要推送具体事件，否则不进 Folo。"
  evidence_locator: "docs/ops/folo-subscription-diff-20260703.md#第二波精确差异"
  next_action: "按需。"
```

---

## 5. 分类重组：只移动，不删除，不导入 OPML

```yaml
- source_id: "Folo 分类重组 v2"
  current_route: move_only_attention_mode_reclassification
  decision_status: recovered_from_pr
  selected_reason: "清理完成后，需要把 27 项按注意力用途重组为 `01 必看触发`、`02 深度阅读`、`03 漫游`、`04 行动提醒`、`05 观察期`。"
  rejected_alternatives:
    - "继续删除"
    - "导入新的 OPML"
    - "把分类移动和新增订阅混在一起"
    - "按主题分类而不是注意力模式"
  route_reason: "route review 明确当前操作只移动、不删除；这一阶段不再导入 OPML，以免把已有订阅的分类移动与新增订阅混在一起。分类是注意力用途，不是价值等级。"
  reuse_rule: "重组阶段只能做一类动作。移动就是移动，新增就是新增，退订就是退订，不能混在同一轮里。"
  evidence_locator: "docs/ops/folo-subscription-route-review-v2.md#当前操作只移动不删除; #边界"
  next_action: "保持 5 类结构。"
```

---

## 6. 五类注意力用途：为什么放这里而不是那里

```yaml
- source_id: "01 必看触发"
  current_route: folo_trigger
  decision_status: recovered_from_pr
  selected_reason: "Google Developers Blog、Google DeepMind News、Obsidian Plugins 放在触发组，因为它们与工具链、开发者平台和 Obsidian 生态相关，可能触发具体工作中的阅读或行动。"
  rejected_alternatives:
    - "放入公共雷达"
    - "与 AI 周报 / Google 新闻 / Obsidian 社区泛流混在一起"
  route_reason: "触发组只处理与当前项目相关的低频、一手或工具链更新；不承担泛浏览。"
  reuse_rule: "来源只有可能触发明确项目、工具或工作流动作时，才进入必看触发。"
  evidence_locator: "docs/ops/folo-subscription-route-review-v2.md#01必看触发"
  next_action: "复查重复与实际触发。"

- source_id: "02 深度阅读"
  current_route: folo_deep_read
  decision_status: recovered_from_pr
  selected_reason: "小众软件、晚点长报道、Last Week in AI、darkreading、Krebs on Security、Steph Ango、有知有行放在深读组，因为它们更需要完整阅读、解释和理解。"
  rejected_alternatives:
    - "当每日新闻流"
    - "按未读数量清空"
  route_reason: "高质量深读不适合高频清空。一次只选少量文章。"
  reuse_rule: "深读来源必须真实值得完整阅读；不以清空未读为目标。"
  evidence_locator: "docs/ops/folo-subscription-route-review-v2.md#02深度阅读"
  next_action: "按打开率复查。"

- source_id: "03 漫游"
  current_route: folo_wander
  decision_status: recovered_from_pr
  selected_reason: "Nat Geo Photo of the Day、NASA APOD、Magnum Photos、中国爬楼联盟、街拍中国放在漫游组，因为它们提供视觉、摄影与陌生领域偶遇。"
  rejected_alternatives:
    - "把视觉漫游未读当债务"
    - "中国爬楼联盟与街拍中国长期无条件并存"
  route_reason: "漫游组只在状态合适时打开，过期可以标记已读；相近视觉源到复查期最多保留一个。"
  reuse_rule: "漫游源要比较媒介体验、审美关系和互补性；相近源不能无限并存。"
  evidence_locator: "docs/ops/folo-subscription-route-review-v2.md#03漫游; #观察期规则"
  next_action: "中国爬楼联盟与街拍中国最多保留一个。"

- source_id: "04 行动提醒"
  current_route: folo_reminder_private_names_omitted
  decision_status: recovered_from_pr
  selected_reason: "4 个来源名称只保留在本地清单，版本化记录保留数量与用途。"
  rejected_alternatives:
    - "把提醒源名称全部写入版本化账本"
    - "没有触发动作仍保留"
  route_reason: "行动提醒的去留标准不是名称是否写入，而是是否确实触发领取、观看或使用。"
  reuse_rule: "提醒源必须触发可执行动作；长期无动作则退出。"
  evidence_locator: "docs/ops/folo-subscription-route-review-v2.md#04行动提醒; #观察期规则"
  next_action: "复查实际触发。"

- source_id: "05 观察期"
  current_route: folo_trial
  decision_status: recovered_from_pr
  selected_reason: "Stack Overflow Blog、Google AI Developers、Logan Kilpatrick、News Minimalist、Hacker News 和 2 个名称省略来源进入观察期。"
  rejected_alternatives:
    - "立即写成长期常驻"
    - "观察期自动续期"
    - "Google AI Developers 与 Google Developers Blog 重复时双保留"
  route_reason: "两周后按打开率、重复度和独特价值判断。没有真实打开或独特价值就退出常驻。"
  reuse_rule: "trial 是有期限的试验，不是软性长期订阅。"
  evidence_locator: "docs/ops/folo-subscription-route-review-v2.md#05观察期; #观察期规则"
  next_action: "2026-07-17 去留判断。"
```

---

## 7. 名称省略 7 项：隐私边界不影响来源去留

```yaml
- source_id: "7 个名称省略来源"
  current_route: private_names_omitted_but_counted
  decision_status: recovered_from_pr
  selected_reason: "Folo 27 项中，20 个名称写入版本化账本，7 个名称当前省略，仅在本地保留名称。"
  rejected_alternatives:
    - "为了完整性强行写出名称"
    - "因为没有写出名称就判定不合格"
    - "把名称省略作为去留条件"
  route_reason: "route review 和 closeout 都明确：7 个来源名称当前省略，是否写入版本化账本不作为来源去留条件。它们的判断依据仍是用途、打开率、触发动作和重复度。"
  reuse_rule: "版本化透明度与隐私边界可以同时存在。可记录数量、用途和复查规则，不必暴露名称。"
  evidence_locator: "docs/ops/folo-subscription-route-review-v2.md#边界; docs/ops/subscription-system-closeout-20260704.md#Folo最终状态"
  next_action: "继续只在本地保留名称。"
```

---

## 8. 引导式复查：固定三个问题，避免页面点评散掉

```yaml
- source_id: "TopHub / Folo 引导式复查协议"
  current_route: review_protocol_for_each_tophub_tag_and_category
  decision_status: recovered_from_pr
  selected_reason: "后续判断始终围绕三个问题：TopHub 是否需要新增 / 替换 / 删除来源；是否需要新增精确追踪器且不混同通知机器人；在 TopHub 永久会员、Folo 免费计划约束下内容如何分工。"
  rejected_alternatives:
    - "只做零散页面点评"
    - "看完一页就立即改真实状态"
    - "把候选写成已执行"
    - "把追踪器和通知机器人合并判断"
  route_reason: "协议要求每个小标签处理时核对节点数、分页和节点性质，与当前 TopHub、Folo 和追踪器比较，给出 TopHub 常驻、替换候选、追踪、Folo 候选、按需访问或排除，并写入对应审查文件。大类结束后才统一回答三个问题并等待用户执行证据。"
  reuse_rule: "每次目录复查都必须先区分：建议执行、保留候选暂不执行、按需页面、明确排除；推荐不是现实。"
  evidence_locator: "docs/ops/tophub-folo-guided-review-20260704.md#固定要回答的三个问题; #每个小标签的处理方式; #每个大类结束后的固定收口"
  next_action: "作为所有后续目录审查协议。"
```

---

## 9. 用户控制页面路线，助手不能夺取浏览权

```yaml
- source_id: "用户控制页面选择与最终路由"
  current_route: user_controlled_review_flow
  decision_status: recovered_from_pr
  selected_reason: "用户掌握页面选择、浏览顺序与是否继续点击的主动权；用户逐个提供目录、详情页、榜单页或追踪结果页。"
  rejected_alternatives:
    - "助手擅自重新浏览整个目录"
    - "把还能继续看页面当答案"
    - "要求用户机械复制所有分页"
    - "以风险、目录规模或模型判断夺取选择权"
  route_reason: "助手负责判断页面能力、是否可加入分组、是否按需、哪些具体对象值得进一步点开，以及更适合 TopHub、追踪、Folo、按需还是排除；但不能替用户夺取浏览路线与最终选择。"
  reuse_rule: "助手可以提出下一步最有信息量的对象，但最终路线由用户控制；没有用户提供页面或文件，不主动假装已审核。"
  evidence_locator: "docs/ops/tophub-folo-guided-review-20260704.md#协作与判断位置; docs/ops/tophub-folo-handoff-after-ai-20260704.md#不可改变的协作规则"
  next_action: "继续遵守。"
```

---

## 10. 推荐、候选、执行证据三者分离

```yaml
- source_id: "推荐 / 候选 / 真实状态分离"
  current_route: evidence_before_state_update
  decision_status: recovered_from_pr
  selected_reason: "TopHub 与 Folo 审核会产生大量候选和推荐动作，但只有用户执行并提供页面证据或明确报告完成后，才能改写真实状态。"
  rejected_alternatives:
    - "在用户执行前改写最终节点总数"
    - "把候选写进真实订阅账本"
    - "Folo 没有变化也为了形式统一改账本"
  route_reason: "协议明确：在用户提供全部订阅节点及相关分组页面，或明确报告完成前，计划仍记录为待执行，不把推荐写成真实订阅状态，不提前改写最终节点总数。只有 Folo 真实发生改变时，才更新订阅源账本和 Folo 分类重组。"
  reuse_rule: "状态文件只记录已执行事实；候选和计划住 docs/ops 或候选清单。"
  evidence_locator: "docs/ops/tophub-folo-guided-review-20260704.md#第二步用户执行; #第三步页面证据确认后更新GitHub"
  next_action: "继续区分。"
```

---

## 11. Apple Podcasts 页面纠正：榜单查询页不是订阅节点

```yaml
- source_id: "Apple Podcasts 排行页面"
  current_route: on_demand_podcast_market_ranking_not_tophub_node
  decision_status: recovered_from_pr
  selected_reason: "页面可切换国家或地区、播客分类、音频或视频播客，按排名列出节目，主要动作是打开 iTunes / Apple Podcasts。"
  rejected_alternatives:
    - "把 iTunes Store 中国区播客 Top100 误记为可加入影游音乐的 TopHub 节点"
    - "把播客榜单当节目策展"
    - "从排行榜直接批量加入 Folo"
  route_reason: "该页面没有显示订阅此节点或加入自定义分组操作。排名只说明 Apple Podcasts 当前榜单位置，不等于节目质量、用户适配度或持续收听价值；发现具体节目后，还要查看近期选题、更新状态、独立站与 RSS，再判断是否进入 Folo。"
  reuse_rule: "榜单查询页是市场发现工具，不是订阅节点；节目关系必须逐个核验。"
  evidence_locator: "docs/ops/tophub-folo-guided-review-20260704.md#Apple Podcasts 排行页面纠正"
  next_action: "按需发现节目。"
```

---

## 12. AI 后交接：候选池不是执行清单

```yaml
- source_id: "AI 后交接 Folo 候选池"
  current_route: folo_review_candidates_not_executed
  decision_status: recovered_from_pr
  selected_reason: "AI 与科技目录完成后形成高优先级、条件和低优先级 Folo 候选池，包括 MacStories、Google Online Security Blog、The MIT Press Reader、夜航船夫、Quanta、科学空间、Simon Willison、Chip Huyen、AI as Normal Technology、Rest of World、Eugene Yan、SemiAnalysis、The Gradient 等。"
  rejected_alternatives:
    - "在 2026-07-17 前修改 Folo 账本"
    - "把 AI 候选全加入 Folo"
    - "把 AI 日报、快讯、热榜、Hugging Face / arXiv 洪流订阅为常驻"
    - "建立人工智能独立组"
  route_reason: "handoff 明确这些均未执行，不要在 2026-07-17 前修改 Folo 账本。AI 目录 18 页、215 节点结论是 TopHub 新增 0、替换 0、删除 0，不建人工智能组，不订阅 AI 日报、快讯、热榜或论文洪流，Folo 只保留少量直接作者候选。"
  reuse_rule: "候选池要到复查日统一比较打开率、重复度和角色互补；AI/Agent 关键词不能直接成为订阅理由。"
  evidence_locator: "docs/ops/tophub-folo-handoff-after-ai-20260704.md#Folo 候选池; #最近完成人工智能"
  next_action: "2026-07-17 统一复查。"
```

---

## 13. TopHub / Folo 分工在科技收口中的复用

```yaml
- source_id: "科技大类后 TopHub / Folo 分工"
  current_route: tophub_handles_public_and_structure_folo_no_duplicate
  decision_status: recovered_from_pr
  selected_reason: "科技大类执行后，The Register、Science Magazine、Counterpoint 和 Apple 维修计划由 TopHub 承担，不在 Folo 重复订阅；快讯、日报、报告库、限免和论坛不使用 RSSHub 常驻。"
  rejected_alternatives:
    - "把 TopHub 新增科技来源再复制到 Folo"
    - "把快讯、日报、报告库、限免、论坛做 RSSHub 常驻"
    - "用 Folo 解决所有科技雷达需求"
  route_reason: "TopHub 负责公共注意力、榜单、跨领域发现、结构观察、官方行动入口和少量精确追踪；Folo 负责少量、持续、独特的一手关系、深读与多媒体原貌。两者重复会浪费免费计划、制造未读堆积。"
  reuse_rule: "一个来源已经由 TopHub 承担公共雷达角色时，Folo 只有在需要完整原文、多媒体或长期作者关系时才重复订阅。"
  evidence_locator: "docs/ops/subscription-system-closeout-20260704.md#系统分工"
  next_action: "避免重复。"
```

---

## 14. 复查日规则：一进一出，首轮最多测试 3 个

```yaml
- source_id: "2026-07-17 Folo 复查"
  current_route: scheduled_review_not_background_execution
  decision_status: recovered_from_pr
  selected_reason: "观察期结束日期为 2026-07-17。"
  rejected_alternatives:
    - "复查前继续删除"
    - "复查前导入 OPML"
    - "一次性加入所有候选"
    - "把浏览器旧书签 48 条候选全部加入"
  route_reason: "收口记录明确：2026-07-17 先判断现有观察来源、重复 AI/开发源、视觉源和行动提醒是否退出；新来源实行一进一出；第一轮最多测试 3 个；候选都是复查队列，不是全部加入清单。"
  reuse_rule: "复查是有上限、有比较、有退出的动作。新来源必须挤掉旧来源或进入有限测试，不通过就回按需。"
  evidence_locator: "docs/ops/subscription-system-closeout-20260704.md#下一复查"
  next_action: "到复查日执行。"
```

---

## 15. 自动化不得接管来源路由

```yaml
- source_id: "Folo / TopHub 自动化边界"
  current_route: automation_disabled_for_source_routing
  decision_status: recovered_from_pr
  selected_reason: "TopHub 和 Folo 涉及外部页面状态、候选、计划、观察期、隐私省略来源和平台手工执行。"
  rejected_alternatives:
    - "自动退订"
    - "自动提交"
    - "自动修改来源状态"
    - "自动把候选写入真实账本"
  route_reason: "收口文件明确：外部页面状态需要页面证据确认，不能仅凭计划推断；自动化不得自动退订、自动提交或接管来源路由。"
  reuse_rule: "自动化只能辅助读取和记录，不能接管路由权、订阅动作或平台状态变更。"
  evidence_locator: "docs/ops/subscription-system-closeout-20260704.md#边界"
  next_action: "保持。"
```

---

## 16. 本补充 ledger 的复用规则

1. Folo 收长期关系，不收公共雷达。
2. 从 45 到 27 是清理结果；到达观察期后不继续无差别删除。
3. 直接 RSS 替换是通道优化，不是新增来源。
4. 被移出 Folo 的来源仍可由 TopHub、按需、原平台或沃壤待机承担。
5. 分类重组只移动，不删除，不导入 OPML。
6. 五类分类是注意力用途，不是价值等级。
7. 名称省略 7 项是隐私 / 本地边界，不影响去留判断。
8. 每个小标签必须回答 TopHub、追踪器、Folo 三个问题。
9. 用户控制页面选择和最终路线；助手不能因为目录大或风险高接管浏览权。
10. 候选、推荐和真实平台执行必须分离。
11. Apple Podcasts 这类榜单查询页不是 TopHub 订阅节点。
12. AI 候选池不是执行清单；2026-07-17 前不改 Folo 账本。
13. TopHub 已承担公共雷达时，Folo 不重复订阅。
14. 复查实行一进一出，第一轮最多测试 3 个。
15. 自动化不得自动退订、自动提交或接管来源路由。

---

## 17. 仍需继续追回

Folo 早期操作链与复查边界已补成 ledger。

继续待办：

- 等用户本地提交 master checklist 与 coverage audit 后，再统一登记本文件和此前新增 ledger；
- 旧 ChatGPT 对话追索所有 partial 条目；
- 若继续补仓库内残余，可转入 `设备生态关注说明 / 文化公共谈话者候选 / 资源总入口` 等非 TopHub 目录但与信息系统相关的文件。
