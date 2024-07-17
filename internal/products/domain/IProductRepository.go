package domain

import "context"

type IProductRepository interface {
	GetByID(ctx context.Context, id int) (*Product, error)
	Create(ctx context.Context, product *Product) (int, error)
	Update(ctx context.Context, product *Product) error
}
