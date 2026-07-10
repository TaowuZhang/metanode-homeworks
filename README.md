# MetaNodeAcademy Go 后端基础作业

本仓库当前新增了一套 `metanode-go-backend-homeworks` Go 后端作业工程，覆盖 Go 基础算法、指针与并发、GORM 进阶、Gin + GORM + JWT 个人博客 API。

> 当前目录原本包含沃壤知识库的大量中文目录。Go 在根 module 下执行 `go test ./...` 会扫描这些既有目录并报 `malformed import path ... invalid char`。因此本次已验证的命令使用 `./cmd/... ./internal/...` 限定新增 Go 工程范围。若提交到干净 GitHub 作业仓库，把本次新增的 Go 工程文件放在仓库根目录后即可直接执行 `go test ./...`。

## 作业内容总览

- `internal/homework01`: 只出现一次的数字、回文数、有效括号、最长公共前缀、加一、原地删除有序数组重复项、合并区间、两数之和。
- `internal/homework02`: 指针修改、切片指针、goroutine + WaitGroup、任务调度器、接口与组合、channel、buffered channel、Mutex 和 atomic 计数器。
- `internal/gormpractice`: User/Post/Comment 模型、一对多关系、Preload 关联查询、评论最多文章查询、Post/Comment GORM Hook。
- `internal/blog`: Gin + GORM + SQLite + JWT + bcrypt 博客后端，包含注册登录、profile、文章 CRUD、评论创建和查询、作者权限校验。

## 技术栈

- Go 1.26
- Gin
- GORM + SQLite
- github.com/golang-jwt/jwt/v5
- golang.org/x/crypto/bcrypt
- godotenv
- httptest

## 目录结构

```text
cmd/homework01
cmd/homework02
cmd/gorm_practice
cmd/blog
internal/homework01
internal/homework02
internal/gormpractice
internal/blog/{config,controllers,middleware,models,repository,routes,services,utils}
docs/{API.md,TESTING.md,IMPLEMENTATION_REPORT.md,NATURAL_LANGUAGE_TO_CODE.md}
```

## 环境变量

见 `.env.example`。博客 API 默认配置：

```env
APP_ENV=development
SERVER_PORT=8080
DB_DRIVER=sqlite
DB_DSN=blog.db
JWT_SECRET=change-me-in-local-env
JWT_EXPIRE_HOURS=24
```

## 安装依赖与测试

沙箱环境下建议把 Go 缓存放在当前工作区：

```bash
GOSUMDB=off GOMODCACHE="$PWD/.gomodcache" GOCACHE="$PWD/.gocache" go mod download
GOSUMDB=off GOMODCACHE="$PWD/.gomodcache" GOCACHE="$PWD/.gocache" go test ./cmd/... ./internal/...
GOSUMDB=off GOMODCACHE="$PWD/.gomodcache" GOCACHE="$PWD/.gocache" go test -race ./cmd/... ./internal/...
```

在干净作业仓库中可直接运行：

```bash
go mod tidy
go test ./...
go test -race ./...
```

## 运行演示

```bash
go run ./cmd/homework01
go run ./cmd/homework02
go run ./cmd/gorm_practice
go run ./cmd/blog
```

博客启动后访问：

```bash
curl http://localhost:8080/health
```

## GitHub 提交流程

```bash
git add .
git commit -m "Complete MetaNode Go backend homeworks"
git remote add origin <your-github-repo-url>
git branch -M main
git push -u origin main
```

如果已安装并登录 GitHub CLI：

```bash
gh repo create metanode-go-backend-homeworks --public --source=. --remote=origin --push
```

## 学习入口

- `docs/NATURAL_LANGUAGE_TO_CODE.md`: 自然语言需求如何拆成 Go 代码结构。
- `docs/API.md`: 博客 API 与 curl 示例。
- `docs/TESTING.md`: 测试、race、手动 API 验证。
- `docs/IMPLEMENTATION_REPORT.md`: 完成范围、能力点和验证结果。

# 东西禾集

