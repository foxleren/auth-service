package service

import (
	"github.com/foxleren/auth-service/backend/internal/models"
	"github.com/foxleren/auth-service/backend/internal/repository"
	"github.com/foxleren/auth-service/backend/pkg/authService"
)

type AuthService struct {
	repo repository.User
}

func NewAuthService(repo repository.User) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) GetUserByEmailAndPassword(email, password string) (*models.User, error) {
	return s.repo.GetUserByEmailAndPassword(email, authService.GenerateHashPassword(password))
}

func (s *AuthService) CreateUser(user *models.UserForCreate) (int, error) {
	user.Password = authService.GenerateHashPassword(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) UpdateUserPassword(email string) (int, error) {
	return s.repo.UpdatePasswordByEmail(email)
}
