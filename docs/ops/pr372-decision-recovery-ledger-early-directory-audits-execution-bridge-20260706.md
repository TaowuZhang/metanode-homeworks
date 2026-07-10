# PR #372 判断链追回：早期目录审计与平台执行桥接（2026-07-06）

## 0. 边界

本文件追回 PR #372 中若干早期目录审计与后续平台执行核验之间的判断链。

纳入文件：

- `docs/ops/tophub-apple-appstore-directory-audit-20260704.md`
- `docs/ops/tophub-development-and-ai-directory-audit-20260704.md`
- `docs/ops/tophub-finance-directory-audit-20260704.md`
- `docs/ops/tophub-newspapers-directory-ledger-20260705.md`
- `docs/ops/tophub-finance-platform-execution-verified-20260705.md`
- `docs/ops/tophub-entertainment-final-reconciliation-20260705.md`
- `资源/雷达/TopHub 全部订阅顺序.md`

边界说明：

- 这些文件里出现过 70、71、73、80、84 等不同阶段数量。数量必须带时间边界。
- 当前真实 TopHub 基线仍以 `资源/雷达/TopHub 全部订阅顺序.md` 的 80 项为准。
- 本文件不重新审核单个目录，不更新控制文件，不合并 PR，不改 Ready。
- 本文件追回的是早期判断如何被后续专项 ledger 或平台执行证据承接。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 总规则：早期审计是判断起点，不是当前状态

```yaml
- source_id: "早期目录审计整体"
  current_route: historical_audit_bridge_to_later_execution_and_current_baseline
  decision_status: recovered_from_pr
  selected_reason: "早期审计文件记录了目录性质、候选判断、按需边界和当时平台数量；后续专项 closure、平台执行核验和全部订阅顺序更新了真实状态。"
  rejected_alternatives:
    - "用早期 70 / 71 / 73 / 84 覆盖当前 80"
    - "把早期候选写成已执行"
    - "把早期目录样本当完整逐页闭环"
  route_reason: "早期审计回答当时为什么不扩张、哪些可按需、哪些要后续细审；真实平台状态必须以执行核验和当前顺序为准。"
  reuse_rule: "任何 ledger 使用早期审计时，必须写清它是判断起点、候选池、阶段快照还是执行证据。当前状态不能从旧数量反推。"
  evidence_locator: "all included early audit and execution files"
  next_action: "作为时间边界规则。"
```

---

## 2. App Store：应用市场天气图，不是应用推荐源

```yaml
- source_id: "TopHub 苹果 / App Store 榜单目录"
  current_route: on_demand_app_market_weather_map
  decision_status: recovered_from_pr
  selected_reason: "该页面可按 iOS App、Mac App、播客、音乐、图书、电影等媒介，国家或地区，分类，iPhone / iPad，以及付费榜、免费榜、赚钱榜切换。它可以用于观察应用市场不同维度的短期变化。"
  rejected_alternatives:
    - "订阅中国区 iOS 全分类付费榜"
    - "订阅中国区 iOS 全分类免费榜"
    - "订阅中国区 iOS 全分类赚钱榜"
    - "开启新上榜通知"
    - "把 App Store 排名当应用质量或购买建议"
  route_reason: "付费榜更容易发现小团队产品但受促销和价格影响；免费榜主要反映大平台分发能力；赚钱榜主要受游戏内购、订阅和商业化能力影响。排名不等于质量、可信度或适合用户。当前已有科技雷达、macOS、Apple Silicon、Folo 官方源和影游音乐承担更重要变化。"
  reuse_rule: "先有具体问题，再用 App Store 榜单做市场样本；真正考虑安装前仍要看开发者、隐私、更新记录、评价与官网。"
  evidence_locator: "docs/ops/tophub-apple-appstore-directory-audit-20260704.md#页面性质; #三种榜单的含义; #当前决定"
  next_action: "按需。"
```

---

## 3. 开发目录早期审计：项目发现与采用判断分开

