# PR #372 判断链追回：社区核心与边界来源（2026-07-06）

## 0. 边界

本文件是 `社区` 大类的补充 ledger，追回 PR #372 中已闭环的社区同族来源判断。

纳入文件：

- `docs/ops/tophub-community-security-closure-20260705.md`
- `docs/ops/tophub-community-v2ex-closure-20260705.md`
- `docs/ops/tophub-community-travel-closure-20260705.md`
- `docs/ops/tophub-community-digital-closure-20260705.md`
- `docs/ops/tophub-community-developer-communities-closure-20260705.md`
- `docs/ops/tophub-community-campus-forums-closure-20260705.md`
- `docs/ops/tophub-community-auto-forums-closure-20260705.md`
- `docs/ops/tophub-community-local-portals-closure-20260705.md`
- `docs/ops/tophub-community-linuxdo-empty-closure-20260705.md`

当前不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 社区来源总规则

```yaml
- source_id: "社区来源总规则"
  current_route: keep_few_low_noise_entry_points_and_task_sources
  decision_status: recovered_from_pr
  selected_reason: "社区目录中有真实经验、故障、路线、工具、职业和地方生活信号；但多数论坛热榜、平台榜单和地方门户同时混入广告、交易、情绪、旧帖、推广和不可核验叙述。"
  rejected_alternatives:
    - "把论坛热榜当作公共温度常驻"
    - "把多个社区高频榜单加入 Folo"
    - "把所有地方门户、校园论坛、汽车论坛变成长期入口"
    - "把平台聚合当作作者订阅关系"
  route_reason: "TopHub 只保留少数低噪声、互补、能提供稳定社区天气的入口；Folo 只保留作者或编辑关系，不保留论坛聚合；具体任务按需调用社区经验。"
  reuse_rule: "以后遇到社区来源，先拆成：低频周报、精华精选、论坛高频榜、地方生活任务、设备经验、旅行路线、作者连续关系。不同层不能互相替代。"
  evidence_locator: "community security / v2ex / travel / digital / developer communities / campus / auto / local portals / linuxdo closure files"
  next_action: "作为社区同族复用规则。"
```

---

## 2. 三个新增动作：先知、V2EX 周报、马蜂窝

```yaml
- source_id: "先知社区｜精华推荐"
  current_route: tophub_technology_radar
  decision_status: recovered_from_pr
  selected_reason: "补当前科技雷达缺少的应用安全、漏洞研究与 AI Agent 安全入口。它覆盖 Web 与云安全、容器、供应链、移动应用、漏洞复现、AI 编码工具、MCP、提示注入和 Agent 工具边界。"
  rejected_alternatives:
    - "看雪论坛｜最新精华"
    - "先知社区无副标题重复节点"
    - "安全论坛全部进入 Folo"
  route_reason: "进入科技雷达，放在 The Register 之后、科普中国网之前。`精华推荐` 比普通最新流更低噪声，不要求逐篇阅读；看雪更窄且包含逆向、内核、移动端和灰度主题，按任务访问；重复节点排除。"
  reuse_rule: "安全社区常驻优先低噪声精华入口；具体系统仍需核对授权、版本、影响范围、厂商公告、CVE、补丁说明和正式报告。"
  evidence_locator: "docs/ops/tophub-community-security-closure-20260705.md#最终决定"
  next_action: "已作为社区动作之一。"

- source_id: "V2EX 周报"
  current_route: tophub_life_and_community
  decision_status: recovered_from_pr
  selected_reason: "它以周频观察开发者、独立创作者和技术使用者一周内持续讨论了什么；同时保留主题和高赞回复，减少只凭标题判断的失真。"
  rejected_alternatives:
    - "V2EX 最热主题"
    - "V2EX 今日热议"
    - "V2EX 技术 / 程序员 / 互联网 / 问与答"
    - "水木社区十大热门话题"
    - "V2EX 高频产品发现板块"
  route_reason: "进入生活与社区，而不是科技雷达。科技雷达负责新闻、产品变化、科学与安全；V2EX 周报补的是具体使用者的经验、故障、选择和工作方式。高频 V2EX 板块推广、账号渠道、中转站和灰色服务较多，转按需。"
  reuse_rule: "技术社区日榜高频噪声大；若要常驻，优先周报或精选，而不是最新、热议、交易和问答流。"
  evidence_locator: "docs/ops/tophub-community-v2ex-closure-20260705.md#最终决定; #为什么选择周报而不是热门或技术节点"
  next_action: "已作为社区动作之一。"

- source_id: "马蜂窝｜热门游记"
  current_route: tophub_life_and_community
  decision_status: recovered_from_pr
  selected_reason: "当前七组缺少旅行经验入口；热门游记包含国内徒步、亲子自驾、国际铁路旅行、澳大利亚、东南亚和区域环线，内容以完整游记和路线经验为主。"
  rejected_alternatives:
    - "携程攻略｜推荐游记"
    - "微信旅行 24h 热文榜"
    - "背包客栈自助旅行论坛｜背包严选好文"
    - "飞客茶馆｜最新热门"
    - "人民网旅游频道"
    - "21财经文旅"
  route_reason: "进入生活与社区，补的是‘人在具体地点怎样走、怎样安排和怎样体验’。背包客栈质量接近但与马蜂窝功能重叠，作为国际旅行高价值按需；飞客偏酒店积分和高消费权益；其他来源陈旧、污染或偏产业。"
  reuse_rule: "旅行经验源负责灵感和路线，不作为最新政策、价格或安全事实；进入计划后回签证、交通、景区、天气和官方安全信息核验。"
  evidence_locator: "docs/ops/tophub-community-travel-closure-20260705.md#最终决定"
  next_action: "已作为社区动作之一。"
```

