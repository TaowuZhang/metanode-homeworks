# PR #372 判断链追回：TopHub 账号、通知、App、Widgets 与功能边界（2026-07-06）

## 0. 边界

本文件追回 PR #372 中 TopHub 账号中心、通知机器人、App 下载、使用指南、首页小部件、过滤、收藏、浏览历史、日报、简报栏目、AI 简报、iDaily.today、热文库与首页功能审计的判断链。

纳入文件：

- `docs/ops/tophub-account-and-notification-audit-20260704.md`
- `docs/ops/tophub-app-and-help-audit-20260704.md`
- `docs/ops/tophub-widgets-and-remaining-settings-audit-20260704.md`
- `docs/ops/tophub-homepage-audit-20260704.md`
- `docs/ops/tophub-feature-audit-20260704.md`
- `资源/雷达/TopHub 使用配置.md`
- `资源/雷达/TopHub 能力地图.md`
- `资源/雷达/信息流最终分工.md`

边界说明：

- 这些文件多为 2026-07-04 的平台能力实测，部分数量记录对应当时阶段，不覆盖后续已更新的当前真实 80 项 TopHub 基线。
- 本文件不重新审查单个来源，不更新平台设置，不启用任何通知或自动化。
- 本文件追回的是：平台功能如何使用、为什么关闭、为什么按需、为什么不能把功能可用误写成来源判断完成。
- 当前不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 平台能力总规则：功能可用不等于应该启用

```yaml
- source_id: "TopHub 平台能力整体"
  current_route: manual_low_noise_feature_boundary
  decision_status: recovered_from_pr
  selected_reason: "TopHub 提供首页、更多目录、小部件、通知机器人、过滤器、收藏、历史、日报、简报、热文库、话题聚合、热点日历、App 和 API 等功能。"
  rejected_alternatives:
    - "因为平台提供功能就逐项启用"
    - "为了把菜单点完而继续扩张操作"
    - "把功能审计结果写成来源选择已经完成"
    - "把热榜、简报、通知和 API 变成自动化信息管道"
  route_reason: "最终边界是低噪声、手动、按需、可核验。TopHub 用于公共温度、发现和按需查看，不承担即时通知、不自动推送追踪结果、不把候选自动变成待办。"
  reuse_rule: "以后遇到平台能力，先问它会增加清晰度，还是增加推送、未读、自动变更、静默隐藏、重复入口或错误责任。不能因功能存在就启用。"
  evidence_locator: "docs/ops/tophub-account-and-notification-audit-20260704.md; docs/ops/tophub-app-and-help-audit-20260704.md; docs/ops/tophub-widgets-and-remaining-settings-audit-20260704.md"
  next_action: "保持低噪声配置。"
```

---

## 2. 通知机器人：保持空订阅、空渠道

```yaml
- source_id: "TopHub 通知机器人"
  current_route: disabled_no_channels_no_subscriptions
  decision_status: recovered_from_pr
  selected_reason: "页面可将节点上榜、上新和追踪结果推送到企业微信、钉钉、飞书、Telegram、Discord 等外部渠道。"
  rejected_alternatives:
    - "配置企业微信、钉钉、飞书、Telegram 或 Discord"
    - "测试推送能力"
    - "把追踪结果自动推送到聊天或办公平台"
    - "让热榜变化触发即时打断"
  route_reason: "通知机器人会把本来按需查看的信息变成主动打断。节点上榜、上新和追踪结果更新频繁，容易重新制造通知负担。重要事项若真的需要提醒，应单独进入日历或专门条件监测，而不是把所有热榜变化推送出去。"
  reuse_rule: "通知只服务明确责任、明确触发条件和明确退出条件。热榜、候选、公共回声和宽泛追踪结果不进通知机器人。"
  evidence_locator: "docs/ops/tophub-account-and-notification-audit-20260704.md#通知机器人"
  next_action: "保持空订阅、空渠道。"
```

---

