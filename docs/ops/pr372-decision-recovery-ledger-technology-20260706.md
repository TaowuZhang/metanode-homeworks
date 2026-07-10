# PR #372 判断链追回：科技大类（第一版，2026-07-06）

## 0. 边界

本文件追回 `TopHub 科技` 大类中已经能从 PR #372 文件直接确认的判断链。

已读取并纳入本版：

- `docs/ops/tophub-technology-final-closeout-20260704.md`

本版尚未覆盖科技同族 closure 的全部细节。以下文件仍需继续追回：

- `tophub-technology-directory-review-20260704.md`
- `tophub-technology-36kr-page-review-20260704.md`
- `tophub-technology-ai-closure-20260704.md`
- `tophub-technology-ai-briefs-closure-20260704.md`
- `tophub-technology-app-closure-20260704.md`
- `tophub-technology-apple-closure-20260704.md`
- `tophub-technology-automotive-closure-20260704.md`
- `tophub-technology-chouti-closure-20260704.md`
- `tophub-technology-coolapk-closure-20260704.md`
- `tophub-technology-digital-closure-20260704.md`
- `tophub-technology-ecommerce-closure-20260704.md`
- `tophub-technology-education-closure-20260704.md`
- `tophub-technology-flash-news-closure-20260704.md`
- `tophub-technology-freebies-closure-20260704.md`
- `tophub-technology-general-tech-closure-20260704.md`
- `tophub-technology-globalization-closure-20260704.md`
- `tophub-technology-google-closure-20260704.md`
- `tophub-technology-huxiu-closure-20260704.md`
- `tophub-technology-ithome-closure-20260704.md`
- `tophub-technology-jandan-closure-20260704.md`
- `tophub-technology-popular-science-closure-20260704.md`
- `tophub-technology-readhub-closure-20260704.md`
- `tophub-technology-reports-closure-20260704.md`
- `tophub-technology-science-closure-20260704.md`
- `tophub-technology-sspai-closure-20260704.md`
- `tophub-technology-techweb-closure-20260704.md`

因此本文件状态是：`partial_recovered_from_pr`。

---

## 1. 科技大类已执行动作

来源证据：`docs/ops/tophub-technology-final-closeout-20260704.md`

科技大类已执行：

- 删除 2 个独立节点；
- 完成 4 组替换；
- 新增 3 个节点；
- 新增 1 个精确追踪；
- Folo 保持 27，不变。

```yaml
- source_id: "Readhub｜热门话题"
  current_route: retired_from_tophub
  decision_status: recovered_from_pr_partial
  selected_reason: "原为科技 / 热点聚合入口。"
  rejected_alternatives:
    - "继续作为独立常驻节点"
  route_reason: "科技最终收口已删除该独立节点；更完整的删除理由需回 `tophub-technology-readhub-closure-20260704.md` 追回。"
  evidence_locator: "docs/ops/tophub-technology-final-closeout-20260704.md#已执行动作-删除2个独立节点"
  next_action: "继续读取 Readhub closure 补逐源理由。"

- source_id: "虎嗅网｜热文"
  current_route: retired_from_tophub
  decision_status: recovered_from_pr_partial
  selected_reason: "原为中文科技 / 商业热文入口。"
  rejected_alternatives:
    - "继续作为独立常驻节点"
  route_reason: "科技最终收口已删除该独立节点；更完整的删除理由需回 `tophub-technology-huxiu-closure-20260704.md` 追回。"
  evidence_locator: "docs/ops/tophub-technology-final-closeout-20260704.md#已执行动作-删除2个独立节点"
  next_action: "继续读取虎嗅 closure 补逐源理由。"
```

---

## 2. 四组替换

