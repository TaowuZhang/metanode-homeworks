# PR #372 判断链追回：原始文件反向覆盖映射（2026-07-06）

## 0. 边界

本文件不是最终完成证明，不替代：

- `docs/ops/pr372-decision-recovery-master-checklist-20260706.md`
- `docs/ops/pr372-decision-recovery-coverage-audit-20260706.md`
- `docs/ops/pr372-decision-recovery-source-index-20260706.md`

本文件只做一件事：从 PR #372 的原始变更文件反向映射到已新增的判断链 ledger，防止两种错误：

1. 在已经沉过判断链的原始文件里反复挖；
2. 把只被总规则覆盖、仍为 partial 或只在用户本地待提交的内容说成已经完整完成。

当前仍不更新 `master-checklist` 与 `coverage-audit`，避免与用户本地待提交内容冲突。

---

## 1. 反向覆盖判准

```yaml
- source_id: "PR #372 原始文件反向覆盖"
  current_route: coverage_map_not_completion_claim
  decision_status: recovered_from_pr
  selected_reason: "PR #372 变更文件数量很大，且包含原始审计文件、平台执行证据、候选清单、控制文件、资源入口和后续新增 ledger。需要一张反向图说明哪些原始文件已经被判断链 ledger 承接。"
  rejected_alternatives:
    - "只统计文件数量"
    - "把所有原始文件都视为已完整迁移"
    - "把已有 ledger 的存在当作逐源判断完成"
    - "把本地未提交的控制文件远程改写"
  route_reason: "覆盖映射只说明‘这个原始文件由哪个 ledger 承接’，不说明其中每个来源都已完整追回。完成状态仍由具体 ledger、source index、coverage audit 与 master checklist 判断。"
  reuse_rule: "以后继续追回时，先查反向覆盖：若原始文件已由专项 ledger 承接，不重复；若只被桥接或总规则覆盖，再判断是否需要追加逐源 ledger。"
  evidence_locator: "PR #372 changed filenames; existing pr372-decision-recovery-ledger-* files"
  next_action: "用作继续追回的导航，不作完成声明。"
```

---

## 2. 控制文件与工作台

| 原始文件组 | 当前承接 | 状态 | 说明 |
|---|---|---|---|
| `pr372-decision-recovery-todo-20260706.md` | 自身 | workbench | 工作台，不是判断账本。 |
| `pr372-decision-recovery-source-index-20260706.md` | 自身 | index_not_ledger | 登记 TopHub / Folo 来源和状态，不替代判断链。 |
| `pr372-decision-recovery-master-checklist-20260706.md` | 用户本地待提交版本 | do_not_remote_update | 不远程改，避免冲突。 |
| `pr372-decision-recovery-coverage-audit-20260706.md` | 用户本地待提交版本 | do_not_remote_update | 不远程改，避免冲突。 |
| `pr372-decision-recovery-ledger-draft-20260706.md` | 第一批 draft ledger | partial_seed | 早期草案，已被后续专项 ledger 扩展，不单独宣布完成。 |

复用规则：控制文件回答“还差什么”；ledger 回答“判断链是什么”。二者不能互相替代。

---

## 3. TopHub / Folo 系统层

| 原始文件组 | 已承接 ledger | 状态 | 说明 |
|---|---|---|---|
| `TopHub 能力地图.md`、`TopHub 使用配置.md`、`TopHub 全部订阅顺序.md` | `ledger-tophub-folo-system-routing`、`ledger-tophub-account-widgets-app-feature-boundary`、`ledger-tophub-overview-discovery-routing`、`ledger-pending-actions-platform-verification-boundary` | covered_with_time_boundary | 系统能力、分组、全部页、通知、小部件、当前 80 基线均已承接。 |
| `tophub-account-and-notification-audit`、`tophub-app-and-help-audit`、`tophub-widgets-and-remaining-settings-audit`、`tophub-homepage-audit`、`tophub-feature-audit` | `ledger-tophub-account-widgets-app-feature-boundary` | covered | 功能边界已追回：平台支持不等于启用。 |
| `tophub-all-page-order-audit`、`tophub-source-discovery-and-entertainment-audit`、`tophub-public-temperature-and-group-map`、`tophub-more-directories-audit` | `ledger-tophub-overview-discovery-routing` | covered | 来源发现层 / 个人路由层边界已追回。 |
| `folo-subscription-diff`、`folo-subscription-route-review-v2`、`tophub-folo-guided-review`、`tophub-folo-handoff-after-ai`、`subscription-system-closeout` | `ledger-folo-early-operation-chain`、`ledger-tophub-folo-system-routing` | covered | Folo 45→27、五类注意力用途、候选复查边界已追回。 |

