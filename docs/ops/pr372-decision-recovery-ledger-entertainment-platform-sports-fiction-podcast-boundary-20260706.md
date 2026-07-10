# PR #372 判断链追回：娱乐平台、体育、网文与播客边界（2026-07-06）

## 0. 边界

本文件补充追回 PR #372 中娱乐目录剩余平台型、社区型、赛事型、网文型和播客型来源的判断链。

纳入文件：

- `docs/ops/tophub-entertainment-bilibili-closure-20260705.md`
- `docs/ops/tophub-entertainment-short-video-closure-20260705.md`
- `docs/ops/tophub-entertainment-acfun-closure-20260705.md`
- `docs/ops/tophub-entertainment-hupu-closure-20260705.md`
- `docs/ops/tophub-entertainment-sports-closure-20260705.md`
- `docs/ops/tophub-entertainment-football-closure-20260705.md`
- `docs/ops/tophub-entertainment-basketball-closure-20260705.md`
- `docs/ops/tophub-entertainment-fiction-closure-20260705.md`
- `docs/ops/tophub-entertainment-podcast-pages-01-04-audit-20260705.md`
- `docs/ops/tophub-entertainment-ximalaya-closure-20260705.md`

边界说明：

- B站、短视频、AcFun、虎扑、体育、足球、篮球、小说、喜马拉雅已按各自文件闭环。
- 播客文件只覆盖前 4 页、48/95 个节点，是阶段记录，不是完整闭环。
- 当前不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 平台榜总规则：保留少数沉淀入口，不订阅分区大流

```yaml
- source_id: "娱乐平台榜总规则"
  current_route: keep_low_noise_platform_entry_points_other_platform_slices_on_demand
  decision_status: recovered_from_pr
  selected_reason: "B站、抖音、AcFun、快手、喜马拉雅、虎扑等平台都能暴露真实用户行为、平台文化、社区争论和作品线索。"
  rejected_alternatives:
    - "把平台日榜、综合热门、热搜、分区榜全部常驻"
    - "把平台榜单加入 Folo"
    - "把平台榜当成作者关系"
    - "把每个分区榜当成独立来源能力"
  route_reason: "保留少量低压力入口：抖音总榜作为大众传播温度，B站每周必看作为沉淀后的视频候选池。其余平台分区榜、短周期榜、商品榜、文章榜和历史累计榜都按需或排除。"
  reuse_rule: "以后遇到平台榜，先判断它是总榜样本、低频精选、分区热榜、商品上新、历史目录、社区帖子还是具体作者/节目。只有总榜样本和低频精选可能常驻；作者/节目另进 Folo 复查。"
  evidence_locator: "bilibili closure / short video closure / acfun closure / ximalaya closure / hupu closure"
  next_action: "无平台榜新增。"
```

---

## 2. B站：为什么只保留每周必看

```yaml
- source_id: "哔哩哔哩｜每周必看"
  current_route: keep_tophub_movie_game_music_low_noise_video_candidate_pool
  decision_status: recovered_from_pr
  selected_reason: "它经过一周传播与平台筛选，更新压力低于日榜，同时覆盖创作、动画、音乐、游戏、知识、生活和公共话题；适合作为低压力的视频候选池，而不是待清空列表。"
  rejected_alternatives:
    - "哔哩哔哩｜全站日榜"
    - "哔哩哔哩｜综合热门"
    - "哔哩哔哩｜热搜"
    - "知识榜、科技数码榜、游戏榜、影视榜、音乐榜、动画榜等分区榜"
    - "B站专栏周榜、月榜、昨天榜、前天榜"
    - "会员购今日上新"
  route_reason: "继续保留在影游音乐。日榜和综合热门更新快、题材混杂、宣发和热点跟随较多；热搜与公共温度重复；分区榜分类漂移明显；专栏时间切片没有形成有意义的编辑差异；会员购是商品上新。"
  reuse_rule: "视频平台只保留低频沉淀入口；具体视频进入原页观看，必要时交给 BibiGPT 摘要。若反复关注同一 UP 主，再按作者关系复查。"
  evidence_locator: "docs/ops/tophub-entertainment-bilibili-closure-20260705.md#节点判断; #最终动作"
  next_action: "保持。"

- source_id: "B站平台榜进入 Folo 的边界"
  current_route: no_folo_for_bilibili_platform_rankings
  decision_status: recovered_from_pr
  selected_reason: "B站榜单是平台榜单、排行、热搜或商品更新入口，不是具体作者、节目、栏目团队或一手连续来源。"
  rejected_alternatives:
    - "把每周必看或分区榜加入 Folo"
    - "用整个平台榜替代 UP 主订阅"
  route_reason: "B站 24 个节点均不进 Folo。平台榜放进 Folo 会制造与 TopHub 重复的未读流，也不符合 Folo 用于少量长期关系源的定位。"
  reuse_rule: "平台榜不进 Folo；从榜单中反复发现同一具体 UP 主、视频播客或创作团队后，再检查主页、RSS、更新频率和长期价值。"
  evidence_locator: "docs/ops/tophub-entertainment-bilibili-closure-20260705.md#Folo"
  next_action: "无。"
```

