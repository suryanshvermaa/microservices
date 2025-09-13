package catalog

import (
	"context"

	"github.com/segmentio/ksuid"
)

type Service interface {
	PostProducts(ctx context.Context, name, description string, price float64) (*Product, error)
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

func (s *catalogService) PostProducts(ctx context.Context, name, description string, price float64) (*Product, error) {
	p := &Product{
		Name:        name,
		Description: description,
		Price:       price,
		ID:          ksuid.New().String(),
	}
	err := s.repository.PutProduct(ctx, *p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *catalogService) GetProductByID(ctx context.Context, id string) (*Product, error) {
	return s.repository.GetProductByID(ctx, id)
}

func (s *catalogService) ListProducts(ctx context.Context, skip uint64, take uint64) ([]Product, error) {
}

func (s *catalogService) ListProductsWithIDs(ctx context.Context, ids []string) ([]Product, error) {
}

func (s *catalogService) SearchProducts(ctx context.Context, query string) ([]Product, error) {
	return s.repo.SearchProducts(ctx, query, 0, 100)
}
