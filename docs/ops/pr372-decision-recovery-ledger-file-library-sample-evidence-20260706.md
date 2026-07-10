# PR #372 判断链追回：File Library 样本级证据补充（2026-07-06）

## 0. 边界

本文件接续最近 7 天 ChatGPT 对话主动追索。

本轮通过 File Library 检索到三类可用证据：

- `粘贴的文本 (1).txt`
- `浏览器链接治理_第一轮候选.csv`
- `订阅源账本.public-candidate-v0.yml`

其中：

- `粘贴的文本 (1).txt` 提供任务边界与验收条件；
- `浏览器链接治理_第一轮候选.csv` 提供浏览器旧链接的逐行去向、理由、置信度和是否查链；
- `订阅源账本.public-candidate-v0.yml` 提供 Folo 候选账本中的 `state_candidate`、`route_candidate`、`why_listen_or_adjust` 和 `decision_status`。

本文件只抽取文件中可见样本，不宣称覆盖全部浏览器链接或全部 Folo 候选源。

当前仍不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 总规则：File Library 样本可以补证据，但不能自动补全全部条目

```yaml
- source_id: "File Library recent evidence samples"
  current_route: sample_level_evidence_not_full_completion
  decision_status: recovered_from_file_library_samples
  selected_reason: "File Library 检索返回了浏览器第一轮候选、公共候选订阅账本和上传粘贴文本，里面有逐行去向、候选状态和任务边界。"
  rejected_alternatives:
    - "把样本级证据说成全量完成"
    - "把候选账本说成真实 Folo 状态"
    - "把浏览器候选 CSV 的所有条目自动迁入正式资源"
    - "把上传文本里的方案建议当当前执行命令"
  route_reason: "这些文件能补充具体样本的 why / route / status，但不是完整旧对话原文，也不是所有来源的完整判断链。"
  reuse_rule: "只抽文件可见条目；未显示或未说明的来源继续标 `needs_item_level_reason`。"
  evidence_locator: "file_library: 粘贴的文本 (1).txt; 浏览器链接治理_第一轮候选.csv; 订阅源账本.public-candidate-v0.yml"
  next_action: "作为样本级证据层。"
```

---

## 2. 账号与邮箱入口：本地私密，不进公开 GitHub

```yaml
- source_id: "北京邮电大学邮箱 / QQ邮箱 / 163邮箱 / 中国移动139邮箱 / 189邮箱"
  current_route: local_private_not_public_github
  decision_status: recovered_from_file_library_sample
  selected_reason: "浏览器第一轮候选表将这些邮箱入口标为 `本地私密／不进公开 GitHub`。"
  rejected_alternatives:
    - "进入公开资源库"
    - "进入 TopHub"
    - "进入 Folo"
    - "作为账号服务公开沉积"
  route_reason: "这些入口涉及账号、登录、私人用途或敏感边界；只在确有需要的本地环境保留。"
  reuse_rule: "账号、邮箱、登录入口默认本地私密，不进入公开 GitHub；只有无敏感的服务门牌才可留 `资源/链接/`。"
  evidence_locator: "file_library: 浏览器链接治理_第一轮候选.csv rows 0-5"
  next_action: "本地保留。"
```

---

## 3. 账号协作入口：本地入口复核，不等于知识资源

```yaml
- source_id: "Gmail / 石墨文档 / Slack / 印象笔记"
  current_route: local_entry_review_not_public_resource
  decision_status: recovered_from_file_library_sample
  selected_reason: "浏览器第一轮候选表将这些账号和协作入口标为 `本地入口复核`。"
  rejected_alternatives:
    - "直接进入公开资源库"
    - "作为知识资源沉积"
    - "作为 Folo / TopHub 来源"
  route_reason: "账号和协作入口是否保留取决于当前真实使用，不进入公开资源库。"
  reuse_rule: "协作工具入口先看当前是否真实使用；本地入口和知识资源分开。"
  evidence_locator: "file_library: 浏览器链接治理_第一轮候选.csv rows 6-10"
  next_action: "本地入口复核。"
```

---

## 4. 已有入口：不重复收录

