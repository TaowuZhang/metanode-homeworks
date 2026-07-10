# PR #372 判断链追回：综合热搜与滚动新闻（2026-07-06）

## 0. 边界

本文件是 `pr372-decision-recovery-ledger-comprehensive-20260706.md` 的补充 ledger。

它只追回两个已经在 PR #372 中闭环的小标签：

- `docs/ops/tophub-comprehensive-hot-search-closure-20260704.md`
- `docs/ops/tophub-comprehensive-rolling-news-closure-20260704.md`

本文件不声称综合大类全量完成。

本文件围绕同一个判断链问题：

> 每个来源背后的判断链——为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 热搜小标签整体判断

来源证据：`docs/ops/tophub-comprehensive-hot-search-closure-20260704.md`

`热搜` 小标签共 16 个节点，两页已经全部查看。

最终结论：

- 不新增任何节点；
- 不替换现有节点；
- 保留 `微博｜热搜榜`；
- 保留现有 `抖音｜总榜`；
- 不进入 Folo；
- 不建立“热搜”“百度热搜”“微博热搜”“抖音热点”等宽泛追踪器。

```yaml
- source_id: "热搜小标签整体"
  current_route: no_new_tophub_no_folo_no_tracker
  decision_status: recovered_from_pr
  selected_reason: "现有公共温度已经通过微博热搜、知乎热榜、抖音总榜、小红书热榜、微信热文、实时榜中榜、编辑媒体和官方议程来源覆盖主要注意力结构。"
  rejected_alternatives:
    - "百度、夸克、搜狗、360、UC 等搜索引擎热搜"
    - "抖音｜热点榜"
    - "知乎｜热搜"
    - "头条搜索｜猜你想搜"
    - "腾讯视频、爱奇艺、哔哩哔哩热搜"
    - "中国搜索热搜"
    - "新浪汽车热搜"
    - "头条热榜总榜"
    - "封面热搜"
    - "热搜相关 Folo 来源"
    - "热搜宽泛追踪器"
  route_reason: "热搜榜能提供的主要增量是搜索词如何命名事件、某个事件被拆成多少近义词、搜索门户偏好的标题形态。这些关系偶尔比较有用，但不足以让多个热搜榜长期占据首页。"
  reuse_rule: "以后遇到新的热搜榜，先判断它是否提供新的注意力结构；如果只是同一事件的关键词拆分、平台内搜索或营销榜，就按需，不常驻。"
  evidence_locator: "docs/ops/tophub-comprehensive-hot-search-closure-20260704.md#核心结论; #十七为什么不需要再加一个搜索引擎热搜; #十八tophubfolo与追踪器; #最终结论"
  next_action: "无；热搜小标签已闭环。"
```

---

## 2. 已保留热搜入口

```yaml
- source_id: "微博｜热搜榜"
  current_route: tophub_public_temperature
  decision_status: recovered_from_pr
  selected_reason: "显示中文社交媒体即时情绪、人物关注和话题拆分方式。它不是提供事实，而是显示社交平台如何切割事件、放大人物、组织情绪词和粉圈内容。"
  rejected_alternatives:
    - "新浪热榜等同源榜单"
    - "再增加多个搜索引擎热搜"
    - "把微博热搜放入 Folo"
  route_reason: "继续放在 `公共温度`；它承担即时情绪和话题拆分观察，不承担事实核验。"
  reuse_rule: "以后遇到社交平台热榜，如果它只复制微博式即时情绪或粉圈话题，不新增；只有提供不同平台结构且当前七分组未覆盖，才考虑替换而不是叠加。"
  evidence_locator: "docs/ops/tophub-comprehensive-hot-search-closure-20260704.md#二微博热搜榜"
  next_action: "无；保留。"

- source_id: "抖音｜总榜"
  current_route: tophub_public_temperature
  decision_status: recovered_from_pr
  selected_reason: "现有节点已经承担抖音平台关系，覆盖事件、生活方式、音乐模板、视频挑战与平台内容结构。"
  rejected_alternatives:
    - "抖音｜热点榜"
  route_reason: "继续保留现有 `抖音｜总榜`，不与 `抖音｜热点榜` 并存。若未来真实使用证明总榜偏离公共热点，再在二者之间做替换评估，而不是叠加。"
  reuse_rule: "以后同平台出现多个榜单，优先一进一出比较，不并存叠加。"
  evidence_locator: "docs/ops/tophub-comprehensive-hot-search-closure-20260704.md#七抖音热点榜; #十八tophubfolo与追踪器"
  next_action: "无；保留总榜，不新增热点榜。"
```

