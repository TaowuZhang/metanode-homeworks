# PR #372 判断链追回：综合门户聚合源（百度 / 腾讯 / 网易 / 搜狐 / 新浪 / ZAKER）（2026-07-06）

## 0. 边界

本文件是 `综合` 大类的补充 ledger，追回 PR #372 中已经闭环的门户聚合小标签：

- `docs/ops/tophub-comprehensive-baidu-closure-20260704.md`
- `docs/ops/tophub-comprehensive-tencent-closure-20260704.md`
- `docs/ops/tophub-comprehensive-netease-closure-20260704.md`
- `docs/ops/tophub-comprehensive-sohu-closure-20260704.md`
- `docs/ops/tophub-comprehensive-sina-closure-20260704.md`
- `docs/ops/tophub-comprehensive-zaker-closure-20260704.md`

本文件不声称综合大类全量完成。腾讯日报 11 个节点在 `tophub-comprehensive-tencent-daily-closure-20260704.md` 中另有完整判断，本文件只沿用腾讯 closure 中的“不重复评价同一内容池”结论。

判断链标准：

> 每个来源背后的判断链——为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 门户聚合源总规则

```yaml
- source_id: "综合门户聚合源总规则"
  current_route: mostly_rejected_or_on_demand
  decision_status: recovered_from_pr
  selected_reason: "百度、腾讯、网易、搜狐、新浪、ZAKER 中确实有少数按需入口和低频候选，但整体主要是门户内容池被点击、评论、滚动、视频、频道、热榜和垂直标签反复切片。"
  rejected_alternatives:
    - "把每个门户小标签当作独立来源"
    - "把点击榜 / 评论榜 / 滚动新闻 / 视频榜 / 垂直频道都加入 TopHub"
    - "把门户聚合页加入 Folo"
    - "给门户名、频道名或宽泛议题建立追踪器"
  route_reason: "TopHub 只保留能补出新功能的少数节点；门户聚合流大多按需或拒绝。Folo 不保存高频门户聚合流，只保存连续作者、编辑刊物、原始来源或低频深读关系。"
  reuse_rule: "以后遇到门户聚合页，先判断它是不是原始来源、稳定编辑产品或方法明确的资料库；如果只是同一内容池的热度 / 点击 / 评论 / 滚动 / 频道切片，就不常驻、不进 Folo、不建宽泛追踪。"
  evidence_locator: "tophub-comprehensive-baidu-closure-20260704.md; tophub-comprehensive-tencent-closure-20260704.md; tophub-comprehensive-netease-closure-20260704.md; tophub-comprehensive-sohu-closure-20260704.md; tophub-comprehensive-sina-closure-20260704.md; tophub-comprehensive-zaker-closure-20260704.md"
  next_action: "无；作为复用规则。"
```

---

## 2. 百度小标签

来源证据：`docs/ops/tophub-comprehensive-baidu-closure-20260704.md`

百度 17 个节点，两页全部查看。结论：不新增任何 TopHub 常驻节点，不进 Folo，不建百度相关宽泛追踪器。

