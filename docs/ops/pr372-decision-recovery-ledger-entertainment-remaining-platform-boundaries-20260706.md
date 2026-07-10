# PR #372 判断链追回：娱乐剩余平台边界补充（2026-07-06）

## 0. 边界

本文件补充追回 PR #372 中娱乐大类若干容易被总类规则吞没的单页 closure 判断链。

纳入文件：

- `docs/ops/tophub-entertainment-dongqiudi-closure-20260705.md`
- `docs/ops/tophub-entertainment-douyin-closure-20260705.md`
- `docs/ops/tophub-entertainment-kuaishou-closure-20260705.md`
- `docs/ops/tophub-entertainment-weread-closure-20260705.md`
- `docs/ops/tophub-entertainment-penti-closure-20260705.md`
- `docs/ops/tophub-entertainment-final-reconciliation-20260705.md`

边界说明：

- 娱乐最终动作、平台体育小说播客边界、影视音乐游戏网文平台边界此前已补 ledger。
- 本文件只补剩余单页中仍有独立判断价值的边界：足球文化按需、短视频平台探针、快手按需研究、微信读书新书榜、喷嚏乐影替换开眼。
- 这些文件中出现的 73 / 74 等数量是当时阶段快照；当前真实状态以 `资源/雷达/TopHub 全部订阅顺序.md` 为准。
- 当前不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 懂球帝：深度有价值，但足球不是长期未读流

```yaml
- source_id: "懂球帝｜深度"
  current_route: high_priority_on_demand_football_culture_and_tactics_entry
  decision_status: recovered_from_pr
  selected_reason: "`深度`明显优于懂球帝其他节点，内容可包含世界杯史、战术复盘、球衣技术史，小国、移民、族群与足球之间的关系，足球与政治符号、国家认同和地方社会，以及海外媒体翻译的球员和球迷故事。"
  rejected_alternatives:
    - "新增懂球帝｜深度为 TopHub 常驻"
    - "把懂球帝今日头条、头条新闻、热门新闻也一起订阅"
    - "把懂球帝榜单加入 Folo"
    - "因重大赛事信息密集而建立足球长期分组"
  route_reason: "深度确有文化阅读价值，但用户当前没有持续足球阅读需求，现有慢读与思想已承担完整阅读窗口；为单一运动增加长期节点会提高固定噪声。它更适合重大赛事、足球文化史、战术和国家社会议题研究时按需打开。"
  reuse_rule: "体育平台中最好的深度栏目也不自动常驻；只有长期赛事 / 俱乐部 / 作者关系成立，才考虑订阅。"
  evidence_locator: "docs/ops/tophub-entertainment-dongqiudi-closure-20260705.md#结论; #按需入口"
  next_action: "按需。"
```

---

## 2. 懂球帝其他节点：足球新闻榜和俱乐部榜不可靠

```yaml
- source_id: "懂球帝｜今日头条 / 早报 / 头条新闻 / 热门新闻 / 联赛与俱乐部节点"
  current_route: rejected_or_event_window_only
  decision_status: recovered_from_pr
  selected_reason: "这些节点能在重大赛事期间提供中文足球关注焦点。"
  rejected_alternatives:
    - "长期订阅今日头条、早报、头条新闻、热门新闻"
    - "订阅中超、山东鲁能泰山、德甲、AC米兰等俱乐部或联赛节点"
  route_reason: "今日头条、头条新闻、热门新闻围绕世界杯、转会、采访和赛果反复排列；早报高度依赖梗、情绪和既有比赛语境；中超、山东鲁能泰山、德甲、AC米兰等节点存在论坛短帖、更新完整性可疑或俱乐部边界不可靠问题。"
  reuse_rule: "体育榜单用于赛事窗口确认注意力，不作为常驻；俱乐部节点只有明确长期追踪某队时才可考虑。事实仍回赛事、俱乐部、记者或官方来源。"
  evidence_locator: "docs/ops/tophub-entertainment-dongqiudi-closure-20260705.md#结论; #最终动作"
  next_action: "排除常驻。"
```

---

