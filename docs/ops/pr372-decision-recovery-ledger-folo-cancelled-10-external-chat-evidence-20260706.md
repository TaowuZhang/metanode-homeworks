# PR #372 判断链追回：Folo 取消 10 项的最近对话与外网事实证据（2026-07-06）

## 0. 边界

本文件接续：

- `docs/ops/pr372-decision-recovery-ledger-external-web-fact-probe-20260706.md`
- `docs/ops/pr372-decision-recovery-ledger-file-library-sample-evidence-20260706.md`
- `docs/ops/pr372-decision-recovery-ledger-folo-early-operation-chain-20260706.md`

本轮使用两类证据：

1. 最近 7 天 ChatGPT 对话检索：确认用户在 2026-07-03 明确要求在 Folo 取消 10 项；
2. 外网搜索 / 页面打开：补当前公开页面事实，用于判断这些来源当前形态。

这 10 项是：

- 少数派
- 36氪 24 小时热榜
- Bing 每日壁纸
- InfoQ 推荐
- 好价品类榜 12 小时
- 好价品类榜 24 小时
- 金十数据
- 人人影视资讯
- New Routes
- Anthropic News

重要边界：

- 最近对话证据证明“当时要求取消这 10 项”；
- 外网证据证明“当前页面形态或搜索状态”；
- 二者合在一起支持判断链，但外网事实不能冒充旧对话原话；
- 抓取失败、搜索不到或命中不可靠的项目明确标 partial / insufficient；
- 当前仍不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 总规则：取消 Folo 不是丢弃价值，而是重新路由

```yaml
- source_id: "Folo 取消 10 项整体"
  current_route: removed_from_folo_but_not_value_erased
  decision_status: recovered_from_recent_chat_plus_external_probe
  selected_reason: "最近对话确认用户明确要求取消这 10 项 Folo 订阅。共同原因不是它们全无价值，而是它们不适合继续占用 Folo 的长期关系 / 未读流位置。"
  rejected_alternatives:
    - "继续留在 Folo 形成未读债务"
    - "取消后视为永久无价值"
    - "把取消动作改写成模型自动决定"
    - "用外网事实冒充旧对话原话"
  route_reason: "这些来源大致分为高频媒体、公共雷达、视觉/壁纸、促销/交易工具、版权风险影视资讯、官方厂商新闻、来源身份不明等类型。取消 Folo 后可转 TopHub、按需、工具、候选复查或直接退休。"
  reuse_rule: "Folo 只留长期关系和真实会读的来源；高频公共雷达、工具、日历、促销、壁纸、厂商公告、身份不明源优先退出 Folo。"
  evidence_kind: recent_chat_instruction_plus_external_current_fact
  evidence_locator: "recent ChatGPT conversation 2026-07-03; external web/open 2026-07-06"
  next_action: "作为取消 10 项总规则。"
```

---

## 2. 少数派：从 Folo 退出，不等于失去任务价值

```yaml
- source_id: "少数派"
  current_route: removed_from_folo_possible_tophub_or_on_demand
  decision_status: recovered_from_recent_chat_plus_external_fact
  selected_reason: "最近对话确认少数派在用户要求取消的 10 项中。外网页面显示少数派首页有共创、PRIME、Matrix、栏目、Pi Store、推荐、全部、讨论、热门、关注等入口，并同时出现编辑内容、用户讨论、会员内容、商品 / Pi Store 等混合形态。"
  rejected_alternatives:
    - "继续作为 Folo 长期未读源"
    - "把少数派整体视为单一内容源"
    - "把少数派商城和少数派内容流混在一起"
    - "因为取消 Folo 就永久排除少数派"
  route_reason: "少数派仍可作为工具、产品、生活方式和社区材料入口，但首页形态混合、更新量较高、内容类型复杂。它更适合 TopHub / 按需 / Matrix 或具体栏目复查，而不是整站占用 Folo 长期关系。"
  reuse_rule: "平台型媒体要拆栏目和用途：编辑深读、社区讨论、商城、工具推荐、会员内容不能整站合并判断。"
  evidence_kind: recent_chat_instruction_plus_external_current_fact
  evidence_locator: "recent ChatGPT conversation 2026-07-03; https://sspai.com/"
  next_action: "退出 Folo；按栏目或任务复查。"
```

