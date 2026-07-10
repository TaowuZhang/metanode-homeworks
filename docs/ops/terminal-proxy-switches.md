# 终端代理开关与端口判定

> 日期：2026-06-13  
> 目的：记录本机终端如何进入可用网络出口，避免把代理端口永久写死，也避免 Clash 与机场官方 App 混用时误判。

## 当前结论

当前本机同时存在两个科学上网工具：

- Clash
- 机场官方 App

本轮实测后，终端工作代理先定为：

```text
Clash HTTP proxy: 127.0.0.1:7890
```

不把代理环境变量写死在 `.zshrc` 顶层。只保留手动开关函数。

原因：

- 终端和浏览器可能走不同出口；
- 两个代理软件同时开时，端口语义容易混；
- 永久 export 会让 Git / curl / Node 永远走某个旧端口；
- 后续切换机场官方 App 时，终端可能仍误走 Clash；
- 代理软件没启动时，新终端会出现隐蔽网络失败。

## 已验证端口

### 端口探测

```text
OPEN   127.0.0.1:7890
OPEN   127.0.0.1:7897
closed 127.0.0.1:7899
closed 127.0.0.1:1080
closed 127.0.0.1:10808
closed 127.0.0.1:20170
closed 127.0.0.1:6152
```

### Notion 连通性测试

`7890` 作为 HTTP proxy：

```text
https://www.notion.com
status=200
time≈1.10s
```

`7897` 作为 HTTP proxy：

```text
LibreSSL SSL_connect: SSL_ERROR_SYSCALL
status=000
time≈5.00s
```

`7897` 作为 SOCKS5 proxy：

```text
LibreSSL SSL_connect: SSL_ERROR_SYSCALL
status=000
time≈5.00s
```

因此：

- `7890` 可以作为终端代理口；
- `7897` 虽然端口 open，但当前不能当作终端 HTTP / SOCKS5 代理口使用；
- 不据此判断机场官方 App 不可用，只能判断它当前不适合作为终端 `curl` 代理口。

## `.zshrc` 当前保留的开关

已加入本机 `~/.zshrc`：

```bash
# ---- proxy switches ----
proxy_clash() {
  export HTTP_PROXY=http://127.0.0.1:7890
  export HTTPS_PROXY=http://127.0.0.1:7890
  unset ALL_PROXY
  echo "proxy on: Clash http://127.0.0.1:7890"
}

proxy_off() {
  unset HTTP_PROXY
  unset HTTPS_PROXY
  unset ALL_PROXY
  echo "proxy off"
}

proxy_status() {
  echo "HTTP_PROXY=${HTTP_PROXY:-}"
  echo "HTTPS_PROXY=${HTTPS_PROXY:-}"
  echo "ALL_PROXY=${ALL_PROXY:-}"
}

proxy_test_p0() {
  proxy_status
  node scripts/connectivity-check.mjs --limit=14 --out=reports/connectivity/p0-current.json
}

proxy_test_full() {
  proxy_status
  node scripts/connectivity-check.mjs --out=reports/connectivity/full-current.json
}
```

## 使用方式

需要 GitHub / Notion / Cloudflare / AI 工具链时：

```bash
proxy_clash
```

确认当前终端代理状态：

```bash
proxy_status
```

跑 P0 工作命脉检查：

```bash
proxy_test_p0
```

跑全量常用网站检查：

```bash
proxy_test_full
```

做完需要回到无代理终端：

```bash
proxy_off
```

## 已验证 P0 / P1 全量结果

在 `proxy_clash` / `127.0.0.1:7890` 下，`scripts/connectivity-check.mjs` 全量通过：

- P0：Notion、Notion API、GitHub、GitHub Raw、GitHub API、Cloudflare、ChatGPT、Claude、Gemini、Google AI Studio、Perplexity 均有响应。
- P1：Google、Google Scholar、YouTube、BibiGPT、NotebookLM、得到、NNGroup、Mobbin、Figma、MDN、Telegram、Discord、X、Reddit、小红书、Steam、豆瓣、Bilibili、AcFun、Apple Music 均有响应。

注意：部分 AI / API 目标返回 `403`、`404`、`421`，这代表服务器有响应，不代表浏览器或正式 API 调用失败。

## 判准

### 当前主工作代理

```text
proxy_clash → 127.0.0.1:7890
```

用于：

- `git fetch` / `git pull`
- GitHub Raw 拉取
- Cloudflare / Wrangler 前置网络
- Notion API / bridge / 自动化
- P0 / P1 可达性检查

### 机场官方 App

当前只保留为：

- 图形界面备用；
- 朋友小白入口候选；
- 手机 / iPad / Windows 官方客户端方向。

暂不把它写入终端代理函数，直到确认可用终端代理端口与协议。

## 不做什么

- 不在 `.zshrc` 顶层永久 export 代理。
- 不把机场官方 App 的 `7897` 当成可用终端代理口。
- 不用浏览器能打开来推断终端能打开。
- 不把一次测速结果当永久线路质量。
- 不把 `reports/connectivity/*.json` 提交进 Git。