---

## 4. 综合类

| 原始文件组 | 已承接 ledger | 状态 | 说明 |
|---|---|---|---|
| `tophub-comprehensive-closeout`、`tophub-comprehensive-zhihu-closure` | `ledger-comprehensive` | covered_partial | 综合 closeout 与知乎链已追回。 |
| 热搜与滚动：`hot-search`、`rolling-news`、`rolling-news-review` | `ledger-comprehensive-hotsearch-rolling` | covered | 热搜 / 滚动新闻不扩容、不进 Folo、不建宽泛追踪器。 |
| 平台注意力：`wechat`、`weibo`、`toutiao` | `ledger-comprehensive-platform-attention` | covered | 微信 / 微博 / 今日头条平台切片边界已追回。 |
| 门户聚合：`baidu`、`tencent`、`netease`、`sohu`、`sina`、`zaker` | `ledger-comprehensive-portal-aggregators` | covered | 门户、研究院、数读、人间、黑猫等路由已追回。 |
| 机构媒体：`people`、`xinhua`、`southern-weekly`、`banyuetan`、`cctv` | `ledger-comprehensive-institutional-media` | covered | 官方 / 机构媒体功能位已追回。 |
| 行动源：`health`、`real-estate`、`local-city`、`tencent-daily` 等 | `ledger-comprehensive-action-sources` | covered | 辟谣、健康、房产、本地宝、腾讯日报等边界已追回。 |
| 军事：`tophub-comprehensive-military-closure` | 用户本地军事 ledger | do_not_remote_update | 用户本地已完成并待提交，不远程改。 |

综合仍保持 partial 的原因不是这些 ledger 没有写，而是军事 ledger 控制文件更新在用户本地、且旧对话迁移仍未统一登记。

---

## 5. 科技 / AI / 开发

| 原始文件组 | 已承接 ledger | 状态 | 说明 |
|---|---|---|---|
| `tophub-technology-*` 全组 | `ledger-technology`、`technology-business-globalization`、`technology-science-infra-devices`、`technology-current-radar-sources`、`technology-task-triggered-and-high-frequency`、`technology-development-and-culture-boundary`、`final-closure-residual-bridges` | covered | 科技目录、当前雷达、任务触发、商业全球化、科学设备、开发文化边界已追回。 |
| `tophub-ai-reassessment-pages-*`、`tophub-ai-page-by-page-ledger`、`tophub-ai-final-closure` | `ledger-ai-reassessment-compression`、`final-closure-residual-bridges` | covered | AI 215 节点闭环、不建 AI 组、Folo 四候选边界已追回。 |
| `tophub-development-reassessment-pages-*`、`tophub-development-page-by-page-ledger`、`tophub-development-page-25-final-closure` | `ledger-development-directory`、`final-closure-residual-bridges` | covered | 开发 290 节点闭环、项目绑定规则已追回。 |
| `tophub-development-and-ai-directory-audit`、`tophub-development-directory-audit` | `early-directory-audits-execution-bridge`、`community-podcast-shopping-dev-finance-bridge` | covered_bridge | 早期 audit 作为判断起点承接，不作为当前状态。 |

---

## 6. 财经 / 报刊 / 购物

| 原始文件组 | 已承接 ledger | 状态 | 说明 |
|---|---|---|---|
| `tophub-finance-reassessment-pages-*`、`finance-page-by-page-ledger`、`finance-final-closure` | `ledger-finance-page-by-page`、`community-podcast-shopping-dev-finance-bridge` | covered | 财经 324 节点闭环、等量替换、工具与 Web3 边界已追回。 |
| `tophub-finance-directory-audit`、`finance-platform-execution-verified` | `early-directory-audits-execution-bridge`、`pending-actions-platform-verification-boundary` | covered_bridge | 早期 audit 与平台执行证据已分层。 |
| `tophub-newspapers-*` | `ledger-newspapers-sampled-closure`、`early-directory-audits-execution-bridge`、`pending-actions-platform-verification-boundary` | covered_sampled | 报刊是代表性抽样闭环，不宣称 873 全量逐项在线核验。 |
| `tophub-shopping-*` | `ledger-shopping-page-by-page`、`community-podcast-shopping-dev-finance-bridge`、`final-closure-residual-bridges` | covered | 购物 95 节点闭环，购物不进注意力流。 |

---

## 7. 娱乐 / 社区 / 专栏 / 政务校务

