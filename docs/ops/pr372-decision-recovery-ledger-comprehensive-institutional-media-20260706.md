# PR #372 判断链追回：综合官方与调查媒体源（人民网 / 新华社 / 南方周末 / 半月谈 / CCTV）（2026-07-06）

## 0. 边界

本文件是 `综合` 大类的补充 ledger，追回 PR #372 中已经闭环的官方媒体、机构媒体与调查解释型媒体小标签：

- `docs/ops/tophub-comprehensive-people-closure-20260704.md`
- `docs/ops/tophub-comprehensive-xinhua-closure-20260704.md`
- `docs/ops/tophub-comprehensive-southern-weekly-closure-20260704.md`
- `docs/ops/tophub-comprehensive-banyuetan-closure-20260704.md`
- `docs/ops/tophub-comprehensive-cctv-closure-20260704.md`

本文件不声称综合大类全量完成。

判断链标准：

> 每个来源背后的判断链——为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 官方与调查媒体源总规则

```yaml
- source_id: "综合官方与调查媒体源总规则"
  current_route: keep_main_entry_plus_task_sources
  decision_status: recovered_from_pr
  selected_reason: "这一组来源包含官方议程、政策核验、地方执行、社会调查、消费治理、人物文化与政策解释等不同功能。"
  rejected_alternatives:
    - "把人民网 / 新华网 / 央视 / 半月谈 / 南方周末所有栏目都常驻"
    - "把同一媒体的栏目切片同时放入 TopHub 和 Folo"
    - "给人民网、新华社、CCTV、时政、国际、社会等宽泛词建立追踪器"
  route_reason: "常驻只保留能代表主入口或补出稳定新功能的节点；栏目切片按任务打开；具体政策、地区、人事、法律、消费、产业或节目问题回原始文件、官方源或具体节目。"
  reuse_rule: "以后遇到官方 / 大媒体目录，先找主入口和互补功能，而不是按栏目全收；能代表每日议程的留一个，能补结构现场的留一个，剩余栏目按任务。"
  evidence_locator: "tophub-comprehensive-people-closure-20260704.md; tophub-comprehensive-xinhua-closure-20260704.md; tophub-comprehensive-southern-weekly-closure-20260704.md; tophub-comprehensive-banyuetan-closure-20260704.md; tophub-comprehensive-cctv-closure-20260704.md"
  next_action: "无；作为复用规则。"
```

---

## 2. 人民网小标签

来源证据：`docs/ops/tophub-comprehensive-people-closure-20260704.md`

页面标注 `人民网 72个`，实际可识别 71 个节点。结论：不新增人民网常驻节点，继续保留 `人民日报｜电子版`，少数栏目按需，不进 Folo，不建宽泛人民网追踪器。

