package application

import (
	"context"
	"fmt"

	"github.com/NeiderFajardo/internal/products/api/models"
	"github.com/NeiderFajardo/internal/products/domain"
)

type IProductService interface {
	GetByID(id int) (*domain.Product, error)
	Create(product *models.ProductRequest) (int, error)
}

type ProductService struct {
	productRepository domain.IProductRepository
}

func NewProductService(repository domain.IProductRepository) IProductService {
	return &ProductService{
		productRepository: repository,
	}
}

func (ph ProductService) GetByID(id int) (*domain.Product, error) {
	result, err := ph.productRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ph ProductService) Create(product *models.ProductRequest) (int, error) {
	if problems := product.Valid(context.Background()); len(problems) > 0 {
		return 0, fmt.Errorf("Invalid %T: %d", product, len(problems))
	}

	productToSave := domain.NewProduct(
		product.Id,
		product.Name,
		product.Description,
		product.Price,
	)
	result, err := ph.productRepository.Create(productToSave)
	if err != nil {
		return 0, err
	}
	return result, nil
}