```yaml
- source_id: "36氪｜24小时热榜"
  current_route: retired_or_replaced_from_tophub
  decision_status: recovered_from_pr_partial
  selected_reason: "原为 36 氪综合热榜入口。"
  rejected_alternatives:
    - "继续使用 24 小时热榜承担科技 / 产业观察"
  route_reason: "已由 `36氪出海｜热门推荐` 替换。最终 closeout 只记录替换结果；具体为什么从 24 小时热榜转为出海视角，需要回 36kr page review / globalization closure 继续追回。"
  evidence_locator: "docs/ops/tophub-technology-final-closeout-20260704.md#完成4组替换"
  next_action: "读取 36kr 与 globalization 相关 closure。"

- source_id: "36氪出海｜热门推荐"
  current_route: tophub_data_and_structure
  decision_status: recovered_from_pr_partial
  selected_reason: "作为全球化 / 出海产业入口进入数据与结构。"
  rejected_alternatives:
    - "36氪｜24小时热榜"
  route_reason: "最终顺序中位于数据与结构，承接科技、产业、企业全球化与出海观察；详细理由需回 36kr / globalization closure。"
  ordering_reason: "位于央视新闻调查之后、第一财经汽车新闻之前。"
  evidence_locator: "docs/ops/tophub-technology-final-closeout-20260704.md#完成4组替换; #数据与结构3040"
  next_action: "补同族文件理由。"

- source_id: "科普中国网｜热点排行"
  current_route: retired_or_replaced_from_tophub
  decision_status: recovered_from_pr_partial
  selected_reason: "原为综合科普热点入口。"
  rejected_alternatives:
    - "继续使用热点排行"
  route_reason: "已由 `科普中国网｜头条` 替换；具体替换理由需回 popular-science closure。"
  evidence_locator: "docs/ops/tophub-technology-final-closeout-20260704.md#完成4组替换"
  next_action: "读取科普 closure。"

- source_id: "科普中国网｜头条"
  current_route: tophub_technology_radar
  decision_status: recovered_from_pr_partial
  selected_reason: "作为科技雷达中的科普入口。"
  rejected_alternatives:
    - "科普中国网｜热点排行"
  route_reason: "最终顺序中进入科技雷达，位于 The Register 之后、果壳科学人之前；完整理由需回 popular-science closure。"
  ordering_reason: "科技雷达第 25 位。"
  evidence_locator: "docs/ops/tophub-technology-final-closeout-20260704.md#完成4组替换; #科技雷达1929"
  next_action: "补同族文件理由。"

- source_id: "果壳｜每日精选"
  current_route: retired_or_replaced_from_tophub
  decision_status: recovered_from_pr_partial
  selected_reason: "原为科学 / 科普类入口。"
  rejected_alternatives:
    - "继续使用每日精选承担科学新闻"
  route_reason: "已由 `Science Magazine｜Latest News` 替换；但最终顺序仍保留 `果壳｜科学人`。具体替换理由需回 science / popular-science closure。"
  evidence_locator: "docs/ops/tophub-technology-final-closeout-20260704.md#完成4组替换"
  next_action: "读取科学相关 closure。"

- source_id: "Science Magazine｜Latest News"
  current_route: tophub_technology_radar
  decision_status: recovered_from_pr_partial
  selected_reason: "作为国际科学新闻入口进入科技雷达。"
  rejected_alternatives:
    - "果壳｜每日精选"
  route_reason: "最终顺序中进入科技雷达，与果壳科学人并列但承担不同科学来源层级；具体理由需回 science closure。"
  ordering_reason: "位于 `果壳｜科学人` 之后、`HelloGitHub｜月刊` 之前。"
  evidence_locator: "docs/ops/tophub-technology-final-closeout-20260704.md#完成4组替换; #科技雷达1929"
  next_action: "补同族文件理由。"

- source_id: "爱范儿｜每日最新"
  current_route: retired_or_replaced_from_tophub
  decision_status: recovered_from_pr_partial
  selected_reason: "原为消费科技 / 数码资讯入口。"
  rejected_alternatives:
    - "继续使用爱范儿每日最新"
  route_reason: "已由 `The Register｜Latest` 替换；具体为什么从消费科技切换为英文企业 IT / 基础设施入口，需回 general tech / ithome / techweb closure。"
  evidence_locator: "docs/ops/tophub-technology-final-closeout-20260704.md#完成4组替换"
  next_action: "读取对应 closure。"

- source_id: "The Register｜Latest"
  current_route: tophub_technology_radar
  decision_status: recovered_from_pr
  selected_reason: "承担企业 IT、基础设施、云、软件、硬件、政策与安全的英文科技入口。"
  rejected_alternatives:
    - "爱范儿｜每日最新"
    - "iThome｜新闻作为新增常驻"
  route_reason: "进入科技雷达；同时解释为什么不新增 iThome：iThome 的企业 IT、资安、云原生和基础设施能力由 The Register、Folo 中的 Dark Reading / Krebs、官方公告与 CVE 共同承担。"
  ordering_reason: "科技雷达中位于 IT之家日榜之后、科普中国网头条之前。"
  evidence_locator: "docs/ops/tophub-technology-final-closeout-20260704.md#完成4组替换; #最终未采用的候选-ithome新闻; #科技雷达1929"
  next_action: "可继续回同族 closure 补更细理由。"
```

---

## 3. 三个新增节点

