# 博客 API 说明

统一成功响应：

```json
{"code":0,"message":"success","data":{}}
```

统一错误响应：

```json
{"code":40001,"message":"invalid username or password"}
```

| 方法 | 路径 | 认证 | 说明 |
|---|---|---|---|
| GET | `/health` | 否 | 健康检查 |
| POST | `/api/v1/auth/register` | 否 | 注册 |
| POST | `/api/v1/auth/login` | 否 | 登录并返回 JWT |
| GET | `/api/v1/profile` | 是 | 当前用户信息 |
| GET | `/api/v1/posts` | 否 | 文章列表 |
| GET | `/api/v1/posts/:id` | 否 | 文章详情 |
| POST | `/api/v1/posts` | 是 | 创建文章 |
| PUT | `/api/v1/posts/:id` | 是 | 作者更新文章 |
| DELETE | `/api/v1/posts/:id` | 是 | 作者删除文章 |
| GET | `/api/v1/posts/:id/comments` | 否 | 查看文章评论，`:id` 即题目中的 `post_id` |
| POST | `/api/v1/posts/:id/comments` | 是 | 创建文章评论，`:id` 即题目中的 `post_id` |

## curl 示例

```bash
curl http://localhost:8080/health

curl -X POST http://localhost:8080/api/v1/auth/register \
  -H 'Content-Type: application/json' \
  -d '{"username":"alice","email":"alice@example.com","password":"password123"}'

TOKEN=$(curl -s -X POST http://localhost:8080/api/v1/auth/login \
  -H 'Content-Type: application/json' \
  -d '{"username":"alice","password":"password123"}' \
  | sed -n 's/.*"token":"\([^"]*\)".*/\1/p')

curl -X POST http://localhost:8080/api/v1/posts \
  -H "Authorization: Bearer $TOKEN" \
  -H 'Content-Type: application/json' \
  -d '{"title":"第一篇文章","content":"这是文章内容"}'

curl http://localhost:8080/api/v1/posts
curl http://localhost:8080/api/v1/posts/1

curl -X PUT http://localhost:8080/api/v1/posts/1 \
  -H "Authorization: Bearer $TOKEN" \
  -H 'Content-Type: application/json' \
  -d '{"title":"更新后的标题","content":"更新后的内容"}'

curl -X POST http://localhost:8080/api/v1/posts/1/comments \
  -H "Authorization: Bearer $TOKEN" \
  -H 'Content-Type: application/json' \
  -d '{"content":"写得很好"}'

curl http://localhost:8080/api/v1/posts/1/comments

curl -X DELETE http://localhost:8080/api/v1/posts/1 \
  -H "Authorization: Bearer $TOKEN"
```

## 请求参数

注册：

```json
{"username":"alice","email":"alice@example.com","password":"password123"}
```

登录：

```json
{"username":"alice","password":"password123"}
```

文章创建/更新：

```json
{"title":"第一篇文章","content":"这是文章内容"}
```

评论创建：

```json
{"content":"写得很好"}
```
