# PR #372 判断链追回：TopHub / Folo 系统总分工（2026-07-06）

## 0. 边界

本文件追回 PR #372 中 TopHub 能力地图、使用配置、全部订阅顺序、Folo 分类重组与订阅源账本所形成的系统级判断链。

纳入文件：

- `资源/雷达/TopHub 能力地图.md`
- `资源/雷达/TopHub 使用配置.md`
- `资源/雷达/TopHub 全部订阅顺序.md`
- `资源/雷达/Folo 分类重组.md`
- `资源/雷达/订阅源账本.yml`
- `资源/雷达/订阅系统.md`
- `资源/雷达/信息流最终分工.md`

边界说明：

- 本文件不是重新审核 80 个 TopHub 节点或 27 个 Folo 来源。
- 单个来源的逐源理由已分别沉积在综合、科技、财经、购物、报刊、娱乐、社区、开发、AI、浏览器与文化等专项 ledger。
- 本文件追回的是：为什么 TopHub / Folo / 沃壤 / 原平台要分层，为什么 7 个用途分组这样用，为什么保持关闭某些自动能力，为什么复查候选不等于已订阅。
- 当前不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 总规则：TopHub 看见回声，Folo 接住原声与媒介，沃壤留下坐标

```yaml
- source_id: "TopHub / Folo / 沃壤总分工"
  current_route: layered_information_system
  decision_status: recovered_from_pr
  selected_reason: "TopHub 当前承担公共注意力、结构观察、内容发现和少量精确追踪；Folo 承担少量稳定的一手关系、深读、多媒体原貌和行动提醒；沃壤只保存关系、判断、复查日期和真正产生的变化。"
  rejected_alternatives:
    - "把 Folo 当 TopHub 的另一份新闻镜像"
    - "把 TopHub 当未读清零系统"
    - "把所有候选来源都写入正式订阅账本"
    - "把平台收藏、OPML、豆瓣、小宇宙、得到、BibiGPT 全部镜像到 GitHub"
  route_reason: "信息流最终分工明确为：TopHub 看见回声，Folo 接住原声与媒介，沃壤留下坐标。TopHub 标题只表示发现；Folo 未读只表示缓存；BibiGPT 转录只表示取水完成；GitHub 与今日报只承接真正形成责任、项目变化或需要跟进的事项。"
  reuse_rule: "任何来源先判断它当前应由 TopHub、Folo、原平台、BibiGPT、沃壤待机还是按需访问承担。订阅关系与内容沉积是两个不同问题。"
  evidence_locator: "资源/雷达/信息流最终分工.md#一句话; #最终路由; 资源/雷达/订阅系统.md#两条管道"
  next_action: "作为系统总边界。"
```

---

## 2. TopHub：为什么是 80 项公共雷达，不是第二个 Folo

```yaml
- source_id: "TopHub 当前 80 项"
  current_route: public_radar_and_discovery_layer
  decision_status: recovered_from_pr_platform_verified
  selected_reason: "TopHub 当前真实订阅为 80 项：七个用途分组共有 79 个内容节点，第 80 项为系统级节点 `专栏｜订阅聚合`。它覆盖公共温度、科技雷达、数据与结构、慢读与思想、生活与社区、影游音乐、视觉与自然，以及系统聚合入口。"
  rejected_alternatives:
    - "继续以 73 或 84 作为当前状态"
    - "把 TopHub 看成少量新闻热榜"
    - "把 TopHub 节点全迁入 Folo"
    - "因单日榜单波动频繁改动结构"
  route_reason: "TopHub 的能力地图说明它横跨综合新闻与舆论、科技 / AI / 企业 IT / 开发、科学科普、统计财经调查、影视游戏音乐、视觉每日图像、书籍影评文化内容、政务教育和重复报道压缩。使用配置和全部订阅顺序又确认：财经和报刊等量替换已执行，四个失效节点已取消不补位，当前真实总量为 80。"
  reuse_rule: "TopHub 用于遇见、标题级扫描、公共回声、结构发现和候选发现；不清未读，不自动制造待办，不因榜单一日波动改结构。"
  evidence_locator: "资源/雷达/TopHub 能力地图.md#结论; 资源/雷达/TopHub 全部订阅顺序.md#当前平台真实顺序; 资源/雷达/TopHub 使用配置.md#复查边界"
  next_action: "当前结构闭环，后续只按证据复查。"
```

---

