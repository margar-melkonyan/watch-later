package service

import (
	"github.com/margar-melkonyan/watch-later.git/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repository *repository.UserRepository
}

type AuthServiceInterface interface {
	SignIn(user *repository.UserRepository) (string, error)
	SignUp(user *repository.UserRepository) (string, error)
	SignOut(user *repository.UserRepository) error
}

func NewAuthService(repository *repository.UserRepository) *AuthService {
	return &AuthService{
		repository: repository,
	}
}

func (a *AuthService) SignIn(repository *repository.User) (string, error) {
	return "", nil
}

func (a *AuthService) SignUp(user *repository.User) error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return err
	}
	user.Password = string(password)

	return a.repository.Create(user)
}

func (a *AuthService) SignOut(user *repository.User) error {
	return nil
}
