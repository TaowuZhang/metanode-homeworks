# 测试说明

## 自动测试

```bash
go mod tidy
go vet ./...
go test ./...
go test -race ./...
```

以上命令均已成功执行。

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
- `UNIQUE constraint failed`: 重复注册同名用户或邮箱。删除本地 `blog.db` 或换用户名。
- `missing bearer token`: 需要登录后在请求头添加 `Authorization: Bearer <token>`。

## 验证结果

- `go mod tidy`：成功。
- `go test ./...`：成功。
- `go test -race ./...`：成功。
- homework01、homework02、gorm_practice：成功运行。
- blog 服务：成功启动。
- `GET /health`：返回 HTTP 200。