```yaml
- source_id: "Notion / Web3 Recruitment Platform"
  current_route: existing_entry_no_duplicate
  decision_status: recovered_from_file_library_sample
  selected_reason: "浏览器第一轮候选表将 Notion 和 Web3 Recruitment Platform 标为 `已有入口／不重复收录`。"
  rejected_alternatives:
    - "重复复制到书签树"
    - "重复写入沃壤链接库"
    - "因为重要就重复登记"
  route_reason: "浏览器或空间里已有明确入口，无需再复制到书签树或沃壤链接库。"
  reuse_rule: "已有 canonical 入口时，重复书签只作为去重证据，不新增正式资源。"
  evidence_locator: "file_library: 浏览器链接治理_第一轮候选.csv rows 8 and 21"
  next_action: "不重复收录。"
```

---

## 5. 旧学校 / 求职 / 阶段性身份入口：历史归档

```yaml
- source_id: "Sydney Uni Outlook / 国家大学生就业服务平台 / BUPT 就业资讯网 / BOSS直聘 / 智联招聘 / 猎聘 / 前程无忧 / 拉勾招聘 / 58同城 / Sydney Uni CareerHub / Sydney Uni Career Centre"
  current_route: historical_archive_not_active_bookmarks
  decision_status: recovered_from_file_library_sample
  selected_reason: "浏览器第一轮候选表将这些学校、就业、求职和学生权益入口标为 `历史归档`。"
  rejected_alternatives:
    - "继续留活动书签"
    - "进入 PR #372 source ledger"
    - "作为当前项目入口"
  route_reason: "它们属于旧学校、求职、学生权益或阶段性身份入口；先退出活动书签，原始导出保留历史。"
  reuse_rule: "阶段性身份入口默认退出活动层；只有当前身份或任务恢复时再重新激活。"
  evidence_locator: "file_library: 浏览器链接治理_第一轮候选.csv rows 1 and 11-20"
  next_action: "历史归档。"
```

---

## 6. 软件资源与商店：多数历史归档，不常驻，不替代官方来源

```yaml
- source_id: "反斗限免 / 果粉GoFans / 老师帮 / NEXT ITELLYOU / Greasy Fork / 数码荔枝 / MacPaw / iPA商店 / Decrypt IPA Store / iOSvizor / 麦氪派 / 哦游Max / Pirate Bay Proxy / FitGirl Repacks / 果核剥壳 / Xclient.info / MacWk / 马可菠萝 / MacZ"
  current_route: historical_archive_or_task_search_not_subscription
  decision_status: recovered_from_file_library_sample
  selected_reason: "浏览器第一轮候选表把多数软件资源、下载、商店、导航集合标为历史归档。"
  rejected_alternatives:
    - "作为浏览器常驻"
    - "进入 Folo / TopHub"
    - "进入公开正式资源库"
    - "作为软件下载默认入口"
  route_reason: "这些入口多属于工具 / 下载 / 导航集合，主要表达‘有很多选项’；软件资源与商店类入口不应常驻，需要时搜索官方来源。原始导出保留历史已足够。"
  reuse_rule: "软件下载、资源集合、商店导航默认不常驻；需要软件时回官方源、App Store、项目主页或可信分发。"
  evidence_locator: "file_library: 浏览器链接治理_第一轮候选.csv rows 22 and 24-43"
  next_action: "历史归档。"
```

---

## 7. IT之家「喜加一」与少数派商城：持续更新候选，但仍不是浏览器日常入口

```yaml
- source_id: "IT之家「喜加一」 / 少数派商城"
  current_route: folo_tophub_candidate_action_trigger
  decision_status: recovered_from_file_library_sample_and_browser_candidate_file
  selected_reason: "浏览器第一轮候选表把它们标为 `Folo／TopHub 候选`，理由是持续更新的来源应由订阅系统承担到来，不在浏览器长期堆媒体首页。"
  rejected_alternatives:
    - "浏览器长期保留为媒体首页"
    - "直接写入正式订阅账本"
    - "当成普通深读来源"
  route_reason: "它们虽然持续更新，但属于促销与行动触发。应该进入复查候选，而不是活动书签或正式订阅事实。"
  reuse_rule: "持续更新的促销 / 商城入口若保留，只能作为动作触发候选；复查时再决定是否有必要进入 TopHub / Folo。"
  evidence_locator: "file_library: 浏览器链接治理_第一轮候选.csv rows 23 and 29; 资源/链接/浏览器订阅候选-2026-07-17.md"
  next_action: "2026-07-17 复查或按任务触发。"
```

