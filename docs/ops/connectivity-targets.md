# 常用网站可达性目标清单 v0

> 目的：把“我平时到底要保证哪些网站能连上”从聊天判断，沉成一张可维护清单。它不保存订阅、节点、UUID、token，也不做公开测速。

## 来源

这版清单来自三类证据：

1. 当前对话里谷明确点名：**Notion** 是重点。
2. 沃壤 / Notion 搜索结果反复出现：Notion、GitHub、本地、Web、AI 工具、发布 / 同步工作台。
3. 资源与项目地层反复出现：ChatGPT、Claude、Gemini、Perplexity、DeepSeek、Kimi、BibiGPT、YouTube、Telegram、Steam、豆瓣、B 站 / 自媒体、Mobbin、Figma、NNGroup 等。

本清单不是一次定死。后续实测时，如果谷说“这个我其实不用 / 这个特别重要”，就按真实生活改。

## 分级

| 层级 | 含义 | 判准 |
|---|---|---|
| P0 | 断了就影响当天工作 | 必须优先保证，失败要立刻换线或排查 |
| P1 | 常用但可短时绕过 | 应保持可用，失败要记录 |
| P2 | 有用但不必每次检查 | 低频抽查即可 |
| CN | 国内常用 / 通常不用代理验证 | 主要用于对照，不作为代理质量核心指标 |

## P0：工作命脉

| 目标 | URL | 为什么在 P0 | 检查点 |
|---|---|---|---|
| Notion App | `https://www.notion.com` | 当前工作台、活现场、AI 协作面 | 首页 / 工作区能进入 |
| Notion API | `https://api.notion.com` | 同步、桥接、自动化依赖 | 返回非网络错误即可，不能超时 |
| GitHub | `https://github.com` | 沃壤事实底盘、PR 流 | 首页 / repo 页面能打开 |
| GitHub Raw | `https://raw.githubusercontent.com` | 脚本、上游源码、raw 文件拉取 | raw 文件域名可达 |
| GitHub API | `https://api.github.com` | GitHub MCP / 自动化 / PR 操作 | API 可达，不能连接失败 |
| Cloudflare Dashboard | `https://dash.cloudflare.com` | Worker / KV 管理 | 登录页或 dashboard 能打开 |
| Cloudflare Workers | `https://workers.cloudflare.com` | Worker 平台文档和运行面 | 页面可达 |

## P0：AI 工作台

| 目标 | URL | 为什么在 P0 | 检查点 |
|---|---|---|---|
| ChatGPT | `https://chatgpt.com` | 常用对话型 AI | 登录页 / 工作台可达 |
| OpenAI API | `https://api.openai.com` | 可能影响工具和集成 | API 域名可达 |
| Claude | `https://claude.ai` | 常用对话 / 代码协作 | 工作台可达 |
| Anthropic API | `https://api.anthropic.com` | Claude API / 工具链潜在依赖 | API 域名可达 |
| Gemini | `https://gemini.google.com` | Google AI 工作台 | 工作台可达 |
| Google AI Studio | `https://aistudio.google.com` | Gemini / 模型试验 | 页面可达 |
| Perplexity | `https://www.perplexity.ai` | 搜索型 AI / 研究 | 首页 / 查询页可达 |

## P1：研究、资料与产品设计

| 目标 | URL | 为什么 | 检查点 |
|---|---|---|---|
| Google Search | `https://www.google.com` | 搜索基础设施 | 首页 + 搜索结果 |
| Google Scholar | `https://scholar.google.com` | 学术检索 | 首页可达 |
| YouTube | `https://www.youtube.com` | 视频材料 / BibiGPT 来源之一 | 首页 / 视频页可达 |
| BibiGPT | `https://bibigpt.co` | 视频 / 音频取水 | 首页 / redirect 可达 |
| NotebookLM | `https://notebooklm.google.com` | source pack / 学习与材料消化 | 工作台可达 |
| 得到 | `https://www.dedao.cn` | GetNote / 得到大脑相关 | 首页可达 |
| Nielsen Norman Group | `https://www.nngroup.com` | UX 资料源 | 文章页可达 |
| Mobbin | `https://mobbin.com` | 产品设计参考 | 首页 / app 页面可达 |
| Figma | `https://www.figma.com` | 设计协作 / 参考 | 文件页 / 首页可达 |
| MDN | `https://developer.mozilla.org` | Web 开发文档 | 文档页可达 |
| Cloudflare Docs | `https://developers.cloudflare.com` | Worker / CFnew 排障 | 文档页可达 |

## P1：社交、沟通与内容流

| 目标 | URL | 为什么 | 检查点 |
|---|---|---|---|
| Telegram Web | `https://web.telegram.org` | 通知 / 通道 / 社群 | Web 入口可达 |
| Discord | `https://discord.com` | 社群 / 开发者社区 | 首页 / app 可达 |
| X | `https://x.com` | 信息流 / 公开讨论 | 首页 / 单贴可达 |
| Reddit | `https://www.reddit.com` | 搜索结果常见 UGC 补位 | 首页 / 帖子可达 |
| 小红书 | `https://www.xiaohongshu.com` | 发布 / 内容观察 | 首页可达 |

## P1：娱乐、影音、游戏与资源

| 目标 | URL | 为什么 | 检查点 |
|---|---|---|---|
| Steam Store | `https://store.steampowered.com` | 游戏库 / 同步 / 商店 | 商店页可达 |
| Steam Community | `https://steamcommunity.com` | 创意工坊 / 社区 | 社区页可达 |
| 豆瓣 | `https://www.douban.com` | 电影 / 书籍 / 音乐资源线 | 首页 / 条目页可达 |
| Bilibili | `https://www.bilibili.com` | 视频 / 自媒体资料 | 首页 / 视频页可达 |
| AcFun | `https://www.acfun.cn` | 资源库已有链接 | 首页可达 |
| Apple Music | `https://music.apple.com` | 音乐资源线 | 页面可达 |

## CN 对照：通常不用代理判断

这些目标更多用来判断“本地网络本身是否正常”，不作为代理质量核心指标：

- `https://www.baidu.com`
- `https://www.bilibili.com`
- `https://www.douban.com`
- `https://www.xiaohongshu.com`
- `https://www.dedao.cn`
- `https://y.qq.com`

## 验收口径

每条线至少记录：

```text
日期：YYYY-MM-DD
网络：家宽 / 手机热点 / 其他
线路：百变小樱 / Flower Dance / Good Old Days
客户端：Clash / sing-box / Shadowrocket / 官方客户端
P0：OK / 部分失败 / 失败
P1：OK / 部分失败 / 失败
最痛失败：哪个网站、什么现象
动作：换节点 / 换线 / 重部署 / 暂不处理
```

## 不做什么

- 不高频自动刷公开测速站。
- 不提交真实订阅、节点、UUID、KV、token。
- 不把一次失败直接判成永久不可用。
- 不用朋友线推断自用线体验。
- 不把国内站可达当成代理正常。