---

## 3. 短视频：抖音总榜保留，其他短视频榜不新增

```yaml
- source_id: "抖音｜总榜"
  current_route: keep_tophub_public_temperature_short_video_mass_attention
  decision_status: recovered_from_pr
  selected_reason: "它用于观察抖音平台当日真正传播的叙事、人物、地方生活、商业内容、非遗、美食和大众情绪。"
  rejected_alternatives:
    - "抖音｜热点榜"
    - "抖音｜视频总榜"
    - "抖音｜体育榜 / 娱乐榜 / 游戏榜 / 时尚榜等细分榜"
  route_reason: "保留总榜。热点榜偏事件词条，不能替代实际视频内容结构；视频总榜出现春节、马年、烟花和旧挑战内容，时间口径失真；细分类榜边界污染严重。"
  reuse_rule: "短视频平台若常驻，只保留一个大众传播样本；细分类榜按事件或研究任务临时打开。"
  evidence_locator: "docs/ops/tophub-entertainment-short-video-closure-20260705.md#结论; #抖音"
  next_action: "保持。"

- source_id: "快手 / 美拍 / 好看视频 / 梨视频 / 360短视频等短视频榜"
  current_route: rejected_or_on_demand_short_video_platform_slices
  decision_status: recovered_from_pr
  selected_reason: "个别平台在下沉市场、短剧、社会榜、挑战榜、视频社区或老内容库方面有观察价值。"
  rejected_alternatives:
    - "快手实时热榜 / 社会榜 / 搜索飙升榜 / 文娱榜 / 有用榜 / 短剧榜常驻"
    - "梨视频 / 好看视频 / 美拍 / 360 短视频榜常驻"
    - "快手指数洞察报告常驻"
  route_reason: "不新增。多数节点与公共温度重复，或者由平台梗、模板、挑战、旧内容、短剧、营销、生活技巧、游戏攻略和跨品类流量变化组成。快手指数洞察报告作为研究特定人群、行业、下沉市场、银发和内容消费问题时的按需资料库。"
  reuse_rule: "短视频榜只作传播样本，不作事实源或作品质量源；平台研究报告按研究任务调用。"
  evidence_locator: "docs/ops/tophub-entertainment-short-video-closure-20260705.md#快手; #结论"
  next_action: "按需或排除。"
```

---

## 4. AcFun：社区差异存在，但榜单不稳定

```yaml
- source_id: "AcFun 15 个节点"
  current_route: on_demand_acg_long_tail_video_community_not_subscription
  decision_status: recovered_from_pr
  selected_reason: "AcFun 仍能提供 B 站以外的 ACG、东方、MMD、翻唱、老歌、日系偶像、复古游戏、主机史、日综搬运和长尾视频线索。"
  rejected_alternatives:
    - "AcFun 全站综合榜 / 三日榜 / 周榜常驻"
    - "AcFun 文章榜 / 鱼塘榜 / 生活榜 / 科技榜常驻"
    - "AcFun 游戏 / 动画 / 音乐 / 舞蹈 / 影视 / 番剧榜常驻"
    - "AcFun 平台榜加入 Folo"
  route_reason: "15 个节点全部不新增。全站榜、三日榜、周榜高度重复；文章、鱼塘、生活榜混入事故、情绪、低质政治评论、SEO、律师营销和零播放；分区榜样本量小、质量起伏大，常有搬运、录播和零散投稿。"
  reuse_rule: "A站更适合作为长尾社区气质和旧内容库按需入口。若反复关注某个具体创作者或栏目，再按 Folo 作者关系复查。"
  evidence_locator: "docs/ops/tophub-entertainment-acfun-closure-20260705.md#是否有值得为用户新增的订阅; #节点判断; #Folo"
  next_action: "按需。"
```

