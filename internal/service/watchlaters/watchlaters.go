package watchlaters

type WatchLatersService struct {
	repository WatchLatersRepository
}

func NewWatchLatersService(repository WatchLatersRepository) *WatchLatersService {
	return &WatchLatersService{
		repository: repository,
	}
}

func (s *WatchLatersService) GetWatchLater(perPage uint16) ([]WatchLater, error) {}

func (s *WatchLatersService) StoreWatchLater() error {}

func (s *WatchLatersService) UpdateWatchLater(id uint64) error {}

func (s *WatchLatersService) DeleteWatchLater(id uint64) error {}

func (s *WatchLatersService) DeleteWatchLaters(id uint64) error {}
