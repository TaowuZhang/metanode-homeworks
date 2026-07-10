# PR #372 判断链追回：娱乐最终动作与分组收口（2026-07-06）

## 0. 边界

本文件追回 PR #372 中 `TopHub > 娱乐` 大类最终真实动作的判断链。

纳入文件：

- `docs/ops/tophub-entertainment-popular-page-01-20260705.md`
- `docs/ops/tophub-entertainment-pending-actions-20260705.md`
- `docs/ops/tophub-entertainment-final-reconciliation-20260705.md`
- `docs/ops/tophub-entertainment-fashion-closure-20260705.md`
- `docs/ops/tophub-entertainment-douban-closure-20260705.md`
- `docs/ops/tophub-entertainment-reading-closure-20260705.md`
- `docs/ops/tophub-entertainment-music-closure-20260705.md`
- `docs/ops/tophub-entertainment-film-tv-closure-20260705.md`
- `docs/ops/tophub-entertainment-entertainment-tag-closure-20260705.md`
- `docs/ops/tophub-entertainment-games-closure-20260705.md`
- `docs/ops/tophub-entertainment-photography-closure-20260705.md`

当前不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 娱乐总规则：不是追星和热搜，而是五类文化入口

```yaml
- source_id: "TopHub > 娱乐最终闭环整体"
  current_route: platform_verified_entertainment_restructured
  decision_status: recovered_from_pr_and_platform_verified
  selected_reason: "娱乐大类最终保留五类功能：作品发现、评论与解释、创作与产业、少量传播温度、视觉漫游与摄影叙事。"
  rejected_alternatives:
    - "常驻明星私生活"
    - "粉圈热搜"
    - "平台重复细分榜"
    - "历史累计榜"
    - "高奢商品连续流"
    - "来源边界不清的节点"
  route_reason: "最终新增 10 个、取消 3 个，TopHub 由 73 增至 80；Folo 保持 27。娱乐不再被理解成泛娱乐热度，而是被拆成数据与结构、慢读与思想、影游音乐、视觉与自然几个不同功能层。"
  reuse_rule: "以后遇到娱乐来源，先判断它补的是作品、评论、产业、传播温度、视觉，还是粉圈/平台热榜/消费广告。只有前五类可能进入系统；后几类按需或排除。"
  evidence_locator: "docs/ops/tophub-entertainment-final-reconciliation-20260705.md#结构判断; #最终动作; docs/ops/tophub-entertainment-pending-actions-20260705.md"
  next_action: "后续只有内容质量变化时重开单个来源复核。"
```

---

## 2. 数据与结构：华丽志

```yaml
- source_id: "华丽志"
  current_route: tophub_data_and_structure_consumer_industry
  decision_status: recovered_from_pr_and_platform_verified
  selected_reason: "它承担时尚产业的结构性变化：品牌、资本、制造、供应链、知识产权、劳动规则和中国消费市场如何变化，而不是新品展示、明星造型、购物灵感或品牌传播。"
  rejected_alternatives:
    - "每日腕表杂志"
    - "每日珠宝杂志"
    - "iBag 包包"
    - "微信时尚 24h 热文榜"
    - "Vogue Hong Kong"
    - "新浪 / 抖音时尚热榜"
    - "A Day Magazine"
    - "GQ 男士网"
    - "AnOther"
    - "HEAVEN RAVEN"
    - "取消 Yoho!潮流志"
  route_reason: "放入数据与结构第 12、全部第 41，位于 Counterpoint Research 之后，作为消费产业垂直观察。Yoho! 潮流志承担青年潮流、产品与生活方式表层观察，华丽志承担产业公司、制造体系、资本动作和市场策略，二者不直接替代。"
  reuse_rule: "时尚源如果只提供新品、明星造型、高奢商品或购物灵感，不常驻；能看到品牌背后的资本、制造、供应链、规则和市场变化，才可能进入数据与结构。"
  evidence_locator: "docs/ops/tophub-entertainment-fashion-closure-20260705.md#新增华丽志; #TopHub新增"
  next_action: "已新增。"
```

---

## 3. 慢读与思想：豆瓣影评、书格、取消非虚构旧榜

