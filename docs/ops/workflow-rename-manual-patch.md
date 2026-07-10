# Workflow rename manual patch｜潮汐体系（已完成 · 归档）

> 状态：✅ 已完成。2026-06-21 核实 main 上 `.github/workflows/cockpit-ignition-wire.yml` 的文件名与正文措辞已全部迁移到位，无待手动执行项。本文件留作迁移历史记录，不再是待办。

## 已完成的迁移（核实于 main @ 2026-06-21）

文件 rename：

```plain text
.github/workflows/cockpit-morning-note.yml
→ .github/workflows/cockpit-ignition-wire.yml   ✅
```

正文措辞替换（逐条核实，均已在 main 落地）：

- `name: cockpit-morning-note` → `name: cockpit-ignition-wire`   ✅
- `潮汐底盘 v2` / `双前台底盘` → `潮汐体系`   ✅
- `open cockpit issues` → `open 中枢户口（label: cockpit）`   ✅
- `烘出早晨纸条` → `烘焙发车引线`   ✅
- `cockpit/morning-note` → `cockpit/ignition-wire`（产物分支）   ✅
- `早晨纸条` / `morning note` → `发车引线` / `ignition wire`   ✅
- `docs/cockpit/morning-note.md` → `docs/cockpit/ignition-wire.md`   ✅
- commit message `bake morning note` → `bake ignition wire`   ✅

## 名实核实

- 当前 workflow：每天 04:00（Asia/Shanghai，cron `0 20 * * *`）从 open 中枢户口（label: cockpit）烘焙 `docs/cockpit/ignition-wire.md`，只 force-push `cockpit/ignition-wire` 分支，**永不碰 main**。
- 因此 main / 本地 feature 分支上看不到 `docs/cockpit/ignition-wire.md` 属**设计内**（产物只活在 `cockpit/ignition-wire` 分支），不是「未建」。要看产物：`git fetch origin cockpit/ignition-wire && git show origin/cockpit/ignition-wire:docs/cockpit/ignition-wire.md`。
- 兼容桩：`notion-github-project-ledger-sync-v0.md` 已迁移为指向 `account-phase-lock-v0.md` 的跳转桩；旧脚本名（project-ledger-sync / sync-project1-fields / audit-project1-fields）均已 shim 到新物理命名实现。

## 历史背景（保留，勿当待办）

此前写入工具不能修改 `.github/workflows/*`（GitHub API 返回 `Resource not accessible by integration`），该 workflow 正文措辞迁移一度挂为「待具备 workflow scope 执行」。**该前置条件已不再适用**——迁移已由有权限的提交完成（rename 见 PR #251）。原手动修复命令（python3 / perl 脚本）已无需执行，故从本文件移除，以免被误读成仍需操作。