```yaml
- source_id: "人民日报｜电子版"
  current_route: tophub_public_temperature
  decision_status: recovered_from_pr
  selected_reason: "承担人民日报当日议程、中央层级政策和公共表述、重要评论、民生、经济、文化与地方稿件的报纸编排。相比人民网几十个栏目切片，电子版更能观察该媒体每天真正选择了什么。"
  rejected_alternatives:
    - "人民网要闻 / 时政 / 国际 / 观点 / 滚动 / 评论栏目全部常驻"
  route_reason: "继续放在 `公共温度`；它代表官方日常议程入口，避免同稿在人民网多个栏目重复出现。"
  reuse_rule: "同一官方媒体有电子版或主编排入口时，优先保留主入口，不叠加大量栏目切片。"
  evidence_locator: "docs/ops/tophub-comprehensive-people-closure-20260704.md#保留现有常驻入口"
  next_action: "无；保留。"

- source_id: "人民网｜社会频道 / 地方领导留言板"
  current_route: on_demand_public_service_and_local_execution
  decision_status: recovered_from_pr
  selected_reason: "社会频道能看到公共服务、就业、住房、儿童、消费纠纷、气象灾害和社会案件；地方领导留言板接近群众投诉、12345、公共服务、就业权益、社区治理和地方政策执行。"
  rejected_alternatives:
    - "生活与社区常驻"
    - "Folo"
  route_reason: "二者都有现实生活价值，但主题过宽、噪声高、旧稿和政策宣传混杂，与半月谈、新闻调查、每周质量报告及公共温度重叠。遇到具体地区、公共服务、投诉反馈或地方执行问题时按需。"
  reuse_rule: "现实问题入口不等于常驻生活流；有具体地区或投诉对象时再打开。"
  evidence_locator: "docs/ops/tophub-comprehensive-people-closure-20260704.md#人民网社会频道; #人民网地方领导留言板; #生活与社区"
  next_action: "按需。"

- source_id: "人民网｜观点*人民网评 / 人民日报重要言论库｜人民时评 / 评论理论类"
  current_route: on_demand_official_commentary
  decision_status: recovered_from_pr
  selected_reason: "可用于核对人民日报系对平台治理、劳动权益、家庭病床、数字生活、消费保护、公共争议和正式评论的表述。"
  rejected_alternatives:
    - "慢读与思想常驻"
    - "多个人民网评论栏目同时订阅"
  route_reason: "评论栏目大量内部重复，同一批人民时评和评论稿跨栏目出现；现有电子版会带出最重要评论，系统已有财新、FT、爱思想、半月谈和南方周末等评论入口。"
  reuse_rule: "评论池多入口只保留主入口和按需检索，不放大单一官方媒体评论权重。"
  evidence_locator: "docs/ops/tophub-comprehensive-people-closure-20260704.md#评论栏目大量内部重复; #不新增活跃但重复的栏目"
  next_action: "按需。"

- source_id: "人民网｜人事频道 / 知识产权 / 环保频道 / 时政频道 / 国家与地区频道"
  current_route: on_demand_policy_region_or_domain_entry
  decision_status: recovered_from_pr
  selected_reason: "人事频道适合核对任免和公示；知识产权适合商标、专利、版权和制度问题；环保频道适合河湖治理、生态、能源转型；时政和国家地区频道适合核对人民网口径。"
  rejected_alternatives:
    - "TopHub 常驻"
    - "Folo"
  route_reason: "这些是明确任务入口，不是每日阅读源；地域跨度大、重复转载多、部分更新不稳。"
  reuse_rule: "具体政策、职位、地区、知识产权或环保问题出现时打开对应任务源；不把任务源变成未读流。"
  evidence_locator: "docs/ops/tophub-comprehensive-people-closure-20260704.md#高价值按需入口"
  next_action: "按需。"

- source_id: "人民网陈旧 / 失真 / 商业污染 / 彩票节点"
  current_route: rejected
  decision_status: recovered_from_pr
  selected_reason: "无常驻价值。"
  rejected_alternatives:
    - "房产 / 文史 / 旅游 / 汽车 / 能源 / 通信 / 健康 / 法制 / 娱乐 / 食品 / 收藏 / 彩票等栏目常驻"
  route_reason: "多频道存在 2016—2022 年旧内容、标签与实际内容不符、商业宣传污染、重复滚动或彩票内容。TopHub 刷新时间不能代表文章新鲜度。"
  reuse_rule: "栏目旧、标签失真、商业宣传或彩票内容是明确排除信号。"
  evidence_locator: "docs/ops/tophub-comprehensive-people-closure-20260704.md#明确排除陈旧失真商业污染或损坏"
  next_action: "无；排除。"
```

---

## 3. 新华社小标签

来源证据：`docs/ops/tophub-comprehensive-xinhua-closure-20260704.md`

新华社 10 个节点完整展示。结论：不新增常驻；时政、地方联播、科技、国际各有按需价值；不进 Folo，不建宽泛新华社追踪器。

