# PR #372 判断链追回：财经逐页 324 节点（2026-07-06）

## 0. 边界

本文件追回 `TopHub > 财经` 目录 27 页、324 个节点逐项复审后的判断链。

纳入文件：

- `docs/ops/tophub-finance-directory-audit-20260704.md`
- `docs/ops/tophub-finance-page-by-page-ledger-20260705.md`
- `docs/ops/tophub-finance-reassessment-pages-01-10-20260705.md`
- `docs/ops/tophub-finance-reassessment-pages-11-20-20260705.md`
- `docs/ops/tophub-finance-reassessment-pages-21-27-20260705.md`
- `docs/ops/tophub-finance-final-closure-20260705.md`
- `docs/ops/tophub-finance-platform-execution-verified-20260705.md`

当前不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 财经目录整体：不是寻找更多行情入口

```yaml
- source_id: "TopHub > 财经 324 节点整体"
  current_route: closed_with_two_equal_replacements
  decision_status: recovered_from_pr
  selected_reason: "财经目录已完整审完 27 页、324 个节点。当前真正缺少的不是更多新闻、快讯或市场意见，而是日本与东亚产业 / 货币政策视角、仍在更新的财新编辑入口、长期投资与资产配置尺度、官方披露与专业研究工具、以及把 Web3 技术 / 安全 / 监管从币价交易流中拆出来的来源。"
  rejected_alternatives:
    - "新增更多 7×24 快讯"
    - "新增荐股、龙虎榜、资金流、盘前盘后和市场复盘"
    - "新增宽泛财经 / 股票 / Web3 / 比特币 / 宏观追踪器"
    - "把财经信息流当成投资动作生成器"
    - "把单个平台的组合调整当成用户投资指令"
  route_reason: "最终 TopHub 只做两项等量替换：取消两个失效财新排行榜，加入财新首页推荐和日经中文每日最新。TopHub 总量仍为 84，数据与结构仍为 12；Folo 实际新增 0。"
  reuse_rule: "以后遇到财经来源，先分六类：官方原始披露、结构分析、研究入口、市场温度、交易社区、Web3。只有官方披露和结构分析可能进入持续候选；研究入口按任务；市场温度最多临时；交易社区排除；Web3 必须拆技术 / 监管 / 安全 / 价格。"
  evidence_locator: "docs/ops/tophub-finance-page-by-page-ledger-20260705.md#核心原则; docs/ops/tophub-finance-final-closure-20260705.md#最终判断"
  next_action: "无新增动作；2026-07-17 复查有知有行。"
```

---

## 2. 两个最终 TopHub 替换