---

## 3. 高价值按需：看雪、V2EX 任务板块、背包客栈、数码专业源

```yaml
- source_id: "看雪论坛｜最新精华"
  current_route: on_demand_reverse_engineering_security_source
  decision_status: recovered_from_pr
  selected_reason: "看雪在逆向工程、内核、移动端、Hook、虚拟化和反作弊领域有独特深度，与用户过去 Android 逆向经历相关。"
  rejected_alternatives:
    - "与先知社区同时常驻"
    - "进入 Folo"
  route_reason: "当前首页需要安全变化雷达，而不是持续深入逆向工程。看雪内容更窄，且夹杂游戏辅助、绕过、去广告等灰度主题；当用户重新进入 Android、内核、恶意样本或深度逆向项目时按需访问。"
  reuse_rule: "窄安全论坛按具体逆向 / 移动安全 / 内核任务启用；不和综合安全精华同时常驻。"
  evidence_locator: "docs/ops/tophub-community-security-closure-20260705.md#暂不新增看雪论坛最新精华"
  next_action: "按需。"

- source_id: "V2EX 任务板块"
  current_route: on_demand_practical_community_knowledge
  decision_status: recovered_from_pr
  selected_reason: "Apple、Android、硬件、DNS、信用卡、汽车、Blog 等板块在具体任务中很有用，可提供设备、网络、跨境支付、订阅、用车和独立博客发现经验。"
  rejected_alternatives:
    - "常驻多个 V2EX 分板块"
    - "把 V2EX Blog 直接当作者订阅"
    - "交易板块进入系统"
  route_reason: "这些板块高频且混入推广、账号、灰色服务和群体偏差。V2EX Blog 只能发现具体作者，只有确认作者身份、长期更新质量和不可替代性后才进入 Folo 复查。交易板块风险与时效高，排除。"
  reuse_rule: "论坛板块按具体任务打开；社区经验不能替代官方政策、银行规则、设备日志、召回、投诉、长期测试或开发者文档。"
  evidence_locator: "docs/ops/tophub-community-v2ex-closure-20260705.md#其他节点的处理"
  next_action: "按需。"

- source_id: "背包客栈自助旅行论坛｜背包严选好文"
  current_route: on_demand_international_travel_planning_source
  decision_status: recovered_from_pr
  selected_reason: "人工严选感更强，行前准备、交通、保险、装备和路线细节更具体，国际旅行与东亚旅行经验较丰富。"
  rejected_alternatives:
    - "与马蜂窝同时常驻"
    - "进入 Folo"
  route_reason: "与马蜂窝功能重叠，当前只需要一个日常旅行发现入口。若未来国际旅行成为连续项目，背包客栈是第一补充来源。"
  reuse_rule: "同类旅行经验源不叠加；国内 / 综合游记常驻一条，国际深度规划按任务补。"
  evidence_locator: "docs/ops/tophub-community-travel-closure-20260705.md#高价值按需来源"
  next_action: "按需。"

- source_id: "数码专业任务源"
  current_route: on_demand_hardware_and_device_task_sources
  decision_status: recovered_from_pr
  selected_reason: "Puget Systems、TechPowerUp、Notebookcheck、UltrabookReview、Tom's Hardware、igor´sLAB、充电头网、数码之家、投影网等，在工作站、GPU、笔记本、充电协议、维修拆解、投影和设备故障任务中有真实价值。"
  rejected_alternatives:
    - "社区数码常驻扩容"
    - "Folo"
    - "用消费数码媒体全站热榜替代"
  route_reason: "当前常驻结构已由少数派、极客公园、TechCrunch、The Verge、IT之家、The Register、Counterpoint、Apple 支持覆盖持续观察层。专业源在购买、升级、维修、性能或兼容性任务中调用，不日常订阅。"
  reuse_rule: "数码源按任务分层：设备体验、专业评测、维修拆解、协议兼容、促销优惠、论坛经验。没有具体设备或购买任务就不常驻。"
  evidence_locator: "docs/ops/tophub-community-digital-closure-20260705.md#高价值按需来源; #结论"
  next_action: "按需。"
```

