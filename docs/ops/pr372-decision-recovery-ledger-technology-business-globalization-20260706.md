# PR #372 判断链追回：科技商业叙事与全球化来源（2026-07-06）

## 0. 边界

本文件是 `科技` 大类的补充 ledger，追回 PR #372 中已闭环的科技商业叙事、聚合快讯和全球化来源判断。

纳入文件：

- `docs/ops/tophub-technology-readhub-closure-20260704.md`
- `docs/ops/tophub-technology-huxiu-closure-20260704.md`
- `docs/ops/tophub-technology-36kr-page-review-20260704.md`
- `docs/ops/tophub-technology-globalization-closure-20260704.md`

当前不更新 `master-checklist` 与 `coverage-audit`，因为用户本地已有这两个文件的待提交修改。后续本地提交后再统一补登记。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 科技商业叙事源总规则

```yaml
- source_id: "科技商业叙事与聚合快讯总规则"
  current_route: keep_one_structural_slot_not_many_heat_feeds
  decision_status: recovered_from_pr
  selected_reason: "科技大类里需要观察科技公司、产业变化、产品和全球化结构，但不需要堆叠多个中文商业热榜、早晚报、快讯和二次聚合流。"
  rejected_alternatives:
    - "Readhub 热门 / 早报 / 24小时 / 一周热门同时常驻"
    - "虎嗅热文 / 最新 / 24小时 / 早晚报常驻"
    - "36氪多个热榜 / 快讯 / AI / 创投 / 视频榜常驻"
    - "把中文商业科技平台搬入 Folo"
  route_reason: "这些来源多数是同一批科技、商业、创投、AI、汽车、消费和社会议题的热度切片。TopHub 只保留能补出稀缺结构能力的来源；Folo 不收高频平台流。"
  reuse_rule: "以后遇到科技商业平台，先问它补的是结构能力还是热度切片；若只是热榜、快讯、早报、视频化或同内容池重排，就按需或退出。"
  evidence_locator: "tophub-technology-readhub-closure-20260704.md; tophub-technology-huxiu-closure-20260704.md; tophub-technology-36kr-page-review-20260704.md; tophub-technology-globalization-closure-20260704.md"
  next_action: "作为科技同族复用规则。"
```

---

## 2. Readhub

```yaml
- source_id: "Readhub｜热门话题"
  current_route: retired_from_tophub
  decision_status: recovered_from_pr
  selected_reason: "原本放在科技雷达，可快速看到商业科技舆论正在扩散什么。"
  rejected_alternatives:
    - "Readhub｜科技动态"
    - "Readhub｜每日早报"
    - "Readhub｜AI"
    - "Readhub｜一周热门"
    - "Readhub｜24 小时热榜"
    - "Readhub｜财经快讯"
    - "Readhub｜汽车"
    - "Readhub｜医疗产业"
  route_reason: "热门话题实际把商业、政策、国际、灾害、体育和科技压缩在一起，不再是科技专题。它与公共温度、36氪、IT之家、极客公园、爱范儿和滚动新闻重叠，且常来自二次转载，不能替代原始来源。"
  reuse_rule: "高压缩快讯流如果不能提供独立编辑判断，只能按需发现关键词，再回原始报道、公告或论文。"
  evidence_locator: "docs/ops/tophub-technology-readhub-closure-20260704.md#一为什么取消Readhub热门话题; #十Readhub的正确位置"
  next_action: "已在科技 final closeout 删除。"

- source_id: "Readhub 其他节点"
  current_route: on_demand_or_rejected_readhub_slices
  decision_status: recovered_from_pr
  selected_reason: "可按需回看某天或某周商业科技舆论集中话题，或发现某事件正在扩散。"
  rejected_alternatives:
    - "科技动态常驻"
    - "每日早报常驻"
    - "AI 常驻"
    - "一周热门 / 24小时热榜常驻"
    - "财经快讯 / 汽车 / 医疗产业常驻"
    - "进入 Folo"
  route_reason: "这些节点大量二次聚合，AI比例过高或标签失真；财经快讯和医疗产业明确不是稳定财经 / 医疗源。时间窗口变化不构成新的来源能力。"
  reuse_rule: "同一聚合器按 AI、财经、汽车、医疗或时间窗口切片时，不当作多个来源；只在具体问题中搜索。"
  evidence_locator: "docs/ops/tophub-technology-readhub-closure-20260704.md#二Readhub科技动态; #四ReadhubAI; #七Readhub财经快讯; #九Readhub医疗产业"
  next_action: "按需或排除。"
```

---