```yaml
- source_id: "财新网｜点击排行榜 / 财新网｜评论排行榜"
  current_route: retired_from_tophub_stale_nodes
  decision_status: recovered_from_pr
  selected_reason: "原本承担财新系中国政策、商业、金融和调查报道入口。"
  rejected_alternatives:
    - "继续保留点击排行榜"
    - "继续保留评论排行榜"
    - "同时新增第三个财新节点"
  route_reason: "两个旧节点被取消的直接事实依据是均已停止更新。排行榜逻辑是否理想只是次要问题；失效节点已经不能继续承担持续入口角色。最终不是增加财新，而是把两个失效流量榜压缩为一个仍在工作的编辑入口。"
  reuse_rule: "节点失效优先于偏好讨论；同一媒体多个榜单失效时，先找一个可工作的编辑入口，而不是继续维护失效热榜。"
  evidence_locator: "docs/ops/tophub-finance-final-closure-20260705.md#最终判断; #TopHub最终执行结果"
  next_action: "已取消。"

- source_id: "财新网｜首页推荐"
  current_route: tophub_data_and_structure
  decision_status: recovered_from_pr
  selected_reason: "替代两个失效财新排行榜，继续承担中国政策、商业、金融与调查报道的财新编辑入口。"
  rejected_alternatives:
    - "财新网｜最新文章"
    - "财新经济频道 / 金融频道 / 政经频道等陈旧映射"
    - "财新点击榜 / 评论榜继续保留"
  route_reason: "放入数据与结构第 3 位。首页推荐比停止更新的点击榜和评论榜更可用；但文件也明确不把“首页推荐”名称本身当精选证据，后续要根据真实内容体验观察噪声与重复。最终在首页推荐和最新文章中选择首页推荐，避免增加第三个财新节点。"
  reuse_rule: "媒体入口要避免同站多栏目叠加；编辑入口可先观察，但必须保留后续按真实使用复审的条件。"
  evidence_locator: "docs/ops/tophub-finance-reassessment-pages-11-20-20260705.md#第12页; docs/ops/tophub-finance-final-closure-20260705.md#财新网"
  next_action: "已加入；后续观察噪声。"

- source_id: "日经中文网｜每日最新"
  current_route: tophub_data_and_structure
  decision_status: recovered_from_pr
  selected_reason: "补足日本、东亚产业、供应链、能源、科技、日元、央行和产业政策视角；与用户在日本生活、全球资产配置和东亚经济观察相关。"
  rejected_alternatives:
    - "联合早报财经节点常驻"
    - "财联社环球 / BBC Business / Bloomberg Markets / FT中文扩容"
    - "更多市场快讯替代亚洲视角"
  route_reason: "放入数据与结构第 4 位，紧接财新之后、FT 中文之前。日经补日本与东亚结构；FT 中文补国际经济与商业解释；联合早报保留任务源，避免与日经和 FT 重叠。"
  reuse_rule: "区域视角来源要看是否补系统缺口。亚洲视角不是多一个新闻源，而是补日本、东亚供应链和货币政策。"
  evidence_locator: "docs/ops/tophub-finance-reassessment-pages-01-10-20260705.md#日经中文网每日最新; docs/ops/tophub-finance-final-closure-20260705.md#日经中文网"
  next_action: "已加入。"
```

---

## 3. 数据与结构中的现有财经 / 产业来源怎样分工

```yaml
- source_id: "FT中文网｜十大热门文章"
  current_route: keep_tophub_data_and_structure_no_expansion
  decision_status: recovered_from_pr
  selected_reason: "承担国际经济、全球政策、商业和市场解释。"
  rejected_alternatives:
    - "扩容更多 FT / 英文财经节点"
    - "用 Bloomberg / BBC Business / Wired Business 替代日常入口"
  route_reason: "保持热门文章入口，不扩容。FT 中文与日经、财新分工：财新看中国政策商业金融调查，日经看日本 / 东亚供应链与货币，FT 看国际经济与全球商业解释。"
  reuse_rule: "已有国际财经解释入口时，不因英文专业源质量高就扩容；专业英文源按任务调用。"
  evidence_locator: "docs/ops/tophub-finance-final-closure-20260705.md#FT中文网; docs/ops/tophub-finance-page-by-page-ledger-20260705.md#来源分工"
  next_action: "无；保留。"

- source_id: "第一财经｜汽车新闻"
  current_route: keep_tophub_data_and_structure_auto_industry
  decision_status: recovered_from_pr
  selected_reason: "继续承担汽车产业、市场结构、政策、安全、供应链、能源转型和消费结构入口。"
  rejected_alternatives:
    - "第一财经头条 / 排行 / 直播 / 视频 / 盘前必读"
    - "财联社汽车同时常驻"
  route_reason: "财经 final 明确第一财经只保留现有汽车新闻；其他第一财经内容已被财新、FT 中文、日经、央视栏目和公共新闻覆盖。财联社汽车与第一财经汽车高度重叠，不新增。"
  reuse_rule: "同一媒体只保留不可替代角色；汽车产业结构由结构源承担，车型和车主经验按任务。"
  evidence_locator: "docs/ops/tophub-finance-final-closure-20260705.md#第一财经; docs/ops/tophub-finance-reassessment-pages-11-20-20260705.md#第16页"
  next_action: "无；保留。"

- source_id: "Counterpoint Research｜最新见解"
  current_route: keep_tophub_data_and_structure_terminal_market_research
  decision_status: recovered_from_pr
  selected_reason: "智能手机、穿戴、物联网、晶圆代工和终端 AI 的产业数据角色清楚；与 Apple、vivo、手机厂商、Apple Silicon、AI 硬件、智能眼镜和消费电子市场相关。"
  rejected_alternatives:
    - "Canalys｜数据分析"
    - "多个报告目录常驻"
  route_reason: "继续保留。Canalys 页面停留在 2025 年，且角色已被 Counterpoint 覆盖；报告目录按任务调用。"
  reuse_rule: "市场研究源常驻要看更新、独立数据能力和长期项目相关性；陈旧页面即使品牌强也不新增。"
  evidence_locator: "docs/ops/tophub-finance-reassessment-pages-21-27-20260705.md#第21页; #第24页; docs/ops/tophub-finance-final-closure-20260705.md#当前数据与结构顺序"
  next_action: "无；保留。"
```

