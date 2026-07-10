# NLM CLI 接入沃壤 · 2026-06-27

状态：可执行设计 / 本地动态验证待 Codex 完成

## 真正的难点

NotebookLM 的关键不在问答入口，而在 source architecture：

> 什么材料值得一起放进同一个 notebook，才能让跨文件连接产生真实增益？

NotebookLM notebook 不应镜像 GitHub 文件夹。文件夹按归属保存材料，notebook 应按“准备反复追问的同一个问题空间”组织材料。

因此：

- `成为/`、`领域/`、`资源/` 是沃壤地层；
- NotebookLM notebook 是临时或长期的认知工作台；
- 两者可以交叉，但不能机械一一对应。

## 当前事实

仓库已经有：

- `wo touch --nlm <route>` 的递碴入口；
- `wo/nlm/manifest.json` 的 pack 与远端绑定账；
- `scripts/nlm-grain-probe.sh` 的本地兼容手柄；
- source-pack 构建脚本与 stale 检查；
- 一套尚未部署的 Worker / runner 图纸。

本轮已修正：

- 现役探针适配 `tmc/nlm` 当前 `nlm chat ... --citations=json`；
- 保留旧 `query notebook` 回退；
- 统一解析旧单 JSON 与现代 JSON Lines；
- 移除脚本内硬编码的 notebook/source 默认值；
- runner 能正确恢复 grain 与 source refs。

## `tmc/nlm` 已经提供什么

不需要在沃壤里复制它的管理面。当前源码已经提供：

- notebook：list/create/delete/rename/description；
- source：add/list/read/delete/refresh/check；
- `source pack`：离线预览将要上传的 txtar；
- `source sync`：发现本地文件、按 hash 判断变化、默认 5MB 自动分块、支持 dry-run、exclude、pre-process 和替换；
- chat：按 `source_ids` 限定材料，并输出带 citation 的 JSON Lines；
- content transforms：summarize/critique/verify/outline/mindmap 等；
- research：fast/deep 外部研究，输出发现的来源与报告；
- `nlm mcp`：直接向 Codex、Claude Code 等暴露 NotebookLM 工具。

这意味着沃壤真正需要建设的是：

1. source 选择与结构；
2. 可重复生成的语料层；
3. 提问与验收协议；
4. 少量沃壤动作与外部 CLI 的接缝。

而不是再写一套 NotebookLM 客户端。

## 成为不再按“最高隐私”处理

`成为/` 不应因为名字私人就默认禁止进入 NotebookLM。

更合理的判断是：

> 这份材料进入 NotebookLM 后，是否能产生足够大的使用增益？

对 `成为/灵感`，答案目前是肯定的：

- 每条文件名本身就是论点；
- 卡片已有核心洞见、来源与证据、已知连接、开放追问和标签；
- 249 条材料之间已经出现引用簇、共同机制和未显式写出的桥；
- NotebookLM 最可能在这里发挥跨卡显影能力。

所以隐私只保留为边界标签，不作为首要阻断器。真正禁止进入远端的仍是登录材料、密钥、明确的 `local-only` 内容和不希望交给第三方服务的原文。

## 第一优先实验：沃壤｜灵感显影

建议建立独立 NotebookLM notebook：

> `沃壤｜灵感显影`

它不是“成为文件夹镜像”，它只回答一个长期问题：

> 这些灵感之间到底有什么结构，它们怎样才能被谷真正用上？

### 首批 source

#### 1. 灵感卡完整语料

新生成层：

`wo/nlm/corpus/becoming-ideas/cards-*.md`

来源仍是：

`成为/灵感/*.md`

每张卡保留：

- 稳定 `idea:<id>`；
- 标题与原路径；
- 标签；
- 核心洞见；
- 来源与证据；
- 已知连接；
- 开放追问。

旧 `becoming-ideas-index.md` 仍可用于轻量本地检索，但它会截断字段，而且构建器没有输出“来源与证据”，不应再承担完整 NotebookLM 语料的职责。

新构建器：

```bash
python3 scripts/build_nlm_ideas_corpus.py
python3 scripts/build_nlm_ideas_corpus.py --check
```

默认每 50 张卡一个 part；总数与 parts 以 `wo/nlm/corpus/becoming-ideas/catalog.json` 为准（首轮实验为 293 张、6 个 source）。这样既不做数百个零碎 source，也不把所有材料压成一张无法定位的大摘要。

#### 2. 可选结构轴

仅在身体动作或操作化问题中临时加入：

`成为/身体.md`

它提供织、筛、刺、搬、领、松、算、拆、画、治玉、登顶、站中间等结构轴。NotebookLM 可以据此把不同主题的灵感映射到相同认知动作，而不只是按标签聚类。

首轮消融显示，`身体.md` 会让一般跨卡问题强行套用身体框架，不应常驻。

#### 3. 临时活问题地图

仅在需要沃壤问题地形时临时加入：

`wo/ask-questions.md`

这里已经明确提出：

- 灵感库有没有自己的地形；
- 引用关系是树、网、密集簇还是孤点；
- 什么能被操作化，什么不能；
- 跨库的桥在哪里。

问题本身能告诉 NotebookLM 哪些连接对沃壤有价值，但也会锚定旧答案。使用前必须核对 `ask-questions.md` 的复核日期、catalog count 与 fingerprint 有效期。

首轮默认基线因此是 corpus-only；这两份材料都是按问题加入的镜片，不是常驻 Source。

### 暂不加入

首轮不要同时加入整个：

- `成为/`；
- `领域/`；
- `资源/`。

