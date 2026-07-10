# PR #372 判断链追回：TopHub 总览、来源发现与路由边界（2026-07-06）

## 0. 边界

本文件追回 PR #372 中 TopHub `全部`页顺序核验、来源发现路径、公共温度与七分组使用地图、更多目录补充审查形成的判断链。

纳入文件：

- `docs/ops/tophub-all-page-order-audit-20260705.md`
- `docs/ops/tophub-source-discovery-and-entertainment-audit-20260704.md`
- `docs/ops/tophub-public-temperature-and-group-map-20260704.md`
- `docs/ops/tophub-more-directories-audit-20260704.md`
- `资源/雷达/TopHub 全部订阅顺序.md`
- `资源/雷达/TopHub 使用配置.md`
- `资源/雷达/TopHub 能力地图.md`

边界说明：

- `tophub-all-page-order-audit-20260705.md`记录的是当时 84 项平台顺序核验；后续四个失效节点已取消不补位，当前真实 TopHub 基线以 `资源/雷达/TopHub 全部订阅顺序.md` 的 80 项为准。
- 本文件不重新审核单个目录，也不更新平台设置。
- 本文件追回的是：全部页、顶部目录、更多目录、首页分组和来源发现路径为什么只是导航与核验能力，不是新一轮扩张命令。
- 当前不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 总规则：来源发现层与个人路由层分开

```yaml
- source_id: "TopHub 来源发现层 / 个人路由层"
  current_route: discovery_layer_plus_user_routing_layer
  decision_status: recovered_from_pr
  selected_reason: "TopHub 顶部栏目和更多目录能用于翻找站内已有节点；七个自定义分组用于把选中的站内节点或外部订阅路由到用户自己的用途容器。"
  rejected_alternatives:
    - "把顶部目录当成个人分组"
    - "把目录规模当成新增理由"
    - "把节点出现在某目录中当作最终归类依据"
    - "把来源发现写成订阅完成"
  route_reason: "文件明确区分两层结构：来源发现层通过综合、科技、娱乐、社区、购物、财经、开发、简报、AI 与更多目录寻找候选；个人路由层将选中的节点放进七个用途分组。顶部目录的分类是平台方粗分，自定义分组才是用户用途判断。"
  reuse_rule: "以后任何来源先判断：它只是被目录发现，还是已经被用户路由。目录名不决定最终分组；只有填补真实缺口、节点口径清楚、与现有来源不重复，才进入路由。"
  evidence_locator: "docs/ops/tophub-source-discovery-and-entertainment-audit-20260704.md#顶部栏目也是订阅源发现目录; #由此形成的原则"
  next_action: "作为来源发现总规则。"
```

---

## 2. `全部`页顺序核验：平台位置证据，不是永久基线

```yaml
- source_id: "TopHub 全部页顺序核验 84 项"
  current_route: historical_platform_order_verification_superseded_by_current_80_baseline
  decision_status: recovered_from_pr_with_temporal_boundary
  selected_reason: "该文件确认当时 `全部`页共有 84 个订阅节点，其中 83 个内容节点属于七个用途分组，1 个系统级节点为 `专栏｜订阅聚合`；13 个新增内容节点已全部进入目标位置。"
  rejected_alternatives:
    - "继续把 84 当作当前真实数量"
    - "把当时生活与社区第 63/64 项保留为当前有效节点"
    - "把顺序核验当成来源价值判断"
  route_reason: "该文件是平台顺序证据，证明新增节点当时已放到约定位置，并说明专栏聚合应放在第 84 项。后续已有失效节点清理，当前真实基线为 80 项；因此本文件只作为历史顺序核验与平台执行证据。"
  reuse_rule: "平台顺序文件必须带时间边界。顺序核验回答‘当时放对了吗’，不回答‘现在还有效吗’。当前状态必须回最新全部订阅顺序或平台核验。"
  evidence_locator: "docs/ops/tophub-all-page-order-audit-20260705.md#结论; #专栏订阅聚合的位置; 资源/雷达/TopHub 全部订阅顺序.md"
  next_action: "不以 84 覆盖当前 80。"
```

---

## 3. 专栏｜订阅聚合为什么放尾部