```yaml
- source_id: "百度小标签整体"
  current_route: no_new_tophub_no_folo_no_tracker
  decision_status: recovered_from_pr
  selected_reason: "唯一具有独立内容属性的是 `百度用户体验中心`，但它更像设计方法与案例档案，适合任务型按需使用。"
  rejected_alternatives:
    - "百度｜实时热点"
    - "百度贴吧｜热议榜"
    - "百度视频电影 / 电视剧 / 综艺 / 动漫榜"
    - "好看视频｜今日热播榜"
    - "百度｜民生榜 / 财经榜 / 热梗榜 / 汽车榜 / 游戏榜 / 小说榜 / NBA焦点新闻"
    - "把百度用户体验中心做常驻"
  route_reason: "百度节点主要是搜索热度、贴吧讨论、视频聚合与垂直榜单；它们不是独立信息来源。百度用户体验中心有任务价值，但缺少稳定更新证据，不进入科技雷达或慢读常驻。"
  reuse_rule: "搜索门户的热度榜、贴吧榜、内容库榜只能按需观察搜索 / 社区 / 内容热度；只有可验证的独立方法档案才作为任务源。"
  evidence_locator: "docs/ops/tophub-comprehensive-baidu-closure-20260704.md#核心判断; #最终结论"
  next_action: "无；百度已闭环。"

- source_id: "百度｜实时热点 / 民生榜 / 财经榜 / 热梗榜"
  current_route: on_demand_search_and_heat_slices
  decision_status: recovered_from_pr
  selected_reason: "可在比较搜索引擎注意力、网络语汇或具体民生 / 财经事件传播时按需打开。"
  rejected_alternatives:
    - "公共温度常驻"
    - "数据与结构常驻"
    - "生活与社区常驻"
  route_reason: "这些榜单只是搜索和热度切片，不能替代事实来源、财经结构来源、民生服务来源或健康核验入口。"
  reuse_rule: "搜索热度只说明有人搜，不说明事实重要、证据等级或行动价值。"
  evidence_locator: "docs/ops/tophub-comprehensive-baidu-closure-20260704.md#热度与社区节点; #民生与财经榜"
  next_action: "按需。"

- source_id: "百度贴吧｜热议榜"
  current_route: on_demand_subculture_or_community_observation
  decision_status: recovered_from_pr
  selected_reason: "可用于研究具体贴吧社区或网络亚文化。"
  rejected_alternatives:
    - "公共温度常驻"
    - "生活与社区常驻"
  route_reason: "它有社区语言和亚文化观察价值，但黑话、侮辱性简称、群体标签、情绪化和阵营化表达较多，讨论热度不能代表事实可靠性。"
  reuse_rule: "社区热议榜只在研究具体社区时打开，不把群体语言和阵营表达变成默认公共温度。"
  evidence_locator: "docs/ops/tophub-comprehensive-baidu-closure-20260704.md#百度贴吧热议榜"
  next_action: "按需。"

- source_id: "好看视频｜今日热播榜"
  current_route: rejected_short_video_heat
  decision_status: recovered_from_pr
  selected_reason: "无。"
  rejected_alternatives:
    - "公共温度"
    - "军事"
    - "科技雷达"
    - "Folo"
  route_reason: "短视频热度混合军事与台海推测、真假难辨故事、投资市场信号、政策询问、美食家庭汽车体验和情绪化表达，无法替代事实核验。"
  reuse_rule: "短视频热播榜若以情绪和二手叙事为主，直接排除或按事件回原始来源。"
  evidence_locator: "docs/ops/tophub-comprehensive-baidu-closure-20260704.md#好看视频"
  next_action: "无；排除。"

- source_id: "百度用户体验中心"
  current_route: on_demand_design_method_archive
  decision_status: recovered_from_pr
  selected_reason: "提供用户研究、体验度量、品牌与交互设计案例，和普通新闻 / 热榜不同。"
  rejected_alternatives:
    - "科技雷达常驻"
    - "慢读与思想常驻"
    - "Folo"
  route_reason: "它是窄领域方法档案，当前页面缺少清晰日期，无法确认 2026 年仍稳定更新；处理 UX、满意度模型、交互设计和产品验证时按需查看。"
  reuse_rule: "方法档案有价值也不等于活跃订阅源；先看更新频率、日期和任务触发。"
  evidence_locator: "docs/ops/tophub-comprehensive-baidu-closure-20260704.md#百度用户体验中心"
  next_action: "按需；未来连续使用再核验订阅方式。"
```

---

## 3. 腾讯小标签

来源证据：`docs/ops/tophub-comprehensive-tencent-closure-20260704.md`

腾讯 21 个节点，两页全部查看。其中 11 个腾讯早报 / 晚报节点沿用 `综合 > 腾讯日报` 结论。

