# GitHub issue draft｜WQB local flywheel runner v0.1

## Title

`[WQB] Ship local flywheel runner for CNHKMCP read → candidate → yellow simulation → handoff`

## Body

### 背景

WQB 不是纯架构项目。平台有时间压力，当前需要一个能在本地跑起来的闭环，而不是只讨论 Cloudflare remote MCP。

MCP 的用途不是“部署后结束”，而是把 Notion 里的当日意图接到本地 runner：

```text
Notion 今日目标
→ 克拆成 read / candidate / simulation ticket
→ 本地跑 read-only 或单次授权 simulation
→ runs/wqb-* 打包
→ 回 Notion 裁决下一刀
```

### Scope

本 issue 交付一个本地可跑包：

- `scripts/wqb-local-flywheel.mjs`
- `docs/ops/wqb-local-flywheel-2026-06-13.md`

### Acceptance Criteria

#### Local runner

- [ ] 支持 HTTP MCP transport：`WQB_MCP_HTTP_URL`
- [ ] 支持 stdio MCP transport：`WQB_MCP_CMD / WQB_MCP_ARGS`
- [ ] `smoke-readonly` 能调用 tools/list 与 green tools
- [ ] `dataset-brief` 能查 dataset brief
- [ ] `field-brief` 能查 field brief
- [ ] `ticket-sim` 只生成 simulation ticket，不运行
- [ ] `run-sim` 必须带 `--yes-i-authorize-simulation`
- [ ] `pack` 能生成 Notion handoff

#### Safety

- [ ] 默认拒绝 red tools
- [ ] 不调用 submit_alpha
- [ ] 不导出 cookie / credential
- [ ] 不保存明文密码
- [ ] 输出做基础脱敏
- [ ] 不部署 Cloudflare
- [ ] 不默认开放 multi simulation

#### Workflow

- [ ] Notion 中给出“今天做 X / 交 Y”后，克能自动拆出本地命令包
- [ ] 本地跑完后，把 `runs/wqb-*/handoff.md` 带回 Notion
- [ ] 稳定字段 / 协议 / runbook 沉 GitHub
- [ ] simulation 现场和账号相关状态留 Notion / 本地

### Non-goals

- 不做 Cloudflare remote MCP 部署
- 不 submit alpha
- 不做 autonomous bulk mining
- 不把 WQB 凭证放进仓库
- 不把平台私密内容全文沉 GitHub