---

## 3. 搜索引擎热搜为什么不选

```yaml
- source_id: "百度｜实时热点"
  current_route: on_demand_search_split_observation
  decision_status: recovered_from_pr
  selected_reason: "可在需要比较搜索引擎如何拆分某一事件时按需打开。"
  rejected_alternatives:
    - "常驻公共温度"
    - "百度热搜追踪器"
  route_reason: "样本中同一足球比赛被拆成大量近义词条，且热度数字呈整齐递减，不能当作可信的跨议题测量尺度。它更像搜索需求与内容分发共同制造的词条扩张。"
  reuse_rule: "搜索引擎热搜若只是把同一事件拆成关键词，并给出不可比较的热度数字，不进常驻。"
  evidence_locator: "docs/ops/tophub-comprehensive-hot-search-closure-20260704.md#三百度实时热点"
  next_action: "按需。"

- source_id: "夸克｜热搜榜 / UC｜热榜"
  current_route: rejected_duplicate_search_pool
  decision_status: recovered_from_pr
  selected_reason: "无常驻价值。"
  rejected_alternatives:
    - "夸克常驻"
    - "UC 常驻"
    - "同时加入多个同源搜索榜"
  route_reason: "两张榜的选题、顺序和标题高度相似，几乎是同一内容池的不同产品入口；与百度、头条和现有公共温度节点相比也没有足够增量。"
  reuse_rule: "同一内容池换产品入口，不算独立视角。以后遇到相似榜单，先看内容池是否真正不同。"
  evidence_locator: "docs/ops/tophub-comprehensive-hot-search-closure-20260704.md#四夸克热搜榜-与-uc热榜"
  next_action: "无；不新增。"

- source_id: "搜狗｜实时热点"
  current_route: rejected_unstable_methodology
  decision_status: recovered_from_pr
  selected_reason: "差异本身不足以构成价值。"
  rejected_alternatives:
    - "常驻公共温度"
  route_reason: "它与其他搜索榜重合度略低，但议题时间口径混杂，部分条目像编辑推荐而非真实实时搜索，且无法判断热度生成机制。缺少稳定口径时，不能承担长期观察节点。"
  reuse_rule: "榜单和主流来源不同不等于值得保留；必须能解释口径和稳定性。"
  evidence_locator: "docs/ops/tophub-comprehensive-hot-search-closure-20260704.md#五搜狗实时热点"
  next_action: "无；不新增。"

- source_id: "360搜索｜实时热点榜单"
  current_route: rejected_low_density_title_stream
  decision_status: recovered_from_pr
  selected_reason: "可显示搜索门户如何标题化事件，但不适合常驻。"
  rejected_alternatives:
    - "公共温度常驻"
  route_reason: "榜单把新闻进一步压缩成刺激性短词组，与微信热文、百度和头条的标题样本重复，信息密度更低。"
  reuse_rule: "标题刺激性更强不是独立信息价值；若只放大情绪词，不常驻。"
  evidence_locator: "docs/ops/tophub-comprehensive-hot-search-closure-20260704.md#六360搜索实时热点榜单"
  next_action: "无；不新增。"
```

---

## 4. 平台热搜为什么不叠加

