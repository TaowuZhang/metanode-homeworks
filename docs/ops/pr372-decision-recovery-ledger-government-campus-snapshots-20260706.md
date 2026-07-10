# PR #372 判断链追回：政务 / 校务快照（2026-07-06）

## 0. 边界

本文件追回 PR #372 中 `政务` 与 `校务 / 校刊` 目录快照的判断链。

纳入文件：

- `docs/ops/tophub-government-affairs-snapshot-closure-20260706.md`
- `docs/ops/tophub-campus-affairs-snapshot-closure-20260706.md`

注意：

- `政务`基于用户提供的 `政务 1063 个`共 30 页快照，属于页级复审闭环；不声称逐条在线核验了 1063 个节点的实时更新状态。
- `校务 / 校刊`只依据用户提供的 `校务 352 个`页面快照；不声称已经逐页审完全部 352 个节点。
- 当前不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 政务目录总规则：上位政策门牌 + 任务工具池

```yaml
- source_id: "TopHub > 政务 1063 节点快照整体"
  current_route: snapshot_closed_with_one_candidate_no_platform_action
  decision_status: recovered_from_pr_snapshot
  selected_reason: "政务快照覆盖中央综合政务、宏观数据、财经监管、法律司法、外交使领馆、应急气象、教育考试、地方政务、辟谣、网信、行业协会、赛事、高校医院和若干误入目录的企业服务节点。目录工具价值很高。"
  rejected_alternatives:
    - "把政务目录整体作为日常信息流"
    - "把 1063 个政务节点写成全量实时核验完成"
    - "把多个部委、地方、使馆、法院、纪委、考试、赛事、协会节点加入 TopHub / Folo"
    - "把误入政务目录的企业服务节点当政务源"
  route_reason: "政务目录更像任务入口库，不适合整包进入日常注意力。最终只形成 1 个 TopHub 新增候选：中国政府网｜最新政策；本轮没有真实平台执行，因此当前真实 TopHub 和 Folo 配置不变。"
  reuse_rule: "以后遇到政务来源，先分层：上位政策门牌、部委任务入口、地方任务入口、驻外任务入口、考试招聘入口、产业监管入口、法律监管入口、政治机关宣传、纪委人事任免、协会竞赛、学校医院、误入目录企业服务。不同层不能互相替代。"
  evidence_locator: "docs/ops/tophub-government-affairs-snapshot-closure-20260706.md#范围; #页码清点; #结论"
  next_action: "若用户实际在平台添加中国政府网最新政策，再写入真实订阅账本。"
```

---

## 2. 中国政府网｜最新政策：为什么是唯一政务候选

```yaml
- source_id: "中国政府网｜最新政策"
  current_route: tophub_candidate_data_and_structure_not_yet_executed
  decision_status: recovered_from_pr_snapshot
  selected_reason: "它是全国性、上位、跨部门的官方政策入口，补当前 TopHub 已有统计局数据源之外的同级政策文本源。它承担规则、制度、行政安排和全国性政策变化发现。"
  rejected_alternatives:
    - "中国政府网｜政策解读"
    - "中国政府网｜图解政策"
    - "中国政府网｜政策说明书"
    - "国务院信息"
    - "国务院会议"
    - "部门动态"
    - "多个中国政府网栏目同时加入"
  route_reason: "建议放入数据与结构，位于国家统计局数据解读之后、财新首页推荐之前。相比解读、图解和说明书，最新政策更接近原始政策入口；相比国务院信息、会议和部门动态，它更像长期承担制度变化发现的门牌。本轮没有真实平台执行，所以只是候选，不写入真实订阅顺序。"
  reuse_rule: "政务常驻只收上位、原始、跨部门的门牌。解读、图解、会议、新闻和部门动态按具体政策任务打开，不多栏目并存。"
  evidence_locator: "docs/ops/tophub-government-affairs-snapshot-closure-20260706.md#是否加入TopHub今日热榜; #如果添加应该放到哪里排第几个"
  next_action: "等待用户实际添加后再更新真实账本。"
```

---

## 3. 为什么不选其他政务来源常驻