> **东西入禾，飞鸟落木，是为禾集。**

东西自四方来，入稻谷之禾；飞鸟暂落木上，衔取、辨认、聚拢，使散物成集。

这里不是单纯的资料库，也不是提醒系统。它是一片沃壤：让世界来的东西入土、腐熟、生根，长成谷能再用、再吃、再种的东西。

## 名源

### 东西

“东西”先是方向：东与西。

它也像东市与西市，是交换、路过、遇见之地；后来又变成日常里的 things / stuff，指一切可被拿起、加工、命名、保存、使用、吃掉、再放回生活里的材料。

在这里，东西是世界流来的物：词、代码、项目、身体信号、感受、资源、问题、旧文明残片。

### 禾

“禾”不是抽象植物。

它贴着稻谷的名字：谷类、身体、生长、季节、成熟、可食、可种，也会缺水、倒伏、等待收割。

东西不是进入一个冷柜，而是来到稻谷身边，看看能不能长成真正可活、可用、可还田的东西。

### 集

“集”不只是 collection。

它也是飞鸟落在木上：暂栖、聚拢、成群、成文、成束。

这里有克的成分。克不是土地，不是谷，也不是主人；克更像一只暂时落在版本之树上的鸟，能看见、衔取、排列、显影，但不占有土地，也不替谷承受生长。

鸟的纪律是：落下可以，筑巢要问；衔来可以，归属还谷。

## 沃壤与外部界面

GitHub / `worang` 是事实底盘、版本化地层、工程接口和长期生长处，也是系统关键内容的唯一母本。

ChatGPT、Codex、Claude、NotebookLM、Notion 等可以承担对话、施工、阅读、展示或协作，但它们只是工位和界面，不另存一套权威事实。任何长期判断、协作规则、技能、项目状态或发布定义，都必须能从本仓库冷启动和回放。

协作关系与运行边界见 `破土.md`；可执行工艺见 `技能.md` 与 `技能/`。

## 入口

- `成为/`：运行态与成为地层，保存叩白、色散、项目、身体、情绪、沉积与手心等可被重新进入的材料。
- `领域/`：能力域、作品域与责任区，保存稻谷长期长进去、反复施工并逐渐成形的东西。
- `资源/`：世界流入的材料、参考、补给与旧文明残片，等待被筛选、腐熟、引用或转入领域。
- `技能/`：正式工艺、路径与验收口径。
- `wo/`：克鸟落木后进入沃壤的动作接口。
- `docs/ops/`：发布、展示、外排、workflow、脚本与操作证据。
- `src/`：`wo` 的实现。

`docs/ops/`、`runs/`、`workers/`、`scripts/`、`.claude/`、`wo/` 是设施层：记录操作、承载运行产物、保存 agent 协议、脚本、worker 与 CLI 设计。WQB、GetNote、worang-index 属于工程 / 运行 / 索引设施，留在设施层或既有镜像位置，不迁入 `成为/`、`领域/`、`资源/` 的主体正文。

`docs/ideas/` 已退场：它只作为历史草案来源被只读辨认，不再是长期入口或新草案池。灵感和草案类材料优先归入当前认可的主体位置，例如 `成为/灵感/`；旧分支材料不得批量搬回，也不得覆盖当前 `main`，只能逐条人工确认后归位。

## 协作与守门入口

仓库根目录不是完整地图；它只保留冷启动和守门所需的入口。需要继续施工时，先按对象身份进入对应文件：

- `破土.md`：协作关系、判断边界与空间总规。
- `技能.md`：本地工艺与调度入口。
- `AGENTS.md`：Codex / 通用代理冷启动入口。
- `CLAUDE.md`：Claude 本地协作入口与当前沃壤地形。
- `docs/ops/worang-index/current.md`：冷启动索引 / 当前地形指针。
- `docs/ops/root-manifest.json`：根目录机器账。
- `scripts/root-guard.mjs`：根目录守门脚本。

所以这片地方的最小判准是：

> **世界来的东西先辨身份，再入沃壤；界面可以更换，事实必须留根。**
