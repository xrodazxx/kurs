package service

import (
	kurs "Kurs"
	"Kurs/pkg/repository"
	"crypto/sha1"
	"fmt"

	"github.com/xrodazxx/App-ToDO/pkg/repository" //xynia
)

const salt = "qwerfvcder290"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}
func (s *AuthService) CreateUser(user kurs.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
