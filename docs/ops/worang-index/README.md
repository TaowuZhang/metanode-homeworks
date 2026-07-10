# 沃壤 Index 维护说明

本目录承接 `worang` 的 AI 冷启动索引。

它不是全仓目录，也不是根目录白名单的替代品。

## 文件分工

- `research-2026-06-20.md`：本轮研究记录，说明为什么这样做。
- `index.schema.json`：机器索引字段契约。
- `current.json`：机器生成索引，允许完整记录全仓事实。
- `current.md`：AI 冷启动人读薄入口，根部必须精要。
- `workflow-handoff.md`：当 workflow 无法直接落地时的交接稿。

`current.json` 的 `git_ref` 是生成时观察到的 source ref，用于说明生成器当时读取的来源；它不是、也不可能承诺为“包含该生成物的 commit”。提交生成物会产生新的 commit SHA，因此消费者不得把这个字段解释为 containing commit。

当索引语义内容未变化时，生成器保留现有 `generated_at` 与 `git_ref`，使重复生成逐字节稳定；当 entries、summary、manifest、warnings 或其他索引语义内容变化时，生成器刷新观察时间和 source ref。

## 维护命令

```bash
node scripts/build-worang-index.mjs
node scripts/check-worang-index.mjs
node scripts/root-guard.mjs
```

完整验证时再跑：

```bash
cargo test
```

## 根部精要规则

`current.json` 可以完整。

`current.md` 必须薄，尤其根部只能保留：

- 必要根。
- 一句话职责。
- 第一跳。
- 机器发现的异常。

不得在 `current.md` 里展开：

- 完整根目录清单。
- `root-entry-allowlist` 正文。
- `root-manifest.json` 全量内容。
- 工程配置、锁文件、workflow、脚本列表。

## 自动维护边界

workflow 可以自动维护：

- `current.json`
- `current.md`
- 可唯一确定目标的仓内相对链接

workflow 只报警，不自动改：

- 语义入口。
- 根目录白名单。
- Notion 链接。
- 外部 URL。
- 需要谷判断的归位。

workflow 不直推 `main`，只开维护 PR。

## 不手改生成物

不要手动同步 `current.json`。要改规则，改脚本或 schema，再重新生成。

`current.md` 虽然可读，但也是生成物；需要改变呈现时，改 `scripts/build-worang-index.mjs` 的渲染规则。