```yaml
- source_id: "知乎｜热搜"
  current_route: rejected_below_zhihu_hotlist
  decision_status: recovered_from_pr
  selected_reason: "无常驻价值；仅按需比较知乎用户搜索与热榜推荐差异。"
  rejected_alternatives:
    - "知乎｜热榜"
  route_reason: "它只是知乎热榜事件的关键词缩写，信息量低于知乎热榜，也无法显示问题如何被组织和解释。继续保留知乎热榜，不添加知乎热搜。"
  reuse_rule: "同平台搜索榜若低于现有内容榜的信息量，不新增。"
  evidence_locator: "docs/ops/tophub-comprehensive-hot-search-closure-20260704.md#八知乎热搜; docs/ops/tophub-comprehensive-zhihu-closure-20260704.md#三知乎热搜"
  next_action: "无；已在知乎小标签闭环。"

- source_id: "哔哩哔哩｜热搜"
  current_route: on_demand_community_reaction
  decision_status: recovered_from_pr
  selected_reason: "能显示 B 站社区此刻在讨论什么，可按需观察特定事件的社区反应。"
  rejected_alternatives:
    - "进入影游音乐常驻"
    - "进入公共温度常驻"
  route_reason: "已有 `哔哩哔哩｜每周必看` 承担较低压力的内容发现；再加入热搜会把影游音乐拉回高频热点流，放入公共温度又与微博、抖音、知乎重复。"
  reuse_rule: "社区热搜如果只是高频反应入口，按事件临时打开；常驻只保留低压内容发现或明确独特公共温度。"
  evidence_locator: "docs/ops/tophub-comprehensive-hot-search-closure-20260704.md#十二哔哩哔哩热搜"
  next_action: "按需。"

- source_id: "腾讯视频｜热搜榜"
  current_route: on_demand_video_platform_search
  decision_status: recovered_from_pr
  selected_reason: "需要看腾讯视频平台内容时按需打开。"
  rejected_alternatives:
    - "影游音乐常驻"
  route_reason: "当前仅列剧集、综艺与动漫名称，且无法区分搜索、播放、营销和站内推荐影响。现有影游音乐分组已有豆瓣电影、IMDb、哔哩哔哩每周必看、游研社、机核等节点，腾讯视频热搜没有足够独特性。"
  reuse_rule: "视频平台站内热搜不等于作品发现；无法区分播放、搜索、营销和推荐时，按需。"
  evidence_locator: "docs/ops/tophub-comprehensive-hot-search-closure-20260704.md#十腾讯视频热搜榜"
  next_action: "按需。"

- source_id: "爱奇艺｜热搜榜"
  current_route: rejected_stale_or_invalid
  decision_status: recovered_from_pr
  selected_reason: "无。"
  rejected_alternatives:
    - "影游音乐常驻"
  route_reason: "当前榜单明显陈旧，包含多部旧剧旧季，不是可信的 2026 年实时热搜，可能是接口停更、历史缓存或错误源。"
  reuse_rule: "榜单时间口径明显陈旧或疑似缓存，不进入任何常驻分组。"
  evidence_locator: "docs/ops/tophub-comprehensive-hot-search-closure-20260704.md#十一爱奇艺热搜榜"
  next_action: "无；拒绝。"
```

---

## 5. 垂直 / 机构 / 失效热搜

