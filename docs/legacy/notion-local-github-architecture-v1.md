# [已废弃] 架构定义 · Notion × 本地 × GitHub 操作规格书

> 来源：Notion 温棚旧终端桥遗留页（《[已废弃] 架构定义 · Notion × 本地 × GitHub 操作规格书》）
> 状态：已废弃，仅作历史参考
> 当前现行架构：Notion 薄中心与工具驾驶舱 + GitHub MCP / PR 流
> 迁移日期：2026-05-18

> 这份文档定义「本地终端桥架构」里的每一个构件、每一种操作。
> 先把规格写死，再动手开发。所有后续搭建以此为准。
> 现已退出现行架构，仅作历史记录保存。

---

## 一、架构总览

三个端，各管一件事：

| 端 | 角色 | 核心职责 | 物理位置 |
|---|---|---|---|
| **Notion** | 指挥中心 | 下发命令、展示结果、调度多窗口、记录交接 | 云端（Notion 工作区） |
| **本地 M1** | 执行引擎 | 轮询命令、跑 Shell、回填结果、唤醒 AI | 老 M1 电脑（插电常开） |
| **GitHub** | 资产仓库 | 存代码、存文档、跑 CI、提 PR、版本管理 | `github.com/dongxi-heji/worang` |

**数据流（两跳）**：克（AI）→ Notion 终端页写 `$` 命令 → 本地 watcher 拉取执行 → 结果回填 Notion → 克读取回填

---

## 二、Notion 端：需要搭建什么？

### 构件清单

| 构件名 | 类型 | 用途 | 是否已有 |
|---|---|---|---|
| **终端（主控台）** | 页面 | 挂领用牌 + 交接槽 + 子页面入口 | ✅ 已有 |
| **终端一～二十** | 子页面 | 每个页面 = 一个虚拟 Shell 槽位 | ✅ 已有 |
| **终端桥状态块** | 子页面 | watcher 心跳、当前主线、异常状态 | ✅ 已有 |
| **终端桥与本地作业规程** | 子页面 | 完整操作规程（进场必读） | ✅ 已有 |
| **本页（架构定义）** | 子页面 | 所有操作的规格定义 | ✅ 本页 |

### 构件 A：终端领用牌

**位置**：终端主控台页面顶部  
**用途**：防止多个 AI 窗口或多个人同时写同一个终端槽位导致冲突。

**结构**：
- **本地写权**：标明当前谁拥有 `git commit/push` 权限（同一时间只有一个窗口可以改代码树）
- **A 车道｜终端一～四**：标明占用者 + 任务简述
- **B 车道｜终端五～八**：同上
- **C 车道｜终端九～十二**：同上
- **D 车道｜终端十三～十六**：同上
- **备用｜终端十七～二十**：波峰溢出用

**操作定义**：

| 操作 | 谁做 | 怎么做 |
|---|---|---|
| 领牌 | AI 进场时 | 找一条「空闲」车道，改成「占用｜克（Notion AI）｜任务简述」 |
| 释放 | AI 任务结束时 | 改回「空闲」 |
| 抢写权 | 需要 git 写操作时 | 在「本地写权」行标明自己，同一时间只允许一个窗口 |

### 构件 B：交接槽

**位置**：终端主控台页面底部（八大交规 callout 上方）  
**用途**：跨窗口 / 跨天 / 中断后的唯一接手入口。不是任务库，不是档案馆。

**操作定义**：

| 操作 | 谁做 | 怎么做 |
|---|---|---|
| 写交接 | AI 中断/收束时 | 带时间戳追加一段：写入者、状态、结论、下一刀 |
| 读交接 | 下一个 AI 窗口进场时 | 只看最新一段，决定接手或忽略 |
| 清理 | 谷定期 | 把已过时的旧交接段删除 |

**格式**：

```text
写入者：克（Notion AI）｜2026-XX-XX HH:MM
状态：○ 已收束 / ◐ 待接手 / ● 紧急
结论：一句话说清当前状态
下一刀：接手者应该先做什么
```

### 构件 C：终端槽位（终端一～二十）

**每个终端页 = 一个虚拟 Shell**

**操作定义**：

| 操作 | 符号 | 含义 |
|---|---|---|
| 投递命令 | `$ command` | AI 写入一条待执行的 Shell 命令 |
| 认领 | `⚡ command` | watcher 已拿到，正在准备执行 |
| 执行中 | `⏳ 执行中` | watcher 追加的执行状态标记 |
| 成功 | `✓ command` | 执行完毕，exit code = 0 |
| 失败 | `✗ command` | 执行完毕，exit code ≠ 0 |
| 输出 | Toggle 子块 | 包含 stdout/stderr 的折叠输出 |