---

## 3. 36氪 24 小时热榜：早期取消，高频商业科技热榜不占 Folo

```yaml
- source_id: "36氪 24 小时热榜"
  current_route: removed_from_folo_tophub_or_on_demand_business_tech_radar
  decision_status: recovered_from_recent_chat_external_partial
  selected_reason: "最近对话确认 `36氪 24 小时热榜` 在用户明确要求取消的 10 项中。外网打开 36氪首页时遇到安全检测页面，未能抓到具体热榜正文；另有公开资料显示 36氪是科技与财经类媒体。"
  rejected_alternatives:
    - "继续作为 Folo 未读源"
    - "把商业科技热榜当长期阅读关系"
    - "用安全检测页内容硬写逐源理由"
  route_reason: "即使 36氪本身有科技商业价值，`24 小时热榜` 这个口径更像高频公共热度 / 商业科技雷达，不适合 Folo。后续科技 ledger 已把 36氪相关判断转向具体功能位，例如出海、深氪或按需，而不是热榜未读。"
  reuse_rule: "媒体品牌和榜单口径分开；热榜优先 TopHub / 按需，深度栏目才有 Folo 复查资格。"
  evidence_kind: recent_chat_instruction_plus_external_partial
  evidence_locator: "recent ChatGPT conversation 2026-07-03; https://www.36kr.com/ external partial"
  next_action: "退出 Folo；若需 36氪，按具体栏目复查。"
```

---

## 4. Bing 每日壁纸：视觉材料不占 Folo 未读流

```yaml
- source_id: "Bing 每日壁纸"
  current_route: removed_from_folo_visual_tool_or_wallpaper_on_demand
  decision_status: recovered_from_recent_chat_external_insufficient
  selected_reason: "最近对话确认 Bing 每日壁纸在用户明确要求取消的 10 项中。本轮外网直接打开 Bing 普通首页和 HPImageArchive 接口没有获得可用正文证据。"
  rejected_alternatives:
    - "继续作为 Folo 未读源"
    - "把每日壁纸当深读来源"
    - "用抓取失败的 Bing 页面硬写内容判断"
  route_reason: "从来源类型看，Bing 每日壁纸属于视觉 / 壁纸 / 每日图像材料，更适合作为视觉漫游或壁纸工具按需调用，而不是文字阅读型 Folo 未读流。当前外网证据不足，只能用最近对话取消事实和类型判断支撑，不补编页面事实。"
  reuse_rule: "每日图像 / 壁纸源默认不进 Folo；如有视觉恢复价值，放视觉与自然、浏览器工具或按需。"
  evidence_kind: recent_chat_instruction_plus_external_insufficient
  evidence_locator: "recent ChatGPT conversation 2026-07-03; external Bing open insufficient"
  next_action: "退出 Folo；如需视觉材料另设视觉入口。"
```

---

## 5. InfoQ 推荐：行业媒体有价值，但推荐流不必形成逐篇未读

```yaml
- source_id: "InfoQ 推荐"
  current_route: removed_from_folo_tophub_or_on_demand_engineering_media
  decision_status: recovered_from_recent_chat_plus_external_fact
  selected_reason: "最近对话确认 InfoQ 推荐在用户要求取消的 10 项中。外网页面显示 InfoQ 首页包含企业动态、行业深度、AI&大模型、出海、后端、芯片&算力、架构、大数据、软件工程、云计算、大前端、管理/文化等栏目，并有编辑精选、专题、会议 / 课程 / 写作社区等混合内容。"
  rejected_alternatives:
    - "继续作为 Folo 推荐流"
    - "把所有工程媒体推荐流都当必读"
    - "把会议、课程、写作社区和技术报道混成单一来源"
  route_reason: "InfoQ 适合了解行业面、技术组织、架构和工程文化，但推荐流混合报道、活动、课程和社区内容，不必形成逐篇未读。更适合 TopHub、按需、具体专题或任务触发。"
  reuse_rule: "工程媒体推荐流先判断是行业雷达还是长期作者关系；行业雷达不自动进 Folo。"
  evidence_kind: recent_chat_instruction_plus_external_current_fact
  evidence_locator: "recent ChatGPT conversation 2026-07-03; https://www.infoq.cn/"
  next_action: "退出 Folo；按需或 TopHub。"
```