## 3. TopHub 账号与功能设置：关闭自动化，避免未读债务和误触发

```yaml
- source_id: "TopHub 账号偏好与关闭项"
  current_route: low_noise_manual_review_configuration
  decision_status: recovered_from_pr_platform_verified
  selected_reason: "当前配置保留每行 4 个卡片，关闭赞助商广告和首页推荐节点，关闭首页校务 / 政务聚合，开启首页专栏节点聚合。通知机器人、微信/飞书等外部通知渠道、自动刷新、未经验证的过滤规则、批量提交信源或 OPML 均保持关闭。"
  rejected_alternatives:
    - "开启通知机器人"
    - "开启外部通知"
    - "开启自动刷新"
    - "启用未经验证的过滤规则"
    - "批量提交信源或 OPML"
    - "把校务 / 政务首页聚合做成日常入口"
  route_reason: "TopHub 是公共注意力与阶段性雷达，不承担未读清零和自动提醒。政务、校务、购物、体育、财经、法务等大量内容是任务触发，不应由通知或自动刷新制造责任。"
  reuse_rule: "凡会把候选、热榜或公共事件变成外部通知、即时责任或自动变更的功能，默认保持关闭；只有具体任务和退出条件明确时才单点启用。"
  evidence_locator: "资源/雷达/TopHub 使用配置.md#已完成的账号偏好; #保持关闭"
  next_action: "保持关闭。"
```

---

## 4. 七个用途分组：不是主题分类，而是打开方式

```yaml
- source_id: "TopHub 七个用途分组"
  current_route: usage_based_groups_not_topic_folders
  decision_status: recovered_from_pr_platform_verified
  selected_reason: "七组不是传统主题收藏夹，而是不同打开方式：公共温度、科技雷达、数据与结构、慢读与思想、生活与社区、影游音乐、视觉与自然。"
  rejected_alternatives:
    - "按 TopHub 原始目录机械保存"
    - "按领域越细越好地拆分"
    - "用一个总首页承载所有用途"
    - "把所有文化内容都塞进娱乐"
  route_reason: "使用配置明确了每组使用方式：公共温度只在需要知道公共世界正在发生什么时打开；科技雷达观察产品、公司、企业 IT、科学、开源及官方设备行动；数据与结构研究统计、经济、制度、全球化、汽车、消费与产业；慢读与思想只在有完整阅读时间时打开；生活与社区保留日常生活、社群与阶段入口；影游音乐发现作品、评论、创作与产业，不追粉圈八卦；视觉与自然用于视觉漫游、摄影叙事、自然与宇宙观察。"
  reuse_rule: "新来源先问它要求什么打开状态：扫一眼、结构研究、完整阅读、生活任务、作品发现、视觉漫游，还是任务触发。按打开方式分组，不按目录名。"
  evidence_locator: "资源/雷达/TopHub 使用配置.md#最终分组; #分组使用方式"
  next_action: "保持七组。"
```

---

## 5. 公共温度：公共世界背景，不是事实裁决

```yaml
- source_id: "公共温度 16 项"
  current_route: tophub_public_temperature
  decision_status: recovered_from_pr_platform_verified
  selected_reason: "公共温度当前包含微博、小红书、知乎、抖音、微信、澎湃、中国新闻周刊、新京报、南方周末、观察者网、参考消息、新闻联播、半月谈和新华每日电讯等，承担公共注意力与平台温度。"
  rejected_alternatives:
    - "用 Folo 逐个订阅公共新闻热榜"
    - "把公共温度当事实核验"
    - "把话题传播量当重要性或行动指令"
    - "继续保留已失效的微信热词、后续 Live、人民日报电子版"
  route_reason: "当前公共温度 16 项来自更新后的真实顺序。报刊等量替换已执行：人民日报电子版由新华每日电讯替换；四个失效节点中的微信热词和后续 Live 已取消不补位。公共温度负责看见世界正在谈什么，事实成立与行动判断仍回原始来源和任务上下文。"
  reuse_rule: "公共温度源只提供传播与公共回声；涉及事实、法律、健康、金融、政治或安全时回官方、原始报道、监管、数据或专业来源核验。"
  evidence_locator: "资源/雷达/TopHub 全部订阅顺序.md#01—16公共温度; #当前失效节点清理; #报刊目录最终方案"
  next_action: "保持。"
```

---

## 6. 科技雷达：产品/公司/基础设施/科学/开源/官方行动分层