```yaml
- source_id: "TopHub 开发目录早期审计"
  current_route: on_demand_development_source_repository_with_hellogithub_monthly_as_valid_low_frequency_entry
  decision_status: recovered_from_pr
  selected_reason: "开发目录约 290 个节点、25 页，混合媒体、社区、教程、项目榜、产品资讯、转载、营销内容和旧节点。"
  rejected_alternatives:
    - "订阅 GitHub Trending Today / Weekly"
    - "订阅 CSDN 今日头条热点"
    - "订阅掘金、人人都是产品经理、开源中国热榜"
    - "把项目上榜速度当采用理由"
    - "立即把 HelloGitHub 加 Folo"
  route_reason: "GitHub Trending 的 stars 和上榜速度代表注意力，不代表成熟度、代码质量、安全性、适配度和维护前景。CSDN 有明显时间漂移。掘金、产品经理、开源中国有按需价值但标题激励、推广和二次报道较多。HelloGitHub 月刊通过内容与节奏验证：每月更新，以项目为单位，编辑化筛选，数量可控；但当时 RSS 未验证，所以更适合保留 TopHub 低频节点或后续再查 Folo。"
  reuse_rule: "开发来源先区分项目发现、社区讨论、教程包装、版本发布、官方文档和采用证据。采用前回仓库活动、issue、release、许可证、维护者和代码。"
  evidence_locator: "docs/ops/tophub-development-and-ai-directory-audit-20260704.md#开发目录; #HelloGitHub; #开发目录结论"
  next_action: "后续已由开发目录 ledger 承接。"
```

---

## 4. AI 目录早期审计：行业注意力层，不是知识源

```yaml
- source_id: "TopHub AI 目录早期审计"
  current_route: high_overlap_ai_attention_layer_on_demand
  decision_status: recovered_from_pr
  selected_reason: "AI 目录约 215 个节点、18 页，把中文 AI 媒体、开发者社区热榜、产品导航与工具推广、国际科技媒体、官方研究博客、课程、代充值、教程和 SEO 内容混在一起。"
  rejected_alternatives:
    - "订阅量子位、AIbase、36氪 AI、AI 产品榜"
    - "订阅 CSDN AI 热榜、AI 工具集、产品经理 AI 学习库"
    - "补入 Google AI Blog 的陈旧 TopHub 节点"
    - "把 MIT Technology Review 两个近似节点同时订阅"
    - "AI 目录整体进入 Folo"
  route_reason: "AI 目录主要功能是看 AI 圈正在炒什么，而不是建立可靠知识源。大量中文 AI 媒体重复模型发布、融资、Agent 产品、算力、机器人和公司新闻；工具集混有会员代充值和商业推广；Google AI Blog 节点可能陈旧；MIT Technology Review 更适合未来查官方 RSS；超神经仅作为 AI4Science 需求出现时的按需候选。"
  reuse_rule: "AI 来源按角色拆：官方 / 作者 / 工程实践 / 研究入口 / 行业注意力 / 工具营销。AI 或 Agent 关键词不能直接成为订阅理由。"
  evidence_locator: "docs/ops/tophub-development-and-ai-directory-audit-20260704.md#AI目录; #AI目录结论"
  next_action: "后续已由 AI reassessment ledger 承接。"
```

---

## 5. 财经目录早期审计：媒体议程、市场情绪和候选材料，不替代核验

```yaml
- source_id: "TopHub 财经目录早期审计"
  current_route: on_demand_finance_source_repository
  decision_status: recovered_from_pr
  selected_reason: "财经目录约 324 个节点、27 页，混合编辑型财经媒体、市场快讯、投资社区、行业报告、直播交易节目、数字货币和高频市场叙事。"
  rejected_alternatives:
    - "整体订阅财经目录"
    - "订阅股吧热榜"
    - "用日排行 / 热文 / 今日话题当重要性排序"
    - "因高频快讯建立持续盯盘习惯"
    - "把社区观点或报告直接当投资判断"
  route_reason: "第一财经、21 财经、华尔街见闻、CBNData、雪球、集思录等都有任务价值，但财经目录只能提供媒体议程、市场情绪和候选材料，不能替代公告、监管文件、统计数据和公司原始披露。当前数据与结构分组当时已覆盖官方数据、中文深度财经新闻、国际宏观评论和产业调查，因此早期审计不新增节点。"
  reuse_rule: "财经来源先区分编辑新闻、快讯、投资社区、报告、直播、Web3 和官方披露。事实核验回公告、监管、统计和公司原始材料。"
  evidence_locator: "docs/ops/tophub-finance-directory-audit-20260704.md#财经目录概况; #使用规则; #当前决定"
  next_action: "后续已由财经 page-by-page ledger 承接。"
```

---

## 6. 博物志 RSS：发现了 RSS，但只是待试听候选