---

## 4. 作者 / 个人来源候选：观察，不立即订阅

```yaml
- source_id: "夜航船夫"
  current_route: folo_review_candidate_technology_in_real_life
  decision_status: recovered_from_pr
  selected_reason: "文章涉及 Obsidian、OpenCode、Daily Note、iCloud 与 Google 日历、Rime、NAS、本地生成图片、项目投产与个人生活记录，具备‘技术进入真实生活’的作者连续性。"
  rejected_alternatives:
    - "立即进入 Folo"
    - "进入 TopHub 标题流"
  route_reason: "只记录为 2026-07-17 复查观察对象，不改变 Folo 27 项。需要确认持续更新、是否以真实实践为主、是否与现有作者源形成明显增量。"
  reuse_rule: "作者型社区 / 数码来源看长期实践关系，不看单篇命中；先观察更新、重复、阅读率和主题稳定性。"
  evidence_locator: "docs/ops/tophub-community-digital-closure-20260705.md#可在Folo复查时观察的个人来源"
  next_action: "Folo 复查。"

- source_id: "Slot 4"
  current_route: folo_review_conditional_device_practice_candidate
  decision_status: recovered_from_pr
  selected_reason: "内容包括自组 PC、复古掌机、迷你主机、Apple 自动化、Switch 维修、2.5G 内网和具体设备折腾，实践性较强。"
  rejected_alternatives:
    - "立即进入 Folo"
    - "进入 TopHub"
  route_reason: "同样只进入观察名单。它是设备实践候选，不是公共科技雷达；是否进入 Folo 要看长期更新、真实阅读和与夜航船夫、少数派、小众软件等来源的重复。"
  reuse_rule: "设备实践作者可作为 Folo 候选，但不能和论坛热榜或消费数码媒体混为一类。"
  evidence_locator: "docs/ops/tophub-community-digital-closure-20260705.md#可在Folo复查时观察的个人来源"
  next_action: "Folo 复查。"

- source_id: "猫和柴的野游"
  current_route: on_demand_north_america_outdoor_travel_author
  decision_status: recovered_from_pr
  selected_reason: "作者经验具体，集中于美国西北、加拿大、夏威夷、徒步和信用卡。"
  rejected_alternatives:
    - "Folo 候选"
    - "替代马蜂窝"
  route_reason: "作者属性成立，但地域和主题过窄，当前长期必要性不足；北美西岸、户外徒步或相关信用卡问题时定向查看。"
  reuse_rule: "旅行作者源必须看地域、主题、更新和长期任务关系；过窄则按需。"
  evidence_locator: "docs/ops/tophub-community-travel-closure-20260705.md#Folo判断; #猫和柴的野游"
  next_action: "按需。"
```

