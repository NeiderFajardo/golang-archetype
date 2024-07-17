package domain

import "github.com/NeiderFajardo/pkg/apierrors"

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Stock       int
}

func NewProduct(id int, name, description string, price float64, stock int) *Product {
	return &Product{
		ID:          id,
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
	}
}

func (p *Product) UpdateStock(stock int) *apierrors.ApiError {
	if stock < 0 {
		return apierrors.BadRequest("Not enough stock", "not_enough_stock", "stock")
	}
	p.Stock = stock
	return nil
}