---

## 4. Folo 唯一复查候选：有知有行

```yaml
- source_id: "有知有行｜全部"
  current_route: folo_review_candidate_long_term_investing_behavior_and_asset_allocation
  decision_status: recovered_from_pr
  selected_reason: "首次补上长期投资、基金、保险、资产配置、组合再平衡、投资行为和数据解释角色。它不以每日价格变化为中心，能把市场波动放回组合、风险和生活。"
  rejected_alternatives:
    - "有知有行｜知行黑板报"
    - "有知有行｜知行小酒馆"
    - "更多市场媒体和每日快讯"
  route_reason: "进入 2026-07-17 Folo 复查候选，不在本轮修改真实 Folo。选择`全部`，因为黑板报偏组合调仓、基金表现和卖出信号，小酒馆内容宽于投资；三者重叠时，`全部`是最完整入口。必须明确不能把有知有行自己的组合动作当成用户投资指令。"
  reuse_rule: "长期投资来源要补时间尺度和行为框架，不是行情预测。平台自有产品和组合动作只能作为材料，不能替代用户投资决策。"
  evidence_locator: "docs/ops/tophub-finance-reassessment-pages-11-20-20260705.md#有知有行全部; docs/ops/tophub-finance-final-closure-20260705.md#Folo最终建议"
  next_action: "2026-07-17 Folo 复查。"
```

---

## 5. 结构候选为什么没有进入常驻

