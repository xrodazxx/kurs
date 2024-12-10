package service

import (
	kurs "Kurs"

	"Kurs/pkg/repository"
)

type Authorization interface {
	CreateUser(user kurs.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
