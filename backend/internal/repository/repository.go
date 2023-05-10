package repository

import (
	"github.com/foxleren/auth-service/backend/internal/models"
	"github.com/foxleren/auth-service/backend/pkg/smtpService"
	"github.com/jmoiron/sqlx"
)

type User interface {
	GetUserByEmailAndPassword(email, password string) (*models.User, error)
	CreateUser(user *models.UserForCreate) (int, error)
	UpdatePasswordByEmail(email string) (int, error)
}

type Repository struct {
	User
}

func NewRepository(db *sqlx.DB, smtp *smtpService.SmtpProvider) *Repository {
	return &Repository{
		User: NewUserPostgres(db, smtp),
	}
}
