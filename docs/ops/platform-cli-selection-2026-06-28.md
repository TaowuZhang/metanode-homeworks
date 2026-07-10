# 平台外手 CLI 选型与接入记录 · 2026-06-28

## 结论

不自造平台发布器。沃壤只维护发布包、平台 adapter、plan/apply 闸门和 receipt；实际平台协议交给成熟上游 CLI。

| 工具 | 当前证据 | 判定 | 沃壤用途 |
|---|---|---|---|
| [`xdevplatform/xurl`](https://github.com/xdevplatform/xurl) | X 官方；MIT；约 1.2k stars；v1.1.1 于 2026-05 发布；官方 OAuth 与 JSON stdout | 试验 | X 第一优先发布外手 |
| [`gitroomhq/postiz-agent`](https://github.com/gitroomhq/postiz-agent) | Postiz 官方 CLI 2.0.15；主仓约 32k stars；OAuth/device flow；覆盖 X、LinkedIn、YouTube、TikTok | 试验 | 先创建 remote draft，再人工提升为 schedule |
| [`dreammis/social-auto-upload`](https://github.com/dreammis/social-auto-upload) | 约 12.9k stars；CLI 覆盖小红书、抖音、B站、视频号、YouTube；活跃更新 | 试验 + live 闸 | 复用 `sau` 参数与 uploader；当前只生成 plan，不由沃壤自动 apply |
| [`white0dew/XiaohongshuSkills`](https://github.com/white0dew/XiaohongshuSkills) | 约 3k stars；明确提示限流/封号风险；支持 CDP、无头与互动自动化 | 暂缓发布能力 | 只借鉴预览/参数，不接互动和无头自动发布 |
| [`comeonzhj/Auto-Redbook-Skills`](https://github.com/comeonzhj/Auto-Redbook-Skills) | 约 1.9k stars；README 明示平台打击 AI 托管运营；有成熟 3:4 卡片渲染 | 借鉴 | 后续可复用卡片生成，不接 Cookie 自动发布 |
| [`TryGhost/Ghost`](https://github.com/TryGhost/Ghost) | MIT；约 54k stars；活跃发布；官方 Admin API、会员、付费订阅、newsletter | 评估 | 个人知识网络/自有长期主站首选候选 |

以上 star 与活跃度是 2026-06-28 快照，不作为合规证明；合规仍以平台官方 API/OAuth 与账号实测为准。

## 平台路由

### 第一优先

- 小红书：`sau` plan + 小红书成品包；当前 `live_allowed=false`。
- X：`xurl` plan；允许在人工核对 plan id 后显式 apply。

### 第二优先

- LinkedIn：Postiz 创建 remote draft；不直接 schedule。
- 脉脉：尚未验证到成熟、可维护、合规的 CLI，保留 draft/last-click adapter，不伪造自动发布能力。

### 国内后续

- 微信公众号：长文草稿 adapter。
- 视频号：`sau tencent` plan。
- 知乎、豆瓣：统一 draft/last-click。
- B站：`sau bilibili` plan。
- 抖音：`sau douyin` plan。

### 国外后续

- YouTube、TikTok：优先 Postiz 官方 OAuth 路线；`sau` 仅作为本地 uploader 备选。
- LinkedIn：Postiz 官方 CLI；底层平台能力以 LinkedIn Posts API 为准。

## 知识网络与知识付费

- 国内按“小报童”理解并登记为 `xiaobot`；另登记小鹅通 `xiaoe-tech`。当前都只生成可迁移草稿，等官方开放接口/导入能力确认后再接。
- 国外首选评估 Ghost：它同时拥有内容主站、Admin API、会员、付费 tiers、newsletter 和数据可迁移性，最符合“个人知识网络 + 付费”。
- Substack 登记为分发出口，但不把它设为 canonical source；未验证到稳定公开写入 API 前保持 last-click。

## 安装与检查

```bash
scripts/publish/install-external-clis.sh
node scripts/publish/external-cli-doctor.mjs
```

固定版本登记：`config/external-publish-tools.v0.json`。工具安装在 Git 忽略的 `target/tools/`。

## 使用

```bash
# 第一批平台：生成 CLI plans，不发布
wo excrete last \
  --platforms xiaohongshu,x-twitter,linkedin,maimai \
  --mode dry_run \
  --prepare \
  --json

# X 显式 live 只生成可执行计划
wo excrete last --platforms x-twitter --mode live --prepare --json

# 人工检查 external-plan.json 后，plan id 必须逐字确认
node scripts/publish/apply-external-plan.mjs \
  --plan ai-browse/out/publish-runs/<event-id>/x-twitter-external-plan.json \
  --confirm-plan-id <plan-id>
```

Postiz 的 LinkedIn/YouTube/TikTok adapter 默认只允许 `last_click` 创建 remote draft，并要求先由用户完成 `postiz auth:login` 与 integration 选择。

## 安全边界

- 不读取 `~/.xurl`、Postiz token、平台 Cookie。
- 不把 secret 写入 Git、Notion、发布包或 receipt。
- 不使用 `xurl --verbose`，不向 Agent 命令传 inline secret。
- `sau` plan 不可被通用 executor 执行；真实平台 upload 仍需单独人工放行。
- 不实现 stealth、仿真人、反检测、自动互动或批量养号。

## 官方依据

- X 官方 xurl 文档：<https://docs.x.com/tools/xurl>
- X 创建 Post API：<https://docs.x.com/x-api/posts/create-post>
- LinkedIn Posts API：<https://learn.microsoft.com/en-us/linkedin/marketing/community-management/shares/posts-api>
- TikTok Content Posting API：<https://developers.tiktok.com/doc/content-posting-api-get-started>
- YouTube `videos.insert`：<https://developers.google.com/youtube/v3/docs/videos/insert>
- Ghost Admin API：<https://docs.ghost.org/admin-api>
- Ghost memberships：<https://docs.ghost.org/members>

## 研究放行记录

- 结果集同质性：中文搜索结果高度集中于浏览器自动化、Cookie、托管运营；英文成熟结果集中于官方 OAuth/API 调度器。两类不能混为同一风险等级。
- SIFT：官方文档/官方仓库优先；README 自报风险仅用于证明其自身边界，不用于证明平台政策全貌。
- CRAAP：X、LinkedIn、TikTok、YouTube、Ghost 的 API 判断由当前官方文档支撑；小红书/脉脉“没有成熟合规 CLI”的判断仅为本轮搜索结果，标记为部分验证。
- 双维标记：`xurl/Postiz/Ghost [来源:可靠] [内容:已验证]`；`sau/小红书 skills [来源:一般] [内容:部分验证]`。
