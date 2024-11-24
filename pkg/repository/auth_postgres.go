package repository

import (
	kurs "Kurs"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

const usersTable = "users"

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user kurs.User) (int, error) {
	var id int

	// Строка запроса для вставки пользователя
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", usersTable)

	// Выполнение запроса
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
func (r *AuthPostgres) GetUser(username, password string) (kurs.User, error) {
	var user kurs.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2 ", usersTable)
	err := r.db.Get(&user, query, username, password)
	return user, err
}