```yaml
- source_id: "豆瓣电影｜最受欢迎的影评"
  current_route: tophub_slow_read_and_thought_chinese_film_discussion
  decision_status: recovered_from_pr_and_platform_verified
  selected_reason: "它显示中文观众正在怎样解释、争论和重新理解具体作品。样本包含隐喻分析、女性主义解读、叙事批评、地域文化反应、纪录片与剧集评论，也有普通长评和粉丝问题。"
  rejected_alternatives:
    - "只保留豆瓣电影一周口碑榜"
    - "把影评放入影游音乐"
    - "用 Cinephilia 迷影完全替代中文观众讨论面"
  route_reason: "进入慢读与思想，而不是影游音乐。一周口碑榜回答‘看什么’，最受欢迎的影评回答‘人们怎样理解和争论它’；Cinephilia 提供更稳定的编辑与专业电影文章，豆瓣影评补当下中文电影讨论面。"
  reuse_rule: "评论源不是作品发现榜。若价值来自解释、争论、社会语境和观众理解，应放慢读与思想；若只是热度或片单，放影游音乐或按需。"
  evidence_locator: "docs/ops/tophub-entertainment-douban-closure-20260705.md#建议新增一; docs/ops/tophub-entertainment-film-tv-closure-20260705.md#豆瓣电影最受欢迎的影评"
  next_action: "已新增。"

- source_id: "书格｜每日好书"
  current_route: tophub_slow_read_and_thought_original_cultural_materials
  decision_status: recovered_from_pr_and_platform_verified
  selected_reason: "它直接提供古籍、画册、书法、谱录和历史图像档案发现；标题目录本身已足以决定是否打开，不是畅销榜、热搜榜、推荐值榜或现代出版营销入口。"
  rejected_alternatives:
    - "微信读书总榜 / 飙升榜 / 热搜榜 / 神作榜"
    - "当当畅销图书榜"
    - "起点 / 纵横 / 七猫 / 推书君等网文榜常驻"
    - "书伴"
  route_reason: "进入慢读与思想，位于豆瓣电影影评之后、器物于我之前。它补的是原始文化材料与数字典藏，不占 Folo 长期关系位。微信读书新书榜继续保留，作为中文数字出版的新书与新刊发现。"
  reuse_rule: "阅读源要区分出版发现、数字阅读新入库、书评、古籍档案、网文市场和销售榜。原始文化材料可以进慢读；销售榜和站内运营指标按需。"
  evidence_locator: "docs/ops/tophub-entertainment-reading-closure-20260705.md#新增书格每日好书; #TopHub决策"
  next_action: "已新增。"

- source_id: "豆瓣｜热门图书-非虚构类"
  current_route: retired_from_tophub_slow_read_static_old_list
  decision_status: recovered_from_pr_and_platform_verified
  selected_reason: "原角色是非虚构阅读发现。"
  rejected_alternatives:
    - "继续保留静态旧榜"
    - "用它继续承担当前非虚构发现"
  route_reason: "页面中的榜单仍由 2021 年出版的《也许你该找个人聊聊》《置身事内》《秦制两千年》《中世纪之美》等长期占据，已经从公共发现入口退化为静态旧书单。微信读书新书榜、豆瓣新书速递、豆瓣最受欢迎的书评和豆瓣热门图书-虚构类已经能承担新书、评论和虚构阅读发现。"
  reuse_rule: "阅读榜如果长期停滞，就不能继续承担当前发现；经典旧书可按需查，不占常驻位置。"
  evidence_locator: "docs/ops/tophub-entertainment-douban-closure-20260705.md#建议取消"
  next_action: "已取消。"
```

---

## 4. 音乐：三条互补入口，取消短周期飙升