```yaml
- source_id: "Apple 支持｜更换和维修扩展计划"
  current_route: tophub_technology_radar
  decision_status: recovered_from_pr
  selected_reason: "低频官方行动入口，捕捉 Apple 维修、更换和扩展计划。"
  rejected_alternatives:
    - "把 Apple 维修计划放进 Folo 制造未读"
    - "用宽泛 Apple 新闻源替代官方行动入口"
  route_reason: "TopHub 承担低成本公共发现和少量官方行动入口；Folo 不重复订阅 Apple 维修计划。"
  ordering_reason: "科技雷达最后一位，位于 HelloGitHub 月刊之后。"
  evidence_locator: "docs/ops/tophub-technology-final-closeout-20260704.md#新增3个节点; #科技雷达1929; #tophub终身版与folo免费版如何分工"
  next_action: "可回 Apple closure 补更细理由。"

- source_id: "第一财经｜汽车新闻"
  current_route: tophub_data_and_structure
  decision_status: recovered_from_pr
  selected_reason: "汽车产业与市场结构入口。"
  rejected_alternatives:
    - "把汽车任务放进 Folo"
    - "用宽泛科技 / 财经快讯替代汽车结构入口"
  route_reason: "TopHub 数据与结构承担汽车与终端市场研究，不制造 Folo 未读债务。"
  ordering_reason: "数据与结构中位于 36氪出海之后、Counterpoint Research 之前。"
  evidence_locator: "docs/ops/tophub-technology-final-closeout-20260704.md#新增3个节点; #数据与结构3040"
  next_action: "可回 automotive closure 补更细理由。"

- source_id: "Counterpoint Research｜最新见解"
  current_route: tophub_data_and_structure
  decision_status: recovered_from_pr
  selected_reason: "终端市场、智能手机、硬件产业与市场研究入口。"
  rejected_alternatives:
    - "把报告库或快讯流放进 Folo"
  route_reason: "TopHub 承担终端市场研究；Folo 不重复订阅 Counterpoint。"
  ordering_reason: "数据与结构中位于第一财经汽车新闻之后。"
  evidence_locator: "docs/ops/tophub-technology-final-closeout-20260704.md#新增3个节点; #数据与结构3040; #tophub终身版与folo免费版如何分工"
  next_action: "可回 reports / digital closure 补更细理由。"
```

---

## 4. Apple Foundation Models 精确追踪

```yaml
- source_id: "Apple Foundation Models"
  current_route: tophub_precise_tracker
  decision_status: recovered_from_pr
  selected_reason: "专属结果页能够捕捉 Apple Foundation Models 与 AFM 模型更新、Apple Machine Learning Research、Swift / SDK / 应用接入、端侧模型与开发案例。"
  rejected_alternatives:
    - "Apple Intelligence"
    - "苹果AI"
    - "Foundation Models 宽词"
    - "追踪机器人"
    - "外部通知"
    - "自动刷新"
  route_reason: "保留为低频精确技术雷达；不扩成宽词，因为结果页同时存在同一事件重复转载、Siri / Gemini / 公司战略泛新闻、第三方评论和低质量二次解读。"
  evidence_locator: "docs/ops/tophub-technology-final-closeout-20260704.md#apple-foundation-models追踪"
  next_action: "无；后续只在具体技术变化出现时处理。"
```

---

## 5. 最终未采用候选

```yaml
- source_id: "iThome｜新闻"
  current_route: on_demand_chinese_enterprise_it_source
  decision_status: recovered_from_pr
  selected_reason: "中文企业 IT 与安全问题的按需入口。"
  rejected_alternatives:
    - "新增为 TopHub 常驻节点"
  route_reason: "其企业 IT、资安、云原生和基础设施能力已经由 The Register、Folo 中的 Dark Reading / Krebs、具体漏洞与产品官方公告和 CVE 共同承担，因此不额外新增。"
  evidence_locator: "docs/ops/tophub-technology-final-closeout-20260704.md#最终未采用的候选-ithome新闻"
  next_action: "无；按需。"

- source_id: "反斗软件｜最新发布"
  current_route: on_demand_software_discovery
  decision_status: recovered_from_pr
  selected_reason: "可作为软件发现的按需入口。"
  rejected_alternatives:
    - "新增为 TopHub 常驻节点"
    - "进入 Folo"
  route_reason: "其软件发现能力与 Folo 中的小众软件、TopHub 中的少数派、App 与软件目录的按需访问重复，因此不新增；反斗软件与限免页面保留按需使用。"
  evidence_locator: "docs/ops/tophub-technology-final-closeout-20260704.md#最终未采用的候选-反斗软件最新发布"
  next_action: "无；按需。"
```

