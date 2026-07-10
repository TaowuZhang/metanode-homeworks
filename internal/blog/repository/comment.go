package repository

import (
	"github.com/TaowuZhang/metanode-go-homeworks/internal/blog/models"

	"gorm.io/gorm"
)

type CommentRepository struct{ db *gorm.DB }

func (r *CommentRepository) Create(comment *models.Comment) error {
	return r.db.Create(comment).Error
}

func (r *CommentRepository) ListByPost(postID uint) ([]models.Comment, error) {
	var comments []models.Comment
	err := r.db.Preload("User").Where("post_id = ?", postID).Order("created_at asc").Find(&comments).Error
	return comments, err
}