```yaml
- source_id: "博物志 RSS"
  current_route: folo_listening_candidate_not_subscribed
  decision_status: recovered_from_pr
  selected_reason: "早期财经目录审计中确认了独立 RSS `https://bowuzhi.fm/feed/audio.xml`。"
  rejected_alternatives:
    - "确认 RSS 后立即订阅"
    - "把 RSS 可用性当节目关系成立"
    - "把候选写进 Folo 真实账本"
  route_reason: "RSS 可用只说明技术上可以订阅，不说明真实收听关系、近期选题质量、未读负担和与小宇宙 / Folo 分工已经成立。文件明确状态为 Folo 待试听候选，暂不订阅。"
  reuse_rule: "播客来源必须先有节目关系、近期质量和实际收听流程；RSS 可用性只是候选条件。"
  evidence_locator: "docs/ops/tophub-finance-directory-audit-20260704.md#博物志-RSS"
  next_action: "待试听。"
```

---

## 7. 报刊目录 ledger：抽样闭环，不是 873 个逐项审完

```yaml
- source_id: "TopHub 报刊目录抽样复核台账"
  current_route: representative_sampling_closure_not_full_exhaustive_review
  decision_status: recovered_from_pr
  selected_reason: "报刊目录总量大，不再要求逐页搬运，采用按类别代表性抽样：中央级、财经报、法制报各看热门第 1 页 12 个节点，合计逐项判断 36 个代表节点。"
  rejected_alternatives:
    - "逐页搬运 873 个节点"
    - "按报纸声望保留失效节点"
    - "把报纸本身值得查等同于 TopHub 抓取节点值得订阅"
    - "为报刊目录扩充订阅总量"
  route_reason: "审核策略是先看热门排序第 1 页，只在发现明确缺口或竞争候选时追加页面，目标是替换和压缩而不是扩充。文化教育报、健康报、都市报、日报/早报/晚报、按省份等因当前无明确缺口或重叠大而停止。"
  reuse_rule: "大型目录可用代表性抽样闭环，但必须写清抽样范围，不能把抽样说成全量逐项在线核验。"
  evidence_locator: "docs/ops/tophub-newspapers-directory-ledger-20260705.md#审核策略; #最终进度"
  next_action: "已由报刊 sampled closure ledger 承接。"
```

---

## 8. 报刊替换：新华每日电讯替代人民日报电子版

```yaml
- source_id: "人民日报电子版 → 新华每日电讯电子报"
  current_route: executed_public_temperature_replacement
  decision_status: recovered_from_pr_platform_verified
  selected_reason: "新华每日电讯样本保持当前更新，更适合作为官方广谱日报摘要；人民日报节点内容约停留在 2024 年。"
  rejected_alternatives:
    - "保留人民日报电子版失效节点"
    - "用人民日报海外版、经济日报、中国青年报、解放军报、健康时报替代"
    - "加入法治日报、检察日报或地方报作为公共广谱日报"
  route_reason: "中央级样本中新华每日电讯是强替换候选。财经报虽有每日经济新闻、21世纪经济报道、经济参考报等任务价值，但与数据与结构层重叠；法制报有效价值窄，适合具体法律任务，不承担全国广谱日报角色。平台执行已完成。"
  reuse_rule: "官方报刊源先看抓取是否仍更新，再看角色是广谱日报、财经任务、法律任务还是地方任务。权威品牌不能覆盖节点失效。"
  evidence_locator: "docs/ops/tophub-newspapers-directory-ledger-20260705.md#最终替换决定; #实际状态"
  next_action: "已执行。"
```

---

## 9. 财经平台执行核验：真实执行证据单独成层

```yaml
- source_id: "财经目录等量替换平台执行"
  current_route: platform_execution_verified
  decision_status: recovered_from_pr_platform_verified
  selected_reason: "用户提供 TopHub 全部订阅节点和数据与结构分组完整文本，核验显示两个旧财新排行榜已不存在，财新首页推荐和日经中文每日最新存在，且均进入数据与结构分组。"
  rejected_alternatives:
    - "只凭计划写成已执行"
    - "只凭单个分组截图推断全部平台状态"
    - "把财经替换与报刊替换混同"
  route_reason: "平台执行核验是独立事实层：已取消财新点击排行榜和评论排行榜，已订阅财新首页推荐与日经中文每日最新，数据与结构仍为 12，TopHub 总量仍为 84。替换依据是两个旧排行榜停止更新，不只是排行榜逻辑不理想。"
  reuse_rule: "任何平台动作完成后都要用全部订阅与分组文本核验；计划、建议和候选不能冒充执行。"
  evidence_locator: "docs/ops/tophub-finance-platform-execution-verified-20260705.md#核验结果; #状态结论"
  next_action: "已执行。"
