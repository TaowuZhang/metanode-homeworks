# PR #372 判断链追回：社区候选迁移与收口边界（2026-07-06）

## 0. 边界

本文件是 `社区` 大类的补充 ledger，追回社区热门第 1 页、pending actions 与已闭环标签之间的候选迁移判断。

本文件回答的是：

- 为什么第 1 页不直接新增；
- 为什么 `马蜂窝｜热门游记`从强候选升级为新增；
- 为什么 `水木社区｜十大热门话题`从弱候选关闭；
- 为什么虎扑不重复审，复用娱乐目录；
- 为什么 LINUXDO 空页面可闭环；
- 为什么社区大类仍不能写成全量闭环。

纳入文件：

- `docs/ops/tophub-community-popular-page-01-20260705.md`
- `docs/ops/tophub-community-pending-actions-20260705.md`
- `docs/ops/tophub-community-travel-closure-20260705.md`
- `docs/ops/tophub-community-campus-forums-closure-20260705.md`
- `docs/ops/tophub-community-v2ex-closure-20260705.md`
- `docs/ops/tophub-community-linuxdo-empty-closure-20260705.md`

当前不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 社区目录第 1 页：为什么不直接新增

```yaml
- source_id: "社区热门第 1 页整体"
  current_route: directory_page_review_not_direct_subscription_action
  decision_status: recovered_from_pr
  selected_reason: "第 1 页暴露了一组候选和重复入口：马蜂窝、水木、吾爱破解、贴吧、虎扑、知乎话题、汽车之家、虫部落、北大未名等。"
  rejected_alternatives:
    - "只看第 1 页就新增马蜂窝或水木"
    - "把社区目录整体灌入生活与社区"
    - "把贴吧、虎扑、知乎话题当公共温度补充"
    - "把吾爱破解当科技雷达或软件发现常驻"
  route_reason: "第 1 页没有一个节点值得在只看第一页时直接加入 TopHub。社区目录不能按目录名整体灌入生活与社区，因为同页同时包含公共情绪、软件论坛、汽车经验、旅行游记、高校论坛和技术社区。必须进入同族页面比较后再决定。"
  reuse_rule: "目录首页只能形成候选与缺口，不直接形成订阅动作；要回同标签完整 closure 比较后再升级或关闭。"
  evidence_locator: "docs/ops/tophub-community-popular-page-01-20260705.md#本页结论; #当前动作"
  next_action: "已完成第 1 页候选迁移；第 2—20 页仍未全量闭环。"
```

---

## 2. 马蜂窝：从第 1 页强候选升级为新增

```yaml
- source_id: "马蜂窝｜热门游记"
  current_route: tophub_life_and_community_upgraded_from_page1_candidate
  decision_status: recovered_from_pr
  selected_reason: "在社区第 1 页，它是唯一明显提供新维度的强候选：具体城市、徒步、自驾、潜水和长途旅行的第一手路线记录。后续旅游标签完整比较后确认，当前七组缺少旅行经验入口，马蜂窝能补完整游记、路线、自驾、徒步和旅行现场经验。"
  rejected_alternatives:
    - "第 1 页直接加入，不等旅游标签比较"
    - "携程攻略推荐游记"
    - "微信旅行热文榜"
    - "背包客栈同时常驻"
    - "飞客茶馆、人民网旅游、文旅产业源替代"
  route_reason: "进入生活与社区，位置在豆瓣小组之后、国家留学网之前；pending actions 后实际顺序为 V2EX 周报之后、国家留学网之前。背包客栈作为国际旅行第一补充按需，不与马蜂窝同时常驻。"
  reuse_rule: "强候选必须经同标签完整比较才能升级；旅行经验入口只提供路线和现场感，不能替代签证、交通、价格、开放时间、天气和安全事实。"
  evidence_locator: "docs/ops/tophub-community-popular-page-01-20260705.md#候选1马蜂窝热门游记; docs/ops/tophub-community-travel-closure-20260705.md#最终决定; docs/ops/tophub-community-pending-actions-20260705.md#已完成动作"
  next_action: "已新增。"
```

---

## 3. 水木：从弱候选关闭

```yaml
- source_id: "水木社区｜十大热门话题"
  current_route: closed_candidate_on_demand_campus_professional_community
  decision_status: recovered_from_pr
  selected_reason: "第 1 页曾保留为弱候选，因为它能看到一部分高校、职业人群和城市中产社区对教育、职业、住房、汽车和生活选择的讨论。"
  rejected_alternatives:
    - "作为泛职业社区入口加入生活与社区"
    - "与 V2EX 周报并存"
    - "把高校背景自动视为高质量来源"
  route_reason: "完整比较 V2EX 与高校论坛后，水木弱候选正式关闭。水木日频筛选弱、标题经常截断、高校与社区黑话多、情感体育八卦混入。V2EX 周报以周频聚合热门主题与高赞回复，更适合承担泛技术职业社区入口。"
  reuse_rule: "弱候选必须能被更完整的同族比较推翻；高校 / 职业人群视角有价值，但不能被误当成社会整体或长期公共入口。具体高校、升学、职业、校友任务再按需。"
  evidence_locator: "docs/ops/tophub-community-popular-page-01-20260705.md#候选2水木社区十大热门话题; docs/ops/tophub-community-campus-forums-closure-20260705.md#最终决定; docs/ops/tophub-community-pending-actions-20260705.md#已关闭的候选"
  next_action: "已关闭；按需。"
```