---

## 5. 虎扑：有样本价值，但不是低噪声入口

```yaml
- source_id: "虎扑 16 个节点"
  current_route: on_demand_male_skewed_chinese_forum_sample
  decision_status: recovered_from_pr
  selected_reason: "虎扑能提供男性用户占比较高的中文体育社区、球迷争论、装备消费、职场困境、汽车品牌战争和散户情绪样本。"
  rejected_alternatives:
    - "步行街 / 生活 / 影音娱乐热帖常驻"
    - "NBA / CBA / 国际足球 / 中国足球论坛常驻"
    - "游戏电竞、运动装备、职场、恋爱、股票、汽车热帖常驻"
    - "虎扑资讯 / 篮球新闻常驻"
    - "任何虎扑榜单进入 Folo"
  route_reason: "16 个榜单不能把这种样本整理成低噪声、稳定、可持续信号。步行街、生活、影音娱乐三榜近乎重复；体育论坛是情绪流而非稳定信息源；股票区投机表达过强；职场区被广告和 SEO 污染；恋爱区依赖性暗示和猎奇；汽车热帖品牌阵营冲突强。"
  reuse_rule: "虎扑只作为社区反应样本。事实回联盟、球队、赛事、官方公告、原始记者、监管、专业测评或财经事实源核验。"
  evidence_locator: "docs/ops/tophub-entertainment-hupu-closure-20260705.md#是否有值得新增; #高优先级按需入口; #明确不采用"
  next_action: "按需。"
```

---

## 6. 体育 / 足球 / 篮球：赛事周期触发，不建永久流