```yaml
- source_id: "科技雷达 12 项"
  current_route: tophub_technology_radar
  decision_status: recovered_from_pr_platform_verified
  selected_reason: "科技雷达当前承担中文科技公司与产品、国际科技公司与平台、中文消费科技、企业 IT / 云 / 数据库 / 安全 / 数据中心 / 软件供应链、综合科普、国际科学新闻、低频开源项目发现、官方维修召回和应用安全精选。"
  rejected_alternatives:
    - "让 Folo 承担泛科技媒体、热门文章和科学新闻"
    - "订阅所有专业安全源和开发源"
    - "把单个漏洞、召回或热帖自动变成个人责任"
  route_reason: "科技雷达由极客公园、少数派、TechCrunch、The Verge、IT之家、The Register、先知社区、科普中国头条、果壳科学人、Science Magazine、HelloGitHub、Apple 支持维修扩展计划组成。它把公共发现放在 TopHub；Folo 只保留一手官方、独立作者、低频周报与当前项目源。"
  reuse_rule: "科技源先判定是消费科技、国际平台、企业基础设施、应用安全、科普、科学新闻、开源策展还是官方行动。公共发现进 TopHub；项目采用、官方源、独立作者再进 Folo 或项目。"
  evidence_locator: "资源/雷达/TopHub 全部订阅顺序.md#17—28科技雷达; 资源/雷达/TopHub 能力地图.md#科技AI企业IT与开发"
  next_action: "保持。"
```

---

## 7. 数据与结构：统计、政策经济、调查、全球化、汽车、终端、消费产业

```yaml
- source_id: "数据与结构 12 项"
  current_route: tophub_data_and_structure
  decision_status: recovered_from_pr_platform_verified
  selected_reason: "数据与结构当前承担正式统计与解释、财经政策商业入口、日本与东亚产业视角、国际财经解释、经济/质量/社会调查、企业全球化、汽车产业、科技终端市场和消费产业结构。"
  rejected_alternatives:
    - "保留停止更新的财新点击/评论排行榜"
    - "加入更多财经快讯、券商策略、论坛和报告库"
    - "用报刊财经版面替代任务源"
    - "把结构来源全部加入 Folo"
  route_reason: "财经等量替换已执行：财新点击和评论排行榜取消，加入财新首页推荐与日经中文每日最新，数据与结构保持 12。财新首页继续按真实内容体验观察，不把名称本身当精选的充分证据。"
  reuse_rule: "数据与结构只收能补结构缺口的入口；快讯、论坛、策略报告和版面型报刊按任务调用。结构源也要防重复和停止更新。"
  evidence_locator: "资源/雷达/TopHub 全部订阅顺序.md#29—40数据与结构; #财经目录最终方案"
  next_action: "保持。"
```

---

## 8. 慢读与思想：完整阅读窗口，不是未读堆积

```yaml
- source_id: "慢读与思想 13 项"
  current_route: tophub_slow_read_and_thought
  decision_status: recovered_from_pr_platform_verified
  selected_reason: "慢读与思想当前承担封面故事、思想文章、知乎日报、虚构阅读、微信读书新书、豆瓣新书、书评、电影文化、影评、古籍档案、器物、每日思想和历史日期入口。"
  rejected_alternatives:
    - "把它当每日必须清空的阅读箱"
    - "继续保留静态旧榜豆瓣非虚构"
    - "把所有书影音榜、书评、影评、古籍和思想内容混进娱乐"
  route_reason: "慢读与思想只在有完整阅读时间时打开。娱乐重构后新增豆瓣电影影评和书格，取消豆瓣非虚构旧榜，使这个分组从作品热度扩展到评论、解释和原始文化材料。"
  reuse_rule: "如果来源需要完整时间和解释性阅读，放慢读；如果只是作品发现，放影游音乐；如果是系统课程，回得到；如果是具体书影音关系，回豆瓣。"
  evidence_locator: "资源/雷达/TopHub 全部订阅顺序.md#41—53慢读与思想; 资源/雷达/TopHub 使用配置.md#分组使用方式"
  next_action: "保持。"
```

---

## 9. 生活与社区：日常任务和社群观察，不是泛论坛流