```yaml
- source_id: "经济观察网 / 21财经数读 / 财经网 / 财富中文 / 未央网 / 科创板日报深度"
  current_route: on_demand_structural_finance_and_industry_sources
  decision_status: recovered_from_pr
  selected_reason: "这些来源分别有公司产业政策长文、数据可视化、金融科技、商业管理、硬科技产业和融资等结构价值。"
  rejected_alternatives:
    - "全部加入数据与结构"
    - "替代财新 / 日经 / FT 中文"
    - "进入 Folo"
  route_reason: "它们进入 A2 或 B+ 任务池，但没有推翻最终架构。当前数据与结构已经由国家统计局、财新、日经、FT、央视调查、36氪出海、第一财经汽车、Counterpoint 和华丽志分工覆盖；A2 来源按具体公司、产业、政策、消费、金融科技或硬科技主题调用。"
  reuse_rule: "结构来源质量高也要比较系统缺口和重复度。A2 不是失败，而是任务源。"
  evidence_locator: "docs/ops/tophub-finance-reassessment-pages-01-10-20260705.md#阶段性候选池; docs/ops/tophub-finance-reassessment-pages-11-20-20260705.md#累计240节点后的候选池; docs/ops/tophub-finance-final-closure-20260705.md#来源分工"
  next_action: "按任务。"

- source_id: "联合早报财经节点"
  current_route: on_demand_asia_chinese_finance_source
  decision_status: recovered_from_pr
  selected_reason: "能补新加坡、东南亚、亚洲华语财经、亚洲公司和区域市场视角。"
  rejected_alternatives:
    - "联合早报全球财经常驻"
    - "联合早报中国财经常驻"
    - "联合早报财经即时常驻"
    - "与日经中文和 FT 中文并存扩容"
  route_reason: "三个早报财经节点只能留一个，但 final 选择任务源。日经已经承担日本与东亚供应链 / 货币视角，FT 中文承担国际经济和商业解释；联合早报按新加坡、东南亚或亚洲华语财经任务调用。"
  reuse_rule: "区域来源要避免亚洲视角重复堆叠；如果已有日经和 FT，联合早报保留任务入口。"
  evidence_locator: "docs/ops/tophub-finance-reassessment-pages-11-20-20260705.md#本批最重要的新结论; docs/ops/tophub-finance-final-closure-20260705.md#联合早报"
  next_action: "按需。"

- source_id: "Bloomberg / BBC Business / Wired Business / Finviz"
  current_route: professional_english_task_sources_not_daily_feed
  decision_status: recovered_from_pr
  selected_reason: "这些英文来源有全球宏观、公司、科技商业、市场聚合和专业任务价值。"
  rejected_alternatives:
    - "替代 FT 中文 / 日经中文"
    - "进入 TopHub 常驻"
    - "进入普通 Folo 未读"
  route_reason: "Bloomberg 受付费和高频限制；BBC Business 与 FT / 日经 / Bloomberg / 联合早报重叠；Wired Business 更像科技社会深读；finviz 是股票研究工具，不是阅读源。"
  reuse_rule: "英文专业源按公司、市场、行业或研究任务调用；工具和付费源不因为质量高就常驻。"
  evidence_locator: "docs/ops/tophub-finance-reassessment-pages-11-20-20260705.md#第15页; docs/ops/tophub-finance-reassessment-pages-21-27-20260705.md#第26页; #第27页"
  next_action: "按任务。"
```

---

## 6. 财经日历与官方披露工具组

```yaml
- source_id: "华尔街见闻｜财经日历"
  current_route: worang_direct_tool_link_not_tophub_feed
  decision_status: recovered_from_pr
  selected_reason: "最终财经日历候选。用于安排一周、核对央行、宏观数据、财报和政策事件。"
  rejected_alternatives:
    - "东方财富财经日历"
    - "财联社提醒 / 投资日历"
    - "把财经日历变成每日订阅"
  route_reason: "路由为沃壤直接工具链接，不进入 TopHub 每日订阅。财经日历是查询工具，不是内容流；只有安排一周或核对事件时打开。"
  reuse_rule: "日历类来源优先工具化：覆盖全球宏观、可按周查看、含数据 / 央行 / 财报 / 政策事件、能筛选、不混入交易建议。"
  evidence_locator: "docs/ops/tophub-finance-reassessment-pages-11-20-20260705.md#财经日历比较池; docs/ops/tophub-finance-final-closure-20260705.md#财经日历"
  next_action: "写入沃壤工具链接。"

- source_id: "官方披露与市场制度工具组"
  current_route: worang_direct_official_tools_not_daily_feed
  decision_status: recovered_from_pr
  selected_reason: "巨潮资讯、上交所、深交所、北交所、全国股转系统、中国货币网、FRB Monetary Policy、中国会计视野网等提供公司公告、交易所规则、债券 / 货币市场公告、央行政策和专业规则核验。"
  rejected_alternatives:
    - "官方规则入口进入每日 TopHub"
    - "用公告二次摘要替代官方披露"
    - "用财联社 / 东方财富公告摘录替代原始入口"
  route_reason: "写入沃壤工具链接，不制造未读。官方入口在具体公司、规则、货币政策、债券、LPR、交易制度、会计审计、再融资和市场制度问题出现时直接查。部分 TopHub 映射旧、把导航当内容或更新异常，因此更应直接用官方站。"
  reuse_rule: "官方披露优先于二次摘要；官方工具是核验入口，不是每日阅读入口。"
  evidence_locator: "docs/ops/tophub-finance-reassessment-pages-21-27-20260705.md#第21页; #第22页; #第23页; docs/ops/tophub-finance-final-closure-20260705.md#官方与研究工具组"
  next_action: "写入沃壤工具链接。"

- source_id: "国家金融与发展实验室"
  current_route: worang_low_frequency_research_index
  decision_status: recovered_from_pr
  selected_reason: "季报、周报、学术报告和研究评价覆盖宏观、金融、数字资产、制度研究、居民收入、物价、地方债和国际金融，是低频高密度研究源。"
  rejected_alternatives:
    - "进入 TopHub 日常流"
    - "拆分多个实验室栏目订阅"
  route_reason: "更适合写入沃壤研究源，在宏观、数字资产、金融风险和监管问题出现时调用；不承担每日阅读或未读清零。"
  reuse_rule: "低频研究机构不是新闻源；按专题调用，不按栏目拆分订阅。"
  evidence_locator: "docs/ops/tophub-finance-reassessment-pages-11-20-20260705.md#第17页; #当前架构判断; docs/ops/tophub-finance-reassessment-pages-21-27-20260705.md#第23页"
  next_action: "沃壤研究索引。"
```

