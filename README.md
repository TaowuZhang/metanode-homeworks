# MetaNodeAcademy Go 后端基础作业

这是一个独立的 MetaNodeAcademy Go 后端作业仓库，覆盖 Go 基础算法、指针与并发、GORM 进阶，以及 Gin + GORM + JWT 个人博客 API。

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

```bash
go mod tidy
go vet ./...
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
git remote add origin git@github.com:TaowuZhang/metanode-go-homeworks.git
git branch -M main
git push -u origin main
```

如果已安装并登录 GitHub CLI：

```bash
gh repo create metanode-go-homeworks --public --source=. --remote=origin --push
```

## 学习入口

- `docs/NATURAL_LANGUAGE_TO_CODE.md`: 自然语言需求如何拆成 Go 代码结构。
- `docs/API.md`: 博客 API 与 curl 示例。
- `docs/TESTING.md`: 测试、race、手动 API 验证。
- `docs/IMPLEMENTATION_REPORT.md`: 完成范围、能力点和验证结果。