```yaml
- source_id: "腾讯小标签整体"
  current_route: no_new_tophub_with_folo_candidate
  decision_status: recovered_from_pr
  selected_reason: "腾讯目录中存在一个值得后续验证的独立低频来源：`腾讯研究院`。同时 `腾讯社交用户体验设计` 与 `腾讯科恩实验室官方博客` 有任务型资料价值。"
  rejected_alternatives:
    - "腾讯新闻热点榜"
    - "腾讯新闻热问"
    - "腾讯财经热点精选"
    - "腾讯科技滚动新闻"
    - "腾讯新闻谷雨实验室"
    - "腾讯新闻娱乐榜"
    - "腾讯早报 / 晚报 11 个切片"
    - "把腾讯整体加入 TopHub 或 Folo"
  route_reason: "腾讯门户热榜、日报和垂直切片重复同一内容池；独立研究 / 设计 / 安全档案按其角色分层，不与门户榜单并列。"
  reuse_rule: "大型平台公司下的研究院、设计团队、安全实验室要与新闻门户切开看：研究院可进低频 Folo 候选，设计 / 安全档案按任务，门户切片多半拒绝。"
  evidence_locator: "docs/ops/tophub-comprehensive-tencent-closure-20260704.md#核心结构判断; #最终结论"
  next_action: "腾讯日报 11 节点仍可回 tencent-daily closure 补细账。"

- source_id: "腾讯研究院"
  current_route: folo_low_frequency_deep_read_candidate
  decision_status: recovered_from_pr
  selected_reason: "讨论技术与人的关系，同时包含调查、圆桌、学者文章和文化产业观察；更新压力低于日更新闻流，更适合选择性深读。"
  rejected_alternatives:
    - "TopHub 常驻"
    - "立即加入 Folo"
    - "腾讯科技滚动新闻"
  route_reason: "它不是热度或快速发现入口，文章更适合形成连续来源关系；但腾讯机构立场需要与其他作者和研究来源并置，且需核验 RSS / 官方订阅 / 更新频率 / 真实阅读价值。"
  reuse_rule: "机构研究院进入 Folo 前必须先验证订阅方式、更新稳定性、是否真实阅读，以及是否只是标题有兴趣。"
  evidence_locator: "docs/ops/tophub-comprehensive-tencent-closure-20260704.md#腾讯研究院; #folo"
  next_action: "Folo 基线复查时验证。"

- source_id: "腾讯社交用户体验设计"
  current_route: on_demand_design_and_product_practice
  decision_status: recovered_from_pr
  selected_reason: "真实的产品设计与用户体验实践资料，覆盖关系链、字体、沉浸式体验、QQ 游戏中心、AIGC 内容生产和工具产品简化。"
  rejected_alternatives:
    - "科技雷达常驻"
    - "慢读与思想常驻"
    - "Folo"
  route_reason: "页面没有日期，无法确认更新稳定；又是很窄的专业栏目，适合处理产品设计、交互、字体、视觉或社交体验问题时按需查看。"
  reuse_rule: "企业设计团队栏目先作为任务资料库，不因方法价值直接常驻。"
  evidence_locator: "docs/ops/tophub-comprehensive-tencent-closure-20260704.md#腾讯社交用户体验设计"
  next_action: "按需；连续使用再核验更新。"

- source_id: "腾讯科恩实验室官方博客"
  current_route: on_demand_security_research_archive
  decision_status: recovered_from_pr
  selected_reason: "奔驰、雷克萨斯、特斯拉、宝马汽车安全研究，Wi-Fi 协议栈、车联网、iOS、macOS、Android Kernel、VMware 漏洞、利用和逃逸技术等内容有明确安全研究价值。"
  rejected_alternatives:
    - "科技雷达常驻"
    - "Folo"
  route_reason: "主要集中在 2017—2019 年，缺少近期持续更新证据；是高价值历史研究档案，不是活跃订阅源。"
  reuse_rule: "安全实验室旧档案价值不能误判为当前活跃来源；项目需要时按漏洞 / 技术点检索。"
  evidence_locator: "docs/ops/tophub-comprehensive-tencent-closure-20260704.md#腾讯科恩实验室官方博客"
  next_action: "按需。"

- source_id: "腾讯原创馆 / 腾讯新闻热问 / 谷雨实验室"
  current_route: on_demand_or_rejected_tencent_columns
  decision_status: recovered_from_pr
  selected_reason: "原创馆可按需找中文插画、字体和商业视觉案例；热问可观察公众追问；谷雨实验室部分题材有采访价值。"
  rejected_alternatives:
    - "视觉与自然常驻"
    - "慢读与思想常驻"
    - "Folo"
  route_reason: "原创馆缺持续更新证据且偏商业视觉；热问不能作为答案来源；谷雨标题持续放大痛苦、冲突和猎奇感，作为连续阅读队列负荷高。"
  reuse_rule: "栏目有个别文章价值时按篇判断，不建立持续队列。"
  evidence_locator: "docs/ops/tophub-comprehensive-tencent-closure-20260704.md#腾讯原创馆首页推荐; #腾讯新闻热问; #腾讯新闻谷雨实验室"
  next_action: "按需。"
```

