# PR #372 判断链追回：浏览器与链接路由（2026-07-06）

## 0. 边界

本文件追回 PR #372 中浏览器清理、旧链接迁居、链接层 / 语境层 / 项目层分工的判断链。

纳入文件：

- `docs/ops/browser-bookmark-cleanup-2026-07-05.md`
- `docs/ops/browser-link-governance-round-2-2026-07-05.md`
- `docs/ops/browser-link-placement-audit-2026-07-05.md`
- `资源/语境/浏览器与链接路由.md`
- `资源/语境/浏览器工具架.md`
- `资源/链接/_index.md`
- `资源/链接/浏览器主题包候选.md`
- `资源/链接/浏览器订阅候选-2026-07-17.md`
- `资源/链接/浏览器项目与学习入口候选.md`
- `资源/链接/浏览器账号与服务入口.md`
- `资源/链接/浏览器待复核与一次性入口.md`
- `资源/链接/浏览器垂直检索候选.md`

当前不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 总规则：浏览器保存抵达，沃壤保存判断

```yaml
- source_id: "浏览器与链接治理整体"
  current_route: browser_for_access_worang_for_links_context_projects
  decision_status: recovered_from_pr
  selected_reason: "597 条旧链接已经完成静态分流：浏览器只保存当前真实入口，URL 候选回到资源/链接，稳定使用语境进入资源/语境，执行和对账证据留在 docs/ops。"
  rejected_alternatives:
    - "把 Chrome / Safari 旧书签完整迁入 Zen"
    - "把 597 条旧链接全部堆进资源/雷达"
    - "把候选 URL 冒充正式资源"
    - "把浏览器书签当成订阅系统"
    - "把执行审计文件当长期知识入口"
  route_reason: "浏览器分工、链接候选、语境规则、项目采用和操作证据各自居留。浏览器导出形成的是 cohort 索引，不替代已有资源/链接.csv 或单条资源/链接/*.md；同名或同 URL 已有正式条目时，复用 canonical 条目。"
  reuse_rule: "以后遇到浏览器旧链接，先问它是抵达入口、外部门牌、稳定使用语境、领域材料、项目证据、账号入口、一次性事项还是历史冗余。不要把所有链接统一迁居。"
  evidence_locator: "docs/ops/browser-link-governance-round-2-2026-07-05.md; docs/ops/browser-link-placement-audit-2026-07-05.md; 资源/语境/浏览器与链接路由.md"
  next_action: "无真实浏览器动作；后续复查候选等待 2026-07-17。"
```

---

## 2. 三个浏览器为什么这样分工