## 3. 抖音总榜：保留一个平台探针，不展开分区榜

```yaml
- source_id: "抖音｜总榜"
  current_route: keep_public_temperature_short_video_probe
  decision_status: recovered_from_pr
  selected_reason: "`抖音｜总榜` 能用一个节点观察中文短视频平台的总体传播温度，同时显现文旅、生活、广告、剧情、非遗、萌宠和社会情绪；与微博、知乎、小红书分别提供不同平台的注意力结构。"
  rejected_alternatives:
    - "用抖音热点榜替换总榜"
    - "新增抖音分区榜"
    - "把抖音作为作品策展入口"
    - "把抖音榜单加入 Folo"
  route_reason: "总榜的混杂性正是公共温度价值：它能看见短视频平台正在采用什么叙事、画面、人物和商业传播形式。热点榜只给事件词与话题热度，信息维度少于总榜。"
  reuse_rule: "短视频平台只保留一个总探针观察平台温度；不要用分区榜模拟专业入口。"
  evidence_locator: "docs/ops/tophub-entertainment-douyin-closure-20260705.md#三个问题; #节点判断"
  next_action: "保留。"
```

---

## 4. 抖音分区榜：标签名不能保证专业质量

```yaml
- source_id: "抖音｜热点榜及 25 个分区榜"
  current_route: on_demand_or_rejected_platform_slices
  decision_status: recovered_from_pr
  selected_reason: "分区榜偶尔可发现具体作者、非遗、乡村手艺、公益项目、手工创造、自然影像、地方生活或平台玩法。"
  rejected_alternatives:
    - "订阅热点榜"
    - "订阅财经榜、科技榜、文化教育榜"
    - "订阅娱乐、明星、时尚、剧情、图文控等粉圈 / 平台玩法榜"
    - "订阅游戏榜、二次元榜"
    - "订阅旅行、美食、户外、汽车等消费导向榜"
    - "由单次上榜直接把创作者加入 Folo"
  route_reason: "财经、科技、文化教育等看似专业的节点标签漂移严重；娱乐、明星、剧情等商业植入、粉圈、剧情营销和平台玩法密集；游戏与二次元主要提供短周期热度，不能替代游研社、机核、B站每周必看或真正作品来源；公益榜混杂不同性质内容，涉及真实项目时必须回具体组织和可核验证据。"
  reuse_rule: "平台分区榜只能发现线索。具体作者、机构或项目必须跨页面反复出现并可核验，才进入 Folo 候选；平台分类名不能当专业质量保证。"
  evidence_locator: "docs/ops/tophub-entertainment-douyin-closure-20260705.md#节点判断"
  next_action: "按需或排除。"
```

---

## 5. 快手：有差异人群样本，但不再占公共温度位

```yaml
- source_id: "快手小标签 10 节点整体"
  current_route: on_demand_platform_population_and_business_ecology_observation
  decision_status: recovered_from_pr
  selected_reason: "快手能提供与抖音不同的人群和传播样本，尤其是地方社会新闻、下沉市场、农村与小城生活、短剧、挑战文化及平台商业生态。"
  rejected_alternatives:
    - "新增快手实时热榜或社会榜为公共温度"
    - "用快手替换抖音总榜"
    - "把快手短剧榜放入影游音乐"
    - "把快手平台榜加入 Folo"
  route_reason: "已有 `抖音｜总榜` 作为短视频公共温度入口，快手实时热榜和社会榜又与微博、抖音、百度、知乎及其他公共温度节点重叠；搜索飙升榜噪声大；文娱榜、短剧榜、挑战榜、月销榜都受平台玩法、粉圈、短剧口径、电商销量和营销影响。"
  reuse_rule: "平台人群差异有研究价值，但不等于必须常驻。研究快手用户结构、短剧市场、直播电商或下沉市场时按需打开。"
  evidence_locator: "docs/ops/tophub-entertainment-kuaishou-closure-20260705.md#三个问题; #最终动作"
  next_action: "按需。"
```

---

## 6. 快手指数洞察报告：研究入口，不进数据与结构