## 3. 虎嗅

```yaml
- source_id: "虎嗅网｜热文"
  current_route: retired_from_tophub
  decision_status: recovered_from_pr
  selected_reason: "原本承担中文科技商业热文入口。"
  rejected_alternatives:
    - "虎嗅网｜最新"
    - "虎嗅网｜24小时"
    - "虎嗅网｜早报"
    - "虎嗅网｜晚报"
    - "把虎嗅放进 Folo"
  route_reason: "虎嗅热文并非纯科技入口，而是科技、商业、消费、社会、人物和情绪议题混合热榜。它与 36氪、极客公园、爱范儿、Readhub、IT之家和公共温度高度重复，继续常驻只会增加相似叙事。"
  reuse_rule: "平台级商业叙事源可以按具体公司或行业事件查文章，但不以全站热榜维护常驻关系。"
  evidence_locator: "docs/ops/tophub-technology-huxiu-closure-20260704.md#一为什么取消虎嗅网热文; #七按需使用方式"
  next_action: "已在科技 final closeout 删除。"

- source_id: "虎嗅其他节点"
  current_route: on_demand_business_narrative_source
  decision_status: recovered_from_pr
  selected_reason: "可在需要寻找中文商业叙事、观察一件事如何被产业媒体解释、或对照 36氪 / 晚点 / 财新叙事时按需使用。"
  rejected_alternatives:
    - "最新常驻"
    - "24小时常驻"
    - "早报 / 晚报常驻"
    - "Folo"
  route_reason: "最新是高频全站流，24小时是综合转载滚动，早晚报是多事件摘要。它们不形成新的科技职责，也会制造持续信息压力。"
  reuse_rule: "早报、晚报、24小时和最新只是时间切片，不构成新来源。若未来发现具体作者或专栏不可替代，只评估作者 / 专栏原始更新。"
  evidence_locator: "docs/ops/tophub-technology-huxiu-closure-20260704.md#三虎嗅网最新; #四虎嗅网24小时; #五虎嗅网早报; #六虎嗅网晚报"
  next_action: "按需。"
```

---

## 4. 36氪：从热榜到全球化结构

```yaml
- source_id: "36氪｜24小时热榜"
  current_route: retired_from_tophub
  decision_status: recovered_from_pr
  selected_reason: "原本能显示中文科技商业公共热度。"
  rejected_alternatives:
    - "36氪｜深氪作为最终常驻"
    - "36氪｜最新"
    - "36氪｜收藏榜 / 综合榜 / 人气榜 / 热议榜"
    - "36氪｜快讯"
    - "36氪｜AI频道"
    - "36氪｜创投频道"
    - "36氪｜资讯推荐"
    - "36氪｜视频榜"
  route_reason: "24小时热榜把科技、创投、商业、职场和情绪内容混在一起，且与极客公园、IT之家、爱范儿、虎嗅、Readhub、TechCrunch、The Verge 等高度重复。最终不再保留相似热榜。"
  reuse_rule: "同一商业科技平台只保留最能补出系统缺口的一个入口；不叠热度榜。"
  evidence_locator: "docs/ops/tophub-technology-36kr-page-review-20260704.md#一为什么不再优先保留24小时热榜"
  next_action: "已被 36氪出海热门推荐替换。"

- source_id: "36氪｜深氪"
  current_route: on_demand_industry_longform
  decision_status: recovered_from_pr
  selected_reason: "比 24 小时热榜更接近 36氪独特能力：公司和产业过程、人物、组织、资本和市场之间的关系、中文创业与产业长报道标题发现。"
  rejected_alternatives:
    - "作为最终常驻替换 36氪热榜"
    - "进入 Folo"
  route_reason: "审完出海后被修订为按需入口。深氪与极客公园、晚点长报道、财新、FT 等存在重叠；36氪同一席位应优先给更稀缺的全球化结构观察。"
  reuse_rule: "阶段性候选要允许被后续文件修订；低频长报道有价值也不必常驻，如果另一来源补出更稀缺结构。"
  evidence_locator: "docs/ops/tophub-technology-36kr-page-review-20260704.md#二为什么选择36氪深氪; docs/ops/tophub-technology-globalization-closure-20260704.md#核心结论"
  next_action: "按需。"

- source_id: "36氪出海｜热门推荐"
  current_route: tophub_data_and_structure
  decision_status: recovered_from_pr
  selected_reason: "补出中国企业和技术进入全球市场的结构视角：消费品牌海外经营、汽车和机器人出海、供应链在全球 AI / 能源 / 基础设施中的位置、区域市场和监管条件。"
  rejected_alternatives:
    - "36氪｜深氪"
    - "36氪出海｜每日最新"
    - "雨果网｜今日头条"
    - "白鲸出海多个节点"
    - "知无不言跨境电商社区三个热榜"
    - "Enjoy出海｜新闻资讯"
  route_reason: "放入数据与结构，而不是科技雷达。它关注市场进入、区域制度、供应链、产业周期和全球经营，不是某款产品发布或当天科技公司新闻。与第一财经汽车新闻不同，它覆盖多行业全球化；汽车只是其中一个样本。"
  reuse_rule: "全球化来源要区分跨境卖家实务、数字内容出海和企业 / 技术全球化。系统缺的是第三种；因此只保留能补结构视角的压缩入口。"
  evidence_locator: "docs/ops/tophub-technology-globalization-closure-20260704.md#核心结论; #一为什么选择36氪出海热门推荐; #建议放入数据与结构; #与第一财经汽车新闻的关系"
  next_action: "已进入数据与结构。"
```