---

## 4. 网易小标签

来源证据：`docs/ops/tophub-comprehensive-netease-closure-20260704.md`

网易 21 个节点，两页全部查看。结论：不新增任何 TopHub 常驻节点，不立即新增 Folo 来源；`网易人间｜记事` 是唯一低优先级、附条件 Folo 试用候选。

```yaml
- source_id: "网易小标签整体"
  current_route: no_new_tophub_with_conditional_folo_candidate
  decision_status: recovered_from_pr
  selected_reason: "网易目录中真正有辨识度的是 `数读` 的数据叙事和 `网易人间` 系列的普通人长篇与社会故事。"
  rejected_alternatives:
    - "网易滚动 / 热榜 / 今日关注"
    - "网易健康频道"
    - "网易人间多个入口同时订阅"
    - "网易娱乐 / 消费 / 情绪专栏常驻"
  route_reason: "门户新闻层情绪化、重复和案件化；数读需核验方法；网易人间内部重复且情绪负荷高。"
  reuse_rule: "人物故事和数据叙事有价值，也要看方法、重复、情绪负荷和现有来源缺口。"
  evidence_locator: "docs/ops/tophub-comprehensive-netease-closure-20260704.md#核心判断; #最终结论"
  next_action: "Folo 复查时再判断网易人间记事。"

- source_id: "网易专栏｜数读"
  current_route: on_demand_data_story_source
  decision_status: recovered_from_pr
  selected_reason: "主题与数据与结构相邻，覆盖大学、教师、研究生专业、城市、营养、职业和产业等数据叙事。"
  rejected_alternatives:
    - "数据与结构常驻"
  route_reason: "标题把复杂数据压缩成单一结论，仅凭标题页无法确认数据来源、时间范围、分母、因果与概括是否可靠；因此只按需，且必须核验原始数据和方法。"
  reuse_rule: "数据叙事标题越强，越要回到原始数据、分母和方法，不因‘数读’之名进入结构层。"
  evidence_locator: "docs/ops/tophub-comprehensive-netease-closure-20260704.md#网易专栏数读"
  next_action: "按需。"

- source_id: "网易人间｜记事"
  current_route: conditional_folo_trial_candidate
  decision_status: recovered_from_pr
  selected_reason: "以普通人的经历为主体，能看到教育、家庭、劳动、疾病、维权、城乡与组织生活，比门户快讯更接近真实生活的连续叙述。"
  rejected_alternatives:
    - "网易人间｜好读"
    - "网易人间｜特写"
    - "网易专栏｜人间"
    - "网易专栏｜大国小民"
    - "网易专栏｜看客"
    - "同时加入多个人间入口"
  route_reason: "疾病、家庭冲突、退学、裁员、监狱、死亡和维权占比高，情绪负荷重；多个入口复用同一文章池。只有 Folo 复查确认缺少普通人第一人称长篇时，才考虑记事一个节点短期试用。"
  reuse_rule: "普通人长篇来源不因‘真实’就加入；必须评估阅读率、情绪成本、与现有新闻调查 / 人物写作来源重复度。"
  evidence_locator: "docs/ops/tophub-comprehensive-netease-closure-20260704.md#网易人间记事; #folo"
  next_action: "Folo 基线复查。"

- source_id: "网易人间其他入口 / 看客 / 大国小民"
  current_route: rejected_or_on_demand_story_pool_duplicates
  decision_status: recovered_from_pr
  selected_reason: "个别作品可能有阅读价值。"
  rejected_alternatives:
    - "好读"
    - "特写"
    - "人间"
    - "大国小民"
    - "看客"
  route_reason: "多个入口大量复用同一文章池，且疾病、死亡、犯罪、家庭冲突、贫困、创伤和身份反差主题过重，不适合作为多节点订阅。"
  reuse_rule: "同一故事池多入口只选一个候选；其他入口按主题检索。"
  evidence_locator: "docs/ops/tophub-comprehensive-netease-closure-20260704.md#网易人间好读; #网易专栏看客; #网易人间特写; #网易专栏人间; #网易专栏大国小民"
  next_action: "按需。"

- source_id: "网易门户新闻 / 健康 / 娱乐消费专栏"
  current_route: rejected_or_on_demand
  decision_status: recovered_from_pr
  selected_reason: "部分栏目有偶发主题价值。"
  rejected_alternatives:
    - "网易新闻滚动 / 实时热榜 / 今日关注"
    - "网易健康频道"
    - "轻松一刻 / 浪潮 / 槽值 / 下划线 / 谈心社 / 胖编怪聊 / 曲一刀 / 今日之声 / 三三有梗 / 哒哒"
  route_reason: "新闻层放大犯罪、冲突和奇闻；健康角色已有科普中国辟谣与半月谈健康；娱乐消费情绪专栏要么标题污染、陈旧、高负荷，要么更适合按主题检索。"
  reuse_rule: "门户专栏若靠情绪、身份反差、猎奇或品牌合作维持，不常驻。"
  evidence_locator: "docs/ops/tophub-comprehensive-netease-closure-20260704.md#门户新闻层; #网易健康频道; #明确不新增的娱乐与消费专栏; #最终结论"
  next_action: "按需或排除。"
```