---

## 5. 论坛与地方任务源：按现实关系触发

```yaml
- source_id: "高校论坛"
  current_route: no_new_on_demand_campus_sources
  decision_status: recovered_from_pr
  selected_reason: "水木、北大未名、北邮人能提供高校背景职业人群、校园生活、升学、校招、保研、城市和技术消费讨论。"
  rejected_alternatives:
    - "水木社区十大热门话题作为泛职业社区入口"
    - "北大未名首页推荐或全站热门常驻"
    - "北邮人十大热门常驻"
    - "Folo"
  route_reason: "不新增任何高校论坛。V2EX 周报胜出，因为周频、保留回复、覆盖技术/独立开发/职业/生活经验，且更贴近用户当前 AI 编程、GitHub、Mac、网络、支付与跨境服务问题。高校论坛依赖校内黑话、地域背景和具体校园事务。"
  reuse_rule: "高校论坛只在具体学校、招生、求职、校友、课程或校园事务出现时使用；泛职业社区入口优先低频周报。"
  evidence_locator: "docs/ops/tophub-community-campus-forums-closure-20260705.md#最终决定; #与V2EX周报的比较"
  next_action: "按需。"

- source_id: "汽车论坛"
  current_route: on_demand_vehicle_purchase_and_owner_experience
  decision_status: recovered_from_pr
  selected_reason: "汽车之家、虎扑汽车、V2EX 汽车等能提供提车、用车、预算、二手车、保险、续航、电费、异响、保养和真实用车体验。"
  rejected_alternatives:
    - "汽车论坛常驻"
    - "虎扑汽车作为汽车公共温度"
    - "V2EX 汽车常驻"
    - "Folo"
  route_reason: "不新增任何汽车论坛。当前系统已通过第一财经汽车新闻与 Counterpoint 观察汽车产业、市场和技术结构；论坛只有在选车、买车、养车、保险、维修或具体车型比较时才有明显增量。V2EX 汽车最有任务价值，但论坛样本有群体偏差，需要和召回、投诉、碰撞、保险、真实成交价和长期测试交叉验证。"
  reuse_rule: "汽车论坛是车主经验样本，不是事实入口。具体车型任务时使用，产业结构继续由结构源承担。"
  evidence_locator: "docs/ops/tophub-community-auto-forums-closure-20260705.md#最终决定; #V2EX汽车"
  next_action: "按需。"

- source_id: "地方门户"
  current_route: on_demand_local_life_and_public_service_sources
  decision_status: recovered_from_pr
  selected_reason: "地方门户在用户正在某地生活、旅行、租房、求学、就医、处理物业或公共服务事项时有价值，可提供交通、物业、住房、环境、学校、公共设施、本地生活经验和公共问题线索。"
  rejected_alternatives:
    - "新增地方门户常驻"
    - "建立独立地方门户分组"
    - "地方论坛进入 Folo"
  route_reason: "49 个节点全部审核后不新增。脱离具体地域后，这些榜单主要是广告、招聘、房产、装修、培训、婚恋、家庭争议、全国新闻重复、口径混乱、乱码或缺少最终状态。深圳论坛和高楼迷可按具体任务使用，但不常驻。"
  reuse_rule: "地方门户只在现实地域关系成立时启用。道路、物业、学位、医院、天气和政策问题，论坛只提供线索，最终核对当地政务、主管部门、气象、教育、交通和正规新闻来源。"
  evidence_locator: "docs/ops/tophub-community-local-portals-closure-20260705.md#最终决定; #路由规则"
  next_action: "按需。"
```

---

## 6. 开发者社区与空页面闭环

