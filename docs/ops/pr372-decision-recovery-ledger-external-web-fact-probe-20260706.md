# PR #372 判断链追回：外网事实核验补充（2026-07-06）

## 0. 边界

本文件回应一个关键问题：找不到旧对话原文时，不能只在 GitHub 和 File Library 里反刍，也应主动搜索外网，获取新的事实性证据。

本轮通过外网搜索与直接打开候选页面，核验了若干浏览器订阅候选 / 行动触发入口 / 财经日历 / 科技产品媒体入口的当前公开页面事实。

重要边界：

- 外网证据只能证明当前页面事实、来源类型、信息形态、是否像工具 / 商城 / 快讯 / 媒体首页；
- 外网证据不能冒充旧 ChatGPT 对话中的原始判断；
- 外网证据可以用来补 `fact_evidence`、`current_page_role` 和 `route_reason_support`；
- 旧判断仍需从最近对话、File Library、GitHub 原始文件或可追溯的旧材料中找；
- 当前仍不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 总规则：外网事实补证，不冒充旧判断

```yaml
- source_id: "external web fact probe"
  current_route: new_fact_evidence_layer_not_old_chat_transcript
  decision_status: external_fact_probe_recovered
  selected_reason: "当旧对话原文、App 六页 closure 或早期 28 项取消原文找不到时，应主动用外网搜索和页面打开补当前事实。"
  rejected_alternatives:
    - "只在 GitHub / File Library 里循环查同一批文件"
    - "找不到旧原文就停在 needs_source_text"
    - "把外网页面事实说成旧对话原话"
    - "用当前页面事实直接宣布历史判断完成"
  route_reason: "外网事实能校验来源当前是什么形态：行动触发、商城、数据工具、媒体首页、产品媒体、快讯、榜单、导航或资料库。它能支持路由，但不能替代旧判断来源。"
  reuse_rule: "以后缺旧证据时，先补外网页面事实，再把 evidence_kind 标成 `external_current_fact`；旧对话证据仍标 `old_chat_evidence`，两者分开。"
  evidence_locator: "web search/open 2026-07-06"
  next_action: "作为事实补证层。"
```

---

## 2. IT之家「喜加一」：外网页面证明它是行动触发流

```yaml
- source_id: "IT之家「喜加一」"
  current_route: action_trigger_candidate_not_folo_deep_read
  decision_status: external_current_fact_supports_existing_route
  selected_reason: "外网打开页面显示标题为‘「喜加一」最新动态’，页面条目以 Epic / Steam 免费领取、限免截止时间、原价和评论数为核心。"
  rejected_alternatives:
    - "作为 Folo 深读源"
    - "作为 TopHub 常驻阅读源"
    - "作为浏览器日常首页"
    - "作为软件质量判断源"
  route_reason: "页面事实支持此前路由：它是促销 / 领取 / 行动触发入口，而非稳定阅读关系。"
  reuse_rule: "限免、领取、促销类来源只在有领取动作时打开；不制造持续未读。"
  evidence_kind: external_current_fact
  evidence_locator: "https://www.ithome.com/zt/xijiayi"
  next_action: "按任务触发或 2026-07-17 候选复查。"
```

---

## 3. 少数派商城：外网页面证明它是商城 / 软件销售入口

```yaml
- source_id: "少数派商城"
  current_route: commerce_action_candidate_not_media_subscription
  decision_status: external_current_fact_supports_existing_route
  selected_reason: "外网打开页面显示其导航为 Pi Store / 商城首页 / 全部分类，并有‘最受欢迎’、‘最新上架’等商品列表，条目以软件名称、用途和价格为核心。"
  rejected_alternatives:
    - "作为 Folo 内容源"
    - "作为 TopHub 常驻科技雷达"
    - "作为浏览器日常媒体入口"
  route_reason: "页面事实支持此前判断：少数派商城是软件 / 模板 / 工具购买入口，不是深读源；可作为行动触发或购买任务入口。"
  reuse_rule: "商城页按购买 / 比价 / 查权益任务打开；不进入持续订阅流。"
  evidence_kind: external_current_fact
  evidence_locator: "https://sspai.com/mall"
  next_action: "按任务触发。"
```

---

## 4. 东方财富财经日历：外网页面证明它属于数据中心工具

```yaml
- source_id: "东方财富财经日历"
  current_route: finance_calendar_tool_not_reading_feed
  decision_status: external_current_fact_supports_existing_route
  selected_reason: "外网打开页面显示它位于东方财富数据中心，页面导航包含行情、数据、财报、研报、新股、资金、经济、特色数据，并在特色数据中列出财经日历。"
  rejected_alternatives:
    - "作为每日阅读流"
    - "作为 Folo 订阅源"
    - "作为 TopHub 常驻内容节点"
    - "作为投资动作生成器"
  route_reason: "页面事实支持此前路由：财经日历是工具 / 数据入口，适合具体事件、宏观数据、财报或市场日程核验，不是阅读流。"
  reuse_rule: "财经日历进工具链接或按需使用；不生成未读和投资动作。"
  evidence_kind: external_current_fact
  evidence_locator: "https://data.eastmoney.com/cjrl/default.html"
  next_action: "工具 / 按需。"
```