```yaml
- source_id: "生活与社区 7 项"
  current_route: tophub_life_and_community
  decision_status: recovered_from_pr_platform_verified
  selected_reason: "生活与社区当前包含下厨房、科普中国辟谣、半月谈健康、豆瓣话题广场、豆瓣小组、V2EX 周报和马蜂窝热门游记，承担日常饮食、辟谣、健康核验、公共生活讨论、技术社区周观察和旅行经验。"
  rejected_alternatives:
    - "保留国家留学网综合项目专栏和 Yoho!潮流志已失效节点"
    - "把地方门户、校园论坛、虎扑、AcFun、快手等平台社区全部常驻"
    - "把社区热榜直接等同于事实或行动"
  route_reason: "当前真实顺序显示国家留学网与 Yoho! 已因失效取消不补位。生活与社区保留少量能稳定补日常生活、健康与社区经验的入口，论坛热帖和地方门户按任务。"
  reuse_rule: "社区源只有周频、低噪声、可复用生活经验或明确任务价值时常驻；其他论坛热榜作为反应样本，不作为事实源。"
  evidence_locator: "资源/雷达/TopHub 全部订阅顺序.md#54—60生活与社区; #当前失效节点清理"
  next_action: "保持。"
```

---

## 10. 影游音乐：作品、评论、创作与产业，不追粉圈

```yaml
- source_id: "影游音乐 12 项"
  current_route: tophub_culture_works_radar
  decision_status: recovered_from_pr_platform_verified
  selected_reason: "影游音乐当前包含 Apple Music 专辑、豆瓣华语新碟、网易云原创榜、豆瓣电影一周口碑、IMDb 热门电影、豆瓣全球口碑剧集、新京报娱乐、B站每周必看、游研社、INDIENOVA、机核、喷嚏乐影。"
  rejected_alternatives:
    - "追粉圈八卦、明星私生活和娱乐平台分区热榜"
    - "把所有音乐榜、影视平台榜、游戏榜都常驻"
    - "把作品发现榜单放 Folo"
  route_reason: "娱乐最终动作将影游音乐从 8 扩到 12，但功能不是泛娱乐，而是作品发现、评论解释、独立游戏设计、游戏文化、中文原创音乐、华语新专辑、剧集口碑、纪录片片名发现和中国文化产业新闻。"
  reuse_rule: "文化作品源先分作品发现、评论、产业、创作过程、平台热度和粉圈流。常驻只留互补作品与产业入口；粉圈和平台分区榜按需或排除。"
  evidence_locator: "资源/雷达/TopHub 全部订阅顺序.md#61—72影游音乐; 资源/雷达/TopHub 使用配置.md#影游音乐"
  next_action: "保持。"
```

---

## 11. 视觉与自然：漫游与审美，不是壁纸/素材仓库

```yaml
- source_id: "视觉与自然 7 项"
  current_route: tophub_visual_nature_wander
  decision_status: recovered_from_pr_platform_verified
  selected_reason: "视觉与自然当前包含 iDaily、中国国家地理、CNU每日精选、胶片的味道、NASA每日星球、北京天文馆每日一图和果壳物种日历，承担视觉漫游、摄影叙事、自然与宇宙观察。"
  rejected_alternatives:
    - "把壁纸、商业摄影、站酷、500px、器材新闻和素材站全部常驻"
    - "因 CNU 新增取消胶片的味道"
    - "把 TopHub 的 NASA 与 Folo APOD 自动视为完全重复"
  route_reason: "视觉源优先低频人工策展、摄影叙事和自然科学观察。CNU 与胶片分工不同；TopHub 的每日图和 Folo 的 APOD 虽相近，但媒介形态、使用场景和订阅呈现并不完全相同。"
  reuse_rule: "视觉源先判定是每日图片、摄影叙事、自然科学、商业作品、壁纸、素材还是器材。只有审美关系和低噪声策展成立时常驻。"
  evidence_locator: "资源/雷达/TopHub 全部订阅顺序.md#73—79视觉与自然; 资源/雷达/TopHub 能力地图.md#视觉摄影与每日图片"
  next_action: "保持。"
```

---

## 12. 系统聚合入口：专栏｜订阅聚合不是 Folo 候选池本身

