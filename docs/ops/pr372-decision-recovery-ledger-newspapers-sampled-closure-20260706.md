# PR #372 判断链追回：报刊目录代表抽样闭环（2026-07-06）

## 0. 边界

本文件追回 PR #372 中 `TopHub > 报刊`目录代表抽样复核后的判断链。

纳入文件：

- `docs/ops/tophub-newspapers-central-page-01-20260705.md`
- `docs/ops/tophub-newspapers-financial-page-01-20260705.md`
- `docs/ops/tophub-newspapers-legal-page-01-20260705.md`
- `docs/ops/tophub-newspapers-directory-ledger-20260705.md`
- `docs/ops/tophub-newspapers-final-closure-20260705.md`
- `docs/ops/tophub-newspapers-platform-execution-verified-20260705.md`

边界说明：

- 报刊目录不采用逐页穷举，而采用按类别的代表性抽样。
- 实际逐项判断：中央级热门第 1 页 12 个、财经报热门第 1 页 12 个、法制报热门第 1 页 12 个，共 36 个代表节点。
- 未声称中央级后 7 页、财经报后 6 页、法制报第 2 页、日报 / 早报 / 晚报 / 都市报 / 健康报 / 老年报 / 文化教育报、按省份目录已经逐项审完。
- 当前不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 报刊总规则：代表抽样找“公共温度缺口”，不是多订报纸

```yaml
- source_id: "TopHub > 报刊目录代表抽样整体"
  current_route: sampled_closure_with_one_platform_verified_replacement
  decision_status: recovered_from_pr_sampled
  selected_reason: "报刊目录中仍有来源价值，但代表样本显示大量电子报 TopHub 节点停留在 2021—2025 年，有效节点之间同题重复严重，地方报持续流主要由会议、活动、表彰、普法和基层简讯构成，财经与法治内容更适合任务检索。"
  rejected_alternatives:
    - "继续逐页翻完所有报刊目录"
    - "因为报纸权威就全部常驻"
    - "把财经报、法制报、地方报加入 Folo"
    - "用报纸数量补公共温度"
  route_reason: "当前系统缺口不是更多报纸，而是替换一个已经失效的全国广谱日报节点。最终只执行一个等量替换：取消人民日报电子版，加入新华每日电讯电子报，TopHub 总量仍为 84。"
  reuse_rule: "以后遇到报刊来源，先区分报纸本身是否有价值与 TopHub 节点是否真实更新。报刊常驻只考虑低频、全国、广谱、真实更新、能补公共温度的入口；财经、法治、地方和行业报按任务调用。"
  evidence_locator: "docs/ops/tophub-newspapers-final-closure-20260705.md#审核范围; #最终决定"
  next_action: "报刊替换已执行；财经目录替换另属财经 ledger。"
```

---

## 2. 人民日报电子版为什么退出

```yaml
- source_id: "人民日报｜电子版"
  current_route: retired_from_tophub_public_temperature
  decision_status: recovered_from_pr_and_platform_verified
  selected_reason: "原角色是全国性官方日报入口，位于公共温度第 18 位。"
  rejected_alternatives:
    - "继续保留失效节点"
    - "因为人民日报权威就忽略抓取失效"
    - "用人民日报海外版替代"
  route_reason: "中央级样本显示人民日报电子报内容约停留在 2024 年 11 月，与现有 TopHub 第 18 项高度疑似同一节点；final closure 确认 TopHub 节点样本约停留在 2024 年，已经不能承担每日公共温度角色。平台核验显示人民日报电子版已不在全部订阅中。"
  reuse_rule: "来源权威不能覆盖抓取失效。今日热榜刷新时间不等于报纸版面日期；节点失效时优先处理节点，而不是继续保留品牌。"
  evidence_locator: "docs/ops/tophub-newspapers-central-page-01-20260705.md#关键发现; docs/ops/tophub-newspapers-final-closure-20260705.md#决定理由; docs/ops/tophub-newspapers-platform-execution-verified-20260705.md#结果"
  next_action: "已取消。"
```