```yaml
- source_id: "中国搜索｜热搜榜"
  current_route: on_demand_official_agenda_search
  decision_status: recovered_from_pr
  selected_reason: "可作为官方议程检索入口按需使用。"
  rejected_alternatives:
    - "公共温度常驻"
  route_reason: "它不是普通公众搜索温度，而是机构化、政策化的议程排序；现有新闻联播、半月谈、新华社按需页和官方数据源更直接承担这一角色。"
  reuse_rule: "官方议程榜若只是中间层，优先看原始官方源或更稳定的官方栏目。"
  evidence_locator: "docs/ops/tophub-comprehensive-hot-search-closure-20260704.md#十三中国搜索热搜榜"
  next_action: "按需。"

- source_id: "新浪汽车｜热搜榜"
  current_route: rejected_auto_marketing_stream
  decision_status: recovered_from_pr
  selected_reason: "无常驻价值。"
  rejected_alternatives:
    - "汽车行业常驻"
    - "公共温度常驻"
  route_reason: "榜单高度被品牌活动、明星代言、发布会话题、智驾营销、车型预售和企业传播口号占据，更像汽车营销与微博话题投放榜。具体购车、车型或政策问题应查官方参数、车主反馈和专业测试。"
  reuse_rule: "垂直行业热搜若主要反映营销投放，不进入结构分组；具体任务回官方参数和专业测试。"
  evidence_locator: "docs/ops/tophub-comprehensive-hot-search-closure-20260704.md#十四新浪汽车热搜榜"
  next_action: "无；不新增。"

- source_id: "头条热榜｜总榜"
  current_route: rejected_duplicate_total榜
  decision_status: recovered_from_pr
  selected_reason: "无常驻价值。"
  rejected_alternatives:
    - "因为叫总榜就加入"
  route_reason: "与微博、百度、夸克、UC 及现有公共温度节点高度重合，并会把同一赛事拆成大量条目；不因名为总榜就赋予更高代表性。"
  reuse_rule: "总榜名义不构成更高价值；仍需看是否提供独立结构。"
  evidence_locator: "docs/ops/tophub-comprehensive-hot-search-closure-20260704.md#十五头条热榜总榜"
  next_action: "无；不新增。"

- source_id: "头条搜索｜猜你想搜"
  current_route: rejected_not_stable_public榜
  decision_status: recovered_from_pr
  selected_reason: "无。"
  rejected_alternatives:
    - "公共温度常驻"
  route_reason: "内容混合天气、机票、新闻、电脑设置、宋词、红酒、武器性能等，更像搜索建议、常见查询和平台推荐的混合；名称本身可能包含推荐或个性化逻辑，不适合公共温度。"
  reuse_rule: "搜索建议 / 猜你想搜 / 个性化推荐不等于公开稳定榜单。"
  evidence_locator: "docs/ops/tophub-comprehensive-hot-search-closure-20260704.md#九头条搜索猜你想搜"
  next_action: "无；明确不新增。"

- source_id: "封面｜热搜"
  current_route: rejected_invalid_node
  decision_status: recovered_from_pr
  selected_reason: "无。"
  rejected_alternatives:
    - "任何常驻分组"
  route_reason: "当前页面只显示网站自律管理承诺书，节点已经失效或接口错误。"
  reuse_rule: "节点失效或接口错误，直接排除，不补编价值。"
  evidence_locator: "docs/ops/tophub-comprehensive-hot-search-closure-20260704.md#十六封面热搜"
  next_action: "无；排除。"
```

---

## 6. 滚动新闻小标签整体判断

来源证据：`docs/ops/tophub-comprehensive-rolling-news-closure-20260704.md`

`滚动新闻` 小标签共 31 个节点，三页已经全部查看。

最终结论：

- 不新增任何节点；
- 继续保留现有 `参考消息｜滚动新闻`；
- `新华网｜滚动新闻` 降为按需核验入口；
- `中新网｜滚动即时新闻` 降为按需入口；
- 其余节点排除或仅在具体领域、地区与任务出现时打开；
- 不把任何滚动新闻源放进 Folo。

```yaml
- source_id: "滚动新闻小标签整体"
  current_route: no_new_tophub_no_folo
  decision_status: recovered_from_pr
  selected_reason: "当前 TopHub 已经有足够的公共快讯与事件线索入口，滚动新闻继续作为发现与横向扫视能力存在，但本标签不再扩容。"
  rejected_alternatives:
    - "把新华网滚动新闻加入常驻"
    - "把中新网滚动即时新闻加入常驻"
    - "把滚动新闻源放进 Folo"
    - "加入地方滚动、体育滚动、门户科技滚动"
  route_reason: "滚动新闻属于 TopHub 的发现与横向扫视能力，不属于 Folo 的连续阅读关系。具体事件持续跟踪时，应使用精确关键词、对应官方源、必要时临时观察，而不是订阅整个滚动新闻流。"
  reuse_rule: "以后遇到滚动新闻源，除非它提供当前系统缺失的稳定角色，否则按需；不要用快讯流填 Folo。"
  evidence_locator: "docs/ops/tophub-comprehensive-rolling-news-closure-20260704.md#三页最终比较; #最终结论; #tophub与folo路由"
  next_action: "无；滚动新闻小标签已闭环。"
```