```yaml
- source_id: "新华社小标签整体"
  current_route: no_new_tophub_with_policy_region_tech_on_demand
  decision_status: recovered_from_pr
  selected_reason: "新华网的价值不在于栏目数量，而在于出现具体政策、地区、产业或国际问题时可用于官方口径核验。"
  rejected_alternatives:
    - "时政 / 财经 / 国际 / 网评 / 军事 / 地方联播 / 科技 / 滚动 / 即时 / 体育全部常驻"
    - "新华社 Folo"
    - "新华社宽泛追踪器"
  route_reason: "10 个节点是新华网内部栏目切片，并大量转载人民日报、央视、环球时报、参考消息、经济日报、科技日报、北京日报等其他媒体。没有一个栏目补出当前系统缺失的常驻能力。"
  reuse_rule: "权威性不是常驻理由；要看该栏目是否补出新功能。官方栏目更适合具体问题核验。"
  evidence_locator: "docs/ops/tophub-comprehensive-xinhua-closure-20260704.md#核心判断; #最终结论"
  next_action: "无；新华社已闭环。"

- source_id: "新华网｜时政频道"
  current_route: on_demand_policy_verification
  decision_status: recovered_from_pr
  selected_reason: "适合核对政策原文、国务院会议、制度变化和官方解释，如医保目录、职业资格、电子商务法、就业优先规划等。"
  rejected_alternatives:
    - "公共温度常驻"
    - "数据与结构常驻"
  route_reason: "页面同时充满党建、纪念、主题宣传、领导人活动和机构报道；不如半月谈时政讲解适合日常理解，也不如具体部委和法规原文精确。"
  reuse_rule: "官方政策栏目用于核验，不用于日常解释；真正判断回部委、法规原文和具体文件。"
  evidence_locator: "docs/ops/tophub-comprehensive-xinhua-closure-20260704.md#新华网时政频道"
  next_action: "按需。"

- source_id: "新华网｜地方联播"
  current_route: on_demand_regional_execution_source
  decision_status: recovered_from_pr
  selected_reason: "能看到地方灾害、天气、公共安全、交通、就业、公积金、产业、乡村、生态与城市发展等区域执行信息。"
  rejected_alternatives:
    - "公共温度常驻"
    - "生活与社区常驻"
  route_reason: "条目上百、党建、地方宣传、文旅、事故、天气和产业混杂，地域分散，对多数城市没有持续现实关系。"
  reuse_rule: "地方联播在具体地区、城市政策、灾害或地方执行任务出现时使用，不做全国生活常驻流。"
  evidence_locator: "docs/ops/tophub-comprehensive-xinhua-closure-20260704.md#新华网地方联播"
  next_action: "按需。"

- source_id: "新华网｜科技频道"
  current_route: on_demand_domestic_tech_policy_source
  decision_status: recovered_from_pr
  selected_reason: "用于观察国内科技产业政策、算力、5G/6G、具身智能、机器人、航天、先进制造和官方科技口径。"
  rejected_alternatives:
    - "科技雷达常驻"
    - "Folo"
  route_reason: "与现有科技雷达高度重叠，同一 AI 和产业政策主题反复出现；它是官方科技聚合，不是一手工程或研究源。"
  reuse_rule: "国内科技政策口径按需看官方聚合，但工程和研究仍回一手源。"
  evidence_locator: "docs/ops/tophub-comprehensive-xinhua-closure-20260704.md#新华网科技频道"
  next_action: "按需。"

- source_id: "新华网｜财经频道 / 即时新闻 / 体育新闻 / 网评 / 国际 / 滚动新闻 / 军事"
  current_route: rejected_or_on_demand_xinhua_slices
  decision_status: recovered_from_pr
  selected_reason: "国际频道可按需核对中国官方国际报道；滚动新闻有通讯社即时性。"
  rejected_alternatives:
    - "常驻 TopHub"
    - "Folo"
  route_reason: "财经频道企业公关污染严重；即时新闻时间和主题不稳定；体育不是长期需求；网评与其他评论入口重叠；国际、滚动和军事与参考消息、央视、人民日报、澎湃、现有军事结论重叠。"
  reuse_rule: "新华社栏目除明确任务外不叠加；具体事件用具体政策、地区、产业计划或国际事件追踪。"
  evidence_locator: "docs/ops/tophub-comprehensive-xinhua-closure-20260704.md#各节点判断"
  next_action: "按需或排除。"
```

---

## 4. 南方周末小标签

来源证据：`docs/ops/tophub-comprehensive-southern-weekly-closure-20260704.md`

南方周末 9 个节点完整展示。现有 `南方周末｜热门文章` 已订阅在公共温度。