```yaml
- source_id: "豆瓣｜华语新碟榜"
  current_route: tophub_movie_game_music_chinese_new_album_discovery
  decision_status: recovered_from_pr_and_platform_verified
  selected_reason: "它提供华语新发行专辑及更偏独立、非主流的唱片发现，能看到 Apple Music 商店榜和网易云短周期单曲传播容易漏掉的华语专辑与独立发行。"
  rejected_alternatives:
    - "豆瓣热门单曲榜"
    - "豆瓣音乐人最热单曲榜"
    - "QQ音乐或网易云大量地区 / 语种 / 类型榜"
  route_reason: "进入影游音乐第 2，位于 Apple Music 专辑排行之后、网易云原创榜之前。Apple Music 承担全球商店与流媒体专辑热度，豆瓣华语新碟承担华语新专辑，网易云原创榜承担中文原创单曲与平台内原创生态。"
  reuse_rule: "音乐常驻只留互补入口：全球专辑热度、华语新专辑、中文原创单曲。单曲、地区、语种、场景、活动榜按需。"
  evidence_locator: "docs/ops/tophub-entertainment-douban-closure-20260705.md#建议新增二; docs/ops/tophub-entertainment-music-closure-20260705.md#确认既有动作"
  next_action: "已新增。"

- source_id: "网易云音乐｜原创榜"
  current_route: tophub_movie_game_music_chinese_original_music
  decision_status: recovered_from_pr_and_platform_verified
  selected_reason: "它比飙升榜更接近中文原创音乐生态，而不是短周期传播波动。"
  rejected_alternatives:
    - "网易云音乐｜飙升榜"
    - "网易云地区 / 类型 / 场景 / 会员 / 车友 / 活动榜"
    - "QQ音乐平台切片榜"
  route_reason: "与网易云飙升榜同位置替换，影游音乐第 3。音乐标签最终只需要 Apple Music 专辑热度 + 豆瓣华语新碟 + 网易云原创单曲三条互补入口，不再增加第四个综合榜。"
  reuse_rule: "短周期飙升榜优先退出；原创生态榜如果能补作品来源结构，可以替代纯传播热度。"
  evidence_locator: "docs/ops/tophub-entertainment-music-closure-20260705.md#确认既有动作; #最终结论"
  next_action: "已新增。"

- source_id: "网易云音乐｜飙升榜"
  current_route: retired_from_tophub_short_cycle_music_spike
  decision_status: recovered_from_pr_and_platform_verified
  selected_reason: "原角色是中文平台歌曲短周期传播。"
  rejected_alternatives:
    - "继续保留飙升榜"
    - "再叠加更多网易云榜单"
  route_reason: "短周期飙升容易受平台传播、热梗和活动影响；在 Apple Music 专辑排行和豆瓣华语新碟已经存在后，更需要原创生态入口而不是传播波动入口。"
  reuse_rule: "音乐榜单看角色互补，不看榜单数量。传播波动榜只按需，不长期常驻。"
  evidence_locator: "docs/ops/tophub-entertainment-music-closure-20260705.md#确认既有动作; #最终结论"
  next_action: "已取消。"
```

---

## 5. 影视与娱乐：剧集、新京报、喷嚏乐影、取消开眼

