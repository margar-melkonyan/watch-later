package service

import (
	"github.com/margar-melkonyan/watch-later.git/internal/repository"
)

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

	if categories == nil {
		categories = []*repository.Category{}
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

func (s *CategoryService) CreateCategory(category *repository.Category) error {
	return s.repository.Create(category)
}

func (s *CategoryService) UpdateCategory(category *repository.Category, id uint64) error {
	return s.repository.Update(category, id)
}

func (s *CategoryService) DeleteCategory(id uint64) error {
	return s.repository.Delete(id)
}

func (s *CategoryService) RestoreCategory(id uint64) error {
	return s.repository.Restore(id)
}