```yaml
- source_id: "南方周末｜热门文章"
  current_route: tophub_public_temperature
  decision_status: recovered_from_pr
  selected_reason: "承担该媒体的横向精选入口，覆盖新闻、人物、文化、科学、生活和南方人物周刊等不同栏目。"
  rejected_alternatives:
    - "今日推荐"
    - "新闻"
    - "人物"
    - "观点"
    - "文化"
    - "影像"
    - "视频"
    - "生活"
  route_reason: "继续只保留热门文章，避免同一家媒体的内容被多个栏目拆开造成重复和权重放大。"
  reuse_rule: "同一媒体总榜 / 推荐 / 栏目切片重复时，保留横向入口，栏目按需。"
  evidence_locator: "docs/ops/tophub-comprehensive-southern-weekly-closure-20260704.md#南方周末热门文章; #最终结论"
  next_action: "无；保留。"

- source_id: "南方周末｜人物 / 文化"
  current_route: on_demand_high_value_columns
  decision_status: recovered_from_pr
  selected_reason: "人物栏目关注人物经历、文化人物、城乡生活和时代侧面；文化栏目包含写作、阅读、表演、文化人物与艺术观察。"
  rejected_alternatives:
    - "Folo"
    - "慢读与思想常驻"
    - "同时订阅南方周末多个栏目"
  route_reason: "两者都有独特气质，但热门文章已经会带出南方人物周刊和文化写作；长期订阅会放大单一媒体权重。"
  reuse_rule: "若未来长期阅读某一栏目，应只选人物或文化其中一个进入观察期，不一次加入多个栏目。"
  evidence_locator: "docs/ops/tophub-comprehensive-southern-weekly-closure-20260704.md#南方周末人物; #南方周末文化; #folo"
  next_action: "按需。"

- source_id: "南方周末其他栏目"
  current_route: rejected_or_on_demand_same_media_slices
  decision_status: recovered_from_pr
  selected_reason: "个别栏目质量不差。"
  rejected_alternatives:
    - "今日推荐"
    - "新闻"
    - "观点"
    - "影像"
    - "视频"
    - "生活"
  route_reason: "与热门文章、澎湃、新京报、中国新闻周刊、参考消息、财新、FT、爱思想、三联、豆瓣、Cinephilia 等现有入口重叠；视频和影像只是媒介形态或新闻摄影，不补新功能。"
  reuse_rule: "单一媒体内部栏目切片不因质量尚可就全加；看是否补出新来源关系。"
  evidence_locator: "docs/ops/tophub-comprehensive-southern-weekly-closure-20260704.md#各节点判断"
  next_action: "按需或不新增。"
```

---

## 5. 半月谈小标签

来源证据：`docs/ops/tophub-comprehensive-banyuetan-closure-20260704.md`

半月谈 15 个节点，两页全部查看。结论：保留 `今日谈`，新增 `健康` 替换 `壹心理`，`时政讲解` 是最高价值政策解释按需入口。

```yaml
- source_id: "半月谈｜今日谈"
  current_route: tophub_public_temperature
  decision_status: recovered_from_pr
  selected_reason: "承担半月谈的综合精选和公共议题入口。"
  rejected_alternatives:
    - "半月谈多个栏目同时常驻"
  route_reason: "继续保留在公共温度；它作为综合入口可覆盖部分政策解释和公共议题，不需要叠加多个半月谈栏目。"
  reuse_rule: "同一机构已有综合入口时，新增栏目必须提供明确独立功能。"
  evidence_locator: "docs/ops/tophub-comprehensive-banyuetan-closure-20260704.md#半月谈最终路由"
  next_action: "无；保留。"

- source_id: "半月谈｜健康"
  current_route: tophub_life_and_community
  decision_status: recovered_from_pr
  selected_reason: "新增健康入口替换壹心理，承担日常健康、慢病、公共健康和风险解释，与科普中国辟谣互补。"
  rejected_alternatives:
    - "壹心理"
    - "网易健康 / 人民网健康 / ZAKER健康 / 微信健康榜等平台健康栏目"
  route_reason: "进入生活与社区；不是平台健康热榜，而是更低噪声的公共健康解释来源。"
  reuse_rule: "健康常驻源必须低噪声、少营销、少猎奇，并能和辟谣来源互补。"
  evidence_locator: "docs/ops/tophub-comprehensive-banyuetan-closure-20260704.md#半月谈最终路由; docs/ops/tophub-comprehensive-closeout-20260704.md#壹心理--半月谈健康"
  next_action: "无；已新增。"

- source_id: "半月谈｜时政讲解"
  current_route: high_value_on_demand_policy_explainer
  decision_status: recovered_from_pr
  selected_reason: "解释政策文件改变了什么、新制度如何运行、监管、投资、社会救助、城市更新、就业和服务业如何落地，能补制度解释能力。"
  rejected_alternatives:
    - "立即加入数据与结构"
    - "同时扩大多个半月谈栏目"
  route_reason: "它有潜力，但今日谈已覆盖部分政策解释；新增健康后半月谈已有两个常驻角色；还需证明更新稳定性、条目密度和长期不可替代性。"
  reuse_rule: "政策解释栏目先按需；连续多次真实使用后再进入观察期，而不是因样本有潜力直接常驻。"
  evidence_locator: "docs/ops/tophub-comprehensive-banyuetan-closure-20260704.md#对时政讲解的最终判断"
  next_action: "按需；未来连续使用再观察。"

- source_id: "半月谈｜基层治理 / 民生话题 / 改革创新 / 区域风采 / 脱贫攻坚等"
  current_route: on_demand_or_rejected_banyuetan_slices
  decision_status: recovered_from_pr
  selected_reason: "基层治理与民生话题有任务价值。"
  rejected_alternatives:
    - "要闻TOP10"
    - "评论"
    - "国际"
    - "科技"
    - "文化"
    - "地方观察"
    - "脱贫攻坚"
    - "企业资讯"
    - "改革创新"
    - "区域风采"
  route_reason: "多数是同一媒体栏目切片；改革创新、区域风采陈旧或宣传性强；脱贫攻坚仍停留旧政策阶段；企业资讯明确排除。"
  reuse_rule: "政策阶段已经变化、栏目仍停留旧语境时，不进入常驻；基层治理和民生按具体问题调用。"
  evidence_locator: "docs/ops/tophub-comprehensive-banyuetan-closure-20260704.md#第2页判断; #半月谈最终路由"
  next_action: "按需或排除。"
```