---

## 3. 新华每日电讯为什么进入公共温度

```yaml
- source_id: "新华每日电讯｜电子报"
  current_route: tophub_public_temperature_verified_replacement
  decision_status: recovered_from_pr_and_platform_verified
  selected_reason: "同为官方全国性日报入口，样本真实更新，编辑密度较高，覆盖全国政策、社会、区域、产业、人物、法治和公共议题，比失效人民日报节点更能承担官方广谱日报角色。"
  rejected_alternatives:
    - "人民日报海外版"
    - "光明日报"
    - "经济日报"
    - "中国青年报"
    - "经济参考报"
    - "法治日报"
    - "财经报或法制报候选替代"
  route_reason: "放入公共温度第 18 位，等量替换人民日报电子版。财经报页面没有出现能替代其全国官方广谱日报角色的节点；法制报页面也确认法治日报过窄，地方法制报无法承担全国公共温度。平台核验显示新华每日电讯已订阅并已加入公共温度分组。"
  reuse_rule: "公共温度报刊入口要广谱而非专业，真实更新而非品牌权威，能覆盖政策、社会、区域、法治和公共议题。专业报刊按任务，不替代广谱日报。"
  evidence_locator: "docs/ops/tophub-newspapers-central-page-01-20260705.md#当前候选池; docs/ops/tophub-newspapers-financial-page-01-20260705.md#候选池变化; docs/ops/tophub-newspapers-legal-page-01-20260705.md#与公共温度层的关系; docs/ops/tophub-newspapers-platform-execution-verified-20260705.md#结果"
  next_action: "已加入。"
```

---

## 4. 中央级旁边来源为什么不选

```yaml
- source_id: "中央级报刊第 1 页其他节点"
  current_route: on_demand_or_rejected_central_newspapers
  decision_status: recovered_from_pr_sampled
  selected_reason: "人民日报海外版、光明日报、经济日报、中国教育报、经济参考报、中国青年报、中国财经报、解放军报、健康时报、科技日报等本身各有来源价值。"
  rejected_alternatives:
    - "人民日报海外版替代人民日报电子版"
    - "光明日报 / 中国教育报 / 科技日报常驻"
    - "经济日报 / 中国财经报继续作为财经结构源"
    - "解放军报 / 健康时报进入公共温度"
  route_reason: "多数节点抓取时效不可靠或停留旧年份。人民日报海外版与人民日报高度重复；光明日报节点半年级滞后；经济日报和中国财经报停在 2022；中国青年报停在 2024；解放军报停在 2023；健康时报时效不透明且已有健康与辟谣来源；科技日报虽有政策和成果报道价值，但样本旧事件显示抓取时效可疑。中国教育报、科技日报、光明日报保留为任务来源。"
  reuse_rule: "中央级报纸先验权威性不能代替节点活性。教育、科技、文化、军事、健康等专业角色按任务查，不占公共温度常驻位置。"
  evidence_locator: "docs/ops/tophub-newspapers-central-page-01-20260705.md#关键发现; #逐项判断; #本页动作"
  next_action: "按需。"
```

---

## 5. 财经报为什么全部按任务

