# NLM 灵感显影本地动态验证 · 2026-06-27

状态：首轮实验完成；不进入自动 sync 设计。

## 范围与边界

- Notebook：`沃壤｜灵感显影`
- Notebook ID：`6e7c59f2-096b-4800-bc26-44e1c2e0d4cc`
- 上传窗口：2026-06-27 23:10:43–23:11:15（UTC+8）
- 仅上传 6 个 corpus part、`成为/身体.md`、`wo/ask-questions.md`。
- 未公开分享 Notebook，未 sync、未删除 Source、未回写 `成为/`、`领域/` 或原始灵感文件。

## 运行环境与真实 CLI 契约

- Go：`go1.26.4 darwin/arm64`；安装前本机没有 Go。
- `tmc/nlm`：`v0.0.0-20260614235420-c19fbf794209`，路径 `/opt/homebrew/bin/nlm`。
- PATH 中原有 `/Users/jiajia/.local/bin/nlm` 是 Python `notebooklm-mcp-cli 0.6.13`，不是 `tmc/nlm`；本轮全部实验显式调用 `/opt/homebrew/bin/nlm`。
- `source pack` 默认阈值为 5,120,000 bytes，支持 `--chunk`、`--exclude`、`--pre-process`。
- `source sync` 是 add/update/skip/delete 计划面，带 hash cache；本轮未调用。
- `chat` 支持 `--source-ids`、`--citations=json`、`--prompt-file -`。
- 真实偏差：`chat <notebook-id> [conversation-id | prompt]` 在已有会话后会把位置参数 prompt 误判成 conversation-id；`--prompt-file -` 可消除输入歧义。
- 真实失败：误判产生的本地会话名含非法字节后，后续回答虽到达 `phase:done`，CLI 仍可能以 `illegal byte sequence` 或 `unexpected EOF` 退出 1。

## 测试与构建

- `cargo test`：58 passed。
- `node --test scripts/lib/nlm-output.test.mjs`：9 passed。
- 两个 Node 语法检查与 `bash -n scripts/nlm-grain-probe.sh`：通过。
- `python3 scripts/test_build_nlm_ideas_corpus.py`：通过。
- `python3 scripts/build_nlm_ideas_corpus.py --check`：`fresh`。
- PR merge ref 上另有两项非 NLM 红灯：root guard 报 `ops/`、`runs/` 未登记；Worang Index 报根段 24 条、内部断链 101。未把全库治理混入本次实验。

构建器在上传前发现并修复三类真实解析错误：

1. `④开放追问/迁移场景与衍生物` 的标题后缀泄漏进正文；
2. 正文中的“语言标签/类型标签”被误当作元数据终止词，造成字段截断；
3. 同行字段、`nn②/③/④` 导出错误及 `①张力点/③对话触发点/③可迁移/④标签` 变体造成串线或漏字段。

修复均先加入失败测试，再修改解析器。全量复核后，源文件中实际存在的已知字段均被捕获；内容哈希、Obsidian URL、导出元数据与字面 `\n` 均未进入 corpus。

## Corpus

总卡片数：293。总生成大小（含 catalog）：633,865 bytes。最大单卡约 9.4 KB。

| 文件 | 卡片 | bytes | SHA-256 | Source ID | 上传时间（UTC） |
|---|---:|---:|---|---|---|
| `cards-001-050.md` | 50 | 99,074 | `8a6b02dee11ca2b23d65e3d6711e4a6570d8c81c3e372cb4533976ab5d6221b6` | `929a438c-77ce-4181-8e90-c1f4480f6085` | 15:10:43 |
| `cards-051-100.md` | 50 | 110,239 | `b8d6783c3d56eb7940d0e7ee63e8b97e1d200785b6c0d13149e94b5d80bbd5e8` | `37e71676-3161-4f82-833f-28d6e961b1f0` | 15:10:47 |
| `cards-101-150.md` | 50 | 111,127 | `c8715f152ffc11b9768060c80cf0ae2f687b8d1ce07e6f172a1218273be26c5e` | `4d8a210a-3302-4480-b8ef-9cfd700c9944` | 15:10:52 |
| `cards-151-200.md` | 50 | 108,083 | `1d8ea4797f18f49b631e9be9e5ee47ba949a433e27899b65d568ade7d19ee1e7` | `00bd91aa-f191-4587-b181-de12e31242cb` | 15:10:57 |
| `cards-201-250.md` | 50 | 113,028 | `63728f13e5ec1087a4ded248f101676a042373d7bf6a35210fe3de272e05010f` | `e666afee-10ab-48a1-a61c-8af9304062f2` | 15:11:02 |
| `cards-251-293.md` | 43 | 91,479 | `0aade373a95d941dc7aaced1603b3dddaef8ae0adb40701c59e682a1001d5797` | `79ab0992-b255-4870-8565-2d751d08f2b1` | 15:11:07 |