---

## 4. 第 1 页其他节点：不添加或按需

```yaml
- source_id: "吾爱破解节点"
  current_route: on_demand_reverse_engineering_and_windows_tool_source
  decision_status: recovered_from_pr
  selected_reason: "技术密度不低，能出现绿色软件、注册授权绕过、反调试、招聘网站反反调试、开源工具和精品软件。"
  rejected_alternatives:
    - "吾爱破解今日热帖常驻"
    - "吾爱破解人气热门常驻"
    - "吾爱破解精品软件区常驻"
    - "进入科技雷达或 Folo"
  route_reason: "不添加。它混合破解下载、授权规避、下载工具、来源不明绿色版本和安全研究；不适合作为普通软件发现入口。需要具体逆向、反调试、Windows 工具或软件安全问题时定向搜索。"
  reuse_rule: "工具社区中出现授权绕过、破解下载、来源不明二进制时，只能按任务和风险核验使用，不能做常驻发现源。"
  evidence_locator: "docs/ops/tophub-community-popular-page-01-20260705.md#吾爱破解今日热帖; #吾爱破解精品软件区; docs/ops/tophub-community-pending-actions-20260705.md#社区热门第1页"
  next_action: "按需。"

- source_id: "百度贴吧 / 虎扑 / 知乎话题等公共情绪节点"
  current_route: rejected_duplicate_or_on_demand_mood_sample
  decision_status: recovered_from_pr
  selected_reason: "可作为某些群体情绪和传播样本。"
  rejected_alternatives:
    - "百度贴吧热议榜常驻"
    - "虎扑步行街热帖作为公共温度"
    - "知乎想法热榜 / 话题榜补充知乎体系"
  route_reason: "它们与现有微博、知乎热榜、抖音、今日热榜和微信热文高度重复，但聚合质量更差。虎扑完整判断复用娱乐目录，社区目录不重复审；知乎话题榜存在聚合质量不稳，现有知乎热榜已经足够承担知乎公共注意力。"
  reuse_rule: "公共情绪节点只作按需样本；低质量重复不因来自社区目录就进入公共温度。跨目录重复入口复用已完成判断，不再另立一套。"
  evidence_locator: "docs/ops/tophub-community-popular-page-01-20260705.md#逐节点判断; docs/ops/tophub-community-pending-actions-20260705.md#虎扑"
  next_action: "按需或排除。"

- source_id: "汽车之家 / 北大未名 / 虫部落"
  current_route: on_demand_or_rejected_page1_community_sources
  decision_status: recovered_from_pr
  selected_reason: "汽车之家可在购车或用车阶段提供车主经验；北大未名对北大校内事务有意义；虫部落能看到部分日常求助。"
  rejected_alternatives:
    - "汽车之家论坛精选日报常驻"
    - "北大未名首页推荐常驻"
    - "虫部落最新热门常驻"
  route_reason: "汽车之家只有在具体车型、购车或自驾任务中有明显增量；北大未名校园语境过窄；虫部落既不是稳定技术社区，也不是可靠生活经验策展，与豆瓣小组功能重叠但质量更不稳定。"
  reuse_rule: "社区来源必须有当前现实关系：车型、校园、城市、任务或项目。没有关系就不常驻。"
  evidence_locator: "docs/ops/tophub-community-popular-page-01-20260705.md#逐节点判断"
  next_action: "按需或排除。"
```

---

## 5. Pending actions：三项动作与排序为什么成立

```yaml
- source_id: "社区 pending actions 三项动作"
  current_route: executed_or_verified_community_actions
  decision_status: recovered_from_pr
  selected_reason: "先知、V2EX 周报、马蜂窝分别补三个不同缺口：安全技术雷达、技术使用者周观察、旅行经验入口。"
  rejected_alternatives:
    - "把社区页其他候选一起加入"
    - "把 V2EX 高频板块与周报并存"
    - "把马蜂窝与背包客栈并存"
    - "把看雪与先知并存"
  route_reason: "pending actions 确认 TopHub 84、科技雷达 12、生活与社区 9、Folo 27。科技雷达中先知位于 The Register 之后、科普中国网之前；生活与社区中 V2EX 周报和马蜂窝位于豆瓣小组之后、国家留学网之前。排序逻辑是日常实用 → 健康核验 → 公共生活讨论 → 技术社区周观察 → 旅行经验 → 留学与潮流专题。"
  reuse_rule: "新增社区源必须补不同缺口，并且排序要解释功能邻接关系；不能只说新增了几个。"
  evidence_locator: "docs/ops/tophub-community-pending-actions-20260705.md#已完成动作; #科技雷达最终顺序; #生活与社区最终顺序; #全部页结果"
  next_action: "无本轮配置动作；只追回判断链。"
```

