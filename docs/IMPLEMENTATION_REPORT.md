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
gofmt -w cmd internal
GOSUMDB=off GOMODCACHE="$PWD/.gomodcache" GOCACHE="$PWD/.gocache" go get github.com/gin-gonic/gin github.com/golang-jwt/jwt/v5 github.com/joho/godotenv golang.org/x/crypto/bcrypt gorm.io/driver/sqlite gorm.io/gorm
GOSUMDB=off GOMODCACHE="$PWD/.gomodcache" GOCACHE="$PWD/.gocache" go test ./cmd/... ./internal/...
GOSUMDB=off GOMODCACHE="$PWD/.gomodcache" GOCACHE="$PWD/.gocache" go test -race ./cmd/... ./internal/...
go run ./cmd/homework01
go run ./cmd/homework02
go run ./cmd/gorm_practice
APP_ENV=test SERVER_PORT=18080 DB_DSN=/private/tmp/metanode_blog_test.db JWT_SECRET=test-secret go run ./cmd/blog
curl -s -i http://localhost:18080/health
```

## 测试结果

- `go test ./cmd/... ./internal/...`：通过。
- `go test -race ./cmd/... ./internal/...`：通过。
- `go run ./cmd/homework01`：通过。
- `go run ./cmd/homework02`：通过。
- `go run ./cmd/gorm_practice`：通过。
- `go run ./cmd/blog` + `/health`：通过，返回 HTTP 200。

## 已知限制

当前目录不是干净 Go 作业仓库，而是既有沃壤仓库。根目录执行 `go test ./...` 和 `go mod tidy` 会扫描中文目录并失败。解决方式是将新增 Go 工程文件提交到干净仓库根目录，或把既有非 Go 内容移出 Go module。
