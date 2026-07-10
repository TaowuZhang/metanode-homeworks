# 论文审阅项目本地搬移记录 · 2026-06-05

## 适用分支

```bash
organize-paper-review-project
```

## 当前状态

本文件最初用于给本地二刀提供 `git mv` 命令。后续本地搬移已经完成并推送到 PR 分支，因此本文现在作为迁移记录与复现说明保留，不再是合入前必须执行的待办。

当前 PR 分支应满足：

- 根目录不再有 `论文审阅项目.md`。
- 根目录不再有 `论文审阅项目/`。
- `成为/项目/论文审阅项目.md` 是项目入口。
- `成为/项目/论文审阅项目/` 包含原项目 Markdown 与图片资产。
- `成为/项目/论文审阅项目.md` 内链接指向同级目录 `论文审阅项目/...`。

## 为什么需要本地搬移

原目录含大量 Markdown 与 Notion 导出的图片资产，并且多个 Markdown 通过相对路径引用这些图片。

因此完整搬移应使用 `git mv` 保留目录结构，而不是只在云端重写 Markdown 文件；否则容易出现图片资产遗漏或相对路径断裂。

## 已采用的搬移原则

```plain text
论文审阅项目/
→ 成为/项目/论文审阅项目/
```

同时更新项目入口：

```plain text
成为/项目/论文审阅项目.md
```

链接从旧根目录相对路径：

```plain text
../../论文审阅项目/...
```

改为同级项目目录路径：

```plain text
论文审阅项目/...
```

## 本地执行时遇到的修正

第一次普通 `git fetch origin` 曾因本地当前分支的远端 ref 不存在而失败。

更稳的抓取方式是显式抓取 PR 分支：

```bash
git fetch origin refs/heads/organize-paper-review-project:refs/remotes/origin/organize-paper-review-project
```

然后再切到施工分支：

```bash
git switch organize-paper-review-project || git switch -c organize-paper-review-project origin/organize-paper-review-project
git pull --ff-only origin organize-paper-review-project
```

## 复现命令（仅用于从旧状态重做）

如果将来需要在旧状态上复现本次迁移，可使用以下形状。若分支已经完成迁移，请不要重复执行 `git mv`。

```bash
set -euo pipefail

git remote -v | grep -E 'dongxi-heji/worang(.git)?' >/dev/null
git fetch origin refs/heads/organize-paper-review-project:refs/remotes/origin/organize-paper-review-project
git switch organize-paper-review-project || git switch -c organize-paper-review-project origin/organize-paper-review-project
git pull --ff-only origin organize-paper-review-project

mkdir -p "成为/项目"
git mv "论文审阅项目" "成为/项目/论文审阅项目"

python3 - <<'PY'
from pathlib import Path
p = Path("成为/项目/论文审阅项目.md")
text = p.read_text(encoding="utf-8")
text = text.replace("../../论文审阅项目/", "论文审阅项目/")
text = text.replace("迁移状态: 已从根目录索引移入项目地层；项目正文与资产待完整迁移到 `成为/项目/论文审阅项目/`。", "迁移状态: 已从根目录完整移入项目地层。")
text = text.replace("> 迁移中：下列链接暂时指向旧根目录下的项目文件夹；完整搬移完成后，应改为指向同级目录 `论文审阅项目/`。", "> 已迁移：下列链接指向同级项目目录 `论文审阅项目/`。")
p.write_text(text, encoding="utf-8")
PY

test ! -e "论文审阅项目"
test -d "成为/项目/论文审阅项目"
test -f "成为/项目/论文审阅项目.md"

if grep -R "../../论文审阅项目/" -n "成为/项目/论文审阅项目.md"; then
  echo "仍有旧链接残留，请检查。"
  exit 1
fi

git status --short
git diff --stat
git add -A
git commit -m "chore: move paper review project under becoming projects"
git push origin organize-paper-review-project
```

## 合入前验证

合入前至少确认：

- PR 文件列表显示原 `论文审阅项目/` 下 Markdown 与图片资产均以 rename 方式进入 `成为/项目/论文审阅项目/`。
- `成为/项目/论文审阅项目.md` 不再引用 `../../论文审阅项目/`。
- 根目录旧 `论文审阅项目.md` 在 PR head 不存在。
- 本轮未改动资源旧目录、主资源 CSV、外部平台写回、cookies、`.env`、token 或 secret。
