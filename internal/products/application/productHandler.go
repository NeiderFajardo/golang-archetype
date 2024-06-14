package application

import (
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
	return ph.productRepository.GetByID(id), nil
}

func (ph ProductService) Create(product *models.ProductRequest) (int, error) {
	return 0, nil
}
