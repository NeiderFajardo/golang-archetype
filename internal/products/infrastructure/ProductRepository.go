package infrastructure

import "github.com/NeiderFajardo/internal/products/domain"

type ProductRepository struct {
}

func NewProductRepository() domain.IProductRepository {
	return &ProductRepository{}
}

func (pr *ProductRepository) GetByID(id int) *domain.Product {
	return &domain.Product{
		ID:   id,
		Name: "Product Name 4",
	}
}
