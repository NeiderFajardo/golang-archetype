package application

import (
	"context"

	"github.com/NeiderFajardo/internal/products/domain"
	"github.com/NeiderFajardo/internal/products/domain/events"
	"github.com/google/uuid"
)

type IProductStockService interface {
	SubtractStock(ctx context.Context, id int, quantity int) error
}

type ProductStockService struct {
	productRepository domain.IProductRepository
	productPublisher  events.IProductEventHandler
}

func NewProductStockService(repository domain.IProductRepository, publisher events.IProductEventHandler) IProductStockService {
	return &ProductStockService{
		productRepository: repository,
		productPublisher:  publisher,
	}
}

func (ps ProductStockService) SubtractStock(ctx context.Context, id int, quantity int) error {
	product, err := ps.productRepository.GetByID(ctx, id)
	if err != nil {
		return err
	}
	updateError := product.UpdateStock(product.Stock - quantity)
	if updateError != nil {
		return updateError
	}
	resultErr := ps.productRepository.Update(ctx, product)
	if resultErr == nil {
		go ps.productPublisher.Notify(events.NewUpdatedStockEvent(uuid.New(), product))
	}
	return resultErr
}