---

## 8. 标准 / 安全参考源：沃壤资源候选，不自动进正式资源库

```yaml
- source_id: "ISO / 国家标准全文公开 / CNVD / FIRST.org / Kali Tools / CyberChef / CyberChef(GCHQ)"
  current_route: worang_resource_candidates_need_reason_before_formal_library
  decision_status: recovered_from_file_library_sample
  selected_reason: "浏览器第一轮候选表把这些标准、安全和工具参考源标为 `沃壤资源候选`，理由是具有长期参考、方法、证据或项目燃料价值。"
  rejected_alternatives:
    - "直接进入正式资源库"
    - "进入 Folo / TopHub"
    - "因为权威就自动沉积"
  route_reason: "候选表明确还需要补‘为什么保留’后再进入资源库。权威性和工具价值只是候选条件，正式沉积仍需使用语境。"
  reuse_rule: "标准、安全、工具参考源先留资源候选；被项目、领域或排障实际调用后再沉正式语境或领域。"
  evidence_locator: "file_library: 浏览器链接治理_第一轮候选.csv rows 463-464 and 503-507"
  next_action: "补使用语境。"
```

---

## 9. 文库 / 搜索 / 工具混合入口：后续复核，不能整体迁移

```yaml
- source_id: "学兔兔 / Bookboon / Clark and Miller / 百度文库 / 360文库 / 道客巴巴 / CTF在线工具"
  current_route: needs_role_split_before_residence
  decision_status: recovered_from_file_library_sample
  selected_reason: "浏览器第一轮候选表将这些旧文件夹内的入口标为后续复核或本地 / 沃壤复核。"
  rejected_alternatives:
    - "按旧文件夹整体迁移"
    - "全部当知识资源"
    - "全部删除"
  route_reason: "旧文件夹混合了媒体、百科、搜索、数据与工具，需按角色拆分，不能整体迁移；官方安全知识可进入沃壤，账号、代理与服务入口只留本地。"
  reuse_rule: "混合旧文件夹先拆角色：资料库、工具、服务、账号、搜索、内容。角色不明不迁居。"
  evidence_locator: "file_library: 浏览器链接治理_第一轮候选.csv rows around 386-392 and 508"
  next_action: "后续复核。"
```

---

## 10. Folo 公共候选账本：候选账本不自动改变 Folo

```yaml
- source_id: "订阅源账本.public-candidate-v0.yml"
  current_route: public_safe_candidate_ledger_not_real_subscription_state
  decision_status: recovered_from_file_library_sample
  selected_reason: "公共候选账本原则明确：这是候选账本，不自动改变 Folo 订阅；私人兴趣与完整关注关系不进入公开仓库；先记录来源关系，再决定是否常驻；RSSHub 路由替换需单独核验。"
  rejected_alternatives:
    - "把候选账本当真实 Folo 状态"
    - "公开完整私人订阅关系"
    - "凭猜测修改 RSSHub 路由"
    - "记录来源后自动常驻"
  route_reason: "该文件是 public-safe candidate ledger，负责候选关系，不是最终真实订阅账本。"
  reuse_rule: "公开仓库只存筛选后的公共候选；真实 Folo 状态另看承重账本，候选变更必须核验。"
  evidence_locator: "file_library: 订阅源账本.public-candidate-v0.yml principles"
  next_action: "作为候选证据。"
```

---

## 11. Folo 候选样本：视觉、科技、论文、AI 官方、深读、公共雷达各有不同路由

