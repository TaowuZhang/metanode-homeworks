package repository

import (
	"metanode-go-backend-homeworks/internal/blog/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
}

type Repositories struct {
	Users    *UserRepository
	Posts    *PostRepository
	Comments *CommentRepository
}

func NewRepositories(db *gorm.DB) Repositories {
	return Repositories{
		Users:    &UserRepository{db: db},
		Posts:    &PostRepository{db: db},
		Comments: &CommentRepository{db: db},
	}
}
