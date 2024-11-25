package service

import (
	kurs "Kurs"
	"Kurs/pkg/repository"
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	salt       = "qwerfvcder290"
	tokenTTL   = 12 * time.Hour       // Время жизни токена
	signingKey = "2j34FACQCJFJS6468F" // ключ для подписи токенов
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// создание нового пользователя, хешируя его пароль
func (s *AuthService) CreateUser(user kurs.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

// генерируем JWT токен для аутентификации
func (s *AuthService) GenerateToken(username, password string) (string, error) {
	// Проверяем, пользователь с такими данными
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	// claims for tokens
	claims := tokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)), // Время истечения токена
			IssuedAt:  jwt.NewNumericDate(time.Now()),               // Время создания токена
		},
		UserId: user.Id,
	}

	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signingKey))
}

// парсинг токена
func (s *AuthService) ParseToken(accesstoken string) (int, error) {
	token, err := jwt.ParseWithClaims(accesstoken, &tokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			// проверка подписи токена
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("Invalid signing method")
			}
			return []byte(signingKey), nil
		})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type")
	}
	return claims.UserId, nil
}

// хеширование пароля пользователя с солью
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
