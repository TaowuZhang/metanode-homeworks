package services

import (
	"errors"

	"metanode-go-backend-homeworks/internal/blog/models"
	"metanode-go-backend-homeworks/internal/blog/repository"
)

var ErrForbidden = errors.New("forbidden")

type PostService struct {
	posts *repository.PostRepository
}

func NewPostService(posts *repository.PostRepository) *PostService {
	return &PostService{posts: posts}
}

func (s *PostService) Create(userID uint, title, content string) (*models.Post, error) {
	if title == "" || content == "" {
		return nil, errors.New("title and content are required")
	}
	post := &models.Post{Title: title, Content: content, UserID: userID}
	return post, s.posts.Create(post)
}

func (s *PostService) List() ([]models.Post, error) {
	return s.posts.List()
}

func (s *PostService) Get(id uint) (*models.Post, error) {
	return s.posts.FindByID(id)
}

func (s *PostService) Update(id, userID uint, title, content string) (*models.Post, error) {
	post, err := s.posts.FindByID(id)
	if err != nil {
		return nil, err
	}
	if post.UserID != userID {
		return nil, ErrForbidden
	}
	if title != "" {
		post.Title = title
	}
	if content != "" {
		post.Content = content
	}
	return post, s.posts.Update(post)
}

func (s *PostService) Delete(id, userID uint) error {
	post, err := s.posts.FindByID(id)
	if err != nil {
		return err
	}
	if post.UserID != userID {
		return ErrForbidden
	}
	return s.posts.Delete(post)
}