```yaml
- source_id: "Zen"
  current_route: main_work_environment
  decision_status: recovered_from_pr
  selected_reason: "Zen 是默认主浏览器，承担工作、思考与 AI 协作，阅读、作品、自我与滋养，主动检索和开放网页浏览三个活动空间。"
  rejected_alternatives:
    - "把 Chrome 或 Safari 旧书签树完整迁入 Zen"
    - "把 Zen 变成所有技术 / AI / 设计 / 下载资源目录"
  route_reason: "Zen 只保存真实反复使用的入口，不接收完整历史书签树。工具架通过正式界面导入 8 条：PDF24、Regex Vis、JSON Crack、Draw.io、Excalidraw、iconfont、Pexels、Freesound。"
  reuse_rule: "主浏览器只放会真实操作、需要一两秒抵达的入口；候选、历史和资料库回沃壤。"
  evidence_locator: "资源/语境/浏览器与链接路由.md#浏览器分工; docs/ops/browser-bookmark-cleanup-2026-07-05.md#Zen"
  next_action: "保持。"

- source_id: "Chrome"
  current_route: chromium_and_public_service_environment
  decision_status: recovered_from_pr
  selected_reason: "Chrome 承担 Google 搜索与 Google 服务、公共网络、热榜、Folo、社交与视频平台入口、Chromium 插件、网页应用、兼容性体验、必要开发与页面验证。"
  rejected_alternatives:
    - "Chrome 承担所有技术、AI、设计和软件下载资源目录"
    - "把旧 Chrome 书签重新迁入 Zen"
  route_reason: "Chrome 书签栏最终只剩工具架、Photopea 与 Apifox；旧书签文件夹全部消失，但标签组 Myself / Job / Network / WQB / Leetcode / Hack 保持不变。"
  reuse_rule: "Chrome 用于公共服务、Google/Chromium 生态与重型网页应用；不再作为泛收藏库。"
  evidence_locator: "资源/语境/浏览器与链接路由.md#Chrome; docs/ops/browser-bookmark-cleanup-2026-07-05.md#Chrome"
  next_action: "保持。"

- source_id: "Safari"
  current_route: clean_isolated_browser_environment
  decision_status: recovered_from_pr
  selected_reason: "Safari 长期角色是与 Zen、Chrome 分离 Cookie、登录态、权限或站点数据的临时访问环境，也是 WebKit / Safari 兼容性验证环境。"
  rejected_alternatives:
    - "Safari 承担深度阅读收藏库"
    - "Safari 承担 RSS 或隐私专用收藏库"
    - "Safari 承担第三套主题书签体系"
  route_reason: "Safari 旧导出保存 303 条书签与 22 条阅读列表作为历史；用户界面已清空：用户书签 0、阅读列表 0、自定义书签文件夹 0。Safari 不承载收藏体系。"
  reuse_rule: "隔离浏览器保持干净；只在需要隔离登录态、权限、Cookie、WebKit 兼容性时使用。"
  evidence_locator: "资源/语境/浏览器与链接路由.md#Safari; docs/ops/browser-bookmark-cleanup-2026-07-05.md#Safari"
  next_action: "保持。"
```

---

## 3. Zen / Chrome 工具架为什么保留这些、不保留旁边那些

```yaml
- source_id: "Zen 工具架 8 项"
  current_route: stable_context_tools_in_zen
  decision_status: recovered_from_pr
  selected_reason: "PDF24、Regex Vis、JSON Crack、Draw.io、Excalidraw、iconfont、Pexels、Freesound 已形成稳定使用语境：PDF 处理、正则调试、JSON 可视化、正式图、快速草图、中文项目图标、图片/视频素材、音效素材。"
  rejected_alternatives:
    - "把所有设计、素材、开发和工具网站都放浏览器书签栏"
    - "用工具候选填满工具架"
  route_reason: "它们进入资源/语境，而非普通链接候选，因为已经知道什么时候、为什么、怎样使用。Draw.io 与 Excalidraw 并存，是因为一个承担正式结构图，一个承担快速思考图。"
  reuse_rule: "同一功能默认只保留一个常用入口；并存必须解释功能差异；停止真实使用后回到链接候选或历史，不用替代品填补空位。"
  evidence_locator: "资源/语境/浏览器工具架.md#Zen工具架; #使用规则"
  next_action: "保持。"

- source_id: "Chrome 工具架 2 项"
  current_route: chromium_heavy_web_app_tools
  decision_status: recovered_from_pr
  selected_reason: "Photopea 是临时图像编辑，适合 Chromium 网页应用环境；Apifox 是 API 设计、调试与文档默认入口。"
  rejected_alternatives:
    - "把 Postman 与 Apifox 并存为常驻"
    - "把 Figma / Docker / Vercel / Cloudflare 等控制台常驻"
  route_reason: "重型 Chromium 网页应用放 Chrome；账号控制台跟随项目，不进入公共工具架。Postman 只有外部团队或既有集合要求时使用；Docker、Vercel、Cloudflare 跟具体开发/部署项目。"
  reuse_rule: "开发控制台和账号后台按项目激活，不作为公共工具架。"
  evidence_locator: "资源/语境/浏览器工具架.md#Chrome工具架; #非常驻; #使用规则"
  next_action: "保持。"
```