```yaml
- source_id: "财经报热门第 1 页"
  current_route: on_demand_financial_newspaper_task_sources
  decision_status: recovered_from_pr_sampled
  selected_reason: "每日经济新闻、21世纪经济报道、经济参考报、中国经营报、证券日报等仍在更新的财经报纸有宏观、产业、资本市场、公司调查、监管和市场事件价值。"
  rejected_alternatives:
    - "经济参考报继续作为 A2 条件候选"
    - "每日经济新闻常驻"
    - "21世纪经济报道常驻"
    - "中国经营报常驻"
    - "证券日报常驻"
    - "继续翻财经报后 6 页寻找更好公共温度源"
  route_reason: "财经报页面没有产生新的 TopHub 持续订阅候选。当前系统已经有财新、FT中文网、日经中文网、央视经济栏目、36氪出海、第一财经汽车、Counterpoint、华丽志；财经目录也已有财新首页推荐 + 日经中文网每日最新的等量替换方案。继续加入综合财经报只会增加同题重复。经济参考报因此由 A2 降为 B+ 任务来源。"
  reuse_rule: "财经报按宏观政策、产业、公司、监管、资本市场任务调用。综合财经日报质量好也不自动常驻；先比较是否突破现有数据与结构 12 位。"
  evidence_locator: "docs/ops/tophub-newspapers-financial-page-01-20260705.md#总结; #当前有效节点之间的比较; #候选池变化"
  next_action: "按任务。"

- source_id: "经济参考报｜电子报"
  current_route: downgraded_to_task_source
  decision_status: recovered_from_pr_sampled
  selected_reason: "政策—产业—资本市场连接清楚，官方政策语境比纯市场媒体完整，当前样本仍在更新。"
  rejected_alternatives:
    - "作为数据与结构新增常驻"
    - "替代日经中文或财新首页"
  route_reason: "财经报第一页提供多个同类有效节点，但没有任何一个证明值得突破当前 12 个数据与结构位置；其宏观政策角色已由国家统计局、日经中文网、财新和央视经济栏目共同覆盖。"
  reuse_rule: "A2 条件候选可以被同族比较降级。政策—产业—资本市场来源在具体宏观或产业任务中使用。"
  evidence_locator: "docs/ops/tophub-newspapers-central-page-01-20260705.md#A2条件候选; docs/ops/tophub-newspapers-financial-page-01-20260705.md#候选池变化"
  next_action: "按需。"
```

---

## 6. 法制报为什么全部按任务

```yaml
- source_id: "法制报热门第 1 页"
  current_route: on_demand_legal_newspaper_task_sources
  decision_status: recovered_from_pr_sampled
  selected_reason: "法制报的有效价值主要来自新法、司法解释、监管与执法变化，法院和检察机关案例与实务，消费、劳动、婚姻家庭、财产、平台与数据问题的具体纠纷。"
  rejected_alternatives:
    - "法治日报常驻"
    - "检察日报常驻"
    - "上海法治报常驻"
    - "地方法制报常驻"
    - "继续查看法制报第 2 页"
  route_reason: "这些内容在具体问题出现时很有用，但不适合作每日公共信息入口。多数法制报头版仍以综合时政、系统会议和地方政法活动为主；地方报之间高度重复，持续订阅会带来党建、表彰、培训、普法活动和基层简讯。"
  reuse_rule: "法律报刊按具体法律、监管、消费、劳动、平台、知识产权、刑事执行、婚姻家庭、财产或地方司法任务调用；不能用头版日更替代法律条文、裁判文书和专业意见。"
  evidence_locator: "docs/ops/tophub-newspapers-legal-page-01-20260705.md#总结; #本页动作"
  next_action: "按任务。"

- source_id: "法治日报 / 检察日报 / 上海法治报"
  current_route: strong_legal_task_sources_not_public_temperature
  decision_status: recovered_from_pr_sampled
  selected_reason: "法治日报全国性和官方性最强；检察日报有案例、学术和检察监督内容；上海法治报实际纠纷密度高，覆盖消费者、平台、劳动、婚姻家庭、物业、破产和民间借贷。"
  rejected_alternatives:
    - "法治日报替代新华每日电讯"
    - "检察日报或上海法治报进入公共温度"
    - "地方司法报刊进入 Folo"
  route_reason: "法治日报角色过窄且头版综合时政比例高；检察日报和上海法治报实务价值较高，但属于明确任务来源；地方法制报无法承担全国广谱日报的公共温度角色。"
  reuse_rule: "全国法治政策、检察监督、生活法律纠纷分别找不同任务源。法律任务源不能替代全国公共温度。"
  evidence_locator: "docs/ops/tophub-newspapers-legal-page-01-20260705.md#本页最值得保留为任务来源的三个节点; #与公共温度层的关系"
  next_action: "按需。"
```

