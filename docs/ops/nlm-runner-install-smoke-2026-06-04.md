# NLM runner install smoke

日期：2026-06-04
状态：候选施工 / 待在有网络环境执行
位置：scripts + docs/ops

## 0. 一句话

`scripts/nlm-runner-install-smoke.sh` 是安全环境里的安装前门：默认不安装、不登录、不碰 NotebookLM；只有显式设置 `NLM_INSTALL_TMC=1` 时，才尝试安装 `tmc/nlm`，随后只跑 help 级 smoke。

## 1. 本刀新增

- `scripts/nlm-runner-install-smoke.sh`

它承接上一刀：

- `scripts/nlm-runner-smoke.sh`
- `scripts/nlm-runner-http.mjs`

## 2. 默认行为

直接执行：

```bash
bash scripts/nlm-runner-install-smoke.sh
```

如果环境里没有 `nlm`，它应返回：

```plain text
status: missing_nlm
```

这不是失败事故，而是安全关门：默认不自动安装外部工具。

## 3. 显式安装 tmc/nlm

在有网络、可安装 Go package、且确认这个环境可以作为 runner 候选时执行：

```bash
NLM_INSTALL_TMC=1 bash scripts/nlm-runner-install-smoke.sh
```

脚本会：

1. 检查 `go`；
2. 设置 `GOBIN`，默认 `$HOME/go/bin`；
3. 执行：

```bash
go install github.com/tmc/nlm/cmd/nlm@latest
```

4. 把 `$GOBIN` 加入本次脚本 PATH；
5. 调用 `scripts/nlm-runner-smoke.sh` 的 help 级 smoke。

## 4. 边界

本脚本不做：

- 不执行 `nlm auth`；
- 不打开浏览器；
- 不保存 Google cookie；
- 不调用 NotebookLM；
- 不上传 / 同步 source；
- 不执行 query；
- 不启动 HTTP runner。

它只回答：

> 这个环境能否安装并识别一个 `nlm` 命令？

## 5. 输出

输出目录：

```plain text
target/nlm-runner-install-smoke/
```

记录：

- install log；
- status JSON；
- 后续 help smoke 的输出。

## 6. 通过后的下一步

若 `install_ok` 且 `scripts/nlm-runner-smoke.sh` 返回 `ok_help_only`：

1. 不立刻 query；
2. 先决定 NotebookLM 登录态放在哪个受控 runner 环境；
3. 再手动完成 `nlm auth` 或该工具要求的登录步骤；
4. 再跑 query 级 smoke；
5. 最后再启动 `scripts/nlm-runner-http.mjs`。

## 7. 仍不做

- 不在 GitHub Actions 里存个人 Google cookie；
- 不把 NotebookLM 登录态提交到仓库；
- 不开放公网 runner；
- 不把 query 结果写回沃壤。
