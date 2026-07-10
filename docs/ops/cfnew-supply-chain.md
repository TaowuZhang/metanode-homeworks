# CFnew 连接供应链

这份文档把 CFnew 从一次性部署，沉成沃壤里的可迁徙供应链。

它不保存任何真实链接。真实订阅、UUID、KV、Cloudflare account、token，只留在本地和密码管理器里。

## 定位

CFnew 在这里不是主入口，也不是公开分发页。

它是：

- 朋友线 / 自用线的备用连接供应链；
- 能重建、能品牌化、能回滚、能验收的一套操作骨架；
- 百变小樱等购买源之外的第二条路；
- Git 里只留下依赖、流程、边界和脚本，不留下秘密。

当前两条线的语义：

| 线 | 用途 | 命名 | 原则 |
|---|---|---|---|
| Flower Dance | 朋友线 | 独立 Worker + 独立 KV + 独立 UUID | 不吃自用线排序，不绑定电信体验 |
| Good Old Days | 自用线 | 独立 Worker + 独立 KV + 独立 UUID | 可按自己的网络体验排序 |

## 上游依赖

上游：`byJoey/cfnew`

当前已确认的硬依赖：

- `compatibility_date = "2026-01-20"`
- 默认取明文源，便于做标题品牌化；
- 不默认加 `nodejs_compat`；
- 不把某次本地试错猜测当成依赖。

本仓库只记录依赖判断，不 vendor 上游 Worker 源码。

## 本仓库留下什么

应该留下：

- runbook；
- env 模板；
- 品牌替换脚本；
- smoke 检查脚本；
- 更新 / 回滚 / 故障边界；
- 哪些值不能进 Git。

不应该留下：

- 真实 UUID；
- 真实订阅链接；
- 管理入口完整 URL；
- KV namespace id；
- Cloudflare account id；
- Cloudflare token；
- 百变小樱账号 token；
- 带真实值的 `wrangler.toml`；
- 下载后的 Worker 源码。

## 本地准备

依赖：

- `bash`
- `python3`
- `curl`
- `wrangler`
- 已登录 Cloudflare 的本地环境

复制模板：

```bash
cp config/cfnew.example.env .env.cfnew.local
```

编辑 `.env.cfnew.local`，填本地真实值。这个文件不提交。

## 拉上游并品牌化

```bash
set -a
source .env.cfnew.local
set +a

scripts/cfnew-fetch-brand.sh
```

默认行为：

- 从上游 main 拉取源文件；
- 写入本地临时目录；
- 把页面标题里的 `订阅中心` 替换成 `${BRAND_NAME} · 订阅中心`；
- 生成 `cfnew-current.js`。

## 部署原则

每条线都独立：

- 独立 Worker；
- 独立 KV；
- 独立 UUID；
- 独立品牌名；
- 独立验收。

不要把朋友线复用自用线的 UUID / KV / 排序。

`wrangler.toml` 应在本地生成或手写，不提交。最小要点：

```toml
name = "flower-dance"
main = "cfnew-current.js"
compatibility_date = "2026-01-20"

[[kv_namespaces]]
binding = "C"
id = "replace_with_local_kv_namespace_id"
```

## smoke 检查

部署后只检查管理页或 UUID 路径，不用根路径判断成败。

```bash
set -a
source .env.cfnew.local
set +a

scripts/cfnew-smoke.sh
```

脚本原则：

- 不打印完整 URL；
- 不打印 UUID；
- 只检查 HTTP 状态和标题片段；
- 不输出页面正文。

## Git 守门

仓库有一条轻量守门脚本：

```bash
node scripts/cfnew-guard.mjs
```

它只做一件事：挡住真实 CFnew / Cloudflare / 订阅秘密和本地部署产物进入 Git。

会拦的东西包括：

- `.env*`、`.secrets/`、`*.uuid`；
- 带真实值的 `wrangler.toml`；
- 下载后的 Worker 源码；
- 真实管理入口 / 订阅链接形态；
- 非占位的 KV / Cloudflare token 赋值。

这不是部署检查，也不碰 Cloudflare；它只是 Git 边界检查。

## 更新流程

1. 看上游 README 和 `wrangler.toml`，确认兼容日期是否变化。
2. 新开本地工作目录，不在旧目录上糊。
3. 拉上游明文源。
4. 品牌化。
5. 本地生成 `wrangler.toml`。
6. 部署到目标 Worker。
7. smoke 检查 UUID 路径。
8. 跑 `node scripts/cfnew-guard.mjs`，确认没有把本地秘密带进 Git。
9. 如果流程判断变了，再开 PR 更新这份 runbook。

## 回滚流程

如果新版本 1101 或管理页异常：

1. 不先猜功能 flag。
2. 先确认上游 README 的硬要求。
3. 先测 UUID 路径，不测根路径。
4. 如果刚部署前有稳定版本，用 Cloudflare deployment 回滚。
5. 如果 Worker / KV 已被污染，清理后重建，不在脏状态上继续叠补丁。
6. 记录故障类型，不记录真实链接。

## 已见故障样本

| 症状 | 判断 | 处理 |
|---|---|---|
| 1101 | Worker 运行时错误；根路径不可靠 | 先测 UUID 路径；回看上游 README；必要时清理重建 |
| 证书域名不匹配 | 测试 Worker 名或域名状态异常 | 删除测试壳；回到正式 Worker |
| 剪贴板复制成命令 | 用户复制了整段命令，不是订阅链接 | 从页面按钮复制订阅链接，不从终端输出里捞 |
| 反复猜 `nodejs_compat` | 未看上游就试错 | 停止猜；以 README 和 `wrangler.toml` 为准 |

## 与百变小樱的关系

百变小樱仍可作为主线客户端 / Clash 订阅源。

CFnew 这条供应链只解决：

- 备用；
- 分线；
- 朋友独立体验；
- 可迁徙和可回滚。

不要把购买源账号、订阅 token、后台地址写进这里。

## 验收

一次 PR / 一次更新算完成，当且仅当：

- 文档说明了依赖和边界；
- 脚本只从 env 取真实值；
- `.gitignore` 覆盖本地秘密和下载产物；
- smoke 不输出 UUID / 订阅链接；
- `cfnew-guard` 通过；
- 没有真实 Cloudflare / 订阅信息进入 Git。
