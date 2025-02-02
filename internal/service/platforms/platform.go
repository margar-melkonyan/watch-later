package platforms

import "github.com/margar-melkonyan/watch-later.git/internal/repository"

type PlatformService struct {
	repository repository.PlatformRepository
}

func NewPlatformRepository(repository repository.PlatformRepository) *PlatformService {
	return &PlatformService{
		repository: repository,
	}
}

func (s *PlatformService) GetPlatforms() ([]*repository.Platform, error) {
	return nil, nil
}

func (s *PlatformService) GetPlatform(id uint64) (*repository.Platform, error) {
	return nil, nil
}

func (s *PlatformService) CreatePlatform(platform *repository.Platform) error {
	return nil
}

func (s *PlatformService) UpdatePlatform(id uint64, platform *repository.Platform) error {
	return nil
}

func (s *PlatformService) DeletePlatform(id uint64) error {
	return nil
}