```yaml
- source_id: "快手指数｜洞察报告"
  current_route: on_demand_platform_report_library
  decision_status: recovered_from_pr
  selected_reason: "它是快手节点中最有研究价值的入口，覆盖新银发人群、汽车、内容消费、服饰、健康、美妆、小镇青年、直播电商和平台生态。"
  rejected_alternatives:
    - "加入数据与结构常驻"
    - "把快手商业报告当独立研究或监管数据"
    - "持续订阅快手报告库"
  route_reason: "它由快手及其商业化体系发布，带有营销目的；报告年份跨度大，大量旧报告长期留在列表；行业分类重复，不能替代独立研究、监管数据或企业原始财务信息。当前没有持续追踪这些行业报告的明确任务。"
  reuse_rule: "平台商业报告库按研究问题调用，必须检查发布方、年份、样本、商业目的和可替代数据。"
  evidence_locator: "docs/ops/tophub-entertainment-kuaishou-closure-20260705.md#节点判断"
  next_action: "按需。"
```

---

## 7. 微信读书新书榜：保留一个平台实际阅读启动信号

```yaml
- source_id: "微信读书｜新书榜"
  current_route: keep_slow_read_recent_platform_reading_signal
  decision_status: recovered_from_pr
  selected_reason: "它能看到微信读书平台近期进入书库并开始获得阅读的作品，样本中同时包含新出版图书、杂志期刊、社会纪实、历史作品、小说与影视原著。"
  rejected_alternatives:
    - "取消微信读书新书榜"
    - "用微信读书总榜或热搜榜替代"
    - "新增微信读书总榜、飙升榜、男生小说榜、女生小说榜、热搜榜、神作榜、神作潜力榜、小说榜"
    - "把微信读书榜单加入 Folo"
  route_reason: "新书榜提供的是微信读书平台上最近出现了什么，与豆瓣新书速递的编辑 / 文化口径不同，能和豆瓣新书速递、豆瓣书评形成‘新进入平台—出版文化筛选—读后观点’的分工。一个节点已足够承担微信读书动态发现。"
  reuse_rule: "阅读平台只保留一个动态发现入口即可；具体作者、出版社、刊物或评论媒体要等持续阅读关系成立后再判断 Folo。"
  evidence_locator: "docs/ops/tophub-entertainment-weread-closure-20260705.md#三个问题; #节点判断"
  next_action: "保留。"
```

---

## 8. 微信读书其他榜：平台推荐值不能冒充质量结论

```yaml
- source_id: "微信读书其他 8 个榜"
  current_route: on_demand_or_rejected_reading_platform_slices
  decision_status: recovered_from_pr
  selected_reason: "总榜、飙升榜、男生 / 女生小说榜和小说榜在特定阅读场景中有按需用途。"
  rejected_alternatives:
    - "新增总榜"
    - "新增飙升榜"
    - "新增男生小说榜或女生小说榜"
    - "新增神作榜 / 神作潜力榜"
    - "新增热搜榜"
  route_reason: "总榜变化慢，接近大众经典书架；飙升榜受短期传播、影视、促销、自助、命理、健康和平台活动影响；男女小说榜与起点、纵横网文榜高度重叠；小说榜混合出版文学、影视原著、网文和畅销经典；神作榜和神作潜力榜的推荐值未呈现样本规模、评价分布和冷启动机制，还混有命理、自助、理财和轻量内容。"
  reuse_rule: "平台推荐值是平台用户反馈，不是高质量结论。网文与小说榜按找书任务临时查，不常驻。"
  evidence_locator: "docs/ops/tophub-entertainment-weread-closure-20260705.md#节点判断; #与起点、纵横和豆瓣的关系"
  next_action: "按需或排除。"
```

---

## 9. 喷嚏乐影：替换开眼，不是再加电影热度榜

