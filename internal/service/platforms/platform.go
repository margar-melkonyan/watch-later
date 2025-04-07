package service

import "github.com/margar-melkonyan/watch-later.git/internal/repository"

type PlatformService struct {
	repository *repository.PlatformRepository
}

func NewPlatformService(repository *repository.PlatformRepository) *PlatformService {
	return &PlatformService{
		repository: repository,
	}
}

func (s *PlatformService) GetPlatforms() ([]*repository.Platform, error) {
	platforms, err := s.repository.GetAll()
	if platforms == nil {
		platforms = []*repository.Platform{}
	}
	return platforms, err
}

func (s *PlatformService) GetPlatform(id uint64) (*repository.Platform, error) {
	return s.repository.Get(id)
}

func (s *PlatformService) CreatePlatform(platform *repository.Platform) error {
	return s.repository.Create(platform)
}

func (s *PlatformService) UpdatePlatform(platform *repository.Platform, id uint64) error {
	return s.repository.Update(platform, id)
}

func (s *PlatformService) DeletePlatform(id uint64) error {
	return s.repository.Delete(id)
}

func (s *PlatformService) RestorePlatform(id uint64) error {
	return s.repository.Restore(id)
}