```yaml
- source_id: "体育大类 18 节点"
  current_route: no_persistent_sports_feed_event_triggered_sources
  decision_status: recovered_from_pr
  selected_reason: "体育信息在世界杯、奥运会、温网、WTT 等大赛期间密度上升；赛事结束后多数节点转向转会传闻、花边、社区争论、比分播报和流量标题。"
  rejected_alternatives:
    - "新增独立体育分组"
    - "体育、世界杯、足球、篮球、乒乓球等宽泛追踪器"
    - "新浪体育、搜狐体育、ZAKER 体育、微信体育、抖音体育等常驻"
  route_reason: "用户当前没有需要每天维护的固定联赛、球队、运动员或比赛日程工作流。更合适结构是重大赛事期间临时调用分析型来源，事实核验调用新华社 / 中新网，社区反应按需查看懂球帝、虎扑或中羽在线，赛事结束后关闭追踪。"
  reuse_rule: "体育来源按赛事周期、球队/运动员关系、装备任务或训练任务触发；不因短期赛事高峰留下永久订阅。"
  evidence_locator: "docs/ops/tophub-entertainment-sports-closure-20260705.md#为什么体育大类本轮不应新增常驻位; #最终结论"
  next_action: "无常驻。"

- source_id: "运动视界 Sports Vision"
  current_route: on_demand_sports_deep_analysis_first_choice
  decision_status: recovered_from_pr
  selected_reason: "它是体育页质量最高、增量最明确的节点，能分析世界杯比赛战术变化、晶片越位机制、球队背景、欧冠英超足总杯赛季总结。"
  rejected_alternatives:
    - "加入 TopHub 常驻"
    - "进入 Folo"
  route_reason: "定位为重大足球赛事期间第一优先级深度入口。当前样本几乎全部集中足球，价值高度依赖赛事周期；用户尚无持续足球阅读工作流，也不足以确认稳定作者关系和非赛事期结构。"
  reuse_rule: "赛事深度源先按重大赛事使用；连续一个赛事周期实际阅读后，再评估常驻或 Folo。"
  evidence_locator: "docs/ops/tophub-entertainment-sports-closure-20260705.md#最高优先级按需入口运动视界SportsVision"
  next_action: "按需。"

- source_id: "懂球帝｜深度"
  current_route: on_demand_football_deep_context_first_choice
  decision_status: recovered_from_pr
  selected_reason: "足球页 17 个节点中唯一持续提供明显解释增量的栏目，覆盖战术、历史、人物、国家、商业和地方文化。"
  rejected_alternatives:
    - "懂球帝今日头条 / 头条新闻 / 热门新闻 / 早报"
    - "懂球帝中超 / 德甲 / AC米兰 / 山东鲁能泰山等球队或联赛节点"
    - "虎扑国际足球 / 中国足球论坛常驻"
    - "FOX Sports Soccer 常驻"
  route_reason: "作为重大赛事、足球文化史、战术和人物故事的最高优先级按需入口，不加入 TopHub 或 Folo。实时与聚合栏目混入花边、篮球、电竞、球员家属和平台活动；球队/联赛节点边界不稳或与用户无持续关系；FOX Sports Soccer 适合世界杯期间英文赛程和北美视角。"
  reuse_rule: "足球源分事实、赛程、深度、社区反应、球队/联赛、梗图和花边。深度按赛事任务打开；论坛观点不能替代录像、官方数据、原始采访或处罚文件。"
  evidence_locator: "docs/ops/tophub-entertainment-football-closure-20260705.md#第一优先级按需入口懂球帝深度; #虎扑两个足球论坛; #懂球帝实时与聚合栏目"
  next_action: "按需。"

- source_id: "篮球 9 节点"
  current_route: on_demand_basketball_event_and_community_reaction
  decision_status: recovered_from_pr
  selected_reason: "NBA / CBA 论坛、篮球新闻和直播吧 NBA 在重大赛事、交易期、选秀、中国男篮、CBA 人事、青训和社区反应中有线索价值。"
  rejected_alternatives:
    - "NBA / CBA / 中国男篮 / 湖人 / 詹姆斯等宽泛追踪器"
    - "虎扑 NBA / CBA 论坛常驻"
    - "百度 NBA 焦点新闻常驻"
    - "直播吧 NBA 篮球新闻常驻"
    - "FOX Sports NBA 常驻"
  route_reason: "9 个节点没有一个同时满足内容边界清晰、信息密度稳定、相比现有综合新闻和官方信息具有持续增量。论坛热帖大量是流言、交易猜测、球迷立场；百度聚合来源不透明；直播吧短讯二次转述多；FOX Sports NBA 当前样本明显包含赔率和博彩网站推荐。"
  reuse_rule: "篮球按季后赛、总决赛、选秀、重大交易、中国男篮大赛和 CBA 人事变化触发；事实回联盟、球队、FIBA、赛事官网、正式采访和可靠记者原文。"
  evidence_locator: "docs/ops/tophub-entertainment-basketball-closure-20260705.md#总体判断; #虎扑社区NBA论坛热帖; #FOXSportsNBA"
  next_action: "按需。"
```

---

## 7. 网文 / 小说：市场榜按需，作者关系进 Folo 复查

