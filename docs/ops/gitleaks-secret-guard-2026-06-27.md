# Gitleaks 密钥门（2026-06-27）

## 它与 `gh` 的区别

- `gh` 操作 GitHub：PR、Issue、Actions、Release、API。
- `gitleaks` 检查 Git、文件和 stdin 中是否出现密码、API key、token、私钥或凭据文件。

一个是外手，一个是门卫。

## 供应决定

- 上游：`gitleaks/gitleaks`
- 首个固定版本：`v8.30.1`
- 本地 macOS 安装：`brew install gitleaks`
- CI 下载固定 release，并使用上游 checksums 文件校验压缩包。
- 上游已将 Gitleaks 定位为 feature complete，后续主要提供安全补丁；新的 Betterleaks 只进入观察名单，不在本 PR 替换成熟工具。

## 沃壤规则

`.gitleaks.toml` 继承 Gitleaks 内置默认规则，并增加：

1. Google OAuth `client_secret` JSON；
2. 沃壤常用外部工具的敏感环境变量赋值；
3. 常见 credential / service-account / token JSON 文件名。

不要为了让 CI 变绿而大范围 allowlist。确认是假阳性时，应留下最窄的路径、提交或具体规则例外，并说明原因。

## 本地使用

```bash
# 扫描当前目录内容
bash scripts/gitleaks-scan.sh worktree

# 扫描暂存区，适合提交前运行
bash scripts/gitleaks-scan.sh staged

# 扫描全部 Git 历史（第一次应在本地执行）
bash scripts/gitleaks-scan.sh history

# 扫描指定提交区间
bash scripts/gitleaks-scan.sh range main..HEAD
```

输出默认完全遮盖命中的 secret。

## CI 边界

`.github/workflows/gitleaks.yml` 只在面向 `main` 的 PR 上运行，并扫描 base commit 到 PR HEAD 的提交差异。

这是有意的：

- 新门先阻止新增泄漏；
- 不在未审计前让历史遗留项阻塞每个 PR；
- 全历史扫描由本地显式执行，发现问题后先轮换凭据，再处理 Git 历史。

CI 不需要也不读取任何真实凭据。

## 发现泄漏时

1. 先撤销或轮换真实凭据；
2. 再删除工作区中的秘密；
3. 判断是否已进入远端 Git 历史；
4. 必要时重写历史并通知所有克隆者重新同步；
5. 不要只提交一个“删除 secret”的新 commit 后就认为已经安全。

## 本 PR 不做

- 不安装本地 pre-commit hook；
- 不自动重写 Git 历史；
- 不上传扫描内容到第三方服务；
- 不引入 Betterleaks；
- 不把 `gitleaks` 变成 `wo` 顶层动作。