```yaml
- source_id: "专栏｜订阅聚合 全部页位置"
  current_route: system_overview_entry_at_end
  decision_status: recovered_from_pr
  selected_reason: "它汇总全部订阅，不属于某一个内容领域。"
  rejected_alternatives:
    - "放进七个用途分组"
    - "放在全部页最前"
    - "把它当成某个主题节点"
  route_reason: "放进七个用途分组会产生错误分类，放在最前会打断内容目录本身的阅读路径；放在最后最符合‘内容来源 + 总览入口’的层级关系，适合完成主题浏览后用于回看、漫游和总览。"
  reuse_rule: "系统总览入口不和内容源混排；聚合入口应在内容之后，不占某个主题分组。"
  evidence_locator: "docs/ops/tophub-all-page-order-audit-20260705.md#专栏订阅聚合的位置"
  next_action: "保持系统级定位。"
```

---

## 4. 节点口径：不能按品牌整包订阅

```yaml
- source_id: "TopHub 节点口径规则"
  current_route: node_granularity_before_subscription_decision
  decision_status: recovered_from_pr
  selected_reason: "同一个媒体或平台可能有热门、最新、日报、周刊、专栏、快讯、标签页和具体栏目等多个节点，不同口径会直接改变信息流。"
  rejected_alternatives:
    - "只问要不要订阅某个品牌"
    - "同一品牌多个节点默认全订"
    - "把热门、最新、快讯和周刊视作相同来源"
  route_reason: "科技目录实测约 415 个节点、35 页，少数派、36氪、果壳、苹果相关来源都存在多个不同口径。热门更适合观察传播与编辑精选，最新更接近连续更新且噪声更多，日报 / 周刊适合周期性查阅，快讯适合事件发现而非长期深读。"
  reuse_rule: "补源时必须问：需要的是这个品牌的哪一种更新口径？选择节点口径后，再比较它与现有成员的增量价值。"
  evidence_locator: "docs/ops/tophub-source-discovery-and-entertainment-audit-20260704.md#来源选择精确到节点口径"
  next_action: "作为补源规则。"
```

---

## 5. 顶部娱乐目录：平台仓库不等于影游音乐

```yaml
- source_id: "TopHub 娱乐目录 1085 节点"
  current_route: large_platform_repository_not_user_group
  decision_status: recovered_from_pr
  selected_reason: "娱乐目录规模约 1085 个节点、91 页，覆盖哔哩哔哩、AcFun、抖音、快手、豆瓣、IMDb、猫眼、腾讯视频、爱奇艺、微信读书、起点、纵横、QQ 音乐、网易云音乐、喜马拉雅、虎扑、懂球帝及短视频、阅读、小说、影视、游戏、体育、播客、音乐、动漫、摄影、时尚等主题。"
  rejected_alternatives:
    - "把娱乐目录等同于个人分组影游音乐"
    - "因为目录里出现某平台就加入影游音乐"
    - "把抖音总榜、豆瓣小组、微信读书总榜、网文榜等全部放进娱乐流"
  route_reason: "娱乐是平台方大仓库，不是单一领域。抖音总榜更适合公共温度，豆瓣小组更适合生活与社区，喷嚏网本周热读可能接近慢读与思想，网文榜只有用户开始找网文时才需要补入，播客策展可能比短视频榜更是未来缺口。"
  reuse_rule: "目录归类不能替代个人路由。娱乐来源要先判断它提供热度、最新、口碑、评论、策展还是完整更新，再决定进入公共温度、影游音乐、慢读与思想、生活与社区、Folo 或按需。"
  evidence_locator: "docs/ops/tophub-source-discovery-and-entertainment-audit-20260704.md#娱乐来源仓库实测"
  next_action: "已由娱乐专项继续细审。"
```

---

## 6. 影游音乐早期实测：成立，但榜单与编辑必须分层