```yaml
- source_id: "小说 / 网文平台榜"
  current_route: on_demand_web_fiction_market_research_not_daily_reading_feed
  decision_status: recovered_from_pr
  selected_reason: "起点、纵横、七猫、推书君、百度小说榜能回答网文市场、类型、平台运营、题材模板、付费结构、完本经典和搜索热度问题。"
  rejected_alternatives:
    - "推书君完本榜 / 畅销榜 / 新书榜常驻"
    - "七猫 14 个男频 / 女频 / 新书 / 完结 / 大热 / 收藏 / 更新 / 原创风云榜常驻"
    - "百度小说榜常驻"
    - "小说关键词追踪器"
  route_reason: "不新增任何 TopHub 节点。完本榜更像网文经典目录；畅销榜和新书榜大量 0 分或指标不清；七猫榜单同一批作品跨榜重复，标题模板化，更新/收藏/原创飞跃衡量站内运营；百度小说榜是搜索热度，不能判断值得读、是否新或趋势来源。"
  reuse_rule: "网文榜用于市场研究和选题观察，不用于制造日常待读库存。想找书时用完本、新书和平台指标按需交叉看。"
  evidence_locator: "docs/ops/tophub-entertainment-fiction-closure-20260705.md#TopHub决策; #推书君三个榜不新增; #七猫十四个榜不新增; #百度小说榜不新增"
  next_action: "按需。"

- source_id: "木遥的窗子"
  current_route: folo_review_candidate_slow_read_author_not_executed
  decision_status: recovered_from_pr
  selected_reason: "它不是小说榜，而是跨数学、系统和社会悖论、科技法律与平台史、音乐演出、现代性、读书和个人经验的稳定个人作者视角。"
  rejected_alternatives:
    - "放入 TopHub"
    - "立即加入 Folo"
    - "与普通小说榜混同"
  route_reason: "进入 2026-07-17 Folo 复查第二优先候选，仅次于基本読書。价值来自同一作者如何跨领域建立联系，不来自标题排名；复查时确认近一年更新频率、官方 RSS 和实际文章质量。"
  reuse_rule: "作者型来源与平台榜分开。若价值来自持续作者视角，应进 Folo 复查；若只是平台排行，按需。"
  evidence_locator: "docs/ops/tophub-entertainment-fiction-closure-20260705.md#Folo决策; #木遥的窗子"
  next_action: "2026-07-17 Folo 复查。"

- source_id: "Velas电波站"
  current_route: folo_observation_low_priority_personal_game_writing
  decision_status: recovered_from_pr
  selected_reason: "可能提供游戏设计、个人写作和轻量周刊的混合漫游价值。"
  rejected_alternatives:
    - "进入 TopHub"
    - "列入正式 Folo 新增顺位"
  route_reason: "当前样本还不能证明 Weekly 是否仍持续、游戏设计文章是否是稳定主线、手记与碎片是否对用户具有持续打开价值，因此只作为低优先级观察。"
  reuse_rule: "个人漫游源需要稳定主线和更新证据；样本弱时观察，不急于订阅。"
  evidence_locator: "docs/ops/tophub-entertainment-fiction-closure-20260705.md#低优先级观察Velas电波站"
  next_action: "观察。"
```

---

## 8. 播客 / 喜马拉雅：节目关系优先，平台榜只找名字

```yaml
- source_id: "播客前四页 48/95 节点"
  current_route: partial_audit_folo_review_pool_not_closure
  decision_status: recovered_from_pr_partial
  selected_reason: "前 48 个节点显示四类对象：具体持续更新节目、平台聚合榜、停更或历史档案节目、英文 AI / 数据节目。"
  rejected_alternatives:
    - "写成播客 95 节点已闭环"
    - "在后四页未收到前决定 TopHub 新增和排位"
    - "把平台榜单等同于节目源"
  route_reason: "这是阶段记录，不是完整闭环。播客的核心价值通常不是榜单热度，而是与节目、主播和编辑判断建立长期关系；具体节目可能进入 Folo，平台榜按需发现。Folo 在 2026-07-17 前仍保持 27 项不变。"
  reuse_rule: "播客必须先确认节目当前更新、近 10 期质量、官方 RSS、实际收听意愿和与现有来源的重复；未完整审核前不得宣称闭环。"
  evidence_locator: "docs/ops/tophub-entertainment-podcast-pages-01-04-audit-20260705.md#状态说明; #阶段性总判断"
  next_action: "等待后四页或 2026-07-17 复查。"

- source_id: "播客高优先级候选：声东击西 / 捕蛇者说 / 博物志 / 所建所闻 / 字谈字畅 / 忽左忽右"
  current_route: folo_review_candidates_audio_relationships_not_executed
  decision_status: recovered_from_pr_partial
  selected_reason: "这些节目分别补国际结构解释、开发者实践、博物馆与地方文化、建筑城市公共空间、字体排印视觉文化、历史长时段解释。"
  rejected_alternatives:
    - "全部立即加入 Folo"
    - "按播客热榜排名加入"
    - "把高价值长节目都留下造成收听债务"
  route_reason: "均为复查候选，不执行。声东击西和忽左忽右价值高但负担重；捕蛇者说与当前 Agent / 开发实践强相关；博物志与所建所闻需比较城市旅行重叠；字谈字畅小众但系统缺口明确。"
  reuse_rule: "播客候选必须比较主题角色和收听负担；能讲清楚角色，不代表一定能持续听。"
  evidence_locator: "docs/ops/tophub-entertainment-podcast-pages-01-04-audit-20260705.md#第一层高优先级Folo复查候选"
  next_action: "复查。"

- source_id: "喜马拉雅 12 个榜单"
  current_route: no_tophub_no_folo_platform_audio_rankings
  decision_status: recovered_from_pr
  selected_reason: "喜马拉雅榜单能暴露具体节目名称、有声书库存、相声评书、影视节目、天文和机构音频候选。"
  rejected_alternatives:
    - "热门免费榜 / 热门订阅榜 / 热门好评榜常驻"
    - "总榜飙升榜 / 新品飙升榜常驻"
    - "音乐免费榜 / 娱乐免费榜 / 有声书免费榜 / 付费榜常驻"
    - "喜马拉雅榜单加入 Folo"
  route_reason: "不新增。榜单更新时间不等于内容当前性，热门免费和订阅榜被历史累计播放支配，总榜和新品飙升高度重复，头条免费榜不是可靠新闻入口，音乐免费榜主要是睡眠、白噪音、DJ合集和翻唱，有声书/相声评书/付费/娱乐榜更适合按具体作品选择。影视免费榜只作为找节目名字的目录使用。"
  reuse_rule: "音频平台榜只找节目名，不订阅榜单。节目级候选要单独核验当前更新、RSS、近 10 期质量和与现有来源重叠。"
  evidence_locator: "docs/ops/tophub-entertainment-ximalaya-closure-20260705.md#三个问题; #Folo; #高优先级按需入口; #明确不采用"
  next_action: "按节目复查。"
```

