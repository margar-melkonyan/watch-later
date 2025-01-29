package repository

import (
	"database/sql"
	"time"
)

type User struct {
	ID           uint64
	Nickname     string
	Firstname    string
	Lastname     string
	Patronymic   string
	Email        string
	Password     string
	RefreshToken string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

type UserRepositoryInterface interface {
	Get(id uint64) (*User, error)
	GetByNickname(nickname string) (*User, error)
	GetByEmail(email string) (*User, error)
	Create(user *User) error
	Update(user *User, id uint64) error
	Delete(id uint64) error
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}
