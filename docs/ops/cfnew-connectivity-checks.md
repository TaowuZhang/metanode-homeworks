# CFnew 可达性与替代检查

这份文档只管个人连接体验的验收，不做公开测速，不做攻击性探测，不把任何订阅或 token 写入仓库。

## 目标

未来判断一条线是否可用，不只看“能打开管理页”。还要看：

- 常用网站是否能稳定打开；
- 延迟是否足够低；
- 是否有替代线路；
- 朋友线和自用线是否能分开调；
- 故障时能不能知道该换哪条线。

## 当前基线

最新线路实测基线：

- `docs/ops/connectivity-baseline-2026-06-13.md`

当前结论：

- 百变小樱当前节点：主工作线合格。
- Flower Dance：朋友线 / 备用线合格。
- Good Old Days：自用备用线合格。
- Clash `127.0.0.1:7890`：当前终端工作代理入口。

## 终端代理入口

当前本机终端代理入口记录在：

- `docs/ops/terminal-proxy-switches.md`

当前结论：

- `127.0.0.1:7890` 可作为 Clash 终端 HTTP proxy。
- `127.0.0.1:7897` 虽然端口 open，但作为 HTTP / SOCKS5 proxy 测 Notion 均失败，暂不纳入终端代理函数。
- `.zshrc` 只保留 `proxy_clash` / `proxy_off` / `proxy_status` / `proxy_test_p0` / `proxy_test_full` 开关，不永久开启代理。

## 目标清单

当前常用网站目标清单已经单独沉到：

- `docs/ops/connectivity-targets.md`
- `config/connectivity-targets.v0.json`

它按 P0 / P1 分层：

- P0：Notion、GitHub、Cloudflare、AI 工作台。
- P1：Google / YouTube / BibiGPT / NotebookLM / UX 与开发资料 / 社交与娱乐资源。

## 检查分层

### 1. Worker 存活

用 `scripts/cfnew-smoke.sh` 检查：

- UUID 路径 HTTP 200；
- 页面标题包含品牌名；
- 不打印完整管理入口。

### 2. 客户端订阅可用

在实际客户端里检查：

- 订阅能导入；
- 节点列表能刷新；
- 选中节点能连通；
- 不把订阅链接贴进 Git、聊天或截图。

### 3. 常用网站半自动验收

先打开终端代理：

```bash
proxy_clash
```

再用本地当前网络出口跑：

```bash
node scripts/connectivity-check.mjs
```

或者用 `.zshrc` 开关：

```bash
proxy_test_p0
proxy_test_full
```

如果要输出一次本地报告：

```bash
node scripts/connectivity-check.mjs --out=reports/connectivity/latest.json
```

注意：

- 这个脚本只从当前设备 / 当前客户端 / 当前代理出口看可达性。
- 它不代表朋友所在地体验。
- 它不提交报告，`reports/` 已被忽略。
- 它不读取订阅链接、UUID、KV 或 token。

### 4. 常用网站手动验收

每次只做少量手动检查，避免滥用公开测速。

建议分类记录，而不是提交真实账号路径：

| 类别 | 检查方式 | 记录 |
|---|---|---|
| 搜索 | 首页能打开，搜索结果能返回 | OK / 慢 / 失败 |
| 文档 | 常用技术文档能打开 | OK / 慢 / 失败 |
| 代码托管 | 仓库页能打开，raw 文件能拉 | OK / 慢 / 失败 |
| 视频 / 图片 | 首页与缩略图能加载 | OK / 慢 / 失败 |
| AI / 工作台 | 登录页或工作台能进入 | OK / 慢 / 失败 |

### 5. 速度体感

优先使用客户端内置延迟测试和少量下载体感。

记录建议：

```text
日期：YYYY-MM-DD
地点 / 网络：例如 家宽 / 手机热点 / 朋友所在地
线路：Flower Dance / Good Old Days / 购买源
客户端：Clash / sing-box / 其他
体感：快 / 可用 / 慢 / 不可用
问题网站类别：代码托管 / 文档 / 视频 / AI / 其他
备注：只写现象，不写 token
```

## 替代策略

不要追求一条线解决所有网络。

建议保持三层：

1. 购买源：省心主线。
2. Flower Dance：朋友独立备用线。
3. Good Old Days：自用备用线。

每条线内部再保留多候选节点或多出口，但不要在仓库里写真实节点秘密。

## 什么时候换线

- 单站失败：先换节点，不急着重部署。
- 多站失败但管理页正常：优先看客户端和出口线路。
- 管理页 / UUID 路径失败：看 Worker / KV / 上游兼容日期。
- P0 多站失败：直接换到另一条线，不在原线里反复耗。
- 新部署后立刻 1101：先回滚或清理重建，不连续猜 flag。

## 不做什么

- 不自动刷公开测速站。
- 不公开分享订阅。
- 不提交真实节点、UUID、订阅 token。
- 不把朋友线的体验套到自用线。
- 不把一次网络波动写成永久规则。