---

## 6. CCTV 小标签

来源证据：`docs/ops/tophub-comprehensive-cctv-closure-20260704.md`

CCTV 14 个节点，两页完整覆盖。现有 TopHub 已订阅 `新闻联播`、`经济半小时`、`每周质量报告`；本次新增 `新闻调查` 到数据与结构。

```yaml
- source_id: "中央电视台｜新闻联播"
  current_route: tophub_public_temperature
  decision_status: recovered_from_pr
  selected_reason: "承担官方议程、重大政策和中央层级公共信息入口。它不是为了覆盖所有新闻，而是观察国家议程与官方优先级。"
  rejected_alternatives:
    - "CCTV 国内新闻"
    - "CCTV 国际新闻"
    - "焦点访谈常驻"
  route_reason: "继续保留在公共温度；国内 / 国际新闻流会增加高频条目但不增加结构。"
  reuse_rule: "央视新闻常驻保留议程主入口，不叠加快讯流。"
  evidence_locator: "docs/ops/tophub-comprehensive-cctv-closure-20260704.md#中央电视台新闻联播"
  next_action: "无；保留。"

- source_id: "中央电视台｜经济半小时"
  current_route: tophub_data_and_structure
  decision_status: recovered_from_pr
  selected_reason: "持续提供产业、企业、技术应用、能源、汽车后市场和经济现场调查，是数据与结构中的实地经济入口。"
  rejected_alternatives:
    - "经济信息联播常驻"
  route_reason: "继续保留；相比日更经济新闻汇总，它更能提供现场调查和结构观察。"
  reuse_rule: "经济类节目优先保留现场调查，不保留只按日期汇总的日更新闻流。"
  evidence_locator: "docs/ops/tophub-comprehensive-cctv-closure-20260704.md#中央电视台经济半小时"
  next_action: "无；保留。"

- source_id: "中央电视台｜每周质量报告"
  current_route: tophub_data_and_structure
  decision_status: recovered_from_pr
  selected_reason: "周频稳定覆盖消费、产品质量、认证、直播乱象、医疗美容、玩具、酒类、噪声和支付等具体问题，功能清晰。"
  rejected_alternatives:
    - "黑猫投诉常驻"
    - "CCTV 法治新闻 / 今日说法常驻"
  route_reason: "继续保留在数据与结构，承担消费、产品与市场治理入口。"
  reuse_rule: "消费治理常驻优先选择周频、调查型、可解释结构的问题入口，而不是高频投诉或案件流。"
  evidence_locator: "docs/ops/tophub-comprehensive-cctv-closure-20260704.md#中央电视台每周质量报告"
  next_action: "无；保留。"

- source_id: "中央电视台｜新闻调查"
  current_route: tophub_data_and_structure
  decision_status: recovered_from_pr
  selected_reason: "周频、更新压力低，以现场、人物和地区为基础，同时呈现政策、执行、行业、社区与普通人的关系；补充社会结构的实地调查层。"
  rejected_alternatives:
    - "更多央视快讯"
    - "焦点访谈常驻"
    - "新闻1+1 常驻"
  route_reason: "进入数据与结构，属于扩容新增而非替换。它补的是国家统计局数据、财新财经报道、经济半小时产业调查和每周质量报告消费调查之外的社会结构现场。"
  reuse_rule: "新增大媒体节目只有在补出明确新功能时才扩容；新闻调查补结构现场，快讯不补。"
  evidence_locator: "docs/ops/tophub-comprehensive-cctv-closure-20260704.md#中央电视台新闻调查; #最终结论"
  next_action: "无；已新增。"

- source_id: "中央电视台｜新闻1+1"
  current_route: high_value_on_demand_explainer_program
  decision_status: recovered_from_pr
  selected_reason: "每天选择一个明确问题做约 26 分钟解释，如灾害、企业出海、陪诊、商标法、能源、就业、高考志愿、平台反内卷和反诈，是当前政策与社会问题快速解释入口。"
  rejected_alternatives:
    - "数据与结构常驻"
    - "Folo"
  route_reason: "日频节目会提高央视权重，且与半月谈时政讲解、新闻调查、经济半小时和每周质量报告有交叉；TopHub 按标题选择观看已经足够，不制造每日未读压力。"
  reuse_rule: "高价值日频解释节目先按需，不因好用就变成每日队列。"
  evidence_locator: "docs/ops/tophub-comprehensive-cctv-closure-20260704.md#中央电视台新闻1+1"
  next_action: "按需。"

- source_id: "CCTV 法治新闻 / 今日说法 / 焦点访谈"
  current_route: on_demand_legal_or_topic_programs
  decision_status: recovered_from_pr
  selected_reason: "法治新闻和今日说法有法律、诈骗、平台和消费权益实用价值；焦点访谈有时提供社会调查。"
  rejected_alternatives:
    - "生活与社区常驻"
    - "数据与结构常驻"
    - "Folo"
  route_reason: "法治新闻案件和通报占比高；今日说法日频且个案化；焦点访谈价值波动并与新闻联播、新闻调查、经济半小时交叉。"
  reuse_rule: "案件型节目按问题打开，不建立连续未读队列。"
  evidence_locator: "docs/ops/tophub-comprehensive-cctv-closure-20260704.md#CCTV央视新闻法治新闻; #中央电视台焦点访谈; #中央电视台今日说法"
  next_action: "按需。"

- source_id: "CCTV 国际 / 国内 / NBA / 中国电视报 / 致富经 / 经济信息联播"
  current_route: rejected_or_on_demand_cctv_slices
  decision_status: recovered_from_pr
  selected_reason: "国际和国内新闻可在具体事件核对央视报道时按需。"
  rejected_alternatives:
    - "公共温度常驻"
    - "影游音乐"
    - "数据与结构常驻"
  route_reason: "国际和国内新闻与参考消息、新华社、人民日报、新闻联播重叠；NBA 是高频赛事流；中国电视报和致富经陈旧；经济信息联播缺少主题筛选且与现有经济来源重叠。"
  reuse_rule: "赛事流、旧电子报、创业致富旧栏目和无主题筛选的日更汇总不常驻。"
  evidence_locator: "docs/ops/tophub-comprehensive-cctv-closure-20260704.md#各节点判断; #最终结论"
  next_action: "按需或排除。"
```