## 3. 账号中心：保持现有登录关系，不新增绑定

```yaml
- source_id: "TopHub 账号中心"
  current_route: stable_account_no_extra_binding
  decision_status: recovered_from_pr
  selected_reason: "账户长期有效，当前登录方式能正常使用；每行 4 个卡片与桌面宽度和首屏快速扫描方式匹配。"
  rejected_alternatives:
    - "新增 QQ、微博、GitHub 或 Google 绑定"
    - "解除现有可用绑定"
    - "为了测试修改密码"
    - "减少或增加卡片密度"
  route_reason: "新增绑定不会改善信息流，只会增加账户关系和后续维护。密码只在确有安全原因、重复使用或怀疑泄露时单独处理，不属于本轮功能审查。每行 4 个卡片继续保持。"
  reuse_rule: "账户设置只在安全、登录失败或明确管理需求出现时变更；不因平台支持更多绑定就添加关系。"
  evidence_locator: "docs/ops/tophub-account-and-notification-audit-20260704.md#账户中心"
  next_action: "保持。"

- source_id: "隐私记录边界"
  current_route: no_personal_identifiers_in_ops_record
  decision_status: recovered_from_pr
  selected_reason: "账号审计只需记录功能和配置结论。"
  rejected_alternatives:
    - "记录用户 UID"
    - "记录昵称"
    - "记录已绑定第三方账号具体名称"
    - "记录其他可识别个人身份信息"
  route_reason: "功能审计不需要保存可识别身份信息。"
  reuse_rule: "ops 记录只记功能状态、判断和后续边界，不保存不必要的账号标识。"
  evidence_locator: "docs/ops/tophub-account-and-notification-audit-20260704.md#隐私记录边界"
  next_action: "继续遵守。"
```

---

## 4. 首页偏好与小部件：展示层，不是来源层

```yaml
- source_id: "TopHub 首页偏好开关"
  current_route: keep_low_noise_homepage_preferences
  decision_status: recovered_from_pr
  selected_reason: "关闭赞助商广告与首页推荐节点可以减少首页额外噪声；校务、政务、专栏等聚合卡可能模糊具体来源边界。"
  rejected_alternatives:
    - "开启首页推荐节点"
    - "开启校务 / 政务聚合"
    - "把聚合卡当成新增来源能力"
  route_reason: "七个自定义分组已经承担主要路由，无需再加平台聚合层。"
  reuse_rule: "首页聚合只在它提高路由清晰度时启用；若模糊来源边界或造成重复，关闭。"
  evidence_locator: "docs/ops/tophub-account-and-notification-audit-20260704.md#首页偏好开关"
  next_action: "保持低噪声。"

- source_id: "首页小部件"
  current_route: display_layer_not_source_layer
  decision_status: recovered_from_pr
  selected_reason: "首页小部件属于展示层，不会增加独立信源，只会把已经订阅的同类内容合并成一张首页卡片。"
  rejected_alternatives:
    - "订阅专栏聚合"
    - "订阅政务聚合"
    - "订阅校务聚合"
    - "订阅公众号聚合"
    - "把小部件当成新来源层"
  route_reason: "专栏聚合会把具体来源混为一张卡，削弱来源边界；政务聚合会与精确保留的少量政务节点重复；校务聚合当前没有持续学校、学院或考试通知需求；公众号尚未形成稳定订阅集合。"
  reuse_rule: "小部件只作展示，不作来源判断。只有已有一组稳定来源且聚合能改善使用时才开启。"
  evidence_locator: "docs/ops/tophub-widgets-and-remaining-settings-audit-20260704.md#首页小部件; #公众号聚合"
  next_action: "不新增首页小部件。"
```

---

## 5. 内容过滤器：保持关闭，不用静默隐藏替代判断

