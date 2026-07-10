# 分发 safety gate 验收｜2026-06-01

时间：2026-06-01  
原始记录分支：`distribution-safety-gate-2026-06-01`  
清洁摘出分支：`docs/distribution-safety-gate-clean-2026-06-01`  
原始验收基线：`refactor/prioritize-knowledge-dirs` @ `9b6f2076 data: canonicalize csv sources`  
清洁落地主线：`main`

## 0. 结论

本轮完成的是 **分发 safety gate / last-click dry-run 验收**，不是正式发布。

结论：

> 六口 last-click dry-run 通过。  
> clean worktree 策略成立。  
> cookies、Playwright、截图、packages 均停留在本地 ignored 层。  
> 未发生真实发布。

下一步不继续扩平台；如要真实发布，必须另选经「结撰 / 编校」放行的主稿，并只选择 1–2 个平台，经人工确认后执行。

## 1. 上游提醒与当前基线

另一个 AI 提醒：`main` 已合并 PR #102，CSV pair 已 canonicalize：

- 普通 `.csv` 是 canonical source。
- 成对 `*_all.csv` 已删除。
- `sense` / `make` / `trace` 已改读 canonical CSV。
- 不应继续依赖 `*_all.csv`。

本轮复核：

- `main` 已包含 `data: canonicalize csv sources`。
- 原始分发 clean worktree 基于 `refactor/prioritize-knowledge-dirs` 的 `9b6f2076`，同样处在 canonical CSV 语义之后。
- 本轮分发 safety gate 不依赖 `*_all.csv`。
- 本页是从原 #103 厚分支中摘出的干净 safety-gate 记录；技能治理、灵感沉积等材料不随本页进入此 PR。

## 2. clean worktree

为避免污染原工作区的大施工现场，本轮新建 clean worktree：

```bash
git worktree add -b distribution-safety-gate-2026-06-01 ../沃壤-distribution-clean origin/refactor/prioritize-knowledge-dirs
cd ../沃壤-distribution-clean
git status --short
git log -1 --oneline
```

验收：

- `git status --short` 为空。
- HEAD 为 `9b6f2076 data: canonicalize csv sources`。
- 原工作区中的 skills / Cargo / crack / bite / 灵感沉积等未提交改动未被触碰。

## 3. 安全闸门检查

### 3.1 ignore 边界

确认以下路径均被 ignore：

```text
cookies/
ai-browse/out/
ai-browse/out/debug/
ai-browse/out/packages/
.env
.env.local
.env.production
```

结果：

- `cookies/` 由 root `.gitignore` 忽略。
- `ai-browse/out/` 由 `ai-browse/.gitignore` 忽略。
- `.env*` 由 root `.gitignore` 忽略。

### 3.2 环境变量残留

检查以下变量是否已在当前 shell 中残留：

```text
XHS_DRY_RUN
TELEGRAM_DRY_RUN
WEIBO_DRY_RUN
DOUBAN_DRY_RUN
ZHIHU_DRY_RUN
JIKE_DRY_RUN
WEIBO_PUBLISH
DOUBAN_PUBLISH
ZHIHU_PUBLISH
JIKE_PUBLISH
XIAOHONGSHU_COOKIES
TELEGRAM_BOT_TOKEN
TELEGRAM_CHAT_ID
TELEGRAM_PROXY
HTTPS_PROXY
HTTP_PROXY
```

结果：无输出。

判定：

> 没有发现长期环境变量残留；dry-run / publish 开关仍需在当次命令中显式给出。

### 3.3 入口脚本

本轮查看入口脚本，未发现隐藏的默认真发脚本或 alias 包装入口。

重点入口仍是：

```text
scripts/wo-excrete.mjs
scripts/xhs-login.mjs
scripts/weibo-login.mjs
scripts/douban-login.mjs
scripts/zhihu-login.mjs
scripts/jike-login.mjs
```

判定：

> 发布风险集中在平台 pipe 如何解释 env；未发现默认一键真发入口。

## 4. 本地 ignored 运行态

### 4.1 cookies

将原工作区本地 cookies 复制进 clean worktree 的 ignored `cookies/`：

```text
cookies/xiaohongshu_cookies.json
cookies/weibo_cookies.json
cookies/douban_cookies.json
cookies/zhihu_cookies.json
cookies/jike_cookies.json
```

`git status --short` 仍为空。

### 4.2 Playwright

clean worktree 初始没有浏览器依赖，登录脚本报：

```text
playwright is required for browser pipes
```

随后在 `ai-browse/` 内安装依赖：

```bash
cd ai-browse
npm install
cd ..
```

确认：

```text
playwright OK
```

`node_modules/` 位于 ignored 本地层，不提交。

## 5. dry-run 命令

测试主稿：

```text
/tmp/worang-distribution-dry-run.md
```

