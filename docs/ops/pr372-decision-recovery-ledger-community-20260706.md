# PR #372 判断链追回：社区大类（第一版，2026-07-06）

## 0. 边界

本文件追回 `TopHub 社区` 大类中已经能从 PR #372 文件直接确认的判断链。

已读取并纳入本版：

- `docs/ops/tophub-community-pending-actions-20260705.md`

本版尚未覆盖社区同族 closure 的全部细节。以下文件仍需继续追回：

- `tophub-community-popular-page-01-20260705.md`
- `tophub-community-security-closure-20260705.md`
- `tophub-community-v2ex-closure-20260705.md`
- `tophub-community-travel-closure-20260705.md`
- `tophub-community-digital-closure-20260705.md`
- `tophub-community-developer-communities-closure-20260705.md`
- `tophub-community-campus-forums-closure-20260705.md`
- `tophub-community-auto-forums-closure-20260705.md`
- `tophub-community-local-portals-closure-20260705.md`
- `tophub-community-linuxdo-empty-closure-20260705.md`
- 虎扑相关判断以娱乐目录虎扑 closure 为准，社区目录不重复建立第二份判断。

因此本文件状态是：`partial_recovered_from_pr`。

特别边界：社区目录 235 个节点没有整体闭环。社区热门总目录第 2—20 页仍是缺口，不能写成 235/235 完整判断。

---

## 1. 社区大类已完成动作

来源证据：`docs/ops/tophub-community-pending-actions-20260705.md`

社区阶段已完成：

- `先知社区｜精华推荐` 加入 `科技雷达`；
- `V2EX 周报` 加入 `生活与社区`；
- `马蜂窝｜热门游记` 加入 `生活与社区`；
- 相关分组排序与全部页位置已完成；
- Folo 保持 27，不变。

```yaml
- source_id: "先知社区｜精华推荐"
  current_route: tophub_technology_radar
  decision_status: recovered_from_pr
  selected_reason: "补足应用安全、漏洞研究、AI Agent、MCP、提示注入和工具信任边界雷达。"
  rejected_alternatives:
    - "先知社区普通节点"
    - "看雪论坛｜最新精华作为同时常驻"
  route_reason: "进入科技雷达第 7 位；普通 `先知社区` 与精华推荐重复，看雪保留为逆向工程、内核和移动安全项目的按需来源。"
  ordering_reason: "科技雷达中位于 The Register 之后、科普中国网头条之前；全部页位置第 25。"
  evidence_locator: "docs/ops/tophub-community-pending-actions-20260705.md#已完成动作-1-先知社区精华推荐"
  next_action: "回 `tophub-community-security-closure-20260705.md` 补完整安全社区比较。"

- source_id: "V2EX 周报"
  current_route: tophub_life_and_community
  decision_status: recovered_from_pr
  selected_reason: "以周频观察开发者、独立创作者与技术使用者持续讨论的问题，避免每日热门、推广和中转站信息形成未读堆积。"
  rejected_alternatives:
    - "V2EX 最热主题"
    - "V2EX 今日热议"
    - "V2EX 首页最新"
    - "V2EX 技术"
    - "V2EX 程序员"
    - "两个重复的分享创造卡片"
    - "交易"
    - "站务"
    - "单个账号历史主题聚合"
  route_reason: "进入生活与社区第 6 位，作为泛技术职业社区的周观察入口，而不是高频社区流。"
  ordering_reason: "放在豆瓣小组讨论精选之后、马蜂窝热门游记之前；全部页位置第 61。"
  evidence_locator: "docs/ops/tophub-community-pending-actions-20260705.md#已完成动作-2-v2ex-周报"
  next_action: "回 `tophub-community-v2ex-closure-20260705.md` 补 V2EX 全节点细节。"

- source_id: "马蜂窝｜热门游记"
  current_route: tophub_life_and_community
  decision_status: recovered_from_pr
  selected_reason: "补足完整游记、路线、自驾、徒步和旅行现场经验。"
  rejected_alternatives:
    - "背包客栈自助旅行论坛｜背包严选好文 同时加入"
  route_reason: "进入生活与社区第 7 位，作为旅行经验入口；背包客栈与其功能部分重叠，本轮不同时加入，若国际旅行成为连续项目则作为第一补充来源。"
  ordering_reason: "放在 V2EX 周报之后、国家留学网综合项目专栏之前；全部页位置第 62。"
  evidence_locator: "docs/ops/tophub-community-pending-actions-20260705.md#已完成动作-3-马蜂窝热门游记; #背包客栈自助旅行论坛背包严选好文"
  next_action: "回 `tophub-community-travel-closure-20260705.md` 补旅行类比较。"
```