---

## 6. 跨目录复用：虎扑与 LINUXDO

```yaml
- source_id: "虎扑社区入口"
  current_route: cross_directory_reuse_entertainment_hupu_closure
  decision_status: recovered_from_pr
  selected_reason: "虎扑在社区目录出现，但其 16 个节点已经在娱乐目录完整审核，包括步行街、生活、NBA、CBA、足球、游戏电竞、职场、恋爱、股票、汽车、资讯与篮球新闻。"
  rejected_alternatives:
    - "社区目录重复建立虎扑判断"
    - "虎扑步行街重新作为社区公共温度候选"
    - "虎扑进入 Folo"
  route_reason: "复用娱乐目录判断：不新增、不替换、不进入 Folo；只在明确赛事、车型、装备或社区情绪任务中按需查看。跨目录出现不等于新来源关系。"
  reuse_rule: "同一平台跨目录出现时，优先复用已完成的完整同源审核，除非新目录出现了真实不同节点或用途。"
  evidence_locator: "docs/ops/tophub-community-pending-actions-20260705.md#跨目录复用与空页闭环; #虎扑; docs/ops/tophub-community-popular-page-01-20260705.md#虎扑社区步行街热帖"
  next_action: "无；不重复审。"

- source_id: "LINUXDO"
  current_route: empty_page_closed_not_quality_judgment
  decision_status: recovered_from_pr
  selected_reason: "当前 TopHub 社区目录下 LINUXDO 页面为空白，可见节点为 0，没有榜单或订阅项可供判断。"
  rejected_alternatives:
    - "把 LINUXDO 继续列为待审缺口"
    - "对 LINUX DO 社区本身补编质量判断"
  route_reason: "按空页面闭环，不新增、不取消、不进入 Folo。未来页面若重新出现节点，只审核届时新增的实际内容。"
  reuse_rule: "空页面闭环只说明页面事实，不评价站点本体；不把空页面当缺口，也不补编。"
  evidence_locator: "docs/ops/tophub-community-linuxdo-empty-closure-20260705.md#页面事实; docs/ops/tophub-community-pending-actions-20260705.md#LINUXDO"
  next_action: "无。"
```

---

## 7. 社区大类仍未全量闭环

```yaml
- source_id: "社区大类全量状态"
  current_route: partial_not_full_closure
  decision_status: recovered_from_pr
  selected_reason: "当前已审核社区热门第 1 页、地方门户 49/49、安全社区 3/3、汽车论坛 7/7、旅游 11/11、数码 61/61、开发者社区 8/8、高校论坛 4/4、V2EX 30/30、虎扑复用、LINUXDO 空页。"
  rejected_alternatives:
    - "写成社区目录 235 节点已全量闭环"
    - "把社区热门第 2—20 页视作已被标签 closure 覆盖"
  route_reason: "pending actions 明确社区目录 235 个节点仍未整体闭环，缺口是社区热门总目录第 2—20 页，以及尚未提供或尚未形成逐项记录的其他实际页面内容。"
  reuse_rule: "已闭环标签不能自动代表目录总页全量完成；覆盖状态要区分标签闭环、目录页闭环、跨目录复用和空页闭环。"
  evidence_locator: "docs/ops/tophub-community-pending-actions-20260705.md#状态"
  next_action: "社区仍为 partial；第 2—20 页不能冒充完成。"
```

---

## 8. 本补充 ledger 的复用规则

1. 目录第 1 页只能形成候选，不直接形成最终动作。
2. 强候选要经过完整同标签比较才能升级；弱候选可以被更好来源关闭。
3. 社区新增必须补不同缺口，并解释组内位置与功能邻接。
4. 跨目录重复入口复用已完成判断，不重复建立第二份账本。
5. 空页面按页面事实闭环，不补编站点价值。
6. 社区全量状态必须区分已闭环标签、未审目录页、跨目录复用和空页闭环。
7. 论坛、社区、地方门户、校园和汽车经验都按现实任务触发；Folo 只处理作者或稳定刊物关系。

---

## 9. 仍需继续追回

社区候选迁移与 pending actions 已补入。

继续待办：

- 社区热门第 2—20 页若后续需要，必须单独提供或抽链；当前不能写成完成；
- 等用户本地提交 master checklist 与 coverage audit 后，再统一登记本文件和此前新增 ledger；
- 财经逐页；
- 购物逐页；
- 政务、校务、专栏、浏览器与链接；
- 旧 ChatGPT 对话追索所有 partial 条目。
