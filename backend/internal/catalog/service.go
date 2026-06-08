package catalog

// Service handles catalog business logic
type Service struct {
	repo Repository
}

// NewService creates a new service
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// GetPublicCatalog returns active products
func (s *Service) GetPublicCatalog() ([]Product, error) {
	return s.repo.GetActiveProducts()
}

// GetProduct retrieves a product by ID
func (s *Service) GetProduct(id string) (*Product, error) {
	return s.repo.GetByID(id)
}
