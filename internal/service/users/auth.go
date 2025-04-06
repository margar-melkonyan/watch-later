package service

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
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

type Claims struct {
	Sub struct {
		Email    string `json:"email"`
		Nickname string `json:"nickname"`
	} `json:"sub"`
	jwt.RegisteredClaims
}

func parseToken(token string, tokenType string) (*Claims, error) {
	var claims Claims
	t, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv(fmt.Sprintf("%v_SECRET", tokenType))), nil
	})

	if claims.ExpiresAt != nil && time.Now().Unix() > claims.ExpiresAt.Unix() {
		return nil, errors.New("token is expired")
	}

	if err != nil || !t.Valid {
		return nil, errors.New("your token is invalid")
	}

	return &claims, nil
}

func CheckTokenIsNotExpired(token string, tokenType string) (*Claims, error) {
	token = strings.TrimSpace(strings.ReplaceAll(token, "Bearer ", ""))

	claims, err := parseToken(token, tokenType)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func getToken(user repository.User, tokenType string) (string, error) {
	seconds := os.Getenv(fmt.Sprintf("%v_DURATION", tokenType))
	duration, err := time.ParseDuration(seconds)

	if err != nil {
		return "", err
	}
	payload := jwt.MapClaims{
		"sub": map[string]interface{}{
			"email":    user.Email,
			"nickname": user.Nickname,
		},
		"exp": time.Now().Add(time.Duration(duration)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	jwtSecret := []byte(os.Getenv(fmt.Sprintf("%v_SECRET", tokenType)))
	t, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return t, nil
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

	if err = bcrypt.CompareHashAndPassword([]byte(currentUser.Password), []byte(user.Password)); err != nil {
		return nil, errors.New("password is not valid")
	}

	accessToken, err := getToken(*currentUser, "JWT_ACCESS_TOKEN")
	if err != nil {
		return nil, err
	}

	refreshToken, err := getToken(*currentUser, "JWT_REFRESH_TOKEN")
	if err != nil {
		return nil, err
	}

	currentUser.RefreshToken = refreshToken
	a.repository.Update(currentUser, currentUser.ID)

	return map[string]string{
		"token":         accessToken,
		"refresh_token": refreshToken,
	}, nil
}

func (a *AuthService) SignUp(user *repository.User) error {
	if _, err := a.repository.GetByEmail(user.Email); err == nil {
		return errors.New("user with this email already exists")
	}

	power, err := strconv.Atoi(os.Getenv("BCRYPT_POWER"))
	if err != nil {
		return errors.New("something went wrong")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), power)
	if err != nil {
		return err
	}
	user.Password = string(password)

	return a.repository.Create(user)
}

func (a *AuthService) CurrentUser(email string) (*common.ResponseUser, error) {
	user, err := a.repository.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	responseUser := &common.ResponseUser{
		ID:         user.ID,
		Email:      user.Email,
		Firstname:  user.Firstname,
		Lastname:   user.Lastname,
		Patronymic: user.Patronymic,
		Nickname:   user.Nickname,
	}

	return responseUser, nil
}

func (a *AuthService) RefreshToken(token string) (map[string]string, error) {
	token = strings.TrimSpace(strings.ReplaceAll(token, "Bearer ", ""))
	claims, err := parseToken(token, "JWT_REFRESH_TOKEN")
	if err != nil {
		return nil, err
	}

	user, err := a.repository.GetByEmail(claims.Sub.Email)

	if user.RefreshToken != token {
		return nil, errors.New("not valid refresh token")
	}

	if err != nil {
		return nil, err
	}

	newToken, err := getToken(*user, "JWT_ACCESS_TOKEN")
	if err != nil {
		return nil, err
	}
	newRefreshToken, err := getToken(*user, "JWT_REFRESH_TOKEN")
	if err != nil {
		return nil, err
	}

	user.RefreshToken = newRefreshToken

	a.repository.Update(user, user.ID)
	return map[string]string{
		"token":         newToken,
		"refresh_token": newRefreshToken,
	}, nil
}
