package service

import (
	"github.com/foxleren/auth-service/backend/internal/models"
	"github.com/foxleren/auth-service/backend/internal/repository"
)

type Auth interface {
	GetUserByEmailAndPassword(email, password string) (*models.User, error)
	CreateUser(user *models.UserForCreate) (int, error)
	UpdateUserPassword(userEmail string) (int, error)
}

type Service struct {
	Auth
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repos.User),
	}
}