```yaml
- source_id: "影游音乐 v0 实测"
  current_route: low_pressure_culture_entertainment_entry
  decision_status: recovered_from_pr
  selected_reason: "早期影游音乐组由 Apple Music、网易云飙升、IMDb、游研社、机核、开眼、B站每周必看和豆瓣一周口碑构成，已经能形成轻松文化娱乐入口。"
  rejected_alternatives:
    - "把它当最近最火都要看的总榜"
    - "一次带走多个榜单和作品"
    - "新增 B站日榜、抖音总榜或网文榜"
  route_reason: "这个分组成立是因为音乐、电影、游戏和视频都有，但没有被拆成过多小分类；首屏同时包含榜单与编辑来源；豆瓣口碑与 IMDb 热门形成中文口碑 / 全球注意力差异；游研社和机核把游戏放回创作、产业和文化语境。后续娱乐专项据此继续替换网易云飙升、开眼等低密度或短周期节点。"
  reuse_rule: "影游音乐用榜单察觉变化，用编辑来源决定是否靠近作品。一次只带走一个作品或一个问题，不把榜单转成收藏任务。"
  evidence_locator: "docs/ops/tophub-source-discovery-and-entertainment-audit-20260704.md#影游音乐分组实测"
  next_action: "早期实测作为后续娱乐重构起点。"
```

---

## 7. 公共温度：平台差异地图，不是事实核验页

```yaml
- source_id: "公共温度 v0 实测"
  current_route: public_attention_and_platform_bias_map
  decision_status: recovered_from_pr
  selected_reason: "公共温度主动保留平台差异、传播偏差和公共情绪，用于同时看见微博、抖音、小红书、知乎、微信热文、澎湃、中国新闻周刊、观察者网、人民日报、央视、参考消息等不同注意力结构。"
  rejected_alternatives:
    - "当成低噪声资讯页"
    - "当成事实核验页"
    - "按热度数字跨平台比较"
    - "顺序读完单个榜单"
    - "形成未读债务"
  route_reason: "公共温度的价值不在于哪条排第一，而在比较同一事件是否跨平台出现、不同平台如何改变标题和情绪、哪些议题只在某平台内部高热，以及机构、媒体和生活平台注意力是否分离。"
  reuse_rule: "公共温度只回答大家正在看什么、平台怎样讲；不回答什么是真的、什么最重要、用户应该做什么。涉及事实、投资、健康、工具或行动时离开热榜找原始来源。"
  evidence_locator: "docs/ops/tophub-public-temperature-and-group-map-20260704.md#公共温度分组结论; #使用规则"
  next_action: "后续以当前 16 项公共温度为准。"
```

---

## 8. 七个分组地图：长期路由容器，不是封闭切片

```yaml
- source_id: "七个自定义分组使用地图"
  current_route: long_term_usage_containers
  decision_status: recovered_from_pr
  selected_reason: "七组分别承担公共注意力地图、科技轻量雷达、数据证据面板、慢读漫游书架、待补强生活容器、文化娱乐入口、视觉自然启动器。"
  rejected_alternatives:
    - "把七组看成当前节点的永久封闭切片"
    - "因当前测试期不加源就理解为未来禁止增加"
    - "根据一天页面样本启动第二轮删除或补源"
  route_reason: "首页审计和公共温度地图都明确：七个自定义分组是长期路由容器，不受当前节点永久限制。新增、删除、替换应基于连续使用证据，而不是一天样本。"
  reuse_rule: "后续新来源可以直接路由到现有七组，不必新建更多分类；但必须说明填补哪个缺口。"
  evidence_locator: "docs/ops/tophub-homepage-audit-20260704.md#分组不是封闭视图; docs/ops/tophub-public-temperature-and-group-map-20260704.md#七个分组的最终用途"
  next_action: "进入真实使用观察。"
```

---

## 9. 更多目录：不为把目录看完而翻页