```yaml
- source_id: "豆瓣｜全球口碑剧集榜"
  current_route: tophub_movie_game_music_current_quality_series_discovery
  decision_status: recovered_from_pr_and_platform_verified
  selected_reason: "当前影游音乐组有电影、音乐、视频、游戏和纪录片入口，但没有稳定承担‘当前有哪些剧集值得看’的节点。全球口碑剧集榜只有 10 项、跨国家跨平台、以口碑而不是播放量筛选，能覆盖华语、美剧、日韩剧、动画及其他地区剧集。"
  rejected_alternatives:
    - "豆瓣电影热门剧集排行榜"
    - "IMDb热门剧集榜"
    - "豆瓣华语口碑剧集榜"
    - "国产 / 美剧 / 日剧 / 韩剧等多个近期热门榜"
    - "腾讯视频、爱奇艺、电视猫、百度、360 的剧集榜"
  route_reason: "进入影游音乐第 6，位于 IMDb 热门电影榜之后、新京报娱乐之前。它承担当期优质剧集入口，不是追剧清单，不自动形成待看债务。"
  reuse_rule: "剧集源优先跨平台、低负担、当期口碑入口；单平台播放量、100 项大榜、地域拆分榜按需。"
  evidence_locator: "docs/ops/tophub-entertainment-film-tv-closure-20260705.md#新增豆瓣全球口碑剧集榜"
  next_action: "已新增。"

- source_id: "新京报｜娱乐"
  current_route: tophub_movie_game_music_chinese_culture_entertainment_industry_news
  decision_status: recovered_from_pr_and_platform_verified
  selected_reason: "它不以明星私生活和热搜搬运为核心，而覆盖新剧开播、创作者访谈、微短剧产业、游戏文化与平台治理、舞台、戏曲、诗歌、展览和电影制作。"
  rejected_alternatives:
    - "新浪娱乐点击量排行"
    - "微博文娱榜"
    - "抖音娱乐榜"
    - "腾讯新闻娱乐榜"
    - "头条娱乐榜"
    - "ZAKER娱乐"
    - "韩星网 / 韓娛最前線"
    - "人民网娱乐频道"
    - "文汇报娱乐常驻"
  route_reason: "进入影游音乐第 7，位于全球口碑剧集榜之后、B站每周必看之前。它补的是中国影视、演出、游戏与文化产业的当期编辑新闻；风险是仍有活动稿、发布稿和机构信息，所以只作为编辑入口，不继续加入其他综合娱乐新闻榜。"
  reuse_rule: "娱乐新闻常驻必须偏作品、创作者和产业，而不是恋情、外貌、粉圈争执、剧宣和明星私生活。"
  evidence_locator: "docs/ops/tophub-entertainment-entertainment-tag-closure-20260705.md#新增新京报娱乐; #不新增的节点"
  next_action: "已新增。"

- source_id: "喷嚏网｜乐影"
  current_route: tophub_movie_game_music_documentary_discovery_replacement_for_kaiyan
  decision_status: recovered_from_pr_and_platform_verified
  selected_reason: "它提供纪录片发现功能，覆盖政治史与战争、人物传记、科技、AI、核电与灾难、艺术史、中国社会与职业、国际关系。"
  rejected_alternatives:
    - "开眼视频｜日报"
    - "喷嚏网｜图卦"
    - "普通影视热榜"
  route_reason: "进入影游音乐第 12，替换开眼视频日报。它只承担片名发现，不承担事实核验、版权判断或正式观看来源保证；喷嚏网图卦的社会情绪流不进入常驻系统。"
  reuse_rule: "纪录片发现源可以常驻，但必须限制为作品发现；事实核验回原始材料、正式平台和可靠资料。"
  evidence_locator: "docs/ops/tophub-entertainment-film-tv-closure-20260705.md#喷嚏网乐影; docs/ops/tophub-entertainment-entertainment-tag-closure-20260705.md#喷嚏网图卦"
  next_action: "已新增。"

- source_id: "开眼视频｜日报"
  current_route: retired_from_tophub_low_density_visual_video_feed
  decision_status: recovered_from_pr_and_platform_verified
  selected_reason: "原角色是视频日报入口。"
  rejected_alternatives:
    - "继续保留开眼视频日报"
    - "用单页样本直接删除而不比较"
  route_reason: "娱乐第 1 页显示该节点已订阅，但样本只出现一条森海塞尔相关内容，信息密度较低；后续影视 / 视觉比较确认喷嚏网乐影能提供更明确的纪录片发现功能，因此以乐影替换。"
  reuse_rule: "视觉视频流如果不能稳定提供作品发现或策展密度，就应被更清晰的作品入口替代。"
  evidence_locator: "docs/ops/tophub-entertainment-popular-page-01-20260705.md#开眼视频日报; docs/ops/tophub-entertainment-film-tv-closure-20260705.md#已经确认的两个影视待办"
  next_action: "已取消。"
```

---

## 6. 游戏：INDIENOVA 与现有游研社 / 机核分工

