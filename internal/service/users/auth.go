package service

import "github.com/margar-melkonyan/watch-later.git/internal/repository"

type AuthService struct {
	repository repository.UserRepository
}

func NewAuthService(repository repository.UserRepository) *AuthService {
	return &AuthService{
		repository: repository,
	}
}

func (a *AuthService) SigIn() {}

func (a *AuthService) SigUp() {}