---

## 7. Web3：技术、监管、安全与币价交易流分开

```yaml
- source_id: "Foresight News｜文章"
  current_route: on_demand_web3_chinese_structural_longform
  decision_status: recovered_from_pr
  selected_reason: "最终中文结构长文首选，角色是稳定币、支付、监管、协议、安全和行业结构长文。相比快讯流更像编辑部原创与专题。"
  rejected_alternatives:
    - "PANews 快讯 / 精选 / 最新快讯"
    - "Odaily 7×24 / 每日头条 / 文章日榜"
    - "链捕手 / ChainFeeds 同时常驻"
    - "进入当前 TopHub 或 Folo"
  route_reason: "保持按需，不进入当前 TopHub 或 Folo。Web3 对用户是职业方向，但不能把币价、杠杆、巨鲸、清算和价格预测变成持续源。中文结构长文按具体协议、监管、安全或行业问题调用。"
  reuse_rule: "Web3 中文源要分结构长文、快讯、链上数据、交易叙事和安全研究；最多选结构角色，不接收价格流。"
  evidence_locator: "docs/ops/tophub-finance-reassessment-pages-11-20-20260705.md#Web3专业候选; docs/ops/tophub-finance-final-closure-20260705.md#Web3"
  next_action: "按需。"

- source_id: "慢雾科技｜技术研究 / 漏洞披露"
  current_route: worang_web3_security_professional_source
  decision_status: recovered_from_pr
  selected_reason: "Web3 安全事件原始复盘，覆盖供应链、凭据窃取、协议攻击和链上安全研究，具有原始性。"
  rejected_alternatives:
    - "慢雾科技｜公司新闻"
    - "Web3 快讯流替代安全研究"
    - "进入普通财经 TopHub"
  route_reason: "写入沃壤专业来源，不选择慢雾公司新闻。它属于安全和技术研究，不是市场价格流，也不要求作为普通财经未读。"
  reuse_rule: "Web3 安全源与 Web3 新闻源分开；技术研究和漏洞披露按安全事件、协议、钱包、供应链或项目任务调用。"
  evidence_locator: "docs/ops/tophub-finance-reassessment-pages-01-10-20260705.md#第8页; #第10页; docs/ops/tophub-finance-final-closure-20260705.md#Web3"
  next_action: "沃壤专业来源。"

- source_id: "CoinDesk｜Today"
  current_route: on_demand_web3_english_original_reporting
  decision_status: recovered_from_pr
  selected_reason: "英文原始报道、监管、代币化和安全内容有价值。"
  rejected_alternatives:
    - "U.Today"
    - "Bitcoin News"
    - "CryptoPanic"
    - "进入常驻英文位置"
  route_reason: "价格新闻占比仍高，不占常驻英文位置；只有英文原始报道、监管或具体行业事件需要时按需。"
  reuse_rule: "英文 Web3 媒体也要控制价格稿比例；有原始报道价值不等于持续订阅。"
  evidence_locator: "docs/ops/tophub-finance-reassessment-pages-11-20-20260705.md#第14页; docs/ops/tophub-finance-final-closure-20260705.md#Web3"
  next_action: "按需。"

- source_id: "Web3 高频价格与交易流"
  current_route: rejected
  decision_status: recovered_from_pr
  selected_reason: "无常驻价值。"
  rejected_alternatives:
    - "BBX 全部节点"
    - "PANews 快讯与最新快讯"
    - "Odaily 7×24"
    - "BitKan"
    - "CryptoPanic"
    - "TechFlow 7×24"
    - "Followin"
    - "U.Today"
    - "Bitcoin News"
    - "巨鲸、杠杆、清算、爆仓、支撑位、价格预测流"
  route_reason: "这些来源高比例混入币价预测、杠杆、巨鲸仓位、代币解锁、交易所上新、ETF、Meme、社交喊单和价格情绪。与用户 Web3 职业方向有关的是协议、监管、基础设施和安全，不是分钟级价格噪声。"
  reuse_rule: "Web3 职业方向不等于交易流。交易流只在具体研究市场情绪或事件时临时看，不能成为长期未读。"
  evidence_locator: "docs/ops/tophub-finance-reassessment-pages-11-20-20260705.md#已明确淘汰的Web3流; docs/ops/tophub-finance-reassessment-pages-21-27-20260705.md#明确不进入; docs/ops/tophub-finance-final-closure-20260705.md#Web3"
  next_action: "排除。"
```

