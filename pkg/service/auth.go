package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/DanjokLion/todo-go"
	"github.com/DanjokLion/todo-go/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const  (
	salt = "askdkalnlamkdmk2324y7823nkgmrke1231"
	singingKey = "jfhjkgjfdsgajksjfk12u845123j3"
	tokenSSL = 18 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authtorization
}

func NewAuthService(repo repository.Authtorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = genPasswordHash(user.Password)
	return s.repo.CreateUser(user)
}	

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, genPasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
		ExpiresAt: time.Now().Add(tokenSSL).Unix(),
		IssuedAt: time.Now().Unix(),
	},
		user.Id,
	})

	return token.SignedString([]byte(singingKey))
}

func (a *AuthService)  ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid sign method")
		}

		return []byte(singingKey), nil
	})
	if err != nil {
		return 0, nil
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserID, nil
}

func genPasswordHash (password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}