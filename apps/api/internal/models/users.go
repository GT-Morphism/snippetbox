package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	HashedPassword []byte    `json:"-"`
	Created_at     time.Time `json:"-"`
}

type UserModel struct {
	DB *pgxpool.Pool
}

func (m *UserModel) Insert(name, email, password string) error {
	return nil
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

func (m *UserModel) Exists(id int) (bool, error) {
	return false, nil
}