```yaml
- source_id: "TopHub 更多目录"
  current_route: candidate_directories_not_expansion_obligation
  decision_status: recovered_from_pr
  selected_reason: "更多目录包含报刊、设计、校务、政务、专栏等巨大候选仓库，能提供大量站内来源发现。"
  rejected_alternatives:
    - "因为目录规模大就扩张当前订阅"
    - "为了把目录看完继续翻页"
    - "把目录存在视为当前系统缺口"
    - "把新增后仍要求反复截图证明位置"
  route_reason: "用户确认既定新增已按约定加入即可记为完成，不再要求反复截图。更多目录只在具体缺口出现时进入。报刊要先核对日期，设计当前缺口不是 AI 教程流，校务只在明确学校或考试阶段使用，政务按政策、立法、预警、监管需求精确添加，专栏逐个核验。"
  reuse_rule: "更多目录是候选仓库。进入前必须带着明确缺口；离开时必须只带走已核验来源，不批量订阅。"
  evidence_locator: "docs/ops/tophub-more-directories-audit-20260704.md#执行状态; #最终决定"
  next_action: "不继续为了把目录看完而翻页。"
```

---

## 10. 报刊 / 设计 / 校务 / 政务 / 专栏在更多目录中的早期边界

```yaml
- source_id: "更多目录｜报刊"
  current_route: on_demand_newspaper_directory_with_date_check
  decision_status: recovered_from_pr
  selected_reason: "报刊目录约 873 个节点，可按省份和中央级、日报、早报、晚报、财经报、法制报、都市报、健康报、老年报、文化教育报等浏览。"
  rejected_alternatives:
    - "因报刊目录规模大而增加传统报纸数量"
    - "只看报纸品牌不核对日期"
  route_reason: "样本显示报刊新鲜度差异极大：每日经济新闻、第一财经日报、21世纪经济报道、新华每日电讯有 2026 内容；人民日报、经济日报、人民日报海外版等有明显旧内容。因此电子报名称不能证明仍更新。"
  reuse_rule: "报刊节点新增前必须验证连续更新；按地区、行业或纸媒版面任务使用。"
  evidence_locator: "docs/ops/tophub-more-directories-audit-20260704.md#报刊目录"
  next_action: "已由报刊专项继续处理。"

- source_id: "更多目录｜设计"
  current_route: on_demand_design_directory_not_visual_recovery_source
  decision_status: recovered_from_pr
  selected_reason: "设计目录包含站酷、优设网、标志情报局、百度用户体验中心、腾讯设计团队、广告门与 Designmodo 等。"
  rejected_alternatives:
    - "当前新增设计节点"
    - "把设计目录当视觉恢复来源"
    - "用 AI 教程流补视觉与自然缺口"
  route_reason: "站酷更接近创作者作品与商业视觉市场；优设更接近设计工具和教程资讯；标志情报局更接近品牌设计观察；企业设计中心更接近案例档案。当前视觉与自然缺口是低文字负担、高质量图像与自然空间，不是更多 AI 设计教程。"
  reuse_rule: "品牌标志、产品体验、设计工具、作品漫游各自按任务调用。"
  evidence_locator: "docs/ops/tophub-more-directories-audit-20260704.md#设计目录"
  next_action: "已由设计专项继续处理。"

- source_id: "更多目录｜校务"
  current_route: on_demand_school_stage_specific_directory
  decision_status: recovered_from_pr
  selected_reason: "校务目录约 352 个节点，包含高校新闻、教务通知、研究生招生、学院公告、学术活动和教育考试机构通知。"
  rejected_alternatives:
    - "新增校务节点"
    - "浏览全目录"
    - "把校园论坛与校务通知混同"
  route_reason: "校务来源高度依赖明确机构与阶段。只有正在申请、报考、就读、合作或关注某学院时，具体通知才有价值；综合高校新闻对当前系统没有稳定增量。"
  reuse_rule: "出现明确学校、项目或考试阶段时，再订阅对应招生网、研究生院、学院或考试院具体栏目。"
  evidence_locator: "docs/ops/tophub-more-directories-audit-20260704.md#校务目录"
  next_action: "已由校务 snapshot 继续处理。"

- source_id: "更多目录｜政务"
  current_route: on_demand_official_policy_legislation_alert_regulation_directory
  decision_status: recovered_from_pr
  selected_reason: "政务目录约 1063 个节点，包含中国政府网、人大、应急广播、12306、部委、证监会、交易所等大量官方来源。"
  rejected_alternatives:
    - "新增广泛政务流"
    - "用官方二字代替用途判断"
    - "同一机构多个近似栏目一起订阅"
  route_reason: "当前国家统计局两个节点已承担数据与结构官方底稿。其他政务源按政策、立法、灾害、铁路出行、金融规则等具体需求精确添加。"
  reuse_rule: "政务源按责任关系和任务触发：政策、立法、预警、出行、金融规则、地方事项。"
  evidence_locator: "docs/ops/tophub-more-directories-audit-20260704.md#政务目录"
  next_action: "已由政务 snapshot 继续处理。"

- source_id: "更多目录｜专栏"
  current_route: high_value_but_noisy_source_repository_individual_verification
  decision_status: recovered_from_pr
  selected_reason: "专栏目录约 3360 个节点，是最丰富也最混杂的来源仓库，存在小众软件、阮一峰、腾讯安全、UX Coffee、博物志、PanSci、MacTalk、Indie Hackers 等价值候选。"
  rejected_alternatives:
    - "批量订阅专栏目录"
    - "因候选有价值就全部加入 Folo"
    - "忽略破解、下载站、影视福利、羊毛、代充、资源聚合和停更旧档案"
  route_reason: "专栏必须逐个核验是否仍在更新、是否有独立站与 RSS、内容是原创、转载、营销还是资源聚合、应进入 TopHub 发现层还是 Folo 持续关系层、与现有来源是否重复。"
  reuse_rule: "专栏是来源仓库，不是订阅池。具体来源必须逐个核验。"
  evidence_locator: "docs/ops/tophub-more-directories-audit-20260704.md#专栏目录"
  next_action: "已由专栏 94 页 snapshot 继续处理。"
```