```yaml
- source_id: "专栏｜订阅聚合"
  current_route: tophub_system_aggregate_entry
  decision_status: recovered_from_pr_platform_verified
  selected_reason: "它是 TopHub 第 80 项系统级节点，承担专栏类系统聚合入口，而不是具体作者或栏目本身。"
  rejected_alternatives:
    - "把 94 页专栏候选逐个加入 TopHub"
    - "把专栏聚合入口当作 Folo 已完成订阅池"
  route_reason: "专栏 1—94 页形成的是候选池与主题池，TopHub 保留上位聚合入口即可。具体作者、节目、官方博客和专业源需要进入 Folo 复查或按任务触发。"
  reuse_rule: "系统聚合入口用于发现和导航；具体来源的长期关系另行判断。"
  evidence_locator: "资源/雷达/TopHub 全部订阅顺序.md#80系统聚合入口"
  next_action: "保持。"
```

---

## 13. 精确追踪器：窄词，不建宽词

```yaml
- source_id: "TopHub 7 个精确追踪器"
  current_route: precise_low_frequency_tracking
  decision_status: recovered_from_pr_platform_verified
  selected_reason: "最终保留 7 个追踪器：Apple Foundation Models、蓝河操作系统、Apple Silicon、小 V Claw、vivo办公套件、Codex CLI、macOS。它们指向明确项目、产品、平台或技术对象。"
  rejected_alternatives:
    - "AI"
    - "Agent"
    - "Apple"
    - "MacBook"
    - "Codex"
    - "vivo"
    - "GeekPwn"
    - "联合国"
    - "马斯克"
    - "罗永浩"
    - "雷军"
    - "巴菲特"
  route_reason: "追踪词必须具体。Apple Foundation Models 只承担 AFM、苹果端侧模型、框架、API 和应用接入；Apple Silicon 承担芯片、MLX、Metal、本地 AI、容器与虚拟化；macOS 承担系统、应用、兼容、安全与日常故障；vivo 拆成蓝河、小 V Claw、vivo 办公套件；Codex CLI 承担本地 CLI、权限、沙箱、MCP 和工作流变化。平台不支持手工修改追踪器显示顺序，顺序不表示优先级。"
  reuse_rule: "追踪器只追明确对象，不追宽泛概念、公众人物、宽品牌或无限责任词。结果页仍需回官方文档、代码仓库、论文或公司公告核验。"
  evidence_locator: "资源/雷达/TopHub 能力地图.md#精确追踪能力; 资源/雷达/TopHub 使用配置.md#最终追踪器"
  next_action: "保持。"
```

---

## 14. TopHub 发现模块：动态、榜中榜、话题聚合、热点日历、热文库各自降级

```yaml
- source_id: "TopHub 发现模块"
  current_route: auxiliary_discovery_modules_with_limits
  decision_status: recovered_from_pr
  selected_reason: "TopHub 可见模块包括动态、榜中榜、话题聚合、热点日历、热文库等，确实能帮助发现公共回声、传播热度、事件聚合、日期灵感和平台趋势。"
  rejected_alternatives:
    - "动态模式替代 Folo"
    - "榜中榜承担每周公共议题摘要"
    - "话题聚合直接回答事实是否成立"
    - "热点日历直接写入 Google Calendar"
    - "热文库默认全部浏览"
  route_reason: "动态模式是背景墙，不适合每日逐条处理；榜中榜只能看粗略传播热度；话题聚合显示传播强度，不代表独立核验；热点日历是低可信度日期灵感库；热文库只有窄筛选时可作为平台趋势侦察。"
  reuse_rule: "发现模块只作线索。事实核验、日历写入、项目责任和长期判断都要回可靠来源或具体任务。"
  evidence_locator: "资源/雷达/TopHub 能力地图.md#已实测的发现模块"
  next_action: "按需使用。"
```

---

## 15. Folo：27 项按注意力用途，不按主题堆叠

```yaml
- source_id: "Folo 27 项分类重组"
  current_route: folo_attention_modes_observation_period
  decision_status: recovered_from_pr_platform_verified
  selected_reason: "Folo 中 27 个订阅已全部进入五个注意力分类：必看触发 3、深度阅读 7、漫游 6、行动提醒 4、观察期 7。当前进入两周观察期，复查日期为 2026-07-17。"
  rejected_alternatives:
    - "按主题堆叠订阅"
    - "用未读数判断来源价值"
    - "观察期自动续期"
    - "观察期内继续删除或导入新的 OPML"
    - "把名称省略的 7 项强行写入版本化账本"
  route_reason: "Folo 分类重组的目标是停止按主题堆叠，改按注意力用途组织。未读数是缓存，不是债务。7 项名称当前有意省略，只记录数量。自动化不得自动退订、自动提交或改变来源状态。"
  reuse_rule: "Folo 来源先归入触发、深读、漫游、行动提醒或观察期；每类按真实处理能力和打开方式复查，不按主题数量扩容。"
  evidence_locator: "资源/雷达/Folo 分类重组.md#总量核对; #使用规则; 资源/雷达/订阅源账本.yml"
  next_action: "2026-07-17 复查。"
```