---

## 5. 搜狐小标签

来源证据：`docs/ops/tophub-comprehensive-sohu-closure-20260704.md`

搜狐 6 个节点，页面完整展示。结论：不新增任何搜狐节点，不进 Folo，不建搜狐相关宽泛追踪器。

```yaml
- source_id: "搜狐小标签整体"
  current_route: rejected_portal_reaggregation
  decision_status: recovered_from_pr
  selected_reason: "无常驻价值。"
  rejected_alternatives:
    - "搜狐｜热评榜"
    - "搜狐｜24小时热点"
    - "搜狐｜搜狐热榜"
    - "搜狐｜科普榜"
    - "搜狐｜视频榜"
    - "搜狐｜体育榜"
  route_reason: "六个节点都是搜狐门户对外部媒体、自媒体和视频内容的再次聚合与热度排序，没有独立编辑来源能力。"
  reuse_rule: "门户再聚合不因换成热评、热点、视频、科普、体育标签而变成新来源。"
  evidence_locator: "docs/ops/tophub-comprehensive-sohu-closure-20260704.md#核心判断; #最终结论"
  next_action: "无；搜狐已闭环。"

- source_id: "搜狐｜科普榜"
  current_route: rejected_misleading_science_source
  decision_status: recovered_from_pr
  selected_reason: "无。"
  rejected_alternatives:
    - "科技雷达"
    - "数据与结构"
    - "生活与社区"
  route_reason: "标题混合误导性比喻、未经充分说明的假设、宇宙学与气候学过度简化、点击诱导表达，可靠科学新闻与伪科学式标题并列；不能作为科普来源。"
  reuse_rule: "带‘科普’标签但标题方法污染严重时，直接排除，回科普中国、果壳、论文、研究机构或权威科学媒体。"
  evidence_locator: "docs/ops/tophub-comprehensive-sohu-closure-20260704.md#搜狐科普榜"
  next_action: "无；明确排除。"

- source_id: "搜狐｜视频榜 / 体育榜"
  current_route: rejected_video_and_sports_heat
  decision_status: recovered_from_pr
  selected_reason: "无。"
  rejected_alternatives:
    - "公共温度"
    - "军事"
    - "影游音乐"
    - "Folo"
  route_reason: "视频榜以战争、外交、台海和冲突性评论视频为主，体育榜是高频赛事、比分、转会和赛后争议流；都不符合当前七分组职责。"
  reuse_rule: "视频化不提升事实质量；体育快讯需用赛事、联盟或球队官方来源。"
  evidence_locator: "docs/ops/tophub-comprehensive-sohu-closure-20260704.md#搜狐视频榜; #搜狐体育榜"
  next_action: "无。"
```