```yaml
- source_id: "部委、地方、监管、法律、应急、使领馆与考试节点"
  current_route: on_demand_government_tool_pool
  decision_status: recovered_from_pr_snapshot
  selected_reason: "这些节点在具体法律、金融、考试、签证、护照、旅行、安全、出入境、地方政务、产业政策、监管、气象、应急、学校申请和医疗任务中非常重要。"
  rejected_alternatives:
    - "外交部发言人流常驻"
    - "地震、气象、消防、应急、交通、核安全、生态环境突发类节点常驻"
    - "中央纪委、地方纪委监委、最高法、最高检、人大、司法部、证监会、交易所问询、央行处罚常驻"
    - "教育部、考试网、IELTS、高校研究生院、就业平台、人事考试常驻"
    - "驻外使馆与领馆通知常驻"
    - "行业协会、竞赛组织、公益组织、企业公告常驻"
  route_reason: "这些来源任务性强、地域 / 身份 / 事件依赖高，不适合日常未读流。重要国际事件可由公共新闻触发；突发预警应由专门预警渠道、天气工具、所在地通知或具体任务触发；法律和监管源按项目、案件、公司、规则或政策任务调用；使领馆只按所在地和近期出入境需求单点使用。"
  reuse_rule: "政务源不按重要性常驻，而按责任关系触发：我是否在该地区、是否涉及该证件 / 考试 / 政策 / 学校 / 监管对象 / 出行事项 / 公司事项。"
  evidence_locator: "docs/ops/tophub-government-affairs-snapshot-closure-20260706.md#不建议加入的类型; #按需政务工具池"
  next_action: "按任务调用。"

- source_id: "误入政务目录的企业与服务节点"
  current_route: route_back_to_product_or_service_tasks
  decision_status: recovered_from_pr_snapshot
  selected_reason: "阿里云、Apple、Toyota、饿百零售开放平台等可能在产品、技术、汽车、维修或服务任务中有价值。"
  rejected_alternatives:
    - "因出现在政务目录中而进入政务订阅"
  route_reason: "它们不属于政务本身。若需要跟踪，应回到对应产品、技术、汽车或服务任务，不作为政务节点处理。"
  reuse_rule: "目录归属不能替代来源本体。误入目录的企业服务源必须回真实领域重审。"
  evidence_locator: "docs/ops/tophub-government-affairs-snapshot-closure-20260706.md#按需政务工具池"
  next_action: "按任务。"
```

---

## 4. 政务为什么不进 Folo

```yaml
- source_id: "政务 Folo 边界"
  current_route: no_folo_for_government_broad_streams
  decision_status: recovered_from_pr_snapshot
  selected_reason: "Folo 当前承担长期作者、专业来源、节目和确实会持续跟随的来源。"
  rejected_alternatives:
    - "把政务节点加入 Folo"
    - "把多个使馆通知加入 Folo"
    - "把考试、地方政务、部委通知加入 Folo"
  route_reason: "政务节点数量巨大，进入 Folo 会形成高未读压力。政务类最好按任务、地区、机构和事件触发；只有出现具体任务时，才可临时单点加入，例如某项考试、某项政策征求意见、某项基金通知、某个监管规则、某个地方事项、某个使领馆通知、某项赛事或某个学校医院公告。"
  reuse_rule: "官方通知监控不是 Folo 的默认职责。只有具体任务成立，才单点临时订阅或建立提醒。"
  evidence_locator: "docs/ops/tophub-government-affairs-snapshot-closure-20260706.md#是否加入FoloFollow"
  next_action: "不新增 Folo 复查候选。"
```

---

## 5. 校务 / 校刊目录：任务入口，不是公共社区常驻

```yaml
- source_id: "TopHub > 校务 / 校刊 352 节点快照整体"
  current_route: snapshot_review_no_tophub_no_folo
  decision_status: recovered_from_pr_snapshot
  selected_reason: "快照包含校园论坛热榜、高校与院系官方节点、招生、教务、学院通知、学术活动和教育局通知。它有工具价值。"
  rejected_alternatives:
    - "把校务目录加入 TopHub 常驻"
    - "把校园论坛热榜作为公共温度或生活与社区常驻"
    - "把官方校务节点加入 Folo"
    - "把快照写成已逐页审完全部 352 个节点"
  route_reason: "本记录只依据用户提供快照，不声称逐页审完 352 个节点。校园论坛热榜主要是校内情绪、征友、二手、日记、求助、offer、实习、食堂、住宿、社团和闲聊；官方校务节点高度依赖具体身份、学校、学院、考试年份和申请任务；快照中北大未名日期显示为 2026-05-08，说明时效性需要逐源核验。"
  reuse_rule: "校务源按具体任务调用：保研、考研、夏令营、开放日、教务、考试、培养方案、实习就业、校园共同体观察、人物 / 学院 / 科研团队线索。没有身份或任务，不常驻。"
  evidence_locator: "docs/ops/tophub-campus-affairs-snapshot-closure-20260706.md#范围; #是否加入TopHub今日热榜; #按需使用规则"
  next_action: "无真实配置动作。"
```

---