```yaml
- source_id: "TopHub 内容过滤器"
  current_route: disabled_no_keywords_no_domains
  decision_status: recovered_from_pr
  selected_reason: "过滤器支持关键词和域名，保存后会过滤首页、热榜页和节点页中匹配内容。"
  rejected_alternatives:
    - "添加 AI / 投资 / 日本 / 健康 / Claude 等宽词"
    - "用过滤器替代来源选择、分组路由和按需跳过"
    - "全局隐藏令人厌烦或高频出现的话题"
  route_reason: "过滤器不是排序或折叠，而是在内容进入视野前静默隐藏。用户看不到哪些内容被挡住；关键词过宽会误伤；同一个词可能同时出现在噪声和真正值得观察的上下文里；结果容易被误认为今天没有这类内容。"
  reuse_rule: "只有高度稳定、无歧义的垃圾词或域名长期反复出现，才考虑临时添加；添加后必须记录它会同时影响首页、热榜页和节点页。"
  evidence_locator: "docs/ops/tophub-widgets-and-remaining-settings-audit-20260704.md#内容过滤器"
  next_action: "保持关闭。"
```

---

## 6. 收藏与浏览历史：轻量痕迹，不是待办仓库

```yaml
- source_id: "TopHub 我的收藏"
  current_route: empty_lightweight_manual_bookmark_layer
  decision_status: recovered_from_pr
  selected_reason: "收藏夹与 App 客户端同步，适合之后明确还要返回的原始页面、正在比较且尚未完成判断的少量条目、临时需要跨网页与 App 接续阅读的内容。"
  rejected_alternatives:
    - "为了测试功能随便收藏"
    - "把收藏夹当内容仓库或待办清单"
    - "收藏所有看起来有点意思的候选"
    - "重复保存已经进入 Folo、沃壤、Apple Music、日历或其他正式系统的内容"
  route_reason: "当前收藏数为 0，继续保持。收藏不是未读债务，也不设清空目标；内容进入更合适长期载体后，应从收藏移除。"
  reuse_rule: "收藏只保存少量有明确返回理由的临时页面；长期判断和证据进入沃壤，持续关系进入 Folo 或原平台。"
  evidence_locator: "docs/ops/tophub-widgets-and-remaining-settings-audit-20260704.md#我的收藏"
  next_action: "保持为空。"

- source_id: "TopHub 浏览历史"
  current_route: passive_recent_trace_not_task_list
  decision_status: recovered_from_pr
  selected_reason: "浏览历史记录最近浏览过的内容，可临时找回刚刚打开过但没有收藏的页面，核对某个候选来源是否确实打开过，在短时间内接续浏览。"
  rejected_alternatives:
    - "清空浏览历史"
    - "把历史记录解释成未完成任务"
    - "把历史当收藏、待办、阅读队列或长期档案"
    - "因出现陌生或旧条目就逐项清理"
  route_reason: "历史是被动留下的近期访问痕迹，不由历史数量制造完成压力。博物志来源标签显示为小黄鱼播客更像 TopHub 内部节点别名或映射问题，不影响独立站和 RSS 验证。"
  reuse_rule: "浏览历史自然滚动即可；需要找回刚看过的页面时再用。"
  evidence_locator: "docs/ops/tophub-widgets-and-remaining-settings-audit-20260704.md#浏览历史"
  next_action: "不清空。"
```

---

## 7. 首页与更多菜单：分组式公共天气面板，不走全部页