先让一个高内聚 notebook 通过跨卡评测。需要某个领域材料来解释灵感时，再按问题补少量 anchor source，而不是一开始吞完整库。

## 为什么不单独建 GitHub 仓库

现在先建专门文件夹，不建第二仓库：

`wo/nlm/corpus/becoming-ideas/`

原因：

- 语料是派生物，源仍在 `成为/灵感`；
- 同仓可以共享路径、hash、PR、stale check；
- 第二仓会产生同步和双份真相；
- NotebookLM notebook 可以独立，不等于 GitHub 必须拆仓。

只有需要独立分享、独立权限或 Google Drive 同步账户边界时，才考虑拆仓。

## 领域与资源随后怎样进入

### `领域/`

领域适合建立独立的长期 notebook，但应由问题空间决定，而不是一级目录数量决定。例如：

- `沃壤｜AI 与人机协作`
- `沃壤｜黑客、防御与模式锁`
- `沃壤｜财务与市场判断`
- `沃壤｜表达与内容生产`

一个 notebook 应包含：

1. 领域内的核心材料；
2. 少量来自灵感库的相关卡；
3. 领域的评价标准或方法文档；
4. 仍未解决的问题。

### `资源/`

资源仍是候选矿场。NotebookLM 可以帮助：

- 聚类；
- 去重；
- 找到与现有灵感或领域的连接；
- 识别值得迁入领域的材料；
- 用 research 补充外部来源。

但资源 notebook 应允许清仓和重建，不应自动成为长期事实底盘。

## CLI 表面：不新增 `wo nlm`

`wo` 的动作保持动词。NLM 是外部工具名，不是沃壤身体动作。

### 保留

```bash
wo touch --nlm becoming "问题"
```

这里的 `--nlm` 是远手选择器；`touch` 仍只递少量 grain。

### 完整跨卡综合

直接使用外部 CLI：

```bash
nlm chat <notebook-id> --source-ids ... "问题"
nlm chat <notebook-id> --citations=json "问题"
```

或把：

```bash
nlm mcp
```

接入本地 Codex。这样 Codex 能直接使用 NotebookLM 的原生 notebook/source/chat/research 能力，不需要沃壤再包一层名词式管理入口。

### source 管理

source 构建放在 `scripts/` 与 `wo/nlm/corpus/`；上传和同步直接调用外部 `nlm source ...`。

只有未来出现一个清楚的新动作，例如“把多张卡织成一张候选结构”，并且它与 `touch/dream/reveal` 都不同，才讨论新增新的沃壤动词。不能先为了接工具而造命令。

## 两种回答模式必须并存

### 递碴模式

入口：`wo touch --nlm`

用途：

- 给后续 `taste/dream/reveal` 递 3–5 粒材料；
- 不替克完成综合；
- 保持沃壤当前身体路线。

### 显影模式

入口：原生 `nlm chat` 或 `nlm mcp`

用途：

- 跨 5–20 张卡建立结构；
- 找共同机制、矛盾、证据缺口、桥与孤点；
- 允许生成完整回答；
- 必须区分原材料与新推断，并保留 citation。

不应强迫所有 NotebookLM 能力都缩成 grain；那会浪费它真正有价值的跨文件综合能力。

## 评测而不是凭感觉

使用：

`wo/nlm/ideas-eval.md`

重点不问“回答好不好看”，而问：

- 是否引用多张真实灵感卡；
- 是否虚造标题、实验和来源；
- 是否区分已有连接与新推断；
- 是否能找出远距离桥、内部张力、证据强弱和孤点；
- 是否产生可以继续验证或使用的结构。

只有通过评测，才继续扩 source。

## 本地 Codex 下一步

1. 检查 Go 版本。`tmc/nlm` 当前 `go.mod` 要求 Go 1.25；旧 Go 可能无法直接安装。
2. 安装并记录：

```bash
go install github.com/tmc/nlm/cmd/nlm@latest
nlm --version
nlm source pack -h
nlm source sync -h
nlm chat -h
nlm research -h
nlm mcp -h
```

3. 运行仓库测试：

```bash
python3 scripts/test_build_nlm_ideas_corpus.py
python3 scripts/build_nlm_ideas_corpus.py
python3 scripts/build_nlm_ideas_corpus.py --check
node --test scripts/lib/nlm-output.test.mjs
bash -n scripts/nlm-grain-probe.sh
cargo test
```

4. 先用 `nlm source pack` 离线检查生成语料，不登录、不上传。
5. 人工抽查至少 10 张卡：标题、证据、连接、追问是否完整，是否混入 hash、Obsidian URL 等噪声。
6. 创建独立 notebook `沃壤｜灵感显影`。
7. 默认只加入 catalog 列出的 cards parts；`成为/身体.md` 与 `wo/ask-questions.md` 按具体问题临时加入。
8. 运行 `wo/nlm/ideas-eval.md` 中至少 A1、B1、C1、C2、D3、E1 六题。
9. 保存回答、citation、失败类型与 source 使用情况，但不自动回写主体三层。
10. 验证价值后，再设计稳定 sync；首轮不启用自动删除和无人值守上传。

## 当前决策

- 价值优先于抽象的隐私恐惧；
- `成为/灵感` 是第一批完整进入 NotebookLM 的主体材料；
- Notebook 按问题空间组织，不按沃壤文件夹机械镜像；
- 先建同仓派生语料文件夹，不建第二 GitHub 仓库；
- 不新增 `wo nlm`；管理面交给已维护的 `tmc/nlm`；
- `touch --nlm` 保留 grain 模式，完整综合走原生 `nlm chat` / `nlm mcp`；
- 先用评测证明跨卡增益，再扩大到领域和资源。