**车道分工**：
- 第 1 槽（终端一/五/九/十三）：主线 / 写入 / Git
- 第 2 槽：只读侦察
- 第 3 槽：验证 / 测试
- 第 4 槽：备用 / 长任务

### 构件 D：终端桥状态块

**用途**：watcher 定期回写的健康页面，让 AI 和人都能一眼看到桥的状态。

**包含字段**：
- 当前主线方向
- 各终端是否可用
- watcher 是否在线
- 最近异常记录

---

## 三、本地端：代码与操作定义

### 核心文件

| 文件 | 位置 | 用途 |
|---|---|---|
| `watch.mjs` | 仓库根目录 | 主脚本：轮询 Notion → 认领 → 执行 → 回填 → 唤醒 |
| `scripts/watch-awake.sh` | scripts/ | 启动脚本：加载 env → 检查 → 启动 watch.mjs |
| `.watch/jobs/*.json` | 本地（gitignore） | Job Ledger：每个命令的执行记录 |
| `.watch.mjs.lock` | 本地（gitignore） | 进程锁：防止两个 watcher 同时跑 |

### 操作定义：watcher 生命周期

| 阶段 | 操作 | 说明 |
|---|---|---|
| **启动** | `~/.local/bin/worang-watch` | 从 Keychain 读 NOTION_TOKEN，启动 watch.mjs + caffeinate |
| **轮询** | 每 5 秒 Notion API | 读取终端主控台的所有子页面，检查哪个槽位有变化 |
| **发现** | 检测 `$` 开头的块 | 找到 pending 命令 |
| **认领** | `$` → `⚡`  • 写 jobId | 防止重复执行；本地写 `.watch/jobs/{jobId}.json` |
| **执行** | `exec(command)` | 在本地 Bash 环境跑命令，捕获 stdout/stderr |
| **回填** | `⚡` → `✓/✗`  • append output | 把结果写回 Notion 终端页 |
| **唤醒** | 剪贴板 + Cmd+V + Enter | 在 Notion 聊天框发"继续"，唤醒 AI |
| **冷却** | 等待最小间隔 | 确保回填已可见，再触发唤醒 |

### 操作定义：命令速度分档

| 档位 | 类型 | 典型命令 | 预计耗时 |
|---|---|---|---|
| **S 档** | 瞬时 | `echo`、`cat`、小 Python | < 1 秒 |
| **M 档** | 仓库读 | `git status`、`grep`、`find` | 2-10 秒 |
| **N 档** | 网络 | `git push`、`git fetch` | 5-45 秒 |
| **L 档** | 长任务 | 构建、批处理、大量下载 | 分钟级 |

### 本地环境要求

- macOS（老 M1）
- Node.js（运行 watch.mjs）
- Git + GitHub CLI (`gh`)（代码操作和 PR）
- `caffeinate`（防休眠）
- Ghostty/Terminal 开启辅助功能权限（允许模拟按键发"继续"）
- `NOTION_TOKEN` 存在 macOS Keychain
- 同一时间只允许一台电脑运行 watcher

---

## 四、GitHub 端：操作定义

### 仓库信息

- **仓库**：`dongxi-heji/worang`
- **主分支**：`main`
- **工作目录**：老 M1 的 iCloud 路径下的 Obsidian/kb/沃壤

### 操作定义：Git 工作流

| 操作 | 在哪里执行 | 命令模板 | 注意事项 |
|---|---|---|---|
| **查看状态** | 任意终端 | `$ git status --short --branch` | 只读，安全 |
| **查看日志** | 任意终端 | `$ git log --oneline -10` | 只读，安全 |
| **查看差异** | 任意终端 | `$ git diff --stat HEAD` | 只读，安全 |
| **暂存文件** | 终端一（独占） | `$ git add "path/to/file"` | 只在终端一 |
| **提交** | 终端一（独占） | `$ git commit -m "type: message"` | 只在终端一 |
| **推送** | 终端一（独占） | `$ git push && git rev-parse HEAD` | 必须带 hash 验证 |
| **验证推送** | AI 侧 | `github.loadFile(ref: hash)` | 击穿 GitHub 缓存 |
| **拉分支** | 终端一（独占） | `$ git checkout -b feature/xxx` | 新功能开发时 |
| **提 PR** | 终端一（独占） | `$ gh pr create --title "..." --body "description"` | 需要 `gh` CLI |
| **查 PR 状态** | 任意终端 | `$ gh pr checks` | 只读 |
| **合并 PR** | 终端一或浏览器 | `$ gh pr merge --squash` 或手动 | 不可逆，需确认 |

### 操作定义：安全红线