内容明确为 dry-run 安全演练，不是正式主稿。

执行：

```bash
REAL_FILE=/tmp/worang-distribution-dry-run.md

cargo run -- touch "$REAL_FILE"
cargo run -- excrete last --platforms xiaohongshu,telegram,weibo,douban,zhihu,jike

XHS_DRY_RUN=1 TELEGRAM_DRY_RUN=1 TELEGRAM_PROXY=http://127.0.0.1:7890 WEIBO_DRY_RUN=1 DOUBAN_DRY_RUN=1 ZHIHU_DRY_RUN=1 JIKE_DRY_RUN=1 node scripts/wo-excrete.mjs
```

注意：

> 这条命令显式设置六口 dry-run。  
> 没有设置任何 `*_PUBLISH=1`。  
> 没有去掉小红书 / Telegram 的 dry-run 开关。

## 6. 验收结果

事件：

```text
392eade6-fc51-41b6-a7f5-753b5a2aa0d8
```

日志时间：

```text
2026-06-01T08:55:07Z
```

| 平台 | 状态 | 判定 |
|---|---|---|
| 小红书 | `delivered` | `XHS_DRY_RUN=1` 下表示 dry-run 截图成功，不是真发 |
| Telegram | `drafted` | 生成 `.telegram.txt`，未实发 |
| 微博 | `drafted` | dry-run / 草稿成功，未发送 |
| 豆瓣 | `drafted` | dry-run / 草稿成功，未发布 |
| 知乎 | `drafted` | dry-run / 草稿成功，未发布 |
| 即刻 | `drafted` | dry-run / 草稿成功，未发布 |

## 7. 本地产物

截图：

```text
ai-browse/out/debug/2026-06-01T08-54-54-154Z-douban-dry-run.png
ai-browse/out/debug/2026-06-01T08-54-54-977Z-jike-dry-run.png
ai-browse/out/debug/2026-06-01T08-54-55-915Z-zhihu-dry-run.png
ai-browse/out/debug/2026-06-01T08-54-55-924Z-weibo-dry-run.png
ai-browse/out/debug/2026-06-01T08-55-06-966Z-xiaohongshu-dry-run.png
```

packages：

```text
ai-browse/out/packages/392eade6-fc51-41b6-a7f5-753b5a2aa0d8.douban.md
ai-browse/out/packages/392eade6-fc51-41b6-a7f5-753b5a2aa0d8.jike.md
ai-browse/out/packages/392eade6-fc51-41b6-a7f5-753b5a2aa0d8.json
ai-browse/out/packages/392eade6-fc51-41b6-a7f5-753b5a2aa0d8.md
ai-browse/out/packages/392eade6-fc51-41b6-a7f5-753b5a2aa0d8.telegram.txt
ai-browse/out/packages/392eade6-fc51-41b6-a7f5-753b5a2aa0d8.weibo.txt
ai-browse/out/packages/392eade6-fc51-41b6-a7f5-753b5a2aa0d8.zhihu.md
```

这些产物均位于 `ai-browse/out/`，不提交。

## 8. 安全判定

本轮验证成立：

1. clean worktree 能隔离原大施工现场。
2. 无危险长期环境变量残留。
3. `cookies/`、`ai-browse/out/`、`.env*` 均被 ignore。
4. 本地 cookies 注入后，浏览器五口可完成 dry-run 截图 / 草稿。
5. Telegram 可完成 drafted。
6. 小红书 `delivered` 在 `XHS_DRY_RUN=1` 下不是正式发布语义。
7. 未发生真实发布。

## 9. 下一步门槛

下一步不是继续扩平台，也不是直接真发。

如进入受控真实发布演练，必须满足：

1. 主稿已经完成「结撰 / 编校」并被人工放行。
2. 先对该主稿跑六口 dry-run。
3. 人眼检查截图与草稿。
4. 只选择 1 个平台，最多 2 个平台。
5. 谷明确确认真实发布。
6. 真实发布开关只在当次命令中显式设置，不写入默认脚本、alias 或长期环境变量。

真实发布开关仍是：

```text
小红书：去掉 XHS_DRY_RUN=1
Telegram：去掉 TELEGRAM_DRY_RUN=1，并配置 token/chat/proxy
微博：WEIBO_PUBLISH=1
豆瓣：DOUBAN_PUBLISH=1
知乎：ZHIHU_PUBLISH=1
即刻：JIKE_PUBLISH=1
```

## 10. 收束

关：A 分发安全线本轮完成，六口 dry-run 通过。  
接：下一轮若继续，应先从内容线选定经「结撰 / 编校」放行的主稿。  
沉：本页记录 2026-06-01 safety gate 验收。  
退：无 cookies、占位路径未替换、Playwright 缺失等中间失败已耗散，不另开任务。