---

## 8. 交易社区、快讯、研报索引与交易工具为什么不常驻

```yaml
- source_id: "7×24 快讯与市场温度流"
  current_route: rejected_or_temporary_query_only
  decision_status: recovered_from_pr
  selected_reason: "能快速显影宏观、能源、地缘、公司公告和市场波动。"
  rejected_alternatives:
    - "华尔街见闻、财联社、金十、新浪财经、东方财富、格隆汇、和讯、金融界、e公司、Readhub 等多个快讯常驻"
    - "用多源交叉验证为由订阅十个快讯入口"
  route_reason: "前十页已经出现十多个几乎等价的快讯入口，主要复制同一组地缘、油价、公司公告和市场波动。当前架构不需要再加任何持续快讯源；更适合的是日历 + 结构解释 + 原始披露。"
  reuse_rule: "快讯用于发现，不用于定案。多平台复制同一新闻不构成多份独立证据。"
  evidence_locator: "docs/ops/tophub-finance-reassessment-pages-01-10-20260705.md#阶段性架构判断; docs/ops/tophub-finance-page-by-page-ledger-20260705.md#核心原则"
  next_action: "排除常驻。"

- source_id: "交易社区、荐股、龙虎榜、异动、资金流、复盘"
  current_route: rejected
  decision_status: recovered_from_pr
  selected_reason: "无常驻价值。"
  rejected_alternatives:
    - "雪球今日话题"
    - "股吧热榜"
    - "淘股吧"
    - "韭研公社"
    - "东方财富社区"
    - "选股通龙虎榜 / 复盘 / 早晚报"
    - "智通市场异动"
    - "证券时报股市一览"
  route_reason: "这些来源密集使用个股喊单、主线、预期差、龙头、涨停、资金流、盘中异动、交易复盘和情绪语言，与长期、分散、定投型投资节奏冲突。"
  reuse_rule: "交易社区可以作为市场情绪研究样本，但不能进入常驻信息系统，也不能生成投资动作。"
  evidence_locator: "docs/ops/tophub-finance-reassessment-pages-01-10-20260705.md#明确排除的类型; docs/ops/tophub-finance-final-closure-20260705.md#明确不做"
  next_action: "排除。"

- source_id: "研报聚合与报告索引"
  current_route: on_demand_research_index_only
  decision_status: recovered_from_pr
  selected_reason: "慧博、侠盾、乌拉邦、东方财富策略报告、CBNData、前瞻经济学人、BigQuant 等有研究索引和资料价值。"
  rejected_alternatives:
    - "研报搬运站持续订阅"
    - "把券商晨报 / 策略研报 / 个股研报变成日常流"
    - "用研报聚合替代原始机构"
  route_reason: "研报索引必须回原券商、机构、报告发布方或数据来源核验。报告不是证据本身；需要检查样本、方法、口径、委托方和商业合作。"
  reuse_rule: "报告源用于找关键词和候选报告，不作为最终证据。券商策略和目标价不进入长期配置系统。"
  evidence_locator: "docs/ops/tophub-finance-reassessment-pages-01-10-20260705.md#阶段性候选池; docs/ops/tophub-finance-reassessment-pages-11-20-20260705.md#累计240节点后的候选池"
  next_action: "按任务。"
```

