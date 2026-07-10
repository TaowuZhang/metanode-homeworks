package services

import (
	"time"

	"github.com/TaowuZhang/metanode-go-homeworks/internal/blog/repository"
)

type Services struct {
	Auth     *AuthService
	Posts    *PostService
	Comments *CommentService
}

func NewServices(repos repository.Repositories, jwtSecret string, jwtTTL time.Duration) Services {
	return Services{
		Auth:     NewAuthService(repos.Users, jwtSecret, jwtTTL),
		Posts:    NewPostService(repos.Posts),
		Comments: NewCommentService(repos.Posts, repos.Comments),
	}
}