---

## 16. Folo 五类：为什么这些来源放这些层

```yaml
- source_id: "Folo｜必看触发"
  current_route: folo_trigger
  decision_status: recovered_from_pr_platform_verified
  selected_reason: "Google Developers Blog、Google DeepMind News、Obsidian Plugins 可以在工作过程中打开，但只处理与当前项目相关的更新。"
  rejected_alternatives:
    - "工作时顺手进入公共雷达和漫游"
    - "把所有 Google / AI / Obsidian 相关源都放入触发"
  route_reason: "触发组是低频、一手、不能轻易漏掉的官方或项目更新；不是无目的浏览区。"
  reuse_rule: "触发源必须与当前项目或工具链相关，并可能触发明确行动；否则不放触发。"
  evidence_locator: "资源/雷达/Folo 分类重组.md#01必看触发; 资源/雷达/订阅系统.md#注意力模式"
  next_action: "复查重复度。"

- source_id: "Folo｜深度阅读"
  current_route: folo_deep_read
  decision_status: recovered_from_pr_platform_verified
  selected_reason: "小众软件、晚点长报道、Last Week in AI、darkreading、Krebs on Security、Steph Ango、有知有行进入晚间阅读橱窗，每次只选一篇，不以清空未读为目标。"
  rejected_alternatives:
    - "把长报道和专业深读当每日待办"
    - "用 TopHub 热榜替代完整阅读"
  route_reason: "深读需要完整时间。高质量不等于高频清空。"
  reuse_rule: "只有真正会完整阅读、且不是 TopHub 可替代的作者、长报道或专业深读，才放深读。"
  evidence_locator: "资源/雷达/Folo 分类重组.md#02深度阅读"
  next_action: "按打开率复查。"

- source_id: "Folo｜漫游"
  current_route: folo_wander
  decision_status: recovered_from_pr_platform_verified
  selected_reason: "Nat Geo Photo of the Day、NASA APOD、Magnum Photos、中国爬楼联盟、街拍中国等用于视觉、摄影与陌生领域漫游。"
  rejected_alternatives:
    - "把漫游未读视作债务"
    - "中国爬楼联盟与街拍中国长期并存不复查"
  route_reason: "漫游只在状态好时打开，过期内容可以直接标记已读。2026-07-17 复查时中国爬楼联盟与街拍中国最多保留一个。"
  reuse_rule: "漫游源必须改善视觉/媒介体验或提供陌生领域偶遇；过期可清空。相近视觉源最多保留互补者。"
  evidence_locator: "资源/雷达/Folo 分类重组.md#03漫游; #2026-07-17复查点"
  next_action: "最多保留一个相近源。"

- source_id: "Folo｜行动提醒"
  current_route: folo_reminder
  decision_status: recovered_from_pr_platform_verified
  selected_reason: "这一组 4 个来源当前只在本地清单保留具体名称，版本化账本记录数量与用途。保留标准是是否真的触发领取、观看或其他具体动作。"
  rejected_alternatives:
    - "把所有免费游戏、播出、项目状态提醒都长期保留"
    - "把无真实触发的提醒继续占订阅位"
  route_reason: "行动提醒没有触发真实动作的来源，应在复查时退出常驻。"
  reuse_rule: "提醒源必须触发可执行动作。若长期没有真实动作，退出 Folo，回 TopHub、原平台或按需。"
  evidence_locator: "资源/雷达/Folo 分类重组.md#04行动提醒; 资源/雷达/订阅源账本.yml#classification"
  next_action: "复查实际触发。"

- source_id: "Folo｜观察期"
  current_route: folo_trial
  decision_status: recovered_from_pr_platform_verified
  selected_reason: "Stack Overflow Blog、Google AI Developers、Logan Kilpatrick、News Minimalist、Hacker News 与两个名称省略来源进入观察期，两周后按打开率、重复度和独特价值判断。"
  rejected_alternatives:
    - "观察期自动续期"
    - "把观察源立即写成长期 active"
    - "Google AI Developers 与 Google Developers Blog 重复时双保留"
  route_reason: "观察期必须到期判断，不自动续期。Google AI Developers 与 Google Developers Blog 重复明显时保留信息更完整的一边；Stack Overflow Blog、News Minimalist、Logan Kilpatrick、Hacker News 没有真实打开或独特价值就退出。"
  reuse_rule: "trial 是有期限的试验，不是软性长期订阅。"
  evidence_locator: "资源/雷达/Folo 分类重组.md#05观察期; #2026-07-17复查点"
  next_action: "2026-07-17 明确去留。"
```

