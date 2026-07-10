package gormpractice

import "gorm.io/gorm"

const (
	CommentStatusHasComments = "有评论"
	CommentStatusNoComments  = "无评论"
)

type User struct {
	gorm.Model
	Username  string `gorm:"uniqueIndex;not null"`
	Email     string `gorm:"uniqueIndex;not null"`
	PostCount int
	Posts     []Post
}

type Post struct {
	gorm.Model
	Title         string `gorm:"not null"`
	Content       string `gorm:"not null"`
	UserID        uint   `gorm:"not null;index"`
	User          User
	CommentStatus string `gorm:"not null;default:无评论"`
	Comments      []Comment
}

type Comment struct {
	gorm.Model
	Content string `gorm:"not null"`
	UserID  uint   `gorm:"not null;index"`
	PostID  uint   `gorm:"not null;index"`
}

func (p *Post) AfterCreate(tx *gorm.DB) error {
	return tx.Model(&User{}).Where("id = ?", p.UserID).UpdateColumn("post_count", gorm.Expr("post_count + ?", 1)).Error
}

func (c *Comment) AfterCreate(tx *gorm.DB) error {
	return tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", CommentStatusHasComments).Error
}

func (c *Comment) AfterDelete(tx *gorm.DB) error {
	var count int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", CommentStatusNoComments).Error
	}
	return nil
}