---

## 4. 597 条旧链接怎样分流

```yaml
- source_id: "597 条旧浏览器链接"
  current_route: split_by_role_not_bulk_migrated
  decision_status: recovered_from_pr
  selected_reason: "旧浏览器导出中有正式资源、候选资源、工具、订阅候选、项目学习入口、账号服务、一次性入口、检索工具和大量历史 / 重复 / 搜索冗余。"
  rejected_alternatives:
    - "全部迁入 Zen"
    - "全部放入资源/雷达"
    - "全部做成正式资源"
    - "自动删除历史、重复或冗余链接"
  route_reason: "最终静态分流为：39 条已存在于资源/链接.csv；37 条主题包候选；17 条工具，常驻语境在资源/语境/浏览器工具架；48 条订阅候选；96 条项目与学习入口；43 条账号与服务入口；71 条待复核与一次性入口；2 条垂直检索候选；244 条历史、重复或搜索冗余只留原始导出。"
  reuse_rule: "旧链接要先按角色分类，再决定居留。历史、重复和搜索冗余不自动删除，只留导出；候选不冒充正式资源。"
  evidence_locator: "docs/ops/browser-link-governance-round-2-2026-07-05.md#当前分流; docs/ops/browser-link-placement-audit-2026-07-05.md#纠正后的归宿"
  next_action: "无。"
```

---

## 5. 订阅候选为什么只留链接层

```yaml
- source_id: "浏览器订阅候选 48 条"
  current_route: link_layer_subscription_review_candidates_for_20260717
  decision_status: recovered_from_pr
  selected_reason: "旧链接中确实存在持续更新媒体、平台和公共入口，如财新、晚点、第一财经、FT 中文、36氪、少数派、IT之家、The Verge、TechCrunch、Bloomberg、FT、WSJ、TED、国家地理等。"
  rejected_alternatives:
    - "立即写入资源/雷达/订阅源账本.yml"
    - "因为是媒体或首页就当合格订阅源"
    - "和 TopHub / Folo 已有来源重复订阅"
    - "促销与行动触发页进入持续订阅"
  route_reason: "这些当前仍是 media / platform 类型链接，居留在资源/链接。只有实际进入 Folo 或 TopHub 后，才更新正式订阅账本。复查边界是：先判断 TopHub 是否已经承担公共温度；Folo 只考虑独特、稳定、真实会读的来源；首页、付费墙、聚合站、财经日历和促销页不自动等同于合格订阅源；2026-07-17 先复查现有 27 个 Folo 来源，再按一进一出原则判断。"
  reuse_rule: "订阅候选不是订阅事实。媒体首页、付费墙、聚合站、日历和促销页都必须经 TopHub/Folo 分工复查。"
  evidence_locator: "资源/链接/浏览器订阅候选-2026-07-17.md#状态; #复查边界"
  next_action: "2026-07-17 复查。"
```

---

## 6. 项目与学习入口为什么不常驻浏览器

```yaml
- source_id: "浏览器项目与学习入口 96 条"
  current_route: project_or_course_activated_link_candidates
  decision_status: recovered_from_pr
  selected_reason: "包含课程、考试、开发、编程、学术研究、AI 文献、公共 API、云平台、Apple Developer、Context7、Tauri、Vite、OpenCompass、Statista、中国裁判文书网、合同示范文本库等真实有用入口。"
  rejected_alternatives:
    - "作为通用书签目录常驻浏览器"
    - "直接沉入领域/或成为/项目"
    - "因为以后可能会用就放 Zen 工具架"
  route_reason: "这些入口不再常驻浏览器，但已在 GitHub 留下真实 URL。进入具体项目、课程或比较任务时，从这里挑选并写入对应项目；实际被某个课程、研究或工程采用后，再沉到对应的领域或成为/项目。"
  reuse_rule: "项目和学习入口按当前项目/课程激活；项目未激活时停留链接层。采用、评估、比较、排除或依赖证据出现后，才进入项目或领域。"
  evidence_locator: "资源/链接/浏览器项目与学习入口候选.md#说明; #居留说明; 资源/链接/_index.md#v0总判断"
  next_action: "按项目触发。"
```