---

## 5. 华尔街见闻财经日历：外网页面可达性有限，但候选类型仍为日历工具

```yaml
- source_id: "华尔街见闻财经日历"
  current_route: finance_calendar_tool_requires_page_specific_recheck
  decision_status: external_current_fact_partial
  selected_reason: "外网打开地址可解析为华尔街见闻 calendar 页面，但本轮抓取没有得到正文行。"
  rejected_alternatives:
    - "因为抓取无正文就判定失效"
    - "因为 URL 可打开就判定为合格订阅源"
    - "作为 Folo 阅读流"
  route_reason: "当前只能确认它作为财经日历 URL 的存在，不能从本轮抓取证明页面完整内容。已有财经 ledger 中将其路由为沃壤工具链接，外网抓取不足以改变该判断。"
  reuse_rule: "抓取无正文的外网页面要标 partial；可打开不等于可用，可用性需浏览器或后续手工核验。"
  evidence_kind: external_current_fact_partial
  evidence_locator: "https://wallstreetcn.com/calendar"
  next_action: "保持工具链接；必要时手工核验页面。"
```

---

## 6. 爱范儿：外网页面证明它是科技 / 产品 / 媒体品牌，而非单一功能位

```yaml
- source_id: "爱范儿"
  current_route: chinese_tech_product_media_candidate_check_function_slot
  decision_status: external_current_fact_supports_existing_route
  selected_reason: "外网打开页面显示爱范儿有 AIGC、产品、公司、商业、汽车、评测、软件等大量分类；品牌描述为关注明日产品的数字潮牌，并有 APPSO、董车会、玩物志等子品牌。"
  rejected_alternatives:
    - "直接加入科技雷达"
    - "直接加入 Folo"
    - "把整个品牌作为单一来源功能位"
    - "用媒体首页替代具体栏目判断"
  route_reason: "页面事实支持此前规则：中文科技媒体候选要按功能位拆，不能因为品牌覆盖产品、软件、汽车、AIGC 就整体加入。"
  reuse_rule: "复查科技媒体时先分栏目和功能位：个人工具、消费电子、汽车、AI、产品评论、商业报道；不能整站打包。"
  evidence_kind: external_current_fact
  evidence_locator: "https://www.ifanr.com/"
  next_action: "候选复查时按栏目拆。"
```

---

## 7. App / 反斗软件：外网搜索没有得到足够事实，不能硬写

```yaml
- source_id: "反斗软件 / App 六页 closure"
  current_route: external_search_insufficient_needs_source_or_page
  decision_status: external_fact_probe_no_reliable_hit
  selected_reason: "本轮外网搜索 `反斗软件 最新发布 App 软件 反斗限免 今日热榜`、`今日热榜 TopHub 反斗软件 最新发布`、`TopHub 今日热榜 App Store 榜单 App 目录 反斗软件` 未得到足够可靠页面事实。"
  rejected_alternatives:
    - "用无关搜索结果补反斗软件判断"
    - "把搜索不到解释成来源不存在"
    - "根据记忆写 App 六页逐源理由"
  route_reason: "外网搜索不足时，只能标证据不足；不能把不相关结果当事实。"
  reuse_rule: "对搜索不到的具体来源，保留 `needs_source_or_page`；下一步要么拿到原始 TopHub 页面 / closure 文件，要么用更精确 URL 手工打开。"
  evidence_kind: external_search_no_reliable_hit
  evidence_locator: "web search queries 2026-07-06"
  next_action: "needs_source_or_page。"
```

---

## 8. 本补充 ledger 的复用规则

1. 外网搜索与页面打开是事实补证层，不是旧对话证据层。
2. 外网页面能证明当前页面类型：行动触发、商城、财经日历工具、科技媒体、导航、榜单等。
3. 外网页面不能直接证明当时为什么选 / 不选，除非旧对话或 ledger 明确引用该事实。
4. 抓取无正文要标 partial，不把可打开当可用。
5. 搜索不到可靠命中时，标 `external_search_no_reliable_hit`，不能硬补。
6. 对促销、商城、财经日历、科技媒体这类来源，新事实已经支持既有路由：按任务、工具、功能位或候选复查，不自动进 Folo / TopHub。
7. 未来继续追 source evidence 时，应并行使用：最近对话、File Library、GitHub、外网事实核验。

---

## 9. 仍需继续追回

外网事实核验已补一层。

继续待办：

- TopHub 早期首轮 28 个取消项原始列表仍未找到；
- App 六页 / 反斗软件仍需原始 TopHub 页面、closure 文件或更精确 URL；
- 浏览器候选和 Folo public candidate 可继续用 File Library 分批抽样；
- 外网事实后续可逐源补，但必须标 `external_current_fact`，不能冒充旧对话原话；
- 仍不更新用户本地待提交的 `master-checklist`、`coverage-audit`、军事边界 ledger。