```yaml
- source_id: "INDIENOVA｜文章"
  current_route: tophub_movie_game_music_indie_game_design_and_creation
  decision_status: recovered_from_pr_and_platform_verified
  selected_reason: "它补独立游戏、游戏设计、叙事方法和开发过程。样本包括 itch 一周游戏汇、每周 Steam 值得关注的游戏、独立游戏通讯、海外独立游戏见闻、多线性叙事美学、Bret Victor 可探索解释、游戏文化经典译丛、复古游戏、独立游戏节和创作现场。"
  rejected_alternatives:
    - "INDIENOVA｜开发"
    - "触乐"
    - "GameRes 游资网"
    - "GameLook / 手游那点事 / 游戏陀螺"
    - "3DM / 篝火营地 / A9VG / 4Gamers / 巴哈姆特 / 腾讯游戏早报 / RPS / Ars Gaming / ZAKER 游戏等高频新闻流"
  route_reason: "进入影游音乐，位于游研社之后、机核之前。游研社偏游戏与玩家文化、行业事件、人物和作品故事；机核偏综合、播客、采访、活动和泛文化；INDIENOVA 偏独立游戏与设计文章。它不是第三个游戏新闻站，而是把游戏部分从知道发生什么扩展到理解作品怎样做出来。"
  reuse_rule: "游戏源按作品文化、综合节目、独立游戏与设计、行业商业、商店折扣、攻略社区分层；常驻只留互补层，新闻流和商店榜按需。"
  evidence_locator: "docs/ops/tophub-entertainment-games-closure-20260705.md#新增INDIENOVA文章; #现有两项为什么继续保留; #游戏新闻与行业媒体判断"
  next_action: "已新增。"

- source_id: "Yuko's Blog"
  current_route: folo_review_candidate_game_design_author_not_executed
  decision_status: recovered_from_pr
  selected_reason: "游戏页新增为 Folo 第三优先候选，属于作者关系而不是 TopHub 热榜。"
  rejected_alternatives:
    - "立即加入 Folo"
    - "进入 TopHub"
  route_reason: "2026-07-17 前不执行，保持 Folo 27。作者型游戏 / 设计来源需要复查更新频率、全文可读性、主题稳定性和与 INDIENOVA、游研社、机核的互补。"
  reuse_rule: "游戏作者源进入 Folo 前，要证明长期关系成立；不能因一页命中就立即订阅。"
  evidence_locator: "docs/ops/tophub-entertainment-games-closure-20260705.md#Folo复查候选"
  next_action: "2026-07-17 Folo 复查。"
```

---

## 7. 视觉与自然：CNU 新增，胶片保留

```yaml
- source_id: "CNU视觉联盟｜每日精选"
  current_route: tophub_visual_and_nature_low_frequency_curated_visual_works
  decision_status: recovered_from_pr_and_platform_verified
  selected_reason: "它是窄的低频人工精选入口，提供作品、构图、色彩和人物视觉入口；页面样本少、信息量克制，比 500px 热门、站酷商业摄影和 CNU 热门/最新/推荐大流更低噪音。"
  rejected_alternatives:
    - "CNU 热门"
    - "CNU 最新"
    - "CNU 推荐"
    - "500px 摄影社区热门作品"
    - "站酷摄影编辑精选"
    - "Bing 壁纸"
    - "Photography Photo of the Day"
  route_reason: "进入视觉与自然第 3，位于中国国家地理之后、胶片的味道之前。只保留这一张 CNU 榜，不同时订阅热门、最新和推荐。"
  reuse_rule: "视觉源优先低频人工策展，不把平台大流、壁纸、商业摄影、器材新闻拉进常驻。"
  evidence_locator: "docs/ops/tophub-entertainment-photography-closure-20260705.md#新增并保留CNU视觉联盟每日精选; #TopHub最终顺序"
  next_action: "已新增。"

- source_id: "胶片的味道"
  current_route: keep_tophub_visual_and_nature_photography_narrative
  decision_status: recovered_from_pr_and_platform_verified
  selected_reason: "它承担摄影叙事、城市生活、胶片器材和画幅方法，连接照片与生活经验。"
  rejected_alternatives:
    - "因 CNU 新增而取消胶片的味道"
    - "仅根据标题文学化判断它已经漂移"
  route_reason: "复核修正后保留。此前仅根据标题列表判断其从摄影漂移到情绪散文，证据不足。文学化标题可能对应照片组、摄影随笔和视觉叙事，且列表仍有宾得 67、Mamiya RB67、135 画幅入门、快门、城市观察等摄影主题。CNU 承担低频人工精选，胶片承担摄影叙事与器材方法，两者不重复。"
  reuse_rule: "不能仅凭标题风格取消视觉来源；摄影源要看实际图文、作品结构和器材/方法内容。"
  evidence_locator: "docs/ops/tophub-entertainment-photography-closure-20260705.md#复核修正为什么保留胶片的味道; #保留的7个节点及角色"
  next_action: "已保留。"

- source_id: "Huiliu"
  current_route: folo_review_candidate_personal_photography_author_not_executed
  decision_status: recovered_from_pr
  selected_reason: "持续以个人摄影集为核心，主题包括阴影、家猫、乡村、城市、树、人、鸟与家庭记忆，更适合建立作者关系和长期跟随。"
  rejected_alternatives:
    - "立即加入 Folo"
    - "放入 TopHub 视觉热榜"
  route_reason: "进入 2026-07-17 Folo 复查候选，不在复查日前执行。复查时确认更新频率、实际图文质量、订阅源完整性，以及与 TopHub 视觉组的互补程度。"
  reuse_rule: "个人摄影作者源属于 Folo 关系，不是 TopHub 热榜；必须确认持续性和互补性。"
  evidence_locator: "docs/ops/tophub-entertainment-photography-closure-20260705.md#Folo复查候选Huiliu"
  next_action: "2026-07-17 Folo 复查。"
```