---

## 6. 新浪小标签

来源证据：`docs/ops/tophub-comprehensive-sina-closure-20260704.md`

新浪 41 个节点，四页全部查看。结论：不新增任何 TopHub 常驻节点，不进 Folo，不建新浪相关追踪器。

```yaml
- source_id: "新浪小标签整体"
  current_route: no_new_tophub_with_on_demand_nodes
  decision_status: recovered_from_pr
  selected_reason: "少数栏目有有限按需用途：黑猫投诉、苹果汇、创事记、财经 7×24、全部滚动、图片排行、娱乐点击榜。"
  rejected_alternatives:
    - "点击量排行 / 评论数排行 / 滚动新闻 / 热榜 / 垂直热榜 / 新浪财经 / 新浪科技 / 新浪汽车系列常驻"
    - "新浪 Folo"
    - "新浪宽泛追踪器"
  route_reason: "41 个节点主要是同一门户内容池按点击量、评论数、滚动时间、热榜热度和栏目标签反复切片；不能形成必须加入个人首页的新能力。"
  reuse_rule: "大型门户的点击、评论、滚动和热榜切片，即使数量多，也通常只是一套内容池重复加权。"
  evidence_locator: "docs/ops/tophub-comprehensive-sina-closure-20260704.md#核心结论; #最终结论"
  next_action: "无；新浪已闭环。"

- source_id: "黑猫投诉｜热点追踪"
  current_route: on_demand_consumer_rights_source
  decision_status: recovered_from_pr
  selected_reason: "新浪组中最高价值的按需节点，触及消费权益和服务纠纷，如酒店退房、演唱会黄牛、智能眼镜隐私、空调维权、外卖打包费、商标争议等。"
  rejected_alternatives:
    - "生活与社区常驻"
    - "立即替换现有节点"
  route_reason: "当前展示主要是媒体文章而非结构化投诉数据或原始投诉流，选题仍受热点传播影响，与每周质量报告有重叠，且常驻投诉风险流容易累积不可行动信息。"
  reuse_rule: "消费维权缺口反复出现时，才考虑替换已有低价值节点；不扩张节点总数。"
  evidence_locator: "docs/ops/tophub-comprehensive-sina-closure-20260704.md#黑猫投诉热点追踪"
  next_action: "具体品牌、平台、酒店、票务、售后、收费或隐私纠纷时按需。"

- source_id: "新浪科技｜苹果汇 / 创事记"
  current_route: on_demand_apple_or_industry_topic_discovery
  decision_status: recovered_from_pr
  selected_reason: "苹果汇在 Apple 发布会、中国区功能或企业生态问题时有按需价值；创事记可做特定产业长文题目发现。"
  rejected_alternatives:
    - "科技雷达常驻"
    - "Folo"
  route_reason: "苹果汇与 IT之家、爱范儿、少数派、The Verge 和 Apple 相关追踪器重复；创事记作者、利益关系和内容纯度不稳定。"
  reuse_rule: "产业题目发现页不等于稳定深读来源；先看具体作者、原媒体、数据与利益关系。"
  evidence_locator: "docs/ops/tophub-comprehensive-sina-closure-20260704.md#新浪科技苹果汇; #新浪科技创事记"
  next_action: "按需。"

- source_id: "新浪财经｜7×24小时全球实时财经新闻直播 / 新浪｜全部滚动新闻"
  current_route: on_demand_event_timeline_or_news_lead
  decision_status: recovered_from_pr
  selected_reason: "重大财经、国际、灾害事件时可临时查看时间线或寻找上游来源。"
  rejected_alternatives:
    - "数据与结构常驻"
    - "Folo"
  route_reason: "频率高、范围宽，混有市场评论、公司股价、论坛发言、体育 AI 营销等；只能作为事件发生时的时间线或线索页，看到来源后回原始媒体。"
  reuse_rule: "7×24 和全部滚动只在事件期查线索，不从头读、不订阅。"
  evidence_locator: "docs/ops/tophub-comprehensive-sina-closure-20260704.md#新浪全部滚动新闻; #新浪财经7×24小时全球实时财经新闻直播"
  next_action: "按事件。"

- source_id: "新浪评论数榜 / 体育滚动 / 博客排行 / 科学探索 / 汽车热搜"
  current_route: rejected_polluted_or_invalid_sina_nodes
  decision_status: recovered_from_pr
  selected_reason: "无。"
  rejected_alternatives:
    - "公共温度"
    - "科技雷达"
    - "影游音乐"
    - "数据与结构"
  route_reason: "评论数榜时间口径失效，体育滚动含博彩彩票内容，博客排行陈旧且有交易诱导，科学探索标签失真，汽车热搜受营销主导。"
  reuse_rule: "历史累计污染、博彩污染、交易诱导、栏目标签失真和营销主导是明确排除信号。"
  evidence_locator: "docs/ops/tophub-comprehensive-sina-closure-20260704.md#评论数排行; #明确排除的失真节点; #最终结论"
  next_action: "无；排除。"
```