---

## 7. 账号与服务入口为什么留链接层

```yaml
- source_id: "浏览器账号与服务入口 43 条"
  current_route: link_layer_account_and_service_entries
  decision_status: recovered_from_pr
  selected_reason: "这些是账号、本地服务、私人用途、邮箱、服务后台、网络服务或待拆分服务 / 知识属性的外部入口。记录站点地址有现实抵达价值。"
  rejected_alternatives:
    - "因为账号属性而从仓库排除"
    - "保存密码、验证码、令牌、Cookie、会话 ID、临时回跳或一次性授权参数"
    - "把账号入口直接沉入项目"
  route_reason: "默认居留地是资源/链接，只保存清理后的站点地址。只有某个入口实际绑定具体项目动作时，才在相应成为/项目中引用。成人、AI 角色、网络、VPN、邮箱、服务后台等入口不因类型敏感就擅自删除，但必须遵守凭证和会话边界。"
  reuse_rule: "账号入口可以保存门牌，不能保存秘密。项目绑定后引用，不提前迁移。"
  evidence_locator: "资源/链接/浏览器账号与服务入口.md#开头; #居留说明; 资源/语境/浏览器与链接路由.md#账号与会话边界"
  next_action: "保持链接层。"
```

---

## 8. 主题包候选为什么不是领域知识

```yaml
- source_id: "浏览器主题包候选 37 条"
  current_route: link_layer_topic_pack_candidates
  decision_status: recovered_from_pr
  selected_reason: "包含百科与文化、学术研究、安全与隐私、技术标准与数据等门牌，如 Internet Archive、SEP、MedlinePlus、Project Gutenberg、ACM、Semantic Scholar、arXiv、NVD、OWASP、MDN、国家数据等。"
  rejected_alternatives:
    - "直接进入领域"
    - "直接成为正式资源/链接.csv"
    - "把同主题链接当成完整知识结构"
  route_reason: "这些按检索角色分组，但尚未成为领域知识，也没有正式进入资源/链接.csv。只保存门牌和候选状态时继续留在资源/链接；实际读过并形成可复用内容时进资源/萃览；知道何时、为何使用时进资源/语境；被长期领域反复调用时才进领域；影响具体项目动作时才进成为/项目。"
  reuse_rule: "主题包只是入口集合，不是知识沉淀。迁居必须等待使用证据。"
  evidence_locator: "资源/链接/浏览器主题包候选.md#开头; #后续迁居门槛"
  next_action: "按复核或任务使用。"
```

---

## 9. 待复核、一次性入口和垂直检索