```

---

## 10. 娱乐最终 reconciliation：目录审核到真实界面执行闭环

```yaml
- source_id: "娱乐大类最终总闭环"
  current_route: executed_entertainment_reconciliation
  decision_status: recovered_from_pr_platform_verified
  selected_reason: "娱乐大类逐页审核、节点新增、旧节点取消和分组排序已经全部完成。最终新增 10 个、取消 3 个，TopHub 由 73 增至 80，Folo 仍为 27。"
  rejected_alternatives:
    - "继续把娱乐动作视为待办"
    - "把娱乐目录继续当无限扩张仓库"
    - "把明星私生活、粉圈热搜、平台重复细分榜、历史累计榜、高奢商品连续流加入常驻"
  route_reason: "最终保留五类功能：作品发现、评论与解释、创作与产业、少量传播温度、视觉漫游与摄影叙事。新增华丽志、豆瓣电影影评、书格、豆瓣华语新碟、网易云原创榜、豆瓣全球口碑剧集、新京报娱乐、INDIENOVA、喷嚏乐影、CNU；取消豆瓣非虚构旧榜、网易云飙升榜、开眼日报；复核后保留胶片的味道。"
  reuse_rule: "娱乐源按作品发现、评论解释、产业创作、传播温度、视觉漫游分层；不追粉圈、私生活、平台切片和历史累计榜。执行闭环后只有实际内容质量变化才重开单源复核。"
  evidence_locator: "docs/ops/tophub-entertainment-final-reconciliation-20260705.md#最终动作; #结构判断; #最终结论"
  next_action: "已执行。"
```

---

## 11. 娱乐 Folo 边界：候选留到 2026-07-17，不改 27

```yaml
- source_id: "娱乐 Folo 候选"
  current_route: folo_review_candidates_not_executed
  decision_status: recovered_from_pr
  selected_reason: "娱乐最终 reconciliation 记录了 Folo 复查候选：基本読書、木遥的窗子、Yuko's Blog、Huiliu；观察：maxOS、Velas电波站。"
  rejected_alternatives:
    - "娱乐闭环后立即加入 Folo"
    - "把候选写成真实订阅"
    - "因 TopHub 娱乐扩到 80 就同步扩 Folo"
  route_reason: "文件明确 2026-07-17 复查前保持 Folo 27 项。候选只是复查队列，不是执行清单。"
  reuse_rule: "TopHub 已完成娱乐执行，不等于 Folo 自动扩容。Folo 候选必须到复查日按一进一出和打开率判断。"
  evidence_locator: "docs/ops/tophub-entertainment-final-reconciliation-20260705.md#Folo边界"
  next_action: "2026-07-17 复查。"
```

---

## 12. 本补充 ledger 的复用规则

1. 早期目录审计是判断起点，不是当前状态。
2. App Store 榜单是市场天气图，不是应用推荐源；榜单通知不启用。
3. GitHub Trending 是项目注意力，不是采用证据。
4. HelloGitHub 月刊通过内容与节奏验证，但 RSS 未验证时不直接进 Folo。
5. AI 目录是行业注意力层，不是可靠知识源；AI/Agent 关键词不能直接订阅。
6. 财经目录提供媒体议程、市场情绪和候选材料，不能替代事实核验与投资判断。
7. RSS 可用只说明可订阅，不说明节目关系成立。
8. 报刊目录采用代表性抽样闭环，不能说成 873 个逐项全量在线核验。
9. 平台执行核验必须单独成层，不能用计划冒充完成。
10. 娱乐最终 reconciliation 是执行闭环；后续只有内容质量变化才重开单源复核。
11. TopHub 执行完成不等于 Folo 自动扩容。

---

## 13. 仍需继续追回

早期目录审计与平台执行桥接已补成 ledger。

继续待办：

- 等用户本地提交 master checklist 与 coverage audit 后，再统一登记本文件和此前新增 ledger；
- 仓库内剩余可查项主要是娱乐若干单页 closure 是否已被前三份娱乐 ledger 全量覆盖；
- 旧 ChatGPT 对话追索所有 partial 条目仍待处理。