```yaml
- source_id: "TopHub 首页"
  current_route: grouped_lightweight_public_weather_panel
  decision_status: recovered_from_pr
  selected_reason: "首页把当前订阅节点按既有顺序以卡片形式铺开，七个分组会直接出现在首页顶部，一屏能并排比较多个平台，也能快速看同一时刻各平台关注点是否重合。"
  rejected_alternatives:
    - "把首页当新内容系统"
    - "在全部页向下滚动作为默认路线"
    - "把每个榜单都当值得看"
    - "把首页变成清空式阅读任务"
  route_reason: "首页与全部内容是同一组节点的不同展示。正确用法是使用单个自定义分组进行短暂扫视，而不是在全部里向下滚动。每次打开只进入一个分组。"
  reuse_rule: "首页只作为轻量入口；看到值得追的内容再进原节点或另行核验。"
  evidence_locator: "docs/ops/tophub-homepage-audit-20260704.md#首页本质; #当前定位; #当前决定"
  next_action: "保留。"

- source_id: "更多菜单"
  current_route: directory_and_structure_tools_not_immediate_expansion
  decision_status: recovered_from_pr
  selected_reason: "更多菜单包含报刊、设计、校务、政务、专栏、苹果、公众号、小部件、自定义分组等入口。"
  rejected_alternatives:
    - "因为更多目录仍存在就立即扩张"
    - "继续探索所有额外目录"
    - "把目录入口当摘要或核验能力"
  route_reason: "报刊、设计、校务、政务、专栏、苹果、公众号是额外节点目录，用于寻找其他公开节点，不是新的摘要或核验能力；小部件是展示工具；自定义分组是结构管理。当前测试期不立即增加节点。"
  reuse_rule: "更多菜单只在明确缺口出现时进入；新增来源必须说明它填补哪个缺口。"
  evidence_locator: "docs/ops/tophub-homepage-audit-20260704.md#更多菜单"
  next_action: "不继续探索额外目录。"
```

---

## 8. 日报、站内 AI 简报、iDaily.today：日期快照与独立付费产品，不是默认早报

```yaml
- source_id: "TopHub 日报"
  current_route: date_snapshot_and_backtracking_index
  decision_status: recovered_from_pr
  selected_reason: "日报页包含 AI 简报、早报聚合、晚报聚合、日报周刊聚合、新闻联播、历史上的今天、摸鱼通知、毒鸡汤和年度进度条。它有查看往日、早晚报目录、日报周刊聚合、新闻联播节目索引等回溯价值。"
  rejected_alternatives:
    - "把日报当每日默认首页"
    - "顺序阅读或清空日报"
    - "把早报、晚报和周刊聚合直接当优先级"
  route_reason: "早报、晚报和日报周刊本质上仍是大量节点更新拼接，财经、体育、汽车、技术、论文、媒体早报和娱乐内容混排，重复很多。日报定位应是免费的日期快照、早晚报目录和跨节点回溯入口。"
  reuse_rule: "需要回顾某一天的信息环境或找某个日期线索时用日报；不作为每天必读。"
  evidence_locator: "docs/ops/tophub-feature-audit-20260704.md#日报"
  next_action: "按需。"

- source_id: "TopHub 站内 AI 简报"
  current_route: three_minute_scan_not_fact_or_action_basis
  decision_status: recovered_from_pr
  selected_reason: "站内 AI 简报确实比整页短，能快速显出体育、商业、公共事件、天气和政策热点。"
  rejected_alternatives:
    - "作为事实核验"
    - "作为行动依据"
    - "作为唯一早报"
    - "因为怕错过就购买额外服务"
  route_reason: "样本中出现主题重复、措辞含混、来源不可见、热榜传播强度与新闻事实和 AI 归纳压在同一层、核心资讯混有猎奇娱乐低行动价值内容。"
  reuse_rule: "AI 简报只用于三分钟扫视；任何事实、行动、健康、金融、法律和项目判断都要回原始来源。"
  evidence_locator: "docs/ops/tophub-feature-audit-20260704.md#TopHub站内AI简报"
  next_action: "不购买额外服务。"

- source_id: "iDaily.today / 今日简报"
  current_route: independent_paid_ai_push_product_on_demand_trial_only
  decision_status: recovered_from_pr
  selected_reason: "它声称从 10,000+ 今日热榜信息源生成 AI 摘要，支持每天早晚推送、官方主题、个性化简报和多种送达方式。"
  rejected_alternatives:
    - "购买月度或年度会员"
    - "开启 3 天试用"
    - "假定 8 元会员包含完整自定义能力"
    - "用邮件、飞书、钉钉等推送替代当前工作流"
  route_reason: "它是独立收费产品和控制台，不应视为 TopHub 会员自动包含的原生能力。当前尚未证明独特增量；官方主题与 TopHub、话题聚合和 Folo 高度重叠；真正可能有价值的个性化信息源与关键词能力可能属于另外按需付费方案；推送会重新制造通知和阅读负担。"
  reuse_rule: "只有出现高度限定的自定义简报任务，且可连续验证能替代某项实际人工巡检，才考虑短期试用。"
  evidence_locator: "docs/ops/tophub-feature-audit-20260704.md#外部服务iDaily.today今日简报"
  next_action: "不试用、不购买。"
```

