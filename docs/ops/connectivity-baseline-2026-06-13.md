# 网络线路实测基线 · 2026-06-13

> 目的：把本轮常用网站可达性实测结果沉到沃壤，形成后续判断主工作线 / 备用线 / 朋友线的基线。  
> 边界：不提交 `reports/connectivity/*.json`，不提交订阅、UUID、KV、节点、token。这里只保留摘要和判断。

## 测试前提

- 本机终端代理入口：`proxy_clash → 127.0.0.1:7890`。
- 检查脚本：`scripts/connectivity-check.mjs`。
- 目标清单：`config/connectivity-targets.v0.json`。
- 检查范围：P0 + P1 全量目标。
- 结果口径：`200/403/404/421` 都代表目标站点有响应；`403/404/421` 对 API 根路径或反爬页面不等于浏览器不可用。

## 总结论

本轮 5 组实测全部达成：

```text
P0 / P1 全部有响应，无连接失败。
```

因此，当前重点不是重部署 Cloudflare Worker，也不是继续修 CFnew，而是把线路运营化：

1. 百变小樱当前节点可作为日常主工作线。
2. Flower Dance 已具备朋友线 / 备用线可用性。
3. Good Old Days 已具备自用备用线可用性。
4. Clash 7890 当前出口是终端可用入口，但具体节点仍会影响 Notion API、X、Discord 等体感。

## 实测组别

| 输出文件 | 线路 / 命名 | 结论 |
|---|---|---|
| `reports/connectivity/full-bbxy-current-node.json` | 百变小樱当前节点 | 全通；当前最适合作为主工作线基线 |
| `reports/connectivity/full-flower-dance.json` | Flower Dance | 全通；可作为朋友线 / 备用线，但 X 偏慢 |
| `reports/connectivity/full-good-old-days.json` | Good Old Days | 全通；可作为自用备用线，整体稳定 |
| `reports/connectivity/full-clash-7890-current.json` | Clash 7890 当前出口复测 | 全通；Notion API 偏慢一次 |
| `reports/connectivity/full-clash-7890-alt.json` | Clash 7890 另一出口 / 节点 | 全通；Notion API 更慢一次，适合备用观察 |

## P0 工作命脉判断

### 百变小樱当前节点

关键 P0：

| 目标 | 时间 |
|---|---:|
| Notion App | 1.49s |
| Notion API | 1.95s |
| GitHub | 1.03s |
| GitHub Raw | 1.32s |
| GitHub API | 0.58s |
| Cloudflare Workers | 1.97s |
| Gemini | 1.88s |
| Google AI Studio | 2.40s |

判断：

> 主工作线合格。Notion / GitHub / Raw / API 都稳，适合作为当前日常默认。

### Flower Dance

关键 P0：

| 目标 | 时间 |
|---|---:|
| Notion App | 1.21s |
| Notion API | 3.55s |
| GitHub | 1.69s |
| GitHub Raw | 1.37s |
| GitHub API | 0.54s |
| Cloudflare Workers | 2.00s |
| Gemini | 1.73s |
| Google AI Studio | 1.87s |

判断：

> 工作备用线合格，也能作为朋友线。Notion API 比百变小樱当前节点慢，但不构成失败。

### Good Old Days

关键 P0：

| 目标 | 时间 |
|---|---:|
| Notion App | 1.14s |
| Notion API | 2.70s |
| GitHub | 0.97s |
| GitHub Raw | 1.62s |
| GitHub API | 0.36s |
| Cloudflare Workers | 1.48s |
| Gemini | 1.10s |
| Google AI Studio | 1.76s |

判断：

> 自用备用线合格。此前 Good Old Days 的 Worker 运行态疑虑，不能再按“未实测”看；至少从当前客户端出口的 P0 / P1 可达性看，它可用。

### Clash 7890 当前 / alt

两次复测都全通，但 Notion API 出现 8.33s / 10.69s。

判断：

> Clash 7890 是可用终端入口，但不同节点的 Notion API 体感差异明显。遇到 Notion API 慢时，优先换节点，不先动 Worker。

## P1 内容与资料判断

### 百变小樱当前节点

亮点：

- YouTube：1.38s
- BibiGPT：0.25s
- Telegram：1.08s
- X：0.96s
- Steam Store：1.95s
- Bilibili：0.20s

判断：

> P1 表现均衡，内容 / 资料 / 娱乐都能用。

### Flower Dance

亮点：

- YouTube：0.89s
- BibiGPT：0.24s
- Figma：1.40s
- Steam Community：0.60s

注意：

- X：5.03s，偏慢。

判断：

> 视频、资料和工作站点可用；若当天大量刷 X，不优先选这条。

### Good Old Days

亮点：

- NNGroup：0.77s
- MDN：0.63s
- Reddit：0.26s
- Apple Music：1.09s

注意：

- Discord：2.68s，略慢。

判断：

> 自用备用线成立。比 Flower Dance 更适合 X，Discord 略慢但可用。

## 当前线路分工

| 角色 | 推荐线路 | 理由 |
|---|---|---|
| 日常主工作线 | 百变小樱当前节点 | Notion / GitHub / AI / P1 全通且均衡 |
| 朋友线 | Flower Dance | 独立线，P0 / P1 全通；朋友不吃自用排序 |
| 自用备用线 | Good Old Days | P0 / P1 全通，适合留作自己备用 |
| 终端入口 | `proxy_clash` / 7890 | 已验证可让 Git / curl / Node 进入可用出口 |

## 操作判准

### 平时

```bash
proxy_clash
```

然后使用 Git / Notion API / Cloudflare / AI 工具链。

### GitHub 拉不动

先看：

```bash
proxy_status
```

如果没开代理，先 `proxy_clash`，不要先改 remote 或重配 SSH。

### Notion API 慢

如果 Notion App 能开、GitHub 能开，但 Notion API 大于 8–10s：

1. 先换节点。
2. 再跑 `proxy_test_p0`。
3. 不先重部署 Worker。

### P0 多站失败

如果 Notion / GitHub / GitHub Raw / Cloudflare 多站同时失败：

1. 换到百变小樱当前节点。
2. 或切 Good Old Days。
3. 再测 P0。
4. 仍失败才排查本地代理 / DNS / 客户端。

## 未完成 / 后续

- 需要在不同真实网络下复测：家宽、手机热点、朋友所在地。
- 需要让朋友在自己的设备上做一次小白入口验收。
- 需要避免把本轮结果永久化：这是 2026-06-13 晚上的基线，不是永恒线路质量。

## 不沉积的东西

- 不提交 `reports/connectivity/*.json`。
- 不提交真实订阅链接。
- 不提交 UUID / KV / token。
- 不提交节点名、机场后台、Cloudflare account。