---

## 7. 报刊为什么不进 Folo

```yaml
- source_id: "报刊目录 Folo 边界"
  current_route: no_folo_candidates_from_newspapers
  decision_status: recovered_from_pr_sampled
  selected_reason: "报刊电子版通常是高频、大体量、版面型来源。"
  rejected_alternatives:
    - "新华每日电讯进入 Folo"
    - "财经报进入 Folo"
    - "法制报进入 Folo"
    - "地方报 / 按省份目录进入 Folo"
  route_reason: "Folo 新增候选为 0，实际仍为 27。报纸电子版不适合再进入 Folo 制造未读负担；TopHub 公共温度只保留一个低频官方广谱入口，具体财经、法治、教育、科技、地方和行业报按任务打开。"
  reuse_rule: "Folo 收长期作者、专业来源和真实会读的少量来源，不收版面型高频报刊。报刊在 TopHub 只能做少量公共温度或数据结构入口。"
  evidence_locator: "docs/ops/tophub-newspapers-final-closure-20260705.md#Folo; docs/ops/tophub-newspapers-central-page-01-20260705.md#本页动作; docs/ops/tophub-newspapers-financial-page-01-20260705.md#本页动作; docs/ops/tophub-newspapers-legal-page-01-20260705.md#本页动作"
  next_action: "无 Folo 动作。"
```

---

## 8. 平台执行状态与财经待办边界

```yaml
- source_id: "报刊平台执行核验"
  current_route: platform_verified_completed_replacement
  decision_status: recovered_from_pr_and_platform_verified
  selected_reason: "根据 TopHub 全部订阅列表和公共温度分组列表核验，新华每日电讯已订阅并加入公共温度，人民日报电子版已不在全部订阅中。"
  rejected_alternatives:
    - "把报刊替换写成未执行"
    - "把财经目录替换写成已执行"
  route_reason: "报刊目录替换已经真实完成：已取消人民日报电子版，已加入新华每日电讯电子报，位置公共温度第 18 位，TopHub 总量仍为 84。同期财经目录仍待执行，因为财新点击排行榜和评论排行榜仍在全部订阅中。"
  reuse_rule: "平台执行必须单独核验。不同目录的待办不能混写：报刊替换完成，不代表财经替换完成。"
  evidence_locator: "docs/ops/tophub-newspapers-platform-execution-verified-20260705.md#结果; #仍待执行"
  next_action: "报刊无待办；财经待办由财经 ledger 处理。"
```

---

## 9. 本补充 ledger 的复用规则

1. 报刊目录是代表抽样闭环，不是全目录逐页穷举。
2. 报纸品牌权威不能覆盖 TopHub 节点失效；刷新时间不等于版面日期。
3. 公共温度只需要一个真实更新、全国广谱、低频官方日报入口。
4. 新华每日电讯替换人民日报电子版，是因为它更能承担当前的官方广谱日报入口，不是因为其他报纸没有价值。
5. 财经报全部按任务：宏观、产业、公司、监管、资本市场，需要时打开。
6. 法制报全部按任务：新法、司法解释、案例、检察监督、消费 / 劳动 / 平台 / 家庭 / 财产纠纷，需要时打开。
7. 专业报刊不能替代公共温度；公共温度也不能替代专业任务源。
8. 报刊电子版不进 Folo，避免版面型高频未读负担。
9. 平台执行状态要单独核验；报刊完成不等于财经完成。

---

## 10. 仍需继续追回

报刊目录代表抽样闭环已补成 ledger。

继续待办：

- 等用户本地提交 master checklist 与 coverage audit 后，再统一登记本文件和此前新增 ledger；
- 娱乐目录大量 closure；
- 旧 ChatGPT 对话追索所有 partial 条目。