---

## 9. 简报栏目、热文库与视频流转：目录与候选，不是阅读正文

```yaml
- source_id: "TopHub 简报栏目"
  current_route: periodic_source_directory_and_subscription_management
  decision_status: recovered_from_pr
  selected_reason: "简报栏目显示约 120 个 Daily 节点，支持热门/最新排序、展开节点条目、单节点订阅取消、批量订阅本页全部节点、返回日报页。"
  rejected_alternatives:
    - "作为每日阅读页"
    - "把热门理解为质量排序"
    - "点击批量订阅本页全部节点"
    - "因为目录里出现某来源就自动加入当前系统"
  route_reason: "简报栏目负责找和管理周期性节点，日报负责聚合这些节点在某一天的更新。它是来源目录，不是经过压缩的简报正文。"
  reuse_rule: "周期性来源从简报栏目发现后，仍需按来源角色复查；绝不批量订阅。"
  evidence_locator: "docs/ops/tophub-feature-audit-20260704.md#顶部简报栏目"
  next_action: "不继续翻简报目录，不批量订阅。"

- source_id: "TopHub 热文库与 B站 / 公众号筛选"
  current_route: limited_platform_candidate_discovery
  decision_status: recovered_from_pr
  selected_reason: "热文库可在窄筛选时发现平台内容候选；B站科技数码可发现产品、硬件、影像和创客视频；B站知识适合陌生领域漫游。"
  rejected_alternatives:
    - "默认浏览 1000 页热文库"
    - "把热度数据当质量"
    - "直接批量订阅 UP 主"
    - "把 B站知识当学术性或可靠性标签"
    - "把公众号科技筛选当中文科技长文入口"
  route_reason: "热文库默认页内容质量跨度极大，重复转载、营销、标题党和低信息密度内容很多；B站分类仍混有广告、整活、标题党、立场先行、健康建议和政治叙事；公众号科技筛选结果大量早安祝福、鸡汤、八卦、娱乐转载和同文多号重复。"
  reuse_rule: "热文库只做有限候选发现；视频候选回原平台看完整页面，值得处理时交给 BibiGPT；同一创作者多次提供独特价值后才进入 Folo 观察期。"
  evidence_locator: "docs/ops/tophub-feature-audit-20260704.md#热文库默认页; #哔哩哔哩; #公众号; #视频统一流转"
  next_action: "按需。"
```

---

## 10. App、帮助、赞助商广告与 API：不制造新任务