| 红线 | 规则 |
|---|---|
| **Git 写操作独占** | 所有 `git add/commit/push/checkout/merge/rebase` 只在终端一 |
| **Push 必带车牌** | `git push && git rev-parse HEAD`，拿 hash 验证 |
| **删除先探测** | 批量删除第一轮只 `-print`，确认后第二轮才 `-delete` |
| **Push 前必确认** | 先在另一终端 `git diff --stat HEAD` 看清楚再推 |
| **串联必带 STEP** | 多步命令用 `echo "=== STEP N ==="` 分隔 |

---

## 五、跨端协作：完整操作流程

### 场景 A：写一个文件并提交到 GitHub

1. **AI 领牌**：在 Notion 终端领用牌拿 A 车道 + 本地写权
2. **AI 投递**：在终端一写 `$ cat << 'EOF' > path/to/file.md ... EOF`
3. **watcher 执行**：本地创建文件
4. **AI 验证**：收到"继续"后，在终端二写 `$ cat path/to/file.md` 确认内容
5. **AI 提交**：在终端一写 `$ git add path/to/file.md && git commit -m "docs: add file" && git push && git rev-parse HEAD`
6. **AI 验证推送**：用 hash 通过 GitHub 工具确认文件已上传
7. **AI 释放**：改领用牌为「空闲」，写交接（如需要）

### 场景 B：开分支 + 改代码 + 提 PR

1. **AI 领牌**：拿写权 + 车道
2. **拉分支**：终端一 `$ git checkout -b feature/xxx`
3. **写代码**：终端一写入文件（heredoc 或 Python 脚本）
4. **验证**：终端二 `$ cat file` 检查
5. **提交**：终端一 `$ git add . && git commit -m "feat: xxx"`
6. **推送**：终端一 `$ git push -u origin feature/xxx && git rev-parse HEAD`
7. **提 PR**：终端一 `$ gh pr create --title "feat: xxx" --body "description"`
8. **等 Review**：谷在浏览器看 PR，或 AI 用 `$ gh pr checks` 查 CI
9. **合并**：谷确认后，终端一 `$ gh pr merge --squash` 或谷手动合并
10. **切回 main**：终端一 `$ git checkout main && git pull`
11. **释放**

### 场景 C：纯调研（不动代码）

1. **AI 领牌**：拿车道，**不拿写权**
2. **AI 用搜索/网页工具**：不经过终端，直接在对话里完成
3. **产出写 Notion 页面**：如果需要留档，写 Notion 页面或交接槽
4. **释放车道**

### 场景 D：中断 / 过夜交接

1. **AI 写交接**：在终端主控台交接槽追加一段（时间戳 + 状态 + 下一刀）
2. **AI 释放**：车道改空闲，写权释放
3. **下一个窗口进场**：先读交接 → 决定接手或忽略 → 领牌 → 继续

---

## 六、未来扩展：如果要做类似项目，Notion 需要搭什么？

如果未来有新项目也想用这套「Notion 指挥 + 本地执行 + GitHub 存储」的模式，**最小搭建清单**如下：

### 必建（4 个构件）

1. **终端主控台页面**
	- 顶部：领用牌（防冲突）
	- 底部：交接槽（防断裂）
	- 中间：终端槽位子页面的入口
2. **终端槽位子页面 × N**
	- 数量按并发需求定（我们用了 20 个，小项目 4-8 个就够）
	- 每个页面就是一个虚拟 Shell
3. **状态块页面**
	- watcher 定期回写心跳
	- 人和 AI 都能一眼看到桥的健康状态
4. **操作规程页面**
	- 就是你现在看到的这份文档的精简版
	- 进场必读，写命令前必看

### 可选（按需加）

1. **异步回环分析页**：当桥出了怪事，把排查过程和结论沉淀在这里
2. **项目文档数据库**：如果项目产出大量文档，单独建数据库管理
3. **CI/CD 看板**：如果 GitHub Actions 复杂，在 Notion 做一个看板同步状态

### 本地端最小搭建

1. 一台常开的电脑（不一定是 M1，任何能跑 Node.js 的都行）
2. `watch.mjs`（或等价的轮询脚本）
3. `NOTION_TOKEN`（Integration Token，在 Notion Developer 页面创建）
4. Git + GitHub CLI
5. 防休眠 + 辅助功能权限（如果需要自动唤醒 AI）

---

## 八、容错与安全最小规则

> 这里只写「不加就会出事」的四条。其他细节等遇到了再补。

### 规则1｜超时与非交互

- 默认每条命令超时 **60 秒**，到点 watcher 直接 kill 子进程。
- 明知是长任务（构建、拉镜像等），命令前加标记 `# long:600`，允许最大 10 分钟。
- **禁止交互式命令**：任何会弹输入提示的命令都不能跑（`npm init`、`gh auth login`、`ssh`、`sudo` 等）。需要用的加非交互参数，例如 `npm init -y`。
- 超时回填格式：`✗ TIMEOUT｜command` + 耗时。

