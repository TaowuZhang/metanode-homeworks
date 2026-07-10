# PR #372 判断链追回：科技当前雷达来源（少数派 / IT之家 / iThome / 极客公园 / TechCrunch / The Verge）（2026-07-06）

## 0. 边界

本文件是 `科技` 大类的补充 ledger，追回当前科技雷达中几类仍未充分落账的现有 / 替换来源判断。

主要证据文件：

- `docs/ops/tophub-technology-sspai-closure-20260704.md`
- `docs/ops/tophub-technology-ithome-closure-20260704.md`
- `docs/ops/tophub-technology-general-tech-closure-20260704.md`
- `docs/ops/pr372-decision-recovery-source-index-20260706.md`
- `资源/链接/极客公园.md`
- `资源/链接/TechCrunch.md`
- `资源/链接/The Verge.md`

注意：

- `少数派` 与 `IT之家 / iThome` 有专门 closure，可追回较完整判断链。
- `极客公园｜每日最新`、`TechCrunch｜Today`、`The Verge｜Today` 在 source index 中是当前科技雷达节点，但本轮没有找到对应专门 closure；只能从 source index、链接卡片和 general-tech 中的比较性证据追回部分判断，不补编完整历史理由。
- 当前不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

判断链标准：为什么选它、为什么不选旁边那个、为什么放这个层、以后遇到同类来源怎样复用。

---

## 1. 当前科技雷达来源总规则

```yaml
- source_id: "当前科技雷达来源总规则"
  current_route: tophub_technology_radar
  decision_status: recovered_from_pr_partial
  selected_reason: "科技雷达不是单一科技媒体列表，而是由几种不同角色组成：中文消费科技、个人工具与独立开发、国际创业与平台公司、国际消费电子与数字文化、企业 IT / 安全基础设施、公共科普、开源项目发现和官方行动入口。"
  rejected_alternatives:
    - "每家科技媒体都保留全站热榜"
    - "同站多个栏目并存"
    - "把作者型来源压成 TopHub 标题流"
    - "把高频科技早报、最新流和热评榜加入 Folo"
  route_reason: "不同来源要补不同层，而不是按品牌名堆叠。TopHub 承担标题发现和横向扫描，Folo 承担连续作者 / 编辑关系。"
  reuse_rule: "以后遇到现有科技雷达调整，先写清它补的角色：消费科技、独立工具、创业融资、数字文化、企业基础设施、科学解释、开源项目、官方行动。不能只说‘科技媒体’。"
  evidence_locator: "docs/ops/pr372-decision-recovery-source-index-20260706.md#22科技雷达1728; technology closure files"
  next_action: "极客公园、TechCrunch、The Verge 仍需未来从旧对话或专门 closure 追回更完整理由。"
```

---

## 2. 少数派：从全站热榜转向 Matrix 热榜

```yaml
- source_id: "少数派｜热门文章"
  current_route: retired_from_tophub_pending_platform_execution
  decision_status: recovered_from_pr
  selected_reason: "原本作为少数派入口，能看到 App、设备、系统、生活方式、数字工具和编辑内容。"
  rejected_alternatives:
    - "继续保留少数派全站热门文章"
    - "少数派｜最新文章"
    - "少数派｜Matrix 最新流"
    - "少数派｜最新上架付费专栏"
    - "少数派｜一派话题广场"
    - "少数派｜#派评"
    - "少数派｜#应用推荐"
    - "少数派｜#派早报"
  route_reason: "热门文章已经扩展为少数派全站生活方式精选，混合 App、影视、旅行、购物、Apple 手记、数码设备、系统测试版、英语学习和消费内容。它与影游音乐、生活与社区、慢读与思想交叉，不能稳定补出科技雷达缺口。"
  reuse_rule: "同一社区型媒体要选最独特的层，而不是保留全站热榜。全站热榜一旦生活方式化，就应退为按需或替换。"
  evidence_locator: "docs/ops/tophub-technology-sspai-closure-20260704.md#一为什么不继续保留少数派热门文章"
  next_action: "由 Matrix 热榜替换。"

- source_id: "少数派｜Matrix热榜"
  current_route: tophub_technology_radar
  decision_status: recovered_from_pr
  selected_reason: "补出少数派最独特的部分：独立开发者、个人工具、本地软件、小型产品、工作流、UX、数字生活实践和从真实需求出发的开发过程。"
  rejected_alternatives:
    - "Matrix 最新流"
    - "#应用推荐"
    - "#派评"
    - "派早报"
    - "最新文章"
  route_reason: "仍放科技雷达，替换而非扩容。它与 HelloGitHub 互补：HelloGitHub 是编辑整理的开源项目发现，Matrix 热榜是个人工具、产品实践和开发过程。选择热榜而非 Matrix 最新流，是为了降低自荐、刚发布小工具和生活随笔噪声。"
  reuse_rule: "社区型工具来源如果要常驻，应选择经过筛选的实践入口，而不是最新流、标签流或商业上新。看到工具不等于安装，不自动转项目任务。"
  evidence_locator: "docs/ops/tophub-technology-sspai-closure-20260704.md#二为什么选择Matrix热榜; #十少数派在系统中的正确角色"
  next_action: "已进入科技大类收口建议。"

- source_id: "少数派按需节点"
  current_route: on_demand_sspai_pages
  decision_status: recovered_from_pr
  selected_reason: "#派评、#应用推荐和话题广场在 App 发现、社区问题和数字生活观察中有价值。"
  rejected_alternatives:
    - "常驻科技雷达"
    - "Folo"
  route_reason: "#派评标题难以判断本期具体 App，#应用推荐重复严重并易制造 App 发现焦虑，话题广场是讨论与投票入口，不是稳定内容源。"
  reuse_rule: "App、工具、扩展、知识管理软件按具体问题打开少数派按需页，不建立持续未读。"
  evidence_locator: "docs/ops/tophub-technology-sspai-closure-20260704.md#七少数派派评; #八少数派应用推荐; #六少数派一派话题广场"
  next_action: "按需。"
```

