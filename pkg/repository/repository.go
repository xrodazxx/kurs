package repository

import (
	kurs "Kurs"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user kurs.User) (int, error)
	GetUser(username, password string) (kurs.User, error)
}
type Repository struct {
	Authorization
	IDevice
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		IDevice:       NewDevice(db),
	}
}
