package authService

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/foxleren/auth-service/backend/internal/models"
	"math/rand"
	"os"
	"time"
)

const tokenTTL = 512 * time.Hour

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func GenerateToken(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{jwt.StandardClaims{
		ExpiresAt: time.Now().Add(tokenTTL).Unix(),
		IssuedAt:  time.Now().Unix(),
	}, userId})

	return token.SignedString([]byte(os.Getenv("SIGNING_KEY")))
}

func GenerateHashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	tmp := fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("SALT_KEY"))))

	return tmp
}

func GeneratePassword(passwordLen int) string {
	rand.Seed(time.Now().UnixNano())
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	password := make([]byte, passwordLen)
	for i := 0; i < passwordLen; i++ {
		password[i] = characters[rand.Intn(len(characters))]
	}
	return string(password)
}

func ParseToken(accessToken string) (*models.UserDataFromToken, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(os.Getenv("SIGNING_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return nil, errors.New("token claims are not of type *tokenClaims")
	}

	userData := models.UserDataFromToken{
		Id: claims.UserId,
	}

	return &userData, nil
}
