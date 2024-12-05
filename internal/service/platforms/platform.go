package platforms

type PlatformService struct {
	repository PlatformRepository,
}

func NewPlatformRepository(repository PlatformRepository) *PlatformService {
	return &PlatformService {
		repository: repository,
	}
}

func (s *PlatformService) GetPlatforms() ([]Platform, error) {}

func (s *PlatformService) GetPlatform(id uint64) (Platform, error) {}

func (s *PlatformService) CreatePlatform(platform Platform) error {}

func (s *PlatformService) UpdatePlatform(id uint64, platform Platform) error {}

func (s *PlatformService) DeletePlatform(id uint64) error {}
