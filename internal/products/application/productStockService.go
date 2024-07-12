package application

import (
	"context"

	"github.com/NeiderFajardo/internal/products/domain"
	"github.com/NeiderFajardo/pkg/apierrors"
)

type IProductStockService interface {
	SubtractStock(ctx context.Context, id int, quantity int) error
}

type ProductStockService struct {
	productRepository domain.IProductRepository
}

func NewProductStockService(repository domain.IProductRepository) IProductStockService {
	return &ProductStockService{
		productRepository: repository,
	}
}

func (ps ProductStockService) SubtractStock(ctx context.Context, id int, quantity int) error {
	product, err := ps.productRepository.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if product.Stock < quantity {
		return apierrors.BadRequest("Not enough stock", "not_enough_stock", "stock")
	}
	resultErr := ps.productRepository.SubtractFromStock(ctx, id, quantity)
	// if resultErr is nil then we can raise and handle product updated event
	return resultErr
}