```yaml
- source_id: "喷嚏网｜乐影"
  current_route: executed_tophub_entertainment_replacement_documentary_and_historical_video_discovery
  decision_status: recovered_from_pr_platform_action_planned_then_executed_in_final_reconciliation
  selected_reason: "本轮提供清晰、可识别的纪录片与历史影像发现，覆盖政治史与人物、国际关系、科技与工业、艺术与文明、社会纪录等主题，给出具体作品名称、年份与主题，用户可以继续查找正式观看来源。"
  rejected_alternatives:
    - "继续保留开眼视频日报"
    - "新增喷嚏网本周热读、图卦、财经风云、铂程斋、乐活、最新"
    - "把乐影加入 Folo"
  route_reason: "开眼视频日报更偏短视频综合推荐，长期密度和主题稳定性较弱。乐影补的是值得找来看的纪录片和历史影像，不是电影热度榜。它只承担发现，不承担事实认证、版权判断或观看入口保证。"
  reuse_rule: "纪录片 / 历史影像发现源可进入影游音乐；但具体作品仍需另行核验来源、版本和版权。平台栏目不是稳定一手制作机构，通常不进 Folo。"
  evidence_locator: "docs/ops/tophub-entertainment-penti-closure-20260705.md#结论; docs/ops/tophub-entertainment-final-reconciliation-20260705.md#最终动作"
  next_action: "已在娱乐最终执行中加入。"
```

---

## 10. 喷嚏其他栏目：公共拼贴和文化材料只能按需

```yaml
- source_id: "喷嚏网其他 6 个栏目"
  current_route: on_demand_or_rejected_mixed_aggregation
  decision_status: recovered_from_pr
  selected_reason: "图卦、铂程斋、本周热读在特定情境下可能有公共情绪、文化材料或站内聚合观察价值。"
  rejected_alternatives:
    - "新增本周热读"
    - "新增图卦"
    - "新增财经风云"
    - "新增铂程斋"
    - "新增乐活"
    - "新增最新"
    - "把铂程斋加入 Folo"
  route_reason: "本周热读混排财经快讯、AI 文章、小说、Suno / NotebookLM 生成内容和商品优选；图卦有公共情绪价值但依赖语境、图片和二次转载；财经风云标题情绪化，不替代数据、公告和财经事实源；铂程斋混有 AI 生成内容、商品优选及站内重复，出处与制作方式不稳定；乐活和最新用途边界不清或重复度最高。"
  reuse_rule: "二次聚合站栏目按具体用途临时打开；涉及事实、财经、AI 生成和长文转载时必须核验作者、原始出处、制作方式和来源边界。"
  evidence_locator: "docs/ops/tophub-entertainment-penti-closure-20260705.md#为什么其他栏目不新增; #按需入口"
  next_action: "按需或排除。"
```

---

## 11. 本补充 ledger 的复用规则

1. 体育平台深度栏目可以按需高优先，但不因单项运动建立长期未读流。
2. 短视频平台保留一个总探针即可，分区榜不能冒充专业入口。
3. 平台分类名不等于内容质量；越像专业标签，越要检查标签漂移。
4. 快手有平台人群差异研究价值，但不因此占第二个短视频公共温度位。
5. 平台商业报告库按研究问题调用，不进数据与结构常驻。
6. 微信读书新书榜保留，因为它补平台实际阅读启动信号；其他榜按需或排除。
7. 平台推荐值、神作值、搜索热度都不是质量结论。
8. 喷嚏乐影替换开眼，补纪录片与历史影像发现，不承担事实、版权或观看入口保证。
9. 二次聚合站可作线索，但作者、出处、制作方式和 AI 生成风险必须核验。
10. 平台榜单通常不进 Folo；只有具体作者、节目、机构或栏目反复证明独特价值，才进入复查候选。

---

## 12. 仍需继续追回

娱乐剩余平台边界已补成 ledger。

继续待办：

- 等用户本地提交 master checklist 与 coverage audit 后，再统一登记本文件和此前新增 ledger；
- 仓库内剩余可查项可继续检查 `tophub-community-and-podcast-directory-audit-20260704.md`、`tophub-development-directory-audit-20260704.md`、`tophub-finance-final-closure-20260705.md` 等是否已被各专项 ledger 完整覆盖；
- 旧 ChatGPT 对话追索所有 partial 条目仍待处理。