## 6. 为什么不选水木 / 未名 / 北邮人常驻

```yaml
- source_id: "校园论坛热榜"
  current_route: on_demand_campus_community_temperature_not_current_subscription
  decision_status: recovered_from_pr_snapshot
  selected_reason: "水木、北大未名、北邮人能在理解某个高校共同体、校内生活、求职、实习、offer、宿舍、食堂、二手和校园情绪时提供局部温度。"
  rejected_alternatives:
    - "水木社区十大热门话题当前加入生活与社区"
    - "北大未名全站热门话题常驻"
    - "北邮人论坛十大热门话题常驻"
    - "校园论坛进入 Folo"
  route_reason: "当前不加入。校园论坛价值更像特定高校共同体的局部温度，不是当前公共温度或生活与社区的必备广谱节点；已由豆瓣话题广场、豆瓣小组、V2EX 周报、知乎和小红书等公共社区源覆盖一部分。"
  reuse_rule: "校园论坛只在需要理解某个高校共同体时临时查看。它不能代表社会整体，也不能替代正式招生、教务、就业和学院通知。"
  evidence_locator: "docs/ops/tophub-campus-affairs-snapshot-closure-20260706.md#是否加入TopHub今日热榜"
  next_action: "按需；未来最多短期试验。"

- source_id: "水木社区｜十大热门话题 校务试验位"
  current_route: future_task_triggered_trial_only
  decision_status: recovered_from_pr_snapshot
  selected_reason: "若未来出现明确的校园共同体观察任务，水木可以作为短期试验位。"
  rejected_alternatives:
    - "本轮真实添加"
    - "作为长期生活与社区节点"
  route_reason: "文件只给未来试验方案：放生活与社区第 7 位，位于 V2EX 周报之后、马蜂窝之前；若连续几次只输出校内私人生活、二手、征友、情绪和职场闲聊，不能补出真实公共判断，就退出。"
  reuse_rule: "试验位必须有退出条件。弱社区候选不能没有时间边界地变成长期订阅。"
  evidence_locator: "docs/ops/tophub-campus-affairs-snapshot-closure-20260706.md#如果添加应该放到哪里排第几个"
  next_action: "本轮不执行。"
```

---

## 7. 校务为什么不进 Folo

```yaml
- source_id: "校务 Folo 边界"
  current_route: no_folo_for_broad_campus_monitoring
  decision_status: recovered_from_pr_snapshot
  selected_reason: "Folo 当前角色是少量长期作者、专业来源、多媒体来源和真正会持续跟随的节目。"
  rejected_alternatives:
    - "校务通知进入 Folo"
    - "高校论坛进入 Folo"
    - "招生、教务、学院公告、就业平台和教育局通知加入 Folo"
  route_reason: "校务通知会形成大量未读负担，而且大部分与用户当前身份和任务无关。只有出现具体任务时，才考虑单点订阅，例如某一年某校研究生招生、某学院通知、某个校园开放日或某个考试安排。"
  reuse_rule: "校务只在具体学校、学院、年份、考试、申请、岗位或项目成立时单点跟踪；不做宽泛校园监控。"
  evidence_locator: "docs/ops/tophub-campus-affairs-snapshot-closure-20260706.md#是否加入FoloFollow"
  next_action: "不新增 Folo 复查候选。"
```

---

## 8. 本补充 ledger 的复用规则

1. 政务是工具价值高、常驻价值窄的目录；不按覆盖率加入。
2. 政务常驻只考虑上位、原始、跨部门门牌；本轮唯一候选是 `中国政府网｜最新政策`，但未真实执行。
3. 部委、地方、使领馆、法律监管、考试、赛事、协会、医院和学校节点按身份、地区、机构、事项和事件触发。
4. 误入政务目录的企业服务节点回产品 / 技术 / 汽车 / 服务任务重审。
5. 政务不进 Folo；具体任务成立时可单点临时订阅或提醒。
6. 校务 / 校刊只依据快照，不写成 352 个节点逐页完成。
7. 校园论坛是局部共同体温度，不是公共温度常驻；水木只保留未来短期试验位与退出条件。
8. 校务通知高度依赖学校、学院、年份、考试、申请、就业和身份关系；无具体任务不常驻。
9. snapshot 只能承认快照范围与证据边界，不能冒充实时在线核验。

---

## 9. 仍需继续追回

政务 / 校务 snapshot 已补成 ledger。

继续待办：

- 等用户本地提交 master checklist 与 coverage audit 后，再统一登记本文件和此前新增 ledger；
- 专栏 94 页；
- 浏览器与链接；
- 旧 ChatGPT 对话追索所有 partial 条目。