```yaml
- source_id: "TopHub App"
  current_route: optional_mobile_client_not_new_information_source
  decision_status: recovered_from_pr
  selected_reason: "App 提供 iOS App Store、Google Play 下载入口，支持 iPhone、iPad 与 Android；网站已经适配移动版，不想安装 App 时可直接使用移动浏览器。"
  rejected_alternatives:
    - "为了完成配置审查而安装 App"
    - "把 App 当新信息源"
    - "为了手机通知或持续刷新而安装"
  route_reason: "App 主要增量是移动端访问便利，不会改变订阅节点、分组、追踪器、收藏与历史等账户数据。当前没有手机通知、持续刷新或移动端高频浏览目标。"
  reuse_rule: "只有经常需要在手机查看分组、收藏或临时浏览时再安装。"
  evidence_locator: "docs/ops/tophub-app-and-help-audit-20260704.md#App下载页; #是否安装"
  next_action: "不安装。"

- source_id: "使用指南"
  current_route: no_more_full_text_help_review_after_real_operation_audit
  decision_status: recovered_from_pr
  selected_reason: "菜单中的使用指南打开 help 页面。"
  rejected_alternatives:
    - "为了覆盖帮助文档每一段文字重复浏览"
  route_reason: "本轮已经通过真实页面逐项完成首页、日报、动态、追踪、榜中榜、热文库、话题、日历、来源目录、订阅分组、过滤、收藏、历史、通知、账户设置与 App 下载的操作审查。真实操作结果比一般性说明更有价值。"
  reuse_rule: "官方帮助用于补疑点，不替代真实界面核验。"
  evidence_locator: "docs/ops/tophub-app-and-help-audit-20260704.md#使用指南"
  next_action: "不逐段复读。"

- source_id: "赞助商广告"
  current_route: ignored_commercial_entry_not_information_workflow
  decision_status: recovered_from_pr
  selected_reason: "入口服务于广告展示或商业合作。"
  rejected_alternatives:
    - "继续审查赞助商广告页面"
    - "购买或查看赞助商服务"
  route_reason: "当前账户已开启会员关闭赞助商广告，该页不属于信息获取、订阅路由或日常使用能力。"
  reuse_rule: "广告合作入口不进入信息流审计，除非要处理商业合作本身。"
  evidence_locator: "docs/ops/tophub-app-and-help-audit-20260704.md#赞助商广告"
  next_action: "不查看。"

- source_id: "TopHub API 开放平台"
  current_route: disabled_until_specific_data_need
  decision_status: recovered_from_pr
  selected_reason: "API 平台可能用于开发自有应用或脚本、批量读取 TopHub 数据、将热榜数据接入沃壤或其他系统、建立自动刷新、分析、通知或存档流程。"
  rejected_alternatives:
    - "注册 API"
    - "申请密钥"
    - "试调接口"
    - "把 TopHub 变成自动化信息管道"
  route_reason: "当前边界是不做自动刷新、不做通知机器人、不把 TopHub 变成新的自动化信息管道，先观察真实使用效果。"
  reuse_rule: "只有出现明确、持续且无法由网页满足的数据需求时，再单独评估 API。"
  evidence_locator: "docs/ops/tophub-app-and-help-audit-20260704.md#API开放平台"
  next_action: "不注册、不试调。"
```

---

## 11. 本补充 ledger 的复用规则

1. 平台支持某功能，不等于当前系统应该启用。
2. 通知机器人、外部推送、自动刷新、API 和 iDaily.today 这类能力会把按需查看变成主动管道，默认关闭。
3. 首页和小部件是展示层，不是新来源层。
4. 内容过滤器是静默隐藏，不用它替代判断。
5. 收藏是极轻量人工书签，不是内容仓库或待办。
6. 浏览历史是被动痕迹，不清空、不追责、不沉积。
7. 日报是日期快照和回溯索引，不是每日默认首页。
8. 站内 AI 简报只能三分钟扫视，不能做事实、行动或唯一早报。
9. 简报栏目是周期性来源目录，不是简报正文，不能批量订阅。
10. 热文库和 B站/公众号筛选只做候选发现，视频流转回原平台、BibiGPT、Folo、沃壤分层处理。
11. App 是移动访问便利，不是新内容能力。
12. API 只有明确数据需求时才评估。

---

## 12. 仍需继续追回

TopHub 账号、通知、App、Widgets 与功能边界已补成 ledger。

继续待办：

- 等用户本地提交 master checklist 与 coverage audit 后，再统一登记本文件和此前新增 ledger；
- 旧 ChatGPT 对话追索所有 partial 条目；
- 如需继续，可补 `TopHub all-page order audit / source discovery / homepage group map / more directories` 等剩余总览文件。 