---

## 2. 跨目录复用与空页闭环

```yaml
- source_id: "虎扑"
  current_route: reuse_entertainment_decision_on_demand
  decision_status: recovered_from_pr
  selected_reason: "社区目录中的虎扑入口与娱乐目录下已审核的虎扑节点重叠。"
  rejected_alternatives:
    - "在社区目录重复建立第二份独立判断"
    - "新增 TopHub"
    - "进入 Folo"
  route_reason: "虎扑 16 个节点已在娱乐目录完整审核；结论仍是不新增、不替换、不进入 Folo，只在明确赛事、车型、装备或社区情绪任务中按需查看。"
  evidence_locator: "docs/ops/tophub-community-pending-actions-20260705.md#跨目录复用与空页闭环-虎扑"
  next_action: "若需要细节，回娱乐目录虎扑 closure；社区目录不重复判断。"

- source_id: "LINUXDO"
  current_route: empty_page_closed
  decision_status: recovered_from_pr
  selected_reason: "用户实际打开社区目录下的 LINUXDO 页面后，页面为空白，当前可见节点为 0。"
  rejected_alternatives:
    - "新增 TopHub"
    - "取消已有节点"
    - "进入 Folo"
  route_reason: "没有榜单或订阅项可供判断，按空页面闭环；未来页面若重新出现节点，只审核届时新增的实际内容。"
  evidence_locator: "docs/ops/tophub-community-pending-actions-20260705.md#跨目录复用与空页闭环-linuxdo"
  next_action: "无；未来有节点时再审。"
```

---

## 3. 已关闭候选

```yaml
- source_id: "水木社区｜十大热门话题"
  current_route: rejected
  decision_status: recovered_from_pr
  selected_reason: "可作为高校 / 老牌社区话题入口，但最终不添加。"
  rejected_alternatives:
    - "加入生活与社区"
    - "替代 V2EX 周报"
  route_reason: "完整比较高校论坛、开发者社区与 V2EX 后，水木日频筛选弱，标题经常截断，高校与社区黑话较多，情感、体育和八卦混入；V2EX 周报以周频聚合热门主题与高赞回复，更适合承担泛技术职业社区入口。"
  evidence_locator: "docs/ops/tophub-community-pending-actions-20260705.md#已关闭的候选-水木社区十大热门话题"
  next_action: "无；不添加。"
```

---

## 4. 高价值按需来源

