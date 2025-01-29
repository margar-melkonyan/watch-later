package repository

import (
	"database/sql"
	"errors"
	"fmt"
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

func (r *UserRepository) getUserByQuery(query string, arg interface{}) (*User, error) {
	user := &User{}
	err := r.db.QueryRow(query, arg).
		Scan(&user.ID,
			&user.Nickname,
			&user.Firstname,
			&user.Lastname,
			&user.Patronymic,
			&user.Email,
			&user.Password,
			&user.RefreshToken,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt,
		)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Get(id uint64) (*User, error) {
	return r.getUserByQuery(
		"SELECT * FROM users WHERE id=$1",
		id,
	)
}

func (r *UserRepository) GetByNickname(nickname string) (*User, error) {
	return r.getUserByQuery(
		"SELECT * FROM users WHERE nickname=$1",
		nickname,
	)
}

func (r *UserRepository) GetByEmail(email string) (*User, error) {
	return r.getUserByQuery(
		"SELECT * FROM users WHERE email=$1",
		email,
	)
}

func (r *UserRepository) Create(user *User) error {
	result, err := r.db.Exec(
		"INSERT INTO users (nickname, first_name, last_name, patronymic_name, email, password, refresh_token) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		user.Nickname,
		user.Firstname,
		user.Lastname,
		user.Patronymic,
		user.Email,
		user.Password,
		user.RefreshToken,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("Nothing to create")
	}

	return nil
}

func (r *UserRepository) Update(user *User, id uint64) error {
	result, err := r.db.Exec(
		"UPDATE users SET nickname=$1, first_name=$2, last_name=$3, patronymic_name=$4, email=$5, password=$6, updated_at=now()   WHERE id=$7",
		user.Nickname,
		user.Firstname,
		user.Lastname,
		user.Patronymic,
		user.Email,
		user.Password,
		id,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return fmt.Errorf("expected 1 row affected, got %d", rowsAffected)
	}

	return nil
}

func (r *UserRepository) Delete(id uint64) error {
	result, err := r.db.Exec("UPDATE users SET deleted_at = now() WHERE id = $1 AND deleted_at IS NULL", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New(fmt.Sprintf("user with %d not found", id))
	}

	return nil
}
