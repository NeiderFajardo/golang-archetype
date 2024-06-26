package application

import (
	"context"
	"fmt"
	"github.com/NeiderFajardo/internal/products/domain"
	"testing"
)

type productRepositoryMock struct{}

func (prm *productRepositoryMock) GetByID(ctx context.Context, id int) (*domain.Product, error) {
	if id != 1 {
		return &domain.Product{}, fmt.Errorf("Error getting product")
	}
	return &domain.Product{
		ID:          1,
		Name:        "Product 1",
		Description: "Description 1",
		Price:       10.5,
	}, nil
}

func (prm *productRepositoryMock) Create(ctx context.Context, product *domain.Product) (int, error) {
	return 0, nil
}

func TestProductService_GetByID(t *testing.T) {
	service := NewProductService(&productRepositoryMock{})
	_, err := service.GetByID(context.Background(), 1)
	if err != nil {
		t.Errorf("Error was not expected")
	}
}

func TestProductService_GetByID_Error(t *testing.T) {
	service := NewProductService(&productRepositoryMock{})
	_, err := service.GetByID(context.Background(), 2)
	if err == nil {
		t.Errorf("Error was expected")
	}
}