| 原始文件组 | 已承接 ledger | 状态 | 说明 |
|---|---|---|---|
| `tophub-entertainment-*` | `ledger-entertainment-final-actions`、`entertainment-platform-sports-fiction-podcast-boundary`、`entertainment-media-music-game-fiction-platforms`、`entertainment-remaining-platform-boundaries`、`pending-actions-platform-verification-boundary` | covered | 娱乐最终动作、平台、体育、网文、播客、影视音乐游戏、剩余平台边界已追回。 |
| `tophub-community-*` | `ledger-community`、`community-core-boundary`、`community-candidate-migration`、`community-podcast-shopping-dev-finance-bridge`、`pending-actions-platform-verification-boundary` | partial | 已完成社区第 1 页、核心边界、候选迁移、已执行三动作；社区热门第 2—20 页仍未提供实际页面，保持 partial。 |
| `tophub-columns-pages-*` | `ledger-columns-94-page-snapshot` | covered_snapshot | 专栏 94 页是页级 snapshot 闭环，不宣称每个节点在线实时逐项验证。 |
| `tophub-government-affairs-snapshot-closure`、`tophub-campus-affairs-snapshot-closure` | `ledger-government-campus-snapshots` | covered_snapshot | 政务 / 校务是 snapshot 边界，不宣称在线实时全量核验。 |

---

## 8. 浏览器 / 资源 / 文化 / 设备

| 原始文件组 | 已承接 ledger | 状态 | 说明 |
|---|---|---|---|
| 浏览器治理 docs/ops 三文件 | `ledger-browser-links-routing`、`root-ops-link-index-boundary` | covered | 浏览器保存抵达，沃壤保存判断。 |
| 六个浏览器候选文件 | `ledger-browser-link-candidate-cohorts` | covered | 订阅、主题包、项目学习、账号服务、待复核、垂直检索六类候选已追回。 |
| `资源.md`、`资源/README.md`、`资源/语境/_index.md`、`资源/链接/_index.md` | `resource-device-cultural-coordinate-routing`、`root-ops-link-index-boundary`、`resource-update-gate-constitution`、`browser-link-candidate-cohorts` | covered | 资源入口、语境、链接、候选确认闸已追回。 |
| `资源/雷达/得到与知识内容接触谱系.md`、`文化内容入口与影音工作流.md`、`豆瓣阅读偏好初步观察.md` | `culture-knowledge-audio-video-routing` | covered | 得到、小宇宙、豆瓣、BibiGPT、影音工作流已追回。 |
| `资源/雷达/设备生态关注说明.md`、`成为/灯塔/文化公共谈话者候选.md` | `resource-device-cultural-coordinate-routing` | covered | 设备追踪词与文化公共谈话者候选已追回。 |

---

## 9. 仍不应宣布完成的部分

```yaml
- source_id: "PR #372 仍未完成声明边界"
  current_route: explicit_non_completion_guard
  decision_status: recovered_from_pr
  selected_reason: "虽然仓库原始文件已有大量 ledger 承接，但还有若干不能宣布完成的边界。"
  rejected_alternatives:
    - "宣布 PR #372 全部判断链完成"
    - "宣布 TopHub 80 + Folo 27 全部逐源判断已完整追回"
    - "宣布旧 ChatGPT 对话已经全部迁移"
    - "远程更新用户本地待提交控制文件"
  route_reason: "用户本地仍有 master checklist / coverage audit / military ledger 等待提交；社区热门第 2—20 页仍未提供页面；旧 ChatGPT 对话追索仍未做；source index 只是索引不是判断账本。"
  reuse_rule: "没有证据的完成不能说完成；控制文件冲突未解决前不远程改控制文件。"
  evidence_locator: "coverage audit boundary, user local pending files, this reverse coverage map"
  next_action: "继续旧 ChatGPT 对话 partial 追索，或等待用户本地提交控制文件后统一登记。"
```

---

## 10. 下一步建议路线

1. **不再重复挖已覆盖的 PR 原始文件**，除非用户指出某个来源判断链仍缺“为什么选 / 为什么不选 / 为什么放此层 / 怎样复用”。
2. **转入旧 ChatGPT 对话 partial 追索**：只追回能从对话材料或已写文件中找到证据的判断，不由当前模型补编。
3. **等待用户本地提交控制文件后统一登记**：master checklist、coverage audit、军事边界 ledger 目前不要远程覆盖。

---

## 11. 本映射的复用规则

1. 先查覆盖映射，再决定是否继续读原始文件。
2. `covered` 表示有 ledger 承接，不表示每个来源都已无缺口。
3. `covered_bridge` 表示早期文件只作为桥接或时间边界，不作为当前状态。
4. `covered_snapshot` 表示基于用户提供 snapshot 或页级样本，不宣称在线实时逐项核验。
5. `partial` 表示还有明确缺口。
6. `do_not_remote_update` 表示用户本地有待提交控制文件或 ledger，当前不远程改。