---

## 7. 已保留与按需滚动新闻入口

```yaml
- source_id: "参考消息｜滚动新闻"
  current_route: tophub_public_temperature
  decision_status: recovered_from_pr
  selected_reason: "承担外媒转述、国际时事与跨国议题观察，和国内门户热搜的角色不同。"
  rejected_alternatives:
    - "新华网｜滚动新闻作为新增常驻"
    - "中新网｜滚动即时新闻作为新增常驻"
  route_reason: "继续保留现有订阅；它提供国际视角与跨国议题线索，而不是单纯国内即时快讯。"
  reuse_rule: "已有一个能承担跨国议题观察的滚动源时，不再叠加更多国内泛滚动快讯。"
  evidence_locator: "docs/ops/tophub-comprehensive-rolling-news-closure-20260704.md#三页最终比较"
  next_action: "无；保留。"

- source_id: "新华网｜滚动新闻"
  current_route: on_demand_fact_checking_entry
  decision_status: recovered_from_pr
  selected_reason: "三页中最好的新增候选，具备通讯社式事实线索能力。"
  rejected_alternatives:
    - "常驻 TopHub"
    - "Folo 连续订阅"
  route_reason: "与参考消息、后续跟踪 Live、新闻联播、人民日报电子版、澎湃热榜等重叠较高；新增会增加大量日常滚动条目，却不会增加当前系统完全缺失的能力。"
  reuse_rule: "事实线索能力强但与现有公共快讯高度重叠时，降为按需核验入口。"
  evidence_locator: "docs/ops/tophub-comprehensive-rolling-news-closure-20260704.md#三页最终比较-新华网滚动新闻"
  next_action: "按需核验。"

- source_id: "中新网｜滚动即时新闻"
  current_route: on_demand_news_entry
  decision_status: recovered_from_pr
  selected_reason: "可作为即时新闻按需入口。"
  rejected_alternatives:
    - "常驻 TopHub"
    - "Folo 连续订阅"
  route_reason: "比新华网更杂，混合地方、体育、生活、时政与即时资讯；若不新增新华网，更没有理由新增中新网。"
  reuse_rule: "若更好的同类候选都不常驻，更杂的候选只能按需。"
  evidence_locator: "docs/ops/tophub-comprehensive-rolling-news-closure-20260704.md#三页最终比较-中新网滚动即时新闻"
  next_action: "按需。"
```

---

## 8. 第三页滚动新闻拒绝项

