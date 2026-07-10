# wo runtime boundary v0

状态：共同运行边界 v0，不是全仓命令登记表。

本页只规定：当 `wo` 或仓库脚本要被本地 AI、自动化、CI 或其他程序稳定调用时，哪些边界必须说清楚并被证据支持。

它不替每个 action 预先定义统一 JSON envelope、统一数字退出码或统一 receipt；这些精确合同由具体 action / adapter 自己拥有。

## 1. 何时进入 runtime v0

满足任一项时，需要显式运行合同：

- 提供 `--json`、JSONL 或其他机器输出；
- 启动外部进程；
- 可能等待 stdin、TTY、认证或确认；
- 由 Agent、脚本或 CI 非交互调用；
- 写入 artifact、receipt 或外部系统；
- 对调用者承诺稳定的退出、超时、信号或兼容行为。

纯人类阅读、没有机器稳定性承诺的本地展示，不必为了形式完整强行套一层 runtime schema。

## 2. 合同局部性

每个稳定机器入口拥有自己的合同：

- schema 名称与版本；
- action-specific 状态；
- 成功 / 无变化 / 失败的语义；
- 精确数字退出码；
- side effect 与 artifact；
- 兼容与升级规则。

共同 runtime v0 只规定底线，不制造一个覆盖所有 action 的巨型总 envelope。

同一 action 的人类模式与机器模式可以不同，但不能在一次调用里含混混用。

具体命令当前做到了什么、候选 PR 走到哪里、哪些缺口尚未闭合，留在各自合同、测试、Issue 与 PR 中；这些会变化的事实不进入共同边界正文。

## 3. 结果与执行分开

共同层只要求在语义上区分三类执行结果，不规定 wire 字段或枚举名称：

- 完成：按照该 action 自己的合同完成；
- 失败：没有按照合同完成；
- 中断：被信号或明确取消中断。

局部合同可以使用 `success`、`failure`、`interrupted` 或自己的等价命名。

`no_change` 可以是 action 结果，同时执行仍属于完成。

`needs_auth`、`rejected`、`temporary` 等可以成为 action-specific 原因或状态，但不要求所有命令共享同一枚举。

## 4. stdout 与 stderr

稳定机器调用必须声明输出形状。

- 成功时，stdout 只包含约定的主结果；
- diagnostics、warnings、进度和安全提示进入 stderr；
- JSON 模式不能在 stdout 混入 prose、颜色、receipt 装饰或 shell export；
- 失败时默认保持 stdout 为空，除非 action 合同明确规定错误 envelope；
- broken pipe 的行为必须被声明并测试，不得无意 panic 或吐出半份成功结果；
- 一个 JSON document 与 JSONL stream 必须明确区分。

## 5. 非交互边界

当局部合同提供 `--no-input` 或等价非交互模式时，它的含义是禁止交互，不是“尽量别问”。

- 非 TTY、`CI=true` 或 Agent 调用不构成授权；
- 缺少输入、认证、权限或审批时必须明确失败；
- wrapper 不得自动发送换行、调用 `yes`、接受默认答案或绕过确认；
- 外部 child stdin 默认关闭，除非合同明确允许并测试输入；
- 可能等待 stdin 的 action 必须在机器入口中显式禁止或显式提供输入。

## 6. 外部进程边界

启动外部进程的 adapter 至少要说明并验证：

- 直接执行还是经过 shell；
- 最大运行时间是单个 child 的预算还是整次调用的总 deadline；
- stdout / stderr 如何读取并限量；
- 父进程 SIGINT / SIGTERM 如何传递；
- timeout 后如何 TERM、等待、KILL、reap；
- 是否可能遗留子孙进程或临时文件；
- 原始 child exit 怎样映射为公共失败，而不是直接泄漏为调用合同；
- 是否自动重试，以及重试预算、幂等前提和歧义结果如何处理；
- diagnostics 如何裁剪与脱敏。

若一次调用会连续启动多个 child，合同必须明确并测试总墙钟上限，不能让每个 child 分别获得一份看似相同的完整 timeout，导致整次调用无界放大。

纯本地计算 action 不需要为了统一形式伪造 process-group 与 receipt 机制。

## 7. 数字退出码

runtime v0 不冻结一套全仓通用 sysexits 表。

共同要求只有：

- `0` 只代表该 action 合同中的成功；
- 非零代表失败或中断；
- timeout 和真实信号中断必须可与普通失败区分；
- adapter 不应把任意第三方退出码原样当成公共合同；
- 精确数字与兼容承诺由 action / adapter 合同定义并测试。

任何局部合同都不能因为采用了一组数字退出码，就把它提升为全部 `wo` action 的数字法典。

## 8. Receipt 不是必选器官

receipt 用于外部进程、外部写入、长运行或确有审计价值的调用；不是每个本地 action 都必须生成。

使用 receipt 时：

- 放在 Git 忽略的运行目录；
- 内容有界、可删除、不可成为第二数据库；
- 不保存 token、环境变量、完整 argv、stdin、完整私人正文、无必要绝对路径或账户标识；
- 保存安全引用、工具版本、执行结果、退出与 artifact / external change 引用；
- receipt 写入失败是否推翻业务成功，必须由该 action 合同明确规定。

## 9. 证据门

一个 PR 只有在与自身范围有关的项目被真实验证后，才可以声称建立了稳定 runtime 合同。

### 所有机器入口

- 正常成功；
- 无变化（若适用）；
- 无效参数或输入；
- stdout / stderr 分离；
- 合同形状与语义可以重复验证，不要求时间戳等动态字段逐字节一致；
- broken pipe；
- 不泄漏秘密和无必要私人路径；
- 文件副作用与网络边界符合声明。

### 额外用于外部进程

- 缺少 dependency；
- 缺少认证 / 权限（若适用）；
- timeout 与总 deadline；
- SIGINT / SIGTERM；
- 子进程组清理与 reap；
- 自动重试与歧义结果路径（若适用）；
- hostile diagnostics 的限量与脱敏；
- 不通过 shell 注入未受控参数。

没有运行的项目不能写成通过。fake adapter 可以证明 process contract，不能证明真实账户、真实权限或上游版本兼容。

后续实现只需要说明自己如何符合本页相关条款，并给出与自身范围相称的真实证据；不要求维护一张包含所有 action、所有 PR 和所有当前缺口的长期总表。

## 10. 采用后的维护方式

- 本页只在共同底线变化时修改；
- action / adapter 细节留在各自合同和测试；
- 当前事实发生变化时，优先修具体合同，不机械刷新全仓矩阵；
- 新 PR 不得用“符合 runtime v0”替代真实测试证据；
- 发现共同层开始吸收具体 action 的字段、状态、命令盘点或 PR 生命周期时，拆回局部合同与施工记录。