附加 Source：

- `身体.md`：`3691f103-3aa7-4d36-a21d-c8807dc81985`，15:11:11Z。
- `ask-questions.md`：`81c655c4-7a59-4959-b168-0b9d64a319ef`，15:11:15Z。

### 离线 pack

六卷分别执行 `nlm source pack`，txtar 文件名准确，输出均为原文件 bytes 加 23-byte marker，逐字节一致。合包 633,168 bytes，中文标题与 `成为/灵感/` 路径完整；未分块、未截断，不需要 pre-process。

### 十张固定随机抽查

随机种子：`20260627`。

| idea | 标题 | 结果 |
|---|---|---|
| `9f979d7fb587` | 证据充分性由论断风险决定 | 七字段正确 |
| `29bca98f367b` | Sensemaking 七属性映射 | 复合第四节标题修复后正确 |
| `3d8f74f87710` | 关系角色固化陷阱 | 七字段正确 |
| `142b1d34cc12` | 工具显影度优先于工具总量 | 复合第四节标题修复后正确 |
| `d6f0e427fde1` | 许可性措辞陷阱 | 字面 `\n` 展开，hash/URL 已排除 |
| `94937b0007f3` | 玩得好的认知困境 | `[待补]` 忠实保留，导出字段排除 |
| `c9c6281c09e7` | 因果时序陷阱 | `④迁移场景与衍生物` 不再粘入连接 |
| `75c2792fdb7e` | 脱离式觉知与去融合 | “语言标签/类型标签”不再截断正文 |
| `8a1e646531ee` | 研究问题层级框架 | 七字段正确 |
| `a80a058ea7f4` | 感知控制降焦虑 | 复合第四节标题修复后正确 |

## A/B/C 消融

固定问题：

> 在当前限定 sources 中，灵感库呈现出什么跨文件结构？请找出至少三组跨卡关系，每组至少引用三张真实灵感卡，并至少包含一组标题、标签和显式连接都不相近但机制可以互相照亮的远距离桥。逐组给出卡片标题或 idea:id，保留 NotebookLM citation，明确区分材料已知连接与新推断，并主动指出反证或证据缺口。不要只按标签聚类；找不到证据时明确说找不到。

| 组 | Source | 回答结果 | 判断 |
|---|---|---|---|
| A | 6 个 corpus part | 完成 3 组：AI 乘法退化、无意志力行动设计、命名/存在门槛远桥；主卡 9 张 | 跨卡覆盖最好，未退化为标签聚类；适合作为默认开放综合 |
| B | A + `身体.md` | 只完成 1 组“数字存在论/碰撞生成真相”即结束 | 身体轴提供“根身/肉身/不可逆”解释，但强行套框架并显著压缩覆盖面 |
| C | B + `ask-questions.md` | 恢复 3 组：AI 主体性、注意力熵增、底座错位远桥 | 问题地图恢复簇/桥任务相关性，但把 Source 中 244 条时期的历史快照当成当前卡数，产生旧问题锚定 |

结论：新增锚点不是单向增益。`身体.md` 应仅在身体动作/操作化问题中按需加入；`ask-questions.md` 应先刷新历史统计，或把问题直接放进 prompt，而不是作为默认常驻 Source。默认推荐 6 个 corpus-only。