---

## 3. IT之家与 iThome：中文消费科技和企业基础设施分层

```yaml
- source_id: "IT之家｜日榜"
  current_route: tophub_technology_radar
  decision_status: recovered_from_pr
  selected_reason: "承担中文消费科技市场的快速、高密度、跨品牌产品雷达，覆盖设备、系统、软件支持、国产硬件、AI 产品、汽车与中文用户讨论。"
  rejected_alternatives:
    - "IT之家｜最新更新"
    - "IT之家｜7天热评"
    - "IT之家｜周榜"
    - "IT之家｜月榜"
    - "IT之家｜IT资讯"
    - "IT之家｜智能汽车"
  route_reason: "继续放科技雷达。日榜虽然有传闻、品牌争议和汽车比例高的问题，但 The Verge 偏国际、爱范儿偏产品叙事、极客公园偏公司和产业、Readhub 偏快讯压缩；IT之家在国产设备、系统版本、设备支持和中文用户讨论上仍有独特覆盖。"
  reuse_rule: "同站时间窗口和评论榜不应并存；保留最能压缩中文消费科技变化的一条。"
  evidence_locator: "docs/ops/tophub-technology-ithome-closure-20260704.md#一为什么继续保留IT之家日榜"
  next_action: "无；保留。"

- source_id: "iThome｜新闻"
  current_route: tophub_technology_radar
  decision_status: recovered_from_pr
  selected_reason: "补企业 IT、资安事件、漏洞响应、云原生、CI/CD、安全基础设施、供应链攻击、凭证窃取、企业 AI 与技术治理。"
  rejected_alternatives:
    - "IT之家｜IT资讯"
    - "Office之家热榜"
    - "The Register 作为唯一企业 IT 入口"
    - "直接进入 Folo"
  route_reason: "放科技雷达，不是公共温度，也不是数据与结构。它不是大陆 IT之家的另一栏目，而是独立媒体。它承担专业技术雷达；只有与用户实际软件、仓库、云服务或项目相关时，才进入项目处理。"
  reuse_rule: "风险与漏洞雷达不能自动制造待办；具体 CVE、产品、版本、供应链事件与用户实际项目相关时，才临时追踪。"
  evidence_locator: "docs/ops/tophub-technology-ithome-closure-20260704.md#二为什么新增iThome新闻"
  next_action: "已进入科技大类收口建议。"

- source_id: "IT之家其余节点"
  current_route: rejected_or_on_demand_ithome_slices
  decision_status: recovered_from_pr
  selected_reason: "Office之家和喜加一有具体任务价值。"
  rejected_alternatives:
    - "最新更新常驻"
    - "热评 / 周榜 / 月榜常驻"
    - "智能汽车常驻"
    - "IT资讯常驻"
    - "喜加一进入 Folo"
  route_reason: "最新更新太高频；热评、周榜、月榜放大品牌争议和历史热度；智能汽车偏新车和品牌公关；IT资讯边界不如 iThome 清楚；喜加一是行动提醒，不是科技信息，等待 Folo 行动提醒复查。"
  reuse_rule: "同站榜单只保留日榜；细分任务页按具体 Office、汽车、游戏领取任务打开。"
  evidence_locator: "docs/ops/tophub-technology-ithome-closure-20260704.md#三IT之家最新更新; #四IT之家7天热评; #十IT之家喜加一"
  next_action: "按需或后续 Folo 复查。"
```

---

## 4. 极客公园：当前科技雷达节点，但判断链仍不完整

```yaml
- source_id: "极客公园｜每日最新"
  current_route: tophub_technology_radar
  decision_status: partial_recovered_from_pr
  selected_reason: "source index 记录它是当前科技雷达第 17 个节点；链接卡片记录其为可复用科技媒体入口。general-tech closure 中也把它作为现有科技雷达比较对象，暗示它承担中文科技公司、产品、产业和新技术叙事入口。"
  rejected_alternatives:
    - "用当前模型补编完整历史理由"
    - "直接判定已完整追回"
  route_reason: "暂时保留为当前科技雷达，但完整判断链未在本轮找到对应 closure。可确认的只是它与 IT之家、爱范儿、The Verge、Readhub 等共同构成现有科技雷达比较系。"
  reuse_rule: "对没有专门 closure 的现有节点，只登记已证实的 route 和缺口；不得用泛泛‘科技媒体’补全。后续应回旧对话或新增专门 closure。"
  evidence_locator: "docs/ops/pr372-decision-recovery-source-index-20260706.md#22科技雷达1728; 资源/链接/极客公园.md; docs/ops/tophub-technology-ithome-closure-20260704.md#一为什么继续保留IT之家日榜"
  next_action: "blocked_partial：需要旧对话或专门极客公园 closure。"
```

