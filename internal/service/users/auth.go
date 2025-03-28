package service

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/margar-melkonyan/watch-later.git/internal/common"
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

func (a *AuthService) SignIn(user *common.SignInUser) (map[string]string, error) {
	currentUser, err := a.repository.GetByEmail(user.Email)
	if err != nil {
		return nil, err
	}
	s := os.Getenv("JWT_ACCESS_TOKEN_DURATION")

	duration, err := time.ParseDuration(s)
	if err != nil {
		return nil, err
	}

	payload := jwt.MapClaims{
		"sub": map[string]interface{}{
			"id":       currentUser.ID,
			"email":    currentUser.Email,
			"nickname": currentUser.Nickname,
		},
		"exp": time.Now().Add(duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	jwtSecret := []byte("very-test")
	t, err := token.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"token": t,
	}, nil
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
