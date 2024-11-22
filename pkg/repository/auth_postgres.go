package repository

import (
	kurs "Kurs"

	"github.com/jwoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
func (r *AuthPostgres) CreateUser(user kurs.User) (int, error) {

}
