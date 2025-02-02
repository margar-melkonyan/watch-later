package users

import "github.com/margar-melkonyan/watch-later.git/internal/repository"

type UserService struct {
	repository repository.UserRepository
}

func NewUserRepository(repository repository.UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) GetUser(id uint64) (*repository.User, error) {
	return nil, nil
}