## 价值评测

| 题 | 分数 | 回答与失败类型 |
|---|---:|---|
| A1 同机制跨领域 | 2/2 | 生物分诊、协议执行、默认选项共同形成“资源匮乏→成本排序→自动剥离” |
| B1 远距离配对 | 1/2 | AI 无距离处理与人类靠气口制造认知距离形成好桥；第二卡未给 `idea:id`，标签为模型推断 |
| B2 身体动作映射 | 2/2 | 织/筛/拆/刺/算/领映射 6 张真实卡，理由不靠关键词；首次过长截断，限长复测完成 |
| C1 内部张力 | 2/2 | 数字主体性 vs 目的论构造；语义方法杠杆 vs 结构性卡点；冲突类型清楚 |
| C2 证据强弱 | 1/2 | 返回三强三弱；最后一张弱证据混用了“核心洞见待补”，越过仅看证据字段的纪律 |
| D3 反向攻击 | 2/2 | 先提出“技能外挂实现等价在场”，再以行为等价错觉、行动三门槛、指令幻觉、隐性知识堵死反攻，缩小到可操作化区 |
| E1 灵感库地形 | 1/2 | 给出簇/桥/孤点/边界，但复用了 `ask-questions.md` 的 244 条时期历史快照（53/22%/76），未与当前 catalog count 293 对账 |

总分：11/14。

### Citation 准确性

- `--citations=json` 的结构化事件能稳定定位到 `cards-*.md` Source，但 `title` 为空，不能直接定位单卡；单卡核验仍依赖回答中的 `idea:<id>`。
- `--citations=tail` 只列 Source ID/文件名，不给卡片行号。
- 部分长回答正文继续生成 `[7]–[27]` 编号，但对应 structured citation 事件并不总是完整出现；不能把漂亮编号当成已核验引用。
- 本轮引用到的 `idea:id` 均可在 corpus 回查；B1 有一张真实卡只给标题未给 ID。

## 双入口

### 递碴模式

命令：

```bash
NLM_BIN=/opt/homebrew/bin/nlm \
NLM_BECOMING=7fe617c9-8250-4e9f-be53-41ea7cd941c1 \
cargo run -- touch --nlm becoming "谷如何接住一次小失败？"
```

结果：成功返回 5 粒“成为｜坐标｜索引”材料，不分析，结尾要求“含着它，回 taste”。粒子入口符合设计。

### 显影模式

直接 `nlm chat --citations=json` 能返回 5+ 卡、多个机制层、反证与证据缺口，未压成 grain，综合价值成立。

但两次完整查询都在生成 `phase:done` 后非零退出：一次 `unexpected EOF`，一次因本地误判会话名触发 `illegal byte sequence`。因此内容路径通过，CLI 进程稳定性未通过。

2026-06-28 已加入严格的本地兼容判定：只有 NDJSON/JSON 完整可解析、收到明确 final/done 事件、最终正文非空、必需引用已收集，且非零退出发生在最终回答之后，才把结果标为 `answer_status: complete`、`usable: true`，同时保留 `process_exit_code` 与 warning。没有 final、截断或坏 JSON、认证失败、Notebook/Source 不存在、引用解析失败、超时或空答案仍按真实失败处理；测试覆盖了“完整答案后退出 1”和这些失败分支。

## 下一阶段判断

- 是否继续接入整个 `领域/`：否。先只按具体问题加入少量领域 anchor；当前锚点已证明会施加偏置。
- 推荐 Source 组合：默认 corpus-only；B2 类问题临时加 `身体.md`；刷新 `ask-questions.md` 的历史统计后，才在地形问题中临时加入。
- 是否进入第二阶段 sync 设计：否。本地已能区分“完整答案 + 进程告警”和真实失败，但上游完成后非零退出仍未消失；同名 Python/Go CLI、prompt/conversation-id 歧义、本地非法会话名、citation 仅到 part、问题地图 stale 也仍是门槛。
- corpus 构建层本身已具备进入下一轮实验的条件；远端无人值守 sync 尚不具备。
