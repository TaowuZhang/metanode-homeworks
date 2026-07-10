package services

import (
	"errors"

	"github.com/TaowuZhang/metanode-go-homeworks/internal/blog/models"
	"github.com/TaowuZhang/metanode-go-homeworks/internal/blog/repository"
)

type CommentService struct {
	posts    *repository.PostRepository
	comments *repository.CommentRepository
}

func NewCommentService(posts *repository.PostRepository, comments *repository.CommentRepository) *CommentService {
	return &CommentService{posts: posts, comments: comments}
}

func (s *CommentService) Create(userID, postID uint, content string) (*models.Comment, error) {
	if content == "" {
		return nil, errors.New("content is required")
	}
	if _, err := s.posts.FindByID(postID); err != nil {
		return nil, err
	}
	comment := &models.Comment{Content: content, UserID: userID, PostID: postID}
	return comment, s.comments.Create(comment)
}

func (s *CommentService) List(postID uint) ([]models.Comment, error) {
	if _, err := s.posts.FindByID(postID); err != nil {
		return nil, err
	}
	return s.comments.ListByPost(postID)
}