---

## 9. 明确排除信号

```yaml
- source_id: "娱乐平台 / 体育 / 网文 / 播客排除信号"
  current_route: reusable_rejection_rules
  decision_status: recovered_from_pr
  selected_reason: "用于后续复用。"
  rejected_alternatives:
    - "用热度、播放量、累计订阅、日榜、短周期飙升、赔率、社区回复、分区名、平台分类来证明常驻价值"
  route_reason: "明确排除：平台日榜和综合热门与低频精选重复；分区榜分类漂移；商品上新不是文化来源；论坛帖子不是作者关系；体育博彩和赔率污染；虎扑股票/恋爱/职场等高噪声社区不进系统；网文站内运营指标不等于阅读价值；喜马拉雅累计播放榜不等于当前节目质量；播客不完整审核不得写成完成。"
  reuse_rule: "平台、赛事、社区和音频来源都要把‘热’与‘值得长期跟随’分开。"
  evidence_locator: "all included entertainment platform/sports/fiction/podcast files"
  next_action: "作为复用规则。"
```

---

## 10. 本补充 ledger 的复用规则

1. 平台榜只保留少数低频沉淀入口或总榜样本，不订阅分区切片。
2. B站保留每周必看；抖音保留总榜；AcFun 和快手等按需。
3. 平台榜不进 Folo；具体 UP 主、创作者、节目或栏目团队另行复查。
4. 虎扑是社区反应样本，不是事实源、财经源、关系源或低噪声生活源。
5. 体育、足球、篮球按赛事周期、球队关系、装备任务、训练任务触发；不建永久体育流。
6. 体育事实回官方、赛事、联盟、球队、正式采访、可靠记者和新华社 / 中新网等事实源。
7. 网文平台榜用于市场研究、题材观察和按需找书，不制造日常待读库存。
8. 作者型读书 / 小说 / 游戏漫游源进入 Folo 复查，不与平台榜混同。
9. 播客看节目关系、RSS、近 10 期质量和实际收听，不看平台榜名次。
10. 播客前四页只是 partial，不得写成 95 节点完成。

---

## 11. 仍需继续追回

娱乐平台、体育、网文与播客边界已补成第二份娱乐 ledger。

继续待办：

- 娱乐剩余专项如果还要细追：IMDb、猫眼、腾讯视频、爱奇艺、QQ 音乐、网易云音乐、起点、纵横、七猫、17173、好游快爆等可再分批补；
- 等用户本地提交 master checklist 与 coverage audit 后，再统一登记本文件和此前新增 ledger；
- 旧 ChatGPT 对话追索所有 partial 条目。
