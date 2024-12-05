package users

type UserService struct {
	repository UserRepository
}

func NewUserRepository(repository UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) GetUser(id uint64) (User, error) {}