```yaml
- source_id: "V2EX 专题节点"
  current_route: on_demand_technical_and_life_threads
  decision_status: recovered_from_pr
  selected_reason: "部分专题节点在具体任务中有高价值，例如 Apple、Android、硬件、DNS、信用卡、汽车、Blog。"
  rejected_alternatives:
    - "批量转入 TopHub"
    - "批量转入 Folo"
  route_reason: "这些节点只在 Mac、iPhone、Apple ID、Root、Termux、eSIM、装机、DNS、跨境支付、购车、独立博客发现等具体问题出现时使用；信用卡等结论仍需核对银行官方政策。"
  evidence_locator: "docs/ops/tophub-community-pending-actions-20260705.md#高价值按需来源-v2ex-专题节点"
  next_action: "按任务调用。"

- source_id: "Puget Systems / TechPowerUp / Notebookcheck / UltrabookReview / Tom's Hardware / igor´sLAB / 充电头网 / 数码之家 / 投影网 / How-To Geek / Daily Camera News / Eurogamer"
  current_route: on_demand_digital_task_sources
  decision_status: recovered_from_pr
  selected_reason: "这些数码 / 硬件来源在本地 AI、视频、摄影后期、GPU 渲染、工作站性能、显卡、主板、电源、散热、显示器、笔记本、USB-C、维修、投影、家庭服务器、相机、游戏平台政策等具体任务中有价值。"
  rejected_alternatives:
    - "进入常驻信息流"
  route_reason: "只在具体任务出现时使用，不进入常驻信息流。"
  evidence_locator: "docs/ops/tophub-community-pending-actions-20260705.md#数码任务入口"
  next_action: "回 `tophub-community-digital-closure-20260705.md` 补逐项比较。"

- source_id: "背包客栈自助旅行论坛｜背包严选好文"
  current_route: on_demand_or_future_supplement
  decision_status: recovered_from_pr
  selected_reason: "国际旅行、行前准备、交通、保险和装备信息较具体。"
  rejected_alternatives:
    - "与马蜂窝热门游记同时加入"
  route_reason: "与马蜂窝功能部分重叠，本轮不同时加入；若国际旅行成为连续项目，它是第一补充来源。"
  evidence_locator: "docs/ops/tophub-community-pending-actions-20260705.md#背包客栈自助旅行论坛背包严选好文"
  next_action: "国际旅行成为连续项目时复查。"

- source_id: "看雪论坛｜最新精华"
  current_route: on_demand_reverse_engineering_source
  decision_status: recovered_from_pr
  selected_reason: "逆向工程、内核、移动端、Hook 和虚拟化深度较高。"
  rejected_alternatives:
    - "与先知社区精华推荐同时常驻"
  route_reason: "当前安全雷达先只加入覆盖更广的先知；用户重新进入 Android、内核或深度逆向项目时按需使用。"
  evidence_locator: "docs/ops/tophub-community-pending-actions-20260705.md#看雪论坛最新精华"
  next_action: "安全 / 逆向任务出现时调用。"

- source_id: "V2EX｜汽车"
  current_route: on_demand_auto_experience_source
  decision_status: recovered_from_pr
  selected_reason: "汽车论坛中最有实际经验价值的节点，适合购车、二手车、保险、续航和具体车型研究。"
  rejected_alternatives:
    - "常驻"
  route_reason: "当前没有明确购车任务，不常驻。"
  evidence_locator: "docs/ops/tophub-community-pending-actions-20260705.md#v2ex汽车"
  next_action: "购车或车型研究时调用。"

- source_id: "深圳论坛｜本周热帖 / 最新热门"
  current_route: on_demand_local_signal
  decision_status: recovered_from_pr
  selected_reason: "可作为深圳交通、物业、住房、学校和公共设施问题的临时传感器。"
  rejected_alternatives:
    - "常驻公共温度"
  route_reason: "论坛只提供线索，实际判断仍需核对主管部门与正式回应。"
  evidence_locator: "docs/ops/tophub-community-pending-actions-20260705.md#深圳论坛本周热帖--最新热门"
  next_action: "有深圳本地问题时调用。"

- source_id: "高楼迷｜今日热帖"
  current_route: on_demand_city_development_source
  decision_status: recovered_from_pr
  selected_reason: "研究具体商业项目、城市开发或铁路规划时可定向搜索。"
  rejected_alternatives:
    - "常驻"
  route_reason: "只作为城市开发 / 规划任务入口，不进入持续信息流。"
  evidence_locator: "docs/ops/tophub-community-pending-actions-20260705.md#高楼迷今日热帖"
  next_action: "具体项目研究时调用。"
```

---

## 5. Folo 复查观察