---

## 6. TopHub / Folo / 按需分工

```yaml
- source_id: "科技大类 TopHub 分工"
  current_route: tophub_low_cost_public_discovery
  decision_status: recovered_from_pr
  selected_reason: "TopHub 承担 73 个低成本公共发现节点，包含科技门户、企业 IT、科学、全球化、汽车与终端市场研究、设备维修计划和少量精确追踪。"
  rejected_alternatives:
    - "把科技快讯、日报、榜单、报告库、限免和论坛交给 Folo / RSSHub"
  route_reason: "TopHub 用于低成本发现和精确追踪，不制造 Folo 未读债务。"
  evidence_locator: "docs/ops/tophub-technology-final-closeout-20260704.md#tophub终身版与folo免费版如何分工"
  next_action: "继续补同族 closure。"

- source_id: "科技大类 Folo 分工"
  current_route: keep_folo_27_unchanged
  decision_status: recovered_from_pr
  selected_reason: "Folo 继续保持 27 个直接关系与观察来源，2026-07-17 前不变。"
  rejected_alternatives:
    - "重复订阅 The Register"
    - "重复订阅 Science Magazine"
    - "重复订阅 Counterpoint"
    - "重复订阅 Apple 维修计划"
    - "用 RSSHub 承接快讯、日报、榜单、报告库、限免和论坛"
  route_reason: "Folo 保存直接关系与观察来源，不接收科技目录里的快讯、榜单、报告库和官方低频行动入口。"
  evidence_locator: "docs/ops/tophub-technology-final-closeout-20260704.md#tophub终身版与folo免费版如何分工"
  next_action: "2026-07-17 复查。"

- source_id: "科技大类按需网页"
  current_route: on_demand_web
  decision_status: recovered_from_pr
  selected_reason: "电商、教育、报告、快讯、评测、硬件购买、竞赛和平台政策只在具体任务中访问。"
  rejected_alternatives:
    - "常驻 TopHub / Folo"
  route_reason: "具体状态回到官方文件、论文、公司公告、监管和项目原文。"
  evidence_locator: "docs/ops/tophub-technology-final-closeout-20260704.md#tophub终身版与folo免费版如何分工"
  next_action: "按任务。"
```

---

## 7. Folo 后续候选队列

```yaml
- source_id: "MacStories / Simon Willison / Rest of World / The MIT Press Reader / Quanta Magazine"
  current_route: folo_review_candidate_queue
  decision_status: recovered_from_pr_partial
  selected_reason: "科技 final closeout 中列为 2026-07-17 Folo 复查优先候选。"
  rejected_alternatives:
    - "全部加入 Folo"
  route_reason: "这是候选队列，不是全部加入清单；新来源实行一进一出，第一轮最多测试 3 个。"
  evidence_locator: "docs/ops/tophub-technology-final-closeout-20260704.md#folo后续"
  next_action: "回对应 AI / 科技 / 阅读来源文件继续追回逐源理由。"

- source_id: "Google AI Developers / Google Developers Blog / Stack Overflow Blog / Logan Kilpatrick / News Minimalist / Hacker News / 私人观察来源 / 中国爬楼联盟 / 街拍中国 / 行动提醒"
  current_route: folo_review_existing_sources
  decision_status: recovered_from_pr_partial
  selected_reason: "科技 final closeout 中列为 2026-07-17 先复查现有来源。"
  rejected_alternatives:
    - "不复查直接扩容"
  route_reason: "先复查是否重复、是否仍有真实触发，尤其 Google AI Developers 与 Google Developers Blog 是否重复，中国爬楼联盟与街拍中国最多保留一个，没有真实触发的行动提醒退出。"
  evidence_locator: "docs/ops/tophub-technology-final-closeout-20260704.md#folo后续"
  next_action: "2026-07-17 或旧对话追索。"
```

---

## 8. 本版仍需继续追回

本文件只完成科技大类 final closeout 的第一版判断链。

仍需继续追回：

- Readhub / 虎嗅删除的完整逐源理由；
- 36kr、科普、科学、爱范儿 / The Register 替换的更细理由；
- Apple、汽车、Counterpoint 新增的同族 closure 证据；
- AI、App、快讯、限免、报告、教育、电商、数码、酷安、Google、TechWeb、少数派等 closure 中的按需 / 拒绝 / 候选关系；
- Folo 后续候选队列的逐源理由。

完成这些以后，才能把科技大类标为 `recovered_from_pr`；目前只能标为 `partial_recovered_from_pr`。