---

## 5. TechCrunch 与 The Verge：当前国际科技媒体基线，但仍需补完整来源链

```yaml
- source_id: "TechCrunch｜Today"
  current_route: tophub_technology_radar
  decision_status: partial_recovered_from_pr
  selected_reason: "source index 记录它是当前科技雷达第 19 个节点；链接卡片记录 TechCrunch 是美国科技创投媒体，以创业公司报道、融资新闻和产品评测著称。general-tech 与 IT之家 closure 多次把 TechCrunch 作为现有国际科技和创业融资叙事来源比较对象。"
  rejected_alternatives:
    - "把 TechCrunch 苹果切片 / 其他切片重复加入"
    - "把 TechCrunch 当企业基础设施源替代 The Register / iThome"
    - "直接判定已完整追回"
  route_reason: "暂时保留为科技雷达中的国际创业、平台公司和科技商业变化入口。它不是 The Register 的替代品；The Register 补企业系统和基础设施，TechCrunch 更靠近创业、产品、融资和平台公司。"
  reuse_rule: "国际科技媒体要拆角色：TechCrunch 偏创业和平台公司，The Verge 偏消费电子和数字文化，The Register 偏企业基础设施。不能因为都是海外科技媒体就互相替代。"
  evidence_locator: "docs/ops/pr372-decision-recovery-source-index-20260706.md#22科技雷达1728; 资源/链接/TechCrunch.md; docs/ops/tophub-technology-general-tech-closure-20260704.md#一最值得进入科技大类收口的节点"
  next_action: "blocked_partial：需要旧对话或专门 TechCrunch closure。"

- source_id: "The Verge｜Today"
  current_route: tophub_technology_radar
  decision_status: partial_recovered_from_pr
  selected_reason: "source index 记录它是当前科技雷达第 20 个节点；链接卡片记录 The Verge 是美国主流科技媒体，报道消费电子、互联网产品与数字文化。general-tech closure 明确不重复新增 The Verge 站内科技切片，因为系统已订阅 Today。"
  rejected_alternatives:
    - "The Verge｜Teches / Sciences 等站内切片"
    - "把 The Verge 当企业基础设施源替代 The Register"
    - "把 The Verge 与 TechCrunch / The Register 混作同一种海外科技媒体"
  route_reason: "暂时保留为国际消费电子、互联网产品和数字文化入口。The Verge Today 是横向入口；站内科技 / 科学切片不新增，避免同源切片重复。"
  reuse_rule: "同一海外媒体已有横向 Today / Front Page 入口时，不再订阅站内科技、科学、Apple 等切片；除非切片补出完全不同功能。"
  evidence_locator: "docs/ops/pr372-decision-recovery-source-index-20260706.md#22科技雷达1728; 资源/链接/The Verge.md; docs/ops/tophub-technology-general-tech-closure-20260704.md#五海外综合科技媒体; docs/ops/tophub-technology-science-closure-20260704.md#The VergeSciences"
  next_action: "blocked_partial：需要旧对话或专门 The Verge closure。"
```

---

## 6. 本补充 ledger 的复用规则

1. 少数派的正确角色是个人工具、独立开发和数字工作流，不是全站生活方式热榜。
2. IT之家日榜承担中文消费科技压缩入口；评论榜、最新流、周月榜不并存。
3. iThome 不是 IT之家同源切片，而是企业 IT / 安全基础设施专业雷达；但风险标题不自动变成待办。
4. 极客公园、TechCrunch、The Verge 当前仍是科技雷达基线，但本轮没有找到专门 closure，必须标 partial，不得补编完整历史链。
5. 国际科技媒体分角色：TechCrunch 偏创业和平台公司，The Verge 偏消费电子和数字文化，The Register 偏企业基础设施。
6. 同一媒体已有横向入口时，不重复订阅站内科技、科学、Apple、产品等切片。

---

## 7. 仍需继续追回

科技大类当前已补入：

- 商业叙事与全球化来源；
- 科学、科普、企业基础设施、Apple、汽车、报告与设备市场结构；
- 当前科技雷达来源中的少数派、IT之家 / iThome，以及极客公园 / TechCrunch / The Verge 的 partial 边界。

继续待办：

- AI、App、快讯、限免、教育、电商、数码、酷安、Google、TechWeb 等 closure；
- HelloGitHub 与开发目录之间的关系；
- 极客公园、TechCrunch、The Verge 的完整判断链需要旧对话或专门 closure 追回；
- 后续等用户本地提交 master checklist 与 coverage audit 后，再登记本文件。
