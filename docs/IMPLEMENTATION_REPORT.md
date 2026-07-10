# 实现报告

## 完成模块

- homework01：实现 8 道算法题，并覆盖普通和边界测试。
- homework02：实现指针、goroutine、任务调度、接口、组合、channel、buffered channel、Mutex、atomic，并补测试和演示。
- gormpractice：实现 User/Post/Comment，一对多关系，AutoMigrate，seed，Preload 查询，评论最多文章查询，Post 创建 Hook 和 Comment 删除 Hook。
- blog：实现 Gin 分层后端，包含配置、模型、仓储、服务、控制器、JWT 中间件、响应工具和路由。

## 课程能力点

- Go 基础语法：切片、map、循环、条件、排序、测试。
- 并发：goroutine、WaitGroup、channel、Mutex、atomic。
- 面向对象风格：接口、多态、结构体组合。
- GORM：模型、外键、一对多、Preload、Hook、SQLite 临时测试库。
- Web 后端：REST API、JWT 登录态、bcrypt 密码哈希、分层架构、httptest。

## 已运行命令

```bash
gofmt -w .
go mod tidy
go test ./...
go test -race ./...
go run ./cmd/homework01
go run ./cmd/homework02
go run ./cmd/gorm_practice
APP_ENV=test SERVER_PORT=18080 DB_DSN=/private/tmp/metanode_blog_test.db JWT_SECRET=test-secret go run ./cmd/blog
curl -s -i http://localhost:18080/health
```

## 测试结果

- `go mod tidy`：成功。
- `go test ./...`：通过。
- `go test -race ./...`：通过。
- `go run ./cmd/homework01`：通过。
- `go run ./cmd/homework02`：通过。
- `go run ./cmd/gorm_practice`：通过。
- `go run ./cmd/blog`：服务成功启动；`GET /health` 返回 HTTP 200。
