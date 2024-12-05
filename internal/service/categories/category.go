package categories

type CategoryService struct {
	repository CategoryRepository,
}

func NewCategoryService(repository CategoryRepository) {
	return &CategoryService{
		repository: repository,
	}
}

func (s *CategoryService) GetCategories() ([]Category, error) {}

func (s *CategoryService) GetCategory(id uint64) (Category, error) {}

func (s *CategoryService) CreateCategory() error {}

func (s *CategoryService) UpdateCategory() error {}

func (s *CategoryService) DeleteCategory(id uint64) error {}
