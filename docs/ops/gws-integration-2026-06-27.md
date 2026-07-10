# Google Workspace CLI 接入判断｜2026-06-27

## 决定

沃壤采用 `googleworkspace/cli`（命令名 `gws`）作为本地 AI 的 Google Workspace 能力依赖。

接入目标不是同步 Google 数据，也不是把沃壤改造成办公自动化系统，而是让本地 Agent 在用户授权下获得结构化、可审计的 Gmail、Calendar、Drive 等操作能力。

## 空间边界

1. 外部材料被明确留下时，默认进入 `资源/` 或其候选区。
2. 邮件和日程可以只是一次操作，不必沉积为仓库材料。
3. `领域/` 与 `成为/` 由本地对话、判断和长期耕耘形成，不由 Workspace 自动路由。
4. Workspace 原始响应、邮件正文、附件和 OAuth 凭据默认不提交 Git。

## CLI 边界

不新增顶层名词命令 `wo gws`。

Google Workspace 与 NotebookLM 一样，是现有动词 `touch` 的远手能力。预期最终入口：

```bash
wo touch --gws calendar +agenda --today
wo touch --gws gmail +triage
wo touch --gws calendar +insert ...       # 默认 dry-run
wo touch --gws --live calendar +insert ...
```

第一阶段先交付独立桥接层：

```bash
node scripts/gws-bridge.mjs --no-input --timeout 30s calendar +agenda --today
```

这样可以先验证上游安装、OAuth、命令稳定性和实际使用价值，再把薄入口接入 Rust CLI。

## 安全姿态

- `auth` 永远在桥接层之外，由人直接运行 `gws auth ...`；
- 读取操作原样透传；
- 识别到写操作时自动补 `--dry-run`；
- 只有显式 `--live` 才允许真实写入；
- 初始 OAuth 权限只选 `drive,gmail,calendar`；
- 上游版本固定为已审查的 `0.22.5`，按月复核；
- 不在第一阶段接 CI、服务账号或远程凭据导出。

## 上游治理判断

`gws` 位于 `googleworkspace` GitHub 组织，但上游 README 仍声明它不是正式受支持的 Google 产品，并提醒 v1.0 前可能发生破坏性变化。

2026 年 6 月的公开报道转述原作者 Justin Poehnelt 的说法：项目走红后他被 Google 解雇，而 Google 几乎同期推动 CLI 官方化。当前公开材料足以确认存在作者、雇佣和项目归属争议，但不足以证明 Google 复制了另一个独立仓库或实施了代码盗用。

工程应对不是猜测内部事实，而是：固定版本、保持薄适配层、保留可替换性，不把沃壤内部模型绑定到上游实现。

## PR 拆分

### PR A｜本地供应链接头（当前）

- `scripts/gws-bridge.mjs`；
- 写操作 dry-run / live 守门；
- `scripts/gws-doctor.sh`；
- 单元测试；
- 固定版本与治理登记；
- 不改 Rust 命令面，不读取真实 Workspace 数据。

### PR B｜`wo touch --gws`

- 在 `touch` 中识别并转交 gws 参数；
- 保留原始 stdout JSON，stderr 只写桥接元数据；
- 增加缺少二进制、未认证、dry-run 与 live 的集成测试；
- 在本机完成真实的只读 Calendar/Gmail smoke。

### PR C｜Agent 使用说明

只收录已经实际使用过的少量动作，例如：

- 查看今天日程；
- 查找特定邮件；
- 起草邮件；
- 预演创建或调整日程；
- 找到会议相关 Drive 文档。

不把上游一百多个 Skills 全量复制进沃壤。

## 验收

```bash
node --test scripts/gws-bridge.test.mjs
node --test scripts/gws-readonly-contract.test.mjs
node --check scripts/gws-bridge.mjs
bash -n scripts/gws-doctor.sh
```

本机真实验收需另外完成：

```bash
gws version
gws auth login --scopes drive,gmail,calendar
node scripts/gws-bridge.mjs --no-input --timeout 30s calendar +agenda --today
```
