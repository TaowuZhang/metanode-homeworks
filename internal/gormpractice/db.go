package gormpractice

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&User{}, &Post{}, &Comment{}); err != nil {
		return nil, err
	}
	return db, nil
}

func Seed(db *gorm.DB) (*User, error) {
	user := &User{Username: "alice", Email: "alice@example.com"}
	if err := db.Create(user).Error; err != nil {
		return nil, err
	}
	post1 := Post{Title: "GORM 入门", Content: "模型、迁移和关联", UserID: user.ID}
	post2 := Post{Title: "Hook 实践", Content: "用钩子维护统计字段", UserID: user.ID}
	if err := db.Create(&post1).Error; err != nil {
		return nil, err
	}
	if err := db.Create(&post2).Error; err != nil {
		return nil, err
	}
	comments := []Comment{
		{Content: "好文章", UserID: user.ID, PostID: post1.ID},
		{Content: "学到了", UserID: user.ID, PostID: post1.ID},
		{Content: "继续写", UserID: user.ID, PostID: post2.ID},
	}
	return user, db.Create(&comments).Error
}

func GetUserPostsWithComments(db *gorm.DB, userID uint) ([]Post, error) {
	var posts []Post
	err := db.Preload("Comments").Where("user_id = ?", userID).Order("id asc").Find(&posts).Error
	return posts, err
}

func GetPostWithMostComments(db *gorm.DB) (*Post, int64, error) {
	type row struct {
		PostID uint
		Count  int64
	}
	var r row
	err := db.Model(&Comment{}).
		Select("post_id, count(*) as count").
		Group("post_id").
		Order("count desc").
		Limit(1).
		Scan(&r).Error
	if err != nil {
		return nil, 0, err
	}
	if r.PostID == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}
	var post Post
	if err := db.Preload("Comments").First(&post, r.PostID).Error; err != nil {
		return nil, 0, err
	}
	return &post, r.Count, nil
}
