package services

import (
	"errors"
	"time"

	"metanode-go-backend-homeworks/internal/blog/models"
	"metanode-go-backend-homeworks/internal/blog/repository"
	"metanode-go-backend-homeworks/internal/blog/utils"
)

type AuthService struct {
	users     *repository.UserRepository
	jwtSecret string
	jwtTTL    time.Duration
}

func NewAuthService(users *repository.UserRepository, jwtSecret string, jwtTTL time.Duration) *AuthService {
	return &AuthService{users: users, jwtSecret: jwtSecret, jwtTTL: jwtTTL}
}

func (s *AuthService) Register(username, email, password string) (*models.User, error) {
	if username == "" || email == "" || len(password) < 6 {
		return nil, errors.New("username, email and at least 6 chars password are required")
	}
	hash, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}
	user := &models.User{Username: username, Email: email, Password: hash}
	return user, s.users.Create(user)
}

func (s *AuthService) Login(username, password string) (string, *models.User, error) {
	user, err := s.users.FindByUsername(username)
	if err != nil || !utils.CheckPassword(user.Password, password) {
		return "", nil, errors.New("invalid username or password")
	}
	token, err := utils.GenerateToken(user.ID, s.jwtSecret, s.jwtTTL)
	return token, user, err
}

func (s *AuthService) Profile(userID uint) (*models.User, error) {
	return s.users.FindByID(userID)
}