---

## 6. 好价品类榜 12 / 24 小时：价格榜是购物任务触发，不是 Folo 关系

```yaml
- source_id: "好价品类榜 12 小时 / 好价品类榜 24 小时"
  current_route: removed_from_folo_purchase_task_trigger
  decision_status: recovered_from_recent_chat_external_partial
  selected_reason: "最近对话确认好价品类榜 12 小时和 24 小时在用户明确取消的 10 项中。本轮外网打开什么值得买首页没有获得正文行，搜索未得到足够可引用页面。"
  rejected_alternatives:
    - "继续作为 Folo 未读源"
    - "把价格榜单生成购买任务"
    - "把抓取失败页面硬写成当前事实"
  route_reason: "从名称和旧判断看，它们属于购物 / 好价 / 时间窗口榜单。购物榜单应由真实需求触发，而不是日常未读流；12 小时 / 24 小时这种窗口更强调短周期价格变化，尤其不适合长期阅读关系。"
  reuse_rule: "价格榜、好价榜、优惠窗口榜只在明确购买任务前调用；任务结束退出。"
  evidence_kind: recent_chat_instruction_plus_external_partial
  evidence_locator: "recent ChatGPT conversation 2026-07-03; https://www.smzdm.com/ external partial"
  next_action: "退出 Folo；购物任务按需。"
```

---

## 7. 金十数据：外网页面证明它是高频交易工具 / 快讯流

```yaml
- source_id: "金十数据"
  current_route: removed_from_folo_on_demand_trading_tool
  decision_status: recovered_from_recent_chat_plus_external_fact
  selected_reason: "最近对话确认金十数据在用户明确取消的 10 项中。外网页面标题写着‘一个交易工具’，导航有快讯、头条、会员、日历、视频、数据，并有重要事件、市场快讯、播报设置、桌面通知、声音提示、内容筛选、热度筛选、交易时钟等功能。"
  rejected_alternatives:
    - "继续作为 Folo 未读源"
    - "把高频市场快讯当日常阅读"
    - "用交易工具生成投资动作"
  route_reason: "外网事实强烈支持它是高频交易 / 快讯 / 工具型来源。它可以在金融现场活跃时按需启用，但不应留在 Folo 形成实时市场噪声。"
  reuse_rule: "市场快讯和交易工具按现场任务启用；不进长期 Folo，不能生成投资动作。"
  evidence_kind: recent_chat_instruction_plus_external_current_fact
  evidence_locator: "recent ChatGPT conversation 2026-07-03; https://www.jin10.com/"
  next_action: "退出 Folo；按需。"
```

---

## 8. 人人影视资讯：版权历史和来源风险，不进 Folo 长期流

```yaml
- source_id: "人人影视资讯"
  current_route: removed_from_folo_retired_or_on_demand_with_rights_risk
  decision_status: recovered_from_recent_chat_plus_external_fact
  selected_reason: "最近对话确认人人影视资讯在用户明确取消的 10 项中。外网公开资料显示人人影视长期与字幕组、影视资源、版权争议、关站和侵权案件有关。"
  rejected_alternatives:
    - "继续作为 Folo 影视资讯源"
    - "作为稳定合法影视信息入口"
    - "用影视资源站替代正式影视评论 / 平台 / 媒体来源"
  route_reason: "该来源存在明显版权与站点历史风险，不适合进入公开订阅系统或长期 Folo。影视发现应回豆瓣、正式平台榜、正规媒体评论、影评或按需查询。"
  reuse_rule: "影视资源站、字幕组、下载站和版权风险源默认退出公开订阅；只在合法、历史研究或材料核验场景下谨慎按需。"
  evidence_kind: recent_chat_instruction_plus_external_current_fact
  evidence_locator: "recent ChatGPT conversation 2026-07-03; external public references on 人人影视"
  next_action: "退出 Folo。"
```