```yaml
- source_id: "待复核与一次性入口 71 条"
  current_route: link_only_low_value_needs_review_or_one_off
  decision_status: recovered_from_pr
  selected_reason: "这些网站仍可能有单次用途、低价值候选或需补证据价值，例如部分词典、百科、AI 检测器、文库、历史上的今天、节假日、Dify 文档、Flowith 社区、微信读书、学习强国、鸠摩搜索等。"
  rejected_alternatives:
    - "迁入资源/雷达"
    - "自动迁入领域或项目"
    - "把候选状态冒充迁移结论"
    - "把鸠摩搜索放入垂直检索重复计数"
  route_reason: "它们当前是 link_only、low_value_candidate、needs_review 或一次性入口。复核后可以继续留在链接层、进入语境、沉入项目，或只保留历史；当前不能把候选状态冒充成迁移结论。鸠摩搜索属于后续复核，不在垂直检索文件重复计数。"
  reuse_rule: "一次性入口用完可回历史；角色未稳定就继续链接层；需要补证据的条目不迁居。"
  evidence_locator: "资源/链接/浏览器待复核与一次性入口.md#开头; #居留说明; 资源/链接/浏览器垂直检索候选.md"
  next_action: "按需要复核。"

- source_id: "浏览器垂直检索候选 2 条"
  current_route: link_layer_vertical_search_candidates
  decision_status: recovered_from_pr
  selected_reason: "问问小宇宙可作为中文播客内容检索候选；Dexa 可作为英文播客与专家内容检索候选。"
  rejected_alternatives:
    - "直接沉入资源/语境"
    - "与鸠摩搜索重复计数"
  route_reason: "只有形成稳定使用语境后，才考虑沉入资源/语境；影响具体研究或项目动作后，再在相应项目中引用。鸠摩仍在待复核与一次性入口中。"
  reuse_rule: "检索工具只有反复解决同类问题后才进入语境；一次发现不构成工具架资格。"
  evidence_locator: "资源/链接/浏览器垂直检索候选.md"
  next_action: "按检索任务观察。"
```

---

## 10. docs/ops 为什么保留执行证据，而不是资源入口

```yaml
- source_id: "浏览器清理与迁移 docs/ops 文件"
  current_route: operations_audit_evidence
  decision_status: recovered_from_pr
  selected_reason: "这些文件记录一次性执行、审计、对账、界面核验和迁移证据。"
  rejected_alternatives:
    - "把执行记录放进资源/雷达"
    - "把自动化日志当完成依据"
    - "把磁盘文件和节点计数覆盖用户界面事实"
  route_reason: "完成标准是用户实际界面符合目标、重启后状态不恢复、不误伤标签组/密码/Cookies/历史/扩展/非目标数据、GitHub 记录与界面一致。磁盘文件、自动化日志和节点计数只能作为辅助证据。"
  reuse_rule: "执行证据住 docs/ops；长期入口住资源/链接或资源/语境；完成依据优先用户可见界面。"
  evidence_locator: "docs/ops/browser-bookmark-cleanup-2026-07-05.md#完成标准; docs/ops/browser-link-placement-audit-2026-07-05.md#结论"
  next_action: "保持。"
```

---

## 11. 本补充 ledger 的复用规则

1. 浏览器保存抵达，Folo / TopHub 保存到来，BibiGPT 保存取水，资源/链接保存门牌，资源/语境保存怎么用，领域保存长期能力，成为/项目保存已影响行动的证据，docs/ops 保存操作事实。
2. Zen 是主工作环境；Chrome 是 Chromium / Google / 公共服务环境；Safari 是干净隔离环境。
3. 主浏览器不接收旧历史书签树；只保留真实反复使用入口。
4. 工具架只收已形成稳定使用语境的工具；同一功能默认只保留一个入口，除非功能差异明确。
5. 账号入口可保存门牌，不保存秘密、Cookie、令牌、会话或一次性授权参数。
6. 订阅候选不是订阅事实；2026-07-17 先复查现有 Folo，再按一进一出判断。
7. 项目与学习入口只有项目 / 课程激活后才进入领域或项目。
8. 主题包是门牌集合，不是知识结构。
9. 待复核和一次性入口不能冒充迁移结论。
10. 浏览器导出 cohort 不替代 canonical 链接记录；同名或同 URL 复用既有正式条目。
11. 不得自动删除、自动合并、把可访问当有价值、把居留建议当迁移命令、把近重复当可合并、把同领域邻居放进去重流程。

---

## 12. 仍需继续追回

浏览器与链接路由已补成 ledger。

继续待办：

- 等用户本地提交 master checklist 与 coverage audit 后，再统一登记本文件和此前新增 ledger；
- 报刊目录专项 ledger；
- 娱乐目录大量 closure；
- 旧 ChatGPT 对话追索所有 partial 条目。
