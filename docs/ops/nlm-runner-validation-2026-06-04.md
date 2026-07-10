# NLM runner local validation

日期：2026-06-04
状态：本地语法 / 关门失败验证
位置：docs/ops

## 0. 一句话

本轮没有接 NotebookLM，也没有安装 `nlm`；只验证 runner scaffold 的语法面与 help 级关门失败是否符合预期。

## 1. 验证环境

执行位置：Notion AI computer sandbox `/data`

限制：sandbox 无法解析 `github.com`，因此不能从 GitHub clone 仓库，也不能安装外部 `tmc/nlm`。

表现：

```plain text
fatal: unable to access 'https://github.com/dongxi-heji/worang.git/': Could not resolve host: github.com
```

处理：降级为把当前分支中的脚本内容写入 sandbox 临时目录，做本地语法检查与关门失败检查。

## 2. 已执行检查

### 2.1 HTTP runner 语法检查

命令：

```bash
node --check scripts/nlm-runner-http.mjs
```

结果：通过。

意义：`scripts/nlm-runner-http.mjs` 至少是 Node 可解析的 ES module。

### 2.2 Smoke 脚本语法检查

命令：

```bash
bash -n scripts/nlm-runner-smoke.sh
```

结果：通过。

意义：`scripts/nlm-runner-smoke.sh` 至少是 Bash 可解析脚本。

### 2.3 Help 级关门失败

命令：

```bash
bash scripts/nlm-runner-smoke.sh
```

结果：

```plain text
沃壤 · NLM runner smoke
mode: help
status: missing_nlm
smoke exit code: 127
```

意义：在没有安装 `nlm` 的环境里，脚本没有误碰 NotebookLM，也没有继续执行网络动作，而是按预期以 `missing_nlm` 退出。

## 3. 当前结论

本轮验证成立的是：

- HTTP runner scaffold 可被 Node 解析；
- smoke 脚本可被 Bash 解析；
- 无 `nlm` 时按预期关门失败；
- 没有发生 NotebookLM 调用；
- 没有写入 cookie / token；
- 没有部署 Worker / runner。

本轮尚未验证：

- `tmc/nlm` 能否安装；
- `nlm --help` 的真实命令形态；
- `list_notebooks`；
- `list_sources`；
- `query_grains`；
- Worker 到 runner 的真实 HTTP 链路。

## 4. 下一步

下一刀应在有网络与安全登录态的 runner 环境里执行：

```bash
# 安装候选 runner
# go install github.com/tmc/nlm/cmd/nlm@latest

# help 级 smoke
bash scripts/nlm-runner-smoke.sh

# 启动 HTTP runner
NLM_RUNNER_TOKEN="<runner-token>" node scripts/nlm-runner-http.mjs

# 健康检查
curl http://127.0.0.1:8788/healthz
```

只有 help 级通过后，才进入 query 级 smoke。

## 5. 边界保留

仍不做：

- 不在聊天或 GitHub 写 Google cookie；
- 不在 GitHub Actions 里长期保存个人 NotebookLM 登录态；
- 不开放公网 runner；
- 不执行 source add / sync / delete / share / deep research；
- 不把 NotebookLM 回答写回沃壤。