```yaml
- source_id: "腾讯科技｜滚动新闻"
  current_route: rejected_portal_tech_aggregate
  decision_status: recovered_from_pr
  selected_reason: "无常驻价值。"
  rejected_alternatives:
    - "科技雷达常驻"
    - "滚动新闻常驻"
  route_reason: "科技、商业、明星、机器人伴侣、苹果传闻、A股、体育营销和公司评论混在一起；大量使用‘消息称’‘曝光’等新闻化标题；与现有科技雷达高度重复。"
  reuse_rule: "门户科技滚动若只是聚合和标题化，不新增；科技问题回科技雷达或原始来源。"
  evidence_locator: "docs/ops/tophub-comprehensive-rolling-news-closure-20260704.md#腾讯科技滚动新闻"
  next_action: "无；不新增。"

- source_id: "京报网｜新闻滚动 / 北青网｜滚动新闻 / 北晚在线｜滚动"
  current_route: on_demand_beijing_local_sources
  decision_status: recovered_from_pr
  selected_reason: "在北京生活、文化活动、城市服务或本地问题阶段可按需查看。"
  rejected_alternatives:
    - "全国公共温度常驻"
  route_reason: "它们混合北京地方政务、天气、党建、文艺、展览、体育、社区宣传和全国新闻；既不是纯北京行动入口，也不是低噪声全国快讯。没有明确北京长期关系时，不进入常驻层。"
  reuse_rule: "地方滚动源只有在用户与该地形成长期关系或具体任务时才按需打开。"
  evidence_locator: "docs/ops/tophub-comprehensive-rolling-news-closure-20260704.md#京报网新闻滚动; #北青网滚动新闻; #北晚在线滚动"
  next_action: "北京相关任务时按需。"

- source_id: "四川在线｜滚动新闻"
  current_route: on_demand_regional_source
  decision_status: recovered_from_pr
  selected_reason: "可在四川地方文化、自然景观或区域新闻任务中按需使用。"
  rejected_alternatives:
    - "常驻公共温度"
  route_reason: "当前样本明显属于区域媒体，没有明确四川长期关系时按需。"
  reuse_rule: "区域媒体不因可更新就常驻；必须有持续地域关系或具体任务。"
  evidence_locator: "docs/ops/tophub-comprehensive-rolling-news-closure-20260704.md#四川在线滚动新闻"
  next_action: "按需。"

- source_id: "新浪体育｜滚动新闻"
  current_route: rejected_gambling_like_sports_stream
  decision_status: recovered_from_pr
  selected_reason: "无。"
  rejected_alternatives:
    - "影游音乐"
    - "公共温度"
  route_reason: "页面绝大部分是彩票预测、排列三、福彩 3D、竞彩情报和投注技巧，不是可靠体育新闻流，也不应进入任何公共分组。"
  reuse_rule: "涉及彩票预测、投注技巧、竞彩情报的体育流直接排除，不包装成体育新闻。"
  evidence_locator: "docs/ops/tophub-comprehensive-rolling-news-closure-20260704.md#新浪体育滚动新闻"
  next_action: "无；明确排除。"

- source_id: "直播吧｜中超滚动新闻"
  current_route: rejected_unstable_sports_label
  decision_status: recovered_from_pr
  selected_reason: "无常驻价值。"
  rejected_alternatives:
    - "影游音乐"
    - "公共温度"
  route_reason: "虽然标题是中超滚动，当前页面实际大量是世界杯、欧洲足球、球员数据和社交媒体转述，标签边界不稳定。若要长期看足球，应选择明确赛事或球队来源，而不是加入综合滚动层。"
  reuse_rule: "体育滚动若标签边界不稳，不常驻；长期看某项赛事时选择明确赛事或球队来源。"
  evidence_locator: "docs/ops/tophub-comprehensive-rolling-news-closure-20260704.md#直播吧中超滚动新闻"
  next_action: "无；不新增。"
```

---

## 9. 本补充 ledger 的复用规则

### 热搜类

1. 不叠加同类热搜榜。
2. 同一事件被拆成大量近义词时，说明榜单在放大重复，不是拓宽视野。
3. 搜索引擎热搜只在比较搜索词命名、事件拆分和标题形态时按需打开。
4. 视频平台站内热搜不等于作品发现。
5. 官方 / 垂直 / 汽车热搜优先按原始官方源、行业源或任务源核验。
6. 热搜不进 Folo，不建宽泛追踪器；具体且有限的事件、政策、产品、公司或人物变化才进追踪。

### 滚动新闻类

1. 滚动新闻属于 TopHub 的横向扫视能力，不属于 Folo 连续阅读关系。
2. 已有足够公共快讯入口时，不新增更多滚动快讯。
3. 高质量但重叠的滚动源降为按需核验入口。
4. 地方滚动源只有在具体地域任务或长期关系出现时使用。
5. 体育 / 彩票 / 竞彩 / 标签不稳的滚动源直接排除或按明确赛事来源重建。

---

## 10. 仍需继续追回

综合大类已经补入：

- closeout 三项动作；
- 知乎小标签；
- 热搜小标签；
- 滚动新闻小标签。

仍需继续追回：

- 微信 / 微博 / 公众号相关判断；
- 百度、腾讯、网易、搜狐、新浪、今日头条、ZAKER 等门户同族 closure；
- 半月谈、央视、人民网、新华社、南方周末等同族 closure；
- 健康、军事、房产、地方城市等 closure；
- `腾讯研究院` 与 `网易人间｜记事` 的更完整理由。

综合大类仍是 `partial_recovered_from_pr`，不能写成全量完成。
