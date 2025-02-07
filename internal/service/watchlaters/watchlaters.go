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

func (s *WatchLaterService) GetWatchLater(perPage uint16) ([]*repository.WatchLater, error) {
	return nil, nil
}

func (s *WatchLaterService) StoreWatchLater() error {
	return nil
}

func (s *WatchLaterService) UpdateWatchLater(id uint64) error {
	return nil
}

func (s *WatchLaterService) DeleteWatchLater(id uint64) error {
	return nil
}