---

## 5. 出海同类为什么不选

```yaml
- source_id: "跨境电商实务来源"
  current_route: on_demand_cross_border_operations
  decision_status: recovered_from_pr
  selected_reason: "雨果网和知无不言社区能提供跨境电商平台规则、店铺运营、广告、账号、合规、侵权、税务和卖家经验。"
  rejected_alternatives:
    - "雨果网｜今日头条常驻"
    - "知无不言 30天 / 当天 / 7天热门常驻"
    - "进入 Folo"
  route_reason: "这些来源服务正在经营跨境店铺的人。用户当前没有持续运营亚马逊、Shopee 或 TikTok Shop 店铺的项目，日常订阅会制造大量无责任关系的规则与风险提醒；社区经验正确性也不稳定。"
  reuse_rule: "跨境电商源只有在真实业务或具体平台故障出现时按需；法律、税务、平台政策和合规必须回正式规则。"
  evidence_locator: "docs/ops/tophub-technology-globalization-closure-20260704.md#三雨果网今日头条; #四知无不言跨境电商社区三个节点"
  next_action: "按需。"

- source_id: "白鲸出海 / Enjoy出海"
  current_route: on_demand_digital_content_globalization
  decision_status: recovered_from_pr
  selected_reason: "可在研究游戏、App、短剧、移动广告、海外数字内容和应用商店政策时按需使用。"
  rejected_alternatives:
    - "白鲸出海今日热文 / 最新文章 / 本周热文 / 7×24 / 本月热文常驻"
    - "Enjoy出海新闻资讯常驻"
    - "进入 Folo"
  route_reason: "白鲸重点并非一般中国企业全球化，而是游戏、App、短剧、移动内容和广告变现；多个节点是同一内容池的时间窗口切片。Enjoy 更偏游戏与 App 全球发行任务源。"
  reuse_rule: "数字内容出海任务出现时按需，发现题目后回 Sensor Tower、公司财报、应用商店数据和平台政策等直接来源。"
  evidence_locator: "docs/ops/tophub-technology-globalization-closure-20260704.md#五白鲸出海五个节点; #六Enjoy出海新闻资讯"
  next_action: "按需。"
```

---

## 6. 本补充 ledger 的复用规则

1. 中文科技商业平台不按数量堆叠；只保留补系统缺口的一项。
2. 热榜、早报、晚报、24小时、最新、视频榜、互动榜是切片，不是新来源关系。
3. Readhub 这类压缩聚合器适合发现关键词，不适合作事实或一手来源。
4. 虎嗅这类商业叙事平台按具体文章和事件使用，不维护全站未读流。
5. 36氪席位最终从国内商业热度转向全球化结构观察；这是复用修订，不是机械保留品牌。
6. 出海来源必须区分：跨境电商实务、数字内容出海、企业与技术全球化。不同层不能互相替代。
7. 涉及平台规则、监管、劳动法、关税、税务、公司经营时，回正式文件、财报、招股书、公司公告或平台政策。

---

## 7. 仍需继续追回

科技大类已补入商业叙事与全球化来源，但仍是 partial。

继续待办：

- 科普 / 科学 / The Register / Apple / 汽车 / Counterpoint 等 final closeout 相关来源；
- 少数派、IT之家、极客公园、TechCrunch、The Verge 等现有科技雷达来源的同族判断；
- AI、App、快讯、限免、报告、教育、电商、数码、酷安、Google、TechWeb 等 closure；
- 后续等用户本地提交 master checklist 与 coverage audit 后，再登记本文件。
