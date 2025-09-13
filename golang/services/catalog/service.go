package catalog

import "context"

type Service interface {
	PostProducts(ctx context.Context, products []Product) error
	GetProductByID(ctx context.Context, id string) (*Product, error)
	ListProducts(ctx context.Context, skip uint64, take uint64) ([]Product, error)
	ListProductsWithIDs(ctx context.Context, ids []string) ([]Product, error)
	SearchProducts(ctx context.Context, query string) ([]Product, error)
}

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type catalogService struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &catalogService{repository: r}
}

func (s *catalogService) PostProducts(ctx context.Context, products []Product) error {

}

func (s *catalogService) GetProductByID(ctx context.Context, id string) (*Product, error) {
}

func (s *catalogService) ListProducts(ctx context.Context, skip uint64, take uint64) ([]Product, error) {
}

func (s *catalogService) ListProductsWithIDs(ctx context.Context, ids []string) ([]Product, error) {
}

func (s *catalogService) SearchProducts(ctx context.Context, query string) ([]Product, error) {
	return s.repo.SearchProducts(ctx, query, 0, 100)
}