---

## 11. 完整系统关系：所有入口各有角色

```yaml
- source_id: "TopHub 完整系统关系"
  current_route: system_map_for_all_entries
  decision_status: recovered_from_pr
  selected_reason: "文件把顶部分类目录、外部订阅、七个自定义分组、首页、话题聚合、日报、热文库、热点日历、追踪器、Folo、BibiGPT、沃壤全部放到同一关系图中。"
  rejected_alternatives:
    - "用一个入口承担全部功能"
    - "让 TopHub 自动沉积到沃壤"
    - "把 Folo、BibiGPT、沃壤互相替代"
  route_reason: "顶部分类目录发现站内来源；外部订阅补 TopHub 没有的来源；七组提供长期路由；首页短暂扫视；话题聚合看传播；日报回溯日期；热文库发现平台候选；热点日历提供日期灵感；追踪器只追精确词；Folo 保存少量连续关系；BibiGPT 处理一次性视频；沃壤沉积真正改变判断、领域或项目的内容。"
  reuse_rule: "每个入口只承担自己的角色。跨层流转必须有理由：发现、查看、核验、处理、复查、沉积。"
  evidence_locator: "docs/ops/tophub-public-temperature-and-group-map-20260704.md#完整系统关系"
  next_action: "作为总览路由图。"
```

---

## 12. 本补充 ledger 的复用规则

1. TopHub 有来源发现层和个人路由层，不能混为一谈。
2. `全部`页顺序核验回答当时是否放对，不回答当前是否有效。
3. 当前真实状态必须以最新平台核验和 `TopHub 全部订阅顺序.md` 为准。
4. 选择来源必须精确到节点口径，不按品牌整包订阅。
5. 顶部目录和更多目录都是候选仓库，不是新增命令。
6. 目录归类不决定个人分组；个人分组按用途路由。
7. 七组是长期路由容器，不是当前节点的永久封闭切片。
8. 公共温度看平台差异，不做事实核验。
9. 影游音乐用榜单发现变化，用编辑来源建立上下文。
10. 更多目录进入前必须带明确缺口，离开时只带走已核验来源。
11. 不为了把目录看完而翻页，不为总量整齐而补源。

---

## 13. 仍需继续追回

TopHub 总览、来源发现与路由边界已补成 ledger。

继续待办：

- 等用户本地提交 master checklist 与 coverage audit 后，再统一登记本文件和此前新增 ledger；
- 旧 ChatGPT 对话追索所有 partial 条目；
- 若继续补仓库内残余，可转入 `folo-subscription-diff / route review / guided review / handoff after AI` 这些 Folo 早期操作链。 