---

## 7. ZAKER 小标签

来源证据：`docs/ops/tophub-comprehensive-zaker-closure-20260704.md`

ZAKER 15 个节点，两页全部查看。结论：不新增任何 TopHub 常驻节点，不进 Folo，不建 ZAKER 相关追踪器。

```yaml
- source_id: "ZAKER小标签整体"
  current_route: no_new_tophub_with_on_demand_nodes
  decision_status: recovered_from_pr
  selected_reason: "ZAKER 看起来比一般门户更像编辑型资讯聚合器，其中精读新闻、互联网、游戏、北京有有限按需价值。"
  rejected_alternatives:
    - "ZAKER 24小时热榜"
    - "ZAKER 科技 / 财经 / 健康 / 国际 / 科学 / 电影 / 体育 / 权威发布等常驻"
    - "ZAKER Folo"
  route_reason: "大量栏目只是对其他媒体、自媒体和机构稿件的再次聚合；标题依赖夸张、冲突、性暗示、身份反差和情绪化判断；多个垂直频道后半段被财经、专利、汽车、IPO、股票和营销稿污染。"
  reuse_rule: "聚合器即使有编辑感，也要看是否回到原始媒体、栏目是否被污染、标题是否过度情绪化。"
  evidence_locator: "docs/ops/tophub-comprehensive-zaker-closure-20260704.md#核心判断; #最终结论"
  next_action: "无；ZAKER 已闭环。"

- source_id: "ZAKER｜精读新闻 / 互联网 / 游戏 / 北京"
  current_route: on_demand_zaker_nodes
  decision_status: recovered_from_pr
  selected_reason: "精读新闻可发现长文题目，互联网可看特定 AI / 互联网公司或产业议题，游戏可偶尔发现触乐、独立游戏或产业文章，北京可用于明确北京本地任务。"
  rejected_alternatives:
    - "常驻 TopHub"
    - "Folo"
  route_reason: "这些节点仍是二次聚合层，标题和频道边界不稳定；有题目价值时应打开原文，判断原始媒体。"
  reuse_rule: "二次聚合器的题目发现可以用，但沉积应回原文和原始来源。"
  evidence_locator: "docs/ops/tophub-comprehensive-zaker-closure-20260704.md#ZAKER精读新闻; #ZAKER互联网; #ZAKER游戏; #ZAKER北京; #tophubfolo与追踪器"
  next_action: "按需。"

- source_id: "ZAKER｜健康有术 / 科学 / 电影 / 体育 / 娱乐 / 权威发布"
  current_route: rejected_polluted_zaker_channels
  decision_status: recovered_from_pr
  selected_reason: "无常驻价值。"
  rejected_alternatives:
    - "生活与社区"
    - "科技雷达"
    - "视觉与自然"
    - "影游音乐"
    - "公共温度"
  route_reason: "健康有术含误导健康标题且后半段被商业财经污染；科学频道不是稳定科学来源；电影和娱乐依赖流量与争议；体育高度重复赛事和情绪化拆分；权威发布标签不能保证每条来自权威源。"
  reuse_rule: "栏目名不能当证据；频道后半段被无关商业 / 财经 / IPO / 汽车污染时，不进领域分组。"
  evidence_locator: "docs/ops/tophub-comprehensive-zaker-closure-20260704.md#ZAKER健康有术; #ZAKER科学; #ZAKER电影; #ZAKER体育; #ZAKER娱乐; #ZAKER权威发布"
  next_action: "无；排除或不新增。"
```

