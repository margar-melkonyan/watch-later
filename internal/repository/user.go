package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type User struct {
	ID                   uint64
	Nickname             string `validate:"required,min=3,max=255"`
	Firstname            string `validate:"required,min=3,max=255"`
	Lastname             string `validate:"required,min=3,max=255"`
	Patronymic           string
	Email                string `validate:"required,email,min=8,max=32"`
	Password             string `validate:"required,eqfield=PasswordConfirmation,min=8,max=32"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,min=8,max=32"`
	RefreshToken         string
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            *time.Time
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
	var user User
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
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Get(id uint64) (*User, error) {
	return r.getUserByQuery(
		"SELECT * FROM users WHERE id=$1 AND deleted_at IS NULL",
		id,
	)
}

func (r *UserRepository) GetByNickname(nickname string) (*User, error) {
	return r.getUserByQuery(
		"SELECT * FROM users WHERE nickname=$1 AND deleted_at IS NULL",
		nickname,
	)
}

func (r *UserRepository) GetByEmail(email string) (*User, error) {
	user, err := r.getUserByQuery(
		"SELECT * FROM users WHERE email=$1 AND deleted_at IS NULL",
		&email,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err // Любая другая ошибка (например, проблема с БД)
	}

	return user, nil
}

func (r *UserRepository) Create(user *User) error {
	result, err := r.db.Exec(
		`INSERT INTO users (nickname, first_name, last_name, patronymic_name, email, password, refresh_token) 
			   VALUES ($1, $2, $3, $4, $5, $6, $7)`,
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
		return errors.New("user was not created")
	}

	return nil
}

func (r *UserRepository) Update(user *User, id uint64) error {
	result, err := r.db.Exec(
		`UPDATE users 
			   SET nickname=$1, first_name=$2, last_name=$3, patronymic_name=$4, email=$5, password=$6, refresh_token=$7, updated_at=now() 
               WHERE id=$8 AND deleted_at IS NULL`,
		user.Nickname,
		user.Firstname,
		user.Lastname,
		user.Patronymic,
		user.Email,
		user.Password,
		user.RefreshToken,
		id,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user was not updated")
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
		return errors.New(fmt.Sprintf("user was not deleted"))
	}

	return nil
}
