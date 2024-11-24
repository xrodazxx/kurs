package service

import (
	kurs "Kurs"

	"github.com/xrodazxx/App-ToDO/pkg/repository"
)

type Authorization interface {
	CreateUser(user kurs.User) (int, error)
	GenerateToke(username, password string) (string, error)
}
type DeviceIot interface {
}
type DeviceData interface {
}
type Service struct {
	Authorization
	DeviceData
	DeviceIot
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
