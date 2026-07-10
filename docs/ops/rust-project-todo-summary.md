# Rust 项目待办汇总 · 沃壤 (wo)

> 生成日期：2026-04-30  
> 更新：2026-06-13，旧 `成为/项目.csv` 已退场；本文中“项目”只指 Rust/CLI 工作事项，不再指向旧成为项目表。
> 来源：源码注释 + 设计文档 + 已知缺失功能

---

## 总体判断

**已完成核心链路**：`spark` / `ask` / `make` / `tend` / `sense` 均已实现 Rust 最小入口，并通过了 checkpoint 验收。

**当前阶段**：基础命令层已就绪，进入深化与扩展阶段。重点在 NLM 接入、情绪库挂钩、数据持久化与功能增强。

---

## 待办汇总表

### ✅ 已完成 (Completed)

| 模块 | 状态 | 待办事项 | 优先级 | 备注 |
|------|------|----------|--------|------|
| 整体架构 | ✅ 完成 | Rust CLI 骨架搭建，命令注册机制 | - | 基于 clap 的子命令架构已稳定 |
| spark | ✅ 完成 | 火柴流程七轮 + 回声实现 | - | 状态存储在 `target/wo-spark/`，烧完即清 |
| ask | ✅ 完成 | 好奇地图五入口（空间/设计/人/关系/做事） | - | 明确为地图而非问题库 |
| make | ✅ 完成 | 极轻入口（泥在案上，有轨径行无轨捏泥） | - | 已停止读取旧 `成为/项目.csv`；行动户口看 GitHub Project 1 / cockpit issue |
| tend | ✅ 完成 | 三问校向（当初以为/现在脚下/目标位置） | - | 第一刀即用 Rust 实现 |
| sense | ✅ 完成 | 最小读法（叩白 + 探筹 + 色散旧痕 + 风铃） | - | 第二刀小步接土已完成 |

---

### 🔄 进行中 (In Progress)

| 模块 | 状态 | 待办事项 | 优先级 | 备注 |
|------|------|----------|--------|------|
| emotion (情绪库) | 🔄 进行中 | **情绪弦接入情绪库**（显式 TODO） | **高** | `sense.rs:32` 标注 TODO；需读取 `成为/情绪.csv`，并将新触碰的关系温度写回 |
| NLM | 🔄 进行中 | 验证 `scripts/nlm-grain-probe.sh` 可用性 | 中 | touch.rs 已调用，但依赖外部 `nlm` CLI 工具，需 sanity check |
| NLM | 🔄 进行中 | 扩展 NLM source packs 覆盖范围 | 中 | 旧 `成为｜项目总览` pack 已退休；后续补充领域/灵感/情绪等 key source |

---

### 📋 未开始 (Not Started)

| 模块 | 状态 | 待办事项 | 优先级 | 备注 |
|------|------|----------|--------|------|
| palm 持久化 | ✅ 完成 | 已从旧 `target/` 临时区迁移至稳定位置 | - | 当前读写 `成为/手心/current.md`；`palm.rs` 与 `sense.rs` 已同步 |
| smell 配置化 | ⚠️ 待办 | 将 `SCENT_WORDS` 数组外部可配置 | 低 | `smell.rs:8-45` 硬编码 36 个关键词，可改为 `wo/scent-words.txt` 或 CSV |
| trace 功能增强 | ⚠️ 待办 | 实现 trace 的 edit / delete / query 操作 | 中 | 当前仅 `append_line`，无法编辑或删除历史色散，也无法查询 |
| make 任务选择 | ⚠️ 待办 | `make` 保持轻入口；如需下一步推荐，只能转向 GitHub Project 1 / cockpit issue | 低 | 不再从旧成为项目表自动推荐 |
| 命令结构重构 | ⚠️ 待办 | 评估 `taste` / `dream` / `reveal` 是否拆为一级命令 | 低 | `sense.md:405-416` 提及可能拆分，当前为 `Command::Taste|Dream|Reveal` 子命令，保留观察 |
| 文档完整性 | ⚠️ 待办 | 补全缺失设计文档：`wo/collab.md` 等真实设计缺口 | 中 | `field-notes.md` 多次提及 collab 六标签与三/四/五标签，但 `wo/collab.md` 仅 1 行标题；palm 持久化口径已收束到 `成为/手心/current.md`，不再补旧 target 文档 |
| 风铃深层链 | ⚠️ 待办 | 实现风铃对情绪库的接入（hook emotion database） | **高** | `sense.md:32-36` 已规划：风铃第五段需挂钩 `成为/情绪_all.csv`，且支持写回 |
| 冰箱贴 | ⚠️ 待办 | 实现冰箱贴 (fridge magnet) 功能，读取灵感/灯塔等精选卡片 | 低 | `README.md` 提及但未实现；不再从旧项目容器取卡片 |
| touch NLM 对话框 | ⚠️ 待办 | 为 `touch --nlm` 增加完整的 source pack 列表输出 | 低 | 当前 `--nlm` 仅支持 `becoming`，`pack_matches` 已有其他 pack 映射但未公开 |
| emo → trace 闭环 | ⚠️ 待办 | 将 `sense` 产生的 grain/taste/dream/reveal 结果写入色散 | 中 | `palm.rs` 仅暂存，未落土；`trace.rs` 仅支持手动 `cargo run -- "text"` |
| 色散 Schema 扩展 | ⚠️ 待办 | 升级 `trace.rs` 以支持 `成为/色散_all.csv` 完整 schema（线索、判断、象限等） | 中 | 当前只写简易两列 (date, content)，忽略其他 metadata |