---

## 17. Folo 移出来源：退出不等于失去价值

```yaml
- source_id: "Folo 移出来源"
  current_route: moved_to_tophub_or_on_demand_or_retired_invalid
  decision_status: recovered_from_pr_platform_verified
  selected_reason: "Lifehacker、Cyber Security News、领研论文计算机、中国气象局每日天气提示、东方财富策略报告、华尔街见闻财经日历、少数派、36氪 24h、Bing 每日壁纸、InfoQ、金十数据、New Routes、Anthropic News、创意文章榜、创意作品榜等被移出或标记为其他状态。"
  rejected_alternatives:
    - "把移出解释为来源永久无价值"
    - "保留所有曾有价值来源导致 Folo 堆叠"
    - "把报告库、财经日历、天气提示、泛新闻、泛安全新闻都留在 Folo"
  route_reason: "移出后的状态分为 TopHub、on_demand、retired_invalid。退出 Folo 只说明它当前不适合占用持续未读位；它仍可能由 TopHub、原平台、按需查询或沃壤待机承担。"
  reuse_rule: "来源价值与当前路由分开。退出 Folo 后必须给去处：TopHub、按需、原平台、待机或失效退休。"
  evidence_locator: "资源/雷达/订阅源账本.yml#second_wave_removed; #removed_named_sources; 资源/雷达/订阅系统.md#来源状态"
  next_action: "按去处使用。"
```

---

## 18. 复查候选不是已订阅

```yaml
- source_id: "复查候选与历史审核证据"
  current_route: review_candidates_not_subscription_facts
  decision_status: recovered_from_pr
  selected_reason: "浏览器旧书签、TopHub 目录审核、专栏 94 页、AI 目录、娱乐和购物等产生大量候选。"
  rejected_alternatives:
    - "把复查候选写成已加入 Folo"
    - "在 2026-07-17 前立即扩容"
    - "导入新的 OPML"
    - "让自动化根据候选自动退订或提交"
  route_reason: "使用配置明确：TopHub 当前结构已经闭环，新来源只有形成独特、持续且可解释功能时才进入；Folo 在 2026-07-17 前保持 27 项不变；复查候选与历史审核证据保存在 docs/ops，不等于已订阅。"
  reuse_rule: "候选必须经过复查日、角色比较、一进一出、打开率与重复度判断后，才可能变成真实来源。"
  evidence_locator: "资源/雷达/TopHub 使用配置.md#复查边界; 资源/雷达/Folo 分类重组.md#边界"
  next_action: "2026-07-17 复查。"
```

---

## 19. 本补充 ledger 的复用规则

1. TopHub 看见回声，Folo 接住原声与媒介，沃壤留下坐标。
2. TopHub 当前真实为 80 项，不继续使用旧的 73/84 状态。
3. TopHub 是公共注意力、结构发现、内容发现和低频精确追踪层，不是未读箱。
4. Folo 当前 27 项按注意力模式分组，不按主题堆叠。
5. 未读数是缓存，不是债务；动态、榜中榜、话题聚合、热点日历、热文库都只是线索工具。
6. 七个 TopHub 分组按打开方式使用，不按目录名机械分类。
7. 精确追踪器只追明确对象，不建宽词、公众人物词和无限责任词。
8. 通知机器人、外部通知、自动刷新、过滤规则、批量提交保持关闭。
9. 复查候选不是订阅事实，docs/ops 证据不是平台执行。
10. 来源退出 Folo 不等于失去价值，必须给出 TopHub、按需、原平台、待机或失效退休的去处。

---

## 20. 仍需继续追回

TopHub / Folo 系统总分工已补成 ledger。

继续待办：

- 等用户本地提交 master checklist 与 coverage audit 后，再统一登记本文件和此前新增 ledger；
- 旧 ChatGPT 对话追索所有 partial 条目；
- 如需继续，可补 `TopHub 当前失效节点清理`、`设计目录`、`视觉与自然审计`、`账号与通知审计` 等剩余专项。 