```yaml
- source_id: "中国爬楼联盟 / 少数派 / InfoQ 推荐 / 领研论文计算机 / Logan Kilpatrick / Google DeepMind News / 晚点长报道 / Last Week in AI / Hacker News"
  current_route: folo_candidate_samples_with_distinct_attention_modes
  decision_status: recovered_from_file_library_sample
  selected_reason: "公共候选账本样本显示，不同来源被赋予不同 attention_mode_candidate、source_role、state_candidate 和 route_candidate。"
  rejected_alternatives:
    - "把所有候选都按同一规则处理"
    - "所有科技源都进入 Folo"
    - "所有公共雷达都形成未读"
  route_reason: "中国爬楼联盟是视觉社区 trial，先看实际打开频率；少数派是科技媒体 active，但更新量可能偏高，需核验官方直连；InfoQ 推荐适合了解行业面，转 TopHub / 按需，不必形成逐篇未读；领研论文聚合范围较宽，适合明确研究问题出现时调用；Logan 是 AI 平台动态个人 / 行业信号，需看是否重复官方公告；Google DeepMind News 低频且有一手信息价值；晚点长报道适合阅读橱窗，不按未读数追全；Last Week in AI 是低频周报；Hacker News 是高频公共技术雷达，不适合形成持续未读。"
  reuse_rule: "Folo 候选要按注意力模式和角色分开：漫游、深读、公共雷达、观察期、触发器、官方源、周报、社区聚合不能同一处理。"
  evidence_locator: "file_library: 订阅源账本.public-candidate-v0.yml sample entries"
  next_action: "按各自 review_wave 复查。"
```

---

## 12. 金十 / 安全媒体样本：高频噪声、专业深读和重复检查要分开

```yaml
- source_id: "金十数据 / Dark Reading / Krebs on Security / Cyber Security News"
  current_route: candidate_samples_finance_noise_vs_security_depth
  decision_status: recovered_from_file_library_sample
  selected_reason: "公共候选账本样本显示：金十数据为金融快讯，on_demand；Dark Reading 为网安媒体 active；Krebs on Security 为个人安全博客 active；Cyber Security News 为 standby。"
  rejected_alternatives:
    - "把金十数据作为持续未读"
    - "把所有安全媒体都保留不裁剪"
    - "忽略 Cyber Security News 与 Dark Reading 的重复"
  route_reason: "金十高频且易制造噪声，只有金融现场活跃时再启用；Dark Reading 是网安领域专业媒体；Krebs 是独立作者、一手调查价值高；Cyber Security News 可能与 Dark Reading 重复，需看实际增量。"
  reuse_rule: "快讯源按现场需要启用；安全深读源看专业性、一手性和重复度。"
  evidence_locator: "file_library: 订阅源账本.public-candidate-v0.yml sample entries around jin10 / security sources"
  next_action: "按需 / 复查重复度。"
```

---

## 13. 本补充 ledger 的复用规则

1. File Library 样本能补具体证据，但不能宣布全量完成。
2. 账号、邮箱和登录入口默认本地私密，不进公开 GitHub。
3. 协作入口按当前真实使用复核，不当知识资源。
4. 已有 canonical 入口不重复收录。
5. 旧学校、求职、学生权益入口进入历史归档。
6. 软件资源、下载集合、商店导航不常驻，任务时回官方源或可信分发。
7. 持续更新的促销 / 商城入口只是动作触发候选，不是深读源。
8. 标准、安全、工具参考源先做沃壤资源候选，补使用语境后再沉积。
9. 混合旧文件夹先拆角色，不能整体迁移。
10. 公共候选账本不自动改变 Folo。
11. Folo 候选必须按注意力模式和来源角色分开，不同 wave 不同处理。
12. 高频快讯、专业安全媒体、一手作者和重复安全媒体要分开判断。

---

## 14. 仍需继续追回

File Library 样本级证据已补一层。

继续待办：

- 若要继续逐项抽浏览器第一轮 CSV，需要分批读取更多行，不能凭当前样本扩写；
- 若要逐项抽 Folo public candidate，需要按 id 分段读取完整文件；
- TopHub 早期首轮 28 个取消项原始列表仍未找到；
- App 六页 closure 原始文件仍未找到；
- 仍不更新用户本地待提交的 `master-checklist`、`coverage-audit`、军事边界 ledger。
