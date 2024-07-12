package application

import (
	"context"

	"github.com/NeiderFajardo/internal/products/api/models"
	"github.com/NeiderFajardo/internal/products/domain"
)

type IProductService interface {
	GetByID(ctx context.Context, id int) (*domain.Product, error)
	Create(ctx context.Context, product *models.ProductRequest) (int, error)
}

type ProductService struct {
	productRepository domain.IProductRepository
}

func NewProductService(repository domain.IProductRepository) IProductService {
	return &ProductService{
		productRepository: repository,
	}
}

func (ph ProductService) GetByID(ctx context.Context, id int) (*domain.Product, error) {
	result, err := ph.productRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ph ProductService) Create(ctx context.Context, product *models.ProductRequest) (int, error) {
	productToSave := domain.NewProduct(
		product.Id,
		product.Name,
		product.Description,
		product.Price,
		product.Stock,
	)
	result, err := ph.productRepository.Create(ctx, productToSave)
	if err != nil {
		return 0, err
	}
	return result, nil
}
