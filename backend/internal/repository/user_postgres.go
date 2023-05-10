package repository

import (
	"errors"
	"fmt"
	models2 "github.com/foxleren/auth-service/backend/internal/models"
	"github.com/foxleren/auth-service/backend/pkg/authService"
	"github.com/foxleren/auth-service/backend/pkg/smtpService"
	"github.com/jmoiron/sqlx"
	"github.com/siruspen/logrus"
)

type UserPostgres struct {
	db   *sqlx.DB
	smtp *smtpService.SmtpProvider
}

func NewUserPostgres(db *sqlx.DB, smtp *smtpService.SmtpProvider) *UserPostgres {
	return &UserPostgres{
		db:   db,
		smtp: smtp,
	}
}

func (p *UserPostgres) GetUserByEmailAndPassword(email, passwordHash string) (*models2.User, error) {
	var user models2.User
	var query = fmt.Sprintf("SELECT id, email FROM %s WHERE email=$1 AND password_hash=$2", usersTable)
	err := p.db.Get(&user, query, email, passwordHash)
	if err != nil {
		logrus.Printf("Level: repos; func GetUserByEmailAndPassword(): error while getting user with email: %s", email)
		return nil, err
	}

	logrus.Printf("Level: repos; func GetUserByEmailAndPassword(): models.Auth=%v", user)

	return &user, err
}

func (p *UserPostgres) CreateUser(user *models2.UserForCreate) (int, error) {
	tx, err := p.db.Begin()
	var userId int

	var addUserQuery = fmt.Sprintf(
		"INSERT INTO %s (email, password_hash) values ($1, $2) RETURNING id", usersTable)
	row := tx.QueryRow(
		addUserQuery,
		user.Email,
		user.Password,
	)
	err = row.Scan(&userId)
	if err != nil {
		logrus.Printf("Level: repos; func CreateUser(): error while creating user with email: %s", user.Email)
		tx.Rollback()
		return 0, err
	}

	logrus.Printf("Level: repos; func CreateUser(): models.Auth.id=%d", userId)

	return userId, tx.Commit()
}

func (p *UserPostgres) UpdatePasswordByEmail(email string) (int, error) {
	var exisingUser models2.User
	var checkOperationQuery = fmt.Sprintf("SELECT * FROM %s WHERE email = $1", usersTable)
	err := p.db.Get(&exisingUser, checkOperationQuery, email)
	if err != nil {
		logrus.Printf("Level: repos; func UpdatePasswordByEmail(): error while checking user with email: %s", email)
		return 0, errors.New("user doesn't exist in db")
	}

	password := authService.GeneratePassword(10)
	hashPassword := authService.GenerateHashPassword(password)

	tx, err := p.db.Begin()
	var updateUserQuery = fmt.Sprintf("UPDATE %s SET password_hash = $1 WHERE id = $2", usersTable)
	_, err = tx.Exec(updateUserQuery,
		hashPassword,
		exisingUser.Id,
	)
	if err != nil {
		logrus.Printf("Level: repos; func UpdatePasswordByEmail(): error while updating password for user with email: %s", email)
		tx.Rollback()
		return 0, err
	}

	err = p.smtp.SendEmail(&smtpService.EmailData{
		Recipient: email,
		Subject:   "Новый пароль",
		Content:   fmt.Sprintf("Пароль: %s", password),
	})
	if err != nil {
		logrus.Printf("Level: repos; func UpdatePasswordByEmail(): error while sending password for user with email: %s", email)
		tx.Rollback()
		return 0, err
	}

	return exisingUser.Id, tx.Commit()
}