---

## 9. New Routes：外网搜索未找到可靠对应源，不能继续占 Folo

```yaml
- source_id: "New Routes"
  current_route: removed_from_folo_unidentified_source_needs_exact_url
  decision_status: recovered_from_recent_chat_external_no_reliable_hit
  selected_reason: "最近对话确认 New Routes 在用户明确取消的 10 项中。本轮外网搜索 `New Routes newsletter travel mobility tech RSS`、`New Routes blog newsletter`、`New Routes Folo RSS New Routes`、`New Routes Substack` 未得到可靠对应结果。"
  rejected_alternatives:
    - "继续作为 Folo 源但不知道它是谁"
    - "用不相关搜索结果补判断"
    - "凭名字推断它的主题"
  route_reason: "来源身份不清时，不应继续占 Folo。若未来要恢复，必须先找到精确 URL、作者、主题、更新形态和与现有来源的差异。"
  reuse_rule: "身份不明来源先退出；恢复前必须补 URL 和来源角色。"
  evidence_kind: recent_chat_instruction_plus_external_search_no_reliable_hit
  evidence_locator: "recent ChatGPT conversation 2026-07-03; web search 2026-07-06"
  next_action: "退出 Folo；needs_exact_url。"
```

---

## 10. Anthropic News：官方厂商公告不等于长期 Folo 关系

```yaml
- source_id: "Anthropic News"
  current_route: removed_from_folo_official_vendor_news_on_demand_or_task_specific
  decision_status: recovered_from_recent_chat_plus_external_fact
  selected_reason: "最近对话确认 Anthropic News 在用户明确取消的 10 项中。外网打开 Anthropic Newsroom 页面可确认它是 Anthropic 官方新闻页。"
  rejected_alternatives:
    - "继续作为 Folo 长期关系"
    - "让厂商公告成为默认 AI 信息源"
    - "把官方 PR 等同于独立判断"
    - "因为官方就自动保留"
  route_reason: "官方厂商公告有事实核验价值，但立场和内容由厂商选择，不适合作为默认 Folo 关系。涉及具体模型、产品、安全事件、API 变化或用户任务时按需打开即可。"
  reuse_rule: "厂商官方新闻作为事实源按需核验，不作为长期判断源；关键事项再回官方页引用。"
  evidence_kind: recent_chat_instruction_plus_external_current_fact
  evidence_locator: "recent ChatGPT conversation 2026-07-03; https://www.anthropic.com/news"
  next_action: "退出 Folo；按需官方核验。"
```

---

## 11. 本补充 ledger 的复用规则

1. 最近对话确认“取消 10 项”的动作，外网补当前事实形态。
2. 取消 Folo 不等于来源永久无价值；是把它从长期关系 / 未读流移出。
3. 平台型媒体要拆栏目和功能位，不能整站判断。
4. 热榜、推荐流、价格榜、市场快讯、厂商公告、壁纸和身份不明源都不适合默认 Folo。
5. 工具型来源进入按需或沃壤工具链接，不制造未读。
6. 官方厂商新闻可作事实核验，不作默认判断源。
7. 抓取失败或搜索不到必须标 partial / no reliable hit，不硬写。
8. 外网事实必须标 `external_current_fact` 或 `external_partial`，不能冒充旧对话原话。

---

## 12. 仍需继续追回

Folo 取消 10 项已补成“最近对话动作 + 外网当前事实”的证据层。

继续待办：

- 若要恢复或精确处理 `New Routes`，必须找到原始 URL；
- `Bing 每日壁纸`、`好价品类榜 12 / 24 小时`、`36氪 24 小时热榜`仍需要更精确页面事实；
- 若要追回旧对话逐字理由，仍需旧对话原文或导出；
- 仍不更新用户本地待提交的 `master-checklist`、`coverage-audit`、军事边界 ledger。
