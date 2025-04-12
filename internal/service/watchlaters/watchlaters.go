package service

import "github.com/margar-melkonyan/watch-later.git/internal/repository"

type WatchLaterService struct {
	repository *repository.WatchLaterRepository
}

func NewWatchLaterService(repository *repository.WatchLaterRepository) *WatchLaterService {
	return &WatchLaterService{
		repository: repository,
	}
}

func (s *WatchLaterService) GetWatchLater(id uint64) (*repository.WatchLater, error) {
	return s.repository.Get(id)
}

func (s *WatchLaterService) GetWatchLaters() ([]*repository.WatchLater, error) {
	watchLaters, err := s.repository.GetAll()
	if len(watchLaters) == 0 {
		return []*repository.WatchLater{}, err
	}

	return watchLaters, err
}

func (s *WatchLaterService) GetWatchLatersByCategory(cateogyrID uint64) ([]*repository.WatchLater, error) {
	watchLaters, err := s.repository.GetByCategory(cateogyrID)

	if len(watchLaters) == 0 {
		watchLaters = []*repository.WatchLater{}
	}

	return watchLaters, err
}

func (s *WatchLaterService) GetWatchLatersByPlatform(platformID uint64) ([]*repository.WatchLater, error) {
	watchLaters, err := s.repository.GetByPlatform(platformID)

	if len(watchLaters) == 0 {
		watchLaters = []*repository.WatchLater{}
	}

	return watchLaters, err
}

func (s *WatchLaterService) StoreWatchLater(form *repository.WatchLater) error {
	return s.repository.Create(form)
}

func (s *WatchLaterService) UpdateWatchLater(form *repository.WatchLater, id uint64) error {
	return s.repository.Update(form, id)
}

func (s *WatchLaterService) DeleteWatchLater(id uint64) error {
	return s.repository.Delete(id)
}
