package order

import (
	"context"
	"time"

	"github.com/segmentio/ksuid"
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
	o := Order{
		ID:        ksuid.New().String(),
		CreatedAt: time.Now().UTC(),
		AccountID: accountID,
		Products:  products,
	}
	var total float64
	for _, p := range products {
		total += p.Price * float64(p.Quantity)
	}
	o.TotalPrice = total
	return s.repository.PutOrder(ctx, o)
}

func (s *orderService) GetOrdersForAccount(ctx context.Context, accountID string) ([]Order, error) {

}