---

## 7. 本补充 ledger 的复用规则

1. 官方 / 大媒体来源不是越多越好；常驻保留主入口和明确互补功能。
2. 电子报、新闻联播、热门文章这类主入口负责观察该媒体当天或横向选择，不再叠加大量栏目切片。
3. 政策、时政、地方、人事、知识产权、环保、法治等栏目适合具体问题按需核验，不形成每日未读队列。
4. 调查型节目只有补出新结构功能时才扩容；`新闻调查` 补社会结构现场，`新闻1+1` 虽高价值但日频，按需。
5. 同一媒体内部评论、理论、人物、文化、健康、地方和行业栏目必须看重复、更新、标签、负荷和已有来源覆盖。
6. 官方宽泛词不进追踪器；追踪器只收具体政策、地区事件、职位任免、法律、产品、公司、项目或有结束条件的事件。

---

## 8. 仍需继续追回

综合大类已经补入：

- closeout 三项动作；
- 知乎小标签；
- 热搜小标签；
- 滚动新闻小标签；
- 微信 / 微博 / 今日头条；
- 百度 / 腾讯 / 网易 / 搜狐 / 新浪 / ZAKER；
- 人民网 / 新华社 / 南方周末 / 半月谈 / CCTV。

仍需继续追回：

- 健康 / 军事 / 房产 / 地方城市；
- 腾讯日报 11 节点细账；
- `腾讯研究院` 与 `网易人间｜记事` 的 Folo 复查条件；
- 综合目录中如仍有 directory review / rolling review 的未落账边界。

综合大类仍是 `partial_recovered_from_pr`，不能写成全量完成。