```yaml
- source_id: "夜航船夫"
  current_route: folo_review_observation_2026-07-17
  decision_status: recovered_from_pr_partial
  selected_reason: "Obsidian、OpenCode、日历、输入法、NAS、本地图片生成与生活实践相关。"
  rejected_alternatives:
    - "立即加入 Folo"
  route_reason: "Folo 保持 27 项不变，到 2026-07-17 再复查。"
  evidence_locator: "docs/ops/tophub-community-pending-actions-20260705.md#数码个人来源观察"
  next_action: "2026-07-17 复查；需回旧对话或数字 closure 补完整理由。"

- source_id: "Slot 4"
  current_route: folo_review_observation_2026-07-17
  decision_status: recovered_from_pr_partial
  selected_reason: "自组 PC、掌机、迷你主机、Apple 自动化、Switch 维修和家庭网络相关。"
  rejected_alternatives:
    - "立即加入 Folo"
  route_reason: "Folo 保持 27 项不变，到 2026-07-17 再复查。"
  evidence_locator: "docs/ops/tophub-community-pending-actions-20260705.md#数码个人来源观察"
  next_action: "2026-07-17 复查；需回旧对话或数字 closure 补完整理由。"
```

---

## 6. 排序与系统聚合节点

```yaml
- source_id: "专栏｜订阅聚合"
  current_route: system_aggregate_node_not_content_group
  decision_status: recovered_from_pr
  selected_reason: "系统级总览入口。"
  rejected_alternatives:
    - "进入七个用途分组"
  route_reason: "不进入七个用途分组，保留在全部页第 84，作为系统级总览入口。"
  evidence_locator: "docs/ops/tophub-community-pending-actions-20260705.md#全部页结果"
  next_action: "后续专栏目录单独追回。"

- source_id: "社区生活与社区排序"
  current_route: tophub_group_ordering_logic
  decision_status: recovered_from_pr
  selected_reason: "排序逻辑为：日常实用 → 健康核验 → 公共生活讨论 → 技术社区周观察 → 旅行经验 → 留学与潮流专题。"
  rejected_alternatives:
    - "把 V2EX 或旅行放在日常实用前"
  route_reason: "V2EX 周报放在豆瓣小组后，马蜂窝放在 V2EX 后，两者均位于国家留学网前。"
  evidence_locator: "docs/ops/tophub-community-pending-actions-20260705.md#生活与社区最终顺序; #全部页结果"
  next_action: "无。"
```

---

## 7. 未完成边界

```yaml
- source_id: "社区目录整体"
  current_route: partial_review_not_full_closure
  decision_status: partial
  selected_reason: "地方门户、安全社区、汽车论坛、旅游、数码、开发者社区、高校论坛、V2EX、虎扑、LINUXDO 等已完成或复用闭环。"
  rejected_alternatives:
    - "声称社区 235 个节点已经全量闭环"
  route_reason: "社区目录 235 个节点的全量审核仍未整体闭环；明确缺口是社区热门总目录第 2—20 页，以及尚未提供或尚未形成逐项记录的其他实际页面内容。"
  evidence_locator: "docs/ops/tophub-community-pending-actions-20260705.md#状态"
  next_action: "继续读取社区同族 closure；不得写成 235/235 完整判断。"
```

---

## 8. 本版仍需继续追回

本文件只完成社区 pending actions 的第一版判断链。

仍需继续追回：

- 安全社区 closure 中先知 / 看雪 / 其他安全源的逐项比较；
- V2EX closure 中 30 个节点的逐项判断；
- travel closure 中马蜂窝 / 背包客栈 / 其他旅行源的比较；
- digital closure 中数码任务入口的逐项判断；
- developer communities、campus forums、auto forums、local portals 的逐项判断；
- LINUXDO 空页 closure 的原始记录；
- 社区热门第 2—20 页缺口是否后续需要继续追索。

完成这些以后，才能把社区大类标为 `recovered_from_pr`；目前只能标为 `partial_recovered_from_pr`。
