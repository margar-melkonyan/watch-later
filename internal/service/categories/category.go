package service

import "github.com/margar-melkonyan/watch-later.git/internal/repository"

type CategoryService struct {
	repository *repository.CategoryRepository
}

func NewCategoryService(repository *repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		repository: repository,
	}
}

func (s *CategoryService) GetCategories() ([]*repository.Category, error) {
	categories, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *CategoryService) GetCategory(id uint64) (*repository.Category, error) {
	category, err := s.repository.Get(id)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *CategoryService) CreateCategory() error {
	return nil
}

func (s *CategoryService) UpdateCategory() error {
	return nil
}

func (s *CategoryService) DeleteCategory(id uint64) error {
	return nil
}
