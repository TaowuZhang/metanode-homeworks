# NLM runner smoke 第一刀

日期：2026-06-04
状态：候选施工 / 未部署
位置：docs/ops

## 0. 一句话

先验证 `nlm` runner 能不能在受控环境里被看见、识别命令形态，并在明确 notebook/source 绑定后执行一次 grain-only query；不先接 Worker，不先处理 cookie，不先开全功能 NotebookLM MCP。

## 1. 本刀新增

- `scripts/nlm-runner-smoke.sh`

原计划同时新增 GitHub Actions 手动 workflow；但当前 GitHub 集成对 `.github/workflows/` 写入返回 `Resource not accessible by integration`，所以本刀先不落 workflow 文件，避免扩大权限面。

## 2. Smoke 分两级

### 2.1 help 级

默认级别，不调用 NotebookLM。

目的：

- 检查 `nlm` 是否存在；
- 保存 `nlm --help`；
- 粗判命令形态：`legacy_query_notebook` 或 `chat_source_ids`；
- 验证 runner 外壳可跑。

本级不需要 NotebookLM cookie、notebook id 或 source id。

### 2.2 query 级

需要显式提供：

- `NLM_SMOKE_NOTEBOOK`
- `NLM_SMOKE_SOURCE_IDS`
- `NLM_SMOKE_QUESTION`

目的：

- 只对指定 source ids 发起 grain-only query；
- 不允许 deep research；
- 不允许 source add / sync；
- 不允许 delete / share；
- 输出保存到 `target/nlm-runner-smoke/`。

## 3. 本地 / runner help 示例

```bash
bash scripts/nlm-runner-smoke.sh
```

如果环境还没有 `nlm`，脚本会以 `missing_nlm` 退出。这是预期的关门失败。

## 4. 本地 / runner query 示例

```bash
NLM_SMOKE_MODE=query \
NLM_SMOKE_NOTEBOOK="<notebook-id>" \
NLM_SMOKE_SOURCE_IDS="<source-id>" \
NLM_SMOKE_QUESTION="沃壤里 NotebookLM 应该扮演什么角色？" \
bash scripts/nlm-runner-smoke.sh
```

脚本会生成固定 prompt：

```plain text
你是图书馆管理员，不是解释者。
只基于指定 source_ids 摘出与问题相关的 3 到 5 条事实碎片。
不要分析，不要推断，不要总结用户，不要建议。
每条必须是一口大小，并尽量保留原文措辞与短来源。
```

## 5. 候选 GitHub Actions workflow

若后续权限允许，可新增手动 workflow：

```yaml
name: NLM Runner Smoke
on:
  workflow_dispatch:
    inputs:
      mode:
        default: help
      install_tmc_nlm:
        default: false
permissions:
  contents: read
```

该 workflow 只应手动触发，不应定时运行；默认 `mode=help`，不碰 NotebookLM。

## 6. 边界

本刀不做：

- 不写 Google cookie；
- 不写 NotebookLM token；
- 不部署 Worker；
- 不自动 sync source pack；
- 不接全功能 NotebookLM MCP；
- 不把回答写回沃壤。

## 7. 下一步

若 help 级通过：

1. 选择实际 runner：优先 `tmc/nlm`，保留 `notebooklm-mcp-cli` 作为参考。
2. 在安全环境里完成一次 query 级 smoke。
3. 把可用 runner 封成 HTTP runner：输入 action + notebook/source ids，输出 grains JSON。
4. 再把 Worker 的 `NLM_RUNNER_URL` 指向该 runner。
