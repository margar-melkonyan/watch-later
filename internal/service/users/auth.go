package users

type AuthService struct {
	repository UserRepository,
}

func NewAuthService (repository UserRepository) *AuthService {
	return &AuthService{
		repository: repository,
	}
}

func (a *AuthService) SigIn() {}

func (a *AuthService) SigUp() {}