### 规则2｜输出截断

- Notion API 单块上限 2000 字符，超了回填会失败。
- watcher 回填前判断：如果 stdout/stderr 总长 > 1500 字符，**只保留最后 1500 字**（错误信息通常在尾部），顶部加一行 `... (截断，完整输出见本地日志)`。
- 完整输出写本地：`.watch/jobs/{jobId}.log`。

### 规则3｜机密红线

- **Notion 终端页不允许出现明文机密**：Token、密码、API Key、私钥、Cookie。不允许 `$ export TOKEN=xxx` 这种写法。
- 机密只能放在：macOS Keychain（首选）或本地 `.env.local`。
- `.gitignore` 必须包含：`.env*`、`.watch/`。
- watcher 读机密只从 Keychain（从 NOTION_TOKEN 现在走的路径）。

### 规则4｜心跳与日志

- **本地日志**：`.watch/bridge.log`，watcher 启动、轮询出错、崩溃都写这里。
- **Notion 心跳**：watcher 每 60 秒在「终端桥状态块」页面回写一行：

```text
最后心跳：2026-XX-XX HH:MM:SS　状态：ONLINE　最近任务：terminal-N / ✓ / 12s
```

- AI 进场第一件事：看心跳是否在 2 分钟内。超过认为桥离线，不要发命令。

### 规则5｜watcher 重启必须外置

- **禁止 watcher 从 Notion 终端槽里重启自己**：终端槽里不得运行 `pkill node .*watch.mjs`、`pkill caffeinate`、`watch-awake.sh restart`、detached restart。
- 重启 watcher 只能由外部执行：谷在 Ghostty 手动执行，或以后另做 launchd / supervisor。
- Notion 终端槽只允许跑短命令、验证命令、仓库读写命令；凡是会杀 watcher、改 watcher 进程生命周期的命令，一律不从终端槽投递。
- 如果 watcher 离线，AI 只能写明离线事实和本地手动命令，不能继续往终端槽堆命令。

### 规则6｜自动唤醒不是核心链路

- 自动唤醒只是一层便利提示，不是闭环正确性的核心依赖。
- 当前 Notion 桌面端 Accessibility 树只稳定暴露到窗口 / 分组 / 窗口按钮层级，没有可稳定定位的 AI composer 控件。
- 因此，`粘贴“继续” + 回车` 只能视为脆弱辅助；失败时不得判断为命令未执行。
- 真正的完成信号以终端页回填、`.watch/jobs/{jobId}.json`、本地日志为准。

---

## 九、已拍板与保留问题

### 已拍板

- [x] **终端一旧内容先保留**：云 MCP 调研残留不清空，作为历史记录留在原处。
- [x] **`gh` CLI 可直接用**：老 M1 已登录，可以用 `gh pr create`、`gh pr checks`、`gh pr merge`。
- [x] **PR 合并权限**：AI 可以在必要时自行 `gh pr merge --squash`，不强制谷人工审批。
- [x] **终端数量暂定 20 个**：数量本身不是问题，关键是车道分工和使用纪律。当前 20 个适合多窗口并行，先不删减。

### 仍需在开发中验证

- [ ] **信息归位规则**：交接槽只放最近交接；有长期价值的信息应沉淀到对应页面或 GitHub 文档，而不是一直堆在交接槽。
- [ ] **自动唤醒可靠性**：当前风险是 Notion 对话框未获得焦点时，模拟粘贴“继续”会失败。第一版先用 AppleScript 激活 Notion 窗口，再粘贴；如果仍不稳，再研究更稳方案。

---

## 十、下一步开发顺序

### 第一步｜先做最小可用 watcher

目标不是一步到位，而是先跑通闭环：

1. Notion 终端页出现 `$ command`
2. 本地 watcher 识别并认领
3. 本地执行命令
4. 结果回填 Notion
5. 尝试唤醒 AI 继续读结果

这一版只支持最基础命令，不急着做复杂并发。

### 第二步｜加四条最小安全规则

在闭环跑通后，立刻加：

1. 超时与非交互
2. 输出截断
3. 机密红线
4. 心跳与日志

### 第三步｜接入 GitHub 工作流

再验证：

1. `git status`
2. `git diff`
3. `git add`
4. `git commit`
5. `git push && git rev-parse HEAD`
6. `gh pr create`
7. `gh pr merge --squash`

### 第四步｜再处理自动唤醒稳定性

先用简单方案：

1. AppleScript 激活 Notion
2. 选中当前对话框
3. 粘贴“继续”
4. 回车

如果这一步不稳，再单独开排错页，不在第一版里过度设计。