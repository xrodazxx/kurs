package service

import "github.com/xrodazxx/App-ToDO/pkg/repository"

type Authorization interface {
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
	return &Service{}
}
