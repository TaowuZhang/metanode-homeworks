# 从自然语言到 Go 代码结构

自然语言需求落到代码，核心是识别输入、输出、状态、约束和边界，再选择 Go 中最直接的表达方式。

## 算法题

题目里的“给定数组”“返回下标”“原地删除”分别对应函数参数、返回值和是否允许新建额外数据结构。例如“两数之和”需要从数组和值找到两个下标，适合用 `map[int]int` 保存已见数字；“删除有序数组重复项”强调原地和 O(1)，自然映射为快慢双指针。

## Go 语法映射

- “修改原变量”映射为指针参数，例如 `AddTen(n *int)`。
- “修改切片元素”可以传 `*[]int`，也可以直接传切片；作业要求指针，所以实现为 `DoubleSlice(nums *[]int)`。
- “两个协程同时执行”映射为 `go func()` 和 `sync.WaitGroup`。
- “两个协程通信”映射为 channel，生产者写入，消费者 range 读取。
- “共享计数器”映射为临界区；复杂共享状态用 `sync.Mutex`，单个整数高频递增可用 `sync/atomic`。
- “面向对象”在 Go 中通常不是继承，而是接口和组合：`Shape` 定义行为，`Rectangle` 和 `Circle` 实现行为；`Employee` 组合 `Person` 复用字段。

## GORM 映射

“实体”映射为 struct，“关系”映射为外键字段和关联切片。“一个用户多篇文章”就是 `User.Posts []Post` 和 `Post.UserID uint`。“查询文章及评论”映射为 `Preload("Comments")`。“创建文章后更新文章数”映射为 `Post.AfterCreate` Hook。“删除最后一条评论后状态为无评论”映射为 `Comment.AfterDelete` 中计数并更新 `Post.CommentStatus`。

## 博客系统映射

博客需求按边界拆成层：

- controller：解析 HTTP 请求，返回 JSON。
- service：处理注册、登录、权限判断、业务规则。
- repository：封装 GORM 查询。
- middleware：JWT 鉴权，把 `userID` 放入 Gin context。
- utils：响应格式、JWT、密码哈希。

## 五个具体例子

1. “只有作者才能更新文章”变成 `if post.UserID != userID { return ErrForbidden }`，controller 将它转换成 HTTP 403。
2. “登录后才能创建文章”变成受保护路由组 `protected.Use(middleware.Auth(jwtSecret))`，没有 Bearer token 会返回 401。
3. “一篇文章有多个评论”变成 `Post.Comments []Comment` 和 `Comment.PostID uint`，查询时用 `Preload("Comments")`。
4. “删除最后一条评论后状态为无评论”变成 `AfterDelete` Hook：统计同一 `post_id` 剩余评论数，等于 0 就更新文章状态。
5. “删除有序数组重复项”变成双指针：`fast` 扫描，遇到新值写到 `slow`，最后返回 `slow` 作为新长度。
