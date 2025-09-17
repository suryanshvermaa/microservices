package order

import (
	"context"
	"time"
)

type Service interface {
	PostOrder(ctx context.Context, accountID string, products []OrderedProduct) error
	GetOrdersForAccount(ctx context.Context, accountID string) ([]Order, error)
}

type Order struct {
	ID         string
	CreatedAt  time.Time
	TotalPrice float64
	AccountID  string
	Products   []OrderedProduct
}

type OrderedProduct struct {
	ID          string
	Name        string
	Description string
	Price       float64
	Quantity    uint32
}

type orderService struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &orderService{repository: r}
}

func (s *orderService) PostOrder(ctx context.Context, accountID string, products []OrderedProduct) error {

}

func (s *orderService) GetOrdersForAccount(ctx context.Context, accountID string) ([]Order, error) {

}
