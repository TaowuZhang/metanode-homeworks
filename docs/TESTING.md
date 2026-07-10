# 测试说明

## 自动测试

当前工作区包含既有中文目录，根目录 `go test ./...` 会被 Go import path 扫描限制拦截。验证新增 Go 工程时使用：

```bash
GOSUMDB=off GOMODCACHE="$PWD/.gomodcache" GOCACHE="$PWD/.gocache" go test ./cmd/... ./internal/...
GOSUMDB=off GOMODCACHE="$PWD/.gomodcache" GOCACHE="$PWD/.gocache" go test -race ./cmd/... ./internal/...
```

在干净作业仓库中：

```bash
go mod tidy
go test ./...
go test -race ./...
```

## 手动运行

```bash
go run ./cmd/homework01
go run ./cmd/homework02
go run ./cmd/gorm_practice
go run ./cmd/blog
```

博客 API：

```bash
APP_ENV=development SERVER_PORT=8080 DB_DSN=blog.db JWT_SECRET=local-secret go run ./cmd/blog
curl http://localhost:8080/health
```

## 常见错误

- `operation not permitted`: Go 默认缓存写入 `$HOME` 被沙箱限制。设置 `GOMODCACHE` 和 `GOCACHE` 到当前目录。
- `malformed import path ... invalid char`: 当前沃壤仓库有中文目录，根 module 的 `./...` 会扫描到它们。提交作业到干净仓库后消失。
- `UNIQUE constraint failed`: 重复注册同名用户或邮箱。删除本地 `blog.db` 或换用户名。
- `missing bearer token`: 需要登录后在请求头添加 `Authorization: Bearer <token>`。
