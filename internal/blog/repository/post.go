package repository

import (
	"metanode-go-backend-homeworks/internal/blog/models"

	"gorm.io/gorm"
)

type PostRepository struct{ db *gorm.DB }

func (r *PostRepository) Create(post *models.Post) error {
	return r.db.Create(post).Error
}

func (r *PostRepository) List() ([]models.Post, error) {
	var posts []models.Post
	err := r.db.Preload("User").Order("created_at desc").Find(&posts).Error
	return posts, err
}

func (r *PostRepository) FindByID(id uint) (*models.Post, error) {
	var post models.Post
	if err := r.db.Preload("User").First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) Update(post *models.Post) error {
	return r.db.Save(post).Error
}

func (r *PostRepository) Delete(post *models.Post) error {
	return r.db.Delete(post).Error
}