---

## 8. Folo 边界

```yaml
- source_id: "娱乐 Folo 边界"
  current_route: folo_unchanged_review_candidates_only
  decision_status: recovered_from_pr_and_platform_verified
  selected_reason: "2026-07-17 复查前保持 Folo 27 项。娱乐目录只形成复查候选，不做立即订阅。"
  rejected_alternatives:
    - "把音乐榜单加入 Folo"
    - "把电影 / 游戏 / 摄影平台榜单加入 Folo"
    - "把所有强候选作者立即加入 Folo"
  route_reason: "最终复查候选为基本読書、木遥的窗子、Yuko's Blog、Huiliu；观察 maxOS、Velas 电波站。音乐榜单、平台榜单、报刊式娱乐新闻和视觉热榜不进 Folo。"
  reuse_rule: "娱乐 Folo 只考虑作者、节目、评论者、稳定长期关系；平台榜单和分类热榜留 TopHub 或按需。"
  evidence_locator: "docs/ops/tophub-entertainment-final-reconciliation-20260705.md#Folo边界; docs/ops/tophub-entertainment-pending-actions-20260705.md#Folo复查候选"
  next_action: "2026-07-17 复查。"
```

---

## 9. 本补充 ledger 的复用规则

1. 娱乐不是粉圈、明星私生活或平台热搜，而是作品、评论、产业、传播温度和视觉漫游。
2. 时尚源只有能看见资本、制造、供应链、规则和市场变化时，才进数据与结构。
3. 影评、书评和长评论放慢读与思想，不放作品发现榜。
4. 音乐常驻只保留互补入口：全球专辑热度、华语新专辑、中文原创单曲。
5. 剧集源优先低负担、跨平台、当期口碑榜，不订阅单平台和地域拆分大榜。
6. 娱乐新闻只留作品、创作者和产业编辑入口；明星私生活与粉圈热搜排除。
7. 游戏源按文化故事、综合节目、独立游戏设计、行业商业、折扣商店、攻略社区分层；常驻只留互补层。
8. 视觉源优先低频人工策展和摄影叙事，不以壁纸、器材新闻、商业摄影或平台大流为常驻。
9. 取消节点必须说明失效、低密度、短周期或静态旧榜原因；不能只说“不喜欢”。
10. Folo 只收长期作者 / 节目 / 评论者关系，不收榜单。

---

## 10. 仍需继续追回

娱乐最终动作、分组排序、10 个新增、3 个取消、胶片保留和 Folo 候选边界已补成第一份 ledger。

继续待办：

- 娱乐剩余小标签专项：B站、短视频、豆瓣、AcFun、虎扑、体育、足球、篮球、网文、播客、喜马拉雅、QQ音乐、网易云音乐、起点/纵横/七猫等可继续分组追回；
- 等用户本地提交 master checklist 与 coverage audit 后，再统一登记本文件和此前新增 ledger；
- 旧 ChatGPT 对话追索所有 partial 条目。