---

## 关键风险与开放问题

| 风险点 | 描述 | 建议 |
|--------|------|------|
| NLM 依赖外部工具 | `touch.rs` 调用 `scripts/nlm-grain-probe.sh` → `nlm` CLI，需确保 NotebookLM 可用 | 增加「NLM 不可用时自动回退到本地 source pack」的优雅降级 |
| 情绪库无人值守 | `情绪.csv` 已有 99 词，但 CLI 未接入，情绪弦仍显示「暂不接库」 | 优先实现情绪读取与热度计算，后续支持写回 |
| Palm 持久化旧风险 | 已收束：当前 palm 读写 `成为/手心/current.md`，不再依赖临时构建目录 | 后续只需避免文档回退到旧路径 |
| smell 词表僵化 | 36 词硬编码在源码中，修改需重新编译 | 抽成配置文件，支持热重载 |
| make 的行动户口边界 | `make` 不再显示旧成为项目清单 | 下一步动作从 GitHub Project 1 / cockpit issue 取，不回灌到成为层 |

---

## 优先级建议

**高优先级**（明确 TODO 且影响核心循环）：
1. 情绪弦接入情绪库 → 完成 sense 五段深层链闭环
2. NLM source pack 扩展 → 让远手真正有土可摸
3. Palm 持久化已完成 → 保持 `成为/手心/current.md` 为稳定手心

**中优先级**（功能增强与健壮性）：
4. trace 查询与编辑 → 色散管理基础能力
5. smell 词表配置化 → 降低维护成本
6. make 行动入口边界 → 保持极轻，必要时只指向 GitHub Project 1 / cockpit issue

**低优先级**（渐进优化）：
7. `taste/dream/reveal` 拆分评估
8. 冰箱贴功能实现
9. source pack 选择对话框完善

---

## 文件定位参考

| 功能模块 | 源码文件 | 设计文档 | 数据文件 |
|----------|----------|----------|----------|
| sense | `src/sense.rs` | `wo/sense.md` | `成为/叩白.csv`、`成为/探筹.csv`、`成为/色散_all.csv`、`成为/认知.csv`、`成为/嫁接.csv` |
| touch | `src/touch.rs` | `wo/touch.md` | `wo/nlm/*.md` |
| palm | `src/palm.rs` | `wo/sense.md`（提及） | `成为/手心/current.md` |
| smell | `src/smell.rs` | `wo/smell.md` | `成为/色散.csv` |
| trace | `src/trace.rs` | `wo/sense.md`（提及） | `成为/色散.csv`、`成为/色散_all.csv` |
| make | `src/make.rs` | `wo/make.md` | GitHub Project 1 / cockpit issue；旧 `成为/项目.csv` 已退休 |
| NLM | `scripts/nlm-grain-probe.sh` | `wo/nlm/README.md` | `wo/nlm/*.md`、`wo/nlm/manifest.json` |
| emotion | 未接入 | `wo/sense.md:32` | `成为/情绪.csv`、`成为/情绪_all.csv` |

---

## 下一步行动建议

1. **先完成情绪库接入**：在 `sense.rs` 的 `print_windbell` 中读取 `成为/情绪.csv`，计算当前认知弦对应的情绪热度，并输出「情绪弦：X（热度：N）」
2. **验证 NLM 脚本**：手动运行 `scripts/nlm-grain-probe.sh "测试问题"`，确认 NotebookLM 环境可用，并检查输出是否符合 `· ` grain 格式
3. **Palm 持久化已完成**：后续只需保持 `palm.rs` / `sense.rs` 指向 `成为/手心/current.md`，并清理旧文档残留
4. **扩展 smell 词表**：将 `SCENT_WORDS` 抽到 `wo/scent-words.txt`，每行一词，程序启动时加载
5. **trace query 命令**：新增 `cargo run -- trace --list` / `--date 2026-04-30` / `--edit <id>` 子命令

---

*最后更新：基于源码深入分析完成，涵盖 15 个 Rust 模块 + 12 个设计文档 + 5 个原始脚本；2026-06-13 已移除旧成为项目表口径。*