```yaml
- source_id: "开发者社区标签整体"
  current_route: no_extra_action_beyond_xianzhi
  decision_status: recovered_from_pr
  selected_reason: "该标签再次确认先知社区的价值，也暴露吾爱破解、链滴、独立开发者社区、失业者联盟、Ruby China、C-Sharpcorner、V2EX 今日热议等不同任务源。"
  rejected_alternatives:
    - "吾爱破解常驻"
    - "链滴热议常驻"
    - "独立开发者社区常驻"
    - "Ruby China / C-Sharpcorner 常驻"
    - "V2EX 今日热议常驻"
    - "Folo"
  route_reason: "本标签不产生除既定先知动作之外的新增。吾爱破解供应链与授权风险高；链滴更偏思源生态；独立开发者社区推广密度过高；Ruby China 当前性不足且无 Ruby/Rails 连续项目；C-Sharpcorner 质量跨度大；V2EX 今日热议已由周报替代。"
  reuse_rule: "开发者社区热榜按具体生态、语言、工具或职业任务使用；平台聚合不变成作者关系。"
  evidence_locator: "docs/ops/tophub-community-developer-communities-closure-20260705.md#最终决定; #逐节点判断"
  next_action: "无额外动作。"

- source_id: "LINUXDO 标签"
  current_route: empty_page_closed
  decision_status: recovered_from_pr
  selected_reason: "无可见节点。"
  rejected_alternatives:
    - "把空页面记成待审核"
    - "对 LINUX DO 社区本身补编质量判断"
  route_reason: "当前 TopHub `LINUXDO` 页面为空白，没有显示任何可见节点、榜单或可订阅项目。因此按空页面闭环，不推断社区本身内容质量。"
  reuse_rule: "空页面只能闭环页面事实，不评价站点本体；未来 TopHub 恢复节点后再按实际内容重审。"
  evidence_locator: "docs/ops/tophub-community-linuxdo-empty-closure-20260705.md#页面事实; #结论"
  next_action: "无。"
```

---

## 7. 明确不进 Folo 的规则

```yaml
- source_id: "社区聚合进入 Folo 的边界"
  current_route: no_folo_for_platform_aggregates
  decision_status: recovered_from_pr
  selected_reason: "Folo 应保存作者、编辑关系或稳定刊物关系，而不是论坛热榜、平台榜单和地方聚合。"
  rejected_alternatives:
    - "V2EX 周报进入 Folo"
    - "先知 / 看雪论坛榜单进入 Folo"
    - "马蜂窝 / 背包客栈榜单进入 Folo"
    - "汽车论坛、高校论坛、地方门户进入 Folo"
  route_reason: "这些来源是聚合流和任务入口，不是单一作者关系。只有夜航船夫、Slot 4、V2EX Blog 中发现的具体独立作者、或其他长期作者关系，才可进入 Folo 复查。"
  reuse_rule: "论坛热榜不进 Folo；从论坛中发现作者后，另按作者身份、更新质量、主题稳定、不可替代性和真实阅读率复查。"
  evidence_locator: "community closure files Folo 判断 sections"
  next_action: "Folo 复查只处理具体作者或稳定刊物。"
```

---

## 8. 本补充 ledger 的复用规则

1. 社区来源要先分层：精华精选、周报、论坛热榜、任务板块、地方门户、作者关系、空页面。
2. `先知社区｜精华推荐`补科技安全雷达，不等于订阅所有安全论坛。
3. `V2EX 周报`补技术使用者的一周社区天气，不等于订阅 V2EX 高频板块。
4. `马蜂窝｜热门游记`补旅行经验入口，不替代签证、交通、价格、天气和安全事实。
5. 高校论坛、汽车论坛、地方门户都按现实任务触发，不常驻。
6. 数码专业源按设备、购买、维修、性能、兼容性任务启用。
7. 社区经验不能替代官方公告、政策、银行规则、设备日志、长期测试、召回、CVE 或厂商说明。
8. 论坛热榜、平台榜单和地方门户不进 Folo；作者型来源另行复查。
9. 空页面只闭环页面事实，不补编站点判断。

---

## 9. 仍需继续追回

社区核心与边界已补入：安全、V2EX、旅游、数码、开发者社区、高校论坛、汽车论坛、地方门户、LINUXDO。

继续待办：

- 社区目录第 1 页与 pending actions 的边界如果需要，可再写一份收口 ledger；
- 等用户本地提交 master checklist 与 coverage audit 后，再统一登记本文件和此前新增 ledger；
- 财经逐页；
- 购物逐页；
- 政务、校务、专栏、浏览器与链接；
- 旧 ChatGPT 对话追索所有 partial 条目。
