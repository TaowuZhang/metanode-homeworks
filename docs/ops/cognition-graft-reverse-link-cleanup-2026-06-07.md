# 认知—嫁接反向 UUID 链接清理 2026-06-07

## 范围

本轮只处理 `成为/认知/` 中已经能被目录实证确认的 `蒸馏为` 链接。

它是上一刀 `docs/ops/becoming-cognition-uuid-link-cleanup-2026-06-07.md` 的反向补齐：上一刀清 `成为/嫁接/` → `成为/认知/`，这一刀清 `成为/认知/` → `成为/嫁接/`。

不做：

- 不全仓扫 UUID。
- 不重写认知库或嫁接库正文。
- 不补造缺失对象。
- 不碰 `成为/项目`、`成为/色散`、论文审阅项目内部链。

## 判断

`成为/嫁接/` 当前已有干净文件名，例如：

- `Analyze.md`
- `Reason.md`
- `Identify.md`
- `Categorize.md`
- `Classify.md`
- `Create.md`
- `Design.md`
- `Evaluate.md`
- `Retrieve.md`
- `Interpret.md`
- `Write.md`
- `Describe.md`

因此，`成为/认知/` 里指向这些嫁接页的旧 Notion 导出链接可以从：

```md
../%E5%AB%81%E6%8E%A5/某页 <uuid>.md
```

改为：

```md
../嫁接/某页.md
```

## 已处理对象

- `思.md` → `../嫁接/Analyze.md`
- `推理.md` → `../嫁接/Reason.md`
- `识.md` → `../嫁接/Identify.md`
- `归类.md` → `../嫁接/Categorize.md`、`../嫁接/Classify.md`
- `想.md` → `../嫁接/Create.md`、`../嫁接/Design.md`
- `判.md` → `../嫁接/Evaluate.md`
- `记.md` → `../嫁接/Retrieve.md`
- `解.md` → `../嫁接/Interpret.md`
- `言.md` → `../嫁接/Write.md`、`../嫁接/Describe.md`

## 继续线

下一刀如果继续，不要扩大成全仓 sweep。优先候选：

1. `成为/色散/` 到 `成为/项目/` 的项目引用；
2. `成为/项目/` 内部子页引用；
3. 论文审阅项目内部阶段页引用。

规则不变：每一刀先确认目标文件真实存在，再替换；不明确就留审计说明，不伪造链接。
