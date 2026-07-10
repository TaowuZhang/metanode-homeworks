package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Password string `gorm:"not null" json:"-"`
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Posts    []Post `json:"posts,omitempty"`
}

type Post struct {
	gorm.Model
	Title    string    `gorm:"not null" json:"title"`
	Content  string    `gorm:"not null" json:"content"`
	UserID   uint      `gorm:"not null;index" json:"user_id"`
	User     User      `json:"user,omitempty"`
	Comments []Comment `json:"comments,omitempty"`
}

type Comment struct {
	gorm.Model
	Content string `gorm:"not null" json:"content"`
	UserID  uint   `gorm:"not null;index" json:"user_id"`
	PostID  uint   `gorm:"not null;index" json:"post_id"`
	User    User   `json:"user,omitempty"`
}
