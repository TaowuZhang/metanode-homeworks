# docs/cockpit · cockpit 与户口产物层

`docs/cockpit/` 保存 cockpit、户口、Project、状态锁与 workflow 证据等机器或分支产物的读法说明。

它不是 `main` 上的稳定知识库，也不是谷每天直接看的 dashboard。这里的许多报告来自 workflow 或专门分支，读它们时必须先确认产物在哪条分支、是否是最新烘焙结果。

## 当前子层

| 子层 / 文件 | 读法 |
|---|---|
| `absolute-bedrock-2026-06-13.md` | cockpit 早期绝对底座记录 |
| `account-phase-lock/` | 账户 / Project phase lock 产物层 |
| `status-phase-lock-inspector/` | 状态字段检查产物层 |
| `workflow-evidence/` | 周度只读异常面；比较 Actions 的 last attempt 与仓库的 last healthy，不给代码打质量分 |
| `docs/cockpit/ignition-wire.md` | branch-only 产物；由 `.github/workflows/cockpit-ignition-wire.yml` 写到 `cockpit/ignition-wire` 分支，不期待在 `main` 出现 |

## 放什么

- workflow 生成或维护的 cockpit / 户口报告；
- Project、issue、phase、status 的对账产物；
- 真实运行证据缺口、反例与复验状态；
- 需要 GitHub diff / 分支保存的机器报告。

## 不放什么

- 谷长期取用的领域 / 资源 / 成为材料；
- 公约、触发语、X 的判断规则；
- 手写 dashboard 或替代现场的读物；
- 逐 PR 质量台账或自动责任裁决。

## 维护原则

1. 产物可以厚，入口必须薄。
2. branch-only 产物必须明说分支，不准写成 main 文件。
3. 报告只显影事实与缺口，不替人决定责任归属。
4. 若某份产物开始承载长期判断，应迁入合适的长期地层，不在 cockpit 固化。