---

## 8. 本补充 ledger 的复用规则

1. 门户聚合页不是原始来源；先看它是否只是外部媒体 / 自媒体 / 机构稿的二次排序。
2. 点击、评论、滚动、视频、热榜、频道标签通常只是同一内容池的不同切片。
3. 栏目名不能替代证据等级；“科普”“科学”“健康”“权威发布”“财经”“AI”“ESG”都必须看实际内容和方法。
4. 有价值的独立栏目优先按任务使用，进入 Folo 前必须验证稳定更新、订阅方式、真实阅读价值和情绪负荷。
5. 情绪化标题、博彩彩票、荐股交易、旧闻累计、栏目污染、营销主导，是拒绝常驻的强信号。
6. 低频深读候选只能一进一出复查，不因为标题有兴趣就加入 Folo。
7. 若二次聚合层发现好题目，应回到原始媒体、作者、数据、政策、公司或事件本身。

---

## 9. 仍需继续追回

综合大类已经补入：

- closeout 三项动作；
- 知乎小标签；
- 热搜小标签；
- 滚动新闻小标签；
- 微信 / 微博 / 今日头条；
- 百度 / 腾讯 / 网易 / 搜狐 / 新浪 / ZAKER。

仍需继续追回：

- 人民网 / 新华 / 南方周末 / 半月谈 / 央视；
- 健康 / 军事 / 房产 / 地方城市；
- 腾讯日报 11 节点的细账；
- `腾讯研究院` 与 `网易人间｜记事` 在 Folo 复查时的更完整条件。

综合大类仍是 `partial_recovered_from_pr`，不能写成全量完成。