---

## 9. 明确排除信号

```yaml
- source_id: "财经排除信号"
  current_route: reusable_rejection_rules
  decision_status: recovered_from_pr
  selected_reason: "用于复用判断，不是单一来源。"
  rejected_alternatives:
    - "以点击量、热度、标题强度、涨跌和交易语言作为常驻理由"
  route_reason: "明确排除：所有重复 7×24 快讯；所有股吧、淘股吧、韭研公社、东方财富社区类交易社区；龙虎榜、异动股、涨停复盘和主线 / 预期差节点；币价预测、永续合约、巨鲸仓位和杠杆喊单；研报搬运站作为持续源；本地代账、楼盘、企业 PR 和 SEO 聚合；更新异常或把导航项当文章的官方节点映射。"
  reuse_rule: "财经源只要出现结构性交易诱导、陈旧映射、导航污染、PR / SEO 聚合、价格预测和喊单，就先排除常驻；具体任务再定向打开。"
  evidence_locator: "docs/ops/tophub-finance-reassessment-pages-01-10-20260705.md#明确排除的类型; docs/ops/tophub-finance-reassessment-pages-21-27-20260705.md#明确不进入"
  next_action: "作为复用规则。"
```

---

## 10. 本补充 ledger 的复用规则

1. 财经来源先分六类：官方原始披露、结构分析、研究入口、市场温度、交易社区、Web3。
2. 财经系统服务长期尺度，不服务每日盯盘、清空快讯或跟随荐股。
3. 财新旧榜退出的直接理由是停止更新；财新首页进入是可工作编辑入口，不是名称神圣化。
4. 日经中文补日本、东亚、供应链、汇率和货币政策；不是又一个快讯源。
5. 有知有行补长期投资、保险、组合和行为尺度；不把平台组合动作当用户指令。
6. 财经日历、官方披露、交易所、央行和研究机构是工具 / 研究源，不是每日未读流。
7. Web3 要分中文结构长文、安全研究、英文原始报道和价格交易流；价格交易流排除。
8. 同一新闻被十个平台复制，不构成十份独立证据。
9. 候选、真实订阅和投资行动严格分开。
10. 只有平台真实操作完成并核验后，才修改真实订阅清单。

---

## 11. 仍需继续追回

财经逐页 324 节点已补成单独 ledger。

继续待办：

- 等用户本地提交 master checklist 与 coverage audit 后，再统一登记本文件和此前新增 ledger；
- 购物逐页；
- 政务、校务、专栏、浏览器与链接；
- 旧 ChatGPT 对话追索所有 partial 条目。
